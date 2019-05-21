package controller

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func Logout(ctx *fasthttp.RequestCtx)  {
	fmt.Fprintf(ctx, "Logout url is %q\n", ctx.RequestURI())
}