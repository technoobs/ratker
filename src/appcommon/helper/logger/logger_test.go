package logger

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

const (
	logName = "test_log.txt"
)

func TestSetup(t *testing.T) {
	// Environment Setup
	workDir, err := os.Getwd()
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Work dir is %s\n", workDir)
	logPath := workDir + "/" + logName

	// Operations
	appLogger := Setup(logPath)

	// Assert
	if appLogger == nil {
		t.Errorf("Expected appLogger to be created")
	}
}

func TestWriteInfo(t *testing.T) {
	// Environment Setup
	workDir, err := os.Getwd()
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Work dir is %s\n", workDir)
	logPath := workDir + "/" + logName

	// Operations
	appLogger := Setup(logPath)
	appLogger.WriteInfo("Test Info")

	// Assert
	file, err := os.Open(logPath)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	defer file.Close()

	for _, eachline := range txtlines {
		if !strings.Contains(eachline, "Test Info") {
			t.Error("Failed to find wanted logging info")
		}
	}
}

func TestWriteWarn(t *testing.T) {
	// Environment Setup
	workDir, err := os.Getwd()
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Work dir is %s\n", workDir)
	logPath := workDir + "/" + logName

	// Operations
	appLogger := Setup(logPath)
	appLogger.WriteWarn("Test Warn")

	// Assert
	file, err := os.Open(logPath)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	defer file.Close()

	for _, eachline := range txtlines {
		if !strings.Contains(eachline, "Test Warn") {
			t.Error("Failed to find wanted logging warning")
		}
	}
}

func TestWriteError(t *testing.T) {
	// Environment Setup
	workDir, err := os.Getwd()
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Work dir is %s\n", workDir)
	logPath := workDir + "/" + logName

	// Operations
	appLogger := Setup(logPath)
	appLogger.WriteError("Test Error")

	// Assert
	file, err := os.Open(logPath)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	defer file.Close()

	for _, eachline := range txtlines {
		if !strings.Contains(eachline, "Test Error") {
			t.Error("Failed to find wanted logging warning")
		}
	}
}
