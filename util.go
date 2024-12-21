package main

import (
	"fmt"
	"strings"
)

func Join[T fmt.Stringer](slice []T, sep string) string {
	switch len(slice) {
	case 0:
		return ""
	case 1:
		return slice[0].String()
	}

	var sb strings.Builder

	sb.WriteString(slice[0].String())
	for _, p := range slice[1:] {
		sb.WriteString(sep)
		sb.WriteString(p.String())
	}

	return sb.String()
}
