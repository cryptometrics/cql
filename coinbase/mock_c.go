package coinbase

import (
	"bytes"
	"cql/client2"
	"io/ioutil"
	"net/http"

	. "github.com/franela/goblin"
	"github.com/golang/mock/gomock"
)

type mockC struct {
	desc      string
	g         *G
	ctrl      *gomock.Controller
	data      []byte
	endpoint  Endpoint
	assertReq func(*G, client2.Request)
	assert    func(*G, client2.Connector)
}

func (ctx *mockC) run() {
	ctx.g.It(ctx.desc, func() {
		c := client2.NewMockC(ctx.ctrl)
		c.EXPECT().Connect().Return(nil)
		c.EXPECT().Do(gomock.Any()).DoAndReturn(
			func(req client2.Request) (*http.Response, error) {
				ctx.g.Assert(req.Endpoint).Equal(ctx.endpoint)
				ctx.assertReq(ctx.g, req)
				stringReader := bytes.NewReader(ctx.data)
				stringReadCloser := ioutil.NopCloser(stringReader)
				return &http.Response{Body: stringReadCloser}, nil
			})
		c.EXPECT().Identifier().Return("client mocker").AnyTimes()

		conn := func() (client2.C, error) { return c, nil }
		ctx.assert(ctx.g, conn)
	})
}
