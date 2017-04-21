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

			log.Println("Requesting page", pageInfo.ID, "with depth", pageInfo.Depth)
			facebookCrawler := facebook.NewCrawler(pageInfo.ID, pageInfo.Depth)
			if err := facebookCrawler.Start(); err != nil {
				fmt.Fprint(w, err)
			} else {
				fmt.Fprint(w, "OK\n")
			}
		})

	return router
}

//  hola, no sabes programar :) ;
