package util

import (
	"fmt"
	"regexp"
)

func WrapURLWithAnchor(text string) string {
	re := regexp.MustCompile(`((https?):\/\/[\w!?/+\-_~;.,*&@#$%()'\[\]]+)`)
	return re.ReplaceAllStringFunc(text, func(url string) string {
		return fmt.Sprintf(`<a href="%s" target="_blank">%s</a>`, url, url)
	})
}
