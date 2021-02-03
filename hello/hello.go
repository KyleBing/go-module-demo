package main

import (
	"fmt"
	"kylebing.cn/greetings"
	"log"
)

func main() {
	/*	msg, err := greetings.Hello("Kyle")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(msg)
	*/

	names := []string{"Kyle", "Tina", "Lucy"}

	msgs, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msgs)
}
