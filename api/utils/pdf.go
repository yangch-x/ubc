package utils

import (
	"io"
	"rsc.io/pdf"
	"strings"
)

func ExtractPDFText(file io.ReaderAt, size int64) (string, error) {
	// Open the PDF file
	pdfReader, err := pdf.NewReader(file, size)
	if err != nil {
		return "", err
	}

	var text strings.Builder
	numPages := pdfReader.NumPage()
	for i := 1; i <= numPages; i++ {
		page := pdfReader.Page(i)
		if page.V.IsNull() {
			continue
		}

		content := page.Content()

		for _, textBlock := range content.Text {
			text.WriteString(textBlock.S)
		}
	}

	return text.String(), nil
}
