package builtin

import (
	"fmt"
	"syscall"

	"github.com/jvikstedt/watchful"
)

type DiskInfo struct{}

func (d DiskInfo) Identifier() string {
	return "disk_info"
}

func (d DiskInfo) Instruction() watchful.Instruction {
	return watchful.Instruction{
		Input: []watchful.Param{
			watchful.Param{Type: watchful.ParamString, Name: "path", Required: true},
		},
		Output: []watchful.Param{
			watchful.Param{Type: watchful.ParamFloat, Name: "all"},
			watchful.Param{Type: watchful.ParamFloat, Name: "free"},
			watchful.Param{Type: watchful.ParamString, Name: "used"},
		},
	}
}

func (d DiskInfo) Execute(params map[string]interface{}) (map[string]watchful.InputValue, error) {
	path, ok := params["path"].(string)
	if !ok {
		return nil, fmt.Errorf("Expected path to be a string but was %T", params["path"])
	}

	disk := diskUsage(path)

	return map[string]watchful.InputValue{
		"all":  watchful.InputValue{Type: watchful.ParamFloat, Val: float64(disk.All) / float64(gb)},
		"free": watchful.InputValue{Type: watchful.ParamFloat, Val: float64(disk.Free) / float64(gb)},
		"used": watchful.InputValue{Type: watchful.ParamFloat, Val: float64(disk.Used) / float64(gb)},
	}, nil
}

type diskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

func diskUsage(path string) (disk diskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	return
}

const (
	b  = 1
	kb = 1024 * b
	mb = 1024 * kb
	gb = 1024 * mb
)
