# Multiple ticker

This library is the efficiently multiple ticker like `time.Ticker`.

```
                                                        |-----> [A, time.Time] /2s
[A:2, B:4, C:6] Interval List => ⏱MultipleTicker => ---|-----> [B, time.Time] /4s
                                                        |-----> [C, time.TIme] /6s
```

For example, you set 2, 4 and 6.  
Ticker interval is 2s(gcd).  
`ticker.C` is received `string` key and `time.Time`.

## Usage

### Install

`$ go get github.com/aakira/multiticker`

### How to use

Set the map with key and time.Duration.

`map[string]time.Duration` (key: string, intervalSeconds: time.Duration)

```go

interval := map[string]time.Duration{"a": 2 * time.Second, "b": 4 * time.Second, "c": 6 * time.Second}
ticker := multiticker.NewMultiTicker(interval)

// Must call this stop function
defer ticker.Stop()

for c := range ticker.C {
    switch c.Key {
    case "a": // receive per 2s
        fmt.Printf("receive: %s, time: %v\n", c.Key, c.Tick)
    case "b": // receive per 4s
        fmt.Printf("receive: %s, time: %v\n", c.Key, c.Tick)
    case "c": // receive per 6s
        fmt.Printf("receive: %s, time: %v\n", c.Key, c.Tick)
    default:
        fmt.Printf("receive: %s, time: %v\n", c.Key, c.Tick)
        return
    }
}

```

## License

```
Copyright (C) 2018 A.Akira

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
