package steps

import (
	"encoding/json"
	"io"

	ostools "github.com/sebastianrau/deployer/pkg/osTools"
)

type Exec struct {
	Command string   `json:"command"`
	Path    string   `json:"path"`
	Params  []string `json:"params"`
}

func UnmarschalExec(step DeployerStep) (*Exec, error) {
	mkdir := Exec{}
	b, _ := json.Marshal(step.Parameter)
	err := json.Unmarshal(b, &mkdir)

	return &mkdir, err
}

func (s *Exec) Exec(v io.Writer) error {
	return ostools.Exec(s.Command, s.Path, v, s.Params...)
}
