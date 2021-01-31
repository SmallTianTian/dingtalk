package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceAdvancedServiceBindRequest() *OapiAttendanceAdvancedServiceBindRequest {
	return &OapiAttendanceAdvancedServiceBindRequest{}
}

type OapiAttendanceAdvancedServiceBindRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceAdvancedServiceBindResponse
	OpUserid        string
	Param           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceAdvancedServiceBindRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceAdvancedServiceBindRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceAdvancedServiceBindRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiAttendanceAdvancedServiceBindRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiAttendanceAdvancedServiceBindRequest) SetParam(param2 string) {
	this.Param = param2
}
func (this *OapiAttendanceAdvancedServiceBindRequest) GetParam() string {
	return this.Param
}
func (this *OapiAttendanceAdvancedServiceBindRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.advanced.service.bind"
}
func (this *OapiAttendanceAdvancedServiceBindRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceAdvancedServiceBindRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceAdvancedServiceBindRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceAdvancedServiceBindRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceAdvancedServiceBindRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["op_userid"] = this.OpUserid
	txtParams["param"] = this.Param
	return txtParams
}
func (this *OapiAttendanceAdvancedServiceBindRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OpUserid, "opUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceAdvancedServiceBindRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceAdvancedServiceBindRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type BindParam struct {
	EndTime    int64  `json:"end_time,omitempty"`
	EntityId   string `json:"entity_id,omitempty"`
	EntityType string `json:"entity_type,omitempty"`
	Name       string `json:"name,omitempty"`
	ServiceId  int64  `json:"service_id,omitempty"`
	StartTime  int64  `json:"start_time,omitempty"`
}
type OapiAttendanceAdvancedServiceBindResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Success bool   `json:"success,omitempty"`
}
