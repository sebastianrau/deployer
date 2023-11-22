package steps

import (
	"encoding/json"
	"io"

	ostools "github.com/sebastianrau/deployer/pkg/osTools"
)

type CreateFolder struct {
	Destination []string `json:"destination"`
}

func UnmarschalCreatefolder(step DeployerStep) (*CreateFolder, error) {
	mkdir := CreateFolder{}
	b, _ := json.Marshal(step.Parameter)
	err := json.Unmarshal(b, &mkdir)

	return &mkdir, err
}

func (s *CreateFolder) Exec(v io.Writer) error {
	for _, d := range s.Destination {
		err := ostools.Mkdir(d, v)
		if err != nil {
			return err
		}
	}
	return nil
}
