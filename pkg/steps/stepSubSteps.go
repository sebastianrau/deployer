package steps

import (
	"encoding/json"
	"io"
)

type SubSteps struct {
	ConfigFile    string `json:"config"`
	DataFile      string `json:"data"`
	EncryptionKey string `json:"encryption"`
	VerboseFlag   bool   `json:"verbose"`
}

func UnmarschalSubSteps(step DeployerStep) (*SubSteps, error) {
	subSteps := SubSteps{}
	b, _ := json.Marshal(step.Parameter)
	err := json.Unmarshal(b, &subSteps)

	return &subSteps, err
}

func (s *SubSteps) Exec(v io.Writer) error {
	subSteps, err := UnmarshalConfigTemplate(s.ConfigFile, s.DataFile, s.EncryptionKey)
	if err != nil {
		return err
	}
	return subSteps.Exceute(v, s.VerboseFlag)
}
