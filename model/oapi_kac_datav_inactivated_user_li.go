package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiKacDatavInactivatedUserListRequest() *OapiKacDatavInactivatedUserListRequest {
	return &OapiKacDatavInactivatedUserListRequest{}
}

type OapiKacDatavInactivatedUserListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiKacDatavInactivatedUserListResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiKacDatavInactivatedUserListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiKacDatavInactivatedUserListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiKacDatavInactivatedUserListRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiKacDatavInactivatedUserListRequest) GetRequest() string {
	return this.Request
}
func (this *OapiKacDatavInactivatedUserListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.kac.datav.inactivated.user.list"
}
func (this *OapiKacDatavInactivatedUserListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiKacDatavInactivatedUserListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiKacDatavInactivatedUserListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiKacDatavInactivatedUserListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiKacDatavInactivatedUserListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiKacDatavInactivatedUserListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiKacDatavInactivatedUserListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiKacDatavInactivatedUserListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type UnactivatedUserRequest struct {
	Cursor int64  `json:"cursor,omitempty"`
	DataId string `json:"data_id,omitempty"`
	Size   int64  `json:"size,omitempty"`
}
type OapiKacDatavInactivatedUserListResponse struct {
	taobao.TaobaoResponse
	Result UnactivatedUserResponse `json:"result,omitempty"`
}
type UnactivatedUserVo struct {
	DeptId      int64  `json:"dept_id,omitempty"`
	DeptName    string `json:"dept_name,omitempty"`
	StaffJobNum string `json:"staff_job_num,omitempty"`
	StaffName   string `json:"staff_name,omitempty"`
	Userid      string `json:"userid,omitempty"`
}
type UnactivatedUserResponse struct {
	Data       []UnactivatedUserVo `json:"data,omitempty"`
	HasMore    bool                `json:"has_more,omitempty"`
	NextCursor int64               `json:"next_cursor,omitempty"`
}
