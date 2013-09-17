package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
	fmt.Println("Testing Loop...")
	for i := 0; i < 10; i++{
		fmt.Println("Loop:", i)
	}
	fmt.Println("END Loop")
	var i string
	for {
		fmt.Printf("Please enter the password: ")
		fmt.Scanf("%s", &i)
		if i == "WfP9Jtb28xxHwRxt"{
			break
		} else {
			fmt.Println("Invalid Password")
		}
	}
	fmt.Println("Identity Verified")
}
