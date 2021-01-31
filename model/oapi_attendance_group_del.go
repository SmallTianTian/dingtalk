package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceGroupDeleteRequest() *OapiAttendanceGroupDeleteRequest {
	return &OapiAttendanceGroupDeleteRequest{}
}

type OapiAttendanceGroupDeleteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceGroupDeleteResponse
	GroupKey        string
	OpUserid        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceGroupDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceGroupDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceGroupDeleteRequest) SetGroupKey(groupKey2 string) {
	this.GroupKey = groupKey2
}
func (this *OapiAttendanceGroupDeleteRequest) GetGroupKey() string {
	return this.GroupKey
}
func (this *OapiAttendanceGroupDeleteRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiAttendanceGroupDeleteRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiAttendanceGroupDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.group.delete"
}
func (this *OapiAttendanceGroupDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceGroupDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceGroupDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceGroupDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceGroupDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["group_key"] = this.GroupKey
	txtParams["op_userid"] = this.OpUserid
	return txtParams
}
func (this *OapiAttendanceGroupDeleteRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.GroupKey, "groupKey"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceGroupDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceGroupDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceGroupDeleteResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
