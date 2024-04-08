package main

import (
	"fmt"

	adapter "gitlab.com/simple-adapter/adapter/api"
)

func main() {
	// Create a new instance of the legacy API with some records.
	legacyAPI := adapter.NewRecordsAPI("record1", "record2", "record3")

	// Create a new instance of the modern API.
	modernAPI := adapter.NewEntriesAPI()

	// Create an instance of the adapter.
	apiAdapter := adapter.NewAdapter(legacyAPI, modernAPI)

	// Convert the legacy records to modern entries using the adapter.
	err := apiAdapter.ConvertRecords()
	if err != nil {
		fmt.Println("Error converting records:", err)
		return
	}

	// Fetch and print the entries from the modern API.
	entries := apiAdapter.ListEntries()
	fmt.Println("Converted Entries:")
	for key, value := range entries {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
}
