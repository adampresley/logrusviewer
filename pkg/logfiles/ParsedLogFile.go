package logfiles

import (
	"encoding/json"
)

/*
A ParsedLogEntry is a key/value pair. The value
can be anything, and the key will always be a string
*/
type ParsedLogEntry map[string]interface{}

/*
A ParsedLogFile is essentially a slice of ParsedLogEntries
*/
type ParsedLogFile []ParsedLogEntry

/*
NewParsedLogEntry creates a new, blank ParsedLogEntry
*/
func NewParsedLogEntry() ParsedLogEntry {
	return make(ParsedLogEntry)
}

/*
NewParsedLogEntryFromJSON creates a new ParsedLogEntry from JSON data
*/
func NewParsedLogEntryFromJSON(data []byte) (ParsedLogEntry, error) {
	var err error
	result := NewParsedLogEntry()

	err = json.Unmarshal(data, &result)
	return result, err
}

/*
NewParsedLogFile creates a new slice of ParsedLogEntries
*/
func NewParsedLogFile() ParsedLogFile {
	return make(ParsedLogFile, 0, 500)
}
