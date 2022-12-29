package ostools

import (
	"fmt"
	"io"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/sebastianRau/deployer/pkg/verbose"
)

func Exec(command string, path string, v io.Writer, arg ...string) error {

	cmd := exec.Command(command, arg...)
	cmd.Stderr = v
	cmd.Stdout = v

	if path != "" {
		cmd.Dir = path
	}

	err := cmd.Run()
	return err
}

func ExecEach(search string, command string, path string, ignoreError bool, v io.Writer, args ...string) error {
	if strings.Contains(search, "**") {
		return fmt.Errorf("Wildcard ** is not allowed here")
	}

	files, err := listFiles(search, true)
	if err != nil {
		return err
	}

	for _, f := range files {
		newArgs := replaceArgs(f, args)
		cmd := exec.Command(command, newArgs...)
		cmd.Stderr = v
		cmd.Stdout = v

		if path != "" {
			cmd.Dir = path
		}

		err := cmd.Run()
		if err != nil {
			if ignoreError {
				verbose.Fprintf(v, "%s", err.Error())
			} else {
				return err
			}
		}
	}

	return err
}

func replaceArgs(file string, args []string) []string {
	var argsReplaced []string
	for _, a := range args {
		ar := a
		base := filepath.Base(file)
		dir := filepath.Dir(file)
		ar = strings.ReplaceAll(ar, "__FILE__", base)
		ar = strings.ReplaceAll(ar, "__DIR__", dir)
		ar = strings.ReplaceAll(ar, "__PATH__", dir+"/"+base)
		argsReplaced = append(argsReplaced, ar)
	}
	return argsReplaced
}
