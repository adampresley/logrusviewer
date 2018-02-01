package logfiles

import (
	"encoding/json"
	"time"
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
func NewParsedLogEntryFromJSON(data []byte, line int) (ParsedLogEntry, error) {
	var err error
	result := NewParsedLogEntry()

	err = json.Unmarshal(data, &result)
	if err == nil {
		result["lineNumber"] = line
	}

	return result, err
}

/*
NewParsedLogFile creates a new slice of ParsedLogEntries
*/
func NewParsedLogFile() ParsedLogFile {
	return make(ParsedLogFile, 0, 500)
}

/*
MakeTimePretty attempts to format a time key in local server time
*/
func (p ParsedLogEntry) MakeTimePretty(key string) string {
	var ok bool
	var err error
	var value interface{}
	var valueString string
	var dateTime time.Time

	if value, ok = p[key]; !ok {
		return ""
	}

	valueString = value.(string)

	if dateTime, err = time.Parse(time.RFC3339, valueString); err != nil {
		return valueString
	}

	return dateTime.In(time.Local).Format("2006-01-02 3:04PM")
}

/*
GetLine returns the ParsedLogEntry at a specified line
*/
func (l ParsedLogFile) GetLine(lineNumber int) ParsedLogEntry {
	return l[lineNumber]
}
