package main

import(
	"os"
	"bufio"
	"fmt"
	"strconv"
)

func main(){
	var finished = false
	var output = 0
	m := make(map[int]bool)
	for (finished == false){
		file, _ := os.Open("input.txt")
		fscanner := bufio.NewScanner(file)
		for fscanner.Scan() {
			as_int, _ := strconv.Atoi(fscanner.Text())
			output += as_int
			if _, ok := m[output]; ok {
				//do something here
				finished = true
				break
			}else{
				m[output]=true
			}
		}
	}
	fmt.Println(output)
}