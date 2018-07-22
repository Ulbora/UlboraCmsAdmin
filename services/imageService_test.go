package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"testing"
)

var imgID int64

//var imgToken = testToken

func TestImageService_getToken(t *testing.T) {
	if testToken == "" {
		req, _ := http.NewRequest("POST", tokenURL, nil)
		client := &http.Client{}
		resp, cErr := client.Do(req)
		if cErr != nil {
			fmt.Print("Client Add err: ")
			fmt.Println(cErr)
		} else {
			defer resp.Body.Close()
			var tres TokenResponse
			decoder := json.NewDecoder(resp.Body)
			error := decoder.Decode(&tres)
			if error != nil {
				log.Println(error.Error())
			} else {
				testToken = tres.Token
				//fmt.Print("token: ")
				//fmt.Println(testToken)
			}
		}
	}
}

func TestImageService_getExt(t *testing.T) {
	var name = "tester.jpg"
	ext := getExt(name)
	fmt.Print("ext: ")
	fmt.Println(ext)
	if ext != "jpg" {
		t.Fail()
	}
}

func TestImageService_stripSpace(t *testing.T) {
	var name = "tes ter .jpg"
	fn := stripSpace(name)
	fmt.Print("name: ")
	fmt.Println(fn)
	if fn != "tester.jpg" {
		t.Fail()
	}
}

func TestImageService_AddImage(t *testing.T) {
	var i ImageService
	i.ClientID = "403"
	i.Host = "http://localhost:3007"
	i.Token = testToken

	imgfile, err := os.Open("./testFiles/upload/test.jpg")
	if err != nil {
		fmt.Println("jpg file not found!")
		os.Exit(1)
	}
	defer imgfile.Close()

	var ui UploadedFile
	ui.Name = "testfile"
	ui.OriginalFileName = imgfile.Name()
	data, err := ioutil.ReadAll(imgfile)
	if err != nil {
		fmt.Println(err)
	}

	cur, err := imgfile.Seek(0, 1)
	if err != nil {
		fmt.Println(err)
	}
	size, err2 := imgfile.Seek(0, 2)
	if err2 != nil {
		fmt.Println(err)
	}
	_, err1 := imgfile.Seek(cur, 0)
	if err1 != nil {
		fmt.Println(err1)
	}

	ui.Size = size
	ui.FileData = data

	res := i.AddImage(&ui)
	fmt.Print("res: ")
	fmt.Println(res)

	if res.Success != true {
		t.Fail()
	} else {
		imgID = res.ID
	}
}

func TestImageService_GetList(t *testing.T) {
	var i ImageService
	i.ClientID = "403"
	i.Host = "http://localhost:3007"
	i.Token = testToken
	res := i.GetList()
	fmt.Print("res: ")
	fmt.Println(res)
	if res == nil {
		t.Fail()
	}
}

func TestImageService_GetList_DeleteImage(t *testing.T) {
	var i ImageService
	i.ClientID = "403"
	i.Host = "http://localhost:3007"
	i.Token = testToken

	res := i.DeleteImage(strconv.FormatInt(imgID, 10))
	fmt.Print("res deleted: ")
	fmt.Println(res)
	addID = res.ID
	if res.Success != true {
		t.Fail()
	}
}
