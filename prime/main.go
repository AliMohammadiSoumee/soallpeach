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

	chunk := num / 8
	turn := 0
	for i := 0; i < num; i += chunk {
		turn++
		go goroutine(i, chunk, end)
	}

	for turn > 0 {
		turn--
		<-end
	}

	ans = ans[:num*2]
	os.Stdout.Write(ans)
}

func goroutine(start, chunk int, end chan string) {
	e := chunk + start
	j := start * 2
	for ind := start; ind < e && ind < num; ind++ {
		n := lines[ind]
		if n < 1000000 {
			if mark[n] {
				ans[j] = 48
				ans[j+1] = 10
			} else {
				ans[j] = 49
				ans[j+1] = 10
			}
		} else {
			for _, p := range primes {
				if p*p > n {
					ans[j] = 49
					ans[j+1] = 10
				} else if n%p == 0 {
					ans[j] = 48
					ans[j+1] = 10
				}
			}
		}
		j += 2
	}
	end <- "e"
}
