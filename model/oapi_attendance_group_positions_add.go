package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceGroupPositionsAddRequest() *OapiAttendanceGroupPositionsAddRequest {
	return &OapiAttendanceGroupPositionsAddRequest{}
}

type OapiAttendanceGroupPositionsAddRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceGroupPositionsAddResponse
	GroupKey        string
	OpUserid        string
	PositionList    string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceGroupPositionsAddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceGroupPositionsAddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceGroupPositionsAddRequest) SetGroupKey(groupKey2 string) {
	this.GroupKey = groupKey2
}
func (this *OapiAttendanceGroupPositionsAddRequest) GetGroupKey() string {
	return this.GroupKey
}
func (this *OapiAttendanceGroupPositionsAddRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiAttendanceGroupPositionsAddRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiAttendanceGroupPositionsAddRequest) SetPositionList(positionList2 string) {
	this.PositionList = positionList2
}
func (this *OapiAttendanceGroupPositionsAddRequest) GetPositionList() string {
	return this.PositionList
}
func (this *OapiAttendanceGroupPositionsAddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.group.positions.add"
}
func (this *OapiAttendanceGroupPositionsAddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceGroupPositionsAddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceGroupPositionsAddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceGroupPositionsAddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceGroupPositionsAddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["group_key"] = this.GroupKey
	txtParams["op_userid"] = this.OpUserid
	txtParams["position_list"] = this.PositionList
	return txtParams
}
func (this *OapiAttendanceGroupPositionsAddRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.GroupKey, "groupKey"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceGroupPositionsAddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceGroupPositionsAddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type Position struct {
	Address   string `json:"address,omitempty"`
	ForeignId string `json:"foreign_id,omitempty"`
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
}
type OapiAttendanceGroupPositionsAddResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
type ErrorInfo struct {
	Code        string     `json:"code,omitempty"`
	FailureList []Position `json:"failure_list,omitempty"`
	Msg         string     `json:"msg,omitempty"`
}
