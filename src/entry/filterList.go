package entry

const (
	GREEN  = "green"
	YELLOW = "yellow"
	GREY   = "grey"
)

/*
Expected body for the FilterList. Expects to receive a list of ObjectRuleList.
Each list represents a single filter. That way, you can still use the API if you are using
the double mode.
*/
type FilterListBodyObject struct {
	Rules []ObjectRuleList `json:"rules"`
}

/*
Filter struct. Expects a letter and a Status.
The letter can be any one letter.
The status must be either grey, yellow, grey
*/
type Filter struct {
	Letter string `json:"letter"`
	Status string `json:"status"`
}

/*
Filter rules + word list. Holds a single list of words.
*/
type ObjectRuleList struct {
	Filter       []Filter `json:"filter"`
	Words        []string `json:"words"`
	RemovedWords []string `json:"removedWords"`
}
