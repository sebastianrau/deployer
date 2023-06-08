package steps

import (
	"encoding/json"
	"io"

	ostools "github.com/sebastianRau/deployer/pkg/osTools"
	"github.com/sebastianRau/deployer/pkg/verbose"
)

type CreateSymlink struct {
	SrcDestMap  map[string]string `json:"src_dest"`
	ignoreError bool
}

func UnmarschalCreateSymlink(step DeployerStep) (*CreateSymlink, error) {
	mksymlink := CreateSymlink{}
	b, _ := json.Marshal(step.Parameter)
	err := json.Unmarshal(b, &mksymlink)

	return &mksymlink, err
}

func (s *CreateSymlink) Exec(v io.Writer) error {
	for src, dst := range s.SrcDestMap {
		err := ostools.MakeSymlink(src, dst, v)

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