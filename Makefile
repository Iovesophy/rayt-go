.PHONY: build
build:
	go build main.go

.PHONY: run
run:
	go run main.go

.PHONY: profile
profile:
	cd cmd/profile && \
	time go run main.go && \
	go tool pprof -png assets/cpu.pprof > assets/out.png
