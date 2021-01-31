package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceApproveCancelRequest() *OapiAttendanceApproveCancelRequest {
	return &OapiAttendanceApproveCancelRequest{}
}

type OapiAttendanceApproveCancelRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceApproveCancelResponse
	ApproveId       string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiAttendanceApproveCancelRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceApproveCancelRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceApproveCancelRequest) SetApproveId(approveId2 string) {
	this.ApproveId = approveId2
}
func (this *OapiAttendanceApproveCancelRequest) GetApproveId() string {
	return this.ApproveId
}
func (this *OapiAttendanceApproveCancelRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiAttendanceApproveCancelRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiAttendanceApproveCancelRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.approve.cancel"
}
func (this *OapiAttendanceApproveCancelRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceApproveCancelRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceApproveCancelRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceApproveCancelRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceApproveCancelRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["approve_id"] = this.ApproveId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiAttendanceApproveCancelRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckMaxLength(this.ApproveId, 100, "approveId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceApproveCancelRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceApproveCancelRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceApproveCancelResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
