package report

// PDFReportGenerator implements the ReportGenerator interface for generating PDF reports.
type PDFReportGenerator struct{}

func (prg PDFReportGenerator) GenerateReport() string {
	return "Generating a PDF report..."
}

// HTMLReportGenerator implements the ReportGenerator interface for generating HTML reports.
type HTMLReportGenerator struct{}

func (hrg HTMLReportGenerator) GenerateReport() string {
	return "Generating an HTML report..."
}
