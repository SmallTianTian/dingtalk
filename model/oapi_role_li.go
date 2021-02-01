package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRoleListRequest() *OapiRoleListRequest {
	return &OapiRoleListRequest{}
}

type OapiRoleListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRoleListResponse
	Offset          int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRoleListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRoleListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRoleListRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *OapiRoleListRequest) GetOffset() int64 {
	return this.Offset
}
func (this *OapiRoleListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiRoleListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiRoleListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.role.list"
}
func (this *OapiRoleListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRoleListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRoleListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRoleListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRoleListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["offset"] = this.Offset
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiRoleListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRoleListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRoleListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRoleListResponse struct {
	taobao.TaobaoResponse
	Result PageVo `json:"result,omitempty"`
}
