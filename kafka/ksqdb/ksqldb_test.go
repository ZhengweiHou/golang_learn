package main

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/thmeitz/ksqldb-go"
	"github.com/thmeitz/ksqldb-go/net"
)

func Test_CreateKsqlDBClient(t *testing.T) {
	options := net.Options{
		// if you need a login, do this; if not its not necessary
		Credentials: net.Credentials{Username: "myuser", Password: "mypassword"},
		// defaults to http://localhost:8088
		BaseUrl: "http://my-super-shiny-ksqldbserver:8082",
		// this is needed, because the ksql api communicates with http2 only
		AllowHTTP: true,
	}

	// only log.Logger is allowed or nil
	// logrus is in maintenance mode, so I'll using zap in the future
	client, err := net.NewHTTPClient(options, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
}

func Test_PullQuery(t *testing.T) {
	options := net.Options{
		// Credentials: net.Credentials{Username: "user", Password: "password"},
		BaseUrl:   "http://localhost:8088",
		AllowHTTP: true,
	}

	kcl, err := ksqldb.NewClientWithOptions(options)
	if err != nil {
		log.Fatal(err)
	}
	defer kcl.Close()

	// query := `select timestamptostring(windowstart,'yyyy-MM-dd HH:mm:ss','Europe/London') as window_start,
	// timestamptostring(windowend,'HH:mm:ss','Europe/London') as window_end, dog_size, dogs_ct
	// from dogs_by_size where dog_size=?;`

	query := `select id,name from hzwkfk_stream emit changes;`

	stmnt, err := ksqldb.QueryBuilder(query)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	qOpts := (&ksqldb.QueryOptions{Sql: *stmnt}).EnablePullQueryTableScan(false)

	_, r, err := kcl.Pull(ctx, *qOpts)
	if err != nil {
		log.Fatal(err)
	}

	var id string
	var name string
	for _, row := range r {

		if row != nil {
			// Should do some type assertions here
			id = row[0].(string)
			name = row[1].(string)
			log.Printf("id:%s,name:%s\n", id, name)
		}
	}
}

func Test_PushQuery(t *testing.T) {
	options := net.Options{
		// Credentials: net.Credentials{Username: "user", Password: "password"},
		BaseUrl:   "http://localhost:8088",
		AllowHTTP: true,
	}

	kcl, err := ksqldb.NewClientWithOptions(options)
	if err != nil {
		log.Fatal(err)
	}
	defer kcl.Close()

	// query := `select timestamptostring(windowstart,'yyyy-MM-dd HH:mm:ss','Europe/London') as window_start,
	// timestamptostring(windowend,'HH:mm:ss','Europe/London') as window_end, dog_size, dogs_ct
	// from dogs_by_size where dog_size=?;`

	query := `select id,name from hzwkfk_stream emit changes;`

	rowChannel := make(chan ksqldb.Row)
	headerChannel := make(chan ksqldb.Header, 1)

	// This Go routine will handle rows as and when they
	// are sent to the channel
	go func() {
		var id string
		var name string
		for row := range rowChannel {
			if row != nil {

				id = row[0].(string)
				name = row[1].(string)
				log.Printf("id:%s,name:%s\n", id, name)
			}
		}
	}()

	go func() {
		for head := range headerChannel {
			log.Printf("head:%vs\n", head)
		}
	}()

	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()

	e := kcl.Push(ctx, ksqldb.QueryOptions{Sql: query}, rowChannel, headerChannel)
	if e != nil {
		log.Fatal(e)
	}
}
