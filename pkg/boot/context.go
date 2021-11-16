package boot

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/acl"
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/session"
	"github.com/rwcoding/goback/pkg/util"
	"github.com/rwcoding/goback/pkg/validator"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type HandleFunc func(ctx *Context)

type Context struct {
	g         *gin.Context
	api       string
	sessionId string
	sign      string
	time      int
	body      []byte //http body

	adminer *models.Adminer
}

func (c *Context) Verify() *api.Response {
	var adminerId uint32 = 0
	needAcl := !acl.IsDefault(c.api)
	if c.api == "" || c.sign == "" {
		return api.NewErrorResponse("参数缺失")
	}

	if needAcl && c.sessionId == "" {
		return api.NewResponse(api.CodeNeedLogin, "会话缺失", nil)
	}

	var buf bytes.Buffer
	buf.Write([]byte(c.api))
	if c.g.Request.Method == http.MethodPost {
		b, err := ioutil.ReadAll(c.g.Request.Body)
		if err != nil {
			return api.NewErrorResponse(err.Error())
		}
		c.body = b
		buf.Write(b)
	}
	buf.Write([]byte(strconv.Itoa(c.time)))

	// 对于不需要登录的访问接口，不验证session，客户端不需要将session和key计入签名参数
	if c.sessionId != "" {
		buf.Write([]byte(c.sessionId))
		sess := session.QuerySession(c.sessionId)
		if sess == nil {
			return api.NewResponse(api.CodeNeedLogin, "会话错误", nil)
		}
		key := sess.SessionValue
		adminerId = sess.AdminerId
		buf.Write([]byte(key))
	}

	if util.Md5Byte(buf.Bytes()) != c.sign {
		return api.NewResponse(api.CodeNeedLogin, "签名错误", nil)
	}

	// 判断用户状态
	// 判断权限
	if needAcl {
		adminer := &models.Adminer{}
		models.GetDb().Take(adminer, adminerId)
		if !adminer.IsOK() {
			return api.NewResponse(api.CodeNeedLogin, "用户被锁定，无法访问", nil)
		}
		c.adminer = adminer

		if !acl.Verify(adminer, c.api) {
			return api.NewResponse(api.CodeNeedLogin, "您没有权限", nil)
		}
	}
	return nil
}

func (c *Context) Render(response *api.Response) {
	c.g.Data(http.StatusOK, "application/json", response.Json())
}

func (c *Context) Error(err error) {
	c.g.Data(http.StatusOK, "application/json", api.NewErrorResponse(err.Error()).Json())
}

func (c *Context) GinCtx() *gin.Context {
	return c.g
}

func (c *Context) NewRequest(r interface{}) error {
	err := json.Unmarshal(c.body, r)
	if err != nil {
		return err
	}
	e := validator.Verify(r)
	if e != nil {
		sb := strings.Builder{}
		for _, v := range e {
			sb.WriteString(v)
		}
		return errors.New(sb.String())
	}
	return nil
}

func (c *Context) GetSession() string {
	return c.sessionId
}

func (c *Context) GetTime() int {
	return c.time
}

func (c *Context) GetAdminer() *models.Adminer {
	return c.adminer
}

func (c *Context) GetRemote() string {
	ip, ok := c.g.RemoteIP()
	if !ok {
		return ""
	}
	return ip.String()
}
