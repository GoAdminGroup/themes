package adminlte

import (
	"github.com/GoAdminGroup/go-admin/modules/language"
	adminTemplate "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/components"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/themes/adminlte/resource"
	"github.com/GoAdminGroup/themes/adminlte/tmpl"
	"html/template"
)

const (
	ColorschemeSkinBlack       = "skin-black"
	ColorschemeSkinBlackLight  = "skin-black-light"
	ColorschemeSkinBlue        = "skin-blue"
	ColorschemeSkinBlueLight   = "skin-blue-light"
	ColorschemeSkinGreen       = "skin-green"
	ColorschemeSkinGreenLight  = "skin-green-light"
	ColorschemeSkinPurple      = "skin-purple"
	ColorschemeSkinPurpleLight = "skin-purple-light"
	ColorschemeSkinRed         = "skin-red"
	ColorschemeSkinRedLight    = "skin-red-light"
	ColorschemeSkinYellow      = "skin-yellow"
	ColorschemeSkinYellowLight = "skin-yellow-light"
)

type Theme struct {
	Name string
	components.Base
}

var Adminlte = Theme{
	Name: "adminlte",
	Base: components.Base{
		Attribute: types.Attribute{
			TemplateList: tmpl.List,
		},
	},
}

func init() {
	adminTemplate.Add("adminlte", &Adminlte)
}

func Get() *Theme {
	return &Adminlte
}

func (*Theme) GetTmplList() map[string]string {
	return tmpl.List
}

func (*Theme) GetTemplate(isPjax bool) (tmpler *template.Template, name string) {
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
		}).Parse(tmpl.List["layout"] +
			tmpl.List["head"] + tmpl.List["header"] + tmpl.List["sidebar"] +
			tmpl.List["footer"] + tmpl.List["js"] + tmpl.List["menu"] +
			tmpl.List["admin_panel"] + tmpl.List["content"])
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
		}).Parse(tmpl.List["admin_panel"] + tmpl.List["content"])
	}

	if err != nil {
		panic(err)
	}

	return
}

func (*Theme) GetAsset(path string) ([]byte, error) {
	path = "template/adminlte/resource" + path
	return resource.Asset(path)
}

func (*Theme) GetAssetList() []string {
	return resource.AssetsList
}
