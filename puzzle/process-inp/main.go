package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s string
	fmt.Scanln(&s)
	fmt.Println(s)

	reader := bufio.NewReader(os.Stdin)
	// n, err := reader.ReadString('\n')
	// if n == "" || err != nil {
	// 	return
	// }
	// fmt.Println(n)

	bytes, _, _ := reader.ReadLine() //处理多行时，注意一下返回值的处理
	str := string(bytes)
	fmt.Println(bytes)
	fmt.Println(str)
}
