package main

import(
	"os"
	"bufio"
	"fmt"
	"strconv"
)

func main(){
	var output = 0
	file, _ := os.Open("input.txt")
    fscanner := bufio.NewScanner(file)
    for fscanner.Scan() {
		as_int, _ := strconv.Atoi(fscanner.Text())
		output += as_int
	}
	fmt.Println(output)
}