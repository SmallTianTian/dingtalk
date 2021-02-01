package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceGroupPositionsQueryRequest() *OapiAttendanceGroupPositionsQueryRequest {
	return &OapiAttendanceGroupPositionsQueryRequest{}
}

type OapiAttendanceGroupPositionsQueryRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceGroupPositionsQueryResponse
	Cursor          string
	GroupKey        string
	OpUserid        string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceGroupPositionsQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceGroupPositionsQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceGroupPositionsQueryRequest) SetCursor(cursor2 string) {
	this.Cursor = cursor2
}
func (this *OapiAttendanceGroupPositionsQueryRequest) GetCursor() string {
	return this.Cursor
}
func (this *OapiAttendanceGroupPositionsQueryRequest) SetGroupKey(groupKey2 string) {
	this.GroupKey = groupKey2
}
func (this *OapiAttendanceGroupPositionsQueryRequest) GetGroupKey() string {
	return this.GroupKey
}
func (this *OapiAttendanceGroupPositionsQueryRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiAttendanceGroupPositionsQueryRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiAttendanceGroupPositionsQueryRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiAttendanceGroupPositionsQueryRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiAttendanceGroupPositionsQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.group.positions.query"
}
func (this *OapiAttendanceGroupPositionsQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceGroupPositionsQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceGroupPositionsQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceGroupPositionsQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceGroupPositionsQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cursor"] = this.Cursor
	txtParams["group_key"] = this.GroupKey
	txtParams["op_userid"] = this.OpUserid
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiAttendanceGroupPositionsQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.GroupKey, "groupKey"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceGroupPositionsQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceGroupPositionsQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceGroupPositionsQueryResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
