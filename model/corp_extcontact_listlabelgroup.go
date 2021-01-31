package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpExtcontactListlabelgroupsRequest() *CorpExtcontactListlabelgroupsRequest {
	return &CorpExtcontactListlabelgroupsRequest{}
}

type CorpExtcontactListlabelgroupsRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            CorpExtcontactListlabelgroupsResponse
	Offset          int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpExtcontactListlabelgroupsRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpExtcontactListlabelgroupsRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *CorpExtcontactListlabelgroupsRequest) GetOffset() int64 {
	return this.Offset
}
func (this *CorpExtcontactListlabelgroupsRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *CorpExtcontactListlabelgroupsRequest) GetSize() int64 {
	return this.Size
}
func (this *CorpExtcontactListlabelgroupsRequest) GetApiMethodName() string {
	return "dingtalk.corp.extcontact.listlabelgroups"
}
func (this *CorpExtcontactListlabelgroupsRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpExtcontactListlabelgroupsRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpExtcontactListlabelgroupsRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpExtcontactListlabelgroupsRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpExtcontactListlabelgroupsRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpExtcontactListlabelgroupsRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["offset"] = this.Offset
	txtParams["size"] = this.Size
	return txtParams
}
func (this *CorpExtcontactListlabelgroupsRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckMaxValue(int(this.Size), 100, "size"); err != nil {
		return err
	}
	return nil
}
func (this *CorpExtcontactListlabelgroupsRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpExtcontactListlabelgroupsRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpExtcontactListlabelgroupsResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
type OpenLabel struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
type OpenLabelGroup struct {
	Color  int64       `json:"color,omitempty"`
	Labels []OpenLabel `json:"labels,omitempty"`
	Name   string      `json:"name,omitempty"`
}
