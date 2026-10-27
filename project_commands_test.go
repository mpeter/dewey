package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestProjectCommands_NoUFInitializer(t *testing.T) {
	foundDeweyIndex := false
	for _, directory := range []string{"command", "commands"} {
		directory = filepath.Join(".opencode", directory)
		entries, err := os.ReadDir(directory)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			t.Fatalf("read project command directory %s: %v", directory, err)
		}
		for _, entry := range entries {
			if entry.IsDir() || filepath.Ext(entry.Name()) != ".md" {
				continue
			}
			name := strings.TrimSuffix(entry.Name(), ".md")
			if name == "uf-init" {
				t.Errorf("retired UF customizer is still advertised in %s", directory)
			}
			if name == "dewey-index" {
				foundDeweyIndex = true
			}
		}
	}
	if !foundDeweyIndex {
		t.Error("Dewey's dewey-index command is missing from the project catalog")
	}
}
