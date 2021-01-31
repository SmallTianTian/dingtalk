package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewSmartworkCheckinRecordGetRequest() *SmartworkCheckinRecordGetRequest {
	return &SmartworkCheckinRecordGetRequest{}
}

type SmartworkCheckinRecordGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            SmartworkCheckinRecordGetResponse
	Cursor          int64
	EndTime         int64
	Size            int64
	StartTime       int64
	TopHttpMethod   string
	TopResponseType string
	UseridList      string
}

func (this *SmartworkCheckinRecordGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *SmartworkCheckinRecordGetRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *SmartworkCheckinRecordGetRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *SmartworkCheckinRecordGetRequest) SetEndTime(endTime2 int64) {
	this.EndTime = endTime2
}
func (this *SmartworkCheckinRecordGetRequest) GetEndTime() int64 {
	return this.EndTime
}
func (this *SmartworkCheckinRecordGetRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *SmartworkCheckinRecordGetRequest) GetSize() int64 {
	return this.Size
}
func (this *SmartworkCheckinRecordGetRequest) SetStartTime(startTime2 int64) {
	this.StartTime = startTime2
}
func (this *SmartworkCheckinRecordGetRequest) GetStartTime() int64 {
	return this.StartTime
}
func (this *SmartworkCheckinRecordGetRequest) SetUseridList(useridList2 string) {
	this.UseridList = useridList2
}
func (this *SmartworkCheckinRecordGetRequest) GetUseridList() string {
	return this.UseridList
}
func (this *SmartworkCheckinRecordGetRequest) GetApiMethodName() string {
	return "dingtalk.smartwork.checkin.record.get"
}
func (this *SmartworkCheckinRecordGetRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *SmartworkCheckinRecordGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *SmartworkCheckinRecordGetRequest) GetTopApiCallType() string {
	return "top"
}
func (this *SmartworkCheckinRecordGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *SmartworkCheckinRecordGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *SmartworkCheckinRecordGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cursor"] = this.Cursor
	txtParams["end_time"] = this.EndTime
	txtParams["size"] = this.Size
	txtParams["start_time"] = this.StartTime
	txtParams["userid_list"] = this.UseridList
	return txtParams
}
func (this *SmartworkCheckinRecordGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Cursor, "cursor"); err != nil {
		return err
	}
	return nil
}
func (this *SmartworkCheckinRecordGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *SmartworkCheckinRecordGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type SmartworkCheckinRecordGetResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
