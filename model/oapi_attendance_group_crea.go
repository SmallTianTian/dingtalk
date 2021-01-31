package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAttendanceGroupCreateRequest() *OapiAttendanceGroupCreateRequest {
	return &OapiAttendanceGroupCreateRequest{}
}

type OapiAttendanceGroupCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceGroupCreateResponse
	Group           string
	OpUserid        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceGroupCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceGroupCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceGroupCreateRequest) SetGroup(group2 string) {
	this.Group = group2
}
func (this *OapiAttendanceGroupCreateRequest) GetGroup() string {
	return this.Group
}
func (this *OapiAttendanceGroupCreateRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiAttendanceGroupCreateRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiAttendanceGroupCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.group.create"
}
func (this *OapiAttendanceGroupCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceGroupCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceGroupCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceGroupCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceGroupCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["group"] = this.Group
	txtParams["op_userid"] = this.OpUserid
	return txtParams
}
func (this *OapiAttendanceGroupCreateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAttendanceGroupCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceGroupCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type Group struct {
	EnableCameraCheck bool   `json:"enable_camera_check,omitempty"`
	EnableFaceBeauty  bool   `json:"enable_face_beauty,omitempty"`
	EnableFaceCheck   bool   `json:"enable_face_check,omitempty"`
	Ext               string `json:"ext,omitempty"`
	LocationOffset    int64  `json:"location_offset,omitempty"`
	Name              string `json:"name,omitempty"`
}
type OapiAttendanceGroupCreateResponse struct {
	taobao.TaobaoResponse
	Result  Group `json:"result,omitempty"`
	Success bool  `json:"success,omitempty"`
}
