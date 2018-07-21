package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("names.txt") // read the file
	if err != nil {
		panic(err)
	}

	str := strings.Split(strings.ToLower(string(file)), ",") // parser
	// fmt.Println(len(str))
	sort.Strings(str)
	fmt.Println(scores(str))

}

func scores(str []string) int {
	var sum, sumname int
	for i := 0; i < len(str); i++ {
		for j := 1; j < (len(str[i]) - 1); j++ {
			sumname += int(str[i][j]) - 'a' + 1
		}
		fmt.Println(i, sumname, str[i])
		sum += sumname * (i + 1)
		sumname = 0
	}
	return sum
}
