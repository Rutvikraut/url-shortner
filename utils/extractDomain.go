package utils

import (
	"fmt"
	"strings"
)

func ExtractDomain(url string) string {
	parts := strings.Split(url, "/")
	fmt.Println(parts[2])
	if len(parts) > 2 {
		return parts[2]
	}
	return ""
}
