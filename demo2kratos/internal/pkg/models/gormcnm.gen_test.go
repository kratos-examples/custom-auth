package models_test

import (
	"testing"

	"github.com/yylego/gormcngen"
	"github.com/yylego/kratos-examples/demo2kratos/internal/pkg/models"
	"github.com/yylego/osexistpath/osmustexist"
	"github.com/yylego/runpath/runtestpath"
)

// Auto generate columns with go generate command
// Support execution via: go generate ./...
// Delete this comment block if auto generation is not needed
//
// 使用 go generate 命令自动生成列定义
// 支持通过以下命令执行：go generate ./...
// 如果不需要自动生成，可以删除此注释块
//
//go:generate go test -v -run TestGenerateColumns
func TestGenerateColumns(t *testing.T) {
	absPath := osmustexist.FILE(runtestpath.SrcPath(t))
	t.Log(absPath)

	objects := []any{
		&models.Admin{},
		&models.Token{},
	}

	options := gormcngen.NewOptions().
		WithColumnClassExportable(true).
		WithColumnsMethodRecvName("c").
		WithColumnsCheckFieldType(true)

	cfg := gormcngen.NewConfigs(objects, options, absPath)
	cfg.Gen()
}
