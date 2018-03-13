package views

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zeenfaizpy/GoBlog/src/config"
	"github.com/zeenfaizpy/GoBlog/src/models"
)

// IndexHandler ...
func IndexHandler(env config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Print("GET IndexHandler")

		posts, err := models.GetAllPosts(env.DB)
		if err != nil {
			log.Print(err)
			http.Error(w, http.StatusText(500), 500)
		}

		for _, post := range posts {
			fmt.Fprintf(w, "%d %s %s \n", post.ID, post.Title, post.Content)
		}

	}
}
