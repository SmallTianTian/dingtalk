package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWorkrecordAddRequest() *OapiWorkrecordAddRequest {
	return &OapiWorkrecordAddRequest{}
}

type OapiWorkrecordAddRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp             OapiWorkrecordAddResponse
	BizId            string
	CreateTime       int64
	FormItemList     string
	OriginatorUserId string
	PcOpenType       int64
	PcUrl            string
	SourceName       string
	Title            string
	TopHttpMethod    string
	TopResponseType  string
	Url              string
	Userid           string
}

func (this *OapiWorkrecordAddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkrecordAddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkrecordAddRequest) SetBizId(bizId2 string) {
	this.BizId = bizId2
}
func (this *OapiWorkrecordAddRequest) GetBizId() string {
	return this.BizId
}
func (this *OapiWorkrecordAddRequest) SetCreateTime(createTime2 int64) {
	this.CreateTime = createTime2
}
func (this *OapiWorkrecordAddRequest) GetCreateTime() int64 {
	return this.CreateTime
}
func (this *OapiWorkrecordAddRequest) SetFormItemList(formItemList2 string) {
	this.FormItemList = formItemList2
}
func (this *OapiWorkrecordAddRequest) GetFormItemList() string {
	return this.FormItemList
}
func (this *OapiWorkrecordAddRequest) SetOriginatorUserId(originatorUserId2 string) {
	this.OriginatorUserId = originatorUserId2
}
func (this *OapiWorkrecordAddRequest) GetOriginatorUserId() string {
	return this.OriginatorUserId
}
func (this *OapiWorkrecordAddRequest) SetPcUrl(pcUrl2 string) {
	this.PcUrl = pcUrl2
}
func (this *OapiWorkrecordAddRequest) GetPcUrl() string {
	return this.PcUrl
}
func (this *OapiWorkrecordAddRequest) SetPcOpenType(pcOpenType2 int64) {
	this.PcOpenType = pcOpenType2
}
func (this *OapiWorkrecordAddRequest) GetPcOpenType() int64 {
	return this.PcOpenType
}
func (this *OapiWorkrecordAddRequest) SetSourceName(sourceName2 string) {
	this.SourceName = sourceName2
}
func (this *OapiWorkrecordAddRequest) GetSourceName() string {
	return this.SourceName
}
func (this *OapiWorkrecordAddRequest) SetTitle(title2 string) {
	this.Title = title2
}
func (this *OapiWorkrecordAddRequest) GetTitle() string {
	return this.Title
}
func (this *OapiWorkrecordAddRequest) SetUrl(url2 string) {
	this.Url = url2
}
func (this *OapiWorkrecordAddRequest) GetUrl() string {
	return this.Url
}
func (this *OapiWorkrecordAddRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiWorkrecordAddRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiWorkrecordAddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workrecord.add"
}
func (this *OapiWorkrecordAddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkrecordAddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkrecordAddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkrecordAddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkrecordAddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_id"] = this.BizId
	txtParams["create_time"] = this.CreateTime
	txtParams["formItemList"] = this.FormItemList
	txtParams["originator_user_id"] = this.OriginatorUserId
	txtParams["pcUrl"] = this.PcUrl
	txtParams["pc_open_type"] = this.PcOpenType
	txtParams["source_name"] = this.SourceName
	txtParams["title"] = this.Title
	txtParams["url"] = this.Url
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiWorkrecordAddRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CreateTime, "createTime"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWorkrecordAddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkrecordAddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWorkrecordAddResponse struct {
	taobao.TaobaoResponse
	Errcode  int64  `json:"errcode,omitempty"`
	Errmsg   string `json:"errmsg,omitempty"`
	RecordId string `json:"record_id,omitempty"`
}
