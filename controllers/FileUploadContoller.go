package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"io"
)

type FileUploadController struct {
	BaseController
}

func (c *FileUploadController)FileUpload(){
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	if r.Method == http.MethodPost{
		file,header,err :=r.FormFile("uploadfile")
		if err !=nil{
			log.Println("uploading error",err)
			http.Error(w,"error in uploading file",http.StatusInternalServerError)
			return
		}
		if _, err := os.Stat("./testUploadImage/"); os.IsNotExist(err) {

			os.Mkdir("./testUploadImage/",os.ModePerm)
		}

		f, err := os.OpenFile("./testUploadImage/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}else{
			fmt.Println("image  uploaded")
		}
		fmt.Println("jst")

		io.Copy(f, file)
		defer file.Close()
		fmt.Fprintf(w,"file uploaded")


	}else{
		c.TplName ="templates/fileUpload.html"
	}


}