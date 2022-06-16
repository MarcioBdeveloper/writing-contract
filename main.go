package main

import (
	"github.com/jung-kurt/gofpdf"
)

func main() {
	err := GeneratePdf("contract.pdf")
	if err != nil {
		panic(err)
	}
}

func GeneratePdf(filename string) error {

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetHeaderFuncMode(func() {
		//pdf.Image("avatar.png", 10, 6, 30, 0, false, "", 0, "")
		pdf.SetY(5)
		pdf.SetFont("Arial", "B", 15)
		pdf.Cell(80, 0, "")
		pdf.CellFormat(30, 10, "ASDSDASDASDASDASDASDAS", "1", 0, "C", false, 0, "")
		pdf.Ln(20)
	}, true)

	pdf.SetFooterFunc(func() {
		pdf.SetY(-15)
		pdf.SetFont("Arial", "I", 8)
		pdf.CellFormat(0, 10, "Zap veiculos multimarcas",
			"", 0, "C", false, 0, "")
	})

	return pdf.OutputFileAndClose(filename)
}

//pdf.SetFont("Arial", "B", 20)

// CellFormat(width, height, text, border, position after, align, fill, link, linkStr)
//pdf.CellFormat(200, 7, "Contrato de Venda", "0", 0, "CM", false, 0, "")

// ImageOptions(src, x, y, width, height, flow, options, link, linkStr)
// pdf.ImageOptions(
// 	"ZAP_VEICULOS.png",
// 	70, 20,
// 	25, 15,
// 	false,
// 	gofpdf.ImageOptions{ImageType: "png", ReadDpi: true},
// 	0,
// 	"",
// )
