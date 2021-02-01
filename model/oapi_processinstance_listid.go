package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiProcessinstanceListidsRequest() *OapiProcessinstanceListidsRequest {
	return &OapiProcessinstanceListidsRequest{}
}

type OapiProcessinstanceListidsRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessinstanceListidsResponse
	Cursor          int64
	EndTime         int64
	ProcessCode     string
	Size            int64
	StartTime       int64
	TopHttpMethod   string
	TopResponseType string
	UseridList      string
}

func (this *OapiProcessinstanceListidsRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessinstanceListidsRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessinstanceListidsRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiProcessinstanceListidsRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiProcessinstanceListidsRequest) SetEndTime(endTime2 int64) {
	this.EndTime = endTime2
}
func (this *OapiProcessinstanceListidsRequest) GetEndTime() int64 {
	return this.EndTime
}
func (this *OapiProcessinstanceListidsRequest) SetProcessCode(processCode2 string) {
	this.ProcessCode = processCode2
}
func (this *OapiProcessinstanceListidsRequest) GetProcessCode() string {
	return this.ProcessCode
}
func (this *OapiProcessinstanceListidsRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiProcessinstanceListidsRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiProcessinstanceListidsRequest) SetStartTime(startTime2 int64) {
	this.StartTime = startTime2
}
func (this *OapiProcessinstanceListidsRequest) GetStartTime() int64 {
	return this.StartTime
}
func (this *OapiProcessinstanceListidsRequest) SetUseridList(useridList2 string) {
	this.UseridList = useridList2
}
func (this *OapiProcessinstanceListidsRequest) GetUseridList() string {
	return this.UseridList
}
func (this *OapiProcessinstanceListidsRequest) GetApiMethodName() string {
	return "dingtalk.oapi.processinstance.listids"
}
func (this *OapiProcessinstanceListidsRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessinstanceListidsRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessinstanceListidsRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessinstanceListidsRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessinstanceListidsRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cursor"] = this.Cursor
	txtParams["end_time"] = this.EndTime
	txtParams["process_code"] = this.ProcessCode
	txtParams["size"] = this.Size
	txtParams["start_time"] = this.StartTime
	txtParams["userid_list"] = this.UseridList
	return txtParams
}
func (this *OapiProcessinstanceListidsRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ProcessCode, "processCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiProcessinstanceListidsRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessinstanceListidsRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiProcessinstanceListidsResponse struct {
	taobao.TaobaoResponse
	Result PageResult `json:"result,omitempty"`
}
