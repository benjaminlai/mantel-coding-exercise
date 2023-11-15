package report

import "github.com/benjaminlai/mantel-coding-exercise/util"

func ReportNumberOfUniqueIPAddresses(IPAddresses []string) int {
	uniqueIPAddresses := make(map[string]bool)

	for _, IPAddress := range IPAddresses {
		uniqueIPAddresses[IPAddress] = true
	}

	return len(uniqueIPAddresses)
}

func ReportMostActiveIPAddresses(IPAddresses []string, count int) []string {
	return util.FilterByTop(IPAddresses, count)
}
