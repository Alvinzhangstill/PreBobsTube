package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	c.Layout = "temp/layout/layout_page.html"
	c.TplName = "temp/pages/home.html"

}
