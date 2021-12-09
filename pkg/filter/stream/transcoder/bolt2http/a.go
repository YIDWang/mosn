package bolt2http

import (
	"context"

	"github.com/valyala/fasthttp"
	"mosn.io/mosn/pkg/filter/stream/transcoder"
	"mosn.io/mosn/pkg/protocol/http"
	"mosn.io/mosn/pkg/protocol/xprotocol/bolt"
	"mosn.io/mosn/pkg/types"
)

func init() {
	transcoder.MustRegister("http2bolt_simple", &http2bolt{})
}

type http2bolt struct{}

func (t *http2bolt) Accept(ctx context.Context, headers types.HeaderMap, buf types.IoBuffer, trailers types.HeaderMap) bool {
	_, ok := headers.(http.RequestHeader)
	return ok
}

func (t *http2bolt) TranscodingRequest(ctx context.Context, headers types.HeaderMap, buf types.IoBuffer, trailers types.HeaderMap) (types.HeaderMap, types.IoBuffer, types.HeaderMap, error) {
	targetRequest := bolt.NewRpcRequest(0, headers, buf)
	return targetRequest, buf, trailers, nil
}

func (t *http2bolt) TranscodingResponse(ctx context.Context, headers types.HeaderMap, buf types.IoBuffer, trailers types.HeaderMap) (types.HeaderMap, types.IoBuffer, types.HeaderMap, error) {
	sourceResponse, ok := headers.(*bolt.Response)
	if !ok {
		// if the response is not bolt response, it maybe come from hijack or send directly response.
		// so we just returns the original data
		return headers, buf, trailers, nil
	}
	targetResponse := fasthttp.Response{}

	// 1. headers
	sourceResponse.Range(func(Key, Value string) bool {
		targetResponse.Header.Set(Key, Value)
		return true
	})
	// 2. status code
	if sourceResponse.ResponseStatus != bolt.ResponseStatusSuccess {
		targetResponse.SetStatusCode(http.InternalServerError)
	}
	return http.ResponseHeader{ResponseHeader: &targetResponse.Header}, buf, trailers, nil
}
