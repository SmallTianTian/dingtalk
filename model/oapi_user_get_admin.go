package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiUserGetAdminRequest() *OapiUserGetAdminRequest {
	return &OapiUserGetAdminRequest{}
}

type OapiUserGetAdminRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiUserGetAdminResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiUserGetAdminRequest) GetTopHttpMethod() string {
	return "GET"
}
func (this *OapiUserGetAdminRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiUserGetAdminRequest) GetApiMethodName() string {
	return "dingtalk.oapi.user.get_admin"
}
func (this *OapiUserGetAdminRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiUserGetAdminRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiUserGetAdminRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiUserGetAdminRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiUserGetAdminRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *OapiUserGetAdminRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiUserGetAdminRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiUserGetAdminRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiUserGetAdminResponse struct {
	taobao.TaobaoResponse
	AdminList []AdminList `json:"admin_list,omitempty"`
	Errcode   int64       `json:"errcode,omitempty"`
	Errmsg    string      `json:"errmsg,omitempty"`
}
type AdminList struct {
	AdminMobile string `json:"admin_mobile,omitempty"`
	SysLevel    int64  `json:"sys_level,omitempty"`
	Userid      string `json:"userid,omitempty"`
}
