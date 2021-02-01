package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiPbpInstanceCreateRequest() *OapiPbpInstanceCreateRequest {
	return &OapiPbpInstanceCreateRequest{}
}

type OapiPbpInstanceCreateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiPbpInstanceCreateResponse
	Active          bool
	BizId           string
	EndTime         int64
	OuterId         string
	StartTime       int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiPbpInstanceCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiPbpInstanceCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiPbpInstanceCreateRequest) SetActive(active2 bool) {
	this.Active = active2
}
func (this *OapiPbpInstanceCreateRequest) GetActive() bool {
	return this.Active
}
func (this *OapiPbpInstanceCreateRequest) SetBizId(bizId2 string) {
	this.BizId = bizId2
}
func (this *OapiPbpInstanceCreateRequest) GetBizId() string {
	return this.BizId
}
func (this *OapiPbpInstanceCreateRequest) SetEndTime(endTime2 int64) {
	this.EndTime = endTime2
}
func (this *OapiPbpInstanceCreateRequest) GetEndTime() int64 {
	return this.EndTime
}
func (this *OapiPbpInstanceCreateRequest) SetOuterId(outerId2 string) {
	this.OuterId = outerId2
}
func (this *OapiPbpInstanceCreateRequest) GetOuterId() string {
	return this.OuterId
}
func (this *OapiPbpInstanceCreateRequest) SetStartTime(startTime2 int64) {
	this.StartTime = startTime2
}
func (this *OapiPbpInstanceCreateRequest) GetStartTime() int64 {
	return this.StartTime
}
func (this *OapiPbpInstanceCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.pbp.instance.create"
}
func (this *OapiPbpInstanceCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiPbpInstanceCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiPbpInstanceCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiPbpInstanceCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiPbpInstanceCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["active"] = this.Active
	txtParams["biz_id"] = this.BizId
	txtParams["end_time"] = this.EndTime
	txtParams["outer_id"] = this.OuterId
	txtParams["start_time"] = this.StartTime
	return txtParams
}
func (this *OapiPbpInstanceCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizId, "bizId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiPbpInstanceCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiPbpInstanceCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiPbpInstanceCreateResponse struct {
	taobao.TaobaoResponse
	BizInstId string `json:"biz_inst_id,omitempty"`
}
