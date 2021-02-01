package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiSmartdeviceVisitorSendnotifyRequest() *OapiSmartdeviceVisitorSendnotifyRequest {
	return &OapiSmartdeviceVisitorSendnotifyRequest{}
}

type OapiSmartdeviceVisitorSendnotifyRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceVisitorSendnotifyResponse
	ReservationId   string
	TopHttpMethod   string
	TopResponseType string
	VisitorNotifyVo string
}

func (this *OapiSmartdeviceVisitorSendnotifyRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceVisitorSendnotifyRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceVisitorSendnotifyRequest) SetReservationId(reservationId2 string) {
	this.ReservationId = reservationId2
}
func (this *OapiSmartdeviceVisitorSendnotifyRequest) GetReservationId() string {
	return this.ReservationId
}
func (this *OapiSmartdeviceVisitorSendnotifyRequest) SetVisitorNotifyVo(visitorNotifyVo2 string) {
	this.VisitorNotifyVo = visitorNotifyVo2
}
func (this *OapiSmartdeviceVisitorSendnotifyRequest) GetVisitorNotifyVo() string {
	return this.VisitorNotifyVo
}
func (this *OapiSmartdeviceVisitorSendnotifyRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.visitor.sendnotify"
}
func (this *OapiSmartdeviceVisitorSendnotifyRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceVisitorSendnotifyRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceVisitorSendnotifyRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceVisitorSendnotifyRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceVisitorSendnotifyRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["reservation_id"] = this.ReservationId
	txtParams["visitor_notify_vo"] = this.VisitorNotifyVo
	return txtParams
}
func (this *OapiSmartdeviceVisitorSendnotifyRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiSmartdeviceVisitorSendnotifyRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceVisitorSendnotifyRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type VisitorNotifyVo struct {
	Content    string `json:"content,omitempty"`
	DeviceId   int64  `json:"device_id,omitempty"`
	Feedback   string `json:"feedback,omitempty"`
	NotifyType int64  `json:"notify_type,omitempty"`
	UserName   string `json:"user_name,omitempty"`
	Userid     string `json:"userid,omitempty"`
}
type OapiSmartdeviceVisitorSendnotifyResponse struct {
	taobao.TaobaoResponse
	Result  bool `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}
