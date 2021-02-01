package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduPeriodListRequest() *OapiEduPeriodListRequest {
	return &OapiEduPeriodListRequest{}
}

type OapiEduPeriodListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduPeriodListResponse
	CampusId        int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduPeriodListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduPeriodListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduPeriodListRequest) SetCampusId(campusId2 int64) {
	this.CampusId = campusId2
}
func (this *OapiEduPeriodListRequest) GetCampusId() int64 {
	return this.CampusId
}
func (this *OapiEduPeriodListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.period.list"
}
func (this *OapiEduPeriodListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduPeriodListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduPeriodListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduPeriodListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduPeriodListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["campus_id"] = this.CampusId
	return txtParams
}
func (this *OapiEduPeriodListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CampusId, "campusId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduPeriodListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduPeriodListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduPeriodListResponse struct {
	taobao.TaobaoResponse
	Result  []PeriodResponse `json:"result,omitempty"`
	Success bool             `json:"success,omitempty"`
}
