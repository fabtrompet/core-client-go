package coreclient

import (
	"os"
	"testing"

	gojson "encoding/json"

	"github.com/datarhei/core-client-go/v16/api"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/require"
)

func TestParseProcessID(t *testing.T) {
	tests := map[string]ProcessID{
		"foo":         NewProcessID("foo", ""),
		"foo@":        NewProcessID("foo", ""),
		"foo@bar":     NewProcessID("foo", "bar"),
		"foo@bar@bar": NewProcessID("foo@bar", "bar"),
	}

	for pid, id := range tests {
		ppid := ParseProcessID(pid)

		require.Equal(t, id, ppid, pid)
	}
}

func BenchmarkProcessList(b *testing.B) {
	data, err := os.ReadFile("./fixtures/processList.json")
	require.NoError(b, err)

	for i := 0; i < b.N; i++ {
		processes := []api.Process{}
		err := json.Unmarshal(data, &processes)
		require.NoError(b, err)
	}
}

func BenchmarkProcessListNativJSON(b *testing.B) {
	data, err := os.ReadFile("./fixtures/processList.json")
	require.NoError(b, err)

	for i := 0; i < b.N; i++ {
		processes := []api.Process{}
		err := gojson.Unmarshal(data, &processes)
		require.NoError(b, err)
	}
}
