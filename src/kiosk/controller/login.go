package controller

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func LoginByEmail(ctx *fasthttp.RequestCtx)  {
	fmt.Fprintf(ctx, "Login url is %q\n", ctx.RequestURI())
}
