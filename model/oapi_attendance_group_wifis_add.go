package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceGroupWifisAddRequest() *OapiAttendanceGroupWifisAddRequest {
	return &OapiAttendanceGroupWifisAddRequest{}
}

type OapiAttendanceGroupWifisAddRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceGroupWifisAddResponse
	GroupKey        string
	OpUserid        string
	TopHttpMethod   string
	TopResponseType string
	WifiList        string
}

func (this *OapiAttendanceGroupWifisAddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceGroupWifisAddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceGroupWifisAddRequest) SetGroupKey(groupKey2 string) {
	this.GroupKey = groupKey2
}
func (this *OapiAttendanceGroupWifisAddRequest) GetGroupKey() string {
	return this.GroupKey
}
func (this *OapiAttendanceGroupWifisAddRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiAttendanceGroupWifisAddRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiAttendanceGroupWifisAddRequest) SetWifiList(wifiList2 string) {
	this.WifiList = wifiList2
}
func (this *OapiAttendanceGroupWifisAddRequest) GetWifiList() string {
	return this.WifiList
}
func (this *OapiAttendanceGroupWifisAddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.group.wifis.add"
}
func (this *OapiAttendanceGroupWifisAddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceGroupWifisAddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceGroupWifisAddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceGroupWifisAddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceGroupWifisAddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["group_key"] = this.GroupKey
	txtParams["op_userid"] = this.OpUserid
	txtParams["wifi_list"] = this.WifiList
	return txtParams
}
func (this *OapiAttendanceGroupWifisAddRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.GroupKey, "groupKey"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceGroupWifisAddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceGroupWifisAddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type Wifi struct {
	ForeignId string `json:"foreign_id,omitempty"`
	MacAddr   string `json:"mac_addr,omitempty"`
	Ssid      string `json:"ssid,omitempty"`
}
type OapiAttendanceGroupWifisAddResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
