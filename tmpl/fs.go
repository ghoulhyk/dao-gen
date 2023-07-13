package tmpl

import (
	"embed"
)

//go:embed fixed
var FixedTemplateFs embed.FS

//go:embed fragment
var FragmentTemplateFs embed.FS

//go:embed orm
var OrmTemplateFs embed.FS
