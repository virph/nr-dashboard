package main

type Data struct {
	PageNumber  int
	PageTitle   string
	MetricTitle string
	StatusField string
	Attribute   DataAttribute
	Filter      DataFilter
	ErrorFilter DataErrorFilter
}

type DataAttribute struct {
	HitCount string
	Latency  string
}

type DataFilter struct {
	RawString       string
	FilterAttribute map[string]DataFilterAttribute
	FilterString    string
}

type DataFilterAttribute struct {
	Key      string
	IsNegate bool
	Values   []string
}

type DataErrorFilter struct {
	RawString    string
	FilterString string
}
