package common

import (
	"html/template"
	"strings"

	"github.com/GoAdminGroup/go-admin/modules/config"
	adminTemplate "github.com/GoAdminGroup/go-admin/template"
)

type BaseTheme struct {
	AssetPaths   map[string]string
	TemplateList map[string]string
}

const Version = "v0.0.39"

func (b *BaseTheme) GetVersion() string {
	return Version
}

func (b *BaseTheme) GetRequirements() []string {
	return []string{">=v1.2.16"}
}

var comps = []string{"datatable.min.js", "form.min.js", "tree.min.js", "treeview.min.js"}

func inArray(s string, arr []string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}

func (b *BaseTheme) GetAssetImportHTML(exclude ...string) template.HTML {
	res := template.HTML("")
	for name, html := range b.AssetPaths {
		if inArray(name, comps) {
			if !inArray(strings.Replace(name, ".min.js", "", -1), exclude) {
				res += GetImportJSTag("/assets" + html)
			}
		}
	}
	return res
}

func (b *BaseTheme) GetHeadHTML() template.HTML {
	res := GetImportJSTag("/assets" + b.AssetPaths["all.min.js"])
	res += GetImportCSSTag("/assets" + b.AssetPaths["all.min.css"])
	return res
}

func (b *BaseTheme) GetFootJS() template.HTML {
	return GetImportJSTag("/assets" + b.AssetPaths["all_2.min.js"])
}

func (b *BaseTheme) Get500HTML() template.HTML {
	return template.HTML(b.TemplateList["500"])
}

func (b *BaseTheme) Get404HTML() template.HTML {
	return template.HTML(b.TemplateList["404"])
}

func (b *BaseTheme) Get403HTML() template.HTML {
	return template.HTML(b.TemplateList["403"])
}

func (b *BaseTheme) GetTemplate(isPjax bool) (tmpl *template.Template, name string) {
	var err error

	if !isPjax {
		name = "layout"
		tmpl, err = template.New("layout").Funcs(adminTemplate.DefaultFuncMap).
			Parse(b.TemplateList["layout"] +
				b.TemplateList["head"] + b.TemplateList["header"] + b.TemplateList["sidebar"] +
				b.TemplateList["footer"] + b.TemplateList["js"] + b.TemplateList["menu"] +
				b.TemplateList["admin_panel"] + b.TemplateList["content"])
	} else {
		name = "content"
		tmpl, err = template.New("content").Funcs(adminTemplate.DefaultFuncMap).
			Parse(b.TemplateList["admin_panel"] + b.TemplateList["content"])
	}

	if err != nil {
		panic(err)
	}

	return
}

func GetImportJSTag(src string) template.HTML {
	if config.GetAssetUrl() != "" {
		return template.HTML(`<script src="` + config.GetAssetUrl() + src + `"></script>`)
	} else {
		return template.HTML(`<script src="` + config.Url(src) + `"></script>`)
	}
}

func GetImportCSSTag(src string) template.HTML {
	if config.GetAssetUrl() != "" {
		return template.HTML(`<link rel="stylesheet" href="` + config.GetAssetUrl() + src + `">`)
	} else {
		return template.HTML(`<link rel="stylesheet" href="` + config.Url(src) + `">`)
	}
}
