package watchful

type Executable interface {
	Identifier() string
	Instruction() Instruction
	Execute(map[string]interface{}) (map[string]interface{}, error)
}

type ParamType int

const (
	ParamInt ParamType = iota
	ParamString
	ParamFloat
	ParamBytes
	ParamAny
)

type Param struct {
	Type     ParamType `json:"type"`
	Name     string    `json:"name"`
	Required bool      `json:"required"`
}

type Instruction struct {
	Dynamic bool    `json:"dynamic"`
	Input   []Param `json:"input"`
	Output  []Param `json:"output"`
}
