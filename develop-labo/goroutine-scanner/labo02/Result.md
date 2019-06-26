
darwin 
:Core 4
:Waittime 500 milisec
:res
```

runtime.CPUS : 4
548  Port Opened
88  Port Opened
445  Port Opened
5900  Port Opened
6258  Port Opened
6263  Port Opened
finish
go run PortScanner_check_CPU.go  5.90s user 14.84s system 72% cpu 28.777 total

```

Centos
:Core 8
:waittime 500 milisec
:res
```

$ time ./PortScan_Cents
runtime.CPUS : 8
25  Port Opened
111  Port Opened
80  Port Opened
22  Port Opened
3306  Port Opened
10050  Port Opened
40989  Port Opened
42642  Port Opened
finish

real	0m0.843s
user	0m3.250s
sys	0m2.500s
$ time ./PortScan_Cents
runtime.CPUS : 8
25  Port Opened
22  Port Opened
111  Port Opened
80  Port Opened
32996  Port Opened
37800  Port Opened
40989  Port Opened
finish

real	0m0.849s
user	0m3.370s
sys	0m2.492s

```

darwin
:Core 4
:waittime 1000 milisec
:res

```

$  time go run PortScanner_check_CPU.go
runtime.CPUS : 4
445  Port Opened
88  Port Opened
548  Port Opened
5900  Port Opened
6258  Port Opened
6263  Port Opened
finish
go run PortScanner_check_CPU.go  5.88s user 13.96s system 35% cpu 55.253 total

```


Centos 
:core 8
:waittime 1000 milisec
:res

```

$ time ./PortScan_Cents
runtime.CPUS : 8
22  Port Opened
25  Port Opened
111  Port Opened
80  Port Opened
3306  Port Opened
42642  Port Opened
finish

real	0m0.900s
user	0m3.316s
sys	0m2.793s
$ time ./PortScan_Cents
runtime.CPUS : 8
22  Port Opened
25  Port Opened
80  Port Opened
111  Port Opened
443  Port Opened
3306  Port Opened
10050  Port Opened
finish

real	0m0.869s
user	0m3.391s
sys	0m2.629s

```

darwin
:Core 4
:waittime 1500 milisec
:res

```

$ time bin/PortScan
runtime.CPUS : 4
445  Port Opened
88  Port Opened
548  Port Opened
5900  Port Opened
6263  Port Opened
6258  Port Opened
finish
bin/PortScan  5.70s user 13.82s system 24% cpu 1:18.87 total
$  time bin/PortScan
runtime.CPUS : 4
88  Port Opened
548  Port Opened
445  Port Opened
5900  Port Opened
6258  Port Opened
6263  Port Opened
finish
bin/PortScan  5.51s user 13.29s system 23% cpu 1:19.78 total

```

Centos
:Core 4
:waittime 1500 milisec
:res

```

$ time PortScan_Cents
runtime.CPUS : 8
25  Port Opened
22  Port Opened
80  Port Opened
111  Port Opened
443  Port Opened
37800  Port Opened
40989  Port Opened
finish

real	0m0.873s
user	0m3.307s
sys	0m2.631s

```
