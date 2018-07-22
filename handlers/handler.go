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
	services "UlboraCmsV3/services"
	"fmt"
	usession "github.com/Ulbora/go-better-sessions"
	oauth2 "github.com/Ulbora/go-oauth2-client"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
	"os"
)

//Handler Handler
type Handler struct {
	Sess      usession.Session
	TokenMap  map[string]*oauth2.Token
	Templates *template.Template
}

type contentAndImages struct {
	Cont *services.Content
	Img  *[]services.Image
}

func (h *Handler) getToken(w http.ResponseWriter, r *http.Request) *oauth2.Token {
	session := h.getSession(w, r)
	var token *oauth2.Token
	if tokenKey := session.Values["accessTokenKey"]; tokenKey != nil {
		token = h.TokenMap[tokenKey.(string)]
	}
	return token
}

func (h *Handler) removeToken(w http.ResponseWriter, r *http.Request) {
	session := h.getSession(w, r)
	tokenKey := session.Values["accessTokenKey"]
	delete(h.TokenMap, tokenKey.(string))
}

func getOauthHost() string {
	var rtn = ""
	if os.Getenv("AUTH_HOST") != "" {
		rtn = os.Getenv("AUTH_HOST")
	} else {
		rtn = "http://localhost:3000"
	}
	return rtn
}

func getOauthRedirectHost() string {
	var rtn = ""
	if os.Getenv("AUTH_REDIRECT_HOST") != "" {
		rtn = os.Getenv("AUTH_REDIRECT_HOST")
	} else {
		rtn = "http://localhost:3000"
	}
	return rtn
}

func getRedirectURI(req *http.Request, path string) string {
	var scheme = req.URL.Scheme
	var serverHost string
	if scheme != "" {
		serverHost = req.URL.String()
	} else {
		serverHost = schemeDefault + req.Host + path
	}
	return serverHost
}

func (h *Handler) getSession(w http.ResponseWriter, r *http.Request) *sessions.Session {
	session, err := h.Sess.GetSession(r)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return session
}
