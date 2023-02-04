package getters

import (
	"testing"
)

func Test_ShouldSuccessfullyGetWords(t *testing.T) {
	// arrange
	//act
	getter := LocalGetter{
		WordPicker{
			WordSize: 5,
		},
	}

	// assert
	if result := getter.GetWords(); len(result) <= 0 {
		t.Errorf("Local Getter could not find word list")
	}
}
