package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartdeviceRemovefaceRequest() *OapiSmartdeviceRemovefaceRequest {
	return &OapiSmartdeviceRemovefaceRequest{}
}

type OapiSmartdeviceRemovefaceRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceRemovefaceResponse
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiSmartdeviceRemovefaceRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceRemovefaceRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceRemovefaceRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiSmartdeviceRemovefaceRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiSmartdeviceRemovefaceRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.removeface"
}
func (this *OapiSmartdeviceRemovefaceRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceRemovefaceRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceRemovefaceRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceRemovefaceRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceRemovefaceRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiSmartdeviceRemovefaceRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Userid, taobao.DATA_OUTGOING_USER_ID); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartdeviceRemovefaceRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceRemovefaceRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartdeviceRemovefaceResponse struct {
	taobao.TaobaoResponse
	Result bool `json:"result,omitempty"`
}
