package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCcoserviceServicegroupIsignoreproblemcheckRequest() *OapiCcoserviceServicegroupIsignoreproblemcheckRequest {
	return &OapiCcoserviceServicegroupIsignoreproblemcheckRequest{}
}

type OapiCcoserviceServicegroupIsignoreproblemcheckRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp               OapiCcoserviceServicegroupIsignoreproblemcheckResponse
	DingtalkId         string
	OpenConversationId string
	TopHttpMethod      string
	TopResponseType    string
}

func (this *OapiCcoserviceServicegroupIsignoreproblemcheckRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCcoserviceServicegroupIsignoreproblemcheckRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCcoserviceServicegroupIsignoreproblemcheckRequest) SetDingtalkId(dingtalkId2 string) {
	this.DingtalkId = dingtalkId2
}
func (this *OapiCcoserviceServicegroupIsignoreproblemcheckRequest) GetDingtalkId() string {
	return this.DingtalkId
}
func (this *OapiCcoserviceServicegroupIsignoreproblemcheckRequest) SetOpenConversationId(openConversationId2 string) {
	this.OpenConversationId = openConversationId2
}
func (this *OapiCcoserviceServicegroupIsignoreproblemcheckRequest) GetOpenConversationId() string {
	return this.OpenConversationId
}
func (this *OapiCcoserviceServicegroupIsignoreproblemcheckRequest) GetApiMethodName() string {
	return "dingtalk.oapi.ccoservice.servicegroup.isignoreproblemcheck"
}
func (this *OapiCcoserviceServicegroupIsignoreproblemcheckRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCcoserviceServicegroupIsignoreproblemcheckRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCcoserviceServicegroupIsignoreproblemcheckRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCcoserviceServicegroupIsignoreproblemcheckRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCcoserviceServicegroupIsignoreproblemcheckRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["dingtalk_id"] = this.DingtalkId
	txtParams["open_conversation_id"] = this.OpenConversationId
	return txtParams
}
func (this *OapiCcoserviceServicegroupIsignoreproblemcheckRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.DingtalkId, "dingtalkId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCcoserviceServicegroupIsignoreproblemcheckRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCcoserviceServicegroupIsignoreproblemcheckRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCcoserviceServicegroupIsignoreproblemcheckResponse struct {
	taobao.TaobaoResponse
	Result  bool `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}
