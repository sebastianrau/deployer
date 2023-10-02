package templating

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"os"
	"text/template"

	yaml "gopkg.in/yaml.v3"
)

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func hasNext(array []interface{}, idx int) bool {
	return idx < len(array)-1
}
func isLast(array []interface{}, idx int) bool {
	return idx == len(array)-1
}
func arrayJoin(array []interface{}, separator string, addLast bool) string {
	var buf strings.Builder
	for i, v := range array {
		if i == len(array)-1 && !addLast {
			buf.WriteString(fmt.Sprintf("%v", v))
		} else {
			buf.WriteString(fmt.Sprintf("%v%s", v, separator))
		}

	}
	return buf.String()
}

func secret(secretText string, keyFile string) string {
	secretFile, err := os.Open(keyFile)
	if err != nil {
		return ""
	}
	defer secretFile.Close()

	keyBytes, _ := io.ReadAll(secretFile)
	return Decrypt(secretText, keyBytes)
}

func ParseTemplateJsonData(templ string, data string) ([]byte, error) {
	m := map[string]interface{}{}

	if data != "" {
		// parse data json file to map
		dataFile, err := os.Open(data)
		if err != nil {
			return []byte(""), err
		}

		dataBytes, _ := io.ReadAll(dataFile)
		if err := json.Unmarshal(dataBytes, &m); err != nil {
			err := yaml.Unmarshal(dataBytes, &m)
			if err != nil {
				return []byte(""), err
			}
		}

	} else {
		// parse empty json as data
		if err := json.Unmarshal([]byte("{ }"), &m); err != nil {
			return []byte(""), err
		}
	}

	templateFile, _ := os.Open(templ)
	templateBytes, _ := io.ReadAll(templateFile)

	funcMap := template.FuncMap{
		"add":       add,
		"sub":       sub,
		"hasNext":   hasNext,
		"isLast":    isLast,
		"arrayjoin": arrayJoin,
		"join":      strings.Join,
		"secret":    secret,
	}

	t, err := template.New("").Funcs(funcMap).Parse(string(templateBytes))
	if err != nil {
		return []byte(""), err
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, m); err != nil {
		fmt.Println("exec error")
		return []byte(""), err
	}

	return tpl.Bytes(), nil
}
