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
	"strconv"
	"strings"
	//"fmt"
	//"fmt"
	"net/http"
)

//HandleAddContent HandleAddContent
func (h *Handler) HandleAddContent(w http.ResponseWriter, r *http.Request) {
	h.Sess.InitSessionStore(w, r)
	session := h.getSession(w, r)
	loggedIn := session.Values["userLoggenIn"]
	token := h.getToken(w, r)
	//fmt.Print("loggedIn in addcontent: ")
	//fmt.Println(loggedIn)
	if loggedIn == nil || !loggedIn.(bool) || token == nil {
		h.loginImplicit(w, r)
	} else {
		clientID := session.Values["clientId"].(string)
		var i services.ImageService
		i.ClientID = clientID
		i.APIClient = getGatewayAPIClient()
		i.APIKey = getGatewayAPIKey()
		i.Token = token.AccessToken
		//fmt.Println(token.AccessToken)
		i.Host = getImageHost()

		res := i.GetList()

		h.Templates.ExecuteTemplate(w, "addContent.html", &res)
	}
}

//HandleNewContent HandleNewContent
func (h *Handler) HandleNewContent(w http.ResponseWriter, r *http.Request) {
	h.Sess.InitSessionStore(w, r)
	session := h.getSession(w, r)
	loggedIn := session.Values["userLoggenIn"]
	token := h.getToken(w, r)
	//fmt.Print("loggedIn in newcontent: ")
	//fmt.Println(loggedIn)
	if loggedIn == nil || !loggedIn.(bool) || token == nil {
		h.loginImplicit(w, r)
	} else {
		clientID := session.Values["clientId"].(string)
		content := r.FormValue("content")
		//fmt.Print("content: ")
		//fmt.Println(content)

		title := r.FormValue("title")
		//fmt.Print("title: ")
		//fmt.Println(title)

		author := r.FormValue("author")
		//fmt.Print("author: ")
		//fmt.Println(author)

		category := r.FormValue("category")
		category = strings.Replace(category, " ", "", -1)
		//fmt.Print("category: ")
		//fmt.Println(category)

		sortOrder := r.FormValue("sortOrder")
		if sortOrder == "" {
			sortOrder = "0"
		}
		//fmt.Print("sortOrder: ")
		//fmt.Println(sortOrder)

		metaKeyWords := r.FormValue("metaKeyWords")
		//fmt.Print("metaKeyWords: ")
		//fmt.Println(metaKeyWords)

		desc := r.FormValue("desc")
		//fmt.Print("desc: ")
		//fmt.Println(desc)
		var ct services.Content
		ct.Text = content
		ct.Title = title
		ct.MetaAuthorName = author
		ct.Category = category
		ct.MetaKeyWords = metaKeyWords
		ct.MetaRobotKeyWords = metaKeyWords
		ct.MetaDesc = desc
		ct.SortOrder, _ = strconv.Atoi(sortOrder)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		//fmt.Println(token.AccessToken)
		var c services.ContentService
		c.ClientID = clientID
		c.APIClient = getGatewayAPIClient()
		c.APIKey = getGatewayAPIKey()

		c.Token = token.AccessToken
		c.Host = getContentHost()
		var res *services.Response
		res = c.AddContent(&ct)
		//fmt.Print("res: ")
		//fmt.Println(res)
		if res.Success || testMode {
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			http.Redirect(w, r, "/addContent", http.StatusFound)
		}
	}
}

