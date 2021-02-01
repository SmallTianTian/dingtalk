package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceGroupWifisQueryRequest() *OapiAttendanceGroupWifisQueryRequest {
	return &OapiAttendanceGroupWifisQueryRequest{}
}

type OapiAttendanceGroupWifisQueryRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceGroupWifisQueryResponse
	Cursor          string
	GroupKey        string
	OpUserid        string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceGroupWifisQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceGroupWifisQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceGroupWifisQueryRequest) SetCursor(cursor2 string) {
	this.Cursor = cursor2
}
func (this *OapiAttendanceGroupWifisQueryRequest) GetCursor() string {
	return this.Cursor
}
func (this *OapiAttendanceGroupWifisQueryRequest) SetGroupKey(groupKey2 string) {
	this.GroupKey = groupKey2
}
func (this *OapiAttendanceGroupWifisQueryRequest) GetGroupKey() string {
	return this.GroupKey
}
func (this *OapiAttendanceGroupWifisQueryRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiAttendanceGroupWifisQueryRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiAttendanceGroupWifisQueryRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiAttendanceGroupWifisQueryRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiAttendanceGroupWifisQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.group.wifis.query"
}
func (this *OapiAttendanceGroupWifisQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceGroupWifisQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceGroupWifisQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceGroupWifisQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceGroupWifisQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cursor"] = this.Cursor
	txtParams["group_key"] = this.GroupKey
	txtParams["op_userid"] = this.OpUserid
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiAttendanceGroupWifisQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.GroupKey, "groupKey"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceGroupWifisQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceGroupWifisQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceGroupWifisQueryResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
