package main

import (
  //elasticsearch6 "github.com/elastic/go-elasticsearch/v6"
  //elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
  "fmt"
  "github.com/ahuigo/glogger"
  //"net/http" "time"
  "bytes"
  "context"
  "encoding/json"
  "strconv"
  "strings"
  "sync"

  "github.com/elastic/go-elasticsearch/v8"
  "github.com/elastic/go-elasticsearch/v8/esapi"
)

var logger = glogger.Glogger
func main(){
    var wg sync.WaitGroup
    cfg := elasticsearch.Config{
      //Username: "foo",
      //Password: "bar",
      Addresses: []string{
        "http://ahui.works:9200",
        //"http://ahui.works:9201",
        //"http://ahui.works:9300",
      },
    }
    es, err := elasticsearch.NewClient(cfg)
    if err!=nil{
        logger.Fatalf("error create es client:%s",err)
    }
    fmt.Println("es:", es)

    // 1. Get cluster info
	res, err := es.Info()
	if err != nil {
		logger.Fatalf("Error getting response: %s", err)
	}
	// Check response status
	if res.IsError() {
		logger.Fatalf("Error: %s", res.String())
	}
	// Deserialize the response into a map.
    var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		logger.Fatalf("Error parsing the response body: %s", err)
	}
	// Print client and server version numbers.
	logger.Infof("Client: %s", elasticsearch.Version)
	logger.Infof("Server: %s", r["version"].(map[string]interface{})["number"])
	logger.Info(strings.Repeat("~", 37))

	// 2. Index documents concurrently
	//
	for i, title := range []string{"Test One", "Test Two"} {
		wg.Add(1)

		go func(i int, title string) {
			defer wg.Done()

			// Build the request body.
			var b strings.Builder
			b.WriteString(`{"title" : "`)
			b.WriteString(title)
			b.WriteString(`"}`)

			// Set up the request object.
			req := esapi.IndexRequest{
				Index:      "test",
				DocumentID: strconv.Itoa(i + 1),
				Body:       strings.NewReader(b.String()),
				Refresh:    "true",
			}

			// Perform the request with the client.
			res, err := req.Do(context.Background(), es)
			if err != nil {
				logger.Fatalf("Error getting response: %s", err)
			}
			defer res.Body.Close()

			if res.IsError() {
				logger.Infof("[%s] Error indexing document ID=%d", res.Status(), i+1)
			} else {
				// Deserialize the response into a map.
				var r map[string]interface{}
				if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
					logger.Infof("Error parsing the response body: %s", err)
				} else {
					// Print the response status and indexed document version.
					logger.Infof("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
				}
			}
		}(i, title)
	}
	wg.Wait()

	logger.Info(strings.Repeat("-", 37))

	// 3. Search for the indexed documents
	//
	// Build the request body.
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "test",
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		logger.Fatalf("Error encoding query: %s", err)
	}

	// Perform the search request.
	res, err = es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("test"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		logger.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			logger.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			logger.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		logger.Fatalf("Error parsing the response body: %s", err)
	}
	// Print the response status, number of results, and request duration.
	logger.Infof(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)
	// Print the ID and document source for each hit.
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		logger.Infof(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}

	logger.Info(strings.Repeat("=", 37))

}
