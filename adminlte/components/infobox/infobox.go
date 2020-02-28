package infobox

import (
	"bytes"
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	adminTemplate "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/themes/adminlte/components"
	"html/template"
	"strings"
)

type InfoBox struct {
	components.Base
	Icon       template.HTML
	Text       template.HTML
	Number     template.HTML
	Content    template.HTML
	Color      template.HTML
	IsHexColor bool
	IsSvg      bool
}

func New() InfoBox {
	return InfoBox{}
}

func (i InfoBox) SetIcon(value template.HTML) InfoBox {
	i.Icon = value
	if strings.Contains(string(value), "svg") {
		i.IsSvg = true
	}
	return i
}

func (i InfoBox) SetText(value template.HTML) InfoBox {
	i.Text = value
	return i
}

func (i InfoBox) SetNumber(value template.HTML) InfoBox {
	i.Number = value
	return i
}

func (i InfoBox) SetContent(value template.HTML) InfoBox {
	i.Content = value
	return i
}

func (i InfoBox) SetColor(value template.HTML) InfoBox {
	i.Color = value
	if strings.Contains(string(value), "#") {
		i.IsHexColor = true
	}
	return i
}

func (i InfoBox) GetTemplate() (*template.Template, string) {
	tmpl, err := template.New("infobox").
		Funcs(adminTemplate.DefaultFuncMap).
		Parse(List["infobox"])

	if err != nil {
		logger.Error("InfoBox GetTemplate Error: ", err)
	}

	return tmpl, "infobox"
}

func (i InfoBox) GetContent() template.HTML {
	buffer := new(bytes.Buffer)
	tmpl, defineName := i.GetTemplate()
	err := tmpl.ExecuteTemplate(buffer, defineName, i)
	if err != nil {
		fmt.Println("ComposeHtml Error:", err)
	}
	return template.HTML(buffer.String())
}
