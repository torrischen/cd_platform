package common

type CreateArgs struct {
	Metadata []byte `json:"metadata"`
	Project  string `json:"project"`
}
