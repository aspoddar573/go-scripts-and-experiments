package helper

import (
	"context"
	"errors"
	"github.com/MindTickle/infracommon/constant/infraconstant"
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
