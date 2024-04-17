package resource

import "embed"

//go:embed assets/dist/*
var AssetFS embed.FS
