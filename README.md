# go-Marshal-testing
```
goos: windows
goarch: amd64
Benchmark_MarshalIndent_MapToBytes-8               33123         35843 ns/op        14180 B/op        203 allocs/op
Benchmark_MarshalIndent_StructToBytes_1-8          69679         14930 ns/op         7225 B/op         33 allocs/op
Benchmark_MarshalIndent_StructToBytes_2-8          79028         14766 ns/op         8154 B/op         27 allocs/op
Benchmark_MarshalIndent_StructToBytes_3-8          90169         13239 ns/op         6954 B/op         12 allocs/op
Benchmark_Marshal_MapToBytes-8                     46640         25062 ns/op         9641 B/op        197 allocs/op
Benchmark_Marshal_StructToBytes_1-8               221845          5257 ns/op         2690 B/op         27 allocs/op
Benchmark_Marshal_StructToBytes_2-8               231732          5040 ns/op         3619 B/op         21 allocs/op
Benchmark_Marshal_StructToBytes_3-8               324344          3632 ns/op         2418 B/op          6 allocs/op
Benchmark_Unmarshal_ToStruct-8                     78730         15071 ns/op         2608 B/op         43 allocs/op
```
