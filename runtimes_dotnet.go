package winruntimesaudit

import (
	"bufio"
	"errors"
	"os/exec"
	"strings"
)

// DotNetRuntime represents a .NET runtime installation on a Windows system.
type DotNetRuntime struct {
	Version   string
	Installed bool
	Path      string
}

// DotNetRuntimesAuditResult holds the results of auditing .NET runtimes.
// It uses dotnet --list-runtimes to gather information about installed runtimes.
func DotNetRuntimesAuditResult() ([]DotNetRuntime, error) {
	cmd := exec.Command("dotnet", "--list-runtimes")
	output, err := cmd.Output()
	if err != nil {
		return nil, errors.New("failed to execute dotnet --list-runtimes: " + err.Error())
	}

	var runtimes []DotNetRuntime
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) < 3 {
			return nil, errors.New("unexpected output format from dotnet --list-runtimes: " + line)
		}
		// Assume format: Name Version [Path]
		version := parts[1]
		path := strings.Trim(parts[2], "[]")
		runtime := DotNetRuntime{
			Version:   version,
			Installed: true,
			Path:      path,
		}
		runtimes = append(runtimes, runtime)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return runtimes, nil
}
