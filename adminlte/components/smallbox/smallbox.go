package smallbox

import (
	"bytes"
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/themes/adminlte/components"
	"html/template"
	"strings"
)

type SmallBox struct {
	components.Base
	Title      template.HTML
	Value      template.HTML
	Url        string
	Color      template.HTML
	IsSvg      bool
	IsHexColor bool
	Icon       template.HTML
}

func New() SmallBox {
	return SmallBox{}
}

func (s SmallBox) SetTitle(value template.HTML) SmallBox {
	s.Title = value
	return s
}

func (s SmallBox) SetValue(value template.HTML) SmallBox {
	s.Value = value
	return s
}

func (s SmallBox) SetColor(value template.HTML) SmallBox {
	s.Color = value
	if strings.Contains(string(value), "#") {
		s.IsHexColor = true
	}
	return s
}

func (s SmallBox) SetIcon(value template.HTML) SmallBox {
	s.Icon = value
	if strings.Contains(string(value), "svg") {
		s.IsSvg = true
	}
	return s
}

func (s SmallBox) SetUrl(value string) SmallBox {
	s.Url = value
	return s
}

func (s SmallBox) GetTemplate() (*template.Template, string) {
	tmpl, err := template.New("smallbox").
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
		Parse(List["smallbox"])

	if err != nil {
		logger.Error("SmallBox GetTemplate Error: ", err)
	}

	return tmpl, "smallbox"
}

func (s SmallBox) GetContent() template.HTML {
	buffer := new(bytes.Buffer)
	tmpl, defineName := s.GetTemplate()
	err := tmpl.ExecuteTemplate(buffer, defineName, s)
	if err != nil {
		fmt.Println("ComposeHtml Error:", err)
	}
	return template.HTML(buffer.String())
}
