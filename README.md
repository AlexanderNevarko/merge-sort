# Parallel merge-sort

Parallel merge-sort algorithm on Go

To run you can either build or just use go's fast build and run option:
```
$ go build main.go
$ ./main
```
or
```
$ go run main.go
```

You can also specify some flags:
* -np: number of threads ("workers") to parallel
* -s: size of the array, using which to test the program
* -h: to see the information about the flags

For example:
```
./main -np=10 -s=250000
```

