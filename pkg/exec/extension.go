package exec

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"

	"github.com/jvikstedt/watchful"
)

// SearchExtensions recursively loads all plugins in target folderPath and
// converts them to watchful.Executable. Then calls callback fn with each of them
func SearchExtensions(folderPath string, fn func(watchful.Executable, error)) error {
	files, err := findFiles(folderPath)
	if err != nil {
		return err
	}

	for _, f := range files {
		p, err := plugin.Open(f)
		if err != nil {
			fn(nil, err)
		}
		s, err := p.Lookup("Extension")
		if err != nil {
			fn(nil, err)
		}

		executable, ok := s.(watchful.Executable)
		if !ok {
			fn(nil, fmt.Errorf("Could not convert to watchful.Executable"))
		}

		fn(executable, nil)
	}

	return nil
}

func findFiles(folderPath string) ([]string, error) {
	var files []string
	err := filepath.Walk(folderPath, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".so" {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}
