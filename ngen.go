package ngen

import (
	"math"
	"math/rand"

	"time"

	"github.com/iancoleman/strcase"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func getRandomStringFromStringArray(arr []string) string {
	randomRawIndex := rand.Float64() * float64(len(arr)-1)
	randomIndex := int(math.Round(randomRawIndex))
	return arr[randomIndex]
}

func randomAdjective() string {
	return getRandomStringFromStringArray(adjectives)
}

func randomNoun() string {
	return getRandomStringFromStringArray(nouns)
}

func Generate() string {
	return strcase.ToCamel(randomAdjective() + "-" + randomNoun())
}
