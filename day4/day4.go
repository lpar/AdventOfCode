package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

const Input = "bgvyzdsv"

func Md5String(stdata string) string {
	data := []byte(stdata)
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}

func Find(zeros string) int {
	n := 1
	for {
		data := fmt.Sprintf("%s%d", Input, n)
		hash := Md5String(data)
		fmt.Printf("%d %s -> %s\n", n, data, hash)
		if strings.HasPrefix(hash, zeros) {
			fmt.Printf("\nSuccess! n = %d\n", n)
			return n
		}
		n++
	}
}

func main() {
	fmt.Println(Md5String("abcdef609043"))
	part1 := Find("00000")
	part2 := Find("000000")
	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1, part2)
}
