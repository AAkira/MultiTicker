TARGETS=$(shell go list ./...)

test: vet
	go test -cover $(TARGETS)

vet:
	go vet $(TARGETS)