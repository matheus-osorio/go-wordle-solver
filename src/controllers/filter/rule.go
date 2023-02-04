package filter

type Rule struct {
	Letter        rune
	In            []int
	NotIn         []int
	NumberOfTimes int
	Exact         bool
}
