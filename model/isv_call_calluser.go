package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewIsvCallCalluserRequest() *IsvCallCalluserRequest {
	return &IsvCallCalluserRequest{}
}

type IsvCallCalluserRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            IsvCallCalluserResponse
	AuthedCorpId    string
	AuthedStaffId   string
	StaffId         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *IsvCallCalluserRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *IsvCallCalluserRequest) SetAuthedCorpId(authedCorpId2 string) {
	this.AuthedCorpId = authedCorpId2
}
func (this *IsvCallCalluserRequest) GetAuthedCorpId() string {
	return this.AuthedCorpId
}
func (this *IsvCallCalluserRequest) SetAuthedStaffId(authedStaffId2 string) {
	this.AuthedStaffId = authedStaffId2
}
func (this *IsvCallCalluserRequest) GetAuthedStaffId() string {
	return this.AuthedStaffId
}
func (this *IsvCallCalluserRequest) SetStaffId(staffId2 string) {
	this.StaffId = staffId2
}
func (this *IsvCallCalluserRequest) GetStaffId() string {
	return this.StaffId
}
func (this *IsvCallCalluserRequest) GetApiMethodName() string {
	return "dingtalk.isv.call.calluser"
}
func (this *IsvCallCalluserRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *IsvCallCalluserRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *IsvCallCalluserRequest) GetTopApiCallType() string {
	return "top"
}
func (this *IsvCallCalluserRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *IsvCallCalluserRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *IsvCallCalluserRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["authed_corp_id"] = this.AuthedCorpId
	txtParams["authed_staff_id"] = this.AuthedStaffId
	txtParams["staff_id"] = this.StaffId
	return txtParams
}
func (this *IsvCallCalluserRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AuthedCorpId, "authedCorpId"); err != nil {
		return err
	}
	return nil
}
func (this *IsvCallCalluserRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *IsvCallCalluserRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type IsvCallCalluserResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
