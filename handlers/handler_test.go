package handlers

import (
	"fmt"
	usession "github.com/Ulbora/go-better-sessions"
	oauth2 "github.com/Ulbora/go-oauth2-client"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHandler_getToken(t *testing.T) {
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
	//session.Save(w, r)
	token := h.getToken(w, r)
	fmt.Print("Token: ")
	fmt.Println(token)
	if token.AccessToken != "bbbnn" {
		t.Fail()
	}
}

func TestHandler_getToken2(t *testing.T) {
	var h Handler
	h.TokenMap = make(map[string]*oauth2.Token)
	var s usession.Session
	h.Sess = s
	r, _ := http.NewRequest("GET", "/challenge?route=challenge&fpath=rs/challenge/en_us?g=g&b=b", nil)
	w := httptest.NewRecorder()
	h.Sess.InitSessionStore(w, r)
	//session, _ := h.Sess.GetSession(r)
	//session.Values["accessTokenKey"] = "123456"
	var resp oauth2.Token
	resp.AccessToken = "bbbnn"
	h.TokenMap["123456"] = &resp
	//session.Save(w, r)
	token := h.getToken(w, r)
	fmt.Print("Token: ")
	fmt.Println(token)
	if token != nil {
		t.Fail()
	}
}

func TestHandler_removeToken(t *testing.T) {
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
	tokena := h.TokenMap["123456"]
	fmt.Print("Token before delete: ")
	fmt.Println(tokena)
	h.removeToken(w, r)
	token := h.TokenMap["123456"]
	// session2, _ := h.Sess.GetSession(r)
	// token := session2.Values["accessTokenKey"]
	fmt.Print("Token after delete: ")
	fmt.Println(token)
	if tokena.AccessToken != "bbbnn" || token != nil {
		t.Fail()
	}
}

func TestHandler_getOauthHost(t *testing.T) {
	host := getOauthHost()
	if host != "http://localhost:3000" {
		t.Fail()
	}
}

func TestHandler_getOauthHost2(t *testing.T) {
	os.Setenv("AUTH_HOST", "12345")
	host := getOauthHost()
	if host != "12345" {
		t.Fail()
	}
	os.Setenv("AUTH_HOST", "")
}

func TestHandler_getOauthRedirectHost(t *testing.T) {
	host := getOauthRedirectHost()
	if host != "http://localhost:3000" {
		t.Fail()
	}
}

func TestHandler_getOauthRedirectHost2(t *testing.T) {
	os.Setenv("AUTH_REDIRECT_HOST", "12345")
	host := getOauthRedirectHost()
	if host != "12345" {
		t.Fail()
	}
	os.Setenv("AUTH_REDIRECT_HOST", "")
}

func TestHandler_getRedirectURI(t *testing.T) {
	r, _ := http.NewRequest("GET", "http://challenge", nil)
	host := getRedirectURI(r, "test")
	fmt.Print("host: ")
	fmt.Println(host)
	if host != "http://challenge" {
		t.Fail()
	}
}

func TestHandler_getRedirectURI2(t *testing.T) {
	r, _ := http.NewRequest("GET", "", nil)
	host := getRedirectURI(r, "test")
	fmt.Print("host bad schema: ")
	fmt.Println(host)
	if host != "http://test" {
		t.Fail()
	}
}

func TestHandler_getSession(t *testing.T) {
	var h Handler
	h.TokenMap = make(map[string]*oauth2.Token)
	var s usession.Session
	h.Sess = s
	r, _ := http.NewRequest("GET", "/challenge?route=challenge&fpath=rs/challenge/en_us?g=g&b=b", nil)
	w := httptest.NewRecorder()
	h.Sess.InitSessionStore(w, r)
	session := h.getSession(w, r)
	if session == nil {
		t.Fail()
	}
}

func TestHandler_getSession2(t *testing.T) {
	var h Handler
	h.TokenMap = make(map[string]*oauth2.Token)
	var s usession.Session
	h.Sess = s
	//r, _ := http.NewRequest("GET", "/challenge?route=challenge&fpath=rs/challenge/en_us?g=g&b=b", nil)
	r, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	session := h.getSession(w, r)
	if session != nil {
		t.Fail()
	}
}
