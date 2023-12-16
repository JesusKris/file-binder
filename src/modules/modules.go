package modules

import (
	"errors"
	"injector/logger"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

func GetExecutableData(names []string) [][]byte {

	logger.Log("Getting executables data...")

	var data [][]byte

	for _, name := range names {
		data = append(data, []byte("h$lL@ w@rLd"))
		data = append(data, GetFileContents(name))

	}
		
	return data
}

func OpenFile(name string) *os.File {

	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Fatal(errors.New("Failed to open file " + name + ":" + err.Error()))
	}

	return file
}

func WriteToFile(name string, file *os.File, content []byte) {

	_, err := file.Write(content)

	if err != nil {
		logger.Fatal(errors.New("Failed to write to file " + name + ":" + err.Error()))

	}

}

func RemoveFile(name string) {

	err := os.Remove(name)
	if err != nil {
		logger.Fatal(errors.New("Failed to delete a file " + name + ":" + err.Error()))

	}
}

func GetFileContents(name string) []byte {

	contents, err := ioutil.ReadFile(name)
	if err != nil {
		logger.Fatal(errors.New("Failed to the contents of file " + name + ":" + err.Error()))
	}
	return contents
}

func StripExtension(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}

func IntToString(nr int) string {
	return strconv.Itoa(nr)
}
