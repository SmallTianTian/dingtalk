package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiExtcontactListRequest() *OapiExtcontactListRequest {
	return &OapiExtcontactListRequest{}
}

type OapiExtcontactListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiExtcontactListResponse
	Offset          int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiExtcontactListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiExtcontactListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiExtcontactListRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *OapiExtcontactListRequest) GetOffset() int64 {
	return this.Offset
}
func (this *OapiExtcontactListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiExtcontactListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiExtcontactListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.extcontact.list"
}
func (this *OapiExtcontactListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiExtcontactListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiExtcontactListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiExtcontactListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiExtcontactListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["offset"] = this.Offset
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiExtcontactListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiExtcontactListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiExtcontactListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiExtcontactListResponse struct {
	taobao.TaobaoResponse
	Errcode int64            `json:"errcode,omitempty"`
	Errmsg  string           `json:"errmsg,omitempty"`
	Results []OpenExtContact `json:"results,omitempty"`
}
