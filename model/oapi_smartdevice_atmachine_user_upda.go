package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiSmartdeviceAtmachineUserUpdateRequest() *OapiSmartdeviceAtmachineUserUpdateRequest {
	return &OapiSmartdeviceAtmachineUserUpdateRequest{}
}

type OapiSmartdeviceAtmachineUserUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceAtmachineUserUpdateResponse
	Param           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartdeviceAtmachineUserUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceAtmachineUserUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceAtmachineUserUpdateRequest) SetParam(param2 string) {
	this.Param = param2
}
func (this *OapiSmartdeviceAtmachineUserUpdateRequest) GetParam() string {
	return this.Param
}
func (this *OapiSmartdeviceAtmachineUserUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.atmachine.user.update"
}
func (this *OapiSmartdeviceAtmachineUserUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceAtmachineUserUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceAtmachineUserUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceAtmachineUserUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceAtmachineUserUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["param"] = this.Param
	return txtParams
}
func (this *OapiSmartdeviceAtmachineUserUpdateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiSmartdeviceAtmachineUserUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceAtmachineUserUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type MachineUsersUpdateRequestVo struct {
	AddUseridList []string `json:"add_userid_list,omitempty"`
	DelUseridList []string `json:"del_userid_list,omitempty"`
	DeviceidList  []string `json:"deviceid_list,omitempty"`
	DevidList     []int64  `json:"devid_list,omitempty"`
}
type OapiSmartdeviceAtmachineUserUpdateResponse struct {
	taobao.TaobaoResponse
}
