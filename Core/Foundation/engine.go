package Foundation

import (
	"WebServer/Core/Logger"
	"WebServer/Core/Settings"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kataras/iris/v12"
	"os"
	"os/signal"
	"reflect"
	"syscall"
	"time"
)

type IEngine interface {
}

type Engine struct {
	Config      *Config
	App         *iris.Application
	StartTime   time.Time
	SignalEvent map[os.Signal]func()
}

var _ IEngine = &Engine{}

func New() *Engine {
	startTime := time.Now()
	engine := &Engine{
		Config: &Config{
			App:     &Settings.AppConf{},
			Web:     &Settings.WebConf{},
			CacheDB: &Settings.CacheDbConf{},
			envMap:  make(map[string]interface{}),
		},
		SignalEvent: make(map[os.Signal]func()),
	}
	engine.StartTime = startTime
	engine.Config.engine = engine
	engine.Config.raw, _ = godotenv.Read()
	engine.Config.systemExternalEnv("app", engine.Config.App)
	engine.Config.systemExternalEnv("web", engine.Config.Web)
	engine.Config.systemExternalEnv("cachedb", engine.Config.CacheDB)
	Logger.SysLog.Info("[Engine] Environment Loaded")
	GetServer().ConnectCacheDbService(engine.Config.CacheDB)
	engine.initEngine()
	return engine
}

func (engine *Engine) initEngine() {
	engine.App = iris.New()
	InitCoreRouters(engine.App)
}

func (engine *Engine) UsingCacheDBService() {
	GetServer().ConnectCacheDbService(engine.Config.CacheDB)
}

func (engine *Engine) RegisterSystem(Cmd string, Instance IWebSystem) {
	registered := GetServer().Register(Cmd, Instance)
	if !registered {
		Logger.SysLog.Warnf(
			"[Engine] Register System: %s, %s Failed, System Command Duplicate!",
			Cmd,
			reflect.TypeOf(Instance),
		)
		return
	}
	Logger.SysLog.Debug(
		"[Engine] Registered System: %s, %s",
		Cmd,
		reflect.TypeOf(Instance),
	)
}

func (engine *Engine) Signal(signal os.Signal, call func()) {
	engine.SignalEvent[signal] = call
}

func (engine *Engine) Serve(app *iris.Application) {
	go func() {
		endPoint := fmt.Sprintf(":%d", engine.Config.Web.HttpPort)
		serveTime := time.Now()
		Logger.SysLog.Infof("[Engine] Serving HTTP(%s) in %dms", endPoint, serveTime.Sub(engine.StartTime).Milliseconds())
		if err := app.Run(iris.Addr(endPoint), iris.WithoutInterruptHandler); err != nil {
			Logger.SysLog.Warnf("[Engine] Stop Serving (%s)", err)
		}
	}()

	signalChan := make(chan os.Signal, 1)
	exitChan := make(chan bool)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT, os.Kill, syscall.SIGKILL, syscall.SIGTERM)
	go func() {
		sig := <-signalChan
		Logger.SysLog.Warnf("[Engine] Caught Signal(%03d)", sig)

		if call, ok := engine.SignalEvent[sig]; ok {
			call()
		}

		if err := app.Shutdown(context.Background()); err != nil {
			Logger.SysLog.Warnf("[Engine] Shutdown Server with ErrorsCode, %s", err)
		}
		exitChan <- true
	}()
	<-exitChan
	Logger.SysLog.Warn("[Engine] Shutdown Server")
}






