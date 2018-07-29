package handlers

/*
 Copyright (C) 2017 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2017 Ken Williamson
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
	"fmt"
	"net/http"

	oauth2 "github.com/Ulbora/go-oauth2-client"
)

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	//s.InitSessionStore(w, r)
	h.loginImplicit(w, r)
}

//HandleLogout HandleLogout
func (h *Handler) HandleLogout(w http.ResponseWriter, r *http.Request) {
	h.removeToken(w, r)
	cookie := &http.Cookie{
		Name:   "ucms-user-session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)

	cookie2 := &http.Cookie{
		Name:   "ulbora_oauth2_server",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie2)
	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *Handler) loginImplicit(w http.ResponseWriter, r *http.Request) {
	//s.InitSessionStore(w, r)
	h.Templates.ExecuteTemplate(w, "login.html", nil)
}

//HandleImplicitLogin HandleImplicitLogin
func (h *Handler) HandleImplicitLogin(w http.ResponseWriter, r *http.Request) {
	h.Sess.InitSessionStore(w, r)
	clientID := r.FormValue("clientId")
	session := h.getSession(w, r)
	// if err != nil {
	// 	fmt.Println(err)
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// } else {
	session.Values["clientId"] = clientID
	//session.Values["testBool"] = true
	serr := session.Save(r, w)
	fmt.Println(serr)
	fmt.Print("clientId: ")
	fmt.Println(clientID)

	fmt.Print("clientId from session: ")
	fmt.Println(session.Values["clientId"].(string))

	var a oauth2.ImplicitAuthorize
	a.ClientID = clientID
	a.OauthHost = getOauthRedirectHost()
	a.RedirectURI = getRedirectURI(r, implicitRedirectURI)
	a.Scope = "write"
	a.State = authCodeState
	a.Req = r
	a.Res = w
	resp := a.ImplicitAuthorize()
	//fmt.Print("RedirectURI: ")
	//fmt.Println(a.RedirectURI)
	//if resp != true {
	fmt.Print("Authorize: ")
	fmt.Println(resp)
	//}
	//fmt.Print("Resp: ")
	//fmt.Println(resp)
	//}
}

//HandleImplicitToken HandleImplicitToken
func (h *Handler) HandleImplicitToken(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	state := r.URL.Query().Get("state")
	//fmt.Println("handle token")
	if state == authCodeState && token != "" {
		//if token != "" {
		//fmt.Println(token)
		//token = resp
		session := h.getSession(w, r)
		// if err != nil {
		// 	fmt.Println(err)
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// } else {
		//session.Values["clientId"] = "616"
		//session.Values["testBool"] = true
		session.Values["userLoggenIn"] = true
		var accKey = generateTokenKey()
		session.Values["accessTokenKey"] = accKey
		var resp oauth2.Token
		resp.AccessToken = token
		h.TokenMap[accKey] = &resp
		//fmt.Print("session id: ")
		//fmt.Println(session.ID)
		err := session.Save(r, w)
		loggedIn := session.Values["userLoggenIn"]
		fmt.Print("loggedIn in index: ")
		fmt.Println(loggedIn)
		fmt.Println(err)
		http.Redirect(w, r, "/", http.StatusFound)
		// decode token and get user id
		//}
		//}
	}
}
