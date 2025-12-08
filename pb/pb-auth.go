package pb

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type PBAuth struct {
	BaseURL string `json:"baseURL"`
}

func (auth *PBAuth) GetPBCollectionsAuthMethods(collection string, fields string) (AuthMethodResponse, error) {
	apiURL := fmt.Sprintf("%s/api/collections/%s/auth-methods/?fields=%s", auth.BaseURL, collection, fields)

	res, err := SendHTTPRequest("GET", apiURL, map[string]string{}, map[string]any{})

	if err != nil {
		return AuthMethodResponse{}, err
	}

	defer res.Body.Close()

	authMethodResponse := AuthMethodResponse{}
	err = json.NewDecoder(res.Body).Decode(&authMethodResponse)

	if err != nil {
		return AuthMethodResponse{}, err
	}

	return authMethodResponse, nil
}

func (auth *PBAuth) AuthWithPasswordForCollection(collection string, expand string, fields string, identity string, password string) (AuthSuccessResponse, error) {

	urlBase := fmt.Sprintf("%s/api/collections/%s/auth-with-password", auth.BaseURL, collection)

	urlBase += AddExpandAndFieldsToURL(expand, fields)

	body := map[string]any{
		"identity": identity,
		"password": password,
	}

	res, err := SendHTTPRequest("POST", urlBase, map[string]string{}, body)

	if err != nil {
		return AuthSuccessResponse{}, err
	}

	responseStatusCode := res.StatusCode

	defer res.Body.Close()

	switch responseStatusCode {
	case http.StatusOK:
		authSuccessResponse := AuthSuccessResponse{}

		json.NewDecoder(res.Body).Decode(&authSuccessResponse)

		return authSuccessResponse, nil
	case http.StatusBadRequest, http.StatusNotFound:
		errorResponse := PocketBaseErrorResponse{}

		json.NewDecoder(res.Body).Decode(&errorResponse)

		return AuthSuccessResponse{}, errors.New(errorResponse.Message)
	}

	return AuthSuccessResponse{}, errors.New("unknown-response")
}

func (auth *PBAuth) RefreshAuth(collection string, token string) (AuthSuccessResponse, error) {
	headers := map[string]string{
		"Authorization": fmt.Sprintf("%s", token),
	}

	apiURL := fmt.Sprintf("%s/api/collections/%s/auth-refresh", auth.BaseURL, collection)

	res, err := SendHTTPRequest("POST", apiURL, headers, map[string]any{})

	if err != nil {
		return AuthSuccessResponse{}, err
	}

	resStatus := res.StatusCode

	if resStatus == http.StatusOK {
		defer res.Body.Close()

		authRefreshSuccess := AuthSuccessResponse{}
		json.NewDecoder(res.Body).Decode(&authRefreshSuccess)

		return authRefreshSuccess, nil
	}

	authErrResponse := PocketBaseErrorResponse{}
	json.NewDecoder(res.Body).Decode(&authErrResponse)

	return AuthSuccessResponse{}, errors.New(authErrResponse.Message)
}
