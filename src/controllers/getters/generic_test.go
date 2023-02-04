package getters

import (
	"testing"
)

func Test_ShouldSelectWordsBySize(t *testing.T) {
	picker := WordPicker{
		WordList: []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa"},
		WordSize: 5,
	}

	picker.SelectWordsBySize()
	if len(picker.WordList) != 1 {
		t.Errorf("Expected result to be length 1, found %d", len(picker.WordList))
	}
}

func Test_ShouldReplaceLatinCharacters(t *testing.T) {
	picker := WordPicker{
		WordList: []string{"áãâàéèẽêíìĩîóòõôúùũû"},
	}

	picker.ReplaceLatinCharacters()

	if picker.WordList[0] != "aaaaeeeeiiiioooouuuu" {
		t.Errorf("Expected latin characters to be replaced. Got %s", picker.WordList[0])
	}
}
