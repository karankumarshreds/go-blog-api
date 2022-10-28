package main

import (
	"github.com/karankumarshreds/go-blog-api/cmd/app"
	_ "github.com/karankumarshreds/go-blog-api/docs"
)

func main() {
	a := app.App{}
	a.Init()
	a.Start()
}
