package text2int

import (
	"fmt"
	"regexp"
	"strings"
)

type Text2Int struct {
	singles map[string]int
	teens   map[string]int
	tens    map[string]int
	scales  map[string]int
}

func (what *Text2Int) Convert(text string) (int, error) {
	what.init()
	sanitized, sanitizeError := what.sanitize(text)
	if sanitizeError != nil {
		return 0, fmt.Errorf("failed to sanitize text: %v", sanitizeError)
	}

	text = sanitized
	if len(text) == 0 {
		return 0, fmt.Errorf("no text to convert")
	}

	var gotSomething bool
	var total, current int
	for len(text) > 0 {
		value, remaining, scale, found := what.match(text)
		if !found {
			if !gotSomething {
				return total, fmt.Errorf("invalid text %q", text)
			}

			return total, nil
		}

		gotSomething = true
		if scale {
			total += current * value
			current = 0
		} else {
			current += value
		}

		text = remaining
	}

	return total + current, nil
}

// match returns the value of the word, the remaining text, and if the word was a scale, if a word was found
func (what *Text2Int) match(text string) (int, string, bool, bool) {
	for word, value := range what.singles {
		if strings.HasPrefix(text, word) {
			return value, text[len(word):], false, true
		}
	}

	for word, value := range what.teens {
		if strings.HasPrefix(text, word) {
			return value, text[len(word):], false, true
		}
	}

	for word, value := range what.tens {
		if strings.HasPrefix(text, word) {
			return value, text[len(word):], false, true
		}
	}

	for word, value := range what.scales {
		if strings.HasPrefix(text, word) {
			return value, text[len(word):], true, true
		}
	}

	return 0, text, false, false
}

func (what *Text2Int) sanitize(text string) (string, error) {
	regex, regexError := regexp.Compile(`[^a-z]+`)
	if regexError != nil {
		return "", fmt.Errorf("failed to compile regex: %v", regexError)
	}

	return regex.ReplaceAllString(strings.ToLower(text), ""), nil
}

func (what *Text2Int) init() {
	what.singles = map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	what.teens = map[string]int{
		"ten":       10,
		"eleven":    11,
		"twelve":    12,
		"thirteen":  13,
		"fourteen":  14,
		"fifteen":   15,
		"sixteen":   16,
		"seventeen": 17,
		"eighteen":  18,
		"nineteen":  19,
	}

	what.tens = map[string]int{
		"twenty":  20,
		"thirty":  30,
		"forty":   40,
		"fifty":   50,
		"sixty":   60,
		"seventy": 70,
		"eighty":  80,
		"ninety":  90,
	}

	what.scales = map[string]int{
		"hundred":  100,
		"thousand": 1000,
		"million":  1000000,
		"billion":  1000000000,
		"trillion": 1000000000000,
	}
}
