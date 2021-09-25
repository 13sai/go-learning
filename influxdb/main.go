package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/sirupsen/logrus"
)

// You can generate a Token from the "Tokens Tab" in the UI
const token = "DrSIOHeQENB1T6o0jA2DpjYT_hs0aMjHszico96AIlLU-KTvJUkrUk0Dx1eOx7t-8wSTL2iMl63Gvh4q3ml9lw=="
const bucket = "test"
const org = "admin"

var client influxdb2.Client

func main() {
	client = influxdb2.NewClient("http://localhost:8086", token)
	a, err := client.Health(context.Background())
	if err != nil {
		logrus.Errorf("conn err: %v", err)
		return
	}

	logrus.Infof("conn success: %v", a.Status)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go write(ctx)

	defer client.Close()
	g := gin.New()

	server := &http.Server{
		Handler: g,
		Addr:    ":8088",
	}

	g.GET("/query", query)

	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	go func() {
		logrus.Infof("listen at %v", server.Addr)
		if err := server.ListenAndServe(); err != nil {
			logrus.Errorf("listen err %v", err)
			cancel()
		}
	}()

	select {
	case <-ctx.Done():
		logrus.Info("webserver exit itself")
		server.Shutdown(context.Background())
	case <-stop:
		logrus.Infof("got exit signal, shutdown webserver")
		server.Shutdown(context.Background())
	}

}

func query(c *gin.Context) {
	queryAPI := client.QueryAPI(org)
	result, err := queryAPI.Query(context.Background(), `from(bucket:"`+bucket+`")|> range(start: -1h) |> filter(fn: (r) => r._measurement == "machine")`)
	str := ""
	if err == nil {
		for result.Next() {
			// Notice when group key has changed
			if result.TableChanged() {
				str += fmt.Sprintf("table: %s\n", result.TableMetadata().String())
			}
			// Access data
			str += fmt.Sprintf("value: %v\n", result.Record().Value())
		}
		// check for an error
		if result.Err() != nil {
			fmt.Printf("query parsing error: %s\n", result.Err().Error())
		}
	} else {
		panic(err)
	}

	c.String(200, "", str)
}

func write(ctx context.Context) {
	var ticker = time.NewTicker(time.Second * 1)
	for {
		select {
		case <-ctx.Done():
			logrus.Infof("loader exit.")
			return
		case <-ticker.C:
			// v, _ := mem.VirtualMemory()
			// cpuPercent, _ := cpu.Percent(time.Second, false)
			parts, _ := disk.Partitions(true)
			diskInfo, _ := disk.Usage(parts[0].Mountpoint)
			data := map[string]interface{}{
				// "cpu": cpuPercent,
				// "memory": v.Free,
				"disk": diskInfo.UsedPercent,
			}

			writeAPI := client.WriteAPIBlocking(org, bucket)

			tags := map[string]string{"abc": "rjgret"}

			pt := influxdb2.NewPoint("machine", tags, data, time.Now())

			err := writeAPI.WritePoint(context.Background(), pt)
			if err != nil {
				logrus.Infof("insert fail: %v", err)
			}
			logrus.Infof("machine %+v", data)
		}
	}
}
