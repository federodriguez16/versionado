package handler

import (
	"os"

	"github.com/federodriguez16/versionado/internal/models"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func Mostrar(c *models.Client, v *models.Versions) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle(*c.Name)
	t.AppendHeader(table.Row{"Producto", "Version"})
	t.AppendRow(table.Row{"Producto A", v.ProductA})
	t.AppendRow(table.Row{"Producto B", v.ProductB})
	t.AppendRow(table.Row{"Producto C", v.ProductC})
	t.AppendRow(table.Row{"Producto D", v.ProductD})
	t.AppendRow(table.Row{"Producto E", v.ProductE})
	t.SetStyle(table.StyleLight)
	t.Style().Title.Align = text.AlignCenter
	t.SetAutoIndex(false)
	t.Render()
}
