package validation

import (
	"errors"
	"fmt"
	"injector/logger"
	"os"
	"runtime"
)

func CheckOS() string {
	logger.Log("Checking OS...")

	switch runtime.GOOS {

	case "windows":
		return "windows"
	case "linux":
		return "linux"
	case "ios":
		return "ios"

	default:
		logger.Fatal(errors.New("Either failed to detect OS or you are on a nonsupported OS system. Supported OS's are Windows, Linux and iOS."))
		return ""
	}
}

func GetArgs() (string, []string) {

	logger.Log("Getting args...")

	if len(os.Args) < 4 {
		logger.Fatal(errors.New("Not enough arguments. You must provide atleast 2 executables. Usage: go run injector.go [TARGET NAME] [EXECUTABLE] [EXECUTABLE] EX: go run injector.go helloworld.exe hello.exe world.exe"))
	}

	var executableNames []string

	for i := 2; i < len(os.Args); i++ {
		executableNames = append(executableNames, os.Args[i])
	}

	return (os.Args[1]), executableNames
}

func ValidateArgs(targetName string, executableNames []string, os string) {

	logger.Log("Validating args...")

	for _, name := range executableNames {
		if !fileExists(name) {
			logger.Fatal(errors.New("File " + name + " does not exist. Please make sure that a file named " + name + " exists."))
		}
	}

	if fileExists(targetName + getOsExecutableExtension(os)) {
		logger.Warning("File " + targetName + " already exists. This file would be overwritten with new data. Are you sure you want to continue? (y/n)")
		var answer string

		for answer != "y" && answer != "n" {
			fmt.Scanln(&answer)
		}

		if answer == "n" {
			logger.Fatal(errors.New("Program terminated..."))
		}
	}

}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !errors.Is(err, os.ErrNotExist)
}

func getOsExecutableExtension(os string) string{
	if os == "windows" {
		return ".exe"
	}

	return ""
}
