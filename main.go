package main

//func main() {
//	pdf := gofpdf.New("P", "mm", "A4", "")
//
//	pdf.AddPage()
//	pdf.SetFont("Arial", "B", 16)
//
//	// Title
//	pdf.Cell(40, 10, "INVOICE")
//
//	// Business Information
//	pdf.SetFont("Arial", "", 12)
//	pdf.Ln(10)
//	pdf.Cell(40, 10, "UNITED BUSINESS CORPORATION")
//	pdf.Ln(6)
//	pdf.Cell(40, 10, "4981 Irwindale Ave., Suite 700")
//	pdf.Ln(6)
//	pdf.Cell(40, 10, "Irwindale, CA 91706")
//	pdf.Ln(6)
//	pdf.Cell(40, 10, "Tel: ( 626 )727-6268")
//	pdf.Ln(6)
//	pdf.Cell(40, 10, "Fax: ( 626 )727-6265")
//
//	// Invoice Details
//	pdf.SetXY(150, 20)
//	pdf.SetFont("Arial", "B", 12)
//	pdf.Cell(40, 10, "Invoice No.:")
//	pdf.Ln(6)
//	pdf.SetXY(150, 26)
//	pdf.Cell(40, 10, "Invoice Date:")
//
//	// Billing and Shipping Information
//	pdf.SetXY(10, 60)
//	pdf.SetFont("Arial", "", 12)
//	pdf.Cell(40, 10, "BILL TO:")
//	pdf.SetXY(100, 60)
//	pdf.Cell(40, 10, "SHIP TO:")
//
//	// Table Headers
//	pdf.SetXY(10, 80)
//	pdf.SetFont("Arial", "B", 12)
//	pdf.CellFormat(30, 10, "UBC PI", "1", 0, "C", false, 0, "")
//	pdf.CellFormat(40, 10, "Ship From", "1", 0, "C", false, 0, "")
//	pdf.CellFormat(40, 10, "Ship To", "1", 0, "C", false, 0, "")
//	pdf.CellFormat(40, 10, "ETD China", "1", 0, "C", false, 0, "")
//	pdf.CellFormat(40, 10, "Bill of Lading", "1", 1, "C", false, 0, "")
//
//	// Table Rows
//	pdf.SetFont("Arial", "", 12)
//	pdf.CellFormat(30, 10, "", "1", 0, "C", false, 0, "")
//	pdf.CellFormat(40, 10, "", "1", 0, "C", false, 0, "")
//	pdf.CellFormat(40, 10, "", "1", 0, "C", false, 0, "")
//	pdf.CellFormat(40, 10, "", "1", 0, "C", false, 0, "")
//	pdf.CellFormat(40, 10, "", "1", 1, "C", false, 0, "")
//
//	// Additional Table
//	pdf.SetXY(10, 100)
//	pdf.SetFont("Arial", "B", 12)
//	pdf.CellFormat(40, 10, "Country of Origin", "1", 0, "C", false, 0, "")
//	pdf.CellFormat(40, 10, "CUSTOMER PO", "1", 0, "C", false, 0, "")
//	pdf.CellFormat(40, 10, "Vessel/Air", "1", 0, "C", false, 0, "")
//	pdf.CellFormat(40, 10, "Term", "1", 0, "C", false, 0, "")
//	pdf.CellFormat(40, 10, "Due Date", "1", 1, "C", false, 0, "")
//
//	// Additional Table Rows
//	pdf.SetFont("Arial", "", 12)
//	pdf.CellFormat(40, 10, "China", "1", 0, "C", false, 0, "")
//	pdf.CellFormat(40, 10, "", "1", 0, "C", false, 0, "")
//	pdf.CellFormat(40, 10, "", "1", 0, "C", false, 0, "")
//	pdf.CellFormat(40, 10, "", "1", 0, "C", false, 0, "")
//	pdf.CellFormat(40, 10, "", "1", 1, "C", false, 0, "")
//
//	// Item Table Headers
//	pdf.SetXY(10, 120)
//	pdf.SetFont("Arial", "B", 12)
//	pdf.CellFormat(20, 10, "PO", "1", 0, "C", false, 0, "")
//	pdf.CellFormat(30, 10, "STYLE NO.", "1", 0, "C", false, 0, "")
//	pdf.CellFormat(40, 10, "STYLE NAME", "1", 0, "C", false, 0, "")
//	pdf.CellFormat(40, 10, "DESCRIPTION", "1", 0, "C", false, 0, "")
//	pdf.CellFormat(20, 10, "COLOR", "1", 0, "C", false, 0, "")
//	pdf.CellFormat(20, 10, "QTY(PC)", "1", 0, "C", false, 0, "")
//	pdf.CellFormat(20, 10, "U/PRICE", "1", 0, "C", false, 0, "")
//	pdf.CellFormat(20, 10, "TOTALUSD", "1", 1, "C", false, 0, "")
//
//	// Save PDF to file
//	err := pdf.OutputFileAndClose("invoice.pdf")
//	if err != nil {
//		panic(err)
//	}
//}

/*
BILL TO:                              SHIP TO:
INEFFABLE MUSIC                         VIA MERCH
829 27TH AVE.                            2164 N. GLASSELL ST.
OAKLAND CA 94601                         ORANGE CA 92864
ATTN: MARINA PETRO                       1. ATTN: STICK FIGURE
                                         2. ATTN: THE MOVEMENT
*/