//HandleUpdateContent HandleUpdateContent
func (h *Handler) HandleUpdateContent(w http.ResponseWriter, r *http.Request) {
	h.Sess.InitSessionStore(w, r)
	session := h.getSession(w, r)
	loggedIn := session.Values["userLoggenIn"]
	token := h.getToken(w, r)
	if loggedIn == nil || !loggedIn.(bool) || token == nil {
		h.loginImplicit(w, r)
	} else {
		clientID := session.Values["clientId"].(string)
		idStr := r.FormValue("id")
		id, _ := strconv.ParseInt(idStr, 10, 0)
		// if errID != nil {
		// 	fmt.Print(errID)
		// }
		//fmt.Print("id: ")
		//fmt.Println(id)

		content := r.FormValue("content")
		//fmt.Print("content: ")
		//fmt.Println(content)

		title := r.FormValue("title")
		//fmt.Print("title: ")
		//fmt.Println(title)

		author := r.FormValue("author")
		//fmt.Print("author: ")
		//fmt.Println(author)

		category := r.FormValue("category")
		category = strings.Replace(category, " ", "", -1)
		//fmt.Print("category: ")
		//fmt.Println(category)

		sortOrder := r.FormValue("sortOrder")
		if sortOrder == "" {
			sortOrder = "0"
		}
		//fmt.Print("sortOrder: ")
		//fmt.Println(sortOrder)

		metaKeyWords := r.FormValue("metaKeyWords")
		//fmt.Print("metaKeyWords: ")
		//fmt.Println(metaKeyWords)

		desc := r.FormValue("desc")
		//fmt.Print("desc: ")
		//fmt.Println(desc)

		archived := r.FormValue("archived")
		//fmt.Print("archived: ")
		//fmt.Println(archived)

		var ct services.Content
		ct.ID = id
		ct.Text = content
		ct.Title = title
		ct.MetaAuthorName = author
		ct.Category = category
		ct.MetaKeyWords = metaKeyWords
		ct.MetaRobotKeyWords = metaKeyWords
		ct.MetaDesc = desc
		ct.SortOrder, _ = strconv.Atoi(sortOrder)
		// if err != nil {
		// 	fmt.Print("sortOrder conversion error: ")
		// 	fmt.Println(err)
		// }
		if archived == "on" {
			ct.Archived = true
		} else {
			ct.Archived = false
		}

		var c services.ContentService
		c.ClientID = clientID
		c.APIClient = getGatewayAPIClient()
		c.APIKey = getGatewayAPIKey()

		c.Token = token.AccessToken
		c.Host = getContentHost()

		var res *services.Response

		res = c.UpdateContent(&ct)

		//fmt.Println(res)
		if res.Success || testMode {
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			fmt.Println("Content update failed")
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}
}

//HandleGetContent HandleGetContent
func (h *Handler) HandleGetContent(w http.ResponseWriter, r *http.Request) {
	h.Sess.InitSessionStore(w, r)
	session := h.getSession(w, r)
	loggedIn := session.Values["userLoggenIn"]
	token := h.getToken(w, r)
	if loggedIn == nil || !loggedIn.(bool) || token == nil {
		h.loginImplicit(w, r)
	} else {
		clientID := session.Values["clientId"].(string)

		id := r.URL.Query().Get("id")
		var c services.ContentService
		c.ClientID = clientID
		c.APIClient = getGatewayAPIClient()
		c.APIKey = getGatewayAPIKey()
		c.Host = getContentHost()
		res := c.GetContent(id, clientID)

		var i services.ImageService
		i.ClientID = clientID
		i.APIClient = getGatewayAPIClient()
		i.APIKey = getGatewayAPIKey()

		i.Token = token.AccessToken
		//fmt.Println(token.AccessToken)
		i.Host = getImageHost()

		ires := i.GetList()

		var ci = new(contentAndImages)
		ci.Cont = res
		ci.Img = ires

		h.Templates.ExecuteTemplate(w, "updateContent.html", &ci)
	}
}

//HandleDeleteContent HandleDeleteContent
func (h *Handler) HandleDeleteContent(w http.ResponseWriter, r *http.Request) {
	h.Sess.InitSessionStore(w, r)
	session := h.getSession(w, r)
	loggedIn := session.Values["userLoggenIn"]
	token := h.getToken(w, r)
	if loggedIn == nil || !loggedIn.(bool) || token == nil {
		h.loginImplicit(w, r)
	} else {
		id := r.URL.Query().Get("id")
		clientID := session.Values["clientId"].(string)
		var c services.ContentService
		c.ClientID = clientID
		c.APIClient = getGatewayAPIClient()
		c.APIKey = getGatewayAPIKey()

		c.Token = token.AccessToken
		c.Host = getContentHost()
		//res := c.DeleteContent(id)
		var res *services.Response
		res = c.DeleteContent(id)

		if !res.Success {
			fmt.Println("Delete content failed on ID: " + id)
			fmt.Print("code: ")
			fmt.Println(res.Code)
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
