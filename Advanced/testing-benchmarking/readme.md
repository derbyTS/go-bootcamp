##File name must \*\_test.go

##To run

```bash
go run test file
```

```bash
go test -bench=. testing_benchmarking_test.go|grep -v 'cpu'
```

```bash
go test -bench=. -benchmem testing_benchmarking_test.go|grep -v 'cpu'
```

##Profiling

```bash
go test -bench=. -memprofile mem.pprof testing_benchmarking_test.go|grep -v 'cpu'
```

##After mem.pprof is generated we can go tool

```bash
╰─ go tool pprof mem.pprof                                                                                           ─╯
File: main.test
Type: alloc_space
Time: 2026-08-09 03:17:35 PST
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 2.44GB, 99.90% of 2.44GB total
Dropped 20 nodes (cum <= 0.01GB)
      flat  flat%   sum%        cum   cum%
    2.44GB 99.90% 99.90%     2.44GB 99.90%  command-line-arguments.GeneratingRandomSlice (inline)
         0     0% 99.90%     2.44GB 99.90%  command-line-arguments.BenchmarkGenerateRandomSlice
         0     0% 99.90%     2.44GB 99.90%  testing.(*B).launch
         0     0% 99.90%     2.44GB 99.90%  testing.(*B).runN
(pprof) list GenerateRandomSlice
Total: 2.44GB
ROUTINE ======================== command-line-arguments.BenchmarkGenerateRandomSlice in /Users/User/Development/Gogogo/
go-bootcamp/Advanced/testing-benchmarking/testing_benchmarking_test.go
         0     2.44GB (flat, cum) 99.90% of Total
         .          .     36:func BenchmarkGenerateRandomSlice(b *testing.B) {
         .          .     37:   for range b.N {
         .     2.44GB     38:           GeneratingRandomSlice(1000)
         .          .     39:   }
         .          .     40:}
         .          .     41:
         .          .     42:func BenchmarkSumOfSlice(b *testing.B) {
         .          .     43:   slice := GeneratingRandomSlice(1000)

```
