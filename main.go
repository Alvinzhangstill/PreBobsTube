package main

import (
	_ "github.com/MobileCPX/PreBobsTube/initial"
	_ "github.com/MobileCPX/PreBobsTube/routers"
	"github.com/astaxie/beego"
)

func main() {
	//initial.InitDatabase()
	beego.Run()
}
