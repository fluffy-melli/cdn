package logger

import (
	"fmt"
	"strings"
)

func ByteToLargestSize(bytes int64) string {
	if bytes >= 1024*1024*1024 {
		return fmt.Sprintf("%.2f GB", float64(bytes)/(1024*1024*1024))
	} else if bytes >= 1024*1024 {
		return fmt.Sprintf("%.2f MB", float64(bytes)/(1024*1024))
	} else if bytes >= 1024 {
		return fmt.Sprintf("%.2f KB", float64(bytes)/1024)
	}
	return fmt.Sprintf("%d Bytes", bytes)
}

func TimeToLargestUnit(nanoseconds int64) string {
	if nanoseconds >= 1_000_000_000 {
		return fmt.Sprintf("%.2fs", float64(nanoseconds)/1_000_000_000)
	} else if nanoseconds >= 1_000_000 {
		return fmt.Sprintf("%.2fms", float64(nanoseconds)/1_000_000)
	} else if nanoseconds >= 1_000 {
		return fmt.Sprintf("%.2fµs", float64(nanoseconds)/1_000)
	}
	return fmt.Sprintf("%dns", nanoseconds)
}

func ShortenString(s string, maxLength int) string {
	if len(s) <= maxLength {
		padding := maxLength - len(s)
		leftPadding := padding / 2
		rightPadding := padding - leftPadding
		return Padding(leftPadding) + s + Padding(rightPadding)
	}
	half := (maxLength - 3) / 2
	return s[:half] + "..." + s[len(s)-half:]
}

func Padding(len int) string {
	padding := ""
	if len > 0 {
		padding = fmt.Sprintf("%*s", len, "")
	}
	return padding
}

func LineString(s string, maxLength int) string {
	if len(s) <= maxLength {
		padding := maxLength - len(s)
		leftPadding := padding / 2
		rightPadding := padding - leftPadding
		return LinePadding(leftPadding) + s + LinePadding(rightPadding)
	}
	half := (maxLength - 3) / 2
	return s[:half] + "..." + s[len(s)-half:]
}

func LinePadding(length int) string {
	if length > 0 {
		return strings.Repeat("-", length)
	}
	return ""
}
