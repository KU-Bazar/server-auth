package Controller

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"

	"go-auth/Config"
)

func GoogleLoginController(w http.ResponseWriter, r *http.Request) {

	ran := make([]byte, 32)
	_, err := rand.Read(ran)
	if err != nil {
		log.Fatal(err)
	}

	state := base64.URLEncoding.EncodeToString(ran)

	temp := Config.GoogleConfig()
	url := temp.AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func GoogleCallback(w http.ResponseWriter, r *http.Request) {

	var googleOAuthConfig = Config.GoogleConfig()

	token, err := googleOAuthConfig.Exchange(context.Background(), r.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	defer response.Body.Close()
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	data := contents

	fmt.Fprintf(w, "UserInfo: %s\n", data)

}
