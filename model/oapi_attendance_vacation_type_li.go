package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceVacationTypeListRequest() *OapiAttendanceVacationTypeListRequest {
	return &OapiAttendanceVacationTypeListRequest{}
}

type OapiAttendanceVacationTypeListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceVacationTypeListResponse
	OpUserid        string
	TopHttpMethod   string
	TopResponseType string
	VacationSource  string
}

func (this *OapiAttendanceVacationTypeListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceVacationTypeListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceVacationTypeListRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiAttendanceVacationTypeListRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiAttendanceVacationTypeListRequest) SetVacationSource(vacationSource2 string) {
	this.VacationSource = vacationSource2
}
func (this *OapiAttendanceVacationTypeListRequest) GetVacationSource() string {
	return this.VacationSource
}
func (this *OapiAttendanceVacationTypeListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.vacation.type.list"
}
func (this *OapiAttendanceVacationTypeListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceVacationTypeListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceVacationTypeListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceVacationTypeListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceVacationTypeListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["op_userid"] = this.OpUserid
	txtParams["vacation_source"] = this.VacationSource
	return txtParams
}
func (this *OapiAttendanceVacationTypeListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OpUserid, "opUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceVacationTypeListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceVacationTypeListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceVacationTypeListResponse struct {
	taobao.TaobaoResponse
	Result  []Result `json:"result,omitempty"`
	Success bool     `json:"success,omitempty"`
}
