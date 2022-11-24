package test

import (
	"fmt"
)

func PrintTestsMessages(msgs []string) {
	for _, m := range msgs {
		fmt.Println("	" + m)
	}
}
