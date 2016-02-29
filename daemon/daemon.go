// http://sysmagazine.com/posts/187668/
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

const (
	envVarName  = "_GO_DAEMON"
	envVarValue = "1"
)

func main() {
	// ?
}

func Reborn(umask uint32, workDir string) (err error) {
	if !WasReborn() {
		var path string
		if path, err = filepath.Abs(os.Args[0]); err != nil {
			return
		}
		cmd := exec.Command(path, os.Args[1:]...)
		envVar := fmt.Sprintf("%s=%s", envVarName, envVarValue)
		cmd.Env = append(os.Environ(), envVar)
		if err = cmd.Start(); err != nil {
			return
		}
		os.Exit(0)
	}
	syscall.Umask(int(umask))
	if len(workDir) == 0 {
		if err = os.Chdir(workDir); err != nil {
			return
		}
	}
	_, err = syscall.Setsid()
	return
}
func WasReborn() bool {
	return os.Getenv(envVarName) == envVarValue
}
