package manager

type ParamType int

const (
	ParamInt ParamType = iota
	ParamString
	ParamFloat
)

type Param struct {
	Type     ParamType `json:"type"`
	Name     string    `json:"name"`
	Required bool      `json:"required"`
}

type Instruction struct {
	Takes   []Param `json:"takes"`
	Returns []Param `json:"returns"`
}
