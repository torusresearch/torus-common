package logging

import (
	"strings"
)

var (
	colorOff    = "\033[0m"
	colorRed    = "\033[0;31m"
	colorGreen  = "\033[0;32m"
	colorOrange = "\033[0;33m"
	colorBlue   = "\033[0;34m"
	colorPurple = "\033[0;35m"
	colorCyan   = "\033[0;36m"
	colorGray   = "\033[0;37m"
)

// mixer mix the color on and off byte with the actual data
func mixer(color string, data string) string {
	var result strings.Builder
	result.WriteString(color)
	result.WriteString(data)
	result.WriteString(colorOff)
	return result.String()
}

func debugMixer(message string) string {
	return mixer(colorPurple, message)
}

func errorMixer(message string) string {
	return mixer(colorRed, message)
}

func fatalMixer(message string) string {
	return mixer(colorRed, message)
}

func warningMixer(message string) string {
	return mixer(colorOrange, message)
}

func infoMixer(message string) string {
	return mixer(colorGreen, message)
}

func traceMixer(message string) string {
	return mixer(colorCyan, message)
}
