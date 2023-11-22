package steps

import (
	"encoding/json"
	"io"

	ostools "github.com/sebastianrau/deployer/pkg/osTools"
)

type ExecEach struct {
	Search      string   `json:"search"`
	Command     string   `json:"command"`
	Path        string   `json:"path"`
	Params      []string `json:"params"`
	ignoreError bool
}

func UnmarschalExecEach(step DeployerStep) (*ExecEach, error) {
	mkdir := ExecEach{}
	b, _ := json.Marshal(step.Parameter)
	err := json.Unmarshal(b, &mkdir)
	mkdir.ignoreError = step.IgnoreError

	return &mkdir, err
}

func (s *ExecEach) Exec(v io.Writer) error {
	return ostools.ExecEach(s.Search, s.Command, s.Path, s.ignoreError, v, s.Params...)
}
