package handlers

/*
 Copyright (C) 2018 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2018 Ken Williamson
 All rights reserved.

 Certain inventions and disclosures in this file may be claimed within
 patents owned or patent applications filed by Ulbora Labs LLC., or third
 parties.

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU Affero General Public License as published
 by the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.

 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU Affero General Public License for more details.

 You should have received a copy of the GNU Affero General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

import (
	services "UlboraCmsAdmin/services"
	"fmt"
	"io/ioutil"
	"net/http"
)

//HandleAddImage HandleAddImage
func (h *Handler) HandleAddImage(w http.ResponseWriter, r *http.Request) {
	h.Sess.InitSessionStore(w, r)
	session := h.getSession(w, r)
	loggedIn := session.Values["userLoggenIn"]
	token := h.getToken(w, r)
	if loggedIn == nil || !loggedIn.(bool) || token == nil {
		h.loginImplicit(w, r)
	} else {
		h.Templates.ExecuteTemplate(w, "imageUpload.html", nil)
	}
}

//HandleImagerUpload HandleImagerUpload
func (h *Handler) HandleImagerUpload(w http.ResponseWriter, r *http.Request) {
	h.Sess.InitSessionStore(w, r)
	session := h.getSession(w, r)
	loggedIn := session.Values["userLoggenIn"]
	token := h.getToken(w, r)
	if loggedIn == nil || !loggedIn.(bool) || token == nil {
		h.loginImplicit(w, r)
	} else {
		clientID := session.Values["clientId"].(string)
		name := r.FormValue("name")
		//fmt.Print("name: ")
		//fmt.Println(name)
		r.ParseMultipartForm(2000000)
		// err := r.ParseMultipartForm(2000000)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		//fmt.Println("before file: ")
		file, handler, _ := r.FormFile("image")
		//fmt.Print("file: ")
		//fmt.Println(file)
		// file, handler, err := r.FormFile("image")
		// if err != nil {
		// 	fmt.Println(err)
		// }
		defer file.Close()

		//fmt.Print("name: ")
		//fmt.Println(handler.Filename)

		// cur, err := file.Seek(0, 1)
		// size, err := file.Seek(0, 2)
		// _, err1 := file.Seek(cur, 0)
		// if err1 != nil {
		// 	fmt.Println(err1)
		// }
		cur, _ := file.Seek(0, 1)
		size, _ := file.Seek(0, 2)
		file.Seek(cur, 0)

		// data, err := ioutil.ReadAll(file)
		// if err != nil {
		// 	fmt.Println(err)
		// }

		data, _ := ioutil.ReadAll(file)

		// fmt.Print("file data: ")
		// fmt.Println(data)

		// fmt.Print("file size: ")
		// fmt.Println(size)

		var i services.ImageService
		i.ClientID = clientID
		i.APIClient = getGatewayAPIClient()
		i.APIKey = getGatewayAPIKey()
		i.Host = getImageHost()
		i.Token = token.AccessToken
		var img services.UploadedFile
		img.Name = name
		img.OriginalFileName = handler.Filename
		img.Size = size
		img.FileData = data
		//fmt.Print("file img: ")
		//fmt.Println(img)
		var res *services.ImageResponse
		res = i.AddImage(&img)
		//fmt.Print("file res: ")
		//fmt.Println(res)
		//res := i.AddImage(&img)
		if res.Success || testMode {
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			fmt.Println("Image upload failed")
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}
}

//HandleImages HandleImages
func (h *Handler) HandleImages(w http.ResponseWriter, r *http.Request) {
	h.Sess.InitSessionStore(w, r)
	session := h.getSession(w, r)
	loggedIn := session.Values["userLoggenIn"]
	tokeni := h.getToken(w, r)
	if loggedIn == nil || !loggedIn.(bool) || tokeni == nil {
		h.loginImplicit(w, r)
	} else {
		clientID := session.Values["clientId"].(string)
		var i services.ImageService
		i.ClientID = clientID
		i.APIClient = getGatewayAPIClient()
		i.APIKey = getGatewayAPIKey()
		i.Host = getImageHost()
		i.Token = tokeni.AccessToken

		res := i.GetList()

		h.Templates.ExecuteTemplate(w, "images.html", &res)
	}
}

//HandleDeleteImage HandleDeleteImage
func (h *Handler) HandleDeleteImage(w http.ResponseWriter, r *http.Request) {
	h.Sess.InitSessionStore(w, r)
	session := h.getSession(w, r)
	loggedIn := session.Values["userLoggenIn"]
	token := h.getToken(w, r)
	if loggedIn == nil || !loggedIn.(bool) || token == nil {
		h.loginImplicit(w, r)
	} else {
		clientID := session.Values["clientId"].(string)
		id := r.URL.Query().Get("id") // r.FormValue("id")
		//id, _ := strconv.ParseInt(idStr, 10, 0)
		var i services.ImageService
		i.ClientID = clientID
		i.APIClient = getGatewayAPIClient()
		i.APIKey = getGatewayAPIKey()
		i.Host = getImageHost()
		i.Token = token.AccessToken
		var res *services.ImageResponse
		res = i.DeleteImage(id)

		//res := i.DeleteImage(id)
		if !res.Success || testMode {
			fmt.Println("Delete image failed on ID: " + id)
			fmt.Print("code: ")
			fmt.Println(res.Code)
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
