package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	pwd, _ := os.Getwd()
	f, err := os.Open(pwd + "\\2024\\day01\\input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	var list1 []int
	var list2 []int

	var total int

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, "   ")
		word1, _ := strconv.Atoi(words[0])
		word2, _ := strconv.Atoi(words[1])

		list1 = append(list1, word1)
		list2 = append(list2, word2)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	for i := range len(list1){
		if list1[i] > list2[i] {
			total += list1[i] - list2[i]
		} else {
			total += list2[i] - list1[i]
		}
	}

	fmt.Println(total)
}



