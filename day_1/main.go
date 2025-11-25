package day1

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

var userName string = "Motadata User"

func getTimeGreeting() string {
	hour := time.Now().Hour()
	
	switch {
	case hour < 12:
		return "Good morning"
	case hour < 17:
		return "Good afternoon"
	case hour < 21:
		return "Good evening"
	default:
		return "Good night"
	}
}

func getEnvironmentInfo() map[string]string {
	info := make(map[string]string)

	hostname, _ := os.Hostname()
	wd, _ := os.Getwd()
	user := os.Getenv("USER")
	if user == "" {
		user = os.Getenv("USERNAME")
	}
	if user == "" {
		user = "Unknown"
	}

	info["Go Version"] = runtime.Version()
	info["Platform"] = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
	info["Hostname"] = hostname
	info["Current Directory"] = wd
	info["Timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	info["CPU Cores"] = fmt.Sprintf("%d", runtime.NumCPU())
	info["Username"] = user
	userName = user
	
	return info
}

func printBanner() {
	banner := "Hello, Motadata!"
	fmt.Println(banner)
}

func printEnvironmentInfo() {
	fmt.Printf("\nEnvironment Information:\n\n")

	info := getEnvironmentInfo()
	for key, value := range info {
		fmt.Printf("  • %-20s: %s\n", key, value)
	}
}

func performGreeting() {
	greeting := getTimeGreeting()
	fmt.Printf("\n%s %s!\n", greeting, userName)
}

func Day1() {
	printBanner()
	printEnvironmentInfo()
	performGreeting()
}
