package main

import (
	"io/ioutil"
	"os"
)

var mark = make([]bool, 1000000)
var ans = make([]byte, 3000000)
var primes = make([]int, 78498)
var lines = make([]int, 1500000)
var num int = 0

func gharbal(end chan string) {
	mark[0] = true
	mark[1] = true
	p := 0
	for i := 2; i < 1000000; i++ {
		if mark[i] {
			continue
		}
		for j := i * 2; j < 1000000; j += i {
			mark[j] = true
		}
		primes[p] = i
		p++
	}
	end <- "e"
}

func read(path string, end chan string) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	number := 0
	l := 0
	for _, b := range bytes {
		b -= 48
		if b == 218 {
			lines[l] = number
			number = 0
			l++
			continue
		}
		number *= 10
		number += int(b)
	}
	lines[l] = number
	l++
	num = l
	end <- "e"
}

func whitespace(end chan string) {
	for i := 1; i < 3000000; i += 2 {
		ans[i] = 10
	}
	end <- "e"
}

func main() {
	path := os.Args[1]

	end := make(chan string, 10)
	go read(path, end)
	go gharbal(end)
	go whitespace(end)

	<-end
	<-end
	<-end

	ind := 0
	for _, i := range lines {
		if i < 1000000 {
			if mark[i] {
				ans[ind] = 48
			} else {
				ans[ind] = 49
			}
		} else {
			ans[ind] = isPrime(i)
		}
		ind += 2
	}
	ans = ans[:num*2]
	os.Stdout.Write(ans)
}

func isPrime(n int) byte {
	for _, p := range primes {
		if p*p > n {
			return 49
		} else if n%p == 0 {
			return 48
		}
	}
	return 49
}
