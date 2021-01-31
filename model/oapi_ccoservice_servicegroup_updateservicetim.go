package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCcoserviceServicegroupUpdateservicetimeRequest() *OapiCcoserviceServicegroupUpdateservicetimeRequest {
	return &OapiCcoserviceServicegroupUpdateservicetimeRequest{}
}

type OapiCcoserviceServicegroupUpdateservicetimeRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp               OapiCcoserviceServicegroupUpdateservicetimeResponse
	EndTime            string
	OpenConversationId string
	StartTime          string
	TimeType           int64
	TopHttpMethod      string
	TopResponseType    string
}

func (this *OapiCcoserviceServicegroupUpdateservicetimeRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCcoserviceServicegroupUpdateservicetimeRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCcoserviceServicegroupUpdateservicetimeRequest) SetEndTime(endTime2 string) {
	this.EndTime = endTime2
}
func (this *OapiCcoserviceServicegroupUpdateservicetimeRequest) GetEndTime() string {
	return this.EndTime
}
func (this *OapiCcoserviceServicegroupUpdateservicetimeRequest) SetOpenConversationId(openConversationId2 string) {
	this.OpenConversationId = openConversationId2
}
func (this *OapiCcoserviceServicegroupUpdateservicetimeRequest) GetOpenConversationId() string {
	return this.OpenConversationId
}
func (this *OapiCcoserviceServicegroupUpdateservicetimeRequest) SetStartTime(startTime2 string) {
	this.StartTime = startTime2
}
func (this *OapiCcoserviceServicegroupUpdateservicetimeRequest) GetStartTime() string {
	return this.StartTime
}
func (this *OapiCcoserviceServicegroupUpdateservicetimeRequest) SetTimeType(timeType2 int64) {
	this.TimeType = timeType2
}
func (this *OapiCcoserviceServicegroupUpdateservicetimeRequest) GetTimeType() int64 {
	return this.TimeType
}
func (this *OapiCcoserviceServicegroupUpdateservicetimeRequest) GetApiMethodName() string {
	return "dingtalk.oapi.ccoservice.servicegroup.updateservicetime"
}
func (this *OapiCcoserviceServicegroupUpdateservicetimeRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCcoserviceServicegroupUpdateservicetimeRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCcoserviceServicegroupUpdateservicetimeRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCcoserviceServicegroupUpdateservicetimeRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCcoserviceServicegroupUpdateservicetimeRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["end_time"] = this.EndTime
	txtParams["open_conversation_id"] = this.OpenConversationId
	txtParams["start_time"] = this.StartTime
	txtParams["time_type"] = this.TimeType
	return txtParams
}
func (this *OapiCcoserviceServicegroupUpdateservicetimeRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.EndTime, "endTime"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCcoserviceServicegroupUpdateservicetimeRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCcoserviceServicegroupUpdateservicetimeRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCcoserviceServicegroupUpdateservicetimeResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
