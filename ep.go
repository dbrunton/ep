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
	// two senses of foo, foo1 and foo2
	foo1 := &term{label: "foo", pos: &partOfSpeech{label: "one", arity: 1} }
	foo2 := &term{label: "foo", pos: &partOfSpeech{label: "two", arity: 2} }

	switch {
		// decide which foo
		case foo1 != nil: fmt.Println(foo1, foo1.pos)
		case foo2 != nil: fmt.Println(foo2)
	}
}

func bar() {
	bar1 := &term{label: "bar", pos: &partOfSpeech{label: "one", arity: 1} }
	fmt.Println(bar1, bar1.pos)
}

