package handlers

import (
	usession "github.com/Ulbora/go-better-sessions"
	oauth2 "github.com/Ulbora/go-oauth2-client"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_handleLogin(t *testing.T) {
	var h Handler
	h.Templates = template.Must(template.ParseFiles("login.html"))
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/test", nil)
	h.handleLogin(w, r)
}

func TestHandler_handleLogout(t *testing.T) {
	var h Handler
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
	h.HandleLogout(w, r)
}

func TestHandler_loginImplicit(t *testing.T) {
	var h Handler
	h.Templates = template.Must(template.ParseFiles("login.html"))
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/test", nil)
	h.loginImplicit(w, r)
}

func TestHandler_handleImplicitLogin(t *testing.T) {
	var h Handler
	h.TokenMap = make(map[string]*oauth2.Token)
	var s usession.Session
	h.Sess = s
	r, _ := http.NewRequest("GET", "/challenge?clientId=12345", nil)
	w := httptest.NewRecorder()
	h.HandleImplicitLogin(w, r)
}

func TestHandler_handleImplicitToken(t *testing.T) {
	var h Handler
	h.TokenMap = make(map[string]*oauth2.Token)
	var s usession.Session
	h.Sess = s
	r, _ := http.NewRequest("GET", "/challenge?token=ghj555&state=ghh66555h", nil)
	w := httptest.NewRecorder()
	h.Sess.InitSessionStore(w, r)
	h.HandleImplicitToken(w, r)
}
