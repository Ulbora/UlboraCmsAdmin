package main

import (
	hand "UlboraCmsAdmin/handlers"
	"fmt"
	usession "github.com/Ulbora/go-better-sessions"
	oauth2 "github.com/Ulbora/go-oauth2-client"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"os"
)

var h hand.Handler

var s usession.Session

var tokenMap map[string]*oauth2.Token
var credentialToken *oauth2.Token

func main() {
	s.MaxAge = sessingTimeToLive
	s.Name = userSession
	if os.Getenv("SESSION_SECRET_KEY") != "" {
		s.SessionKey = os.Getenv("SESSION_SECRET_KEY")
	} else {
		s.SessionKey = "115722gggg14ddfg4567"
	}
	h.Sess = s
	h.TokenMap = make(map[string]*oauth2.Token)

	h.Templates = template.Must(template.ParseFiles("./static/index.html", "./static/login.html", "./static/header.html",
		"./static/navbarLogin.html", "./static/footer.html", "./static/navbar.html", "./static/addContent.html",
		"./static/updateContent.html", "./static/imageUpload.html", "./static/images.html"))
	router := mux.NewRouter()

	router.HandleFunc("/", h.HandleAdminIndex).Methods("GET")
	router.HandleFunc("/loginUser", h.HandleImplicitLogin).Methods("POST")
	router.HandleFunc("/tokenImplicitHandler", h.HandleImplicitToken).Methods("GET")
	router.HandleFunc("/logout", h.HandleLogout).Methods("GET")

	router.HandleFunc("/addContent", h.HandleAddContent).Methods("GET")
	router.HandleFunc("/newContent", h.HandleNewContent).Methods("POST")
	router.HandleFunc("/getContent", h.HandleGetContent).Methods("GET")
	router.HandleFunc("/updateContent", h.HandleUpdateContent).Methods("POST")
	router.HandleFunc("/deleteContent", h.HandleDeleteContent).Methods("GET")

	router.HandleFunc("/addImage", h.HandleAddImage).Methods("GET")
	router.HandleFunc("/uploadImage", h.HandleImagerUpload).Methods("POST")
	router.HandleFunc("/images", h.HandleImages).Methods("GET")
	router.HandleFunc("/deleteImage", h.HandleDeleteImage).Methods("GET")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	fmt.Println("Ulbora CMS Admin!")
	fmt.Println("Listening on :8060...")
	http.ListenAndServe(":8060", router)

}
