package steps

import (
	"encoding/json"
	"io"

	ostools "github.com/sebastianrau/deployer/pkg/osTools"
	"github.com/sebastianrau/deployer/pkg/verbose"
)

type Delete struct {
	Destination []string `json:"destination"`
	ignoreError bool
}

func UnmarschalDelete(step DeployerStep) (*Delete, error) {
	rm := Delete{}
	b, _ := json.Marshal(step.Parameter)
	err := json.Unmarshal(b, &rm)
	rm.ignoreError = step.IgnoreError
	return &rm, err
}

func (s *Delete) Exec(v io.Writer) error {
	for _, d := range s.Destination {
		err := ostools.Delete(d, v)
		if err != nil {
			if s.ignoreError {
				verbose.Fprintf(v, "%s/n", err.Error())
			} else {
				return err
			}
		}
	}
	return nil
}
