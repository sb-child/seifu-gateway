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
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					controller.Hello,
				)
			})
			s.Start()
			gproc.AddSigHandlerShutdown(func(sig os.Signal) {
				g.Log().Infof(ctx, "%s Signal received, stopping service...", sig.String())
				s.Shutdown()
				// stop other services
				// ---
			})
			gproc.Listen()
			g.Wait()
			// wait other services
			// ---
			return nil
		},
	}
)
