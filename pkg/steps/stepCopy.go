package steps

import (
	"encoding/json"
	"io"

	ostools "github.com/sebastianRau/deployer/pkg/osTools"
	"github.com/sebastianRau/deployer/pkg/verbose"
)

type Copy struct {
	SrcDestMap     map[string]string `json:"src_dest"`
	IgnoreDotFiles bool              `json:"ignoreDotFiles"`
	ignoreError    bool
}

func UnmarschalCopy(step DeployerStep) (*Copy, error) {
	copy := Copy{}
	b, _ := json.Marshal(step.Parameter)
	err := json.Unmarshal(b, &copy)
	copy.ignoreError = step.IgnoreError

	return &copy, err
}

func (s *Copy) Exec(v io.Writer) error {
	for src, dst := range s.SrcDestMap {
		err := ostools.Copy(src, dst, s.IgnoreDotFiles, v)

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
