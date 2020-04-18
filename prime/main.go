package main

import (
	"io/ioutil"
	"log"
	"os"
	"time"
)

const bigestPrime = 100000

var mark = make([]byte, bigestPrime)
var ans = make([]byte, 3000000)
var primes = make([]int, 10000)
var num int

func gharbal() {
	for i := 0; i < bigestPrime; i++ {
		mark[i] = '0'
	}
	mark[0] = '1'
	mark[1] = '1'
	p := 0
	for i := 2; i < bigestPrime; i++ {
		if mark[i] == '1' {
			continue
		}
		for j := i * 2; j < bigestPrime; j += i {
			mark[j] = '1'
		}
		primes[p] = i
		p++
	}
}

func read(path string, end chan string) {
	t1 := time.Now()
	bytes, _ := ioutil.ReadFile(path)
	t2 := time.Now()

	log.Println("--->", t2.Sub(t1))

	n := 0
	for _, b := range bytes {
		if b != '\n' {
			n = n*10 + int(b) - '0'
			continue
		}
		if n < bigestPrime {
			ans[num] = mark[n]
		} else {
			ans[num] = isPrime(n)
		}
		n = 0
		num += 2
	}

	if n < bigestPrime {
		ans[num] = mark[n]
	} else {
		ans[num] = isPrime(n)
	}
	end <- "e"
}

func main() {
	path := os.Args[1]

	t0 := time.Now()
	gharbal()
	t1 := time.Now()
	end := make(chan string, 10)
	go read(path, end)

	for i := 1; i < 3000000; i += 2 {
		ans[i] = 10
	}
	<-end
	t2 := time.Now()

	ans = ans[:num]
	os.Stdout.Write(ans)
	t3 := time.Now()
	log.Println(t1.Sub(t0))
	log.Println(t2.Sub(t1))
	log.Println(t3.Sub(t2))
}

func isPrime(n int) byte {
	for _, p := range primes {
		if p*p > n {
			return '1'
		} else if n%p == 0 {
			return '0'
		}
	}
	return '1'
}
