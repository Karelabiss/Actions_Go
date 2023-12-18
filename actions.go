package Actions

import "fmt"

// Print Line for Success of process. "arg" as message of type string.
func Success(arg string) {
	fmt.Println("[+] - " + arg)
}

// Print line for Failure of process. "arg" as message of type string.
func Failure(arg string) {
	fmt.Println("[-] - " + arg)
}

// Print Line for Error of Process. "arg" as message of type string.
func Error(arg string) {
	fmt.Println("[!] - " + arg)
}

// Print Line for Info of Process. "arg" as message of type string.
func Info(arg string) {
	fmt.Println("[?] - " + arg)
}

// Print Line for Debug of Process. "arg" as message of type string.
func Debug(arg string) {
	fmt.Println("[DEBUG] - " + arg)
}
