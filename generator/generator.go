package generator

import "github.com/luenci/golang-project-template/generator/content"

// TemplateMap is the map of templates register.
var (
	TemplateMap = map[string]string{
		".pre-commit-config.yaml": content.PreCommit,
		".gitignore":              content.GitIgnore,
		"Dockerfile":              content.Dockerfile,
		"main.go":                 content.MainCLT,
	}

	ArgsTemplateMap = map[string]string{
		"go.mod":   content.Module,
		"Makefile": content.Makefile,
	}
)
