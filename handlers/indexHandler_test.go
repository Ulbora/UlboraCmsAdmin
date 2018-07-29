package handlers

import (
	usession "github.com/Ulbora/go-better-sessions"
	oauth2 "github.com/Ulbora/go-oauth2-client"
	"html/template"
	//"html/template"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_handleAdminIndex(t *testing.T) {
	var h Handler
	h.Templates = template.Must(template.ParseFiles("index.html"))
	h.TokenMap = make(map[string]*oauth2.Token)
	var s usession.Session
	h.Sess = s
	r, _ := http.NewRequest("GET", "/challenge?route=challenge&fpath=rs/challenge/en_us?g=g&b=b", nil)
	w := httptest.NewRecorder()
	h.Sess.InitSessionStore(w, r)
	session, _ := h.Sess.GetSession(r)
	session.Values["accessTokenKey"] = "123456"
	var resp oauth2.Token
	resp.AccessToken = "bbbnn"
	h.TokenMap["123456"] = &resp
	h.HandleAdminIndex(w, r)
}

func TestHandler_handleAdminIndex2(t *testing.T) {
	var h Handler
	h.Templates = template.Must(template.ParseFiles("index.html"))
	h.TokenMap = make(map[string]*oauth2.Token)
	var s usession.Session
	h.Sess = s
	r, _ := http.NewRequest("GET", "/challenge?route=challenge&fpath=rs/challenge/en_us?g=g&b=b", nil)
	w := httptest.NewRecorder()
	h.Sess.InitSessionStore(w, r)
	session, _ := h.Sess.GetSession(r)
	session.Values["accessTokenKey"] = "123456"
	session.Values["userLoggenIn"] = true
	session.Values["clientId"] = "123"
	var resp oauth2.Token
	resp.AccessToken = "bbbnn"
	h.TokenMap["123456"] = &resp
	h.HandleAdminIndex(w, r)
}
