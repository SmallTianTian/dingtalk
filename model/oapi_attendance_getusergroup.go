package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceGetusergroupRequest() *OapiAttendanceGetusergroupRequest {
	return &OapiAttendanceGetusergroupRequest{}
}

type OapiAttendanceGetusergroupRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceGetusergroupResponse
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiAttendanceGetusergroupRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceGetusergroupRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceGetusergroupRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiAttendanceGetusergroupRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiAttendanceGetusergroupRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.getusergroup"
}
func (this *OapiAttendanceGetusergroupRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceGetusergroupRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceGetusergroupRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceGetusergroupRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceGetusergroupRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiAttendanceGetusergroupRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Userid, taobao.DATA_OUTGOING_USER_ID); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceGetusergroupRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceGetusergroupRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceGetusergroupResponse struct {
	taobao.TaobaoResponse
	Result AtGroupFullForTopVo `json:"result,omitempty"`
}
type AtGroupFullForTopVo struct {
	Classes []AtClassVo `json:"classes,omitempty"`
	GroupId int64       `json:"group_id,omitempty"`
	Name    string      `json:"name,omitempty"`
	Type    string      `json:"type,omitempty"`
}
