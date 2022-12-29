package steps

import (
	"encoding/json"
	"io"

	filetools "github.com/sebastianRau/deployer/pkg/fileTools"
)

type FileWriter struct {
	Filename string   `json:"filename"`
	Text     []string `json:"text"`
}

func UnmarschalFileWriter(step DeployerStep) (*FileWriter, error) {
	mkdir := FileWriter{}
	b, _ := json.Marshal(step.Parameter)
	err := json.Unmarshal(b, &mkdir)

	return &mkdir, err
}

func (s *FileWriter) Exec(v io.Writer) error {
	return filetools.WriteFile(s.Filename, s.Text, v)
}
