package utils

import (
	"bytes"
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"os"
	"strings"
)

// Table1Row represents a row in the first table
type Table1Row struct {
	ShipFrom   string
	CustomerPO string
	UBCPI      string
	ShipVia    string
	Term       string
}

// Table2Row represents a row in the second table
type Table2Row struct {
	CountryOfOrigin string
	VesselName      string
	BillOfLading    string
	ETDChina        string
}

// Table3Row represents a row in the third table
type Table3Row struct {
	PO          string
	StyleName   string
	StyleCode   string
	Description string
	Color       string
	Size        string
	Qty         string
	UPrice      string
	TotalUSD    string
}

var (
	smallNumbers = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	tens         = []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	thousands    = []string{"", "Thousand", "Million", "Billion"}
)

func buildTitle(pdf *gofpdf.Fpdf, address, invoice, invoiceOne, billTo, shipTo []string) {
	pdf.Ln(15)
	// Company information and invoice title
	pdf.SetFont("Arial", "B", 18)
	pdf.SetX(145)
	pdf.CellFormat(80, 10, "INVOICE", "0", 1, "L", false, 0, "")
	pdf.SetFont("Arial", "B", 10)
	pdf.SetX(20)
	pdf.CellFormat(80, 10, "UNITED BUSINESS CORPORATION", "0", 0, "L", false, 0, "")
	pdf.Ln(8)
	pdf.SetFont("Arial", "", 8)
	index := len(address)
	if len(address) < len(invoice) {
		index = len(invoice)
	}
	for i := 0; i < index; i++ {
		pdf.SetX(20)
		if i >= 0 && i < len(address) {
			pdf.CellFormat(125, 5, address[i], "0", 0, "L", false, 0, "")
		} else {
			pdf.CellFormat(125, 5, "", "0", 0, "L", false, 0, "")

		}
		if i >= 0 && i < len(invoice) {
			pdf.CellFormat(10, 5, invoice[i], "0", 0, "L", false, 0, "")
			pdf.CellFormat(45, 5, invoiceOne[i], "0", 0, "R", false, 0, "")
		} else {
			pdf.CellFormat(10, 5, "", "0", 0, "L", false, 0, "")
		}
		pdf.Ln(5)
	}
	pdf.Ln(5)
	// bill to & ship to title
	pdf.SetX(20)
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(125, 10, "BILL TO:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(10, 10, "SHIP TO:", "0", 0, "L", false, 0, "")
	pdf.Ln(8)
	pdf.SetFont("Arial", "", 8)
	index = len(billTo)
	if len(billTo) < len(shipTo) {
		index = len(shipTo)
	}
	for i := 0; i < index; i++ {
		pdf.SetX(20)
		if i >= 0 && i < len(billTo) {
			pdf.CellFormat(125, 5, billTo[i], "0", 0, "L", false, 0, "")
		} else {
			pdf.CellFormat(125, 5, "", "0", 0, "L", false, 0, "")
		}
		if i >= 0 && i < len(shipTo) {
			pdf.CellFormat(10, 5, shipTo[i], "0", 0, "L", false, 0, "")
		} else {
			pdf.CellFormat(10, 5, "", "0", 0, "L", false, 0, "")
		}
		pdf.Ln(5)
	}

}

// buildTable creates tables in the PDF with the given data
func buildTable(pdf *gofpdf.Fpdf, table1Data []Table1Row, table2Data []Table2Row, table3Data []Table3Row, lastStr, totalStr, subStr string) {
	// Table 1
	pdf.SetX(20)
	pdf.SetFont("Arial", "B", 5)
	pdf.CellFormat(26, 4, "Ship From", "1", 0, "C", false, 0, "")
	pdf.CellFormat(26, 4, "Customer PO", "1", 0, "C", false, 0, "")
	pdf.CellFormat(26, 4, "UBC PI", "1", 0, "C", false, 0, "")
	pdf.CellFormat(47, 4, "", "1", 0, "C", false, 0, "")
	pdf.CellFormat(25, 4, "Ship Via", "1", 0, "C", false, 0, "")
	pdf.CellFormat(30, 4, "Term", "1", 1, "C", false, 0, "")

	pdf.SetFont("Arial", "", 5)
	for _, row := range table1Data {
		pdf.SetX(20)
		pdf.CellFormat(26, 4, row.ShipFrom, "1", 0, "C", false, 0, "")
		pdf.CellFormat(26, 4, row.CustomerPO, "1", 0, "C", false, 0, "")
		pdf.CellFormat(26, 4, row.UBCPI, "1", 0, "C", false, 0, "")
		pdf.CellFormat(47, 4, "", "1", 0, "C", false, 0, "")
		pdf.CellFormat(25, 4, row.ShipVia, "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 4, row.Term, "1", 1, "C", false, 0, "")
	}

	// Add a row without inner borders between Table 1 and Table 2
	pdf.SetX(20)
	pdf.CellFormat(26, 4, "", "LR", 0, "C", false, 0, "")
	pdf.CellFormat(26, 4, "", "LR", 0, "C", false, 0, "")
	pdf.CellFormat(26, 4, "", "LR", 0, "C", false, 0, "")
	pdf.CellFormat(47, 4, "", "LR", 0, "C", false, 0, "")
	pdf.CellFormat(25, 4, "", "LR", 0, "C", false, 0, "")
	pdf.CellFormat(30, 4, "", "LR", 1, "C", false, 0, "")
	// Table 2
	pdf.SetX(20)
	pdf.SetFont("Arial", "B", 5)
	pdf.CellFormat(26, 4, "Country of Origin", "1", 0, "C", false, 0, "")
	pdf.CellFormat(26, 4, "Vessel/Flight Name", "1", 0, "C", false, 0, "")
	pdf.CellFormat(26, 4, "Bill of Lading No.", "1", 0, "C", false, 0, "")
	pdf.CellFormat(47, 4, "", "1", 0, "C", false, 0, "")
	pdf.CellFormat(25, 4, "", "1", 0, "C", false, 0, "")
	pdf.CellFormat(30, 4, "ETD CHINA", "1", 1, "C", false, 0, "")

	pdf.SetFont("Arial", "", 5)
	for _, row := range table2Data {
		pdf.SetX(20)
		pdf.CellFormat(26, 4, row.CountryOfOrigin, "1", 0, "C", false, 0, "")
		pdf.CellFormat(26, 4, row.VesselName, "1", 0, "C", false, 0, "")
		pdf.CellFormat(26, 4, row.BillOfLading, "1", 0, "C", false, 0, "")
		pdf.CellFormat(47, 4, "", "1", 0, "C", false, 0, "")
		pdf.CellFormat(25, 4, "", "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 4, row.ETDChina, "1", 1, "C", false, 0, "")
	}
	pdf.Ln(3)

	// Table 3
	pdf.SetX(20)
	pdf.SetFont("Arial", "B", 5)
	pdf.CellFormat(30, 4, "PO#", "1", 0, "L", false, 0, "")
	pdf.CellFormat(12, 4, "STYLE#", "1", 0, "L", false, 0, "")
	pdf.CellFormat(12, 4, "STYLE NAME", "1", 0, "L", false, 0, "")
	pdf.CellFormat(66, 4, "FABRICATION", "1", 0, "L", false, 0, "")
	pdf.CellFormat(15, 4, "COLOR", "1", 0, "C", false, 0, "")
	pdf.CellFormat(10, 4, "QTY(PC)", "1", 0, "C", false, 0, "")
	pdf.CellFormat(20, 4, "RATE", "1", 0, "C", false, 0, "")
	pdf.CellFormat(20, 4, "TOTAL", "1", 1, "C", false, 0, "")

	pdf.SetFont("Arial", "", 5)
	for _, row := range table3Data {
		pdf.SetX(20)
		pdf.CellFormat(30, 4, row.PO, "LR", 0, "L", false, 0, "")
		pdf.CellFormat(12, 4, row.StyleCode, "LR", 0, "L", false, 0, "")
		pdf.CellFormat(12, 4, row.StyleName, "LR", 0, "L", false, 0, "")
		pdf.CellFormat(66, 4, row.Description, "LR", 0, "L", false, 0, "")
		pdf.CellFormat(15, 4, row.Color, "LR", 0, "C", false, 0, "")
		pdf.CellFormat(10, 4, row.Qty, "LR", 0, "C", false, 0, "")

		// UPrice with $ left-aligned and number right-aligned
		pdf.CellFormat(5, 4, "$", "L", 0, "L", false, 0, "")
		pdf.CellFormat(15, 4, row.UPrice, "R", 0, "R", false, 0, "")

		// TotalUSD with $ left-aligned and number right-aligned
		pdf.CellFormat(5, 4, "$", "L", 0, "L", false, 0, "")
		pdf.CellFormat(15, 4, row.TotalUSD, "R", 1, "R", false, 0, "")
	}

	// Add the total row within the same table, aligned with the columns
	pdf.SetX(20)
	pdf.SetFont("Arial", "B", 5)
	// 添加最左边的表格框
	pdf.CellFormat(20, 4, "", "LT", 0, "C", false, 0, "")
	pdf.CellFormat(90, 4, "", "T", 0, "C", false, 0, "") // Empty cells to align with previous columns
	pdf.CellFormat(15, 4, "TOTAL", "1", 0, "C", false, 0, "")
	pdf.CellFormat(15, 4, totalStr, "1", 0, "C", false, 0, "")
	pdf.CellFormat(20, 4, "$", "1", 0, "L", false, 0, "")
	pdf.CellFormat(20, 4, subStr, "1", 1, "R", false, 0, "")

	// Add a new table for the total cartons and total amount text with only the outer border
	pdf.SetX(20)
	pdf.SetFont("Arial", "B", 5)
	pdf.MultiCell(180, 4, lastStr, "1", "L", false)
}

func convertToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	words := ""

	if num < 0 {
		words = "Minus "
		num = -num
	}

	for i := 0; num > 0; i++ {
		if num%1000 != 0 {
			words = convertHundreds(num%1000) + thousands[i] + " " + words
		}
		num /= 1000
	}

	return strings.TrimSpace(words)
}

func convertHundreds(num int) string {
	words := ""

	if num > 99 {
		words += smallNumbers[num/100] + " Hundred "
		num %= 100
	}

	if num > 19 {
		words += tens[num/10] + " "
		num %= 10
	}

	if num > 0 {
		words += smallNumbers[num] + " "
	}

	return words
}

func convertFloatToWords(amount float64) string {
	intPart := int(amount)
	decimalPart := int((amount - float64(intPart)) * 100)

	dollarWords := convertToWords(intPart)
	centWords := convertToWords(decimalPart)

	result := fmt.Sprintf("TOTAL USD %s DOLLARS", strings.ToUpper(dollarWords))
	if decimalPart > 0 {
		result += fmt.Sprintf(" AND %s CENTS", strings.ToUpper(centWords))
	}
	result += " ONLY"

	return result
}

func BuildInvoicePdf(table1Data []Table1Row, table2Data []Table2Row, table3Data []Table3Row,
	address, invoice, invoiceOne, billTo, shipTo []string, lastStr, totalStr, subStr string) (*bytes.Buffer, error) {

	pdf := gofpdf.NewCustom(&gofpdf.InitType{
		UnitStr: "mm",
		Size:    gofpdf.SizeType{Wd: 215, Ht: 280}, // Set page size to 21.5 cm x 18 cm
	})
	pdf.AddPage()

	buildTitle(pdf, address, invoice, invoiceOne, billTo, shipTo)
	pdf.Ln(10)
	// Build the tables
	for i := 0; i < 40-len(table3Data); i++ {
		t := Table3Row{UPrice: "", TotalUSD: "-"}
		table3Data = append(table3Data, t)
	}
	buildTable(pdf, table1Data, table2Data, table3Data, lastStr, totalStr, subStr)

	// 将PDF内容写入字节缓冲区
	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}

	// 将字节缓冲区的内容写入文件
	file, err := os.Create("invoice.pdf")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	_, err = file.Write(buf.Bytes())
	if err != nil {
		return nil, err
	}

	return &buf, nil

}

func ConvertFloatToWords(amount float64) string {
	intPart := int(amount)
	decimalPart := int((amount - float64(intPart)) * 100)

	dollarWords := convertToWords(intPart)
	centWords := convertToWords(decimalPart)

	result := fmt.Sprintf("TOTAL USD %s DOLLARS", strings.ToUpper(dollarWords))
	if decimalPart > 0 {
		result += fmt.Sprintf(" AND %s CENTS", strings.ToUpper(centWords))
	}
	result += " ONLY"

	return result
}
