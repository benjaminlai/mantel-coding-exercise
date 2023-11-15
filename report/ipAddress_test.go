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

func Test_ShouldReportMostActiveIPAddresses(t *testing.T) {
	mockUniqueIPAddresses := []string{
		"192.0.0.1",
		"192.0.0.2",
		"192.0.0.3",
		"192.0.0.1",
		"192.0.0.2",
		"192.0.0.2",
	}

	expectedTopItem := "192.0.0.2"
	expectedSecondTopItem := "192.0.0.1"
	expectedThirdTopItem := "192.0.0.3"

	result := ReportMostActiveIPAddresses(mockUniqueIPAddresses, 3)
	if result[0] != expectedTopItem {
		t.Errorf("Expected Item: %s, Got: %s", result[0], expectedTopItem)
	}
	if result[1] != expectedSecondTopItem {
		t.Errorf("Expected Item: %s, Got: %s", result[1], expectedSecondTopItem)
	}
	if result[2] != expectedThirdTopItem {
		t.Errorf("Expected Item: %s, Got: %s", result[2], expectedThirdTopItem)
	}
}
