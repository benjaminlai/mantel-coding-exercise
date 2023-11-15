package util

import "testing"

func Test_ShouldReporTopItemsInList(t *testing.T) {
	mockItemsList := []string{
		"test1",
		"test2",
		"test3",
		"test2",
		"test2",
		"test3",
	}
	expectedTopItem := "test2"
	expectedSecondTopItem := "test3"
	expectedThirdTopItem := "test1"

	result := FilterByTop(mockItemsList, 3)
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
