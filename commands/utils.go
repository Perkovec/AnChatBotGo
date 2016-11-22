package commands

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var alphabetG = []string{"а", "о", "и", "е", "ё", "э", "ы", "у", "ю", "я"}
var alphabetS = []string{"б", "в", "г", "д", "ж", "з", "к", "л", "м", "н", "п", "р", "с", "т", "ф", "х", "ц", "ч", "ш", "щ"}

func TimeDiff(diff int) string {
	var text string

	if diff < 60 {
		text = strconv.Itoa(diff) + " секунд"
	} else {
		minutes := int(diff / 60)
		if minutes < 60 {
			text = strconv.Itoa(minutes) + " минут"
		} else {
			hours := int(minutes / 60)
			if hours < 24 {
				text = strconv.Itoa(hours) + " часов"
			} else {
				text = strconv.Itoa(int(hours/24)) + " дней"
			}
		}
	}

	return text
}

func NumberToLetters(num int) string {
	var result string
	for num > 0 {
		num--
		remainder := num % 26
		digit := remainder + 97
		result = fmt.Sprintf("%c", digit) + result
		num = (num - remainder) / 26
	}
	return strings.ToUpper(result)
}

func LettersToNumber(val string) int {
	val = reverse(val)
	values := []rune(strings.ToLower(val))
	var num = 0
	for i := 0; i < len(values); i++ {
		num += (int(values[i]) - 96) * int(math.Pow(26, float64(i)))
	}
	return num
}

func Nickname(syllableCount int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	if syllableCount < 1 {
		return ""
	}

	var output = ""
	for i := 0; i < syllableCount; i += 1 {
		output += alphabetS[rand.Intn(len(alphabetS))]
		output += alphabetG[rand.Intn(len(alphabetG))]
	}

	return strings.Title(output)
}

func reverse(value string) string {
	data := []rune(value)
	result := []rune{}
	for i := len(data) - 1; i >= 0; i-- {
		result = append(result, data[i])
	}
	return string(result)
}
