package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduUserRelationListRequest() *OapiEduUserRelationListRequest {
	return &OapiEduUserRelationListRequest{}
}

type OapiEduUserRelationListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduUserRelationListResponse
	ClassId         int64
	PageNo          int64
	PageSize        int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduUserRelationListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduUserRelationListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduUserRelationListRequest) SetClassId(classId2 int64) {
	this.ClassId = classId2
}
func (this *OapiEduUserRelationListRequest) GetClassId() int64 {
	return this.ClassId
}
func (this *OapiEduUserRelationListRequest) SetPageNo(pageNo2 int64) {
	this.PageNo = pageNo2
}
func (this *OapiEduUserRelationListRequest) GetPageNo() int64 {
	return this.PageNo
}
func (this *OapiEduUserRelationListRequest) SetPageSize(pageSize2 int64) {
	this.PageSize = pageSize2
}
func (this *OapiEduUserRelationListRequest) GetPageSize() int64 {
	return this.PageSize
}
func (this *OapiEduUserRelationListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.user.relation.list"
}
func (this *OapiEduUserRelationListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduUserRelationListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduUserRelationListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduUserRelationListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduUserRelationListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["class_id"] = this.ClassId
	txtParams["page_no"] = this.PageNo
	txtParams["page_size"] = this.PageSize
	return txtParams
}
func (this *OapiEduUserRelationListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ClassId, "classId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduUserRelationListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduUserRelationListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduUserRelationListResponse struct {
	taobao.TaobaoResponse
	Result  Result `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
