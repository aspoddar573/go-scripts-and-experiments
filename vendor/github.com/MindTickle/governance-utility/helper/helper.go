package helper

import (
	"context"
	"errors"
	"github.com/MindTickle/infracommon/constant/infraconstant"
	"net/url"
	"os"
	"strings"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func ExtractFromContext(ctx context.Context, field string) (string, error) {
	if ctx == nil {
		return "", errors.New("Context is empty")
	}

	value := ctx.Value(field)
	if value == nil {
		return "", errors.New("Value is nil")
	}
	return value.(string), nil
}

func ExtractInfraConstantFromContext(ctx context.Context, field infraconstant.InfraConstant) (string, error) {
	if ctx == nil {
		return "", errors.New("Context is empty")
	}

	value := ctx.Value(field)
	if value == nil {
		return "", errors.New("Value is nil")
	}
	return value.(string), nil
}

func IsEmpty(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

func Min(val1 int, val2 int) int {
	if val1 < val2 {
		return val1
	}
	return val2
}

func ExtractBaseUrlWithPath(ctx context.Context, urlString string) (string, error) {
	urlObject, err := url.Parse(urlString)
	if urlString == "" || err != nil {
		return "", err
	}
	return urlObject.Scheme + "://" + urlObject.Host + urlObject.Path, nil
}

func GetCommonIdsInt64(list1 []int64, list2 []int64) []int64 { // returns Ids in order of list2
	commonIds := make([]int64, 0)
	list1Map := make(map[int64]bool, 0)
	for _, value := range list1 {
		list1Map[value] = true
	}
	for _, value := range list2 {
		if _, exists := list1Map[value]; exists {
			commonIds = append(commonIds, value)
		}
	}
	return commonIds
}
