package api

type SelectorCond struct {
	Key       string   `json:"key"`
	Operation string   `json:"operation"`
	Value     []string `json:"value"`
}

type SelectorCondList struct {
	Cond []SelectorCond `json:"cond"`
}
