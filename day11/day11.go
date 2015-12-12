package main

import (
	"fmt"
	"strings"
)

const PasswordLength = 8

type PasswordData [PasswordLength]int

const Allowed = "abcdefghjkmnpqrstuvwxyz"
const Base = len(Allowed)

func StringToPasswordData(pw string) (result PasswordData) {
	i := 0
	for _, c := range pw {
		result[i] = strings.IndexRune(Allowed, c)
		i++
	}
	return result
}

func PasswordDataToString(ppw *PasswordData) string {
	pw := *ppw
	result := make([]byte, PasswordLength)
	for i, c := range pw {
		result[i] = Allowed[c]
	}
	return string(result)
}

// Increment increments the supplied password data in place
func Increment(ppw *PasswordData) {
	carry := 1
	for i := PasswordLength - 1; i >= 0 && carry != 0; i-- {
		ppw[i]++
		if ppw[i] >= Base {
			ppw[i] -= Base
			carry = 1
		} else {
			carry = 0
		}
	}
}

func IsSecure(ppw *PasswordData) bool {
	pw := *ppw
	var straight bool
	var pairs int
	var lastpair int = -1
	for i, _ := range pw {
		if i > 2 {
			// Check for a straight
			if pw[i]-pw[i-1] == 1 && pw[i-1]-pw[i-2] == 1 {
				straight = true
			}
		}
		// Check for pair
		if i > 1 {
			if pw[i] == pw[i-1] && i-lastpair > 2 {
				pairs++
				lastpair = i
			}
		}
	}
	return pairs > 1 && straight
}

func Test(pw string) {
	fmt.Println(pw)
	pwd := StringToPasswordData(pw)
	fmt.Println(IsSecure(&pwd))
}

func NextPassword(ppw *PasswordData) {
	for {
		Increment(ppw)
		if IsSecure(ppw) {
			fmt.Println("Secure:   " + PasswordDataToString(ppw))
			break
		}
		fmt.Println("Insecure: " + PasswordDataToString(ppw))
	}
}

func main() {

	pw := StringToPasswordData("hxbxwxba")
	NextPassword(&pw)
	NextPassword(&pw)

}
