package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiV2SafeSetenableRequest() *OapiV2SafeSetenableRequest {
	return &OapiV2SafeSetenableRequest{}
}

type OapiV2SafeSetenableRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiV2SafeSetenableResponse
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiV2SafeSetenableRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiV2SafeSetenableRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiV2SafeSetenableRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiV2SafeSetenableRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiV2SafeSetenableRequest) GetApiMethodName() string {
	return "dingtalk.oapi.v2.safe.setenable"
}
func (this *OapiV2SafeSetenableRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiV2SafeSetenableRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiV2SafeSetenableRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiV2SafeSetenableRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiV2SafeSetenableRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiV2SafeSetenableRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Userid, taobao.DATA_OUTGOING_USER_ID); err != nil {
		return err
	}
	return nil
}
func (this *OapiV2SafeSetenableRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiV2SafeSetenableRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiV2SafeSetenableResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
