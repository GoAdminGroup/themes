package chart_legend

import (
	"bytes"
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/logger"
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
		Funcs(template.FuncMap{
			"lang":     language.Get,
			"langHtml": language.GetFromHtml,
			"link": func(cdnUrl, prefixUrl, assetsUrl string) string {
				if cdnUrl == "" {
					return prefixUrl + assetsUrl
				}
				return cdnUrl + assetsUrl
			},
			"isLinkUrl": func(s string) bool {
				return (len(s) > 7 && s[:7] == "http://") || (len(s) > 8 && s[:8] == "https://")
			},
		}).
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
