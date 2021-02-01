package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiProcessinstanceListRequest() *OapiProcessinstanceListRequest {
	return &OapiProcessinstanceListRequest{}
}

type OapiProcessinstanceListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessinstanceListResponse
	Cursor          int64
	EndTime         int64
	ProcessCode     string
	Size            int64
	StartTime       int64
	TopHttpMethod   string
	TopResponseType string
	UseridList      string
}

func (this *OapiProcessinstanceListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessinstanceListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessinstanceListRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiProcessinstanceListRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiProcessinstanceListRequest) SetEndTime(endTime2 int64) {
	this.EndTime = endTime2
}
func (this *OapiProcessinstanceListRequest) GetEndTime() int64 {
	return this.EndTime
}
func (this *OapiProcessinstanceListRequest) SetProcessCode(processCode2 string) {
	this.ProcessCode = processCode2
}
func (this *OapiProcessinstanceListRequest) GetProcessCode() string {
	return this.ProcessCode
}
func (this *OapiProcessinstanceListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiProcessinstanceListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiProcessinstanceListRequest) SetStartTime(startTime2 int64) {
	this.StartTime = startTime2
}
func (this *OapiProcessinstanceListRequest) GetStartTime() int64 {
	return this.StartTime
}
func (this *OapiProcessinstanceListRequest) SetUseridList(useridList2 string) {
	this.UseridList = useridList2
}
func (this *OapiProcessinstanceListRequest) GetUseridList() string {
	return this.UseridList
}
func (this *OapiProcessinstanceListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.processinstance.list"
}
func (this *OapiProcessinstanceListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessinstanceListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessinstanceListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessinstanceListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessinstanceListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cursor"] = this.Cursor
	txtParams["end_time"] = this.EndTime
	txtParams["process_code"] = this.ProcessCode
	txtParams["size"] = this.Size
	txtParams["start_time"] = this.StartTime
	txtParams["userid_list"] = this.UseridList
	return txtParams
}
func (this *OapiProcessinstanceListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ProcessCode, "processCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiProcessinstanceListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessinstanceListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiProcessinstanceListResponse struct {
	taobao.TaobaoResponse
	Result PageResult `json:"result,omitempty"`
}
