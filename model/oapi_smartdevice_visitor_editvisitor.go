package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartdeviceVisitorEditvisitorRequest() *OapiSmartdeviceVisitorEditvisitorRequest {
	return &OapiSmartdeviceVisitorEditvisitorRequest{}
}

type OapiSmartdeviceVisitorEditvisitorRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceVisitorEditvisitorResponse
	ReservationId   string
	TopHttpMethod   string
	TopResponseType string
	VisitorVo       string
}

func (this *OapiSmartdeviceVisitorEditvisitorRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceVisitorEditvisitorRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceVisitorEditvisitorRequest) SetReservationId(reservationId2 string) {
	this.ReservationId = reservationId2
}
func (this *OapiSmartdeviceVisitorEditvisitorRequest) GetReservationId() string {
	return this.ReservationId
}
func (this *OapiSmartdeviceVisitorEditvisitorRequest) SetVisitorVo(visitorVo2 string) {
	this.VisitorVo = visitorVo2
}
func (this *OapiSmartdeviceVisitorEditvisitorRequest) GetVisitorVo() string {
	return this.VisitorVo
}
func (this *OapiSmartdeviceVisitorEditvisitorRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.visitor.editvisitor"
}
func (this *OapiSmartdeviceVisitorEditvisitorRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceVisitorEditvisitorRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceVisitorEditvisitorRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceVisitorEditvisitorRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceVisitorEditvisitorRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["reservation_id"] = this.ReservationId
	txtParams["visitor_vo"] = this.VisitorVo
	return txtParams
}
func (this *OapiSmartdeviceVisitorEditvisitorRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ReservationId, "reservationId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartdeviceVisitorEditvisitorRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceVisitorEditvisitorRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartdeviceVisitorEditvisitorResponse struct {
	taobao.TaobaoResponse
	Result bool `json:"result,omitempty"`
}
