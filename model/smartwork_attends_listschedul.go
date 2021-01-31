package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
	"time"
)

func NewSmartworkAttendsListscheduleRequest() *SmartworkAttendsListscheduleRequest {
	return &SmartworkAttendsListscheduleRequest{}
}

type SmartworkAttendsListscheduleRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            SmartworkAttendsListscheduleResponse
	Offset          int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
	WorkDate        time.Time
}

func (this *SmartworkAttendsListscheduleRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *SmartworkAttendsListscheduleRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *SmartworkAttendsListscheduleRequest) GetOffset() int64 {
	return this.Offset
}
func (this *SmartworkAttendsListscheduleRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *SmartworkAttendsListscheduleRequest) GetSize() int64 {
	return this.Size
}
func (this *SmartworkAttendsListscheduleRequest) SetWorkDate(workDate2 time.Time) {
	this.WorkDate = workDate2
}
func (this *SmartworkAttendsListscheduleRequest) GetWorkDate() time.Time {
	return this.WorkDate
}
func (this *SmartworkAttendsListscheduleRequest) GetApiMethodName() string {
	return "dingtalk.smartwork.attends.listschedule"
}
func (this *SmartworkAttendsListscheduleRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *SmartworkAttendsListscheduleRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *SmartworkAttendsListscheduleRequest) GetTopApiCallType() string {
	return "top"
}
func (this *SmartworkAttendsListscheduleRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *SmartworkAttendsListscheduleRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *SmartworkAttendsListscheduleRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["offset"] = this.Offset
	txtParams["size"] = this.Size
	txtParams["work_date"] = this.WorkDate
	return txtParams
}
func (this *SmartworkAttendsListscheduleRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.WorkDate, "workDate"); err != nil {
		return err
	}
	return nil
}
func (this *SmartworkAttendsListscheduleRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *SmartworkAttendsListscheduleRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type SmartworkAttendsListscheduleResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
