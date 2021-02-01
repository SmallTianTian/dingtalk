package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRoleAddrolegroupRequest() *OapiRoleAddrolegroupRequest {
	return &OapiRoleAddrolegroupRequest{}
}

type OapiRoleAddrolegroupRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRoleAddrolegroupResponse
	Name            string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRoleAddrolegroupRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRoleAddrolegroupRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRoleAddrolegroupRequest) SetName(name2 string) {
	this.Name = name2
}
func (this *OapiRoleAddrolegroupRequest) GetName() string {
	return this.Name
}
func (this *OapiRoleAddrolegroupRequest) GetApiMethodName() string {
	return "dingtalk.oapi.role.addrolegroup"
}
func (this *OapiRoleAddrolegroupRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRoleAddrolegroupRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRoleAddrolegroupRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRoleAddrolegroupRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRoleAddrolegroupRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["name"] = this.Name
	return txtParams
}
func (this *OapiRoleAddrolegroupRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Name, "name"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRoleAddrolegroupRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRoleAddrolegroupRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRoleAddrolegroupResponse struct {
	taobao.TaobaoResponse
	GroupId int64 `json:"groupId,omitempty"`
}
