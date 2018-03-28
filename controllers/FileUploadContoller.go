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
	"image/jpeg"
	"github.com/nfnt/resize"
	"image/gif"
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


		msec := strconv.FormatInt(time.Now().UnixNano() / 1000000,10)

		//for processing json files
		if mimeType =="application/octet-stream"{
			fmt.Println("inside json")


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

		f, err := os.OpenFile("./testUploadImage/"+msec+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("image 4 error",err)
			return
		}
		fmt.Println("jst")

		io.Copy(f, file)
		defer file.Close()
		fmt.Fprintf(w,"file  uploaded")
	//for creating thumbnails of uploading image


		if mimeType =="image/jpeg" {
			thumbfile, err := os.Open("./testUploadImage/" + msec + header.Filename)
			if err != nil {
				fmt.Println("the error5")
				log.Fatal(err)
			}

			// decode jpeg into image.Image
			img, err := jpeg.Decode(thumbfile)
			if err != nil {
				fmt.Println("the error2")
				log.Fatal(err)

			}
			thumbfile.Close()

			// resize to width 100  and height 100 using Lanczos resampling
			// and preserve aspect ratio
			m := resize.Resize(100, 100, img, resize.Lanczos3)

			out, err := os.Create("./testUploadImage/" + "thumb" + msec + header.Filename)
			if err != nil {
				fmt.Println("the error1")
				log.Fatal(err)
			}
			defer out.Close()

			// write new image to file
			jpeg.Encode(out, m, nil)
		}else if mimeType =="image/png"{
			thumbfile, err := os.Open("./testUploadImage/" + msec + header.Filename)
			if err != nil {
				fmt.Println("the error5")
				log.Fatal(err)
			}

			// decode png into image.Image
			img, err := png.Decode(thumbfile)
			if err != nil {
				fmt.Println("the error2")
				log.Fatal(err)

			}
			thumbfile.Close()

			// resize to width 100  and height 100 using Lanczos resampling
			// and preserve aspect ratio
			m := resize.Resize(100, 100, img, resize.Lanczos3)

			out, err := os.Create("./testUploadImage/" + "thumb" + msec + header.Filename)
			if err != nil {
				fmt.Println("the error1")
				log.Fatal(err)
			}
			defer out.Close()

			// write new image to file
			png.Encode(out, m)
		}else if mimeType =="image/gif" {
			thumbfile, err := os.Open("./testUploadImage/" + msec + header.Filename)
			if err != nil {
				fmt.Println("the error5")
				log.Fatal(err)
			}

			// decode png into image.Image
			img, err := gif.Decode(thumbfile)
			if err != nil {
				fmt.Println("the error2")
				log.Fatal(err)

			}
			thumbfile.Close()

			// resize to width 100  and height 100 using Lanczos resampling
			// and preserve aspect ratio
			m := resize.Resize(100, 100, img, resize.Lanczos3)

			out, err := os.Create("./testUploadImage/" + "thumb" + msec + header.Filename)
			if err != nil {
				fmt.Println("the error1")
				log.Fatal(err)
			}
			defer out.Close()

			// write new image to file
			gif.Encode(out, m,nil)


		}
		}else{
		c.TplName ="templates/fileUpload.html"
	}


}