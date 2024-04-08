package adapter

import (
	"errors"

	"github.com/google/uuid"
)

// APIAdapter - this is the interface that the adapter will implement.
type APIAdapter interface {
	ConvertRecords() error
	ListEntries() map[string]string
}

// Adapter - this is the struct that implements the APIAdapter interface.
// This is the adapter and one of the key points of the adapter pattern.
// This is the struct that will be used to convert the legacy api to the modern api.
type Adapter struct {
	legacy LegacyAPI
	modern ModernAPI
}

// ConvertRecords - this is the method that implements the APIAdapter interface.
func (a *Adapter) ConvertRecords() error {
	if len(a.legacy.Records()) <= 0 {
		return errors.New("no records to convert")
	}

	// iterate over the legacy records and add them to the modern api.
	for _, v := range a.legacy.Records() {
		a.modern.AddEntry(uuid.NewString(), v)
	}

	return nil
}

// ListEntries - returns the entries in the modern api
func (a *Adapter) ListEntries() map[string]string {
	return a.modern.Entries()
}

// NewAdapter - this is the constructor for the Adapter struct.
func NewAdapter(legacy LegacyAPI, modern ModernAPI) APIAdapter {
	return &Adapter{
		legacy: legacy,
		modern: modern,
	}
}
