package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiSmartdeviceVisitorAddvisitorRequest() *OapiSmartdeviceVisitorAddvisitorRequest {
	return &OapiSmartdeviceVisitorAddvisitorRequest{}
}

type OapiSmartdeviceVisitorAddvisitorRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceVisitorAddvisitorResponse
	TopHttpMethod   string
	TopResponseType string
	VisitorVo       string
}

func (this *OapiSmartdeviceVisitorAddvisitorRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceVisitorAddvisitorRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceVisitorAddvisitorRequest) SetVisitorVo(visitorVo2 string) {
	this.VisitorVo = visitorVo2
}
func (this *OapiSmartdeviceVisitorAddvisitorRequest) GetVisitorVo() string {
	return this.VisitorVo
}
func (this *OapiSmartdeviceVisitorAddvisitorRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.visitor.addvisitor"
}
func (this *OapiSmartdeviceVisitorAddvisitorRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceVisitorAddvisitorRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceVisitorAddvisitorRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceVisitorAddvisitorRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceVisitorAddvisitorRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["visitor_vo"] = this.VisitorVo
	return txtParams
}
func (this *OapiSmartdeviceVisitorAddvisitorRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiSmartdeviceVisitorAddvisitorRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceVisitorAddvisitorRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type DidoVisitorVo struct {
	AppointedEndtime   int64    `json:"appointed_endtime,omitempty"`
	AppointedStarttime int64    `json:"appointed_starttime,omitempty"`
	ExtraInfo          string   `json:"extra_info,omitempty"`
	MediaId            string   `json:"media_id,omitempty"`
	Mobile             string   `json:"mobile,omitempty"`
	NotifyUserList     []string `json:"notify_user_list,omitempty"`
	RecognizeEndtime   int64    `json:"recognize_endtime,omitempty"`
	RecognizeStarttime int64    `json:"recognize_starttime,omitempty"`
	UserName           string   `json:"user_name,omitempty"`
	UserType           string   `json:"user_type,omitempty"`
	Userid             string   `json:"userid,omitempty"`
}
type OapiSmartdeviceVisitorAddvisitorResponse struct {
	taobao.TaobaoResponse
	Result string `json:"result,omitempty"`
}
