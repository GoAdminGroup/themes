package chart_legend

import (
	"bytes"
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	adminTemplate "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/themes/adminlte/components"
	"html/template"
)

type ChartLegend struct {
	components.Base
	Data []map[string]string
}

func New() ChartLegend {
	return ChartLegend{}
}

func (c ChartLegend) SetData(value []map[string]string) ChartLegend {
	c.Data = value
	return c
}

func (c ChartLegend) GetTemplate() (*template.Template, string) {
	tmpl, err := template.New("chart-legend").
		Funcs(adminTemplate.DefaultFuncMap).
		Parse(List["chart-legend"])

	if err != nil {
		logger.Error("ChartLegend GetTemplate Error: ", err)
	}

	return tmpl, "chart-legend"
}

func (c ChartLegend) GetContent() template.HTML {
	buffer := new(bytes.Buffer)
	tmpl, defineName := c.GetTemplate()
	err := tmpl.ExecuteTemplate(buffer, defineName, c)
	if err != nil {
		fmt.Println("ComposeHtml Error:", err)
	}
	return template.HTML(buffer.String())
}
