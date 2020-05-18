package common

import (
	"github.com/GoAdminGroup/go-admin/modules/config"
	"html/template"
	"strings"
)

type BaseTheme struct {
	AssetPaths map[string]string
}

const Version = "v0.0.33"

func (b *BaseTheme) GetVersion() string {
	return Version
}

func (b *BaseTheme) GetRequirements() []string {
	return []string{"v1.2.11"}
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

func (b *BaseTheme) Get404Page() template.HTML   { return "" }
func (b *BaseTheme) GetErrorPage() template.HTML { return "" }

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
