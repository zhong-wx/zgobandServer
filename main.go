package main

import (
	"./gen-go/zgobandRPC"
	"./handler"
	"./models"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"os"
)

func main() {
	models.Open()
	defer models.Close()

	ret := runServer()
	if(ret != nil) {
		fmt.Println("runSever fail:", ret.Error())
	}
}

func runServer() error{
	addr := "localhost:9091"
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transport, err := thrift.NewTServerSocket(addr)
	if(err != nil) {
		fmt.Println("err:" , err.Error())
		os.Exit(-1)
	}

	gameHallHandler := handler.NewGameHallHandler()
	loginAndRegHandler := handler.NewLoginAndRegHandler()
	gameOperatorHandler := handler.NewGameOperatorHandler()
	gameOperatorProcessor := zgobandRPC.NewGameOperatorProcessor(gameOperatorHandler)
	gameHallHandlerProcessor := zgobandRPC.NewGameHallProcessor(gameHallHandler)
	loginAndRegHandlerProcessor := zgobandRPC.NewLoginAndRegProcessor(loginAndRegHandler)

	multiProcessor := thrift.NewTMultiplexedProcessor()
	multiProcessor.RegisterProcessor("GameHall", gameHallHandlerProcessor)
	multiProcessor.RegisterProcessor("LoginAndReg", loginAndRegHandlerProcessor)
	multiProcessor.RegisterProcessor("GameOperator", gameOperatorProcessor)
	server := thrift.NewTSimpleServer4(multiProcessor, transport, transportFactory, protocolFactory)

	fmt.Println("starting the simple server... on", addr)
	return server.Serve()
}