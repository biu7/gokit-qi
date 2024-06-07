package response

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"net/http"
)

const (
	ContextResponse = "gin_response"
)

const (
	CodeOK        = http.StatusOK
	CodeError     = http.StatusBadRequest
	CodeAuthError = http.StatusUnauthorized
)

var marshaller = protojson.MarshalOptions{
	AllowPartial:      true,
	UseEnumNumbers:    true,
	EmitDefaultValues: true,
}

func ProtoJSON(c *gin.Context, code int, data proto.Message, msg string) {
	// for logging
	c.Set(ContextResponse, &CommonResponse{
		Code:    int32(code),
		Message: msg,
	})
	var anyData *anypb.Any
	if data != nil {
		anyData, _ = anypb.New(data)
	}
	b, _ := marshaller.Marshal(&CommonResponse{
		Code:    int32(code),
		Message: msg,
		Data:    anyData,
	})
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.Render(http.StatusOK, render.String{
		Format: "%s",
		Data:   []any{string(b)},
	})
}

func Success(c *gin.Context, body proto.Message) {
	ProtoJSON(c, CodeOK, body, "success")
}

func Fail(c *gin.Context, err error) {
	ProtoJSON(c, CodeError, nil, err.Error())
}

func AuthFail(c *gin.Context, err error) {
	ProtoJSON(c, CodeAuthError, nil, err.Error())
}

func SetResponseStatus(c *gin.Context, code int, msg string) {
	c.Set(ContextResponse, &CommonResponse{
		Code:    int32(code),
		Message: msg,
	})
}

func GetResponseStatus(c *gin.Context) (int32, string) {
	respStatus, ok := c.Get(ContextResponse)
	if !ok {
		return 0, ""
	}
	commonResp, ok := respStatus.(*CommonResponse)
	if !ok {
		return 0, ""
	}
	return commonResp.GetCode(), commonResp.GetMessage()
}
