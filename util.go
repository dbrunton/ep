package main

import(
	"regexp"
)

/* UTIL FUNCTIONS */

func split(input string) []string {
	rx := regexp.MustCompile("[^ ]+")
	return rx.FindAllString(input, -1)
}

func chomp(input string) string {
	rx := regexp.MustCompile("\n")
	return rx.ReplaceAllString(input, "")
}
