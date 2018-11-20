package table

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"sort"
)

const (
	tableHeaderName = "Name"
	tableHeaderSignature = "Unterschrift"

	columnWidthName = 95.0
	columnWidthSignature = 95.0
	columnHeight = 9.0

	defaultFont = "Times"
	defaultFontSize = 16.0
	defaultPaperSize = "A4"
	defaultUnit = "mm"
	defaultOrientation = "P"
	defaultAlignment = "C"
	defaultBorderWidth = "1"

	pageTitle = "Gruppe %s - Übung %d"

	defaultNote = "Teilnahme ausschließlich mit Anmeldung"
	defaultNoteFontSize = 20.0
)

func nextLine(pdf *gofpdf.Fpdf) {
	pdf.Ln(-1)
}

func header(pdf *gofpdf.Fpdf) *gofpdf.Fpdf {
	pdf.SetFont(defaultFont, "B", defaultFontSize)
	pdf.CellFormat(columnWidthName, columnHeight, tableHeaderName, defaultBorderWidth, 0, defaultAlignment, false, 0, "")
	pdf.CellFormat(columnWidthSignature, columnHeight, tableHeaderSignature, defaultBorderWidth, 0, defaultAlignment, false, 0, "")
	nextLine(pdf)
	return pdf;
}

func table(pdf *gofpdf.Fpdf, students []string) *gofpdf.Fpdf {
	pdf.SetFont(defaultFont, "", defaultFontSize)
	translator := pdf.UnicodeTranslatorFromDescriptor("")

	// Sort student names in ascending order
	sort.Strings(students)

	// Print row for each student
	for _, student := range students {
		xBefore, yBefore := pdf.GetXY()
		pdf.MultiCell(columnWidthName, columnHeight, translator(student), defaultBorderWidth, defaultAlignment, false)
		yAfter := pdf.GetY()
		pdf.MoveTo(xBefore + columnWidthName, yBefore)
		pdf.MultiCell(columnWidthSignature, yAfter - yBefore, "", defaultBorderWidth, defaultAlignment, false)
	}

	return pdf
}

func note(pdf *gofpdf.Fpdf) *gofpdf.Fpdf {
	translator := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.SetFont(defaultFont, "BU", defaultNoteFontSize)
	pdf.CellFormat(pdf.GetLineWidth() * 1000, 16, translator(defaultNote), "", 0, defaultAlignment, false, 0, "")
	return pdf
}

func Generate(appointment int, group string, students []string) *gofpdf.Fpdf  {
	pdf := gofpdf.New(defaultOrientation, defaultUnit, defaultPaperSize, "")
	pdf.AddPage()
	pdf.SetFont(defaultFont, "B", defaultFontSize)

	// Generate page title
	translator := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.CellFormat(pdf.GetLineWidth() * 1000, 16, translator(fmt.Sprintf(pageTitle, group, appointment)), "", 0, defaultAlignment, false, 0, "")
	nextLine(pdf)

	// Generate table header
	pdf = header(pdf)

	// Generate table content
	pdf = table(pdf, students)

	nextLine(pdf)

	pdf = note(pdf)

	return pdf
}
