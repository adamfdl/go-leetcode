package main

import (
	"bytes"
)

func main() {
	defangIPaddr("1.2.3.4")

}

func defangIPaddr(address string) string {
	defanged := &bytes.Buffer{}

	for i := 0; i < len(address); i++ {
		temp := string(address[i])
		if temp == "." {
			defanged.WriteString("[.]")
		} else {
			defanged.WriteString(temp)
		}
	}

	return defanged.String()
}
