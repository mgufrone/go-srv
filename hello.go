package main

import (
	"fmt"
	"regexp"
	"strings"
)

type Hello struct {
}

func (h *Hello) Greet(name string) string {
	numberRegex := regexp.MustCompile(`\d`)
	result := fmt.Sprintf("Hello, %s", strings.Title(name))
	if numberRegex.MatchString(name) {
		result = fmt.Sprintf("%s. A weird name, but okay.", result)
	}
	return result
}
