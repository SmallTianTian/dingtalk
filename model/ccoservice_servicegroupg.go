package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCcoserviceServicegroupGetRequest() *CcoserviceServicegroupGetRequest {
	return &CcoserviceServicegroupGetRequest{}
}

type CcoserviceServicegroupGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            CcoserviceServicegroupGetResponse
	OpenGroupId     string
	TopHttpMethod   string
	TopResponseType string
}

func (this *CcoserviceServicegroupGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CcoserviceServicegroupGetRequest) SetOpenGroupId(openGroupId2 string) {
	this.OpenGroupId = openGroupId2
}
func (this *CcoserviceServicegroupGetRequest) GetOpenGroupId() string {
	return this.OpenGroupId
}
func (this *CcoserviceServicegroupGetRequest) GetApiMethodName() string {
	return "dingtalk.ccoservice.servicegroup.get"
}
func (this *CcoserviceServicegroupGetRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CcoserviceServicegroupGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CcoserviceServicegroupGetRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CcoserviceServicegroupGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CcoserviceServicegroupGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CcoserviceServicegroupGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["open_group_id"] = this.OpenGroupId
	return txtParams
}
func (this *CcoserviceServicegroupGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OpenGroupId, "openGroupId"); err != nil {
		return err
	}
	return nil
}
func (this *CcoserviceServicegroupGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CcoserviceServicegroupGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CcoserviceServicegroupGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64                    `json:"errcode,omitempty"`
	Errmsg  string                   `json:"errmsg,omitempty"`
	Result  ServiceConversationModel `json:"result,omitempty"`
}
type ServiceConversationModel struct {
	ConversationType int64  `json:"conversation_type,omitempty"`
	OwnerDingtalkId  string `json:"owner_dingtalk_id,omitempty"`
	OwnerName        string `json:"owner_name,omitempty"`
	OwnerNick        string `json:"owner_nick,omitempty"`
	OwnerUserid      string `json:"owner_userid,omitempty"`
	Title            string `json:"title,omitempty"`
	Type             int64  `json:"type,omitempty"`
}
