package card

import (
	"bytes"
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	adminTemplate "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/themes/sword/components"
	"html/template"
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
		Funcs(adminTemplate.DefaultFuncMap).
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
