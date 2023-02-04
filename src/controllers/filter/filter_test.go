package filter

import (
	"testing"

	"github.com/matheus-osorio/go-term-solver/src/entry"
)

func Test_ShouldCreateFilter(t *testing.T) {
	t.Parallel()
	// arrange
	filter := WordFilter{}
	rawFilter := []entry.Filter{
		{
			Letter: "a",
			Status: entry.GREEN,
		},
		{
			Letter: "a",
			Status: entry.YELLOW,
		},
		{
			Letter: "a",
			Status: entry.GREY,
		},
		{
			Letter: "b",
			Status: entry.YELLOW,
		},
		{
			Letter: "c",
			Status: entry.GREY,
		},
	}
	// act
	filter.CreateFilter(rawFilter)

	// assert
	specificFilter := filter.Rules[0]

	if !specificFilter.Exact {
		t.Errorf("The filter was supposed to qualify this rule as Exact")
	}

	if specificFilter.NumberOfTimes != 2 {
		t.Errorf("The filter was supposed to find 2 instances of the letter in the filter")
	}

	if specificFilter.In[0] != 0 {
		t.Errorf("The filter was supposed to get the \"In\" rule to the first Index")
	}

	if specificFilter.NotIn[0] != 1 {
		t.Errorf("The filter was supposed to get the \"NotIn\" rule to the second Index")
	}
}

func Test_ShouldVerifyValidWord(t *testing.T) {
	t.Parallel()
	// arrange
	rawFilter := []entry.Filter{
		{
			Letter: "a",
			Status: entry.GREEN,
		},
		{
			Letter: "a",
			Status: entry.YELLOW,
		},
		{
			Letter: "a",
			Status: entry.GREY,
		},
		{
			Letter: "b",
			Status: entry.YELLOW,
		},
		{
			Letter: "c",
			Status: entry.GREY,
		},
	}

	filter := WordFilter{}

	// act
	filter.CreateFilter(rawFilter)

	result1 := filter.IsWordValid("bbaad")
	result2 := filter.IsWordValid("aabdd")
	result3 := filter.IsWordValid("abaad")
	result4 := filter.IsWordValid("ababd")
	result5 := filter.IsWordValid("abacd")
	result6 := filter.IsWordValid("abadd")

	// assert

	if result1 {
		t.Errorf("Result 1 should be invalidated by the first rule")
	}

	if result2 {
		t.Errorf("Result 2 should be invalidated by the second filter")
	}

	if result3 {
		t.Errorf("Result 3 should be invalidated by the third filter")
	}

	if result4 {
		t.Errorf("Result 4 should be invalidated by the forth filter")
	}

	if result5 {
		t.Errorf("Result 5 should be invalidated by the fifth filter")
	}

	if !result6 {
		t.Errorf("Result 6 should be valid")
	}
}

func Test_ShouldFilterWords(t *testing.T) {
	t.Parallel()
	// arrange
	rawFilter := []entry.Filter{
		{
			Letter: "a",
			Status: entry.GREEN,
		},
		{
			Letter: "a",
			Status: entry.YELLOW,
		},
		{
			Letter: "a",
			Status: entry.GREY,
		},
		{
			Letter: "b",
			Status: entry.YELLOW,
		},
		{
			Letter: "c",
			Status: entry.GREY,
		},
	}

	wordList := []string{
		"bbaad", // out by rule: 1
		"aabdd", // out by rule: 2
		"abaad", // out by rule: 3
		"ababd", // out by rule: 4
		"abacd", // out by rule: 5
		"abadd", //valid
	}

	// act
	filter := WordFilter{}
	filter.CreateFilter(rawFilter)
	result := filter.FilterWords(wordList)

	if value := len(result); value != 1 || result[0] != "abadd" {
		t.Errorf("Expected 1 word to not be filtered out. Got %d", value)
	}
}
