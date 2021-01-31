package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCrmMenuGetRequest() *OapiCrmMenuGetRequest {
	return &OapiCrmMenuGetRequest{}
}

type OapiCrmMenuGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCrmMenuGetResponse
	ClientType      string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCrmMenuGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCrmMenuGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCrmMenuGetRequest) SetClientType(clientType2 string) {
	this.ClientType = clientType2
}
func (this *OapiCrmMenuGetRequest) GetClientType() string {
	return this.ClientType
}
func (this *OapiCrmMenuGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.crm.menu.get"
}
func (this *OapiCrmMenuGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCrmMenuGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCrmMenuGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCrmMenuGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCrmMenuGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["client_type"] = this.ClientType
	return txtParams
}
func (this *OapiCrmMenuGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ClientType, "clientType"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCrmMenuGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCrmMenuGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCrmMenuGetResponse struct {
	taobao.TaobaoResponse
	Result  OnlineNavigationModel `json:"result,omitempty"`
	Success bool                  `json:"success,omitempty"`
}
type OnlineNavigationModel struct {
	FormUuid     string `json:"form_uuid,omitempty"`
	Icon         string `json:"icon,omitempty"`
	IsNew        int64  `json:"is_new,omitempty"`
	LinkUrl      string `json:"link_url,omitempty"`
	ListOrder    int64  `json:"list_order,omitempty"`
	MobileHidden int64  `json:"mobile_hidden,omitempty"`
	Name         string `json:"name,omitempty"`
	NavType      string `json:"nav_type,omitempty"`
	ParentId     int64  `json:"parent_id,omitempty"`
	PcHidden     int64  `json:"pc_hidden,omitempty"`
}
