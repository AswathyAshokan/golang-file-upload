package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"io"
	"io/ioutil"
	"encoding/json"
	"encoding/base64"
	"image/png"
	"time"
	"strconv"
	"bytes"
)

type FileUploadController struct {
	BaseController
}

func (c *FileUploadController)FileUpload(){
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	type Config struct {
		Database struct {
			Picture     string `json:"picture"`

		} `json:"database"`

	}
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
		mimeType := header.Header.Get("Content-Type")
		fmt.Println("typeeeeee",mimeType)
		if mimeType =="application/octet-stream"{


			if _, err := os.Stat("./testUploadJson/"); os.IsNotExist(err) {

				os.Mkdir("./testUploadJson/",os.ModePerm)
			}
			jsonFile, err := os.OpenFile("./testUploadJson/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err)
				return
			}
			io.Copy(jsonFile, file)
			defer file.Close()

			file, e := ioutil.ReadFile("./testUploadJson/"+header.Filename)
			if e != nil {
				fmt.Println("File error: %v\n", e)
				os.Exit(1)
			}
			fmt.Println("%s\n", string(file))


			var config Config
			configFile, err := os.Open("./testUploadJson/"+header.Filename)
			defer configFile.Close()
			if err != nil {
				fmt.Println(err.Error())
			}

			jsonParser := json.NewDecoder(configFile)
			jsonParser.Decode(&config)
			msec := strconv.FormatInt(time.Now().UnixNano() / 1000000,10)


			fmt.Println("file we got is ",config.Database.Picture)
			unbased, err := base64.StdEncoding.DecodeString(config.Database.Picture)
			if err != nil {
				fmt.Println("error",err)
				panic("Cannot decode b64")
			}

			r := bytes.NewReader(unbased)
			im, err := png.Decode(r)
			if err != nil {
				fmt.Println("error",err)
				panic("Bad png")
			}

			f, err := os.OpenFile("./testUploadImage/"+"example"+msec+".png", os.O_WRONLY|os.O_CREATE, 0777)
			if err != nil {
				fmt.Println("error",err)
				panic("Cannot open file")
			}

			png.Encode(f, im)


		}

		f, err := os.OpenFile("./testUploadImage/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("image 4 error",err)
			return
		}
		fmt.Println("jst")

		io.Copy(f, file)
		defer file.Close()
		fmt.Fprintf(w,"file  uploaded")


	}else{
		c.TplName ="templates/fileUpload.html"
	}


}