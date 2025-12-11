package pb

import "fmt"

func ConstructQueryStringForAPI(options PocketBaseListOptions) (output string, hasOptions bool) {
	queryString := "?"

	formattedOptions := make([]string, 0)

	formattedOptions = append(formattedOptions, fmt.Sprintf("page=%d", options.Page))
	formattedOptions = append(formattedOptions, fmt.Sprintf("perPage=%d", options.PerPage))

	if len(options.Expand) > 0 {
		formattedOptions = append(formattedOptions, fmt.Sprintf("expand=%s", options.Expand))
	}

	if len(options.Fields) > 0 {
		formattedOptions = append(formattedOptions, fmt.Sprintf("fields=%s", options.Fields))
	}

	if len(options.Filter) > 0 {
		formattedOptions = append(formattedOptions, fmt.Sprintf("filter=(%s)", options.Filter))
	}

	if len(options.Sort) > 0 {
		formattedOptions = append(formattedOptions, fmt.Sprintf("sort=%s", options.Sort))
	}

	if len(formattedOptions) > 0 {
		for i, v := range formattedOptions {
			if i >= len(formattedOptions)-1 {
				queryString += v
				continue
			}

			queryString += v + "&"
		}
	}

	return queryString, len(formattedOptions) > 0
}
