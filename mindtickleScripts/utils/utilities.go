package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MindTickle/governance-utility/logger"
	"github.com/MindTickle/storageprotos/pb/tickleDbSqlStore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func UnmarshalString(str string, obj interface{}) error {
	err := json.Unmarshal([]byte(str), obj)
	return err
}

func GetOrgIdForEnv(env string) int64 {
	if env == "integration" {
		return 1212991592195620381
	} else if env == "staging" {
		return 1212991592195620381
	} else if env == "prod" {
		return 1212991592195620381
	} else if env == "prod-us" {
		return 1212991592195620381
	}
	return 0
}

func GetCompanyIdForEnv(env string) int64 {
	if env == "integration" {
		return 1212991592654170730
	} else if env == "staging" {
		return 1212991592654170730
	} else if env == "prod" {
		return 1212991592654170730
	} else if env == "prod-us" {
		return 1212991592654170730
	}
	return 0
}

func ConvertRowToDbModel(ctx context.Context, row *tickleDbSqlStore.SqlRow) ([]byte, error) {
	if row == nil {
		logger.Errorf(ctx, logger.NewFacets(), "tickledb row is nil")
		return nil, status.Errorf(codes.Internal, "tickledb row is nil")
	}
	data := row.Data
	rowMap := make(map[string]string)
	for key, value := range data {
		rowMap[key] = string(value)
	}
	jsonString, err := json.Marshal(rowMap)
	if err != nil {
		logger.Errorf(ctx, logger.NewFacets(), "error in marshalling row : %+v, error: %+v", rowMap, err)
		return nil, status.Errorf(codes.Internal, "error in marshalling tickledb row: %+v", rowMap)
	}
	return jsonString, nil
}

func GetRequestContext(orgId int64, env string, namespace string) *tickleDbSqlStore.RequestContext {
	return &tickleDbSqlStore.RequestContext{
		TenantId:  orgId,
		Namespace: namespace,
		Env:       env,
	}
}

func CreateSearchQuery(tableName string, templateList []string) string {
	query := "select * from " + tableName + " where id in "
	inClause := "("
	for idx, item := range templateList {
		inClause += item
		if idx != len(templateList)-1 {
			inClause += ", "
		}
	}
	inClause += ")"
	query += inClause
	fmt.Printf("Query is %s\n", query)
	return query
}
