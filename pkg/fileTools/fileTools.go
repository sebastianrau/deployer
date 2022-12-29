package filetools

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/sebastianRau/deployer/pkg/templating"
	"github.com/sebastianRau/deployer/pkg/verbose"
)

func ReplaceAll(src string, dest string, repalcements map[string]string, v io.Writer) error {
	verbose.Fprintf(v, "Replace All started\n")

	verbose.Fprintf(v, "read File: %s\n", src)
	lines, err := readFile(src)
	if err != nil {
		return err
	}
	verbose.Fprintf(v, "got %d lines\n", len(lines))

	for i := range lines {
		lines[i] = replaceText(lines[i], repalcements)

	}

	verbose.Fprintf(v, "writing  file: %s\n", dest)
	err = writeFile(dest, strings.Join(lines, "\n"))

	if err != nil {
		return err
	}

	return nil
}

func ParseTemplate(src string, data string, dest string, v io.Writer) error {
	verbose.Fprintf(v, "parsing template started\n")
	verbose.Fprintf(v, "read template: %s\n", src)
	verbose.Fprintf(v, "read data: %s\n", data)

	lines, err := templating.ParseTemplateJsonData(src, data)
	if err != nil {
		return err
	}

	verbose.Fprintf(v, "got %d lines\n", strings.Count(lines, "\n"))

	verbose.Fprintf(v, "writing file: %s\n", dest)
	err = writeFile(dest, lines)
	if err != nil {
		return err
	}

	return nil
}

func WriteFile(dst string, text []string, v io.Writer) error {
	s := strings.Join(text, "\n")

	verbose.Fprintf(v, "Replace All started\n")
	verbose.Fprintf(v, "got %d lines\n", strings.Count(s, "\n"))
	verbose.Fprintf(v, "writing file: %s\n", dst)

	return writeFile(dst, s)
}

func readFile(src string) ([]string, error) {
	var fileLines []string
	readFile, err := os.Open(src)
	defer readFile.Close()

	if err != nil {
		return fileLines, err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	readFile.Close()

	return fileLines, nil
}

func writeFile(dest string, text string) error {
	file, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(text)
	file.Sync()
	file.Close()

	return err
}

func replaceText(s string, repalcements map[string]string) string {
	for k, v := range repalcements {
		s = strings.ReplaceAll(s, k, v)
	}
	return s
}
