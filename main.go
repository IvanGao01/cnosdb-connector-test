package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"

	_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
	client "github.com/influxdata/influxdb1-client/v2"
)

func main() {
	CnosDBQuery()
}

func InfluxDBQuery() *Content {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	}
	defer c.Close()

	q := client.NewQuery("SELECT * FROM cpu limit 10", "telegraf", "")
	if response, err := c.Query(q); err == nil && response.Error() == nil {
		return ParseInfluxDBResult(response.Results[0])
	}
	return nil
}

func ParseInfluxDBResult(result client.Result) *Content {
	content := &Content{Title: nil, Table: nil}
	content.Title = result.Series[0].Columns
	content.Table = result.Series[0].Values
	return content
}

type Content struct {
	Title []string
	Table [][]interface{}
}

func Diff(source Content, target Content) string {
	return ""
}

func CnosDBQuery() *Content {
	// curl -i -u "username:password" -H "Accept: application/json" -XPOST ""http://localhost:31007/api/v1/sql\?db=telegraf -d 'SELECT * from cpu limit 10'
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, "http://localhost:31007/api/v1/sql?db=telegraf", bytes.NewReader([]byte("SELECT * from cpu limit 10")))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("username", "password")
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(resp.Body)
	fmt.Println(reader)
	return nil
}
