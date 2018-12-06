package main

import(
	"os"
	"bufio"
	"fmt"
)

func main(){

	var TwoCount = 0
	var ThreeCount = 0
	
	file, _ := os.Open("input.txt")
    fscanner := bufio.NewScanner(file)
    for fscanner.Scan() {
		m := make(map[rune]int)
		line := fscanner.Text()
		for _, char := range line {
			if _, ok := m[char]; ok {
				m[char]++
			}else{
				m[char]++
			}
		}
		is_a_two := false;
		is_a_three := false;
		for _, val := range m{
			if (val==2){
				is_a_two = true
			} else if(val==3){
				is_a_three = true
			}
		}
		if is_a_two {
			TwoCount++
		}
		if is_a_three {
			ThreeCount++
		}
	}
	fmt.Println(TwoCount * ThreeCount)
}