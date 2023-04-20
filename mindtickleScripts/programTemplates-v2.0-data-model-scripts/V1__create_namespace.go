package main

import (
	"fmt"
	"os"

	helper "github.com/MindTickle/tickledb-data-automation"
)

func main() {
	resp, err := helper.CreateNamespace(
		os.Getenv("conf_track"),
		os.Getenv("conf_database_name"),
		(os.Getenv("conf_host") + ":" + os.Getenv("conf_port")),
		os.Getenv("node_ip"),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(resp)
}
