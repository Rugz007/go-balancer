package util

import "github.com/valyala/fasthttp"

func IsHostAlive(h *fasthttp.Client, url string) bool {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodGet)
	resp := fasthttp.AcquireResponse()
	err := h.Do(req, resp)

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	if err != nil {
		return false
	}
	return resp.StatusCode() == 200
}
