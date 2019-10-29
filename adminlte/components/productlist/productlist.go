package productlist

import (
	"bytes"
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/themes/adminlte/components"
	"html/template"
)

type ProductList struct {
	components.Base
	Data []map[string]string
}

func New() ProductList {
	return ProductList{}
}

func (p ProductList) SetData(value []map[string]string) ProductList {
	p.Data = value
	return p
}

func (p ProductList) GetTemplate() (*template.Template, string) {
	tmpl, err := template.New("productlist").
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
		Parse(List["productlist"])

	if err != nil {
		logger.Error("ProductList GetTemplate Error: ", err)
	}

	return tmpl, "productlist"
}

func (p ProductList) GetContent() template.HTML {
	buffer := new(bytes.Buffer)
	tmpl, defineName := p.GetTemplate()
	err := tmpl.ExecuteTemplate(buffer, defineName, p)
	if err != nil {
		fmt.Println("ComposeHtml Error:", err)
	}
	return template.HTML(buffer.String())
}
