package report

import "testing"

func Test_ShouldReportUniqueIPAddressCount(t *testing.T) {
	mockUniqueIPAddresses := []string{
		"192.0.0.1",
		"192.0.0.2",
		"192.0.0.3",
		"192.0.0.1",
		"192.0.0.2",
	}
	expectedUniqueIPAddressCount := 3

	result := ReportNumberOfUniqueIPAddresses(mockUniqueIPAddresses)
	if result != expectedUniqueIPAddressCount {
		t.Errorf("Expected Count: %d, Got: %d", expectedUniqueIPAddressCount, result)
	}
}
