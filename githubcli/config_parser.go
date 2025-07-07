package githubcli

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

//go:generate faux --interface ConfigParser --output fakes/config_parser.go
type ConfigParser interface {
	Parse(path string) (hasGhConfig bool, err error)
}

type ConfigParserImpl struct{}

func NewConfigParser() ConfigParser {
	return ConfigParserImpl{}
}

func (p ConfigParserImpl) Parse(path string) (bool, error) {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, fmt.Errorf("failed to parse requirements file: %w", err)
	}
	defer file.Close()

	// Look for GitHub CLI related requirements
	ghRe := regexp.MustCompile(`(?i)(github-cli|gh-cli|gh)`)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if ghRe.MatchString(line) {
			return true, nil
		}
	}

	return false, nil
} 