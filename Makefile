GOTEST_FLAGS=-cpu=4 -benchmem -benchtime=5s

benchmark:
	go test $(GOTEST_FLAGS) -bench "^Benchmark"
