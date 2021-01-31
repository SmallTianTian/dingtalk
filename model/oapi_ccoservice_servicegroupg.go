package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCcoserviceServicegroupGetRequest() *OapiCcoserviceServicegroupGetRequest {
	return &OapiCcoserviceServicegroupGetRequest{}
}

type OapiCcoserviceServicegroupGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCcoserviceServicegroupGetResponse
	OpenGroupId     string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCcoserviceServicegroupGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCcoserviceServicegroupGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCcoserviceServicegroupGetRequest) SetOpenGroupId(openGroupId2 string) {
	this.OpenGroupId = openGroupId2
}
func (this *OapiCcoserviceServicegroupGetRequest) GetOpenGroupId() string {
	return this.OpenGroupId
}
func (this *OapiCcoserviceServicegroupGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.ccoservice.servicegroup.get"
}
func (this *OapiCcoserviceServicegroupGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCcoserviceServicegroupGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCcoserviceServicegroupGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCcoserviceServicegroupGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCcoserviceServicegroupGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["open_group_id"] = this.OpenGroupId
	return txtParams
}
func (this *OapiCcoserviceServicegroupGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OpenGroupId, "openGroupId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCcoserviceServicegroupGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCcoserviceServicegroupGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCcoserviceServicegroupGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64                    `json:"errcode,omitempty"`
	Errmsg  string                   `json:"errmsg,omitempty"`
	Result  ServiceConversationModel `json:"result,omitempty"`
}
