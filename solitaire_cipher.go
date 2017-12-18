package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

const SIZE = 5

func split_to_fix_size(input string, padding string) []string {
	strlen := len(input)
	group_num := 1

	if strlen/SIZE > 1 {
		group_num = strlen / SIZE
	}
	result := make([]string, group_num)
	// if less than 5, padding first group
	if strlen < SIZE {
		result[0] = input + strings.Repeat("X", SIZE-strlen)
	} else { // padding last group
		index := 0
		for i := 0; i < group_num-1; i++ {
			index = i * SIZE
			result[i] = input[index : index+SIZE]
			fmt.Println(result[i])
		}
		index = (group_num - 1) * SIZE
		result[group_num-1] = input[index : strlen-1]
	}
	return result
}

func main() {
	//read from stdin
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	// discard non A-Z character
	valid_char := regexp.MustCompile(`[a-zA-Z]`)
	texts := valid_char.FindAllString(text, -1)
	if texts == nil {
		log.Fatal("no a-z character found")
	}
	text = strings.Join(texts, "")
	// to upper string
	text = strings.ToUpper(text)
	// split string to 5-char group
	result := split_to_fix_size(text, "X")
	fmt.Println(result)
}
