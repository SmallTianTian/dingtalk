package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceTokenGetRequest() *OapiAttendanceTokenGetRequest {
	return &OapiAttendanceTokenGetRequest{}
}

type OapiAttendanceTokenGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceTokenGetResponse
	OpUserid        string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiAttendanceTokenGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceTokenGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceTokenGetRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiAttendanceTokenGetRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiAttendanceTokenGetRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiAttendanceTokenGetRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiAttendanceTokenGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.token.get"
}
func (this *OapiAttendanceTokenGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceTokenGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceTokenGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceTokenGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceTokenGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["op_userid"] = this.OpUserid
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiAttendanceTokenGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OpUserid, "opUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceTokenGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceTokenGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceTokenGetResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
type TokenVO struct {
	ExpireTime int64  `json:"expire_time,omitempty"`
	Token      string `json:"token,omitempty"`
}
