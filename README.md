# winruntimesaudit

A Go package for auditing installed Visual C++ Redistributable and .NET runtimes on Windows systems.

## Installation

To install the package, run:

```bash
go get github.com/miroslav-matejovsky/winruntimesaudit
```

## Requirements

- Go 1.25.4 or later
- Windows operating system
- For .NET runtime auditing: .NET SDK installed and available in PATH

## Usage

```go
package main

import (
    "fmt"
    "log"

    "github.com/miroslav-matejovsky/winruntimesaudit"
)

func main() {
    result, err := winruntimesaudit.DoAudit()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("VC Redistributables:")
    for _, vc := range result.VCRedistRuntimes {
        status := "not installed"
        if vc.Installed {
            status = "installed"
        }
        fmt.Printf("- %s (%s): %s at %s\n", vc.Version, vc.Architecture, status, vc.Path)
    }

    fmt.Println("\n.NET Runtimes:")
    for _, dotnet := range result.DotNetRuntimes {
        fmt.Printf("- %s %s at %s\n", dotnet.Type, dotnet.Version, dotnet.Location)
    }
}
```

## API

### `DoAudit() (*AuditResult, error)`

Performs a complete audit of both VC Redistributables and .NET runtimes.

Returns an `AuditResult` struct containing:

- `VCRedistRuntimes`: Slice of `VCRedistRuntime` structs
- `DotNetRuntimes`: Slice of `DotNetRuntime` structs

### `VCRedistRuntime`

Represents a Visual C++ Redistributable runtime:

- `Version`: Version string (e.g., "14.40.33810.00")
- `Architecture`: Architecture string ("x86" or "x64")
- `Installed`: Boolean indicating installation status
- `Path`: Registry path where the runtime is registered

### `DotNetRuntime`

Represents a .NET runtime:

- `Type`: Runtime type (e.g., "Microsoft.AspNetCore.App", "Microsoft.NETCore.App")
- `Version`: Version string
- `Location`: Installation path

## Building and Testing

This project uses [Mage](https://magefile.org/) for build automation.

### Prerequisites

- Install Mage: `go install github.com/magefile/mage@latest`
- Install golangci-lint: Follow instructions at <https://golangci-lint.run/usage/install/>
- Install gotestsum: `go install gotest.tools/gotestsum@latest`

### Commands

- `mage tidy`: Clean up go.mod and go.sum
- `mage fmt`: Format code with gofmt
- `mage vet`: Run go vet
- `mage lint`: Run golangci-lint
- `mage test`: Run tests with gotestsum
- `mage all`: Run all quality checks and tests

## Contributing

Contributions are welcome! Please ensure that:

- Code is formatted with `gofmt`
- All tests pass
- New functionality includes appropriate tests
- Code passes linting checks

## License

This project is licensed under the MIT License.
