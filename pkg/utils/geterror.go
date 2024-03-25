package utils

import "strings"

func ExtractError(message string) string{
	parts:=strings.Split(message,"desc =")
	if len(parts) >1{
		return strings.TrimSpace(parts[1])
	}
	return message
}