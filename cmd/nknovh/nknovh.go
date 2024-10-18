package main

import (
	"context"
	"log"
	"nknovh/internal/jobs"
	"time"

	x "nknovh-engine"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var e x.NKNOVH
	if err := e.Build(); err != nil {
		log.Fatal("Build() returned error: ", err.Error())
	}
	conf := e.GetConfig()

	job, err := jobs.NewJobsEngine(conf.Db, ctx, 10*time.Second, e.GetLogger(), conf.PoxyAddress, conf.ProxyLogin, conf.ProxyPassword, conf.NodesPath)
	if err != nil {
		panic(err)
	}
	job.Run()
	go func() {
		if err := e.Run(); err != nil {
			log.Fatal("Run() returned error: ", err.Error())
		}
	}()

	e.Listen()
	cancel()

}
