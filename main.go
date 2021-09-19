package ngen

import (
	"math"
	"math/rand"
	"os"
)

var adjectives []string
var adverbs []string
var nouns []string
var verbs []string

func init() {
	var err error
	adjectives, err = readLines("./ngen/words/adjectives.txt")
	if err != nil {
		os.Exit(1)
	}
	adverbs, err = readLines("./words/adverbs.txt")
	if err != nil {
		os.Exit(1)
	}
	nouns, err = readLines("./words/nouns.txt")
	if err != nil {
		os.Exit(1)
	}
	verbs, err = readLines("./words/verbs.txt")
	if err != nil {
		os.Exit(1)
	}
}

func Generate() string {
	randomRawAdjectiveIndex := rand.Float64() * float64(len(adjectives)-1)
	randomAdjectiveIndex := int(math.Round(randomRawAdjectiveIndex))
	return adjectives[randomAdjectiveIndex]
}
