package facebook

import (
	"fmt"
	"models"
	"strconv"
)

type facebookNode struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	About string `json:"about"`

	Category string `json:"category"`

	CategoryList []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"category_list"`

	FanCount int64 `json:"fan_count"`

	Location struct {
		City    string `json:"city"`
		Country string `json:"country"`
		ZIP     string `json:"zip"`
	} `json:"location"`
}

func (fbNode *facebookNode) isEmpty() bool {
	return len(fbNode.ID) == 0
}

func (fbNode *facebookNode) ToCrawlerNode() *models.Node {

	// Parse id
	var id int64
	id, _ = strconv.ParseInt(fbNode.ID, 10, 64)

	// Join main category and category list.
	categories := []string{fbNode.Category}
	for _, subCategories := range fbNode.CategoryList {
		categories = append(categories, subCategories.Name)
	}

	return &models.Node{
		ID:          id,
		Platform:    "facebook",
		Name:        fbNode.Name,
		Description: fbNode.About,
		Categories:  categories,
		FanCount:    fbNode.FanCount,
		City:        fbNode.Location.City,
		Country:     fbNode.Location.Country,
		ZIP:         string(fbNode.Location.ZIP),
	}
}

type facebookEdges struct {
	Edges  []facebookNode `json:"data"`
	Paging struct {
		Previous string `json:"previous"`
		Next     string `json:"next"`
	} `json:"paging"`
}

type graphAPIError struct {
	InnerError struct {
		Message        string `json:"message"`
		Type           string `json:"type"`
		Code           int64  `json:"code"`
		ErrorSubcode   int64  `json:"error_subcode"`
		ErrorUserTitle string `json:"error_user_title"`
		ErrorUserMsg   string `json:"error_user_msg"`
		FBTraceID      string `json:"fbtrace_id"`
	} `json:"error"`
}

func (e *graphAPIError) isEmpty() bool {
	return len(e.InnerError.Message) == 0
}

func (e graphAPIError) Error() string {
	return fmt.Sprintf("Graph API error %d - %s: %s", e.InnerError.Code, e.InnerError.Type, e.InnerError.Message)
}
