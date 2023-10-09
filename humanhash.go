// Humanhash establishes a human-readable representation of a digest
package humanhash

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"github.com/google/uuid"
)

var (
	wordList = defaultWordList
)

// Defines a custom wordlist (Should have 256 words)
func Wordlist(path string) {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wordList = append(wordList, scanner.Text())
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
}

// Given digest will be converted to a human-hash.
// wordcount specifies the output of words in the hash (normally 4)
// seperator will specify how the words will be seperated (normally "-")
//
// WARNING:
// Human-hashes aren't safe! The chance of a collision with 256 words is 256^wordcount (With 4 words 1 : ~4.3 Billion)
// Only use them for a better-readable version of a modern hash function.
func Humanize(digest []byte, wordcount int, seperator string) (string, error) {
	compressed, err := compress(digest, wordcount)
	if err != nil {
		return "", err
	}

	result := make([]string, len(compressed))

	for i, comp_byte := range compressed {
		result[i] = wordList[comp_byte]
	}

	return strings.Join(result, seperator), err
}

// Compresses the given digest into <wordcount> segments
func compress(digest []byte, wordcount int) ([]byte, error) {
	res := make([]byte, wordcount)
	length := len(digest)

	if int(wordcount) > length {
		return []byte{}, errors.New("wordcount higher than digest length")
	}

	segmentSize := length / wordcount
	remainder := length % wordcount

	for i := 0; i < wordcount; i++ {
		start := i * segmentSize
		end := start + segmentSize

		if i == wordcount-1 {
			end += remainder
		}

		for _, b := range digest[start:end] {
			res[i] ^= b
		}
	}

	return res, nil
}

// Defines a pair of UUID and the related Humanhash
func UUID() (string, uuid.UUID, error) {
	id := uuid.New()
	humanhash, err := Humanize(id[:], 4, "-")
	return humanhash, id, err
}
