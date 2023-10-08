package humanhash

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

var (
	wordList = []string{}
)

func init() {
	Wordlist("./wordlists/default.txt")
}

// Define a custom wordlist (Should have 256 words)
// Default is ./wordlists/default.txt
func Wordlist(path string) {
	file, err := os.Open(path) // Replace "file.txt" with the path to your file
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
