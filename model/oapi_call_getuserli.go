package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCallGetuserlistRequest() *OapiCallGetuserlistRequest {
	return &OapiCallGetuserlistRequest{}
}

type OapiCallGetuserlistRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCallGetuserlistResponse
	Offset          int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCallGetuserlistRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCallGetuserlistRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCallGetuserlistRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *OapiCallGetuserlistRequest) GetOffset() int64 {
	return this.Offset
}
func (this *OapiCallGetuserlistRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiCallGetuserlistRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiCallGetuserlistRequest) GetApiMethodName() string {
	return "dingtalk.oapi.call.getuserlist"
}
func (this *OapiCallGetuserlistRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCallGetuserlistRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCallGetuserlistRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCallGetuserlistRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCallGetuserlistRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["offset"] = this.Offset
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiCallGetuserlistRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCallGetuserlistRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCallGetuserlistRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCallGetuserlistResponse struct {
	taobao.TaobaoResponse
	Result Result `json:"result,omitempty"`
}
