package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/itchyny/base58-go"
)

func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

func GenerateShortURL(initialLink string, userId string, customText string) string {
	if customText != "" {
		//strip of any whitespace and convert to lowercase
		customText = strings.ToLower(customText)
		customText = strings.Replace(customText, " ", "", -1)
		return customText
	}
	urlHashBytes := sha256Of(initialLink + userId)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	return finalString[:8]
}
