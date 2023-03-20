package getters

import (
	"os"
	"testing"
)

func Test_ShouldDownloadS3File(t *testing.T) {
	// arrange
	if os.Getenv("ENVIRONMENT") != "local" {
		t.Skip("Test should only be executed locally")
	}

	t.Setenv("BUCKET_NAME", "wordle-languages")
	t.Setenv("AWS_DEFAULT_REGION", "us-east-1")
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
