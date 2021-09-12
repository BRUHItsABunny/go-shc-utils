package decoder

import (
	"fmt"
	"strconv"
	"strings"
)

func ToOriginalString(in string) string {
	runes := []rune{}
	temp := []string{}
	for i, char := range in[5:] {
		temp = append(temp, string(char))
		if (i+1)%2 == 0 {
			runite, err := strconv.Atoi(strings.Join(temp, ""))
			if err != nil {
				fmt.Println(err)
				break
			}
			runes = append(runes, rune(runite+45))
			temp = []string{}
		}
	}

	return string(runes)
}

func EnsurePadding(in string) string {
	rem := len(in) % 4
	for rem > 0 {
		rem--
		in += "="
	}
	return in
}
