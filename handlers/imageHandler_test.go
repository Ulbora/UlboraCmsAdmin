package handlers

import (
	"fmt"
	usession "github.com/Ulbora/go-better-sessions"
	oauth2 "github.com/Ulbora/go-oauth2-client"
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_handleAddImage(t *testing.T) {
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
	h.handleAddImage(w, r)
}

func TestHandler_handleAddImage2(t *testing.T) {
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
	var resp oauth2.Token
	resp.AccessToken = "bbbnn"
	h.TokenMap["123456"] = &resp
	h.handleAddImage(w, r)
}

func TestHandler_handleImagerUpload(t *testing.T) {
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
	h.handleImagerUpload(w, r)
}

func TestHandler_handleImagerUpload2(t *testing.T) {
	var h Handler
	testMode = true
	h.Templates = template.Must(template.ParseFiles("index.html"))
	h.TokenMap = make(map[string]*oauth2.Token)
	var s usession.Session
	h.Sess = s
	r := newTestMultipartRequest(t) //, _ := http.NewRequest("PUT", "http://challenge?image=challenge&fpath=rs/challenge/en_us?g=g&b=b", nil)
	w := httptest.NewRecorder()
	h.Sess.InitSessionStore(w, r)
	session, _ := h.Sess.GetSession(r)
	session.Values["accessTokenKey"] = "123456"
	session.Values["userLoggenIn"] = true
	session.Values["clientId"] = "123"
	var resp oauth2.Token
	resp.AccessToken = "bbbnn"
	h.TokenMap["123456"] = &resp
	h.handleImagerUpload(w, r)
}

func TestHandler_handleImagerUpload3(t *testing.T) {
	var h Handler
	testMode = false
	h.Templates = template.Must(template.ParseFiles("index.html"))
	h.TokenMap = make(map[string]*oauth2.Token)
	var s usession.Session
	h.Sess = s
	r := newTestMultipartRequest(t) //, _ := http.NewRequest("PUT", "http://challenge?image=challenge&fpath=rs/challenge/en_us?g=g&b=b", nil)
	w := httptest.NewRecorder()
	h.Sess.InitSessionStore(w, r)
	session, _ := h.Sess.GetSession(r)
	session.Values["accessTokenKey"] = "123456"
	session.Values["userLoggenIn"] = true
	session.Values["clientId"] = "123"
	var resp oauth2.Token
	resp.AccessToken = "bbbnn"
	h.TokenMap["123456"] = &resp
	h.handleImagerUpload(w, r)
}

func newTestMultipartRequest(t *testing.T) *http.Request {
	b := strings.NewReader(strings.Replace(message, "\n", "\r\n", -1))
	req, err := http.NewRequest("POST", "/", b)
	if err != nil {
		t.Fatal("NewRequest:", err)
	}
	ctype := fmt.Sprintf(`multipart/form-data; boundary="%s"`, boundary)
	req.Header.Set("Content-type", ctype)
	return req
}

const (
	fileaContents = "This is a test file."
	filebContents = "Another test file."
	textaValue    = "foo"
	textbValue    = "bar"
	boundary      = `MyBoundary`
)
const message = `
--MyBoundary
Content-Disposition: form-data; name="image"; filename="image.txt"
Content-Type: text/plain

` + fileaContents + `
--MyBoundary
Content-Disposition: form-data; name="fileb"; filename="fileb.txt"
Content-Type: text/plain

` + filebContents + `
--MyBoundary
Content-Disposition: form-data; name="texta"

` + textaValue + `
--MyBoundary
Content-Disposition: form-data; name="textb"

` + textbValue + `
--MyBoundary--
`
