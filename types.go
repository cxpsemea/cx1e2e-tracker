package main

type ReportSettings struct {
	E2ESuffix string `json:"E2ESuffix"`
}

type ReportTestDetails struct {
	Name     string
	Source   string
	Test     string
	Duration float64
	Result   string
	ID       uint64
}

type Report struct {
	Settings ReportSettings      `json:"Settings"`
	Details  []ReportTestDetails `json:"Details"`
}

type HistoryMetric struct {
	Min, Max, Average float64
	Count             uint
}

type History struct {
	Pass map[string]HistoryMetric
}

type comparisonEntry struct {
	TestName             string
	Result1, Result2     string
	Duration1, Duration2 float64
	ID                   uint64
}
