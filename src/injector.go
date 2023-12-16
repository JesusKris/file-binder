package main

import (
	"injector/logger"
	"injector/modules"
	"injector/validation"
)

func main() {

	os := validation.CheckOS()

	targetName, executableNames := validation.GetArgs()

	validation.ValidateArgs(targetName, executableNames, os)

	excutablesData := modules.GetExecutableData(executableNames)

	modules.CreateMother(targetName)

	modules.AppendToMother(targetName, excutablesData)

	logger.Log("Successfully combined " + modules.IntToString(len(executableNames)) + " executables.")
}
