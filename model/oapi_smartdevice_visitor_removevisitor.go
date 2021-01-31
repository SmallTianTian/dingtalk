package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartdeviceVisitorRemovevisitorRequest() *OapiSmartdeviceVisitorRemovevisitorRequest {
	return &OapiSmartdeviceVisitorRemovevisitorRequest{}
}

type OapiSmartdeviceVisitorRemovevisitorRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceVisitorRemovevisitorResponse
	ReservationId   string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartdeviceVisitorRemovevisitorRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceVisitorRemovevisitorRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceVisitorRemovevisitorRequest) SetReservationId(reservationId2 string) {
	this.ReservationId = reservationId2
}
func (this *OapiSmartdeviceVisitorRemovevisitorRequest) GetReservationId() string {
	return this.ReservationId
}
func (this *OapiSmartdeviceVisitorRemovevisitorRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.visitor.removevisitor"
}
func (this *OapiSmartdeviceVisitorRemovevisitorRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceVisitorRemovevisitorRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceVisitorRemovevisitorRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceVisitorRemovevisitorRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceVisitorRemovevisitorRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["reservation_id"] = this.ReservationId
	return txtParams
}
func (this *OapiSmartdeviceVisitorRemovevisitorRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ReservationId, "reservationId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartdeviceVisitorRemovevisitorRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceVisitorRemovevisitorRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartdeviceVisitorRemovevisitorResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  bool   `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
