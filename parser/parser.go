package parser

import (
	"bufio"
	"log"
	"os"
	"regexp"

	"github.com/benjaminlai/mantel-coding-exercise/models"
)

func ParseLogsFromFile(fileName string) models.LogEntries {
	ipAddresses := make([]string, 0)
	urlPaths := make([]string, 0)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parsedLogEntry := parseLogLineContent(line)
		ipAddresses = append(ipAddresses, parsedLogEntry.IPAddress)
		urlPaths = append(urlPaths, parsedLogEntry.URLPath)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return models.LogEntries{
		IPAddresses: ipAddresses,
		URLPaths:    urlPaths,
	}
}

func parseLogLineContent(line string) models.LogEntry {
	logFileLineRegex := regexp.MustCompile(LOG_FILE_LINE_REGEX)
	matches := logFileLineRegex.FindStringSubmatch(line)

	return models.LogEntry{
		IPAddress: matches[1],
		URLPath:   matches[3],
	}
}
