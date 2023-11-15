package main

type DashboardBuilder struct {
	StrTemplate    string
	DashboardTitle string
	SourceData     []Data

	BillboardData        []BillboardTemplatePage
	DashboardBuilderData []DashboardBuilderDataPage
	TotalPage            int
}

type DashboardBuilderDataPage struct {
	PageTitle    string
	Metrics      []Data
	TotalMetrics int
}

type BillboardTemplatePage struct {
	Title string
	Rows  []BillboardTemplateRow
}

type BillboardTemplateRow struct {
	Sections []BillboardTemplateSection
}

type BillboardTemplateSection struct {
	Title string
	Items []BillboardTemplateItem
}

type BillboardTemplateItem struct {
	Column int
	Row    int
	Width  int
	Height int

	Title  string
	Type   string
	Metric Data
	IsLast bool
}
