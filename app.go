package main

import (
	WebCore "WebServer/Core/Foundation"
	"WebServer/Features/Sample"
)

func main(){
	engine := WebCore.New()
	engine.RegisterSystem("sample", Sample.New()) // sample api system
	engine.Serve(engine.App)
}










