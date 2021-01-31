package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCcoserviceEntranceSendnotifyRequest() *OapiCcoserviceEntranceSendnotifyRequest {
	return &OapiCcoserviceEntranceSendnotifyRequest{}
}

type OapiCcoserviceEntranceSendnotifyRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCcoserviceEntranceSendnotifyResponse
	AppId           int64
	Content         string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiCcoserviceEntranceSendnotifyRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCcoserviceEntranceSendnotifyRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCcoserviceEntranceSendnotifyRequest) SetAppId(appId2 int64) {
	this.AppId = appId2
}
func (this *OapiCcoserviceEntranceSendnotifyRequest) GetAppId() int64 {
	return this.AppId
}
func (this *OapiCcoserviceEntranceSendnotifyRequest) SetContent(content2 string) {
	this.Content = content2
}
func (this *OapiCcoserviceEntranceSendnotifyRequest) GetContent() string {
	return this.Content
}
func (this *OapiCcoserviceEntranceSendnotifyRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiCcoserviceEntranceSendnotifyRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiCcoserviceEntranceSendnotifyRequest) GetApiMethodName() string {
	return "dingtalk.oapi.ccoservice.entrance.sendnotify"
}
func (this *OapiCcoserviceEntranceSendnotifyRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCcoserviceEntranceSendnotifyRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCcoserviceEntranceSendnotifyRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCcoserviceEntranceSendnotifyRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCcoserviceEntranceSendnotifyRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["app_id"] = this.AppId
	txtParams[taobao.DATA_CONTENT] = this.Content
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiCcoserviceEntranceSendnotifyRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AppId, "appId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCcoserviceEntranceSendnotifyRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCcoserviceEntranceSendnotifyRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCcoserviceEntranceSendnotifyResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
