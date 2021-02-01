package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiOpenencryptHeartbeatRequest() *OapiOpenencryptHeartbeatRequest {
	return &OapiOpenencryptHeartbeatRequest{}
}

type OapiOpenencryptHeartbeatRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiOpenencryptHeartbeatResponse
	Appid           int64
	Extension       string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiOpenencryptHeartbeatRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiOpenencryptHeartbeatRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiOpenencryptHeartbeatRequest) SetAppid(appid2 int64) {
	this.Appid = appid2
}
func (this *OapiOpenencryptHeartbeatRequest) GetAppid() int64 {
	return this.Appid
}
func (this *OapiOpenencryptHeartbeatRequest) SetExtension(extension2 string) {
	this.Extension = extension2
}
func (this *OapiOpenencryptHeartbeatRequest) GetExtension() string {
	return this.Extension
}
func (this *OapiOpenencryptHeartbeatRequest) GetApiMethodName() string {
	return "dingtalk.oapi.openencrypt.heartbeat"
}
func (this *OapiOpenencryptHeartbeatRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiOpenencryptHeartbeatRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiOpenencryptHeartbeatRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiOpenencryptHeartbeatRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiOpenencryptHeartbeatRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["appid"] = this.Appid
	txtParams["extension"] = this.Extension
	return txtParams
}
func (this *OapiOpenencryptHeartbeatRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Appid, "appid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiOpenencryptHeartbeatRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiOpenencryptHeartbeatRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiOpenencryptHeartbeatResponse struct {
	taobao.TaobaoResponse
	Result  bool `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}
