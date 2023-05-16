package sqlClient

import (
	"github.com/MindTickle/storageprotos/pb/tickleDbSqlStore"
	"google.golang.org/grpc"
)

const PROD_TICKLE_DB_HP = "tickledbsqlstore.internal-grpc.prod.mindtickle.com:80"
const STAGING_TICKLE_DB_HP = "tickledbsqlstore.internal-grpc.staging.mindtickle.com:80"
const INTEGRATION_TICKLE_DB_HP = "tickledbsqlstore.internal-grpc.integration.mindtickle.com:80"

func GetTickleDBClient(env string) tickleDbSqlStore.SqlStoreClient {
	storemanagerURL := "tickledbsqlstore.internal-grpc." + env + ".mindtickle.com:80"
	tickledbServer, _ := grpc.Dial(storemanagerURL, grpc.WithInsecure())
	tickledbClient := tickleDbSqlStore.NewSqlStoreClient(tickledbServer)
	return tickledbClient
}

func GetProdTickleDBClient() tickleDbSqlStore.SqlStoreClient {
	storemanagerURL := PROD_TICKLE_DB_HP
	tickledbServer, _ := grpc.Dial(storemanagerURL, grpc.WithInsecure())
	tickledbClient := tickleDbSqlStore.NewSqlStoreClient(tickledbServer)
	return tickledbClient
}

func GetIntegrationTickleDBClient() tickleDbSqlStore.SqlStoreClient {
	storemanagerURL := INTEGRATION_TICKLE_DB_HP
	tickledbServer, _ := grpc.Dial(storemanagerURL, grpc.WithInsecure())
	tickledbClient := tickleDbSqlStore.NewSqlStoreClient(tickledbServer)
	return tickledbClient
}

func GetStagingTickleDBClient() tickleDbSqlStore.SqlStoreClient {
	storemanagerURL := STAGING_TICKLE_DB_HP
	tickledbServer, _ := grpc.Dial(storemanagerURL, grpc.WithInsecure())
	tickledbClient := tickleDbSqlStore.NewSqlStoreClient(tickledbServer)
	return tickledbClient
}
