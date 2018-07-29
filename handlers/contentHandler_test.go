package handlers

import (
	usession "github.com/Ulbora/go-better-sessions"
	oauth2 "github.com/Ulbora/go-oauth2-client"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_handleAddContent(t *testing.T) {
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
	h.HandleAddContent(w, r)
}

func TestHandler_handleAddContent2(t *testing.T) {
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
	h.HandleAddContent(w, r)
}

func TestHandler_handleNewContent(t *testing.T) {
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
	h.HandleNewContent(w, r)
}

func TestHandler_handleNewContent2(t *testing.T) {
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
	h.HandleNewContent(w, r)
}

func TestHandler_handleNewContent3(t *testing.T) {
	testMode = true
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
	h.HandleNewContent(w, r)
	testMode = false
}

func TestHandler_handleUpdateContent(t *testing.T) {
	//testMode = true
	var h Handler
	h.Templates = template.Must(template.ParseFiles("index.html"))
	h.TokenMap = make(map[string]*oauth2.Token)
	var s usession.Session
	h.Sess = s
	r, _ := http.NewRequest("GET", "/challenge?archived=on&fpath=rs/challenge/en_us?g=g&b=b", nil)
	w := httptest.NewRecorder()
	h.Sess.InitSessionStore(w, r)
	session, _ := h.Sess.GetSession(r)
	session.Values["accessTokenKey"] = "123456"
	//session.Values["userLoggenIn"] = true
	session.Values["clientId"] = "123"
	var resp oauth2.Token
	resp.AccessToken = "bbbnn"
	h.TokenMap["123456"] = &resp
	h.HandleUpdateContent(w, r)
	testMode = false
}

func TestHandler_handleUpdateContent2(t *testing.T) {
	testMode = true
	var h Handler
	h.Templates = template.Must(template.ParseFiles("index.html"))
	h.TokenMap = make(map[string]*oauth2.Token)
	var s usession.Session
	h.Sess = s
	r, _ := http.NewRequest("GET", "/challenge?archived=on&fpath=rs/challenge/en_us?g=g&b=b", nil)
	w := httptest.NewRecorder()
	h.Sess.InitSessionStore(w, r)
	session, _ := h.Sess.GetSession(r)
	session.Values["accessTokenKey"] = "123456"
	session.Values["userLoggenIn"] = true
	session.Values["clientId"] = "123"
	var resp oauth2.Token
	resp.AccessToken = "bbbnn"
	h.TokenMap["123456"] = &resp
	h.HandleUpdateContent(w, r)
	testMode = false
}

func TestHandler_handleUpdateContent3(t *testing.T) {
	//testMode = true
	var h Handler
	h.Templates = template.Must(template.ParseFiles("index.html"))
	h.TokenMap = make(map[string]*oauth2.Token)
	var s usession.Session
	h.Sess = s
	r, _ := http.NewRequest("GET", "/challenge?archived=off&fpath=rs/challenge/en_us?g=g&b=b", nil)
	w := httptest.NewRecorder()
	h.Sess.InitSessionStore(w, r)
	session, _ := h.Sess.GetSession(r)
	session.Values["accessTokenKey"] = "123456"
	session.Values["userLoggenIn"] = true
	session.Values["clientId"] = "123"
	var resp oauth2.Token
	resp.AccessToken = "bbbnn"
	h.TokenMap["123456"] = &resp
	h.HandleUpdateContent(w, r)
	testMode = false
}

func TestHandler_handleGetContent(t *testing.T) {
	//testMode = true
	var h Handler
	h.Templates = template.Must(template.ParseFiles("index.html"))
	h.TokenMap = make(map[string]*oauth2.Token)
	var s usession.Session
	h.Sess = s
	r, _ := http.NewRequest("GET", "/challenge?archived=on&fpath=rs/challenge/en_us?g=g&b=b", nil)
	w := httptest.NewRecorder()
	h.Sess.InitSessionStore(w, r)
	session, _ := h.Sess.GetSession(r)
	session.Values["accessTokenKey"] = "123456"
	//session.Values["userLoggenIn"] = true
	session.Values["clientId"] = "123"
	var resp oauth2.Token
	resp.AccessToken = "bbbnn"
	h.TokenMap["123456"] = &resp
	h.HandleGetContent(w, r)
	testMode = false
}

func TestHandler_handleGetContent2(t *testing.T) {
	//testMode = true
	var h Handler
	h.Templates = template.Must(template.ParseFiles("index.html"))
	h.TokenMap = make(map[string]*oauth2.Token)
	var s usession.Session
	h.Sess = s
	r, _ := http.NewRequest("GET", "/challenge?archived=on&fpath=rs/challenge/en_us?g=g&b=b", nil)
	w := httptest.NewRecorder()
	h.Sess.InitSessionStore(w, r)
	session, _ := h.Sess.GetSession(r)
	session.Values["accessTokenKey"] = "123456"
	session.Values["userLoggenIn"] = true
	session.Values["clientId"] = "123"
	var resp oauth2.Token
	resp.AccessToken = "bbbnn"
	h.TokenMap["123456"] = &resp
	h.HandleGetContent(w, r)
	testMode = false
}

func TestHandler_handleDeleteContent(t *testing.T) {
	//testMode = true
	var h Handler
	h.Templates = template.Must(template.ParseFiles("index.html"))
	h.TokenMap = make(map[string]*oauth2.Token)
	var s usession.Session
	h.Sess = s
	r, _ := http.NewRequest("GET", "/challenge?archived=on&fpath=rs/challenge/en_us?g=g&b=b", nil)
	w := httptest.NewRecorder()
	h.Sess.InitSessionStore(w, r)
	session, _ := h.Sess.GetSession(r)
	session.Values["accessTokenKey"] = "123456"
	//session.Values["userLoggenIn"] = true
	session.Values["clientId"] = "123"
	var resp oauth2.Token
	resp.AccessToken = "bbbnn"
	h.TokenMap["123456"] = &resp
	h.HandleDeleteContent(w, r)
	testMode = false
}

func TestHandler_handleDeleteContent2(t *testing.T) {
	//testMode = true
	var h Handler
	h.Templates = template.Must(template.ParseFiles("index.html"))
	h.TokenMap = make(map[string]*oauth2.Token)
	var s usession.Session
	h.Sess = s
	r, _ := http.NewRequest("GET", "/challenge?archived=on&fpath=rs/challenge/en_us?g=g&b=b", nil)
	w := httptest.NewRecorder()
	h.Sess.InitSessionStore(w, r)
	session, _ := h.Sess.GetSession(r)
	session.Values["accessTokenKey"] = "123456"
	session.Values["userLoggenIn"] = true
	session.Values["clientId"] = "123"
	var resp oauth2.Token
	resp.AccessToken = "bbbnn"
	h.TokenMap["123456"] = &resp
	h.HandleDeleteContent(w, r)
	testMode = false
}
