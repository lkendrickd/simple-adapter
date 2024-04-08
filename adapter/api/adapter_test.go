package adapter

import (
	"testing"
)

func TestAdapter(t *testing.T) {
	// create a new instance of the legacy api.
	legacyAPI := NewRecordsAPI("record1", "record2", "record3")

	// create a new instance of the modern api.
	modernAPI := NewEntriesAPI()

	apiAdapter := NewAdapter(legacyAPI, modernAPI)

	// table driven test case
	tests := []struct {
		name string
		want map[string]string
	}{
		{
			name: "Test Adapter",
			want: map[string]string{
				"421f4b43-4d80-40fa-af8c-3369ea47d032": "record1",
				"4d5b63c4-2127-4592-be6b-22a777a5be7f": "record2",
				"72e155ab-0ffd-4f88-98e5-30dbc9dfac34": "record3",
			},
		},
	}

	// loop over the tests.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := apiAdapter.ConvertRecords(); err != nil {
				t.Errorf("TestAdapter() error = %v", err)
			}

			if len(apiAdapter.ListEntries()) != len(tt.want) {
				t.Errorf("TestAdapter() = %v, want %v", apiAdapter.ListEntries(), tt.want)
			}
		})
	}

	// print the entries from the modern api.
	for k, v := range apiAdapter.ListEntries() {
		t.Logf("key: %s, value: %s", k, v)
	}

}
