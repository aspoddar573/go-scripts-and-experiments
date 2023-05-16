package fetchAllActiveCompanies

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MindTickle/governance-utility/storage/tickledb"
	"github.com/MindTickle/storageprotos/pb/tickleDbSqlStore"
	"goScriptsAndExperiments/mindtickleScripts/templates/programTemplates-v1.0/sqlClient"
	"goScriptsAndExperiments/mindtickleScripts/templates/utils"
)

const (
	NAMESPACE = "platform"
)

type CompanyObject struct {
	CompanyId int64 `json:"company_id,string"`
}

func FetchAllCompaniesForTrack(track string) ([]CompanyObject, error) {
	var allCompaniesForTrack []CompanyObject

	ctx := context.Background()
	targetSqlStoreClient := sqlClient.GetTickleDBClient(track)

	searchRes, err := targetSqlStoreClient.Search(ctx, &tickleDbSqlStore.SearchRequest{
		RequestContext: utils.GetRequestContext(1, track, NAMESPACE),
		SqlStatement:   "SELECT ci.company_id as company_id FROM company_info ci where (ci.company_type != 'DELETED')",
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for _, row := range searchRes.Rows {
		rowBytes, err := tickledb.ConvertRowToDbModel(ctx, row)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		rowData := &CompanyObject{}
		err = json.Unmarshal(rowBytes, rowData)
		if err != nil {
			fmt.Println(rowBytes)
			fmt.Println(err)
			return nil, err
		}

		allCompaniesForTrack = append(allCompaniesForTrack, *rowData)
	}
	return allCompaniesForTrack, nil

}
