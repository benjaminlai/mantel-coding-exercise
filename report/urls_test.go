package report

import "testing"

func Test_ShouldReportVistedURLS(t *testing.T) {
	mockUniqueIPAddresses := []string{
		"/test2",
		"/test3",
		"test1",
		"/test2",
		"/test3",
		"/test3",
	}

	expectedTopItem := "/test3"
	expectedSecondTopItem := "/test2"
	expectedThirdTopItem := "test1"

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
