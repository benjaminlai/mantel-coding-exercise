package parser

import (
	"regexp"
)

type LogEntry struct {
	IPAddress string
	URLPath   string
}

func ParseLogLineContent(line string) LogEntry {
	logFileLineRegex := regexp.MustCompile(LOG_FILE_LINE_REGEX)
	matches := logFileLineRegex.FindStringSubmatch(line)

	return LogEntry{
		IPAddress: matches[1],
		URLPath:   matches[3],
	}
}
