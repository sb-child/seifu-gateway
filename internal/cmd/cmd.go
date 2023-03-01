package cmd

import (
	"context"
	"os"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gproc"

	"seifu-gw/internal/controller"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func:  startHttpServer,
	}
)

func startHttpServer(ctx context.Context, parser *gcmd.Parser) (err error) {
	// main server is used for frontend
	mainServer := g.Server()
	mainServer.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		group.Bind(
			// todo
			controller.Hello,
		)
	})
	// proxy servers are handle HTTP requests, do some scanning,
	// then forward to origin server or redirect to main server
	// todo
	// ---
	// start the main server
	mainServer.Start()
	// start proxy servers
	// todo
	// ---
	// if a shutdown signal received
	gproc.AddSigHandlerShutdown(func(sig os.Signal) {
		g.Log().Infof(ctx, "%s Signal received, stopping service...", sig.String())
		// stop the main server
		mainServer.Shutdown()
		// stop proxy servers
		// todo
		// ---
		// stop other services
		// ---
	})
	// listen for signals
	gproc.Listen()
	// wait for http servers stopping
	g.Wait()
	// wait for other services stopping
	// ---
	return nil
}
