package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewSmartworkAttendsGetleaveapprovedurationRequest() *SmartworkAttendsGetleaveapprovedurationRequest {
	return &SmartworkAttendsGetleaveapprovedurationRequest{}
}

type SmartworkAttendsGetleaveapprovedurationRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            SmartworkAttendsGetleaveapprovedurationResponse
	FromDate        time.Time
	ToDate          time.Time
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *SmartworkAttendsGetleaveapprovedurationRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *SmartworkAttendsGetleaveapprovedurationRequest) SetFromDate(fromDate2 time.Time) {
	this.FromDate = fromDate2
}
func (this *SmartworkAttendsGetleaveapprovedurationRequest) GetFromDate() time.Time {
	return this.FromDate
}
func (this *SmartworkAttendsGetleaveapprovedurationRequest) SetToDate(toDate2 time.Time) {
	this.ToDate = toDate2
}
func (this *SmartworkAttendsGetleaveapprovedurationRequest) GetToDate() time.Time {
	return this.ToDate
}
func (this *SmartworkAttendsGetleaveapprovedurationRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *SmartworkAttendsGetleaveapprovedurationRequest) GetUserid() string {
	return this.Userid
}
func (this *SmartworkAttendsGetleaveapprovedurationRequest) GetApiMethodName() string {
	return "dingtalk.smartwork.attends.getleaveapproveduration"
}
func (this *SmartworkAttendsGetleaveapprovedurationRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *SmartworkAttendsGetleaveapprovedurationRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *SmartworkAttendsGetleaveapprovedurationRequest) GetTopApiCallType() string {
	return "top"
}
func (this *SmartworkAttendsGetleaveapprovedurationRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *SmartworkAttendsGetleaveapprovedurationRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *SmartworkAttendsGetleaveapprovedurationRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["from_date"] = this.FromDate
	txtParams["to_date"] = this.ToDate
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *SmartworkAttendsGetleaveapprovedurationRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.FromDate, "fromDate"); err != nil {
		return err
	}
	return nil
}
func (this *SmartworkAttendsGetleaveapprovedurationRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *SmartworkAttendsGetleaveapprovedurationRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type SmartworkAttendsGetleaveapprovedurationResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
