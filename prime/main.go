package main

import (
	"io/ioutil"
	"os"
)

const bigestPrime = 100000

var mark = make([]byte, bigestPrime)
var ans = make([]byte, 3000000)
var primes = make([]int, 10000)
var num int

func gharbal() {
	for i := 0; i < bigestPrime; i++ {
		mark[i] = '1'
	}
	mark[0] = '0'
	mark[1] = '0'
	p := 0
	for i := 2; i < bigestPrime; i++ {
		if mark[i] == '0' {
			continue
		}
		for j := i * 2; j < bigestPrime; j += i {
			mark[j] = '0'
		}
		primes[p] = i
		p++
	}
}

func read(path string, end chan string) {
	bytes, _ := ioutil.ReadFile(path)

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
	num+=2
	end <- "e"
}

func main() {
	path := os.Args[1]

	gharbal()
	end := make(chan string, 10)
	go read(path, end)

	for i := 1; i < 3000000; i += 2 {
		ans[i] = 10
	}
	<-end

	ans = ans[:num]
	os.Stdout.Write(ans)
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
