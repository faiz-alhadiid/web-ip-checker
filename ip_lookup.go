package main

import (
	"os/exec"
	"regexp"
	"strings"
)

type digResult struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func getDigResult(endpoint string) (string, error) {
	cmd := exec.Command("dig", endpoint)
	b, err := cmd.Output()

	return string(b), err
}
func parseDigResult(rawDig string) (result []digResult) {
	re := regexp.MustCompile(`(?s);; ANSWER SECTION:(.*?);;`)
	matched := re.FindString(rawDig)

	matched = strings.ReplaceAll(matched, ";; ANSWER SECTION:", "")
	matched = strings.ReplaceAll(matched, ";;", "")
	matched = strings.TrimSpace(matched)

	list := strings.Split(matched, "\n")
	for _, val := range list {
		splitted := strings.Split(val, "\t")

		if len(splitted) < 2 {
			continue
		}

		result = append(result, digResult{
			Type:  splitted[len(splitted)-2],
			Value: splitted[len(splitted)-1],
		})
	}

	return result

}
