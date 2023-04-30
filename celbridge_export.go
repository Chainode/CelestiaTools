package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	localHeight = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "bridge_local_height",
		Help: "Local height of the Celestia node",
	})

	networkHeight = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "bridge_network_height",
		Help: "Network height of the Celestia node",
	})
)

func init() {
	prometheus.MustRegister(localHeight)
	prometheus.MustRegister(networkHeight)
}

func main() {
	listenPort := flag.String("listen.port", "8380", "port to listen on")
	endpoint := flag.String("endpoint", "http://localhost:26658", "endpoint to connect to")
	p2pNetwork := flag.String("p2p.network", "blockspacerace", "network to use")
	flag.Parse()

	http.Handle("/metrics", promhttp.Handler())

	go func() {
		client := &http.Client{}
		authToken := getAuthToken(*p2pNetwork)
		for {
			updateMetrics(client, authToken, *endpoint)
			time.Sleep(5 * time.Second)
		}
	}()

	fmt.Printf("Celestia Bridge Exporter started on port %s\n", *listenPort)
	http.ListenAndServe(":"+*listenPort, nil)
}

func updateMetrics(client *http.Client, authToken, endpoint string) {
	local, network, err := getHeights(client, authToken, endpoint)
	if err != nil {
		fmt.Println("Error getting heights:", err)
		return
	}

	localHeight.Set(float64(local))
	networkHeight.Set(float64(network))

}

func getHeights(client *http.Client, authToken, endpoint string) (int, int, error) {
	local := getHeight(client, authToken, "header.LocalHead", endpoint)
	network := getHeight(client, authToken, "header.NetworkHead", endpoint)

	return local, network, nil
}

func getHeight(client *http.Client, authToken, method, endpoint string) int {
	reqData := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  method,
		"params":  []interface{}{},
	}
	reqBytes, _ := json.Marshal(reqData)

	req, _ := http.NewRequest("POST", endpoint, bytes.NewReader(reqBytes))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authToken))

	resp, err := client.Do(req)
	if err != nil {
		return 0
	}
	defer resp.Body.Close()

	respBytes, _ := ioutil.ReadAll(resp.Body)
	var respData map[string]interface{}
	json.Unmarshal(respBytes, &respData)

	heightStr := respData["result"].(map[string]interface{})["header"].(map[string]interface{})["height"].(string)
	height, err := strconv.Atoi(heightStr)
	if err != nil {
		fmt.Printf("Error converting height to int: %v\n", err)
		return 0
	}
	return height
}

func getAuthToken(p2pNetwork string) string {
	out, err := exec.Command("celestia", "bridge", "auth", "admin", "--p2p.network", p2pNetwork).Output()
	if err != nil {
		fmt.Println("Error getting auth token:", err)
		return ""
	}

	return strings.TrimSpace(string(out))
}
