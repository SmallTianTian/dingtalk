package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceGroupWifisRemoveRequest() *OapiAttendanceGroupWifisRemoveRequest {
	return &OapiAttendanceGroupWifisRemoveRequest{}
}

type OapiAttendanceGroupWifisRemoveRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceGroupWifisRemoveResponse
	GroupKey        string
	OpUserid        string
	TopHttpMethod   string
	TopResponseType string
	WifiKeyList     string
}

func (this *OapiAttendanceGroupWifisRemoveRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceGroupWifisRemoveRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceGroupWifisRemoveRequest) SetGroupKey(groupKey2 string) {
	this.GroupKey = groupKey2
}
func (this *OapiAttendanceGroupWifisRemoveRequest) GetGroupKey() string {
	return this.GroupKey
}
func (this *OapiAttendanceGroupWifisRemoveRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiAttendanceGroupWifisRemoveRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiAttendanceGroupWifisRemoveRequest) SetWifiKeyList(wifiKeyList2 string) {
	this.WifiKeyList = wifiKeyList2
}
func (this *OapiAttendanceGroupWifisRemoveRequest) GetWifiKeyList() string {
	return this.WifiKeyList
}
func (this *OapiAttendanceGroupWifisRemoveRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.group.wifis.remove"
}
func (this *OapiAttendanceGroupWifisRemoveRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceGroupWifisRemoveRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceGroupWifisRemoveRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceGroupWifisRemoveRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceGroupWifisRemoveRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["group_key"] = this.GroupKey
	txtParams["op_userid"] = this.OpUserid
	txtParams["wifi_key_list"] = this.WifiKeyList
	return txtParams
}
func (this *OapiAttendanceGroupWifisRemoveRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.GroupKey, "groupKey"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceGroupWifisRemoveRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceGroupWifisRemoveRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceGroupWifisRemoveResponse struct {
	taobao.TaobaoResponse
	Result  Result `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
