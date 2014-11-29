package problem59

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var currPass = "aaa"

func getPass() (s string, done bool) {
	p := currPass
	currPass = rollPass(currPass)
	if currPass == "aaa" {
		return "", true
	}
	return p, false
}

func rollPass(s string) string {
	if len(s) == 0 {
		return ""
	}
	if s[0] == 'z' {
		return "a" + rollPass(s[1:])
	}
	return string(byte(s[0])+1) + s[1:]
}

func decrypt(text string, pass string) string {
	store := make([]byte, 0, len(text))
	var i int
	for i < len(text) {
		grab := len(text) - i
		if grab > len(pass) {
			grab = len(pass)
		}
		store = append(store, decryptPart(text, pass, i, grab)...)
		i += grab
	}
	return string(store)
}

func decryptPart(text string, pass string, index int, grab int) []byte {
	store := make([]byte, 0, grab)
	for i := 0; i < grab; i++ {
		store = append(store, byte(text[index+i])^byte(pass[i]))
	}
	return store
}

func Solve() string {
	f, err := os.Open("problem59/p059_cipher.txt")
	if err != nil {
		log.Fatal(err)
	}
	c := csv.NewReader(f)
	strs, err := c.Read()
	if err != nil {
		log.Fatal(err)
	}
	tmp := make([]byte, 0, len(strs))
	for _, s := range strs {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		tmp = append(tmp, byte(i))
	}
	v := string(tmp)
	for pass, done := getPass(); !done; pass, done = getPass() {
		t := decrypt(v, pass)
		if strings.Contains(t, "Gospel") { // hehe
			sum := 0
			for _, b := range t {
				sum += int(b)
			}
			return fmt.Sprintf("%d", sum)
		}
	}
	panic("unreachable")
	return ""
}
