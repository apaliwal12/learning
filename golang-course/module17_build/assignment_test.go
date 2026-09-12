package main

import (
	"strings"
	"testing"
)

func TestLinuxBuildCommand(t *testing.T) {
	cmd := LinuxBuildCommand()
	
	if !strings.Contains(cmd, "GOOS=linux") {
		t.Error("Command missing GOOS=linux")
	}
	if !strings.Contains(cmd, "GOARCH=amd64") {
		t.Error("Command missing GOARCH=amd64")
	}
	if !strings.Contains(cmd, "go build") {
		t.Error("Command missing 'go build'")
	}
	if !strings.Contains(cmd, "-o server") {
		t.Error("Command missing '-o server'")
	}
}

func TestLDFlagsCommand(t *testing.T) {
	cmd := LDFlagsCommand()
	
	if !strings.Contains(cmd, "-ldflags") {
		t.Error("Command missing '-ldflags'")
	}
	if !strings.Contains(cmd, "-X") {
		t.Error("Command missing '-X'")
	}
	if !strings.Contains(cmd, "main.Version=v1.5.0") && !strings.Contains(cmd, "main.Version='v1.5.0'") {
		t.Error("Command missing proper variable injection for main.Version")
	}
}
