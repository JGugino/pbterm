package pb

import "fmt"

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
