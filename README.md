## Usage 
```
cd cmd/cli
# Run the program without any parameter
# By default it uses sample-large.csv and search for app Id 374
go run main.go 
# Run the program with CSV file path and app Id
go run main.go -file sample-small.csv 333
```
## Benchmarks
```
pkg: flexrea/cmd/cli
BenchmarkSampleLargeFile-8             1        7709017754 ns/op        2829706752 B/op 44033414 allocs/op
BenchmarkSampleSmallFile-8            20          71033827 ns/op        28266340 B/op     440647 allocs/op
```



