package config

import (
	"fmt"
	"os"
	"testing"
)

func TestGetAppDirPath(t *testing.T) {
	dir, _ := GetAppDirPath()

	expectedDir, _ := os.Getwd()

	if expectedDir != fmt.Sprintf("%s/core/config", dir) {
		t.Fatalf("Expected path %s, but got %s", expectedDir, dir)
	}
}

func TestGetDataDirPath(t *testing.T) {
	appDir, _ := GetAppDirPath()
	dir := GetDataDirPath()

	expectedDir := fmt.Sprintf("%s/data", appDir)

	if expectedDir != dir {
		t.Fatalf("Expected path %s, but got %s", expectedDir, dir)
	}
}

func TestGetTargetDotEnvFilePath(t *testing.T) {
	path := GetTargetDotEnvFilePath()

	expectedPath := fmt.Sprintf("%s/.env", GetDataDirPath())

	if expectedPath != path {
		t.Fatalf("Expected path %s, but got %s", expectedPath, path)
	}
}
