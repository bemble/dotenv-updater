package dotenv

import (
	"os"
	"testing"
)

const TEST_FILE = "dotenv_test"

func beforeTestCase() {
	d1 := []byte("VAR1=\"foo\"\nVAR2=\"bar\"\n")
	err := os.WriteFile(TEST_FILE, d1, 0644)
	if err != nil {
		panic(err)
	}
}

func afterTestCase() {
	err := os.Remove(TEST_FILE)
	if err != nil {
		panic(err)
	}
}

func TestSetVar(t *testing.T) {
	beforeTestCase()
	defer afterTestCase()

	SetVar(TEST_FILE, "VAR3", "foobar")

	expectedContent := "VAR1=\"foo\"\nVAR2=\"bar\"\nVAR3=\"foobar\"\n"

	content, err := os.ReadFile(TEST_FILE)

	if err != nil {
		panic(err)
	}

	if expectedContent != string(content) {
		t.Fatalf("expected content %s, but got %s", expectedContent, content)
	}
}

func TestUpdateVar(t *testing.T) {
	beforeTestCase()
	defer afterTestCase()

	SetVar(TEST_FILE, "VAR3", "foobar")
	SetVar(TEST_FILE, "VAR2", "foo")

	expectedContent := "VAR1=\"foo\"\nVAR2=\"foo\"\nVAR3=\"foobar\"\n"

	content, err := os.ReadFile(TEST_FILE)

	if err != nil {
		panic(err)
	}

	if expectedContent != string(content) {
		t.Fatalf("expected content %s, but got %s", expectedContent, content)
	}
}

func TestDeleteVar(t *testing.T) {
	beforeTestCase()
	defer afterTestCase()

	DeleteVar(TEST_FILE, "VAR1")

	expectedContent := "VAR2=\"bar\"\n"

	content, err := os.ReadFile(TEST_FILE)

	if err != nil {
		panic(err)
	}

	if expectedContent != string(content) {
		t.Fatalf("expected content %s, but got %s", expectedContent, content)
	}
}
func TestDeleteNotSetVar(t *testing.T) {
	beforeTestCase()
	defer afterTestCase()

	DeleteVar(TEST_FILE, "VAR_UNSET")

	expectedContent := "VAR1=\"foo\"\nVAR2=\"bar\"\n"

	content, err := os.ReadFile(TEST_FILE)

	if err != nil {
		panic(err)
	}

	if expectedContent != string(content) {
		t.Fatalf("expected content %s, but got %s", expectedContent, content)
	}
}
