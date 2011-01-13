package main

import(
	"fmt"
	"bufio"
	"os"
	"io/ioutil"
	"regexp"
)

func main() {
	ch := make(chan string)

	for {
		// Prompt
		fmt.Printf("ep> ")

		// Tokenize it here
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		terms := split(input)

		// Dispatch it here
		for _, term := range terms {
			// Every term gets dispatched here
		}
	}
}

func ls(ch chan string) {
	go func(){
		files, _ := ioutil.ReadDir(".")
		ret := ""
		for _, v := range files {
			ret += fmt.Sprintf("%s\n", v.Name)
		}
		ch <- ret
	}()
}


/* UTIL FUNCTIONS */

func split(input string) []string {
	rx := regexp.MustCompile("[^ ]+")
	return rx.FindAllString(input, -1)
}
