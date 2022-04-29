package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "相关接口文档",
        "title": "gin示例 API",
        "contact": {},
        "license": {},
        "version": "0.0.1"
    },
    "host": "127.0.0.1:8080",
    "paths": {}
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

type s struct{}

var SwaggerInfo swaggerInfo

func (s *s) ReadDoc() string {

	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}
func init() {
	swag.Register(swag.Name, &s{})
}
