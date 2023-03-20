package getters

import (
	"testing"
)

func Test_ShouldDownloadS3File(t *testing.T) {
	// arrange
	t.Setenv("BUCKET_NAME", "wordle-languages")
	getter := S3Getter{
		WordPicker: WordPicker{
			WordSize: 5,
		},
		Language: "pt-br",
	}

	// act
	wordList := getter.GetWords()

	// assert
	if len(wordList) <= 0 {
		t.Error("Should have gotten word size above 0!")
	}
}
