package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCollectionFormCreateRequest() *OapiCollectionFormCreateRequest {
	return &OapiCollectionFormCreateRequest{}
}

type OapiCollectionFormCreateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCollectionFormCreateResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCollectionFormCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCollectionFormCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCollectionFormCreateRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiCollectionFormCreateRequest) GetRequest() string {
	return this.Request
}
func (this *OapiCollectionFormCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.collection.form.create"
}
func (this *OapiCollectionFormCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCollectionFormCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCollectionFormCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCollectionFormCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCollectionFormCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiCollectionFormCreateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCollectionFormCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCollectionFormCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type ProcessVisibleValueVo struct {
	VisibleType  int64  `json:"visible_type,omitempty"`
	VisibleValue string `json:"visible_value,omitempty"`
}
type FormSchemaSettingVo struct {
	BizType        int64     `json:"biz_type,omitempty"`
	CollectionType int64     `json:"collection_type,omitempty"`
	EndTime        time.Time `json:"end_time,omitempty"`
	FormType       int64     `json:"form_type,omitempty"`
	LoopDayOfWeeks []int64   `json:"loop_day_of_weeks,omitempty"`
	LoopTime       string    `json:"loop_time,omitempty"`
	ReplyTime      bool      `json:"reply_time,omitempty"`
	SubSource      int64     `json:"sub_source,omitempty"`
}
type BehaviorTarget struct {
	Behavior string `json:"behavior,omitempty"`
	FieldId  string `json:"field_id,omitempty"`
}
type BehaviorLinkageVo struct {
	Targets []BehaviorTarget `json:"targets,omitempty"`
	Value   string           `json:"value,omitempty"`
}
type ComponentPropOptionVo struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}
type ComponentMatrixDefVo struct {
	Alias string `json:"alias,omitempty"`
	Key   string `json:"key,omitempty"`
	Name  string `json:"name,omitempty"`
}
type ComponentPropVo struct {
	BehaviorLinkage []BehaviorLinkageVo     `json:"behavior_linkage,omitempty"`
	BizAlias        string                  `json:"biz_alias,omitempty"`
	Cols            []ComponentMatrixDefVo  `json:"cols,omitempty"`
	Id              string                  `json:"id,omitempty"`
	Label           string                  `json:"label,omitempty"`
	Options         []ComponentPropOptionVo `json:"options,omitempty"`
	Placeholder     string                  `json:"placeholder,omitempty"`
	Required        bool                    `json:"required,omitempty"`
	Rows            []ComponentMatrixDefVo  `json:"rows,omitempty"`
}
type FormComponentVo struct {
	ComponentName string          `json:"component_name,omitempty"`
	Props         ComponentPropVo `json:"props,omitempty"`
}
type FormContentVo struct {
	Items []FormComponentVo `json:"items,omitempty"`
}
type SwFormVisibleValueVo struct {
	CidEncrypted int64  `json:"cid_encrypted,omitempty"`
	VisibleType  int64  `json:"visible_type,omitempty"`
	VisibleValue string `json:"visible_value,omitempty"`
}
type SaveFormSchemaRequest struct {
	Content            string                  `json:"content,omitempty"`
	CustomSetting      FormSchemaSettingVo     `json:"custom_setting,omitempty"`
	FormContent        FormContentVo           `json:"form_content,omitempty"`
	Icon               string                  `json:"icon,omitempty"`
	Memo               string                  `json:"memo,omitempty"`
	Name               string                  `json:"name,omitempty"`
	ProcessVisibleList []ProcessVisibleValueVo `json:"process_visible_list,omitempty"`
	Userid             string                  `json:"userid,omitempty"`
	VisibleValueList   []SwFormVisibleValueVo  `json:"visible_value_list,omitempty"`
}
type OapiCollectionFormCreateResponse struct {
	taobao.TaobaoResponse
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
