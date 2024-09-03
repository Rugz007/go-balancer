package proxy

import (
	types "github.com/rugz007/go-balancer/types"
	"github.com/valyala/fasthttp"
)

type BackendProxy struct {
	types.BackendProxy
	Backend types.Backend
	Client  *fasthttp.Client
}

func CreateProxyClient(backend types.Backend) types.BackendProxy {
	client := &fasthttp.Client{
		MaxConnsPerHost:     backend.MaxConns,
		MaxConnDuration:     backend.MaxConnDuration,
		MaxIdleConnDuration: backend.MaxIdleConnDuration,
	}

	return types.BackendProxy{
		Backend: backend,
		Client:  client,
	}
}

func HandleRequestViaProxy(bp *types.BackendProxy, ctx *fasthttp.RequestCtx) {
	proxyUrl := bp.Backend.Url + string(ctx.Request.URI().RequestURI())
	req := &ctx.Request
	resp := &ctx.Response

	req.SetRequestURI(proxyUrl)
	req.Header.SetMethodBytes(ctx.Method())
	req.Header.SetHostBytes(ctx.Host())
	req.Header.Set("X-Forwarded-For", ctx.RemoteIP().String())

	if bp.Backend.HealthCheckPath != "" && string(ctx.Path()) == bp.Backend.HealthCheckPath {
		resp.SetStatusCode(fasthttp.StatusOK)
		return
	}

	err := bp.Client.Do(req, resp)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusServiceUnavailable)
	}
}
