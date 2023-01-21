package main

import (
	tiktokbackendbase "github.com/demonwuwen/tiktokBackend/kitex_gen/tiktokBackendBase/baseapiservice"
	"log"
)

func main() {
	svr := tiktokbackendbase.NewServer(new(BaseApiServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
