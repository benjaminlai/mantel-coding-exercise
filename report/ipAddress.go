package report

func ReportNumberOfUniqueIPAddresses(IPAddresses []string) int {
	uniqueIPAddresses := make(map[string]bool)

	for _, IPAddress := range IPAddresses {
		uniqueIPAddresses[IPAddress] = true
	}

	return len(uniqueIPAddresses)
}
