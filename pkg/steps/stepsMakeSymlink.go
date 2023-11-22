package steps

import (
	"encoding/json"
	"fmt"
	"io"

	ostools "github.com/sebastianrau/deployer/pkg/osTools"
	"github.com/sebastianrau/deployer/pkg/verbose"
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
		fmt.Printf("creating symlink %s --> %s\n", src, dst)

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
