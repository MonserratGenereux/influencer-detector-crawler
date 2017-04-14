package facebook

type facebookNode struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	Category string `json:"category,omitempty"`

	CategoryList []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"category_list"`

	About    string `json:"about"`
	FanCount int64  `json:"fan_count"`

	Location struct {
		City    string `json:"city"`
		Country string `json:"country"`
		ZIP     string `json:"zip"`
	} `json:"location"`
}

func (fn *facebookNode) isEmpty() bool {
	return len(fn.ID) == 0
}
