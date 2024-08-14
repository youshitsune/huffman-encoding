package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Tree struct {
	left  rune
	right *Tree
	base  rune
}

func NewBase(left, right rune) *Tree {
	return &Tree{
		left:  left,
		right: nil,
		base:  right,
	}
}

func NewTree(left rune, right *Tree) *Tree {
	return &Tree{
		left:  left,
		right: right,
		base:  0,
	}
}

func (t *Tree) Next() *Tree {
	return t.right
}

func compress(x rune, tree *Tree) string {
	buf := ""
	t := tree
	for {
		if t.base != 0 {
			if x == t.left {
				buf += "0"
				break
			} else if x == t.base {
				buf += "1"
				break
			}
		} else {
			if x == t.left {
				buf += "0"
				break
			} else {
				buf += "1"
				t = t.Next()
			}
		}
	}

	return buf
}

func length_encoding(x string) string {
	var r string
	prev := x[0]
	counter := 1
	for i := 1; i < len(x); i++ {
		if prev == x[i] {
			counter++
		} else {
			r += fmt.Sprintf("%s%s", strconv.Itoa(counter), string(prev))
			prev = x[i]
			counter = 1
		}

		if i+1 == len(x) {
			r += fmt.Sprintf("%s%s", strconv.Itoa(counter), string(x[len(x)-1]))
		}
	}
	return r
}

func serialize(x string) []byte {
	r := make([]byte, 0)
	for i := range x {
		t, _ := strconv.Atoi(string(x[i]))
		r = append(r, byte(t))
	}
	return r
}

func main() {
	var str string

	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		str = scanner.Text()
	}

	strings.ToLower(str)

	freq := make(map[rune]int, 0)
	keys := make([][]int, 0)

	for i := range str {
		freq[rune(str[i])]++
	}

	for k, v := range freq {
		keys = append(keys, []int{int(k), v})
	}

	sort.SliceStable(keys, func(i, j int) bool {
		if keys[i][1] == keys[j][1] {
			return keys[i][0] < keys[j][0]
		}
		return keys[i][1] < keys[j][1]
	})

	var t *Tree
	for i := 0; i < len(keys)-1; i++ {
		if i < 1 {
			t = NewBase(rune(keys[1][0]), rune(keys[0][0]))
		} else {
			t = NewTree(rune(keys[i+1][0]), t)
		}
	}

	var cstr string

	for i := range str {
		cstr += compress(rune(str[i]), t)
	}

	fmt.Println(cstr)
	fmt.Println(length_encoding(cstr))
	fmt.Println(serialize(length_encoding(cstr)))
}
