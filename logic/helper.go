package logic

import (
	"crypto/sha256"
	"encoding/base32"
	"strings"
)

// TODO: implement changes to format 10 character strings to tiny url
// compatible strings

// input: qwertyuiop
// output: qwe-rty-uiop
func formatUrlHash(str string) (string, error) {
	if len(str) != 10 {
		return "", ErrQuerryLengthInvalid
	}

	var s strings.Builder
	for i, char := range str {
		if i == 3 || i == 6 {
			s.WriteString("-")
		}

		s.WriteRune(char)
	}
	return s.String(), nil
}

func getHasedValue(url string) string {

	h := sha256.New()
	h.Write([]byte(url))
	bs := h.Sum(nil)

	encoded := base32.StdEncoding.EncodeToString(bs)
	lowerEncoded := strings.ToLower(encoded)

	return lowerEncoded[:10]
}
