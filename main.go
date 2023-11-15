package main

import (
	"fmt"

	"github.com/benjaminlai/mantel-coding-exercise/parser"
	"github.com/benjaminlai/mantel-coding-exercise/report"
)

func main() {
	logEntries := parser.ParseLogsFromFile("resources/sample.log")

	// Find number of unique IP Addresses
	uniqueIPAddressCount := report.ReportNumberOfUniqueIPAddresses(logEntries.IPAddresses)
	fmt.Println("Unique IP Address Count", uniqueIPAddressCount)

	// Find top 3 most visited URLS
	mostVisitedURLS := report.ReportMostVisitedURLS(logEntries.IPAddresses, 3)
	fmt.Println("Top 3 most visited URLS", mostVisitedURLS)

	// Find top 3 most active IP Addresses
	mostActiveIPAddresses := report.ReportMostActiveIPAddresses(logEntries.IPAddresses, 3)
	fmt.Println("Top 3 most active IP Addresses", mostActiveIPAddresses)
}
