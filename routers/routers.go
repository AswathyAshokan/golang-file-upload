package routers

import (
	"github.com/astaxie/beego"
	"fileUpload/controllers"
	"fmt"
)

func init(){
	fmt.Println("jjjjj")
	beego.Router("/fileUpload",&controllers.FileUploadController{},"*:FileUpload")
}