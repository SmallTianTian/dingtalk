package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceGroupPositionsRemoveRequest() *OapiAttendanceGroupPositionsRemoveRequest {
	return &OapiAttendanceGroupPositionsRemoveRequest{}
}

type OapiAttendanceGroupPositionsRemoveRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceGroupPositionsRemoveResponse
	GroupKey        string
	OpUserid        string
	TopHttpMethod   string
	TopResponseType string
	WifiKeyList     string
}

func (this *OapiAttendanceGroupPositionsRemoveRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceGroupPositionsRemoveRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceGroupPositionsRemoveRequest) SetGroupKey(groupKey2 string) {
	this.GroupKey = groupKey2
}
func (this *OapiAttendanceGroupPositionsRemoveRequest) GetGroupKey() string {
	return this.GroupKey
}
func (this *OapiAttendanceGroupPositionsRemoveRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiAttendanceGroupPositionsRemoveRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiAttendanceGroupPositionsRemoveRequest) SetWifiKeyList(wifiKeyList2 string) {
	this.WifiKeyList = wifiKeyList2
}
func (this *OapiAttendanceGroupPositionsRemoveRequest) GetWifiKeyList() string {
	return this.WifiKeyList
}
func (this *OapiAttendanceGroupPositionsRemoveRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.group.positions.remove"
}
func (this *OapiAttendanceGroupPositionsRemoveRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceGroupPositionsRemoveRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceGroupPositionsRemoveRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceGroupPositionsRemoveRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceGroupPositionsRemoveRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["group_key"] = this.GroupKey
	txtParams["op_userid"] = this.OpUserid
	txtParams["wifi_key_list"] = this.WifiKeyList
	return txtParams
}
func (this *OapiAttendanceGroupPositionsRemoveRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.GroupKey, "groupKey"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceGroupPositionsRemoveRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceGroupPositionsRemoveRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceGroupPositionsRemoveResponse struct {
	taobao.TaobaoResponse
	Result  Result `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
