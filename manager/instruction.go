package manager

type ParamType int

const (
	ParamInt ParamType = iota
	ParamString
	ParamFloat
	ParamBytes
)

type Param struct {
	Type     ParamType `json:"type"`
	Name     string    `json:"name"`
	Required bool      `json:"required"`
}

type Instruction struct {
	Dynamic bool    `json:"dynamic"`
	Takes   []Param `json:"takes"`
	Returns []Param `json:"returns"`
}
