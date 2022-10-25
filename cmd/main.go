package main

import "github.com/karankumarshreds/go-blog-api/cmd/app"

func main() {
	a := app.App{}
	a.Init()
	a.Start()
}
