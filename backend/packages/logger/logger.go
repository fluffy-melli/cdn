package logger

import (
	"fmt"
	"strconv"
	"time"
)

func LOAD_Cache_FILE(file string, i, max int, loadbyte, maxbyte int64, elapsed time.Duration) {
	filei := Fg_BrightGreen + "[LOAD] " + Fg_BrightYellow + ShortenString(file, 15) + Fg_BrightBlue + ShortenString("["+fmt.Sprintf("%d", i+1)+"/"+fmt.Sprintf("%d", max)+"] ", 12)
	bytes := ShortenString("("+ByteToLargestSize(loadbyte)+"/"+ByteToLargestSize(maxbyte)+")", 25)
	paddingLength := len("(1023 Bytes/1023 Bytes)") - len(bytes) + 1
	padding := ""
	if paddingLength > 0 {
		padding = fmt.Sprintf("%*s", paddingLength, "")
	}
	dls := padding + Fg_BrightCyan + " " + TimeToLargestUnit(elapsed.Nanoseconds())
	fmt.Println(filei + bytes + dls)
}

func PASS_Cache_FILE(file string, i, max int, loadbyte, maxbyte int64, elapsed time.Duration) {
	filei := Fg_BrightMagenta + "[PASS] " + Fg_BrightYellow + ShortenString(file, 15) + Fg_BrightBlue + ShortenString("["+fmt.Sprintf("%d", i+1)+"/"+fmt.Sprintf("%d", max)+"] ", 12)
	bytes := ShortenString("("+ByteToLargestSize(loadbyte)+"/"+ByteToLargestSize(maxbyte)+")", 25)
	paddingLength := len("(1023 Bytes/1023 Bytes)") - len(bytes) + 1
	padding := ""
	if paddingLength > 0 {
		padding = fmt.Sprintf("%*s", paddingLength, "")
	}
	dls := padding + Fg_BrightCyan + " " + TimeToLargestUnit(elapsed.Nanoseconds())
	fmt.Println(filei + bytes + dls)
}

func ERR_Cache_FILE(file string, i, max int, loadbyte, maxbyte int64, elapsed time.Duration) {
	filei := Fg_BrightRed + "[ERR]  " + Fg_BrightYellow + ShortenString(file, 15) + Fg_BrightBlue + ShortenString("["+fmt.Sprintf("%d", i+1)+"/"+fmt.Sprintf("%d", max)+"] ", 12)
	bytes := ShortenString("("+ByteToLargestSize(loadbyte)+"/"+ByteToLargestSize(maxbyte)+")", 25)
	paddingLength := len("(1023 Bytes/1023 Bytes)") - len(bytes) + 1
	padding := ""
	if paddingLength > 0 {
		padding = fmt.Sprintf("%*s", paddingLength, "")
	}
	dls := padding + Fg_BrightCyan + " " + TimeToLargestUnit(elapsed.Nanoseconds()) + Reset
	fmt.Println(filei + bytes + dls)
}

func END_Cache(cachec, passc int, size int64) {
	msg := Fg_BrightYellow + ShortenString("["+fmt.Sprintf("%d", cachec)+"/"+fmt.Sprintf("%d", cachec+passc)+"] | (Cached/All) | "+ByteToLargestSize(size), 70) + Reset
	//dls := Fg_BrightBlue + "(" + TimeToLargestUnit(elapsed.Nanoseconds()) + ")" + Reset
	fmt.Println(msg)
}

func READ_FILE(file string, cacheing bool, loadbyte int64, elapsed time.Duration) {
	filei := Fg_BrightGreen + "[LOAD] " + Fg_BrightYellow + ShortenString(file, 15) + Fg_BrightBlue + ShortenString("["+strconv.FormatBool(cacheing)+"] ", 12)
	bytes := ShortenString("("+ByteToLargestSize(loadbyte)+") ", 25) + " "
	dls := Fg_BrightCyan + TimeToLargestUnit(elapsed.Nanoseconds()) + Reset
	fmt.Println(filei + bytes + dls)
}

func UPLOAD_FILE(file string, i, max int, loadbyte, maxbyte int64, elapsed time.Duration) {
	filei := Fg_BrightBlue + "[UPLD] " + Fg_BrightYellow + ShortenString(file, 15) + Fg_BrightBlue + ShortenString("["+fmt.Sprintf("%d", i+1)+"/"+fmt.Sprintf("%d", max)+"] ", 12)
	bytes := ShortenString("("+ByteToLargestSize(loadbyte)+"/"+ByteToLargestSize(maxbyte)+")", 25)
	paddingLength := len("(1023 Bytes/1023 Bytes)") - len(bytes) + 1
	padding := ""
	if paddingLength > 0 {
		padding = fmt.Sprintf("%*s", paddingLength, "")
	}
	dls := padding + Fg_BrightCyan + " " + TimeToLargestUnit(elapsed.Nanoseconds()) + Reset
	fmt.Println(filei + bytes + dls)
}

func ERR_READ_FILE(file string, cacheing bool, loadbyte int64, elapsed time.Duration) {
	filei := Fg_BrightRed + "[ERR]  " + Fg_BrightYellow + ShortenString(file, 15) + Fg_BrightBlue + ShortenString("["+strconv.FormatBool(cacheing)+"] ", 12)
	bytes := ShortenString("("+ByteToLargestSize(loadbyte)+") ", 25) + " "
	dls := Fg_BrightCyan + TimeToLargestUnit(elapsed.Nanoseconds()) + Reset
	fmt.Println(filei + bytes + dls)
}
