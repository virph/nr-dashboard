package main

import (
	"bytes"
	"fmt"
	"text/template"
)

func (d *DashboardBuilder) ProcessData() {
	d.DashboardBuilderData = make([]DashboardBuilderDataPage, 0)

	dataByPageNumber := make(map[int][]Data, 0)
	for _, sd := range d.SourceData {
		dataByPageNumber[sd.PageNumber] = append(dataByPageNumber[sd.PageNumber], sd)
	}

	for i := 1; i < len(dataByPageNumber)+1; i++ {
		if len(dataByPageNumber[i]) == 0 {
			continue
		}

		d.DashboardBuilderData = append(d.DashboardBuilderData, DashboardBuilderDataPage{
			PageTitle:    dataByPageNumber[i][0].PageTitle,
			Metrics:      dataByPageNumber[i],
			TotalMetrics: len(dataByPageNumber[i]),
		})
	}
	d.TotalPage = len(d.DashboardBuilderData)

	var (
		rowsPerPage       = 5
		columnsPerSection = 12
		sectionsPerRow    = 2

		currPage    = 1
		currRow     = 1
		currSection = 1
	)

	sections := make([]BillboardTemplateSection, 0)
	rows := make([]BillboardTemplateRow, 0)
	d.BillboardData = make([]BillboardTemplatePage, 0)
	for _, pg := range d.DashboardBuilderData {
		bSection := BillboardTemplateSection{
			Title: pg.PageTitle,
			Items: make([]BillboardTemplateItem, 0),
		}

		bSection.Items = append(bSection.Items, BillboardTemplateItem{
			Column: 1 + ((currSection - 1) * columnsPerSection),
			Row:    currRow,
			Height: 1,
			Title:  pg.PageTitle,
			Type:   "section_title",
			Width:  2,
		})

		currColumn := 3
		for _, m := range pg.Metrics {
			bSection.Items = append(bSection.Items, BillboardTemplateItem{
				Column: currColumn + ((currSection - 1) * columnsPerSection),
				Row:    currRow,
				Height: 1,
				Title:  m.MetricTitle,
				Type:   "metric",
				Width:  1,
				Metric: m,
			})
			currColumn = currColumn + 1
		}
		if currColumn <= columnsPerSection {
			bSection.Items = append(bSection.Items, BillboardTemplateItem{
				Column: currColumn + ((currSection - 1) * columnsPerSection),
				Row:    currRow,
				Height: 1,
				Title:  " ",
				Type:   "space",
				Width:  columnsPerSection - currColumn + 1,
			})
		}

		sections = append(sections, bSection)
		// next
		currSection = currSection + 1
		if currSection > sectionsPerRow {
			rows = append(rows, BillboardTemplateRow{
				Sections: sections,
			})
			sections = make([]BillboardTemplateSection, 0)

			currSection = 1
			currRow = currRow + 1
		}
		if currRow > rowsPerPage {
			d.BillboardData = append(d.BillboardData, BillboardTemplatePage{
				Title: fmt.Sprintf("Summary - %d", currPage),
				Rows:  rows,
			})
			rows = make([]BillboardTemplateRow, 0)

			currRow = 1
			currPage = currPage + 1
		}
	}
	d.BillboardData = append(d.BillboardData, BillboardTemplatePage{
		Title: fmt.Sprintf("Summary - %d", currPage),
		Rows:  rows,
	})
	for i := 0; i < len(d.BillboardData); i++ {
		d.BillboardData[i].Rows[len(d.BillboardData[i].Rows)-1].Sections[len(d.BillboardData[i].Rows[len(d.BillboardData[i].Rows)-1].Sections)-1].Items[len(d.BillboardData[i].Rows[len(d.BillboardData[i].Rows)-1].Sections[len(d.BillboardData[i].Rows[len(d.BillboardData[i].Rows)-1].Sections)-1].Items)-1].IsLast = true
	}

}

func (d DashboardBuilder) Build() (*bytes.Buffer, error) {
	funcMap := template.FuncMap{
		"getRows": func(idx, start int) int {
			return start + (idx * 3)
		},
	}

	tmpl, err := template.
		New("dashboard-template").
		Delims("[[", "]]").
		Funcs(funcMap).
		Parse(d.StrTemplate)
	if err != nil {
		return nil, err
	}

	buff := bytes.NewBuffer(nil)
	return buff, tmpl.Execute(buff, d)
}
