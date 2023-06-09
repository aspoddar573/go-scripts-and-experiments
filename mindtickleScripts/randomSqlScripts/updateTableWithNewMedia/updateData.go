package updateTableWithNewMedia

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MindTickle/storageprotos/pb/tickleDbSqlStore"
	"goScriptsAndExperiments/mindtickleScripts/templates/programTemplates-experiments/sqlClient"
	"goScriptsAndExperiments/mindtickleScripts/templates/utils"
	"math/rand"
)

type ProgramTemplateCreator struct {
	CreatorId int64 `json:"creator_id,omitempty,int64"`
}

type TemplateDbRow struct {
	Id               string `json:"id,omitempty"`
	ThumbnailMediaId string `json:"thumbnail_media_id,omitempty"`
}

type Template struct {
	Metadata TemplateMetadata `json:"template_metadata"`
}

type TemplateMetadata struct {
	ThumbnailImageId string `json:"thumbnail_image_id,omitempty"`
	ThumbnailVideoId string `json:"thumbnail_video_id,omitempty"`
}

type TemplateCategory struct {
	DisplayOrder int64 `json:"display_order,omitempty,int"`
}

func GetOrgId(val string) int64 {
	if val == "integration" {
		return 1583043504811348643
	} else if val == "staging" {
		return 1646473740207779828
	} else if val == "prod" {
		return 1526086301659646176
	} else if val == "prod-us" {
		return 1
	} else {
		return 1
	}
}

func GetVideoMediaIds(val string) []string {
	if val == "integration" {
		return []string{"", "2304190627499909126", "1648574157447366648", "1648574459184001353", "1648574569667619405"}
	} else if val == "staging" {
		return []string{"", "230419132747302177", "2304191328245575041", "230419132939097139", "230419132637275946"}
	} else if val == "prod" {
		return []string{""}
	} else if val == "prod-us" {
		return []string{""}
	} else {
		return []string{""}
	}
}

func UpdateThumbnailData() {
	targetEnv := "staging"

	videoMediaIds := GetVideoMediaIds(targetEnv)
	n := len(videoMediaIds)

	ctx := context.Background()
	targetSqlStoreClient := sqlClient.GetTickleDBClient(targetEnv)

	searchRes, err := targetSqlStoreClient.Search(ctx, &tickleDbSqlStore.SearchRequest{
		RequestContext: utils.GetRequestContext(1, targetEnv, "templates"),
		SqlStatement:   "select * from template",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(searchRes)
	rows := searchRes.Rows
	fmt.Printf("%d rows fetched!\n", len(rows))

	for _, row := range rows {
		rowDataBytes, _ := utils.ConvertRowToDbModel(ctx, row)
		templateDbRow := &TemplateDbRow{}
		err = json.Unmarshal(rowDataBytes, templateDbRow)
		if err != nil {
			fmt.Printf("Error occurred while unmarshalling program templates: %+v\n", err)
		}

		idx := rand.Int() % n
		template := Template{
			Metadata: TemplateMetadata{
				ThumbnailImageId: templateDbRow.ThumbnailMediaId,
				ThumbnailVideoId: videoMediaIds[idx],
			},
		}
		rowValue, err := json.Marshal(template)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(template)
		updateRes, err := targetSqlStoreClient.UpdateRowById(ctx, &tickleDbSqlStore.UpdateRowByIdRequest{
			RequestContext: utils.GetRequestContext(GetOrgId(targetEnv), targetEnv, "templates"),
			Project:        "templates",
			TableName:      "template",
			Row: &tickleDbSqlStore.Row{
				Id: templateDbRow.Id,
				RowValue: &tickleDbSqlStore.RowValue{
					RowInBytes: rowValue,
					AuthMeta: &tickleDbSqlStore.AuthMeta{
						GlobalContextId: "aspoddar-local",
						AuthId:          "_mt_admin",
					},
				},
			},
			Mask: &tickleDbSqlStore.Mask{Fields: []string{"template_metadata"}},
		})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(updateRes)

	}
}
