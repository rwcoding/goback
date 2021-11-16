package login

import (
	"bytes"
	"encoding/base64"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"github.com/rwcoding/goback/pkg/session"
	"github.com/rwcoding/goback/pkg/util"
	"golang.org/x/image/font"
	"golang.org/x/image/font/inconsolata"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/png"
)

type captchaRequest struct {
	ctx *boot.Context

	Type  string `validate:"required"`
	ImgId string `validate:"omitempty,len=32" json:"img_id"`
}

type captchaResponse struct {
	Img   string `json:"img"`
	ImgId string `json:"img_id"`
}

func NewApiCaptcha(ctx *boot.Context) boot.Logic {
	return &captchaRequest{ctx: ctx}
}

func (request *captchaRequest) Run() *api.Response {
	dst := image.NewRGBA(image.Rect(0, 0, 80, 30))

	x := 20
	y := 20
	label := util.RandString(4)
	col := color.RGBA{200, 100, 0, 255}
	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}

	d := &font.Drawer{
		Dst:  dst,
		Src:  image.NewUniform(col),
		Face: inconsolata.Regular8x16, //Regular8x16
		Dot:  point,
	}
	d.DrawString(label)

	target := &bytes.Buffer{}
	err := png.Encode(target, dst)
	if err != nil {
		return api.NewErrorResponse("图片编码错误")
	}

	res := &captchaResponse{}
	res.Img = "data:image/png;base64," + base64.StdEncoding.EncodeToString(target.Bytes())

	sess := session.NewKVSession(label, request.ImgId)
	if sess == nil {
		return api.NewErrorResponse("生成图片错误")
	}

	res.ImgId = sess.SessionId

	return api.NewDataResponse(res)
}
