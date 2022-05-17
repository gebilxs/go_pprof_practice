package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"time"

	"bytedance/byte_dance_projects/day_3_go-pprof-practice/animal"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetOutput(os.Stdout)

	runtime.GOMAXPROCS(1)              //限制CPU的使用数目
	runtime.SetMutexProfileFraction(1) //开启锁调用跟踪
	runtime.SetBlockProfileRate(1)     // 开启阻塞调用跟踪

	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	for {
		for _, v := range animal.AllAnimals {
			v.Live()
		}
		time.Sleep(time.Second)
	}
}
