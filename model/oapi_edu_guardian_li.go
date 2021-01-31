package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduGuardianListRequest() *OapiEduGuardianListRequest {
	return &OapiEduGuardianListRequest{}
}

type OapiEduGuardianListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduGuardianListResponse
	ClassId         int64
	PageNo          int64
	PageSize        int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduGuardianListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduGuardianListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduGuardianListRequest) SetClassId(classId2 int64) {
	this.ClassId = classId2
}
func (this *OapiEduGuardianListRequest) GetClassId() int64 {
	return this.ClassId
}
func (this *OapiEduGuardianListRequest) SetPageNo(pageNo2 int64) {
	this.PageNo = pageNo2
}
func (this *OapiEduGuardianListRequest) GetPageNo() int64 {
	return this.PageNo
}
func (this *OapiEduGuardianListRequest) SetPageSize(pageSize2 int64) {
	this.PageSize = pageSize2
}
func (this *OapiEduGuardianListRequest) GetPageSize() int64 {
	return this.PageSize
}
func (this *OapiEduGuardianListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.guardian.list"
}
func (this *OapiEduGuardianListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduGuardianListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduGuardianListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduGuardianListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduGuardianListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["class_id"] = this.ClassId
	txtParams["page_no"] = this.PageNo
	txtParams["page_size"] = this.PageSize
	return txtParams
}
func (this *OapiEduGuardianListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ClassId, "classId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduGuardianListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduGuardianListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduGuardianListResponse struct {
	taobao.TaobaoResponse
	Result  PageResult `json:"result,omitempty"`
	Success bool       `json:"success,omitempty"`
}
