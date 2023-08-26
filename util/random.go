package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

const alphabet = "zyxwvutsrqponmlkjihgfedcba"

var loremwords [30]string = [30]string{"adipiscing", "amet", "congue", "consectetur", "cursus", "dolor", "eget", "elit", "eleifend", "erat", "faucibus", "ipsum", "jpsum", "lectus", "ligula", "lorem", "luctus", "massa", "mattis", "mauris", "nec", "orci", "pellentesque", "sed", "semper", "sit", "sapien", "suscipit", "tempus"}

func init() {
	rand.NewSource(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomUUID() string {
	// Generate a UUID using the library function
	id, err := uuid.NewRandom()
	if err != nil {
		fmt.Printf("Failed to generate UUID: %v\n", err)
		return ""
	}
	return id.String()
}

func RandomLorem(num int64) string {
	lorem := ""

	for i := int64(0); i < num; i++ {
		lorem += loremwords[rand.Intn(30)] + " "
	}

	return lorem
}

func RandomWords() int64 {
	return RandomInt(5, 8)
}

func RandomSentence() string {
	words := RandomWords()
	text := RandomLorem(words)
	return strings.TrimSpace(strings.ToUpper(text[0:1])+text[1:]) + "."
}
