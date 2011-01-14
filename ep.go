package main

import(
	"fmt"
	"bufio"
	"os"
)

func main() {

	for {
		// Prompt
		fmt.Printf("ep> ")

		// Tokenize it here
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		terms := split(chomp(input))

		// Dispatch it here
		for _, term := range terms {
			switch term {
				case "foo": foo()
				case "bar": bar()
			}
		}
	}
}

// A term is a label plus a part of speech
// For bootstrapping purposes, no name collisions
type term struct {
	label string
	pos *partOfSpeech
}

func (t *term) String() string {
	return fmt.Sprintf("label: %s, arity: %d", t.label, t.pos.Arity())
}

func foo() {
	// one and two will be in a grammar somewhere someday
	one := &partOfSpeech{label: "one", arity: 1}
	two := &partOfSpeech{label: "two", arity: 2}

	// two senses of foo, foo1 and foo2
	foo1 := &term{label: "foo", pos: one} // Definition 1
	foo2 := &term{label: "foo", pos: two}

	switch {
		// decide which foo
		case foo1 != nil: fmt.Println(foo1, foo1.pos)
		case foo2 != nil: fmt.Println(foo2)
	}
}

func bar() {
	one := &partOfSpeech{label: "one", arity: 1}
	bar1 := &term{label: "bar", pos: one}
	fmt.Println(bar1)
}

