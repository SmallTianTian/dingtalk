package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiSmartdeviceAtmachineGetByUseridRequest() *OapiSmartdeviceAtmachineGetByUseridRequest {
	return &OapiSmartdeviceAtmachineGetByUseridRequest{}
}

type OapiSmartdeviceAtmachineGetByUseridRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceAtmachineGetByUseridResponse
	Param           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartdeviceAtmachineGetByUseridRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceAtmachineGetByUseridRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceAtmachineGetByUseridRequest) SetParam(param2 string) {
	this.Param = param2
}
func (this *OapiSmartdeviceAtmachineGetByUseridRequest) GetParam() string {
	return this.Param
}
func (this *OapiSmartdeviceAtmachineGetByUseridRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.atmachine.get_by_userid"
}
func (this *OapiSmartdeviceAtmachineGetByUseridRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceAtmachineGetByUseridRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceAtmachineGetByUseridRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceAtmachineGetByUseridRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceAtmachineGetByUseridRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["param"] = this.Param
	return txtParams
}
func (this *OapiSmartdeviceAtmachineGetByUseridRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiSmartdeviceAtmachineGetByUseridRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceAtmachineGetByUseridRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type UserMachineInfoRequestVo struct {
	Offset int64  `json:"offset,omitempty"`
	Size   int64  `json:"size,omitempty"`
	Userid string `json:"userid,omitempty"`
}
type OapiSmartdeviceAtmachineGetByUseridResponse struct {
	taobao.TaobaoResponse
	Result MachineInfoResultVo `json:"result,omitempty"`
}
