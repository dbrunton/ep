package main

import "fmt"

// Every term has a part of speech:
type partOfSpeech struct {
	label string
	arity int
	neighbors []*partOfSpeech
}

func (p *partOfSpeech) Arity() int {	// how many neighbors
//	return len(p.neighbors)
	return p.arity
}

func (p *partOfSpeech) String() string {
	return fmt.Sprintf("Part of Speech: %s", p.label)
}

