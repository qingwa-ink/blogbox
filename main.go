package main

import (
	_ "blog/models"
	_ "blog/routers"
	_ "blog/services"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
