package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		return
	}

	goFilename := args[1]

	printf("go file is %s", goFilename)

	if !strings.HasSuffix(goFilename, ".go") {
		return
	}

	execName := strings.Replace(goFilename, ".go", "", -1)
	winName := strings.Replace(goFilename, ".go", ".exe", -1)

	// build mac
	cmd := exec.Command("go", "build", "-i", goFilename)
	_ = cmd.Run()

	gzFile(execName, execName+"_"+"darwin")

	// build linux
	cmd2 := exec.Command("go", "build", goFilename)

	_ = os.Setenv("CGO_ENABLED", "0")
	_ = os.Setenv("GOOS", "linux")
	_ = os.Setenv("GOARCH", "amd64")
	cmd2.Stdout = os.Stdout
	_ = cmd2.Run()

	gzFile(execName, execName+"_"+"linux")

	_ = os.Setenv("CGO_ENABLED", "0")
	_ = os.Setenv("GOOS", "windows")
	_ = os.Setenv("GOARCH", "amd64")
	cmd3 := exec.Command("go", "build", "-i", goFilename)
	_ = cmd3.Run()

	gzFile(winName, execName+"_"+"windows")

	_ = os.Remove(execName)
	_ = os.Remove(winName)
}

func gzFile(fileName, targetName string) {
	cmd := exec.Command("tar", "cvjf", targetName+".tar.gz", fileName)
	_ = cmd.Run()
}

func printf(format string, a ...interface{}) {
	outputString := fmt.Sprintf(format, a...)
	_, _ = fmt.Fprintf(os.Stdout, outputString+"\n")
}
