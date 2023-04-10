package color

import "runtime"

var Reset = "\033[0m"
var Black = "\033[0;30m"
var Red = "\033[0;31m"
var Green = "\033[0;32m"
var Orange = "\033[0;33m"
var Blue = "\033[0;34m"
var Purple = "\033[0;35m"
var Cyan = "\033[0;36m"
var LightGray = "\033[0;37m"
var DarkGray = "\033[1;30m"
var LightRed = "\033[1;31m"
var LightGreen = "\033[1;32m"
var Yellow = "\033[1;33m"
var LightBlue = "\033[1;34m"
var LightPurple = "\033[1;35m"
var LightCyan = "\033[1;36m"
var White = "\033[1;37m"

func init() {
	if runtime.GOOS == "windows" {
		Disable()
	}
}

func Disable() {
	Reset = ""
	Black = ""
	Red = ""
	Green = ""
	Orange = ""
	Blue = ""
	Purple = ""
	Cyan = ""
	LightGray = ""
	DarkGray = ""
	LightRed = ""
	LightGreen = ""
	Yellow = ""
	LightBlue = ""
	LightPurple = ""
	LightCyan = ""
	White = ""
}
