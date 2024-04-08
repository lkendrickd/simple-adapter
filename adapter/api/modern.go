package adapter

import "errors"

// ModernAPI - the new api interface to implement. This represents the modern api.
type ModernAPI interface {
	Entries() map[string]string
	AddEntry(key, value string) error
}

// EntriesAPI - this is the struct that implements the API interface. This represents the modern api.
// This wraps the legacy api and implements the modern api. This is the adapter and one of the key points of the adapter pattern.
type EntriesAPI struct {
	entries map[string]string
}

// Entries - this is the method that implements the ModernAPI interface. This represents the modern api call that returns a list of entries.
// this will call the legacy api and convert the results to the modern api entries instead of records.
func (e *EntriesAPI) Entries() map[string]string {
	return e.entries
}

// AddEntry - this is the method that implements the ModernAPI interface,
// This will add an entry into the entries map
func (e *EntriesAPI) AddEntry(key, value string) error {
	if key == "" || value == "" {
		return errors.New("key and value cannot be empty")
	}

	e.entries[key] = value

	return nil
}

// NewEntries - this is the constructor for the EntriesAPI struct. This is used to create a new instance of the EntriesAPI struct.
func NewEntriesAPI() ModernAPI {
	return &EntriesAPI{
		entries: make(map[string]string),
	}
}
