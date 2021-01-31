package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMiniappMiniappversionQueryRequest() *OapiMiniappMiniappversionQueryRequest {
	return &OapiMiniappMiniappversionQueryRequest{}
}

type OapiMiniappMiniappversionQueryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMiniappMiniappversionQueryResponse
	ModelKey        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiMiniappMiniappversionQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMiniappMiniappversionQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMiniappMiniappversionQueryRequest) SetModelKey(modelKey2 string) {
	this.ModelKey = modelKey2
}
func (this *OapiMiniappMiniappversionQueryRequest) GetModelKey() string {
	return this.ModelKey
}
func (this *OapiMiniappMiniappversionQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.miniapp.miniappversion.query"
}
func (this *OapiMiniappMiniappversionQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMiniappMiniappversionQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMiniappMiniappversionQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMiniappMiniappversionQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMiniappMiniappversionQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["model_key"] = this.ModelKey
	return txtParams
}
func (this *OapiMiniappMiniappversionQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ModelKey, "modelKey"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMiniappMiniappversionQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMiniappMiniappversionQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMiniappMiniappversionQueryResponse struct {
	taobao.TaobaoResponse
	Data    MiniAppVersionDoModel `json:"data,omitempty"`
	Errcode int64                 `json:"errcode,omitempty"`
	Errmsg  string                `json:"errmsg,omitempty"`
}
type MiniAppVersionDoModel struct {
	AppId           string `json:"appId,omitempty"`
	AuditFinishTime int64  `json:"audit_finish_time,omitempty"`
	AuditSubmitTime int64  `json:"audit_submit_time,omitempty"`
	BuildSource     string `json:"build_source,omitempty"`
	ClientType      int64  `json:"client_type,omitempty"`
	Description     string `json:"description,omitempty"`
	ExpVersion      int64  `json:"exp_version,omitempty"`
	ExtraInfo       string `json:"extra_info,omitempty"`
	GmtCreate       int64  `json:"gmt_create,omitempty"`
	GmtModified     int64  `json:"gmt_modified,omitempty"`
	GrayStartTime   int64  `json:"gray_start_time,omitempty"`
	GrayStrategy    string `json:"gray_strategy,omitempty"`
	Id              int64  `json:"id,omitempty"`
	InstId          int64  `json:"inst_id,omitempty"`
	IsDeleted       int64  `json:"is_deleted,omitempty"`
	OfflineTime     int64  `json:"offline_time,omitempty"`
	PackageUrl      string `json:"package_url,omitempty"`
	Pid             string `json:"pid,omitempty"`
	RollbackTime    int64  `json:"rollback_time,omitempty"`
	ShelfTime       int64  `json:"shelf_time,omitempty"`
	Status          string `json:"status,omitempty"`
	SubStatus       string `json:"sub_status,omitempty"`
	TemplateExtra   string `json:"template_extra,omitempty"`
	Version         string `json:"version,omitempty"`
}
