# go-sample

## buffer_bench
goでのバッファリングについてシステムコールのチューニングをした
Macのdtraceで確かめました。

```
go build buffer_bench.go
sudo dtruss -c ./buffer_bench
```