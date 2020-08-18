package separation

import (
	"io/ioutil"

	"github.com/GoAdminGroup/go-admin/modules/config"
	adminTemplate "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/components"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/themes/common"
	"github.com/GoAdminGroup/themes/sword/resource"
)

type Theme struct {
	ThemeName string
	components.Base
	*common.BaseTheme
}

var Sword = Theme{
	ThemeName: "sword_sep",
	Base: components.Base{
		Attribute: types.Attribute{
			TemplateList: common.SepTemplateList,
			Separation:   true,
		},
	},
	BaseTheme: &common.BaseTheme{
		AssetPaths:   resource.AssetPaths,
		TemplateList: common.SepTemplateList,
		Separation:   true,
	},
}

func init() {
	adminTemplate.Add("sword_sep", &Sword)
}

func Get() *Theme {
	return &Sword
}

func (t *Theme) Name() string {
	return t.ThemeName
}

func (t *Theme) GetTmplList() map[string]string {
	return common.SepTemplateList
}

func (t *Theme) GetAsset(path string) ([]byte, error) {
	return ioutil.ReadFile(config.GetAssetRootPath() + path)
}

func (t *Theme) GetAssetList() []string {
	return resource.AssetsList
}
