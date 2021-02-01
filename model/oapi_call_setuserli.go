package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCallSetuserlistRequest() *OapiCallSetuserlistRequest {
	return &OapiCallSetuserlistRequest{}
}

type OapiCallSetuserlistRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCallSetuserlistResponse
	StaffIdList     string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCallSetuserlistRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCallSetuserlistRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCallSetuserlistRequest) SetStaffIdList(staffIdList2 string) {
	this.StaffIdList = staffIdList2
}
func (this *OapiCallSetuserlistRequest) GetStaffIdList() string {
	return this.StaffIdList
}
func (this *OapiCallSetuserlistRequest) GetApiMethodName() string {
	return "dingtalk.oapi.call.setuserlist"
}
func (this *OapiCallSetuserlistRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCallSetuserlistRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCallSetuserlistRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCallSetuserlistRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCallSetuserlistRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["staff_id_list"] = this.StaffIdList
	return txtParams
}
func (this *OapiCallSetuserlistRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.StaffIdList, "staffIdList"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCallSetuserlistRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCallSetuserlistRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCallSetuserlistResponse struct {
	taobao.TaobaoResponse
}
