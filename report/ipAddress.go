package report

import "github.com/benjaminlai/mantel-coding-exercise/util"

func ReportNumberOfUniqueIPAddresses(ipAddresses []string) int {
	uniqueIPAddresses := make(map[string]bool)

	for _, ipAddress := range ipAddresses {
		uniqueIPAddresses[ipAddress] = true
	}

	return len(uniqueIPAddresses)
}

func ReportMostActiveIPAddresses(ipAddresses []string, count int) []string {
	return util.FilterByTop(ipAddresses, count)
}
