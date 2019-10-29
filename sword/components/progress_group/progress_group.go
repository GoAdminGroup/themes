package progress_group

import (
	"bytes"
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/themes/sword/components"
	"html/template"
	"strings"
)

type ProgressGroup struct {
	components.Base
	Title       template.HTML
	Molecular   int
	Denominator int
	Color       template.HTML
	IsHexColor  bool
	Percent     int
}

func New() ProgressGroup {
	return ProgressGroup{}
}

func (p ProgressGroup) SetTitle(value template.HTML) ProgressGroup {
	p.Title = value
	return p
}

func (p ProgressGroup) SetColor(value template.HTML) ProgressGroup {
	p.Color = value
	if strings.Contains(string(value), "#") {
		p.IsHexColor = true
	}
	return p
}

func (p ProgressGroup) SetPercent(value int) ProgressGroup {
	p.Percent = value
	return p
}

func (p ProgressGroup) SetDenominator(value int) ProgressGroup {
	p.Denominator = value
	return p
}

func (p ProgressGroup) SetMolecular(value int) ProgressGroup {
	p.Molecular = value
	return p
}

func (p ProgressGroup) GetTemplate() (*template.Template, string) {
	tmpl, err := template.New("progress-group").
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
		Parse(List["progress-group"])

	if err != nil {
		logger.Error("ProgressGroup GetTemplate Error: ", err)
	}

	return tmpl, "progress-group"
}

func (p ProgressGroup) GetContent() template.HTML {
	buffer := new(bytes.Buffer)
	tmpl, defineName := p.GetTemplate()
	err := tmpl.ExecuteTemplate(buffer, defineName, p)
	if err != nil {
		fmt.Println("ComposeHtml Error:", err)
	}
	return template.HTML(buffer.String())
}
