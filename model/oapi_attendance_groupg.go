package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceGroupGetRequest() *OapiAttendanceGroupGetRequest {
	return &OapiAttendanceGroupGetRequest{}
}

type OapiAttendanceGroupGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceGroupGetResponse
	GroupKey        string
	OpUserid        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceGroupGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceGroupGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceGroupGetRequest) SetGroupKey(groupKey2 string) {
	this.GroupKey = groupKey2
}
func (this *OapiAttendanceGroupGetRequest) GetGroupKey() string {
	return this.GroupKey
}
func (this *OapiAttendanceGroupGetRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiAttendanceGroupGetRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiAttendanceGroupGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.group.get"
}
func (this *OapiAttendanceGroupGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceGroupGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceGroupGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceGroupGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceGroupGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["group_key"] = this.GroupKey
	txtParams["op_userid"] = this.OpUserid
	return txtParams
}
func (this *OapiAttendanceGroupGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.GroupKey, "groupKey"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceGroupGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceGroupGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceGroupGetResponse struct {
	taobao.TaobaoResponse
	Result  Group `json:"result,omitempty"`
	Success bool  `json:"success,omitempty"`
}
