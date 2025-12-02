package winruntimesaudit

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDotNetRuntimesAudit(t *testing.T) {
	dotnet, err := DotNetRuntimesAuditResult()
	require.NoError(t, err)
	require.NotNil(t, dotnet)
	for _, d := range dotnet {
		t.Logf("Version: %s, Installed: %t\nPath: %s", d.Version, d.Installed, d.Path)
		require.NotEmpty(t, d.Version)
		require.NotEmpty(t, d.Path)
	}

}
