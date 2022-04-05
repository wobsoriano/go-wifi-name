package wifiname

import (
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)


func WifiName() string {
	platform := runtime.GOOS

	switch platform {
		case "darwin":
			return forDarwin()
		case "linux":
			return forLinux()
		case "win32":
			return forWindows()
		default:
			panic("Not supported")
	}
}

func forDarwin() string {
	cmd := "/System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport"
	args := "-I"
	rawOutput, err := exec.Command(cmd, args).Output()
	panicIf(err)
	
	output := string(rawOutput)
	re := regexp.MustCompile(`(?m)^\s*SSID: (.+)\s*$`)
	matches := re.FindAllStringSubmatch(output, -1)

	if len(matches) == 0 {
		panic("Could not get SSID")
	}

	output = matches[0][1]

	return output
}

func forLinux() string {
	cmd := "iwgetid"
	args := "--raw"
	rawOutput, err := exec.Command(cmd, args).Output()
	panicIf(err)
	
	output := string(rawOutput)
	output = strings.Replace(output, "\n", "", -1)

	if len(output) == 0 {
		panic("Could not get SSID")
	}

	return output
}

func forWindows() string {
	rawOutput, err := exec.Command("netsh", "wlan", "show", "interface").Output()
	panicIf(err)
	
	output := string(rawOutput)
	re := regexp.MustCompile(`(?m)^\s*SSID\s*: (.+)\s*$`)
	matches := re.FindAllStringSubmatch(output, 1)

	if len(matches) == 0 {
		panic("Could not get SSID")
	}

	output = matches[0][1]

	return output
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}