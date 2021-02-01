package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCallRemoveuserlistRequest() *OapiCallRemoveuserlistRequest {
	return &OapiCallRemoveuserlistRequest{}
}

type OapiCallRemoveuserlistRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCallRemoveuserlistResponse
	StaffIdList     string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCallRemoveuserlistRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCallRemoveuserlistRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCallRemoveuserlistRequest) SetStaffIdList(staffIdList2 string) {
	this.StaffIdList = staffIdList2
}
func (this *OapiCallRemoveuserlistRequest) GetStaffIdList() string {
	return this.StaffIdList
}
func (this *OapiCallRemoveuserlistRequest) GetApiMethodName() string {
	return "dingtalk.oapi.call.removeuserlist"
}
func (this *OapiCallRemoveuserlistRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCallRemoveuserlistRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCallRemoveuserlistRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCallRemoveuserlistRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCallRemoveuserlistRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["staff_id_list"] = this.StaffIdList
	return txtParams
}
func (this *OapiCallRemoveuserlistRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.StaffIdList, "staffIdList"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCallRemoveuserlistRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCallRemoveuserlistRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCallRemoveuserlistResponse struct {
	taobao.TaobaoResponse
}
