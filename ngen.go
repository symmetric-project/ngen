package ngen

import (
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/iancoleman/strcase"
)

var adjectives []string
var adverbs []string
var nouns []string
var verbs []string

func init() {
	rand.Seed(time.Now().UnixNano())
	var err error
	adjectives, err = readLines("./words/adjectives.txt")
	if err != nil {
		log.Fatal(err)
	}
	adverbs, err = readLines("./words/adverbs.txt")
	if err != nil {
		log.Fatal(err)
	}
	nouns, err = readLines("./words/nouns.txt")
	if err != nil {
		log.Fatal(err)
	}
	verbs, err = readLines("./words/verbs.txt")
	if err != nil {
		log.Fatal(err)
	}
}

func getRandomStringFromStringArray(arr []string) string {
	randomRawIndex := rand.Float64() * float64(len(arr)-1)
	randomIndex := int(math.Round(randomRawIndex))
	return arr[randomIndex]
}

func adjective() string {
	return getRandomStringFromStringArray(adjectives)
}

func adverb() string {
	return getRandomStringFromStringArray(adverbs)
}

func noun() string {
	return getRandomStringFromStringArray(nouns)
}

func verb() string {
	return getRandomStringFromStringArray(verbs)
}

func Generate() string {
	return strcase.ToCamel(adjective() + "-" + noun())
}
