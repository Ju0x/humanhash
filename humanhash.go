package humanhash

import (
	"errors"
	"fmt"
	"strings"
)

// Default list
var wordList = []string{
	"ack", "alabama", "alanine", "alaska", "alpha", "angel", "apart", "april",
	"arizona", "arkansas", "artist", "asparagus", "aspen", "august", "autumn",
	"avocado", "bacon", "bakerloo", "batman", "beer", "berlin", "beryllium",
	"black", "blossom", "blue", "bluebird", "bravo", "bulldog", "burger",
	"butter", "california", "carbon", "cardinal", "carolina", "carpet", "cat",
	"ceiling", "charlie", "chicken", "coffee", "cola", "cold", "colorado",
	"comet", "connecticut", "crazy", "cup", "dakota", "december", "delaware",
	"delta", "diet", "don", "double", "early", "earth", "east", "echo",
	"edward", "eight", "eighteen", "eleven", "emma", "enemy", "equal",
	"failed", "fanta", "fifteen", "fillet", "finch", "fish", "five", "fix",
	"floor", "florida", "football", "four", "fourteen", "foxtrot", "freddie",
	"friend", "fruit", "gee", "georgia", "glucose", "golf", "green", "grey",
	"hamper", "happy", "harry", "hawaii", "helium", "high", "hot", "hotel",
	"hydrogen", "idaho", "illinois", "india", "indigo", "ink", "iowa",
	"island", "item", "jersey", "jig", "johnny", "juliet", "july", "jupiter",
	"kansas", "kentucky", "kilo", "king", "kitten", "lactose", "lake", "lamp",
	"lemon", "leopard", "lima", "lion", "lithium", "london", "louisiana",
	"low", "magazine", "magnesium", "maine", "mango", "march", "mars",
	"maryland", "massachusetts", "may", "mexico", "michigan", "mike",
	"minnesota", "mirror", "mississippi", "missouri", "mobile", "mockingbird",
	"monkey", "montana", "moon", "mountain", "muppet", "music", "nebraska",
	"neptune", "network", "nevada", "nine", "nineteen", "nitrogen", "north",
	"november", "nuts", "october", "ohio", "oklahoma", "one", "orange",
	"oranges", "oregon", "oscar", "oven", "oxygen", "papa", "paris", "pasta",
	"pennsylvania", "pip", "pizza", "pluto", "potato", "princess", "purple",
	"quebec", "queen", "quiet", "red", "river", "robert", "robin", "romeo",
	"rugby", "sad", "salami", "saturn", "september", "seven", "seventeen",
	"shade", "sierra", "single", "sink", "six", "sixteen", "skylark", "snake",
	"social", "sodium", "solar", "south", "spaghetti", "speaker", "spring",
	"stairway", "steak", "stream", "summer", "sweet", "table", "tango", "ten",
	"tennessee", "tennis", "texas", "thirteen", "three", "timing", "triple",
	"twelve", "twenty", "two", "uncle", "undress", "uniform", "uranus", "utah",
	"vegan", "venus", "vermont", "victor", "video", "violet", "virginia",
	"washington", "west", "whiskey", "white", "william", "winner", "winter",
	"wisconsin", "wolfram", "wyoming", "xray", "yankee", "yellow", "zebra",
	"zulu",
}

// Given digest will be converted to a human-hash.
// wordcount specifies the output of words in the hash (normally 4)
// seperator will specify how the words will be seperated (normally "-")
func Humanize(digest []byte, wordcount int, seperator string) (string, error) {
	compressed, err := compress(digest, wordcount)
	if err != nil {
		return "", err
	}

	var result []string

	for _, comp_byte := range compressed {
		result = append(result, wordList[comp_byte])
	}

	return strings.Join(result, seperator), err
}

func compress(digest []byte, wordcount int) ([]byte, error) {
	res := make([]byte, wordcount)

	if int(wordcount) > len(digest) {
		return []byte{}, errors.New("wordcount higher than digest length")
	}

	segments := make([][]byte, wordcount)
	segmentSize := len(digest) / wordcount

	for i := 0; i < wordcount; i++ {
		for b := 0; b < segmentSize; b++ {
			segments[i] = append(segments[i], digest[(i*segmentSize)+b])
		}
	}
	fmt.Println(segments)
	// Compress using XOR per segment
	for i, segment := range segments {
		for _, b := range segment {
			res[i] ^= b
		}
	}

	return res, nil
}
