package services

import (
	cm "UlboraCmsAdmin/common"
	b64 "encoding/base64"
	"net/http"
	"strings"
)

//ImageService service
type ImageService struct {
	Token     string
	ClientID  string
	APIClient string
	APIKey    string
	UserID    string
	Hashed    string
	Host      string
}

// Image the image info
type Image struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Size          int64  `json:"size"`
	FileExtension string `json:"fileExtension"`
	ClientID      int64  `json:"clientId"`
	ImageURL      string `json:"imageUrl"`
}

// UploadedFile file
type UploadedFile struct {
	Name             string
	Size             int64
	OriginalFileName string
	FileData         []byte
}

type imageFile struct {
	Name          string `json:"name"`
	Size          int64  `json:"size"`
	FileExtension string `json:"fileExtension"`
	FileData      string `json:"fileData"`
}

//ImageResponse res
type ImageResponse struct {
	Success bool   `json:"success"`
	ID      int64  `json:"id"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

//AddImage add image
func (i *ImageService) AddImage(image *UploadedFile) *ImageResponse {
	var rtn = new(ImageResponse)
	var addURL = i.Host + "/rs/image/add"
	s64 := b64.StdEncoding.EncodeToString(image.FileData)
	igf := new(imageFile)
	igf.FileData = s64
	igf.Name = stripSpace(image.Name)
	igf.Size = image.Size
	igf.FileExtension = getExt(stripSpace(image.OriginalFileName))
	aJSON := cm.GetJSONEncode(igf)
	req, fail := cm.GetRequest(addURL, http.MethodPost, aJSON)
	if !fail {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+i.Token)
		req.Header.Set("u-client-id", i.APIClient)
		req.Header.Set("clientId", i.ClientID)
		req.Header.Set("u-api-key", i.APIKey)
		code := cm.ProcessServiceCall(req, &rtn)
		rtn.Code = code
	}
	return rtn
}

//GetList add image
func (i *ImageService) GetList() *[]Image {
	var rtn = make([]Image, 0)
	var gURL = i.Host + "/rs/image/list/100"
	//fmt.Println(gURL)
	req, fail := cm.GetRequest(gURL, http.MethodGet, nil)
	if !fail {
		req.Header.Set("Authorization", "Bearer "+i.Token)
		req.Header.Set("u-client-id", i.APIClient)
		req.Header.Set("clientId", i.ClientID)
		req.Header.Set("u-api-key", i.APIKey)
		cm.ProcessServiceCall(req, &rtn)
	}
	return &rtn
}

// DeleteImage delete image
func (i *ImageService) DeleteImage(id string) *ImageResponse {
	var rtn = new(ImageResponse)
	var gURL = i.Host + "/rs/image/delete/" + id
	//fmt.Println(gURL)
	req, fail := cm.GetRequest(gURL, http.MethodDelete, nil)
	if !fail {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+i.Token)
		req.Header.Set("u-client-id", i.APIClient)
		req.Header.Set("clientId", i.ClientID)
		req.Header.Set("u-api-key", i.APIKey)
		code := cm.ProcessServiceCall(req, &rtn)
		rtn.Code = code
	}
	return rtn
}

func getExt(name string) string {
	var rtn string
	i := strings.LastIndex(name, ".")
	rtn = name[i+1:]
	return rtn
}

func stripSpace(name string) string {
	var rtn string
	rtn = strings.Replace(name, " ", "", -1)
	return rtn
}
