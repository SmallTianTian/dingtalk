package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiProcessGettodonumRequest() *OapiProcessGettodonumRequest {
	return &OapiProcessGettodonumRequest{}
}

type OapiProcessGettodonumRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessGettodonumResponse
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiProcessGettodonumRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessGettodonumRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessGettodonumRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiProcessGettodonumRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiProcessGettodonumRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.gettodonum"
}
func (this *OapiProcessGettodonumRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessGettodonumRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessGettodonumRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessGettodonumRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessGettodonumRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiProcessGettodonumRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Userid, taobao.DATA_OUTGOING_USER_ID); err != nil {
		return err
	}
	return nil
}
func (this *OapiProcessGettodonumRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessGettodonumRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiProcessGettodonumResponse struct {
	taobao.TaobaoResponse
	Count   int64  `json:"count,omitempty"`
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
