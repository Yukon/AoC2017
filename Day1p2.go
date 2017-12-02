package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("Enter captcha: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	captcha := make([]int, len(input)-1)
	for i, elm := range input[:len(input)-1] {
		str := string(elm)
		num, _ := strconv.ParseInt(str, 10, 8)
		captcha[i] = int(num)
	}

	result := 0
	for i, elm := range captcha {
		if elm == captcha[(i+len(captcha)/2)%len(captcha)] {
			result += elm
		}
	}

	fmt.Printf("%d\n", result)
}
