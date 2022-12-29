package steps

import (
	"encoding/json"
	"fmt"
	"io"

	gittools "github.com/sebastianRau/deployer/pkg/gitTools"
)

type GitUpdate struct {
	Url            string `json:"url"`
	Directory      string `json:"dir"`
	PrivateKeyFile string `json:"keyfile"`
	Password       string `json:"password"`
	ErrOnNoUpdate  bool   `json:"ErrorOnNoUpdate"`
}

func UnmarschalGitUpdate(step DeployerStep) (*GitUpdate, error) {
	mkdir := GitUpdate{}
	b, _ := json.Marshal(step.Parameter)
	err := json.Unmarshal(b, &mkdir)
	return &mkdir, err
}

func (s *GitUpdate) Exec(v io.Writer) error {
	_, new, err := gittools.Update(
		s.Url, s.Directory,
		s.PrivateKeyFile,
		s.Password,
		v,
	)

	if err != nil {
		return err
	}

	if !new && s.ErrOnNoUpdate {
		return fmt.Errorf("no upadte from gut repo")
	}

	return nil
}
