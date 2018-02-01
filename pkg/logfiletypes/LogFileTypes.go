package logfiletypes

// LogFileType describes a type of log file
type LogFileType int

const (
	// LogrusJSON is a Logrus JSON log file
	LogrusJSON LogFileType = iota
)
