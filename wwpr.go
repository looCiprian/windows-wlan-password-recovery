package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
	"time"
)

//Cross compiling for Windows GOOS=windows GOARCH=386 go build -o wwpr.exe wwpr.go
func main() {
	if checkWindowsOS() {
		fmt.Println("Windows OS detected")

		profiles := getProfiles()
		profileKey := getPassword(profiles)
		fmt.Println(profileKey)
		fmt.Println(len(profileKey))
		if len(profileKey) != 0 {
			printProfileKey(profileKey)
		} else {
			printAll(profiles)
		}

	} else {
		fmt.Println("No Windows OS detected")
	}
	fmt.Println("Press enter to close")
	fmt.Scanf("h")
}

func checkWindowsOS() bool {
	if runtime.GOOS == "windows" {
		return true
	}
	return false
}

func getProfiles() []string {
	cmd := exec.Command("cmd", "/C", "netsh.exe", "wlan", "show", "profiles")
	out, _ := cmd.CombinedOutput()

	var profiles []string
	var regex = regexp.MustCompile(`.*: (.*)`)
	var result = regex.FindAllStringSubmatch(string(out), -1)
	for _, s := range result {
		//fmt.Println(s)
		for j := 1; j < len(s); j++ {
			profiles = append(profiles, s[j][:len(s[j])-1])
		}
	}
	return profiles

}

func getPassword(profiles []string) map[string]string {
	name := ""
	var keys = make(map[string]string)
	for _, profile := range profiles {
		name = "name=" + profile
		cmd := exec.Command("cmd", "/C", "netsh.exe", "wlan", "show", "profiles", name, "key=clear")
		out, _ := cmd.CombinedOutput()

		var regex = regexp.MustCompile(`Key Content.*: (.*)`)
		var result = regex.FindAllStringSubmatch(string(out), -1)
		for _, s := range result {
			//fmt.Println(s)
			for j := 1; j < len(s); j++ {
				//keys = append(keys, s[j][:len(s[j])-1])
				keys[profile] = s[j][:len(s[j])-1]
			}
		}

	}
	return keys
}

func printProfileKey(profileKey map[string]string) {
	for profile, key := range profileKey {
		fmt.Printf("Network Name: %s \t Password: %s\n", profile, key)
	}
}

func printAll(profiles []string) {
	fmt.Println("No English terminal detected...")
	fmt.Println("Printing all raw data...")
	time.Sleep(1 * time.Second)

	for _, profile := range profiles {
		name := "name=" + profile
		cmd := exec.Command("cmd", "/C", "netsh.exe", "wlan", "show", "profiles", name, "key=clear")
		out, _ := cmd.CombinedOutput()
		fmt.Println(string(out))
	}
}
