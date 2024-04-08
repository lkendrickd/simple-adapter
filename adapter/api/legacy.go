package adapter

// LegacyAPI - the legacy api unterface to implement. This represents the old api
type LegacyAPI interface {
	Records() []string
}

// RecordsAPI - this is the struct that implements the LegacyAPI interface. This represents the old api.
type RecordsAPI struct {
	records []string
}

// Records - this is the method that implements the LegacyAPI interface. This represents the old api call that returns a list of records.
func (r *RecordsAPI) Records() []string {
	return r.records
}

// NewRecordsAPI - this is the constructor for the RecordsAPI struct. This is used to create a new instance of the RecordsAPI struct.
func NewRecordsAPI(s ...string) LegacyAPI {
	return &RecordsAPI{
		records: s,
	}
}
