package main

import (
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/template/components"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/themes/sword"
	"github.com/GoAdminGroup/themes/sword/resource"
	"html/template"
	"strings"
)

type Theme struct {
	Name string
	components.Base
}

var Sword = Theme{
	Name: "sword",
	Base: components.Base{
		Attribute: types.Attribute{
			TemplateList: sword.TemplateList,
		},
	},
}

func (Theme) GetTmplList() map[string]string {
	return sword.TemplateList
}

func (Theme) GetTemplate(isPjax bool) (tmpler *template.Template, name string) {
	var err error

	if !isPjax {
		name = "layout"
		tmpler, err = template.New("layout").Funcs(template.FuncMap{
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
			"renderJS": func(s template.JS, old, repl template.HTML) template.JS {
				return template.JS(strings.Replace(string(s), string(old), string(repl), -1))
			},
			"divide": func(a, b int) int {
				return a / b
			},
		}).Parse(sword.TemplateList["layout"] +
			sword.TemplateList["head"] + sword.TemplateList["header"] + sword.TemplateList["sidebar"] +
			sword.TemplateList["footer"] + sword.TemplateList["js"] + sword.TemplateList["menu"] +
			sword.TemplateList["admin_panel"] + sword.TemplateList["content"])
	} else {
		name = "content"
		tmpler, err = template.New("content").Funcs(template.FuncMap{
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
			"renderJS": func(s template.JS, old, repl template.HTML) template.JS {
				return template.JS(strings.Replace(string(s), string(old), string(repl), -1))
			},
			"divide": func(a, b int) int {
				return a / b
			},
		}).Parse(sword.TemplateList["admin_panel"] + sword.TemplateList["content"])
	}

	if err != nil {
		panic(err)
	}

	return
}

func (Theme) GetAsset(path string) ([]byte, error) {
	path = "resource" + path
	return resource.Asset(path)
}

func (Theme) GetAssetList() []string {
	return resource.AssetsList
}
