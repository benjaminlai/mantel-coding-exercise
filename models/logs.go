package models

type LogEntry struct {
	IPAddress string
	URLPath   string
}

type LogEntries struct {
	IPAddresses []string
	URLPaths    []string
}
