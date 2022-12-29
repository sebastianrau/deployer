package steps

import (
	"encoding/json"
	"io"
)

type SubSteps struct {
	ConfigFile  string `json:"config"`
	DataFile    string `json:"data"`
	Out         io.Writer
	VerboseFlag bool
}

func UnmarschalSubSteps(step DeployerStep, out io.Writer, verboseFlag bool) (*SubSteps, error) {
	mkdir := SubSteps{}
	b, _ := json.Marshal(step.Parameter)
	err := json.Unmarshal(b, &mkdir)
	mkdir.Out = out
	mkdir.VerboseFlag = verboseFlag

	return &mkdir, err
}

func (s *SubSteps) Exec(v io.Writer) error {
	config, err := UnmarshalConfigTemplate(s.ConfigFile, s.DataFile)
	if err != nil {
		return err
	}
	return config.Exceute(s.Out, s.VerboseFlag)
}
