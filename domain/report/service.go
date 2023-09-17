package report

// Interface Segregation Principle (ISP)
// Open/Closed Principle (OCP)
// Open for extension but closed for modification.

// ReportGenerator is open for extension
// because you can create new implementations
// (e.g., for generating XML or CSV reports)
// without modifying the existing interface.
type ReportGenerator interface {
	GenerateReport() string
}

// Dependency Inversion Principle (DIP)
// High-level and Low-level modules should depend on abstractions
// rather than concrete details. Abstractions should not depend on details.

// ReportService class is also open for extension
// because it depends on the ReportGenerator interface.
// It's also mean that the high-level ReportService
// is not tightly coupled to specific report generator
// implementations (PDFReportGenerator or HTMLReportGenerator).
type ReportService struct {
	Generator ReportGenerator
}

func (rs ReportService) GenerateAndSendReport() string {
	return rs.Generator.GenerateReport()
}

// Bad practice for Dependency Inversion Principle (DIP)

type ViolateDIPReportService struct {
	PDFGenerator  PDFReportGenerator  // Violating DIP by depending on a concrete implementation
	HTMLGenerator HTMLReportGenerator // Violating DIP by depending on a concrete implementation
}

func (rs ViolateDIPReportService) ViolateDIPGenerateAndSendReport(format string) string {
	switch format {
	case "PDF":
		return rs.PDFGenerator.GenerateReport()
	case "HTML":
		return rs.HTMLGenerator.GenerateReport()
	default:
		return "Unsupported format"
	}
}
