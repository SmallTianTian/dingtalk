package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceCorpConfirmRequest() *OapiAttendanceCorpConfirmRequest {
	return &OapiAttendanceCorpConfirmRequest{}
}

type OapiAttendanceCorpConfirmRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceCorpConfirmResponse
	CorpId          string
	CorpList        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceCorpConfirmRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceCorpConfirmRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceCorpConfirmRequest) SetCorpId(corpId2 string) {
	this.CorpId = corpId2
}
func (this *OapiAttendanceCorpConfirmRequest) GetCorpId() string {
	return this.CorpId
}
func (this *OapiAttendanceCorpConfirmRequest) SetCorpList(corpList2 string) {
	this.CorpList = corpList2
}
func (this *OapiAttendanceCorpConfirmRequest) GetCorpList() string {
	return this.CorpList
}
func (this *OapiAttendanceCorpConfirmRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.corp.confirm"
}
func (this *OapiAttendanceCorpConfirmRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceCorpConfirmRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceCorpConfirmRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceCorpConfirmRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceCorpConfirmRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["corp_id"] = this.CorpId
	txtParams["corp_list"] = this.CorpList
	return txtParams
}
func (this *OapiAttendanceCorpConfirmRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CorpId, "corpId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceCorpConfirmRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceCorpConfirmRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type Corp struct {
	AdminName   string `json:"admin_name,omitempty"`
	AdminPhone  string `json:"admin_phone,omitempty"`
	StaffAmount int64  `json:"staff_amount,omitempty"`
}
type OapiAttendanceCorpConfirmResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  Result `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
