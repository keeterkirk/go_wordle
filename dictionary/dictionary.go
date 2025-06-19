package dictionary

import "math/rand/v2"

func RandomWord() string {

	temp := getWords()
	arrayCount := len(temp)
	picker := rand.IntN(arrayCount)
	return temp[picker]
}

func getWords() []string {
	return []string{
		"award",
		"apple",
		"grape",
		"peach",
		"berry",
	}
}
