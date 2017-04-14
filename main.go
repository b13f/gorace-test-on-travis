package main

import (
	"os"
	"runtime"
	"time"
)

type s struct {
	m map[int]int
}

func main() {
	println(runtime.GOMAXPROCS(0))

	time.AfterFunc(10*time.Second, func() {
		os.Exit(0)
	})

	m := map[int]int{
		0: 0, 1: 1, 2: 3, 4: 5,
	}

	var sss, sss2 s
	sss.m = m
	sss2 = sss
	for i := 0; i < 3; i++ {
		go func() {
			for {
				for a, b := range sss2.m {
					a = a + b
					runtime.Gosched()
				}
			}
		}()
	}
	for {
		m[0] = m[0] + 1
		runtime.Gosched()
	}
}
