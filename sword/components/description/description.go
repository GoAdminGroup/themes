package description

import (
	"bytes"
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	adminTemplate "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/themes/sword/components"
	"html/template"
)

type Description struct {
	components.Base
	Name    string
	Border  string
	Number  template.HTML
	Title   template.HTML
	Arrow   string
	Color   template.HTML
	Percent template.HTML
}

func New() Description {
	return Description{}
}

func (c Description) SetNumber(value template.HTML) Description {
	c.Number = value
	return c
}

func (c Description) SetTitle(value template.HTML) Description {
	c.Title = value
	return c
}

func (c Description) SetArrow(value string) Description {
	c.Arrow = value
	return c
}

func (c Description) SetPercent(value template.HTML) Description {
	c.Percent = value
	return c
}

func (c Description) SetColor(value template.HTML) Description {
	c.Color = value
	return c
}

func (c Description) SetBorder(value string) Description {
	c.Border = value
	return c
}

func (c Description) GetTemplate() (*template.Template, string) {
	tmpl, err := template.New("description").
		Funcs(adminTemplate.DefaultFuncMap).
		Parse(List["description"])

	if err != nil {
		logger.Error("Description GetTemplate Error: ", err)
	}

	return tmpl, "description"
}

func (c Description) GetContent() template.HTML {
	buffer := new(bytes.Buffer)
	tmpl, defineName := c.GetTemplate()
	err := tmpl.ExecuteTemplate(buffer, defineName, c)
	if err != nil {
		fmt.Println("ComposeHtml Error:", err)
	}
	return template.HTML(buffer.String())
}
