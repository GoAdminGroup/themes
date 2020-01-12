package card

import (
	"bytes"
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/themes/sword/components"
	"html/template"
	"strings"
)

type Card struct {
	components.Base
	Title    string
	SubTitle string
	Content  template.HTML
	Action   template.HTML
	Footer   template.HTML
}

func New() Card {
	return Card{}
}

func (c Card) SetTitle(title string) Card {
	c.Title = title
	return c
}

func (c Card) SetSubTitle(subTitle string) Card {
	c.SubTitle = subTitle
	return c
}

func (c Card) SetContent(content template.HTML) Card {
	c.Content = content
	return c
}

func (c Card) SetAction(action template.HTML) Card {
	c.Action = action
	return c
}

func (c Card) SetFooter(footer template.HTML) Card {
	c.Footer = footer
	return c
}

func (c Card) GetTemplate() (*template.Template, string) {
	tmpl, err := template.New("card").
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
			"render": func(s, old, repl template.HTML) template.HTML {
				return template.HTML(strings.Replace(string(s), string(old), string(repl), -1))
			},
		}).
		Parse(List["card"])

	if err != nil {
		logger.Error("Login GetTemplate Error: ", err)
	}

	return tmpl, "card"
}

func (c Card) GetContent() template.HTML {
	buffer := new(bytes.Buffer)
	tmpl, defineName := c.GetTemplate()
	err := tmpl.ExecuteTemplate(buffer, defineName, c)
	if err != nil {
		fmt.Println("ComposeHtml Error:", err)
	}
	return template.HTML(buffer.String())
}
