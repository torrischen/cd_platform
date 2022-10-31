package common

type SelectorCond struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type SelectorCondList struct {
	Cond []SelectorCond `json:"cond"`
}
