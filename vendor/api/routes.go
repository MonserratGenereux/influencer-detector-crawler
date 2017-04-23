package api

import (
	"crawler/facebook"
	"fmt"
	"models"
	"net/http"

	"encoding/json"

	"log"

	"github.com/julienschmidt/httprouter"
)

var (
	router *httprouter.Router
)

// GetRoutes returns router with the defined API routes.
func GetRoutes() *httprouter.Router {
	if router == nil {
		router = buildRoutes()
	}

	return router
}

func buildRoutes() *httprouter.Router {
	router := httprouter.New()

	// Root
	router.GET("/", func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Wizeline Bots Analytics API up and running\n")
	})

	router.POST("/page",
		func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {

			var pageInfo models.PageRequest
			err := json.NewDecoder(req.Body).Decode(&pageInfo)
			if err != nil {
				fmt.Fprint(w, "Error parsing page request")
				return
			}

			msg := fmt.Sprintf("Received crawling request for page %s in depth %d\n",
				pageInfo.ID,
				pageInfo.Depth)

			log.Println(msg)
			fmt.Fprint(w, msg)

			facebookCrawler := facebook.NewCrawler(pageInfo.ID, pageInfo.Depth)

			log.Printf("Starting page %s in depth %d\n", pageInfo.ID, pageInfo.Depth)
			err = facebookCrawler.Start()
			if err != nil {
				log.Println(err)
			}

			log.Printf("Finished page %s in depth %d\n", pageInfo.ID, pageInfo.Depth)
		})

	return router
}

//  hola, no sabes programar :) ;
