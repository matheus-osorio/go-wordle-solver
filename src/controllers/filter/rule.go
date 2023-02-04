package filter

/*
 * The struct is a smarter method of checking what words actually are valid
 * A green rule will always add to the "In" slice and increase number of times
 * A yellow rule will always add to the "NotIn" slice and increase the number of times
 * A grey rule sets the exact as "true"
 *
 * When checking, it is verified in the following way.
 * The In rule verifies that the letter in the position is equal
 * The NotIn rule only verifies if the letter in the position is different
 * After both are verified, the number of appearances is checked
 * If the number is lower than expected, it is not valid.
 * if the number is different than expect AND the Exact is set to true, it is not valid
 * If the number is equal or bigger than expected AND the Exact is set to false, it is valid
 */
type Rule struct {
	Letter        rune
	In            []int
	NotIn         []int
	NumberOfTimes int
	Exact         bool
}
