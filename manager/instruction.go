package manager

type ParamType int

const (
	ParamInt ParamType = iota
	ParamString
	ParamFloat
)

type Param struct {
	ParamType
	Name     string
	Required bool
}

type Instruction struct {
	Takes   []Param
	Returns []Param
}
