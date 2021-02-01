# go-Marshal-testing
```
Benchmark_MarshalIndent_MapToBytes-8            32977       35099 ns/op      14179 B/op      203 allocs/op
Benchmark_Marshal_MapToBytes-8                  47414       24888 ns/op       9641 B/op      197 allocs/op
Benchmark_MarshalIndent_StructToBytes_1-8       78408       14960 ns/op       7226 B/op       33 allocs/op
Benchmark_MarshalIndent_StructToBytes_2-8       81266       14714 ns/op       8155 B/op       27 allocs/op
Benchmark_MarshalIndent_StructToBytes_3-8       90296       13028 ns/op       6954 B/op       12 allocs/op
Benchmark_Marshal_StructToBytes_3-8            352725        3484 ns/op       2418 B/op        6 allocs/op
Benchmark_Unmarshal_ToStruct-8                  80941       14640 ns/op       2608 B/op       43 allocs/op
Benchmark_Unmarshal_ToMap-8                     82516       14811 ns/op       5112 B/op      150 allocs/op
```
