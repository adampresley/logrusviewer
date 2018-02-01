package ui

import (
	"github.com/adampresley/logrusviewer/pkg/logfiles"
)

/*
ViewState is used to pass information along to pages
during render
*/
type ViewState struct {
	IsError    bool
	LogEntries logfiles.ParsedLogFile
	LogFile    string
	Message    string
	Title      string
}

/*
NewViewState returns a view state struct intialized with a title
*/
func NewViewState(title string) *ViewState {
	return &ViewState{
		IsError:    false,
		LogEntries: logfiles.NewParsedLogFile(),
		LogFile:    "",
		Message:    "",
		Title:      title,
	}
}
