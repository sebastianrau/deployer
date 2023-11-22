package steps

import (
	"encoding/json"
	"io"

	filetools "github.com/sebastianrau/deployer/pkg/fileTools"
)

type ReplaceText struct {
	File           string            `json:"destination"`
	ReplacementMap map[string]string `json:"map"`
}

func UnmarschalReplaceText(step DeployerStep) (*ReplaceText, error) {
	mkdir := ReplaceText{}
	b, _ := json.Marshal(step.Parameter)
	err := json.Unmarshal(b, &mkdir)

	return &mkdir, err
}

func (s *ReplaceText) Exec(v io.Writer) error {
	return filetools.ReplaceAll(s.File, s.File, s.ReplacementMap, v)
}
