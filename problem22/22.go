package problem22

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
)

type lexer struct {
	name  chan string
	input string
	point int
}

type lexFn func(*lexer) lexFn

func lexNames(n chan string, input string) {
	l := &lexer{name: n, input: input}
	for s := lexStart; s != nil; {
		s = s(l)
	}
	close(l.name)
}

func lexStart(l *lexer) lexFn {
	if l.point >= len(l.input) {
		return nil
	}
	if l.input[l.point] == '"' {
		l.point++
		return lexWord
	}
	l.point++
	return lexStart
}

func lexWord(l *lexer) lexFn {
	var buff []byte
	for l.input[l.point] != '"' {
		buff = append(buff, l.input[l.point])
		l.point++
	}
	l.point++
	l.name <- string(buff)
	return lexStart
}

// broken ):
// func brokenSort(things []string) {
// 	var recurse func([]string, int, int)
// 	recurse = func(a []string, low, high int) {
// 		if low >= high {
// 			return
// 		}
// 		mid := (low + high) / 2
// 		pivot := things[mid]
// 		a[mid], a[high] = a[high], a[mid]
// 		store := low
// 		for i := low; i < high-1; i++ {
// 			if a[i] < pivot {
// 				a[i], a[store] = a[store], a[i]
// 				store++
// 			}
// 		}
// 		a[store], a[high] = a[high], a[store]
// 		recurse(a, low, store-1)
// 		recurse(a, store+1, high)
// 	}
// 	recurse(things, 0, len(things)-1)
// }

func val(str string) int {
	var sum int
	for _, c := range str {
		sum += int(c) - 'A' + 1
	}
	return sum
}

func Solve() string {
	c := make(chan string)
	in, err := ioutil.ReadFile("./problem22/p022_names.txt")
	if err != nil {
		log.Fatal(err)
	}
	go lexNames(c, string(in))
	names := []string{}
	for n := range c {
		names = append(names, n)
	}
	sort.Strings(names)
	var sum int
	for i, name := range names {
		sum += (i + 1) * val(name)
	}
	return fmt.Sprint(sum)

}
