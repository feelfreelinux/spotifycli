package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/feelfreelinux/spotifycli/spotifycli/gui"
	"github.com/jroimartin/gocui"
	"github.com/shibukawa/configdir"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

type config struct {
	Token        string `json:"clientId"`
	Secret       string `json:"secretKey"`
	RefreshToken string `json:"refreshToken"`
	TokenType    string `json:"tokenType"`
}

func loadConfig() (cfg config, err error) {
	dir := configdir.New("feelfreelinux", "spotifycli")
	cfgFile := dir.QueryFolderContainsFile("config.json")

	if cfgFile == nil { // file not found, ask user for data
		// TODO: add nice gui
		fmt.Println("--- Copy spotify auth token from app and paste ---")
		fmt.Print("ClientId: ")
		fmt.Scanln(&cfg.Token)
		fmt.Print("SecretId: ")
		fmt.Scanln(&cfg.Secret)
		cfg.TokenType = ""
		cfg.RefreshToken = ""
		var jsonFile *os.File
		jsonFile, err = dir.QueryFolders(configdir.Global)[0].Create("config.json")
		if err != nil {
			return
		}
		encoder := json.NewEncoder(jsonFile)
		encoder.SetIndent("", "    ")
		err = encoder.Encode(&cfg)
		if err != nil {
			return
		}
	} else {
		var jsonData []byte
		jsonData, err = cfgFile.ReadFile("config.json")
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(jsonData, &cfg)
		if err != nil {
			return
		}
	}
	return
}

const redirectURI = "http://localhost:8080/callback"

var (
	auth  = spotify.NewAuthenticator(redirectURI, spotify.ScopeUserReadCurrentlyPlaying, spotify.ScopeUserReadPlaybackState, spotify.ScopeUserModifyPlaybackState)
	ch    = make(chan *spotify.Client)
	state = "abc123"
)

func main() {
	cfg, cfgErr := loadConfig()
	if cfgErr != nil {
		log.Panicf("Failed to load config %v", cfgErr)
	}

	auth.SetAuthInfo(cfg.Token, cfg.Secret)

	if cfg.RefreshToken == "" {
		// We'll want these variables sooner rather than later
		var client *spotify.Client

		http.HandleFunc("/callback", completeAuth)

		go func() {
			url := auth.AuthURL(state)
			fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

			// wait for auth to complete
			client = <-ch

			openGui(client)
		}()

		http.ListenAndServe(":8080", nil)
	} else {
		token := &oauth2.Token{
			TokenType:    cfg.TokenType,
			RefreshToken: cfg.RefreshToken,
		}
		client := auth.NewClient(token)
		openGui(&client)
	}

}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.Token(state, r)

	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}

	cfg, cfgErr := loadConfig()
	if cfgErr != nil {
		return
	}

	cfg.RefreshToken = tok.RefreshToken
	cfg.TokenType = tok.TokenType

	var jsonFile *os.File

	dir := configdir.New("feelfreelinux", "spotifycli")
	jsonFile, err = dir.QueryFolders(configdir.Global)[0].Create("config.json")
	if err != nil {
		return
	}
	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(&cfg)
	if err != nil {
		return
	}

	client := auth.NewClient(tok)
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "Login Completed!")
	ch <- &client
}

func openGui(client *spotify.Client) {
	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	err = gui.CreateMainView(g, client)
	if err != nil {
		panic(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
