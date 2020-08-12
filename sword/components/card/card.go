package card

import (
	"html/template"

	"github.com/GoAdminGroup/go-admin/modules/utils"
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

	ID        string
	BodyID    string
	TopID     string
	ContentID string
	FooterID  string
}

func New() Card {
	UUID := utils.Uuid(10)
	return Card{
		BaseComponent: &adminTemplate.BaseComponent{
			Name:     "card",
			HTMLData: List["card"],
		},
		ID:        UUID,
		BodyID:    UUID + "_body",
		TopID:     UUID + "_top",
		ContentID: UUID + "_content",
		FooterID:  UUID + "_footer",
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

func (c Card) BindAction(action types.Action) Card {
	c.BindActionTo(action, "#"+c.ID)
	return c
}

func (c Card) BindActionToBody(action types.Action) Card {
	c.BindActionTo(action, "#"+c.BodyID)
	return c
}

func (c Card) BindActionToTop(action types.Action) Card {
	c.BindActionTo(action, "#"+c.TopID)
	return c
}

func (c Card) BindActionToContent(action types.Action) Card {
	c.BindActionTo(action, "#"+c.ContentID)
	return c
}

func (c Card) BindActionToFooter(action types.Action) Card {
	c.BindActionTo(action, "#"+c.FooterID)
	return c
}

func (c Card) GetContent() template.HTML { return c.GetContentWithData(c) }
