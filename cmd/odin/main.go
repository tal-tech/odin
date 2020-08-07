package main

import (
	"odin/app"
	"odin/version"

	bs "github.com/tal-tech/hera/bootstrap"
	"github.com/tal-tech/hera/rpcxserver"
	rpcxplugin "github.com/tal-tech/odinPlugin"
	"github.com/tal-tech/xtools/addrutil"
	"github.com/tal-tech/xtools/confutil"
	"github.com/tal-tech/xtools/flagutil"
)

func main() {
	/*
	 * Usage: show app version
	 * Command: bin/odin -v
	 */
	showVersion := flagutil.GetVersion()
	if *showVersion {
		version.Version()
		return
	}
	/*
	 * init config
	 */
	confutil.InitConfig()

	addr, err := addrutil.Extract("")
	if err != nil {
		panic(err)
	}
	/*
	 * init server
	 */
	s := rpcxserver.NewServer(rpcxserver.Addr(addr))

	s.AddBeforeServerStartFunc(bs.InitLoggerWithConf(), bs.InitTraceLogger("Odin", "0.1"))
	/*
	 * Optional Plugin
	 * 1.Perf: s.AddBeforeServerStartFunc(bs.InitPerfutil())
	 * 2.Pprof: s.AddBeforeServerStartFunc(bs.InitPprof())
	 * 3.Expvar: s.AddBeforeServerStartFunc(bs.InitExpvar())
	 * 4.Maxfd: s.AddBeforeServerStartFunc(bs.InitMaxFd())
	 */

	s.AddBeforeServerStartFunc(s.InitConfig(),
		s.InitRegistry(),
		s.InitRpcxPlugin(rpcxplugin.RatelimitOption(), rpcxplugin.TraceOption(), rpcxplugin.ContextOption(), rpcxplugin.RecoveryOption(), rpcxplugin.CostWarnOption()),
		s.DisableHTTPGateway(),
		s.RegisterServiceWithPlugin("Odin", app.NewService(), ""))
	s.AddAfterServerStopFunc(bs.CloseLogger())
	err = s.Serve()
	if err != nil {
		panic(err)
	}
}
