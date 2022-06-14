package web

import (
	"fmt"
	"mtui/app"
	"mtui/public"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/vearutop/statigz"
	"github.com/vearutop/statigz/brotli"
)

func Setup(a *app.App) error {
	r := mux.NewRouter()

	api := NewApi(a)
	r.HandleFunc("/api/ws", api.Websocket)
	r.HandleFunc("/api/login", api.DoLogout).Methods(http.MethodDelete)
	r.HandleFunc("/api/login", api.DoLogin).Methods(http.MethodPost)
	r.HandleFunc("/api/login", api.GetLogin).Methods(http.MethodGet)
	r.HandleFunc("/api/features", api.GetFeatures).Methods(http.MethodGet)
	r.HandleFunc("/api/areas", Secure(api.GetAreas)).Methods(http.MethodGet)
	r.HandleFunc("/api/mail/list", Secure(api.GetMails)).Methods(http.MethodGet)
	r.HandleFunc("/api/mail/{sender}/{time}", Secure(api.DeleteMail)).Methods(http.MethodDelete)
	r.HandleFunc("/api/mail/{sender}/{time}/read", Secure(api.MarkRead)).Methods(http.MethodPost)
	r.HandleFunc("/api/mail/checkrecipient/{recipient}", Secure(api.CheckValidRecipient)).Methods(http.MethodGet)
	r.HandleFunc("/api/mail/send/{recipient}", Secure(api.SendMail)).Methods(http.MethodPost)
	r.HandleFunc("/api/mail/contacts", Secure(api.GetContacts)).Methods(http.MethodGet)
	r.HandleFunc("/api/changepw", Secure(api.ChangePassword)).Methods(http.MethodPost)
	r.HandleFunc("/api/bridge/execute_chatcommand", Secure(api.ExecuteChatcommand)).Methods(http.MethodPost)
	r.HandleFunc("/api/bridge", CheckApiKey(os.Getenv("APIKEY"), a.Bridge.HandlePost)).Methods(http.MethodPost)
	r.HandleFunc("/api/bridge", CheckApiKey(os.Getenv("APIKEY"), a.Bridge.HandleGet)).Methods(http.MethodGet)

	// start tan login listener
	go api.TanSetListener(a.Bridge.AddHandler())

	// static files
	if os.Getenv("WEBDEV") == "true" {
		fmt.Println("using live mode")
		fs := http.FileServer(http.FS(os.DirFS("public")))
		r.PathPrefix("/").HandlerFunc(fs.ServeHTTP)

	} else {
		fmt.Println("using embed mode")
		r.PathPrefix("/").Handler(statigz.FileServer(public.Webapp, brotli.AddEncoding))
	}

	http.Handle("/", r)
	return nil
}
