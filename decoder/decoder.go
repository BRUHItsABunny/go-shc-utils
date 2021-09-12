package decoder

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"fmt"
	"io"
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

func DecodeData(in string) ([]byte, error) {
	dataBytes, err := base64.URLEncoding.DecodeString(EnsurePadding(in))

	if err == nil {
		reader := flate.NewReader(bytes.NewReader(dataBytes))
		bodyBytes, err := io.ReadAll(reader)
		return bodyBytes, err
	}

	return nil, err
}
