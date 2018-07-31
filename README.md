# ORM Benchmark

A benchmark to compare the performance of golang orm package.

## Results (2014-1-7)

### Environment

* Aliyun Intel Xeon E5-2630 2.30GHz (4 core)
* 8G RAM
* CentOS 6.5
* go version go1.2 linux/amd64
* [Go-MySQL-Driver Latest](https://github.com/go-sql-driver/mysql)

### MySQL

* MySQL 5.5.34 for Linux on x86_64
* MySQL-server-5.5.34-1.rhel5.x86_64.rpm
* Config in my.cnf

### ORMs

All package run in no-cache mode.

* [Beego ORM](http://beego.me/docs/mvc/model/overview.md) latest in branch [develop](https://github.com/astaxie/beego/tree/develop)
* [xorm](https://github.com/lunny/xorm) latest
* [Hood](https://github.com/eaigner/hood) latest
* [Qbs](https://github.com/coocood/qbs) latest (Disabled stmt cache / [patch](https://gist.github.com/slene/8297019) / [full](https://gist.github.com/slene/8297565))

### Run

```go
go get github.com/beego/orm-benchmark
orm-benchmark -multi=20 -orm=all
```

### Reports

#### Sample 1

```
40000 times - Insert
       raw:     7.71s       192768 ns/op     552 B/op     12 allocs/op
      hood:    11.78s       294520 ns/op   13039 B/op    316 allocs/op
       orm:    12.13s       303365 ns/op    1937 B/op     40 allocs/op
    gendry:    12.54s       313596 ns/op    1953 B/op     36 allocs/op
      gorm:    16.48s       411881 ns/op    6142 B/op    124 allocs/op

 10000 times - MultiInsert 100 row
       orm:    15.55s      1555473 ns/op  147150 B/op   1534 allocs/op
       raw:    16.73s      1673108 ns/op  110803 B/op    811 allocs/op
    gendry:    17.22s      1722363 ns/op  148934 B/op    943 allocs/op
      gorm:     Not support multi insert
      hood:     Not support multi insert

 40000 times - Update
       raw:    10.51s       262697 ns/op     616 B/op     14 allocs/op
       orm:    12.80s       319930 ns/op    1929 B/op     40 allocs/op
    gendry:    15.17s       379272 ns/op    3026 B/op     62 allocs/op
      hood:    15.93s       398332 ns/op   13039 B/op    316 allocs/op
      gorm:    35.66s       891479 ns/op   15646 B/op    335 allocs/op

 80000 times - Read
       raw:     9.87s       123378 ns/op    1464 B/op     40 allocs/op
      hood:    15.86s       198311 ns/op    4418 B/op     93 allocs/op
       orm:    16.99s       212386 ns/op    2837 B/op    100 allocs/op
    gendry:    17.56s       219441 ns/op    4266 B/op    100 allocs/op
      gorm:    20.12s       251512 ns/op    6482 B/op    147 allocs/op

 40000 times - MultiRead limit 100
       raw:    14.77s       369221 ns/op   34848 B/op   1324 allocs/op
       orm:    26.50s       662504 ns/op   85269 B/op   4291 allocs/op
    gendry:    32.90s       822476 ns/op  141368 B/op   3574 allocs/op
      hood:    48.69s      1217279 ns/op  234267 B/op   9611 allocs/op
      gorm:    50.87s      1271681 ns/op  247494 B/op   6210 allocs/op
```


### Contact

Maintain by [nlimpid](https://github.com/nlimpid)