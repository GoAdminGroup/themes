package card

import (
	"html/template"

	adminTemplate "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
)

type Card struct {
	*adminTemplate.BaseComponent

	Title    string
	SubTitle string
	Content  template.HTML
	Action   template.HTML
	Footer   template.HTML
}

func New() Card {
	return Card{
		BaseComponent: &adminTemplate.BaseComponent{
			Name:     "card",
			HTMLData: List["card"],
		},
	}
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

func (c Card) AddButton(button types.Button) Card {
	c.Footer, c.JS = button.Content()
	c.Callbacks = append(c.Callbacks, button.GetAction().GetCallbacks())
	return c
}

func (c Card) SetFooter(footer template.HTML) Card {
	c.Footer = footer
	return c
}

func (c Card) GetContent() template.HTML { return c.GetContentWithData(c) }
