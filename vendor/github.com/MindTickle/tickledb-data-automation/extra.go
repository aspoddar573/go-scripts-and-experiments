package extra

import (
	"context"
	"fmt"

	"github.com/MindTickle/storageprotos/pb/tickleDb"
	"google.golang.org/grpc"
)

func CreateCollection(host string, collection *tickleDb.Collection) (*tickleDb.CreateCollectionResponse, error) {
	client, err := GetClient(host)
	if err != nil {
		return nil, err
	}
	resp, err := client.CreateCollection(context.Background(), &tickleDb.CreateCollectionRequest{
		Collection: collection,
	})
	if err != nil {
		return resp, err
	}
	fmt.Print("Successfully created collection")
	return resp, nil
}

func CreateNamespace(track string, namespace string, host string, nodeIP string) (resp *tickleDb.InitNamespaceOnTrackResponse, err error) {
	client, err := GetClient(host)
	if err != nil {
		return nil, err
	}

	if nodeIP != "" {
		resp, err = client.InitNamespaceOnTrack(context.Background(), &tickleDb.InitNamespaceOnTrackRequest{
			Env:       track,
			Namespace: namespace,
			NodeInfo:  &tickleDb.Node{Host: nodeIP},
		})
	} else {
		resp, err = client.InitNamespaceOnTrack(context.Background(), &tickleDb.InitNamespaceOnTrackRequest{
			Env:       track,
			Namespace: namespace,
		})
	}

	if err != nil {
		return resp, err
	}
	fmt.Print("Successfully created namespace")
	return nil, nil
}

func CreateTable(host string, table *tickleDb.Table) (*tickleDb.CreateTableResponse, error) {
	client, err := GetClient(host)
	if err != nil {
		return nil, err
	}
	resp, err := client.CreateTable(context.Background(), &tickleDb.CreateTableRequest{
		Table: table,
	})
	if err != nil {
		return resp, err
	}
	fmt.Print("Successfully created table")
	return resp, nil
}

func UpdateTable(host string, table *tickleDb.UpdateTableRequest) (*tickleDb.UpdateTableResponse, error) {
	client, err := GetClient(host)
	if err != nil {
		return nil, err
	}
	resp, err := client.UpdateTable(context.Background(), table)
	if err != nil {
		return resp, err
	}
	fmt.Print("Successfully Updated table")
	return resp, nil
}

func DeleteTable(host string, table *tickleDb.DeleteTableRequest) (*tickleDb.DeleteTableResponse, error) {
	client, err := GetClient(host)
	if err != nil {
		return nil, err
	}
	resp, err := client.DeleteTable(context.Background(), table)
	if err != nil {
		return resp, err
	}
	fmt.Print("Successfully Deleted table")
	return resp, nil
}

func GetClient(host string) (tickleDb.StoreManagerClient, error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("error connecting TickleDB sql store :%s", err)
		return nil, err
	}
	client := tickleDb.NewStoreManagerClient(conn)
	return client, nil
}
