package conf

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Path() {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	path = path[:index]
}
