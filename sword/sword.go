package sword

import (
	adminTemplate "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/components"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/themes/common"
	"github.com/GoAdminGroup/themes/sword/resource"
	"html/template"
)

type Theme struct {
	Name string
	components.Base
	common.BaseTheme
}

var Sword = Theme{
	Name: "sword",
	Base: components.Base{
		Attribute: types.Attribute{
			TemplateList: TemplateList,
		},
	},
}

func init() {
	adminTemplate.Add("sword", &Sword)
}

func Get() *Theme {
	return &Sword
}

func (*Theme) GetTmplList() map[string]string {
	return TemplateList
}

func (*Theme) GetTemplate(isPjax bool) (tmpl *template.Template, name string) {
	var err error

	if !isPjax {
		name = "layout"
		tmpl, err = template.New("layout").Funcs(adminTemplate.DefaultFuncMap).
			Parse(TemplateList["layout"] +
				TemplateList["head"] + TemplateList["header"] + TemplateList["sidebar"] +
				TemplateList["footer"] + TemplateList["js"] + TemplateList["menu"] +
				TemplateList["admin_panel"] + TemplateList["content"])
	} else {
		name = "content"
		tmpl, err = template.New("content").Funcs(adminTemplate.DefaultFuncMap).
			Parse(TemplateList["admin_panel"] + TemplateList["content"])
	}

	if err != nil {
		panic(err)
	}

	return
}

func (*Theme) GetAsset(path string) ([]byte, error) {
	path = "resource" + path
	return resource.Asset(path)
}

func (*Theme) GetAssetList() []string {
	return resource.AssetsList
}
