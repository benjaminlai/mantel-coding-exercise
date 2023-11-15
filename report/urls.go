package report

import "github.com/benjaminlai/mantel-coding-exercise/util"

func ReportMostVisitedURLS(urls []string, count int) []string {
	return util.FilterByTop(urls, count)
}
