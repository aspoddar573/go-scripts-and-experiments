package tickledb

import (
	"encoding/json"
	"github.com/MindTickle/storageprotos/pb/tickleDbSqlStore"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/MindTickle/governance-utility/logger"
)

/*
	QueryTemplateFile corresponds to multiple templated queries written to fetch data from tickledb sql store
*/
type ViewName string

func GetRequestContext(orgId int64, env string, namespace string) *tickleDbSqlStore.RequestContext {
	return &tickleDbSqlStore.RequestContext{
		TenantId:  orgId,
		Namespace: namespace,
		Env:       env,
	}
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

func GetFilterField(fieldName string, fieldValue interface{}, filterPredicate tickleDbSqlStore.Filter_Predicate) *tickleDbSqlStore.Filter {
	var value *tickleDbSqlStore.Value = nil

	switch fieldValue.(type) {
	case []int64:
		value = &tickleDbSqlStore.Value{
			ValueType: &tickleDbSqlStore.Value_Int64ArrayValue{
				Int64ArrayValue: &tickleDbSqlStore.Int64Array{
					Int64Array: fieldValue.([]int64),
				},
			},
		}
	case []string:
		value = &tickleDbSqlStore.Value{
			ValueType: &tickleDbSqlStore.Value_StringArrayValue{
				StringArrayValue: &tickleDbSqlStore.StringArray{
					StringArray: fieldValue.([]string),
				},
			},
		}
	case bool:
		value = &tickleDbSqlStore.Value{
			ValueType: &tickleDbSqlStore.Value_BooleanValue{BooleanValue: fieldValue.(bool)},
		}
	case int32:
		value = &tickleDbSqlStore.Value{
			ValueType: &tickleDbSqlStore.Value_Int32Value{Int32Value: fieldValue.(int32)},
		}
	case int64:
		value = &tickleDbSqlStore.Value{
			ValueType: &tickleDbSqlStore.Value_Int64Value{Int64Value: fieldValue.(int64)},
		}
	case string:
		value = &tickleDbSqlStore.Value{
			ValueType: &tickleDbSqlStore.Value_StringValue{StringValue: fieldValue.(string)},
		}
	case float64:
		value = &tickleDbSqlStore.Value{
			ValueType: &tickleDbSqlStore.Value_DoubleValue{DoubleValue: fieldValue.(float64)},
		}
	case nil:
		value = &tickleDbSqlStore.Value{
			ValueType: &tickleDbSqlStore.Value_NullValue{NullValue: structpb.NullValue_NULL_VALUE},
		}
	}

	filterField := &tickleDbSqlStore.Filter{
		Predicate: filterPredicate,
		Field:     fieldName,
		Value:     value,
	}

	return filterField
}
