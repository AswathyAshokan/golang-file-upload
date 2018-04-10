package routers

import (
	"github.com/astaxie/beego"
	"fileUpload/controllers"

)

func init(){

	beego.Router("/fileUpload",&controllers.FileUploadController{},"*:FileUpload")
	beego.Router("/fileUpload/{:imageUrl}",&controllers.FileUploadController{},"*:FileUpload")

}