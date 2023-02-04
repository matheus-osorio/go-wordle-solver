package getters

import (
	"fmt"
	"os"
	"strconv"
)

// Factory for the getters package.
func CreateGetter(wordSize uint, language string) GetterInterface {
	environmentString := os.Getenv("IS_OFFLINE")
	isOffline, err := strconv.ParseBool(environmentString)
	if err != nil {
		fmt.Println("WARNING:COULD NOT GET IS OFFLINE VARIABLE: ", environmentString)
		isOffline = false
	}

	if isOffline {
		fmt.Println("Getting from local word list")
		return LocalGetter{
			WordPicker{
				WordSize: wordSize,
			},
		}
	}

	return S3Getter{
		WordPicker: WordPicker{
			WordSize: wordSize,
		},
		Language: language,
	}
}
