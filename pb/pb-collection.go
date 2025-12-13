package pb

import "fmt"

type CollectionType int

const (
	BaseCollection CollectionType = iota
	AuthCollection
	ViewCollection
)

func (c CollectionType) String() string {
	switch c {
	case 0:
		return "base"
	case 1:
		return "auth"
	case 2:
		return "view"
	default:
		return "base"
	}
}

type CollectionOptions struct {
	Name       string           `json:"name"`
	Type       CollectionType   `json:"type"`
	Fields     []map[string]any `json:"fields"`
	System     bool             `json:"system"`
	ListRule   string           `json:"listRule"`
	ViewRule   string           `json:"viewRule"`
	CreateRule string           `json:"createRule"`
	UpdateRule string           `json:"updateRule"`
	DeleteRule string           `json:"deleteRule"`
	ViewQuery  string           `json:"viewQuery"`
}

type PBCollection struct {
	BaseURL string
}

// ### CREATE COLLECTION ###
func (collection *PBCollection) CreateNewCollection(token string) {
	apiUrl := fmt.Sprintf("%s/api/collections", collection.BaseURL)

	res, err := SendAuthenticatedHTTPRequest("POST", apiUrl, map[string]string{}, map[string]any{}, token)

	if err != nil {
		return
	}

	fmt.Println(res)
}

// ### UPDATE COLLECTION ###
func (collection *PBCollection) UpdateCollection(token string, desiredCollection string) {
	apiUrl := fmt.Sprintf("%s/api/collections/%s", collection.BaseURL, desiredCollection)
	res, err := SendAuthenticatedHTTPRequest("PATCH", apiUrl, map[string]string{}, map[string]any{}, token)

	if err != nil {
		return
	}

	fmt.Println(res)

}

// ### VIEW COLLECTION ###

func (collection *PBCollection) ScaffoldCollections(token string, desiredCollection string) {
	apiUrl := fmt.Sprintf("%s/api/collections/meta/scaffolds", collection.BaseURL)
	res, err := SendAuthenticatedHTTPRequest("GET", apiUrl, map[string]string{}, map[string]any{}, token)

	if err != nil {
		return
	}

	fmt.Println(res)

}

func (collection *PBCollection) ViewCollection(token string, desiredCollection string) {
	apiUrl := fmt.Sprintf("%s/api/collections/%s", collection.BaseURL, desiredCollection)
	res, err := SendAuthenticatedHTTPRequest("GET", apiUrl, map[string]string{}, map[string]any{}, token)

	if err != nil {
		return
	}

	fmt.Println(res)

}

func (collection *PBCollection) ListCollections(token string) {
	apiUrl := fmt.Sprintf("%s/api/collections", collection.BaseURL)
	res, err := SendAuthenticatedHTTPRequest("GET", apiUrl, map[string]string{}, map[string]any{}, token)

	if err != nil {
		return
	}

	fmt.Println(res)

}

// ### DELETE COLLECTION ###
func (collection *PBCollection) DeleteCollection(token string, desiredCollection string) {
	apiUrl := fmt.Sprintf("%s/api/collections/%s", collection.BaseURL, desiredCollection)
	res, err := SendAuthenticatedHTTPRequest("DELETE", apiUrl, map[string]string{}, map[string]any{}, token)

	if err != nil {
		return
	}

	fmt.Println(res)
}

func (collection *PBCollection) TruncateCollection(token string, desiredCollection string) {
	apiUrl := fmt.Sprintf("%s/api/collections/%s/truncate", collection.BaseURL, desiredCollection)
	res, err := SendAuthenticatedHTTPRequest("DELETE", apiUrl, map[string]string{}, map[string]any{}, token)

	if err != nil {
		return
	}

	fmt.Println(res)
}
