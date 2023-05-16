package main

import (
	"fmt"

	"github.com/go-pdf/fpdf"
)

func main() {
	pdf := fpdf.New("P", "mm", "A4", "")

	pdf.SetTopMargin(30)
	pdf.SetHeaderFuncMode(func() {
		pdf.Image("avatar.png", 10, 5, 18, 0, false, "", 0, "")
		pdf.SetY(5)
		pdf.SetFont("Arial", "B", 25)
		pdf.Cell(25, 0, "")
		pdf.CellFormat(160, 20, "CONTRATO DE VENDA", "1", 1, "C", false, 0, "")
	}, true)

	pdf.SetFooterFunc(func() {
		pdf.SetY(-15)
		pdf.SetFont("Arial", "I", 8)
		pdf.CellFormat(0, 10, fmt.Sprintf("Page %d/{nb}", pdf.PageNo()), "", 0, "C", false, 0, "")
	})

	pdf.AliasNbPages("")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(70, 10, "")
	pdf.CellFormat(50, 10, "EMPRESA", "2", 1, "C", false, 0, "")
	pdf.Ln(10)

	pdf.SetFont("Arial", "I", 12)
	pdf.Cell(1, 10, "")
	pdf.CellFormat(0, 10, "Razao Social:Minha Loja CNPJ:98.654.987/0001-01", "", 1, "L", false, 0, "")
	pdf.Cell(1, 10, "")
	pdf.CellFormat(0, 10, "ENDEREÇO: AV Meio da Rua BAIRRO: Mangabeira CEP:58059434", "", 1, "L", false, 0, "")
	pdf.Cell(1, 10, "")
	pdf.CellFormat(0, 10, "CIDADE: Joao Pessoa- PB TELEFONE:(83) 89988998 VENDEDOR: Fulano", "", 1, "L", false, 0, "")

	err := pdf.OutputFileAndClose("contract.pdf")
	if err != nil {
		fmt.Print("Output file error:", err)
	}
}
