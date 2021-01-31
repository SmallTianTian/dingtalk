package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiUserBatchdeleteRequest() *OapiUserBatchdeleteRequest {
	return &OapiUserBatchdeleteRequest{}
}

type OapiUserBatchdeleteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiUserBatchdeleteResponse
	TopHttpMethod   string
	TopResponseType string
	Useridlist      []string
}

func (this *OapiUserBatchdeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiUserBatchdeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiUserBatchdeleteRequest) SetUseridlist(useridlist2 []string) {
	this.Useridlist = useridlist2
}
func (this *OapiUserBatchdeleteRequest) GetUseridlist() []string {
	return this.Useridlist
}
func (this *OapiUserBatchdeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.user.batchdelete"
}
func (this *OapiUserBatchdeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiUserBatchdeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiUserBatchdeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiUserBatchdeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiUserBatchdeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["useridlist"] = this.Useridlist
	return txtParams
}
func (this *OapiUserBatchdeleteRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiUserBatchdeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiUserBatchdeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiUserBatchdeleteResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
