package setTemplatesV2Flag

import (
	"bytes"
	"context"
	"encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"sync"
	"time"
)

type TemplatesV2FlagObject struct {
	TemplatesV2Flag bool `json:"templatesV2Enabled"`
}

type CeClientImpl struct {
	url     string
	timeout time.Duration
}

var onceCe sync.Once
var ceRestClient CeClientImpl

func GetContentEngineUrl(track string) string {
	return "http://cen-svc-ce.internal." + track + ".mindtickle.com"
}

func GetContentEngineClient(track string) CeClientImpl {
	onceCe.Do(func() {
		ceRestClient = CeClientImpl{
			url:     GetContentEngineUrl(track),
			timeout: 10 * time.Second,
		}
	})
	return ceRestClient
}

func SetTemplatesV2FlagForCompany(client CeClientImpl, companyId string, flag bool) error {

	templatesV2FlagObject := TemplatesV2FlagObject{
		TemplatesV2Flag: flag,
	}
	reqBodyBytes, err := json.Marshal(templatesV2FlagObject)

	ctx := context.Background()

	postReq, err := http.NewRequestWithContext(ctx, http.MethodPost, client.url+"/company/default_settings/update?company="+companyId, bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		return err
	}
	postReq.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(postReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = status.Errorf(codes.Internal, resp.Status)
		return err
	}
	return nil
}
