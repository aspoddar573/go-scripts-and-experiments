package urlParser

import (
	"fmt"
	"net/url"
)

func ExtractBaseUrlWithPath(urlString string) (string, error) {
	urlObject, err := url.Parse(urlString)
	if err != nil {
		fmt.Printf("Error occurred while parsing url: %+v!\n", err)
		return "", err
	}
	urlWithPath := urlObject.Scheme + "://" + urlObject.Host + urlObject.Path
	fmt.Printf("Extracted url is: %s", urlWithPath)
	return urlWithPath, nil
}
