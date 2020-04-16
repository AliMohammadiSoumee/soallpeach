package main

import (
	"io/ioutil"
	"os"
)

var mark = make([]bool, 1000000)
var ans = make([]uint32, 1500000)
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

	for i := range bytes {
		bytes[i] -= 48
	}

	number := 0
	l := 0
	for _, b := range bytes {
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

func main() {
	path := os.Args[1]

	end := make(chan string, 10)
	go read(path, end)
	go gharbal(end)
	<-end
	<-end

	chunk := num / 10
	for i := 0; i < num; i += chunk {
		go goroutine(i, chunk, end)
	}

	for i := 0; i < 10; i++ {
		<-end
	}

	ans = ans[:num]
	answer := join(ans)
	os.Stdout.Write(answer)
}
func goroutine(start, chunk int, end chan string) {
	e := chunk + start
	for ind := start; ind < e && ind < num; ind++ {
		n := lines[ind]
		isPrime(n, ind)
	}
	end <- "e"
}

func isPrime(n, ind int) {
	if n < 1000000 {
		if mark[n] {
			ans[ind] = 0
		} else {
			ans[ind] = 1
		}
	} else {
		for _, p := range primes {
			if p >= 46340 || p*p > n {
				ans[ind] = 1
			} else if n%p == 0 {
				ans[ind] = 0
			}
		}
	}
}

func join(elems []uint32) []byte {
	j := 0
	bs := make([]byte, len(elems)*2)
	for _, i := range elems {
		if i == 0 {
			bs[j] = 48
		} else {
			bs[j] = 49
		}
		bs[j+1] = 10
		j += 2
	}
	return bs
}
