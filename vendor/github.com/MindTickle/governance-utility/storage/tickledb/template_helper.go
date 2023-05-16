package tickledb

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"text/template"

	"github.com/MindTickle/governance-utility/logger"
)

var templateHelper *TemplateHelper
var once sync.Once

/**
Function will return a TemplateHelper for entity.
Templates must be stored in sql_template directory. SQL templates must use the file extension .t.sql

e.g.:
.
├── caller.go
├── sql_template
│   ├── list_data.t.sql
│   ├── save_data.t.sql
*/
func GetTemplateHelper(templatesPath string) *TemplateHelper {
	once.Do(func() {
		var err error
		templateHelper, err = NewTemplateHelper(templatesPath)
		if err != nil {
			panic(fmt.Sprintf("Failed to initialize template helper due to err %s", err))
		}
	})
	return templateHelper
}

func NewTemplateHelper(dir string) (*TemplateHelper, error) {
	viewTemplateMap := make(map[ViewName]string)
	if err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".sql" {
			viewName := strings.Replace(info.Name(), ".t.sql", "", 1)
			viewTemplate, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			viewTemplateMap[ViewName(viewName)] = string(viewTemplate)
		}
		return nil

	}); err != nil {
		return nil, err
	}
	return &TemplateHelper{viewTemplateMap: viewTemplateMap}, nil
}

type TemplateHelper struct {
	viewTemplateMap map[ViewName]string
}

func (t *TemplateHelper) GetTemplate(view ViewName) (string, error) {

	if val, ok := t.viewTemplateMap[view]; ok {
		return val, nil
	}
	return "", status.Error(codes.Internal, fmt.Sprintf("view template: %s not present", view))
}

var funcMap = template.FuncMap{
	"isStrPtrEmpty": func(str *string) bool {
		if str == nil || strings.TrimSpace(*str) == "" {
			return true
		}
		return false
	},
	"isIntPtrEmpty": func(ptr *int32) bool {
		if ptr == nil || *ptr == 0 {
			return true
		}
		return false
	},
	"isInt64PtrEmpty": func(ptr *int64) bool {
		if ptr == nil || *ptr == 0 {
			return true
		}
		return false
	},
	"buildInt64ForInClause": func(items []int64) string {
		value := ""
		for _, item := range items {
			value = value + strconv.FormatInt(item, 10) + ","
		}
		return value[:(len(value) - 1)]
	},
	"buildStringForInClause": func(items []string) string {
		return "'" + strings.Join(items, "', '") + "'"
	},
	"buildIntForInClause": func(items []string) string {
		return strings.Join(items, ", ")
	},
	"getInt32FromInt32Pointer": func(ptr *int32) int32 {
		return *ptr
	},
	"addInt32": func(nums ...int32) int32 {
		sum := int32(0)
		for _, num := range nums {
			sum = sum + num
		}
		return sum
	},
	"buildStructSliceForInClause": func(items reflect.Value) (string, error) {
		if items.Kind() != reflect.Slice {
			return "", fmt.Errorf("valid arg for buildStructSliceForInClause is []fmt.Stringer")
		}
		stringVals := make([]string, 0)
		for idx := 0; idx < items.Len(); idx++ {
			stringVals = append(stringVals, items.Index(idx).Interface().(fmt.Stringer).String())
		}
		return "'" + strings.Join(stringVals, "', '") + "'", nil
	},
}

func (t *TemplateHelper) BuildSqlQueryFromTemplate(ctx context.Context, viewName ViewName, filter interface{}) (*string, error) {
	sb := strings.Builder{}

	viewTemplate, err := t.GetTemplate(viewName)
	if err != nil {
		logger.Errorf(ctx, logger.NewFacets().AddField("view", viewName), "view template not present")
		return nil, status.Errorf(codes.Internal, "unable to build query for view %s", viewName)
	}

	sqlTemplate, err := template.New("Test").
		Funcs(funcMap).
		Parse(viewTemplate)
	if err != nil {
		logger.Errorf(ctx, logger.NewFacets().AddField("template", viewTemplate).AddField("error", err),
			"failed to initialize template")
		return nil, status.Errorf(codes.Internal, "unable to build query for view %s", viewName)
	}

	if err = sqlTemplate.Execute(&sb, filter); err != nil {
		logger.Errorf(ctx, logger.NewFacets().AddField("view", viewName).AddField("error", err),
			"failed to build query")
		return nil, status.Errorf(codes.Internal, "unable to build query for view %s", viewName)
	}

	query := strings.Replace(sb.String(), "\n", " ", -1)

	logger.Debugf(ctx, logger.NewFacets().AddField("query", query).AddField("view", viewName).AddField("template", viewTemplate),
		"successfully built query")

	return &query, nil
}
