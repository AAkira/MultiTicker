package multiticker

import (
	"time"
)

type MultiTicker struct {
	C       <-chan Tick // send label and time
	c       chan Tick
	items   []tickItem
	closeCh chan struct{}
}

func NewMultiTicker(intervalSecond map[string]time.Duration) *MultiTicker {
	g := 1
	m := 0
	for _, v := range intervalSecond {
		s := int(v.Seconds())
		g = gcd(g, s)
		m = max(m, s)
	}

	items := make([]tickItem, 0, len(intervalSecond))
	for k, v := range intervalSecond {
		items = append(items, tickItem{key: k, gcdPoint: int(v.Seconds()) / g})
	}
	c := make(chan Tick, 1)

	t := &MultiTicker{
		C:       c,
		c:       c,
		items:   items,
		closeCh: make(chan struct{}),
	}
	go t.start(g, m)

	return t
}

func (m *MultiTicker) start(gcd int, max int) {
	defer close(m.c)

	ticker := time.NewTicker(time.Duration(gcd) * time.Second)
	defer ticker.Stop()

	loopCount := 1
	resetCount := max / gcd
	for {
		select {
		case <-m.closeCh:
			return
		case t := <-ticker.C:
			for _, v := range m.items {
				if loopCount%v.gcdPoint != 0 {
					continue
				}

				// non-blocking send
				select {
				case m.c <- Tick{v.key, t}:
					time.Sleep(1 * time.Microsecond)
				default:
				}
			}

			if loopCount == resetCount {
				loopCount = 1
			} else {
				loopCount++
			}
		}
	}
}

// call this function only once
func (m *MultiTicker) Stop() {
	close(m.closeCh)
}

type Tick struct {
	Key  string
	Tick time.Time
}

type tickItem struct {
	key      string
	gcdPoint int
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
