package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Health struct {
	ClusterName             string `json:"cluster_name"`
	Status                  string `json:"status"`
	TimedOut                bool   `json:"timedout"`
	NumberOfNodes           int    `json:"number_of_nodes"`
	NumberOfDataNodes       int    `json:"number_of_data_nodes"`
	ActivePrimaryShards     int    `json:"active_primary_shards"`
	ActiveShards            int    `json:"active_shards"`
	RelocatingShards        int    `json:"relocating_shards"`
	InitializingShards      int    `json:"initializing_shards"`
	UnassignedShards        int    `json:"unassigned_shards"`
	DelayedUnassignedShards int    `json:"delayed_unassigned_shards"`
	NumberOfPendingTasks    int    `json:"number_of_pending_tasks"`
	NumberOfInFlightFetch   int    `json:"number_of_in_flight_fetch"`
}

func main() {
	esAddr := os.Getenv("ES_ADDR")
	if esAddr == "" {
		if restAddr := os.Getenv("NOMAD_ADDR_rest"); restAddr != "" {
			esAddr = fmt.Sprintf("http://%s", os.Getenv("NOMAD_ADDR_rest"))
		}
	}
	if esAddr == "" {
		esAddr = "http://127.0.0.1:9200"
	}

	health := getLocalHealth(esAddr)

	fmt.Println("=================================")
	fmt.Printf("Cluster is %s\n", health.Status)
	fmt.Println("=================================")
	fmt.Printf("number_of_nodes: %d\n", health.NumberOfNodes)
	fmt.Printf("number_of_data_nodes: %d\n", health.NumberOfDataNodes)
	fmt.Printf("active_primary_shards: %d\n", health.ActivePrimaryShards)
	fmt.Printf("active_shards: %d\n", health.ActiveShards)
	fmt.Printf("relocating_shards: %d\n", health.RelocatingShards)
	fmt.Printf("initializing_shards: %d\n", health.InitializingShards)
	fmt.Printf("unassigned_shards: %d\n", health.UnassignedShards)
	fmt.Printf("delayed_unassigned_shards: %d\n", health.DelayedUnassignedShards)
	fmt.Printf("number_of_pending_tasks: %d\n", health.NumberOfPendingTasks)
	fmt.Printf("number_of_in_flight_fetch: %d\n", health.NumberOfInFlightFetch)

	if health.Status != "green" {
		os.Exit(1)
	}
}

func getLocalHealth(addr string) Health {
	url := fmt.Sprintf("%s/_cluster/health?local=true", addr)
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var data Health
	json.Unmarshal(body, &data)

	return data
}
