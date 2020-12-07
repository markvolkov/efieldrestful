package access

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
)

func getSupportedSpecialChars() []string {
	return []string{"#", "$"}
}

func getSupportedNormalChars() []string {
	return []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
}

func isInt(i string) bool {
	return regexp.MustCompile("^[0-9]+$").MatchString(i)
}

func GenerateAccessCode() string {
	const codeLength = 5
	var specialChars = getSupportedSpecialChars()
	var normalChars = getSupportedNormalChars()
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := ""
	for i := 0; i < codeLength; i++ {
		if random.Intn(15) < 1 {
			result += specialChars[random.Intn(len(specialChars))]
		} else {
			next := normalChars[random.Intn(len(normalChars))]
			if random.Intn(2) == 0 && !isInt(next) {
				result += strings.ToUpper(next)
			} else {
				result += next
			}
		}
	}
	return result
}