package steps

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/sebastianRau/deployer/pkg/templating"

	yaml "gopkg.in/yaml.v3"
)

type StepTypes string

const (
	CreateFolderType  StepTypes = "mkdir"
	CopyType          StepTypes = "cp"
	DeleteType        StepTypes = "rm"
	ReplaceTextType   StepTypes = "replaceText"
	FileWriterType    StepTypes = "fileWriter"
	ExecType          StepTypes = "exec"
	EcecEachFileType  StepTypes = "execEachFile"
	GitUpdateType     StepTypes = "gitUpdate"
	SubStepsType      StepTypes = "subSteps"
	CreateSymlinkType StepTypes = "symlink"
)

type DeployerStep struct {
	// mandatory
	Type      string      `json:"type" yaml:"type"`
	Parameter interface{} `json:"parameter" yaml:"parameter"`
	// optinal
	Description string `json:"description" yaml:"description"`
	IgnoreError bool   `json:"ignoreError" yaml:"ignoreError"`
}

type JsonConfig struct {
	Steps []DeployerStep `json:"steps" yaml:"steps"`
}

type ExceutableStep interface {
	Exec(v io.Writer) error
}

func UnmarshalConfigTemplate(templ string, data string) (*JsonConfig, error) {
	jsonResult := JsonConfig{}

	fmt.Println("Reading File")
	jsonConfigFile, err := templating.ParseTemplateJsonData(templ, data)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	err = json.Unmarshal(jsonConfigFile, &jsonResult)
	if err != nil {
		err = yaml.Unmarshal(jsonConfigFile, &jsonResult)
		if err != nil {
			return nil, err
		}
	}

	if jsonResult.Steps == nil {
		return nil, fmt.Errorf("No Steps defined")
	}

	for i, s := range jsonResult.Steps {
		if s.Type == "" {
			return nil, fmt.Errorf("No Type for Step #%d defined", i)
		}
		if s.Parameter == nil {
			return nil, fmt.Errorf("No Parameter for Step #%d (type: %s) defined", i, s.Type)
		}

		switch StepTypes(s.Type) {

		case CreateFolderType:
			s.Parameter, err = UnmarschalCreatefolder(s)
		case CopyType:
			s.Parameter, err = UnmarschalCopy(s)
		case DeleteType:
			s.Parameter, err = UnmarschalDelete(s)
		case FileWriterType:
			s.Parameter, err = UnmarschalFileWriter(s)
		case ReplaceTextType:
			s.Parameter, err = UnmarschalReplaceText(s)
		case ExecType:
			s.Parameter, err = UnmarschalExec(s)
		case EcecEachFileType:
			s.Parameter, err = UnmarschalExecEach(s)
		case GitUpdateType:
			s.Parameter, err = UnmarschalGitUpdate(s)
		case SubStepsType:
			s.Parameter, err = UnmarschalSubSteps(s)
		case CreateSymlinkType:
			s.Parameter, err = UnmarschalCreateSymlink(s)
		default:
			err = fmt.Errorf("cant parse type: %s", s.Type)
		}
		if err != nil {
			fmt.Printf("Step: %v", s.Parameter)
			return nil, fmt.Errorf("error in step %s: %s", s.Description, err.Error())
		}
	}

	return &jsonResult, nil
}

func (c *JsonConfig) Marshal() (string, error) {
	result, err := json.Marshal(&c)
	return string(result), err
}

func (c *JsonConfig) MarshalIndent(ident string) (string, error) {
	result, err := json.MarshalIndent(&c, "", ident)
	return string(result), err
}

func (c *JsonConfig) Exceute(out io.Writer, verboseFlag bool) error {
	return c.exec(out, verboseFlag, false)
}

func (c *JsonConfig) ExceuteDirectVerbose(out io.Writer) error {
	return c.exec(out, false, true)
}

func (c *JsonConfig) exec(out io.Writer, verbose bool, directOut bool) error {

	for _, s := range c.Steps {

		var ex ExceutableStep
		var err error

		// Exec Step
		if ex != nil {
			var write io.Writer
			v := bytes.NewBufferString("")

			if verbose {
				write = v
			} else if directOut {
				write = out
			} else {
				write = nil
			}

			fmt.Fprint(out, formatStep(s.Description, s.Type, (err != nil), 80))
			err = ex.Exec(write)
			if err != nil && !s.IgnoreError {
				fmt.Fprintf(out, "%v\n", err.Error())
				return err
			}

			if verbose {
				_, err = out.Write(v.Bytes())
				if err != nil {
					panic(err.Error())
				}
			}
		}

	}
	return nil
}

func formatStep(sName string, sType string, err bool, len int) string {
	result := "OK"
	if err {
		result = "ERROR"
	}

	return fmt.Sprintf("%-*s %s\n", len, sName, result)
}

func (c *JsonConfig) WriteConfigTo(out io.Writer) error {
	buf, err := yaml.Marshal(&c)
	if err != nil {
		return err
	}
	_, err = out.Write(buf)
	return err

}

func (c *JsonConfig) WriteConfigToFile(filename string) error {

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = c.WriteConfigTo(file)
	if err != nil {
		return err
	}
	err = file.Sync()
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}
