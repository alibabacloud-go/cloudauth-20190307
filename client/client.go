// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/v2/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type CompareFaceVerifyRequest struct {
	Crop                         *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	OuterOrderNo                 *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	ProductCode                  *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	SceneId                      *int64  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	SourceCertifyId              *string `json:"SourceCertifyId,omitempty" xml:"SourceCertifyId,omitempty"`
	SourceFaceContrastPicture    *string `json:"SourceFaceContrastPicture,omitempty" xml:"SourceFaceContrastPicture,omitempty"`
	SourceFaceContrastPictureUrl *string `json:"SourceFaceContrastPictureUrl,omitempty" xml:"SourceFaceContrastPictureUrl,omitempty"`
	SourceOssBucketName          *string `json:"SourceOssBucketName,omitempty" xml:"SourceOssBucketName,omitempty"`
	SourceOssObjectName          *string `json:"SourceOssObjectName,omitempty" xml:"SourceOssObjectName,omitempty"`
	TargetCertifyId              *string `json:"TargetCertifyId,omitempty" xml:"TargetCertifyId,omitempty"`
	TargetFaceContrastPicture    *string `json:"TargetFaceContrastPicture,omitempty" xml:"TargetFaceContrastPicture,omitempty"`
	TargetFaceContrastPictureUrl *string `json:"TargetFaceContrastPictureUrl,omitempty" xml:"TargetFaceContrastPictureUrl,omitempty"`
	TargetOssBucketName          *string `json:"TargetOssBucketName,omitempty" xml:"TargetOssBucketName,omitempty"`
	TargetOssObjectName          *string `json:"TargetOssObjectName,omitempty" xml:"TargetOssObjectName,omitempty"`
}

func (s CompareFaceVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s CompareFaceVerifyRequest) GoString() string {
	return s.String()
}

func (s *CompareFaceVerifyRequest) SetCrop(v string) *CompareFaceVerifyRequest {
	s.Crop = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetOuterOrderNo(v string) *CompareFaceVerifyRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetProductCode(v string) *CompareFaceVerifyRequest {
	s.ProductCode = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetSceneId(v int64) *CompareFaceVerifyRequest {
	s.SceneId = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetSourceCertifyId(v string) *CompareFaceVerifyRequest {
	s.SourceCertifyId = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetSourceFaceContrastPicture(v string) *CompareFaceVerifyRequest {
	s.SourceFaceContrastPicture = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetSourceFaceContrastPictureUrl(v string) *CompareFaceVerifyRequest {
	s.SourceFaceContrastPictureUrl = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetSourceOssBucketName(v string) *CompareFaceVerifyRequest {
	s.SourceOssBucketName = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetSourceOssObjectName(v string) *CompareFaceVerifyRequest {
	s.SourceOssObjectName = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetTargetCertifyId(v string) *CompareFaceVerifyRequest {
	s.TargetCertifyId = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetTargetFaceContrastPicture(v string) *CompareFaceVerifyRequest {
	s.TargetFaceContrastPicture = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetTargetFaceContrastPictureUrl(v string) *CompareFaceVerifyRequest {
	s.TargetFaceContrastPictureUrl = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetTargetOssBucketName(v string) *CompareFaceVerifyRequest {
	s.TargetOssBucketName = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetTargetOssObjectName(v string) *CompareFaceVerifyRequest {
	s.TargetOssObjectName = &v
	return s
}

type CompareFaceVerifyResponseBody struct {
	Code         *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *CompareFaceVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s CompareFaceVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CompareFaceVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *CompareFaceVerifyResponseBody) SetCode(v string) *CompareFaceVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *CompareFaceVerifyResponseBody) SetMessage(v string) *CompareFaceVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *CompareFaceVerifyResponseBody) SetRequestId(v string) *CompareFaceVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CompareFaceVerifyResponseBody) SetResultObject(v *CompareFaceVerifyResponseBodyResultObject) *CompareFaceVerifyResponseBody {
	s.ResultObject = v
	return s
}

type CompareFaceVerifyResponseBodyResultObject struct {
	CertifyId   *string  `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	Passed      *string  `json:"Passed,omitempty" xml:"Passed,omitempty"`
	VerifyScore *float32 `json:"VerifyScore,omitempty" xml:"VerifyScore,omitempty"`
}

func (s CompareFaceVerifyResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s CompareFaceVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *CompareFaceVerifyResponseBodyResultObject) SetCertifyId(v string) *CompareFaceVerifyResponseBodyResultObject {
	s.CertifyId = &v
	return s
}

func (s *CompareFaceVerifyResponseBodyResultObject) SetPassed(v string) *CompareFaceVerifyResponseBodyResultObject {
	s.Passed = &v
	return s
}

func (s *CompareFaceVerifyResponseBodyResultObject) SetVerifyScore(v float32) *CompareFaceVerifyResponseBodyResultObject {
	s.VerifyScore = &v
	return s
}

type CompareFaceVerifyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CompareFaceVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CompareFaceVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s CompareFaceVerifyResponse) GoString() string {
	return s.String()
}

func (s *CompareFaceVerifyResponse) SetHeaders(v map[string]*string) *CompareFaceVerifyResponse {
	s.Headers = v
	return s
}

func (s *CompareFaceVerifyResponse) SetStatusCode(v int32) *CompareFaceVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *CompareFaceVerifyResponse) SetBody(v *CompareFaceVerifyResponseBody) *CompareFaceVerifyResponse {
	s.Body = v
	return s
}

type CompareFacesRequest struct {
	SourceImageType  *string `json:"SourceImageType,omitempty" xml:"SourceImageType,omitempty"`
	SourceImageValue *string `json:"SourceImageValue,omitempty" xml:"SourceImageValue,omitempty"`
	TargetImageType  *string `json:"TargetImageType,omitempty" xml:"TargetImageType,omitempty"`
	TargetImageValue *string `json:"TargetImageValue,omitempty" xml:"TargetImageValue,omitempty"`
}

func (s CompareFacesRequest) String() string {
	return tea.Prettify(s)
}

func (s CompareFacesRequest) GoString() string {
	return s.String()
}

func (s *CompareFacesRequest) SetSourceImageType(v string) *CompareFacesRequest {
	s.SourceImageType = &v
	return s
}

func (s *CompareFacesRequest) SetSourceImageValue(v string) *CompareFacesRequest {
	s.SourceImageValue = &v
	return s
}

func (s *CompareFacesRequest) SetTargetImageType(v string) *CompareFacesRequest {
	s.TargetImageType = &v
	return s
}

func (s *CompareFacesRequest) SetTargetImageValue(v string) *CompareFacesRequest {
	s.TargetImageValue = &v
	return s
}

type CompareFacesResponseBody struct {
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CompareFacesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CompareFacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CompareFacesResponseBody) GoString() string {
	return s.String()
}

func (s *CompareFacesResponseBody) SetCode(v string) *CompareFacesResponseBody {
	s.Code = &v
	return s
}

func (s *CompareFacesResponseBody) SetData(v *CompareFacesResponseBodyData) *CompareFacesResponseBody {
	s.Data = v
	return s
}

func (s *CompareFacesResponseBody) SetMessage(v string) *CompareFacesResponseBody {
	s.Message = &v
	return s
}

func (s *CompareFacesResponseBody) SetRequestId(v string) *CompareFacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CompareFacesResponseBody) SetSuccess(v bool) *CompareFacesResponseBody {
	s.Success = &v
	return s
}

type CompareFacesResponseBodyData struct {
	ConfidenceThresholds *string  `json:"ConfidenceThresholds,omitempty" xml:"ConfidenceThresholds,omitempty"`
	SimilarityScore      *float32 `json:"SimilarityScore,omitempty" xml:"SimilarityScore,omitempty"`
}

func (s CompareFacesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CompareFacesResponseBodyData) GoString() string {
	return s.String()
}

func (s *CompareFacesResponseBodyData) SetConfidenceThresholds(v string) *CompareFacesResponseBodyData {
	s.ConfidenceThresholds = &v
	return s
}

func (s *CompareFacesResponseBodyData) SetSimilarityScore(v float32) *CompareFacesResponseBodyData {
	s.SimilarityScore = &v
	return s
}

type CompareFacesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CompareFacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CompareFacesResponse) String() string {
	return tea.Prettify(s)
}

func (s CompareFacesResponse) GoString() string {
	return s.String()
}

func (s *CompareFacesResponse) SetHeaders(v map[string]*string) *CompareFacesResponse {
	s.Headers = v
	return s
}

func (s *CompareFacesResponse) SetStatusCode(v int32) *CompareFacesResponse {
	s.StatusCode = &v
	return s
}

func (s *CompareFacesResponse) SetBody(v *CompareFacesResponseBody) *CompareFacesResponse {
	s.Body = v
	return s
}

type ContrastFaceVerifyRequest struct {
	CertName               *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertNo                 *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	CertType               *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	CertifyId              *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	Crop                   *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	DeviceToken            *string `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
	EncryptType            *string `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	FaceContrastFile       *string `json:"FaceContrastFile,omitempty" xml:"FaceContrastFile,omitempty"`
	FaceContrastPicture    *string `json:"FaceContrastPicture,omitempty" xml:"FaceContrastPicture,omitempty"`
	FaceContrastPictureUrl *string `json:"FaceContrastPictureUrl,omitempty" xml:"FaceContrastPictureUrl,omitempty"`
	Ip                     *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Mobile                 *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Model                  *string `json:"Model,omitempty" xml:"Model,omitempty"`
	OssBucketName          *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	OssObjectName          *string `json:"OssObjectName,omitempty" xml:"OssObjectName,omitempty"`
	OuterOrderNo           *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	ProductCode            *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	SceneId                *int64  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	UserId                 *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ContrastFaceVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s ContrastFaceVerifyRequest) GoString() string {
	return s.String()
}

func (s *ContrastFaceVerifyRequest) SetCertName(v string) *ContrastFaceVerifyRequest {
	s.CertName = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetCertNo(v string) *ContrastFaceVerifyRequest {
	s.CertNo = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetCertType(v string) *ContrastFaceVerifyRequest {
	s.CertType = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetCertifyId(v string) *ContrastFaceVerifyRequest {
	s.CertifyId = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetCrop(v string) *ContrastFaceVerifyRequest {
	s.Crop = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetDeviceToken(v string) *ContrastFaceVerifyRequest {
	s.DeviceToken = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetEncryptType(v string) *ContrastFaceVerifyRequest {
	s.EncryptType = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetFaceContrastFile(v string) *ContrastFaceVerifyRequest {
	s.FaceContrastFile = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetFaceContrastPicture(v string) *ContrastFaceVerifyRequest {
	s.FaceContrastPicture = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetFaceContrastPictureUrl(v string) *ContrastFaceVerifyRequest {
	s.FaceContrastPictureUrl = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetIp(v string) *ContrastFaceVerifyRequest {
	s.Ip = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetMobile(v string) *ContrastFaceVerifyRequest {
	s.Mobile = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetModel(v string) *ContrastFaceVerifyRequest {
	s.Model = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetOssBucketName(v string) *ContrastFaceVerifyRequest {
	s.OssBucketName = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetOssObjectName(v string) *ContrastFaceVerifyRequest {
	s.OssObjectName = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetOuterOrderNo(v string) *ContrastFaceVerifyRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetProductCode(v string) *ContrastFaceVerifyRequest {
	s.ProductCode = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetSceneId(v int64) *ContrastFaceVerifyRequest {
	s.SceneId = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetUserId(v string) *ContrastFaceVerifyRequest {
	s.UserId = &v
	return s
}

type ContrastFaceVerifyAdvanceRequest struct {
	CertName               *string   `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertNo                 *string   `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	CertType               *string   `json:"CertType,omitempty" xml:"CertType,omitempty"`
	CertifyId              *string   `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	Crop                   *string   `json:"Crop,omitempty" xml:"Crop,omitempty"`
	DeviceToken            *string   `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
	EncryptType            *string   `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	FaceContrastFileObject io.Reader `json:"FaceContrastFile,omitempty" xml:"FaceContrastFile,omitempty"`
	FaceContrastPicture    *string   `json:"FaceContrastPicture,omitempty" xml:"FaceContrastPicture,omitempty"`
	FaceContrastPictureUrl *string   `json:"FaceContrastPictureUrl,omitempty" xml:"FaceContrastPictureUrl,omitempty"`
	Ip                     *string   `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Mobile                 *string   `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Model                  *string   `json:"Model,omitempty" xml:"Model,omitempty"`
	OssBucketName          *string   `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	OssObjectName          *string   `json:"OssObjectName,omitempty" xml:"OssObjectName,omitempty"`
	OuterOrderNo           *string   `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	ProductCode            *string   `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	SceneId                *int64    `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	UserId                 *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ContrastFaceVerifyAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ContrastFaceVerifyAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ContrastFaceVerifyAdvanceRequest) SetCertName(v string) *ContrastFaceVerifyAdvanceRequest {
	s.CertName = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetCertNo(v string) *ContrastFaceVerifyAdvanceRequest {
	s.CertNo = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetCertType(v string) *ContrastFaceVerifyAdvanceRequest {
	s.CertType = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetCertifyId(v string) *ContrastFaceVerifyAdvanceRequest {
	s.CertifyId = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetCrop(v string) *ContrastFaceVerifyAdvanceRequest {
	s.Crop = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetDeviceToken(v string) *ContrastFaceVerifyAdvanceRequest {
	s.DeviceToken = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetEncryptType(v string) *ContrastFaceVerifyAdvanceRequest {
	s.EncryptType = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetFaceContrastFileObject(v io.Reader) *ContrastFaceVerifyAdvanceRequest {
	s.FaceContrastFileObject = v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetFaceContrastPicture(v string) *ContrastFaceVerifyAdvanceRequest {
	s.FaceContrastPicture = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetFaceContrastPictureUrl(v string) *ContrastFaceVerifyAdvanceRequest {
	s.FaceContrastPictureUrl = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetIp(v string) *ContrastFaceVerifyAdvanceRequest {
	s.Ip = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetMobile(v string) *ContrastFaceVerifyAdvanceRequest {
	s.Mobile = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetModel(v string) *ContrastFaceVerifyAdvanceRequest {
	s.Model = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetOssBucketName(v string) *ContrastFaceVerifyAdvanceRequest {
	s.OssBucketName = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetOssObjectName(v string) *ContrastFaceVerifyAdvanceRequest {
	s.OssObjectName = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetOuterOrderNo(v string) *ContrastFaceVerifyAdvanceRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetProductCode(v string) *ContrastFaceVerifyAdvanceRequest {
	s.ProductCode = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetSceneId(v int64) *ContrastFaceVerifyAdvanceRequest {
	s.SceneId = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetUserId(v string) *ContrastFaceVerifyAdvanceRequest {
	s.UserId = &v
	return s
}

type ContrastFaceVerifyResponseBody struct {
	Code         *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *ContrastFaceVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s ContrastFaceVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ContrastFaceVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *ContrastFaceVerifyResponseBody) SetCode(v string) *ContrastFaceVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *ContrastFaceVerifyResponseBody) SetMessage(v string) *ContrastFaceVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *ContrastFaceVerifyResponseBody) SetRequestId(v string) *ContrastFaceVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ContrastFaceVerifyResponseBody) SetResultObject(v *ContrastFaceVerifyResponseBodyResultObject) *ContrastFaceVerifyResponseBody {
	s.ResultObject = v
	return s
}

type ContrastFaceVerifyResponseBodyResultObject struct {
	CertifyId    *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	IdentityInfo *string `json:"IdentityInfo,omitempty" xml:"IdentityInfo,omitempty"`
	MaterialInfo *string `json:"MaterialInfo,omitempty" xml:"MaterialInfo,omitempty"`
	Passed       *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	SubCode      *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
}

func (s ContrastFaceVerifyResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s ContrastFaceVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *ContrastFaceVerifyResponseBodyResultObject) SetCertifyId(v string) *ContrastFaceVerifyResponseBodyResultObject {
	s.CertifyId = &v
	return s
}

func (s *ContrastFaceVerifyResponseBodyResultObject) SetIdentityInfo(v string) *ContrastFaceVerifyResponseBodyResultObject {
	s.IdentityInfo = &v
	return s
}

func (s *ContrastFaceVerifyResponseBodyResultObject) SetMaterialInfo(v string) *ContrastFaceVerifyResponseBodyResultObject {
	s.MaterialInfo = &v
	return s
}

func (s *ContrastFaceVerifyResponseBodyResultObject) SetPassed(v string) *ContrastFaceVerifyResponseBodyResultObject {
	s.Passed = &v
	return s
}

func (s *ContrastFaceVerifyResponseBodyResultObject) SetSubCode(v string) *ContrastFaceVerifyResponseBodyResultObject {
	s.SubCode = &v
	return s
}

type ContrastFaceVerifyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ContrastFaceVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ContrastFaceVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s ContrastFaceVerifyResponse) GoString() string {
	return s.String()
}

func (s *ContrastFaceVerifyResponse) SetHeaders(v map[string]*string) *ContrastFaceVerifyResponse {
	s.Headers = v
	return s
}

func (s *ContrastFaceVerifyResponse) SetStatusCode(v int32) *ContrastFaceVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *ContrastFaceVerifyResponse) SetBody(v *ContrastFaceVerifyResponseBody) *ContrastFaceVerifyResponse {
	s.Body = v
	return s
}

type CreateAuthKeyRequest struct {
	AuthYears    *int32  `json:"AuthYears,omitempty" xml:"AuthYears,omitempty"`
	BizType      *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	Test         *bool   `json:"Test,omitempty" xml:"Test,omitempty"`
	UserDeviceId *string `json:"UserDeviceId,omitempty" xml:"UserDeviceId,omitempty"`
}

func (s CreateAuthKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAuthKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateAuthKeyRequest) SetAuthYears(v int32) *CreateAuthKeyRequest {
	s.AuthYears = &v
	return s
}

func (s *CreateAuthKeyRequest) SetBizType(v string) *CreateAuthKeyRequest {
	s.BizType = &v
	return s
}

func (s *CreateAuthKeyRequest) SetTest(v bool) *CreateAuthKeyRequest {
	s.Test = &v
	return s
}

func (s *CreateAuthKeyRequest) SetUserDeviceId(v string) *CreateAuthKeyRequest {
	s.UserDeviceId = &v
	return s
}

type CreateAuthKeyResponseBody struct {
	AuthKey   *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAuthKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAuthKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAuthKeyResponseBody) SetAuthKey(v string) *CreateAuthKeyResponseBody {
	s.AuthKey = &v
	return s
}

func (s *CreateAuthKeyResponseBody) SetRequestId(v string) *CreateAuthKeyResponseBody {
	s.RequestId = &v
	return s
}

type CreateAuthKeyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAuthKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAuthKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAuthKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateAuthKeyResponse) SetHeaders(v map[string]*string) *CreateAuthKeyResponse {
	s.Headers = v
	return s
}

func (s *CreateAuthKeyResponse) SetStatusCode(v int32) *CreateAuthKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAuthKeyResponse) SetBody(v *CreateAuthKeyResponseBody) *CreateAuthKeyResponse {
	s.Body = v
	return s
}

type CreateVerifySettingRequest struct {
	BizName     *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	BizType     *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	GuideStep   *bool   `json:"GuideStep,omitempty" xml:"GuideStep,omitempty"`
	PrivacyStep *bool   `json:"PrivacyStep,omitempty" xml:"PrivacyStep,omitempty"`
	ResultStep  *bool   `json:"ResultStep,omitempty" xml:"ResultStep,omitempty"`
	Solution    *string `json:"Solution,omitempty" xml:"Solution,omitempty"`
}

func (s CreateVerifySettingRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVerifySettingRequest) GoString() string {
	return s.String()
}

func (s *CreateVerifySettingRequest) SetBizName(v string) *CreateVerifySettingRequest {
	s.BizName = &v
	return s
}

func (s *CreateVerifySettingRequest) SetBizType(v string) *CreateVerifySettingRequest {
	s.BizType = &v
	return s
}

func (s *CreateVerifySettingRequest) SetGuideStep(v bool) *CreateVerifySettingRequest {
	s.GuideStep = &v
	return s
}

func (s *CreateVerifySettingRequest) SetPrivacyStep(v bool) *CreateVerifySettingRequest {
	s.PrivacyStep = &v
	return s
}

func (s *CreateVerifySettingRequest) SetResultStep(v bool) *CreateVerifySettingRequest {
	s.ResultStep = &v
	return s
}

func (s *CreateVerifySettingRequest) SetSolution(v string) *CreateVerifySettingRequest {
	s.Solution = &v
	return s
}

type CreateVerifySettingResponseBody struct {
	BizName   *string   `json:"BizName,omitempty" xml:"BizName,omitempty"`
	BizType   *string   `json:"BizType,omitempty" xml:"BizType,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Solution  *string   `json:"Solution,omitempty" xml:"Solution,omitempty"`
	StepList  []*string `json:"StepList,omitempty" xml:"StepList,omitempty" type:"Repeated"`
}

func (s CreateVerifySettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVerifySettingResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVerifySettingResponseBody) SetBizName(v string) *CreateVerifySettingResponseBody {
	s.BizName = &v
	return s
}

func (s *CreateVerifySettingResponseBody) SetBizType(v string) *CreateVerifySettingResponseBody {
	s.BizType = &v
	return s
}

func (s *CreateVerifySettingResponseBody) SetRequestId(v string) *CreateVerifySettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVerifySettingResponseBody) SetSolution(v string) *CreateVerifySettingResponseBody {
	s.Solution = &v
	return s
}

func (s *CreateVerifySettingResponseBody) SetStepList(v []*string) *CreateVerifySettingResponseBody {
	s.StepList = v
	return s
}

type CreateVerifySettingResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateVerifySettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVerifySettingResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVerifySettingResponse) GoString() string {
	return s.String()
}

func (s *CreateVerifySettingResponse) SetHeaders(v map[string]*string) *CreateVerifySettingResponse {
	s.Headers = v
	return s
}

func (s *CreateVerifySettingResponse) SetStatusCode(v int32) *CreateVerifySettingResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVerifySettingResponse) SetBody(v *CreateVerifySettingResponseBody) *CreateVerifySettingResponse {
	s.Body = v
	return s
}

type DescribeDeviceInfoRequest struct {
	BizType         *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	CurrentPage     *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DeviceId        *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	ExpiredEndDay   *string `json:"ExpiredEndDay,omitempty" xml:"ExpiredEndDay,omitempty"`
	ExpiredStartDay *string `json:"ExpiredStartDay,omitempty" xml:"ExpiredStartDay,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	UserDeviceId    *string `json:"UserDeviceId,omitempty" xml:"UserDeviceId,omitempty"`
}

func (s DescribeDeviceInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeviceInfoRequest) SetBizType(v string) *DescribeDeviceInfoRequest {
	s.BizType = &v
	return s
}

func (s *DescribeDeviceInfoRequest) SetCurrentPage(v int32) *DescribeDeviceInfoRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDeviceInfoRequest) SetDeviceId(v string) *DescribeDeviceInfoRequest {
	s.DeviceId = &v
	return s
}

func (s *DescribeDeviceInfoRequest) SetExpiredEndDay(v string) *DescribeDeviceInfoRequest {
	s.ExpiredEndDay = &v
	return s
}

func (s *DescribeDeviceInfoRequest) SetExpiredStartDay(v string) *DescribeDeviceInfoRequest {
	s.ExpiredStartDay = &v
	return s
}

func (s *DescribeDeviceInfoRequest) SetPageSize(v int32) *DescribeDeviceInfoRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDeviceInfoRequest) SetUserDeviceId(v string) *DescribeDeviceInfoRequest {
	s.UserDeviceId = &v
	return s
}

type DescribeDeviceInfoResponseBody struct {
	CurrentPage    *int32                                        `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DeviceInfoList *DescribeDeviceInfoResponseBodyDeviceInfoList `json:"DeviceInfoList,omitempty" xml:"DeviceInfoList,omitempty" type:"Struct"`
	PageSize       *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount     *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDeviceInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeviceInfoResponseBody) SetCurrentPage(v int32) *DescribeDeviceInfoResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDeviceInfoResponseBody) SetDeviceInfoList(v *DescribeDeviceInfoResponseBodyDeviceInfoList) *DescribeDeviceInfoResponseBody {
	s.DeviceInfoList = v
	return s
}

func (s *DescribeDeviceInfoResponseBody) SetPageSize(v int32) *DescribeDeviceInfoResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDeviceInfoResponseBody) SetRequestId(v string) *DescribeDeviceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeviceInfoResponseBody) SetTotalCount(v int32) *DescribeDeviceInfoResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDeviceInfoResponseBodyDeviceInfoList struct {
	DeviceInfo []*DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Repeated"`
}

func (s DescribeDeviceInfoResponseBodyDeviceInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceInfoResponseBodyDeviceInfoList) GoString() string {
	return s.String()
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfoList) SetDeviceInfo(v []*DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo) *DescribeDeviceInfoResponseBodyDeviceInfoList {
	s.DeviceInfo = v
	return s
}

type DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo struct {
	BeginDay     *string `json:"BeginDay,omitempty" xml:"BeginDay,omitempty"`
	BizType      *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	DeviceId     *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	ExpiredDay   *string `json:"ExpiredDay,omitempty" xml:"ExpiredDay,omitempty"`
	UserDeviceId *string `json:"UserDeviceId,omitempty" xml:"UserDeviceId,omitempty"`
}

func (s DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo) GoString() string {
	return s.String()
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo) SetBeginDay(v string) *DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo {
	s.BeginDay = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo) SetBizType(v string) *DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo {
	s.BizType = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo) SetDeviceId(v string) *DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo {
	s.DeviceId = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo) SetExpiredDay(v string) *DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo {
	s.ExpiredDay = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo) SetUserDeviceId(v string) *DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo {
	s.UserDeviceId = &v
	return s
}

type DescribeDeviceInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDeviceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDeviceInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeviceInfoResponse) SetHeaders(v map[string]*string) *DescribeDeviceInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeviceInfoResponse) SetStatusCode(v int32) *DescribeDeviceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeviceInfoResponse) SetBody(v *DescribeDeviceInfoResponseBody) *DescribeDeviceInfoResponse {
	s.Body = v
	return s
}

type DescribeFaceVerifyRequest struct {
	CertifyId         *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	PictureReturnType *string `json:"PictureReturnType,omitempty" xml:"PictureReturnType,omitempty"`
	SceneId           *int64  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s DescribeFaceVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceVerifyRequest) GoString() string {
	return s.String()
}

func (s *DescribeFaceVerifyRequest) SetCertifyId(v string) *DescribeFaceVerifyRequest {
	s.CertifyId = &v
	return s
}

func (s *DescribeFaceVerifyRequest) SetPictureReturnType(v string) *DescribeFaceVerifyRequest {
	s.PictureReturnType = &v
	return s
}

func (s *DescribeFaceVerifyRequest) SetSceneId(v int64) *DescribeFaceVerifyRequest {
	s.SceneId = &v
	return s
}

type DescribeFaceVerifyResponseBody struct {
	Code         *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *DescribeFaceVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s DescribeFaceVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFaceVerifyResponseBody) SetCode(v string) *DescribeFaceVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeFaceVerifyResponseBody) SetMessage(v string) *DescribeFaceVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeFaceVerifyResponseBody) SetRequestId(v string) *DescribeFaceVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFaceVerifyResponseBody) SetResultObject(v *DescribeFaceVerifyResponseBodyResultObject) *DescribeFaceVerifyResponseBody {
	s.ResultObject = v
	return s
}

type DescribeFaceVerifyResponseBodyResultObject struct {
	DeviceRisk   *string `json:"DeviceRisk,omitempty" xml:"DeviceRisk,omitempty"`
	DeviceToken  *string `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
	IdentityInfo *string `json:"IdentityInfo,omitempty" xml:"IdentityInfo,omitempty"`
	MaterialInfo *string `json:"MaterialInfo,omitempty" xml:"MaterialInfo,omitempty"`
	Passed       *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	SubCode      *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	Success      *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFaceVerifyResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeFaceVerifyResponseBodyResultObject) SetDeviceRisk(v string) *DescribeFaceVerifyResponseBodyResultObject {
	s.DeviceRisk = &v
	return s
}

func (s *DescribeFaceVerifyResponseBodyResultObject) SetDeviceToken(v string) *DescribeFaceVerifyResponseBodyResultObject {
	s.DeviceToken = &v
	return s
}

func (s *DescribeFaceVerifyResponseBodyResultObject) SetIdentityInfo(v string) *DescribeFaceVerifyResponseBodyResultObject {
	s.IdentityInfo = &v
	return s
}

func (s *DescribeFaceVerifyResponseBodyResultObject) SetMaterialInfo(v string) *DescribeFaceVerifyResponseBodyResultObject {
	s.MaterialInfo = &v
	return s
}

func (s *DescribeFaceVerifyResponseBodyResultObject) SetPassed(v string) *DescribeFaceVerifyResponseBodyResultObject {
	s.Passed = &v
	return s
}

func (s *DescribeFaceVerifyResponseBodyResultObject) SetSubCode(v string) *DescribeFaceVerifyResponseBodyResultObject {
	s.SubCode = &v
	return s
}

func (s *DescribeFaceVerifyResponseBodyResultObject) SetSuccess(v string) *DescribeFaceVerifyResponseBodyResultObject {
	s.Success = &v
	return s
}

type DescribeFaceVerifyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeFaceVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFaceVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceVerifyResponse) GoString() string {
	return s.String()
}

func (s *DescribeFaceVerifyResponse) SetHeaders(v map[string]*string) *DescribeFaceVerifyResponse {
	s.Headers = v
	return s
}

func (s *DescribeFaceVerifyResponse) SetStatusCode(v int32) *DescribeFaceVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFaceVerifyResponse) SetBody(v *DescribeFaceVerifyResponseBody) *DescribeFaceVerifyResponse {
	s.Body = v
	return s
}

type DescribeOssUploadTokenResponseBody struct {
	OssUploadToken *DescribeOssUploadTokenResponseBodyOssUploadToken `json:"OssUploadToken,omitempty" xml:"OssUploadToken,omitempty" type:"Struct"`
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOssUploadTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssUploadTokenResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOssUploadTokenResponseBody) SetOssUploadToken(v *DescribeOssUploadTokenResponseBodyOssUploadToken) *DescribeOssUploadTokenResponseBody {
	s.OssUploadToken = v
	return s
}

func (s *DescribeOssUploadTokenResponseBody) SetRequestId(v string) *DescribeOssUploadTokenResponseBody {
	s.RequestId = &v
	return s
}

type DescribeOssUploadTokenResponseBodyOssUploadToken struct {
	Bucket   *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	EndPoint *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	Expired  *int64  `json:"Expired,omitempty" xml:"Expired,omitempty"`
	Key      *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Path     *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Secret   *string `json:"Secret,omitempty" xml:"Secret,omitempty"`
	Token    *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s DescribeOssUploadTokenResponseBodyOssUploadToken) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssUploadTokenResponseBodyOssUploadToken) GoString() string {
	return s.String()
}

func (s *DescribeOssUploadTokenResponseBodyOssUploadToken) SetBucket(v string) *DescribeOssUploadTokenResponseBodyOssUploadToken {
	s.Bucket = &v
	return s
}

func (s *DescribeOssUploadTokenResponseBodyOssUploadToken) SetEndPoint(v string) *DescribeOssUploadTokenResponseBodyOssUploadToken {
	s.EndPoint = &v
	return s
}

func (s *DescribeOssUploadTokenResponseBodyOssUploadToken) SetExpired(v int64) *DescribeOssUploadTokenResponseBodyOssUploadToken {
	s.Expired = &v
	return s
}

func (s *DescribeOssUploadTokenResponseBodyOssUploadToken) SetKey(v string) *DescribeOssUploadTokenResponseBodyOssUploadToken {
	s.Key = &v
	return s
}

func (s *DescribeOssUploadTokenResponseBodyOssUploadToken) SetPath(v string) *DescribeOssUploadTokenResponseBodyOssUploadToken {
	s.Path = &v
	return s
}

func (s *DescribeOssUploadTokenResponseBodyOssUploadToken) SetSecret(v string) *DescribeOssUploadTokenResponseBodyOssUploadToken {
	s.Secret = &v
	return s
}

func (s *DescribeOssUploadTokenResponseBodyOssUploadToken) SetToken(v string) *DescribeOssUploadTokenResponseBodyOssUploadToken {
	s.Token = &v
	return s
}

type DescribeOssUploadTokenResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeOssUploadTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeOssUploadTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssUploadTokenResponse) GoString() string {
	return s.String()
}

func (s *DescribeOssUploadTokenResponse) SetHeaders(v map[string]*string) *DescribeOssUploadTokenResponse {
	s.Headers = v
	return s
}

func (s *DescribeOssUploadTokenResponse) SetStatusCode(v int32) *DescribeOssUploadTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOssUploadTokenResponse) SetBody(v *DescribeOssUploadTokenResponseBody) *DescribeOssUploadTokenResponse {
	s.Body = v
	return s
}

type DescribeVerifyResultRequest struct {
	BizId   *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
}

func (s DescribeVerifyResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeVerifyResultRequest) SetBizId(v string) *DescribeVerifyResultRequest {
	s.BizId = &v
	return s
}

func (s *DescribeVerifyResultRequest) SetBizType(v string) *DescribeVerifyResultRequest {
	s.BizType = &v
	return s
}

type DescribeVerifyResultResponseBody struct {
	AuthorityComparisionScore *float32                                  `json:"AuthorityComparisionScore,omitempty" xml:"AuthorityComparisionScore,omitempty"`
	FaceComparisonScore       *float32                                  `json:"FaceComparisonScore,omitempty" xml:"FaceComparisonScore,omitempty"`
	IdCardFaceComparisonScore *float32                                  `json:"IdCardFaceComparisonScore,omitempty" xml:"IdCardFaceComparisonScore,omitempty"`
	Material                  *DescribeVerifyResultResponseBodyMaterial `json:"Material,omitempty" xml:"Material,omitempty" type:"Struct"`
	RequestId                 *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VerifyStatus              *int32                                    `json:"VerifyStatus,omitempty" xml:"VerifyStatus,omitempty"`
}

func (s DescribeVerifyResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVerifyResultResponseBody) SetAuthorityComparisionScore(v float32) *DescribeVerifyResultResponseBody {
	s.AuthorityComparisionScore = &v
	return s
}

func (s *DescribeVerifyResultResponseBody) SetFaceComparisonScore(v float32) *DescribeVerifyResultResponseBody {
	s.FaceComparisonScore = &v
	return s
}

func (s *DescribeVerifyResultResponseBody) SetIdCardFaceComparisonScore(v float32) *DescribeVerifyResultResponseBody {
	s.IdCardFaceComparisonScore = &v
	return s
}

func (s *DescribeVerifyResultResponseBody) SetMaterial(v *DescribeVerifyResultResponseBodyMaterial) *DescribeVerifyResultResponseBody {
	s.Material = v
	return s
}

func (s *DescribeVerifyResultResponseBody) SetRequestId(v string) *DescribeVerifyResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifyResultResponseBody) SetVerifyStatus(v int32) *DescribeVerifyResultResponseBody {
	s.VerifyStatus = &v
	return s
}

type DescribeVerifyResultResponseBodyMaterial struct {
	FaceGlobalUrl *string                                             `json:"FaceGlobalUrl,omitempty" xml:"FaceGlobalUrl,omitempty"`
	FaceImageUrl  *string                                             `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty"`
	FaceMask      *bool                                               `json:"FaceMask,omitempty" xml:"FaceMask,omitempty"`
	FaceQuality   *string                                             `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty"`
	IdCardInfo    *DescribeVerifyResultResponseBodyMaterialIdCardInfo `json:"IdCardInfo,omitempty" xml:"IdCardInfo,omitempty" type:"Struct"`
	IdCardName    *string                                             `json:"IdCardName,omitempty" xml:"IdCardName,omitempty"`
	IdCardNumber  *string                                             `json:"IdCardNumber,omitempty" xml:"IdCardNumber,omitempty"`
	VideoUrls     []*string                                           `json:"VideoUrls,omitempty" xml:"VideoUrls,omitempty" type:"Repeated"`
}

func (s DescribeVerifyResultResponseBodyMaterial) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyResultResponseBodyMaterial) GoString() string {
	return s.String()
}

func (s *DescribeVerifyResultResponseBodyMaterial) SetFaceGlobalUrl(v string) *DescribeVerifyResultResponseBodyMaterial {
	s.FaceGlobalUrl = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterial) SetFaceImageUrl(v string) *DescribeVerifyResultResponseBodyMaterial {
	s.FaceImageUrl = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterial) SetFaceMask(v bool) *DescribeVerifyResultResponseBodyMaterial {
	s.FaceMask = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterial) SetFaceQuality(v string) *DescribeVerifyResultResponseBodyMaterial {
	s.FaceQuality = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterial) SetIdCardInfo(v *DescribeVerifyResultResponseBodyMaterialIdCardInfo) *DescribeVerifyResultResponseBodyMaterial {
	s.IdCardInfo = v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterial) SetIdCardName(v string) *DescribeVerifyResultResponseBodyMaterial {
	s.IdCardName = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterial) SetIdCardNumber(v string) *DescribeVerifyResultResponseBodyMaterial {
	s.IdCardNumber = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterial) SetVideoUrls(v []*string) *DescribeVerifyResultResponseBodyMaterial {
	s.VideoUrls = v
	return s
}

type DescribeVerifyResultResponseBodyMaterialIdCardInfo struct {
	Address       *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Authority     *string `json:"Authority,omitempty" xml:"Authority,omitempty"`
	BackImageUrl  *string `json:"BackImageUrl,omitempty" xml:"BackImageUrl,omitempty"`
	Birth         *string `json:"Birth,omitempty" xml:"Birth,omitempty"`
	EndDate       *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	FrontImageUrl *string `json:"FrontImageUrl,omitempty" xml:"FrontImageUrl,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Nationality   *string `json:"Nationality,omitempty" xml:"Nationality,omitempty"`
	Number        *string `json:"Number,omitempty" xml:"Number,omitempty"`
	StartDate     *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeVerifyResultResponseBodyMaterialIdCardInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyResultResponseBodyMaterialIdCardInfo) GoString() string {
	return s.String()
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) SetAddress(v string) *DescribeVerifyResultResponseBodyMaterialIdCardInfo {
	s.Address = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) SetAuthority(v string) *DescribeVerifyResultResponseBodyMaterialIdCardInfo {
	s.Authority = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) SetBackImageUrl(v string) *DescribeVerifyResultResponseBodyMaterialIdCardInfo {
	s.BackImageUrl = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) SetBirth(v string) *DescribeVerifyResultResponseBodyMaterialIdCardInfo {
	s.Birth = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) SetEndDate(v string) *DescribeVerifyResultResponseBodyMaterialIdCardInfo {
	s.EndDate = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) SetFrontImageUrl(v string) *DescribeVerifyResultResponseBodyMaterialIdCardInfo {
	s.FrontImageUrl = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) SetName(v string) *DescribeVerifyResultResponseBodyMaterialIdCardInfo {
	s.Name = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) SetNationality(v string) *DescribeVerifyResultResponseBodyMaterialIdCardInfo {
	s.Nationality = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) SetNumber(v string) *DescribeVerifyResultResponseBodyMaterialIdCardInfo {
	s.Number = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) SetStartDate(v string) *DescribeVerifyResultResponseBodyMaterialIdCardInfo {
	s.StartDate = &v
	return s
}

type DescribeVerifyResultResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVerifyResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVerifyResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeVerifyResultResponse) SetHeaders(v map[string]*string) *DescribeVerifyResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeVerifyResultResponse) SetStatusCode(v int32) *DescribeVerifyResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVerifyResultResponse) SetBody(v *DescribeVerifyResultResponseBody) *DescribeVerifyResultResponse {
	s.Body = v
	return s
}

type DescribeVerifySDKRequest struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeVerifySDKRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifySDKRequest) GoString() string {
	return s.String()
}

func (s *DescribeVerifySDKRequest) SetTaskId(v string) *DescribeVerifySDKRequest {
	s.TaskId = &v
	return s
}

type DescribeVerifySDKResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SdkUrl    *string `json:"SdkUrl,omitempty" xml:"SdkUrl,omitempty"`
}

func (s DescribeVerifySDKResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifySDKResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVerifySDKResponseBody) SetRequestId(v string) *DescribeVerifySDKResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifySDKResponseBody) SetSdkUrl(v string) *DescribeVerifySDKResponseBody {
	s.SdkUrl = &v
	return s
}

type DescribeVerifySDKResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVerifySDKResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVerifySDKResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifySDKResponse) GoString() string {
	return s.String()
}

func (s *DescribeVerifySDKResponse) SetHeaders(v map[string]*string) *DescribeVerifySDKResponse {
	s.Headers = v
	return s
}

func (s *DescribeVerifySDKResponse) SetStatusCode(v int32) *DescribeVerifySDKResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVerifySDKResponse) SetBody(v *DescribeVerifySDKResponseBody) *DescribeVerifySDKResponse {
	s.Body = v
	return s
}

type DescribeVerifyTokenRequest struct {
	BizId                *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType              *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	CallbackSeed         *string `json:"CallbackSeed,omitempty" xml:"CallbackSeed,omitempty"`
	CallbackUrl          *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	FaceRetainedImageUrl *string `json:"FaceRetainedImageUrl,omitempty" xml:"FaceRetainedImageUrl,omitempty"`
	FailedRedirectUrl    *string `json:"FailedRedirectUrl,omitempty" xml:"FailedRedirectUrl,omitempty"`
	IdCardBackImageUrl   *string `json:"IdCardBackImageUrl,omitempty" xml:"IdCardBackImageUrl,omitempty"`
	IdCardFrontImageUrl  *string `json:"IdCardFrontImageUrl,omitempty" xml:"IdCardFrontImageUrl,omitempty"`
	IdCardNumber         *string `json:"IdCardNumber,omitempty" xml:"IdCardNumber,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PassedRedirectUrl    *string `json:"PassedRedirectUrl,omitempty" xml:"PassedRedirectUrl,omitempty"`
	UserId               *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserIp               *string `json:"UserIp,omitempty" xml:"UserIp,omitempty"`
	UserPhoneNumber      *string `json:"UserPhoneNumber,omitempty" xml:"UserPhoneNumber,omitempty"`
	UserRegistTime       *int64  `json:"UserRegistTime,omitempty" xml:"UserRegistTime,omitempty"`
}

func (s DescribeVerifyTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyTokenRequest) GoString() string {
	return s.String()
}

func (s *DescribeVerifyTokenRequest) SetBizId(v string) *DescribeVerifyTokenRequest {
	s.BizId = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetBizType(v string) *DescribeVerifyTokenRequest {
	s.BizType = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetCallbackSeed(v string) *DescribeVerifyTokenRequest {
	s.CallbackSeed = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetCallbackUrl(v string) *DescribeVerifyTokenRequest {
	s.CallbackUrl = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetFaceRetainedImageUrl(v string) *DescribeVerifyTokenRequest {
	s.FaceRetainedImageUrl = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetFailedRedirectUrl(v string) *DescribeVerifyTokenRequest {
	s.FailedRedirectUrl = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetIdCardBackImageUrl(v string) *DescribeVerifyTokenRequest {
	s.IdCardBackImageUrl = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetIdCardFrontImageUrl(v string) *DescribeVerifyTokenRequest {
	s.IdCardFrontImageUrl = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetIdCardNumber(v string) *DescribeVerifyTokenRequest {
	s.IdCardNumber = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetName(v string) *DescribeVerifyTokenRequest {
	s.Name = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetPassedRedirectUrl(v string) *DescribeVerifyTokenRequest {
	s.PassedRedirectUrl = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetUserId(v string) *DescribeVerifyTokenRequest {
	s.UserId = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetUserIp(v string) *DescribeVerifyTokenRequest {
	s.UserIp = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetUserPhoneNumber(v string) *DescribeVerifyTokenRequest {
	s.UserPhoneNumber = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetUserRegistTime(v int64) *DescribeVerifyTokenRequest {
	s.UserRegistTime = &v
	return s
}

type DescribeVerifyTokenResponseBody struct {
	OssUploadToken *DescribeVerifyTokenResponseBodyOssUploadToken `json:"OssUploadToken,omitempty" xml:"OssUploadToken,omitempty" type:"Struct"`
	RequestId      *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VerifyPageUrl  *string                                        `json:"VerifyPageUrl,omitempty" xml:"VerifyPageUrl,omitempty"`
	VerifyToken    *string                                        `json:"VerifyToken,omitempty" xml:"VerifyToken,omitempty"`
}

func (s DescribeVerifyTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyTokenResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVerifyTokenResponseBody) SetOssUploadToken(v *DescribeVerifyTokenResponseBodyOssUploadToken) *DescribeVerifyTokenResponseBody {
	s.OssUploadToken = v
	return s
}

func (s *DescribeVerifyTokenResponseBody) SetRequestId(v string) *DescribeVerifyTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifyTokenResponseBody) SetVerifyPageUrl(v string) *DescribeVerifyTokenResponseBody {
	s.VerifyPageUrl = &v
	return s
}

func (s *DescribeVerifyTokenResponseBody) SetVerifyToken(v string) *DescribeVerifyTokenResponseBody {
	s.VerifyToken = &v
	return s
}

type DescribeVerifyTokenResponseBodyOssUploadToken struct {
	Bucket   *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	EndPoint *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	Expired  *int64  `json:"Expired,omitempty" xml:"Expired,omitempty"`
	Key      *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Path     *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Secret   *string `json:"Secret,omitempty" xml:"Secret,omitempty"`
	Token    *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s DescribeVerifyTokenResponseBodyOssUploadToken) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyTokenResponseBodyOssUploadToken) GoString() string {
	return s.String()
}

func (s *DescribeVerifyTokenResponseBodyOssUploadToken) SetBucket(v string) *DescribeVerifyTokenResponseBodyOssUploadToken {
	s.Bucket = &v
	return s
}

func (s *DescribeVerifyTokenResponseBodyOssUploadToken) SetEndPoint(v string) *DescribeVerifyTokenResponseBodyOssUploadToken {
	s.EndPoint = &v
	return s
}

func (s *DescribeVerifyTokenResponseBodyOssUploadToken) SetExpired(v int64) *DescribeVerifyTokenResponseBodyOssUploadToken {
	s.Expired = &v
	return s
}

func (s *DescribeVerifyTokenResponseBodyOssUploadToken) SetKey(v string) *DescribeVerifyTokenResponseBodyOssUploadToken {
	s.Key = &v
	return s
}

func (s *DescribeVerifyTokenResponseBodyOssUploadToken) SetPath(v string) *DescribeVerifyTokenResponseBodyOssUploadToken {
	s.Path = &v
	return s
}

func (s *DescribeVerifyTokenResponseBodyOssUploadToken) SetSecret(v string) *DescribeVerifyTokenResponseBodyOssUploadToken {
	s.Secret = &v
	return s
}

func (s *DescribeVerifyTokenResponseBodyOssUploadToken) SetToken(v string) *DescribeVerifyTokenResponseBodyOssUploadToken {
	s.Token = &v
	return s
}

type DescribeVerifyTokenResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVerifyTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVerifyTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyTokenResponse) GoString() string {
	return s.String()
}

func (s *DescribeVerifyTokenResponse) SetHeaders(v map[string]*string) *DescribeVerifyTokenResponse {
	s.Headers = v
	return s
}

func (s *DescribeVerifyTokenResponse) SetStatusCode(v int32) *DescribeVerifyTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVerifyTokenResponse) SetBody(v *DescribeVerifyTokenResponseBody) *DescribeVerifyTokenResponse {
	s.Body = v
	return s
}

type DetectFaceAttributesRequest struct {
	BizType       *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	MaterialValue *string `json:"MaterialValue,omitempty" xml:"MaterialValue,omitempty"`
}

func (s DetectFaceAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesRequest) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesRequest) SetBizType(v string) *DetectFaceAttributesRequest {
	s.BizType = &v
	return s
}

func (s *DetectFaceAttributesRequest) SetMaterialValue(v string) *DetectFaceAttributesRequest {
	s.MaterialValue = &v
	return s
}

type DetectFaceAttributesResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DetectFaceAttributesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DetectFaceAttributesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseBody) SetCode(v string) *DetectFaceAttributesResponseBody {
	s.Code = &v
	return s
}

func (s *DetectFaceAttributesResponseBody) SetData(v *DetectFaceAttributesResponseBodyData) *DetectFaceAttributesResponseBody {
	s.Data = v
	return s
}

func (s *DetectFaceAttributesResponseBody) SetMessage(v string) *DetectFaceAttributesResponseBody {
	s.Message = &v
	return s
}

func (s *DetectFaceAttributesResponseBody) SetRequestId(v string) *DetectFaceAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectFaceAttributesResponseBody) SetSuccess(v bool) *DetectFaceAttributesResponseBody {
	s.Success = &v
	return s
}

type DetectFaceAttributesResponseBodyData struct {
	FaceInfos *DetectFaceAttributesResponseBodyDataFaceInfos `json:"FaceInfos,omitempty" xml:"FaceInfos,omitempty" type:"Struct"`
	ImgHeight *int32                                         `json:"ImgHeight,omitempty" xml:"ImgHeight,omitempty"`
	ImgWidth  *int32                                         `json:"ImgWidth,omitempty" xml:"ImgWidth,omitempty"`
}

func (s DetectFaceAttributesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseBodyData) SetFaceInfos(v *DetectFaceAttributesResponseBodyDataFaceInfos) *DetectFaceAttributesResponseBodyData {
	s.FaceInfos = v
	return s
}

func (s *DetectFaceAttributesResponseBodyData) SetImgHeight(v int32) *DetectFaceAttributesResponseBodyData {
	s.ImgHeight = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyData) SetImgWidth(v int32) *DetectFaceAttributesResponseBodyData {
	s.ImgWidth = &v
	return s
}

type DetectFaceAttributesResponseBodyDataFaceInfos struct {
	FaceAttributesDetectInfo []*DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfo `json:"FaceAttributesDetectInfo,omitempty" xml:"FaceAttributesDetectInfo,omitempty" type:"Repeated"`
}

func (s DetectFaceAttributesResponseBodyDataFaceInfos) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseBodyDataFaceInfos) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfos) SetFaceAttributesDetectInfo(v []*DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfo) *DetectFaceAttributesResponseBodyDataFaceInfos {
	s.FaceAttributesDetectInfo = v
	return s
}

type DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfo struct {
	FaceAttributes *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes `json:"FaceAttributes,omitempty" xml:"FaceAttributes,omitempty" type:"Struct"`
	FaceRect       *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect       `json:"FaceRect,omitempty" xml:"FaceRect,omitempty" type:"Struct"`
}

func (s DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfo) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfo) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfo) SetFaceAttributes(v *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfo {
	s.FaceAttributes = v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfo) SetFaceRect(v *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfo {
	s.FaceRect = v
	return s
}

type DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes struct {
	Blur       *float32                                                                                     `json:"Blur,omitempty" xml:"Blur,omitempty"`
	Facequal   *float32                                                                                     `json:"Facequal,omitempty" xml:"Facequal,omitempty"`
	Facetype   *string                                                                                      `json:"Facetype,omitempty" xml:"Facetype,omitempty"`
	Glasses    *string                                                                                      `json:"Glasses,omitempty" xml:"Glasses,omitempty"`
	Headpose   *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose `json:"Headpose,omitempty" xml:"Headpose,omitempty" type:"Struct"`
	Integrity  *int32                                                                                       `json:"Integrity,omitempty" xml:"Integrity,omitempty"`
	Respirator *string                                                                                      `json:"Respirator,omitempty" xml:"Respirator,omitempty"`
	Smiling    *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling  `json:"Smiling,omitempty" xml:"Smiling,omitempty" type:"Struct"`
}

func (s DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetBlur(v float32) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Blur = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetFacequal(v float32) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Facequal = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetFacetype(v string) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Facetype = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetGlasses(v string) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Glasses = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetHeadpose(v *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Headpose = v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetIntegrity(v int32) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Integrity = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetRespirator(v string) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Respirator = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetSmiling(v *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Smiling = v
	return s
}

type DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose struct {
	PitchAngle *float32 `json:"PitchAngle,omitempty" xml:"PitchAngle,omitempty"`
	RollAngle  *float32 `json:"RollAngle,omitempty" xml:"RollAngle,omitempty"`
	YawAngle   *float32 `json:"YawAngle,omitempty" xml:"YawAngle,omitempty"`
}

func (s DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) SetPitchAngle(v float32) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose {
	s.PitchAngle = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) SetRollAngle(v float32) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose {
	s.RollAngle = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) SetYawAngle(v float32) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose {
	s.YawAngle = &v
	return s
}

type DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling struct {
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Value     *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling) SetThreshold(v float32) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling {
	s.Threshold = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling) SetValue(v float32) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling {
	s.Value = &v
	return s
}

type DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect) SetHeight(v int32) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect {
	s.Height = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect) SetLeft(v int32) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect {
	s.Left = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect) SetTop(v int32) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect {
	s.Top = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect) SetWidth(v int32) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect {
	s.Width = &v
	return s
}

type DetectFaceAttributesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectFaceAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectFaceAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponse) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponse) SetHeaders(v map[string]*string) *DetectFaceAttributesResponse {
	s.Headers = v
	return s
}

func (s *DetectFaceAttributesResponse) SetStatusCode(v int32) *DetectFaceAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectFaceAttributesResponse) SetBody(v *DetectFaceAttributesResponseBody) *DetectFaceAttributesResponse {
	s.Body = v
	return s
}

type Id2MetaVerifyRequest struct {
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	ParamType   *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	UserName    *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s Id2MetaVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s Id2MetaVerifyRequest) GoString() string {
	return s.String()
}

func (s *Id2MetaVerifyRequest) SetIdentifyNum(v string) *Id2MetaVerifyRequest {
	s.IdentifyNum = &v
	return s
}

func (s *Id2MetaVerifyRequest) SetParamType(v string) *Id2MetaVerifyRequest {
	s.ParamType = &v
	return s
}

func (s *Id2MetaVerifyRequest) SetUserName(v string) *Id2MetaVerifyRequest {
	s.UserName = &v
	return s
}

type Id2MetaVerifyResponseBody struct {
	Code         *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *Id2MetaVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s Id2MetaVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s Id2MetaVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *Id2MetaVerifyResponseBody) SetCode(v string) *Id2MetaVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *Id2MetaVerifyResponseBody) SetMessage(v string) *Id2MetaVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *Id2MetaVerifyResponseBody) SetRequestId(v string) *Id2MetaVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *Id2MetaVerifyResponseBody) SetResultObject(v *Id2MetaVerifyResponseBodyResultObject) *Id2MetaVerifyResponseBody {
	s.ResultObject = v
	return s
}

type Id2MetaVerifyResponseBodyResultObject struct {
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
}

func (s Id2MetaVerifyResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s Id2MetaVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *Id2MetaVerifyResponseBodyResultObject) SetBizCode(v string) *Id2MetaVerifyResponseBodyResultObject {
	s.BizCode = &v
	return s
}

type Id2MetaVerifyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Id2MetaVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s Id2MetaVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s Id2MetaVerifyResponse) GoString() string {
	return s.String()
}

func (s *Id2MetaVerifyResponse) SetHeaders(v map[string]*string) *Id2MetaVerifyResponse {
	s.Headers = v
	return s
}

func (s *Id2MetaVerifyResponse) SetStatusCode(v int32) *Id2MetaVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *Id2MetaVerifyResponse) SetBody(v *Id2MetaVerifyResponseBody) *Id2MetaVerifyResponse {
	s.Body = v
	return s
}

type InitFaceVerifyRequest struct {
	AuthId                     *string `json:"AuthId,omitempty" xml:"AuthId,omitempty"`
	Birthday                   *string `json:"Birthday,omitempty" xml:"Birthday,omitempty"`
	CallbackToken              *string `json:"CallbackToken,omitempty" xml:"CallbackToken,omitempty"`
	CallbackUrl                *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	CertName                   *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertNo                     *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	CertType                   *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	CertifyId                  *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	CertifyUrlStyle            *string `json:"CertifyUrlStyle,omitempty" xml:"CertifyUrlStyle,omitempty"`
	CertifyUrlType             *string `json:"CertifyUrlType,omitempty" xml:"CertifyUrlType,omitempty"`
	Crop                       *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	EncryptType                *string `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	FaceContrastPicture        *string `json:"FaceContrastPicture,omitempty" xml:"FaceContrastPicture,omitempty"`
	FaceContrastPictureUrl     *string `json:"FaceContrastPictureUrl,omitempty" xml:"FaceContrastPictureUrl,omitempty"`
	Ip                         *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	MetaInfo                   *string `json:"MetaInfo,omitempty" xml:"MetaInfo,omitempty"`
	Mobile                     *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Mode                       *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Model                      *string `json:"Model,omitempty" xml:"Model,omitempty"`
	OssBucketName              *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	OssObjectName              *string `json:"OssObjectName,omitempty" xml:"OssObjectName,omitempty"`
	OuterOrderNo               *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	ProcedurePriority          *string `json:"ProcedurePriority,omitempty" xml:"ProcedurePriority,omitempty"`
	ProductCode                *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ReadImg                    *string `json:"ReadImg,omitempty" xml:"ReadImg,omitempty"`
	ReturnUrl                  *string `json:"ReturnUrl,omitempty" xml:"ReturnUrl,omitempty"`
	SceneId                    *int64  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	SuitableType               *string `json:"SuitableType,omitempty" xml:"SuitableType,omitempty"`
	UserId                     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ValidityDate               *string `json:"ValidityDate,omitempty" xml:"ValidityDate,omitempty"`
	VoluntaryCustomizedContent *string `json:"VoluntaryCustomizedContent,omitempty" xml:"VoluntaryCustomizedContent,omitempty"`
}

func (s InitFaceVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s InitFaceVerifyRequest) GoString() string {
	return s.String()
}

func (s *InitFaceVerifyRequest) SetAuthId(v string) *InitFaceVerifyRequest {
	s.AuthId = &v
	return s
}

func (s *InitFaceVerifyRequest) SetBirthday(v string) *InitFaceVerifyRequest {
	s.Birthday = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCallbackToken(v string) *InitFaceVerifyRequest {
	s.CallbackToken = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCallbackUrl(v string) *InitFaceVerifyRequest {
	s.CallbackUrl = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCertName(v string) *InitFaceVerifyRequest {
	s.CertName = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCertNo(v string) *InitFaceVerifyRequest {
	s.CertNo = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCertType(v string) *InitFaceVerifyRequest {
	s.CertType = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCertifyId(v string) *InitFaceVerifyRequest {
	s.CertifyId = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCertifyUrlStyle(v string) *InitFaceVerifyRequest {
	s.CertifyUrlStyle = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCertifyUrlType(v string) *InitFaceVerifyRequest {
	s.CertifyUrlType = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCrop(v string) *InitFaceVerifyRequest {
	s.Crop = &v
	return s
}

func (s *InitFaceVerifyRequest) SetEncryptType(v string) *InitFaceVerifyRequest {
	s.EncryptType = &v
	return s
}

func (s *InitFaceVerifyRequest) SetFaceContrastPicture(v string) *InitFaceVerifyRequest {
	s.FaceContrastPicture = &v
	return s
}

func (s *InitFaceVerifyRequest) SetFaceContrastPictureUrl(v string) *InitFaceVerifyRequest {
	s.FaceContrastPictureUrl = &v
	return s
}

func (s *InitFaceVerifyRequest) SetIp(v string) *InitFaceVerifyRequest {
	s.Ip = &v
	return s
}

func (s *InitFaceVerifyRequest) SetMetaInfo(v string) *InitFaceVerifyRequest {
	s.MetaInfo = &v
	return s
}

func (s *InitFaceVerifyRequest) SetMobile(v string) *InitFaceVerifyRequest {
	s.Mobile = &v
	return s
}

func (s *InitFaceVerifyRequest) SetMode(v string) *InitFaceVerifyRequest {
	s.Mode = &v
	return s
}

func (s *InitFaceVerifyRequest) SetModel(v string) *InitFaceVerifyRequest {
	s.Model = &v
	return s
}

func (s *InitFaceVerifyRequest) SetOssBucketName(v string) *InitFaceVerifyRequest {
	s.OssBucketName = &v
	return s
}

func (s *InitFaceVerifyRequest) SetOssObjectName(v string) *InitFaceVerifyRequest {
	s.OssObjectName = &v
	return s
}

func (s *InitFaceVerifyRequest) SetOuterOrderNo(v string) *InitFaceVerifyRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *InitFaceVerifyRequest) SetProcedurePriority(v string) *InitFaceVerifyRequest {
	s.ProcedurePriority = &v
	return s
}

func (s *InitFaceVerifyRequest) SetProductCode(v string) *InitFaceVerifyRequest {
	s.ProductCode = &v
	return s
}

func (s *InitFaceVerifyRequest) SetReadImg(v string) *InitFaceVerifyRequest {
	s.ReadImg = &v
	return s
}

func (s *InitFaceVerifyRequest) SetReturnUrl(v string) *InitFaceVerifyRequest {
	s.ReturnUrl = &v
	return s
}

func (s *InitFaceVerifyRequest) SetSceneId(v int64) *InitFaceVerifyRequest {
	s.SceneId = &v
	return s
}

func (s *InitFaceVerifyRequest) SetSuitableType(v string) *InitFaceVerifyRequest {
	s.SuitableType = &v
	return s
}

func (s *InitFaceVerifyRequest) SetUserId(v string) *InitFaceVerifyRequest {
	s.UserId = &v
	return s
}

func (s *InitFaceVerifyRequest) SetValidityDate(v string) *InitFaceVerifyRequest {
	s.ValidityDate = &v
	return s
}

func (s *InitFaceVerifyRequest) SetVoluntaryCustomizedContent(v string) *InitFaceVerifyRequest {
	s.VoluntaryCustomizedContent = &v
	return s
}

type InitFaceVerifyResponseBody struct {
	Code         *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *InitFaceVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s InitFaceVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitFaceVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *InitFaceVerifyResponseBody) SetCode(v string) *InitFaceVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *InitFaceVerifyResponseBody) SetMessage(v string) *InitFaceVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *InitFaceVerifyResponseBody) SetRequestId(v string) *InitFaceVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitFaceVerifyResponseBody) SetResultObject(v *InitFaceVerifyResponseBodyResultObject) *InitFaceVerifyResponseBody {
	s.ResultObject = v
	return s
}

type InitFaceVerifyResponseBodyResultObject struct {
	CertifyId  *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	CertifyUrl *string `json:"CertifyUrl,omitempty" xml:"CertifyUrl,omitempty"`
}

func (s InitFaceVerifyResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s InitFaceVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *InitFaceVerifyResponseBodyResultObject) SetCertifyId(v string) *InitFaceVerifyResponseBodyResultObject {
	s.CertifyId = &v
	return s
}

func (s *InitFaceVerifyResponseBodyResultObject) SetCertifyUrl(v string) *InitFaceVerifyResponseBodyResultObject {
	s.CertifyUrl = &v
	return s
}

type InitFaceVerifyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InitFaceVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InitFaceVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s InitFaceVerifyResponse) GoString() string {
	return s.String()
}

func (s *InitFaceVerifyResponse) SetHeaders(v map[string]*string) *InitFaceVerifyResponse {
	s.Headers = v
	return s
}

func (s *InitFaceVerifyResponse) SetStatusCode(v int32) *InitFaceVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *InitFaceVerifyResponse) SetBody(v *InitFaceVerifyResponseBody) *InitFaceVerifyResponse {
	s.Body = v
	return s
}

type LivenessFaceVerifyRequest struct {
	CertifyId              *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	Crop                   *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	DeviceToken            *string `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
	FaceContrastPicture    *string `json:"FaceContrastPicture,omitempty" xml:"FaceContrastPicture,omitempty"`
	FaceContrastPictureUrl *string `json:"FaceContrastPictureUrl,omitempty" xml:"FaceContrastPictureUrl,omitempty"`
	Ip                     *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Mobile                 *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Model                  *string `json:"Model,omitempty" xml:"Model,omitempty"`
	OssBucketName          *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	OssObjectName          *string `json:"OssObjectName,omitempty" xml:"OssObjectName,omitempty"`
	OuterOrderNo           *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	ProductCode            *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	SceneId                *int64  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	UserId                 *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s LivenessFaceVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s LivenessFaceVerifyRequest) GoString() string {
	return s.String()
}

func (s *LivenessFaceVerifyRequest) SetCertifyId(v string) *LivenessFaceVerifyRequest {
	s.CertifyId = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetCrop(v string) *LivenessFaceVerifyRequest {
	s.Crop = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetDeviceToken(v string) *LivenessFaceVerifyRequest {
	s.DeviceToken = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetFaceContrastPicture(v string) *LivenessFaceVerifyRequest {
	s.FaceContrastPicture = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetFaceContrastPictureUrl(v string) *LivenessFaceVerifyRequest {
	s.FaceContrastPictureUrl = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetIp(v string) *LivenessFaceVerifyRequest {
	s.Ip = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetMobile(v string) *LivenessFaceVerifyRequest {
	s.Mobile = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetModel(v string) *LivenessFaceVerifyRequest {
	s.Model = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetOssBucketName(v string) *LivenessFaceVerifyRequest {
	s.OssBucketName = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetOssObjectName(v string) *LivenessFaceVerifyRequest {
	s.OssObjectName = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetOuterOrderNo(v string) *LivenessFaceVerifyRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetProductCode(v string) *LivenessFaceVerifyRequest {
	s.ProductCode = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetSceneId(v int64) *LivenessFaceVerifyRequest {
	s.SceneId = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetUserId(v string) *LivenessFaceVerifyRequest {
	s.UserId = &v
	return s
}

type LivenessFaceVerifyResponseBody struct {
	Code         *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *LivenessFaceVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s LivenessFaceVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LivenessFaceVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *LivenessFaceVerifyResponseBody) SetCode(v string) *LivenessFaceVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *LivenessFaceVerifyResponseBody) SetMessage(v string) *LivenessFaceVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *LivenessFaceVerifyResponseBody) SetRequestId(v string) *LivenessFaceVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *LivenessFaceVerifyResponseBody) SetResultObject(v *LivenessFaceVerifyResponseBodyResultObject) *LivenessFaceVerifyResponseBody {
	s.ResultObject = v
	return s
}

type LivenessFaceVerifyResponseBodyResultObject struct {
	CertifyId    *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	MaterialInfo *string `json:"MaterialInfo,omitempty" xml:"MaterialInfo,omitempty"`
	Passed       *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	SubCode      *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
}

func (s LivenessFaceVerifyResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s LivenessFaceVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *LivenessFaceVerifyResponseBodyResultObject) SetCertifyId(v string) *LivenessFaceVerifyResponseBodyResultObject {
	s.CertifyId = &v
	return s
}

func (s *LivenessFaceVerifyResponseBodyResultObject) SetMaterialInfo(v string) *LivenessFaceVerifyResponseBodyResultObject {
	s.MaterialInfo = &v
	return s
}

func (s *LivenessFaceVerifyResponseBodyResultObject) SetPassed(v string) *LivenessFaceVerifyResponseBodyResultObject {
	s.Passed = &v
	return s
}

func (s *LivenessFaceVerifyResponseBodyResultObject) SetSubCode(v string) *LivenessFaceVerifyResponseBodyResultObject {
	s.SubCode = &v
	return s
}

type LivenessFaceVerifyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *LivenessFaceVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LivenessFaceVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s LivenessFaceVerifyResponse) GoString() string {
	return s.String()
}

func (s *LivenessFaceVerifyResponse) SetHeaders(v map[string]*string) *LivenessFaceVerifyResponse {
	s.Headers = v
	return s
}

func (s *LivenessFaceVerifyResponse) SetStatusCode(v int32) *LivenessFaceVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *LivenessFaceVerifyResponse) SetBody(v *LivenessFaceVerifyResponseBody) *LivenessFaceVerifyResponse {
	s.Body = v
	return s
}

type Mobile3MetaDetailVerifyRequest struct {
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	Mobile      *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	ParamType   *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	UserName    *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s Mobile3MetaDetailVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s Mobile3MetaDetailVerifyRequest) GoString() string {
	return s.String()
}

func (s *Mobile3MetaDetailVerifyRequest) SetIdentifyNum(v string) *Mobile3MetaDetailVerifyRequest {
	s.IdentifyNum = &v
	return s
}

func (s *Mobile3MetaDetailVerifyRequest) SetMobile(v string) *Mobile3MetaDetailVerifyRequest {
	s.Mobile = &v
	return s
}

func (s *Mobile3MetaDetailVerifyRequest) SetParamType(v string) *Mobile3MetaDetailVerifyRequest {
	s.ParamType = &v
	return s
}

func (s *Mobile3MetaDetailVerifyRequest) SetUserName(v string) *Mobile3MetaDetailVerifyRequest {
	s.UserName = &v
	return s
}

type Mobile3MetaDetailVerifyResponseBody struct {
	Code         *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *Mobile3MetaDetailVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s Mobile3MetaDetailVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s Mobile3MetaDetailVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *Mobile3MetaDetailVerifyResponseBody) SetCode(v string) *Mobile3MetaDetailVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *Mobile3MetaDetailVerifyResponseBody) SetMessage(v string) *Mobile3MetaDetailVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *Mobile3MetaDetailVerifyResponseBody) SetRequestId(v string) *Mobile3MetaDetailVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *Mobile3MetaDetailVerifyResponseBody) SetResultObject(v *Mobile3MetaDetailVerifyResponseBodyResultObject) *Mobile3MetaDetailVerifyResponseBody {
	s.ResultObject = v
	return s
}

type Mobile3MetaDetailVerifyResponseBodyResultObject struct {
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
}

func (s Mobile3MetaDetailVerifyResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s Mobile3MetaDetailVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *Mobile3MetaDetailVerifyResponseBodyResultObject) SetBizCode(v string) *Mobile3MetaDetailVerifyResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *Mobile3MetaDetailVerifyResponseBodyResultObject) SetIspName(v string) *Mobile3MetaDetailVerifyResponseBodyResultObject {
	s.IspName = &v
	return s
}

func (s *Mobile3MetaDetailVerifyResponseBodyResultObject) SetSubCode(v string) *Mobile3MetaDetailVerifyResponseBodyResultObject {
	s.SubCode = &v
	return s
}

type Mobile3MetaDetailVerifyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Mobile3MetaDetailVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s Mobile3MetaDetailVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s Mobile3MetaDetailVerifyResponse) GoString() string {
	return s.String()
}

func (s *Mobile3MetaDetailVerifyResponse) SetHeaders(v map[string]*string) *Mobile3MetaDetailVerifyResponse {
	s.Headers = v
	return s
}

func (s *Mobile3MetaDetailVerifyResponse) SetStatusCode(v int32) *Mobile3MetaDetailVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *Mobile3MetaDetailVerifyResponse) SetBody(v *Mobile3MetaDetailVerifyResponseBody) *Mobile3MetaDetailVerifyResponse {
	s.Body = v
	return s
}

type Mobile3MetaSimpleVerifyRequest struct {
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	Mobile      *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	ParamType   *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	UserName    *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s Mobile3MetaSimpleVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s Mobile3MetaSimpleVerifyRequest) GoString() string {
	return s.String()
}

func (s *Mobile3MetaSimpleVerifyRequest) SetIdentifyNum(v string) *Mobile3MetaSimpleVerifyRequest {
	s.IdentifyNum = &v
	return s
}

func (s *Mobile3MetaSimpleVerifyRequest) SetMobile(v string) *Mobile3MetaSimpleVerifyRequest {
	s.Mobile = &v
	return s
}

func (s *Mobile3MetaSimpleVerifyRequest) SetParamType(v string) *Mobile3MetaSimpleVerifyRequest {
	s.ParamType = &v
	return s
}

func (s *Mobile3MetaSimpleVerifyRequest) SetUserName(v string) *Mobile3MetaSimpleVerifyRequest {
	s.UserName = &v
	return s
}

type Mobile3MetaSimpleVerifyResponseBody struct {
	Code         *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *Mobile3MetaSimpleVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s Mobile3MetaSimpleVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s Mobile3MetaSimpleVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *Mobile3MetaSimpleVerifyResponseBody) SetCode(v string) *Mobile3MetaSimpleVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *Mobile3MetaSimpleVerifyResponseBody) SetMessage(v string) *Mobile3MetaSimpleVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *Mobile3MetaSimpleVerifyResponseBody) SetRequestId(v string) *Mobile3MetaSimpleVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *Mobile3MetaSimpleVerifyResponseBody) SetResultObject(v *Mobile3MetaSimpleVerifyResponseBodyResultObject) *Mobile3MetaSimpleVerifyResponseBody {
	s.ResultObject = v
	return s
}

type Mobile3MetaSimpleVerifyResponseBodyResultObject struct {
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
}

func (s Mobile3MetaSimpleVerifyResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s Mobile3MetaSimpleVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *Mobile3MetaSimpleVerifyResponseBodyResultObject) SetBizCode(v string) *Mobile3MetaSimpleVerifyResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *Mobile3MetaSimpleVerifyResponseBodyResultObject) SetIspName(v string) *Mobile3MetaSimpleVerifyResponseBodyResultObject {
	s.IspName = &v
	return s
}

type Mobile3MetaSimpleVerifyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Mobile3MetaSimpleVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s Mobile3MetaSimpleVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s Mobile3MetaSimpleVerifyResponse) GoString() string {
	return s.String()
}

func (s *Mobile3MetaSimpleVerifyResponse) SetHeaders(v map[string]*string) *Mobile3MetaSimpleVerifyResponse {
	s.Headers = v
	return s
}

func (s *Mobile3MetaSimpleVerifyResponse) SetStatusCode(v int32) *Mobile3MetaSimpleVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *Mobile3MetaSimpleVerifyResponse) SetBody(v *Mobile3MetaSimpleVerifyResponseBody) *Mobile3MetaSimpleVerifyResponse {
	s.Body = v
	return s
}

type ModifyDeviceInfoRequest struct {
	BizType      *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	DeviceId     *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	Duration     *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	ExpiredDay   *string `json:"ExpiredDay,omitempty" xml:"ExpiredDay,omitempty"`
	UserDeviceId *string `json:"UserDeviceId,omitempty" xml:"UserDeviceId,omitempty"`
}

func (s ModifyDeviceInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDeviceInfoRequest) GoString() string {
	return s.String()
}

func (s *ModifyDeviceInfoRequest) SetBizType(v string) *ModifyDeviceInfoRequest {
	s.BizType = &v
	return s
}

func (s *ModifyDeviceInfoRequest) SetDeviceId(v string) *ModifyDeviceInfoRequest {
	s.DeviceId = &v
	return s
}

func (s *ModifyDeviceInfoRequest) SetDuration(v string) *ModifyDeviceInfoRequest {
	s.Duration = &v
	return s
}

func (s *ModifyDeviceInfoRequest) SetExpiredDay(v string) *ModifyDeviceInfoRequest {
	s.ExpiredDay = &v
	return s
}

func (s *ModifyDeviceInfoRequest) SetUserDeviceId(v string) *ModifyDeviceInfoRequest {
	s.UserDeviceId = &v
	return s
}

type ModifyDeviceInfoResponseBody struct {
	BeginDay     *string `json:"BeginDay,omitempty" xml:"BeginDay,omitempty"`
	BizType      *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	DeviceId     *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	ExpiredDay   *string `json:"ExpiredDay,omitempty" xml:"ExpiredDay,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserDeviceId *string `json:"UserDeviceId,omitempty" xml:"UserDeviceId,omitempty"`
}

func (s ModifyDeviceInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDeviceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDeviceInfoResponseBody) SetBeginDay(v string) *ModifyDeviceInfoResponseBody {
	s.BeginDay = &v
	return s
}

func (s *ModifyDeviceInfoResponseBody) SetBizType(v string) *ModifyDeviceInfoResponseBody {
	s.BizType = &v
	return s
}

func (s *ModifyDeviceInfoResponseBody) SetDeviceId(v string) *ModifyDeviceInfoResponseBody {
	s.DeviceId = &v
	return s
}

func (s *ModifyDeviceInfoResponseBody) SetExpiredDay(v string) *ModifyDeviceInfoResponseBody {
	s.ExpiredDay = &v
	return s
}

func (s *ModifyDeviceInfoResponseBody) SetRequestId(v string) *ModifyDeviceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDeviceInfoResponseBody) SetUserDeviceId(v string) *ModifyDeviceInfoResponseBody {
	s.UserDeviceId = &v
	return s
}

type ModifyDeviceInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDeviceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDeviceInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDeviceInfoResponse) GoString() string {
	return s.String()
}

func (s *ModifyDeviceInfoResponse) SetHeaders(v map[string]*string) *ModifyDeviceInfoResponse {
	s.Headers = v
	return s
}

func (s *ModifyDeviceInfoResponse) SetStatusCode(v int32) *ModifyDeviceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDeviceInfoResponse) SetBody(v *ModifyDeviceInfoResponseBody) *ModifyDeviceInfoResponse {
	s.Body = v
	return s
}

type VerifyMaterialRequest struct {
	BizId               *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType             *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	FaceImageUrl        *string `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty"`
	IdCardBackImageUrl  *string `json:"IdCardBackImageUrl,omitempty" xml:"IdCardBackImageUrl,omitempty"`
	IdCardFrontImageUrl *string `json:"IdCardFrontImageUrl,omitempty" xml:"IdCardFrontImageUrl,omitempty"`
	IdCardNumber        *string `json:"IdCardNumber,omitempty" xml:"IdCardNumber,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	UserId              *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s VerifyMaterialRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyMaterialRequest) GoString() string {
	return s.String()
}

func (s *VerifyMaterialRequest) SetBizId(v string) *VerifyMaterialRequest {
	s.BizId = &v
	return s
}

func (s *VerifyMaterialRequest) SetBizType(v string) *VerifyMaterialRequest {
	s.BizType = &v
	return s
}

func (s *VerifyMaterialRequest) SetFaceImageUrl(v string) *VerifyMaterialRequest {
	s.FaceImageUrl = &v
	return s
}

func (s *VerifyMaterialRequest) SetIdCardBackImageUrl(v string) *VerifyMaterialRequest {
	s.IdCardBackImageUrl = &v
	return s
}

func (s *VerifyMaterialRequest) SetIdCardFrontImageUrl(v string) *VerifyMaterialRequest {
	s.IdCardFrontImageUrl = &v
	return s
}

func (s *VerifyMaterialRequest) SetIdCardNumber(v string) *VerifyMaterialRequest {
	s.IdCardNumber = &v
	return s
}

func (s *VerifyMaterialRequest) SetName(v string) *VerifyMaterialRequest {
	s.Name = &v
	return s
}

func (s *VerifyMaterialRequest) SetUserId(v string) *VerifyMaterialRequest {
	s.UserId = &v
	return s
}

type VerifyMaterialResponseBody struct {
	AuthorityComparisionScore *float32                            `json:"AuthorityComparisionScore,omitempty" xml:"AuthorityComparisionScore,omitempty"`
	IdCardFaceComparisonScore *float32                            `json:"IdCardFaceComparisonScore,omitempty" xml:"IdCardFaceComparisonScore,omitempty"`
	Material                  *VerifyMaterialResponseBodyMaterial `json:"Material,omitempty" xml:"Material,omitempty" type:"Struct"`
	RequestId                 *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VerifyStatus              *int32                              `json:"VerifyStatus,omitempty" xml:"VerifyStatus,omitempty"`
	VerifyToken               *string                             `json:"VerifyToken,omitempty" xml:"VerifyToken,omitempty"`
}

func (s VerifyMaterialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyMaterialResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyMaterialResponseBody) SetAuthorityComparisionScore(v float32) *VerifyMaterialResponseBody {
	s.AuthorityComparisionScore = &v
	return s
}

func (s *VerifyMaterialResponseBody) SetIdCardFaceComparisonScore(v float32) *VerifyMaterialResponseBody {
	s.IdCardFaceComparisonScore = &v
	return s
}

func (s *VerifyMaterialResponseBody) SetMaterial(v *VerifyMaterialResponseBodyMaterial) *VerifyMaterialResponseBody {
	s.Material = v
	return s
}

func (s *VerifyMaterialResponseBody) SetRequestId(v string) *VerifyMaterialResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyMaterialResponseBody) SetVerifyStatus(v int32) *VerifyMaterialResponseBody {
	s.VerifyStatus = &v
	return s
}

func (s *VerifyMaterialResponseBody) SetVerifyToken(v string) *VerifyMaterialResponseBody {
	s.VerifyToken = &v
	return s
}

type VerifyMaterialResponseBodyMaterial struct {
	FaceGlobalUrl *string                                       `json:"FaceGlobalUrl,omitempty" xml:"FaceGlobalUrl,omitempty"`
	FaceImageUrl  *string                                       `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty"`
	FaceMask      *string                                       `json:"FaceMask,omitempty" xml:"FaceMask,omitempty"`
	FaceQuality   *string                                       `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty"`
	IdCardInfo    *VerifyMaterialResponseBodyMaterialIdCardInfo `json:"IdCardInfo,omitempty" xml:"IdCardInfo,omitempty" type:"Struct"`
	IdCardName    *string                                       `json:"IdCardName,omitempty" xml:"IdCardName,omitempty"`
	IdCardNumber  *string                                       `json:"IdCardNumber,omitempty" xml:"IdCardNumber,omitempty"`
}

func (s VerifyMaterialResponseBodyMaterial) String() string {
	return tea.Prettify(s)
}

func (s VerifyMaterialResponseBodyMaterial) GoString() string {
	return s.String()
}

func (s *VerifyMaterialResponseBodyMaterial) SetFaceGlobalUrl(v string) *VerifyMaterialResponseBodyMaterial {
	s.FaceGlobalUrl = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterial) SetFaceImageUrl(v string) *VerifyMaterialResponseBodyMaterial {
	s.FaceImageUrl = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterial) SetFaceMask(v string) *VerifyMaterialResponseBodyMaterial {
	s.FaceMask = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterial) SetFaceQuality(v string) *VerifyMaterialResponseBodyMaterial {
	s.FaceQuality = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterial) SetIdCardInfo(v *VerifyMaterialResponseBodyMaterialIdCardInfo) *VerifyMaterialResponseBodyMaterial {
	s.IdCardInfo = v
	return s
}

func (s *VerifyMaterialResponseBodyMaterial) SetIdCardName(v string) *VerifyMaterialResponseBodyMaterial {
	s.IdCardName = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterial) SetIdCardNumber(v string) *VerifyMaterialResponseBodyMaterial {
	s.IdCardNumber = &v
	return s
}

type VerifyMaterialResponseBodyMaterialIdCardInfo struct {
	Address       *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Authority     *string `json:"Authority,omitempty" xml:"Authority,omitempty"`
	BackImageUrl  *string `json:"BackImageUrl,omitempty" xml:"BackImageUrl,omitempty"`
	Birth         *string `json:"Birth,omitempty" xml:"Birth,omitempty"`
	EndDate       *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	FrontImageUrl *string `json:"FrontImageUrl,omitempty" xml:"FrontImageUrl,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Nationality   *string `json:"Nationality,omitempty" xml:"Nationality,omitempty"`
	Number        *string `json:"Number,omitempty" xml:"Number,omitempty"`
	StartDate     *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s VerifyMaterialResponseBodyMaterialIdCardInfo) String() string {
	return tea.Prettify(s)
}

func (s VerifyMaterialResponseBodyMaterialIdCardInfo) GoString() string {
	return s.String()
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) SetAddress(v string) *VerifyMaterialResponseBodyMaterialIdCardInfo {
	s.Address = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) SetAuthority(v string) *VerifyMaterialResponseBodyMaterialIdCardInfo {
	s.Authority = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) SetBackImageUrl(v string) *VerifyMaterialResponseBodyMaterialIdCardInfo {
	s.BackImageUrl = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) SetBirth(v string) *VerifyMaterialResponseBodyMaterialIdCardInfo {
	s.Birth = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) SetEndDate(v string) *VerifyMaterialResponseBodyMaterialIdCardInfo {
	s.EndDate = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) SetFrontImageUrl(v string) *VerifyMaterialResponseBodyMaterialIdCardInfo {
	s.FrontImageUrl = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) SetName(v string) *VerifyMaterialResponseBodyMaterialIdCardInfo {
	s.Name = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) SetNationality(v string) *VerifyMaterialResponseBodyMaterialIdCardInfo {
	s.Nationality = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) SetNumber(v string) *VerifyMaterialResponseBodyMaterialIdCardInfo {
	s.Number = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) SetStartDate(v string) *VerifyMaterialResponseBodyMaterialIdCardInfo {
	s.StartDate = &v
	return s
}

type VerifyMaterialResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *VerifyMaterialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VerifyMaterialResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyMaterialResponse) GoString() string {
	return s.String()
}

func (s *VerifyMaterialResponse) SetHeaders(v map[string]*string) *VerifyMaterialResponse {
	s.Headers = v
	return s
}

func (s *VerifyMaterialResponse) SetStatusCode(v int32) *VerifyMaterialResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyMaterialResponse) SetBody(v *VerifyMaterialResponseBody) *VerifyMaterialResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("cloudauth"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CompareFaceVerifyWithOptions(request *CompareFaceVerifyRequest, runtime *util.RuntimeOptions) (_result *CompareFaceVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Crop)) {
		body["Crop"] = request.Crop
	}

	if !tea.BoolValue(util.IsUnset(request.OuterOrderNo)) {
		body["OuterOrderNo"] = request.OuterOrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		body["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceCertifyId)) {
		body["SourceCertifyId"] = request.SourceCertifyId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceFaceContrastPicture)) {
		body["SourceFaceContrastPicture"] = request.SourceFaceContrastPicture
	}

	if !tea.BoolValue(util.IsUnset(request.SourceFaceContrastPictureUrl)) {
		body["SourceFaceContrastPictureUrl"] = request.SourceFaceContrastPictureUrl
	}

	if !tea.BoolValue(util.IsUnset(request.SourceOssBucketName)) {
		body["SourceOssBucketName"] = request.SourceOssBucketName
	}

	if !tea.BoolValue(util.IsUnset(request.SourceOssObjectName)) {
		body["SourceOssObjectName"] = request.SourceOssObjectName
	}

	if !tea.BoolValue(util.IsUnset(request.TargetCertifyId)) {
		body["TargetCertifyId"] = request.TargetCertifyId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetFaceContrastPicture)) {
		body["TargetFaceContrastPicture"] = request.TargetFaceContrastPicture
	}

	if !tea.BoolValue(util.IsUnset(request.TargetFaceContrastPictureUrl)) {
		body["TargetFaceContrastPictureUrl"] = request.TargetFaceContrastPictureUrl
	}

	if !tea.BoolValue(util.IsUnset(request.TargetOssBucketName)) {
		body["TargetOssBucketName"] = request.TargetOssBucketName
	}

	if !tea.BoolValue(util.IsUnset(request.TargetOssObjectName)) {
		body["TargetOssObjectName"] = request.TargetOssObjectName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CompareFaceVerify"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CompareFaceVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CompareFaceVerify(request *CompareFaceVerifyRequest) (_result *CompareFaceVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CompareFaceVerifyResponse{}
	_body, _err := client.CompareFaceVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CompareFacesWithOptions(request *CompareFacesRequest, runtime *util.RuntimeOptions) (_result *CompareFacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SourceImageType)) {
		body["SourceImageType"] = request.SourceImageType
	}

	if !tea.BoolValue(util.IsUnset(request.SourceImageValue)) {
		body["SourceImageValue"] = request.SourceImageValue
	}

	if !tea.BoolValue(util.IsUnset(request.TargetImageType)) {
		body["TargetImageType"] = request.TargetImageType
	}

	if !tea.BoolValue(util.IsUnset(request.TargetImageValue)) {
		body["TargetImageValue"] = request.TargetImageValue
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CompareFaces"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CompareFacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CompareFaces(request *CompareFacesRequest) (_result *CompareFacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CompareFacesResponse{}
	_body, _err := client.CompareFacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ContrastFaceVerifyWithOptions(request *ContrastFaceVerifyRequest, runtime *util.RuntimeOptions) (_result *ContrastFaceVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Model)) {
		query["Model"] = request.Model
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertName)) {
		body["CertName"] = request.CertName
	}

	if !tea.BoolValue(util.IsUnset(request.CertNo)) {
		body["CertNo"] = request.CertNo
	}

	if !tea.BoolValue(util.IsUnset(request.CertType)) {
		body["CertType"] = request.CertType
	}

	if !tea.BoolValue(util.IsUnset(request.CertifyId)) {
		body["CertifyId"] = request.CertifyId
	}

	if !tea.BoolValue(util.IsUnset(request.Crop)) {
		body["Crop"] = request.Crop
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceToken)) {
		body["DeviceToken"] = request.DeviceToken
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptType)) {
		body["EncryptType"] = request.EncryptType
	}

	if !tea.BoolValue(util.IsUnset(request.FaceContrastFile)) {
		body["FaceContrastFile"] = request.FaceContrastFile
	}

	if !tea.BoolValue(util.IsUnset(request.FaceContrastPicture)) {
		body["FaceContrastPicture"] = request.FaceContrastPicture
	}

	if !tea.BoolValue(util.IsUnset(request.FaceContrastPictureUrl)) {
		body["FaceContrastPictureUrl"] = request.FaceContrastPictureUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		body["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		body["Mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.OssBucketName)) {
		body["OssBucketName"] = request.OssBucketName
	}

	if !tea.BoolValue(util.IsUnset(request.OssObjectName)) {
		body["OssObjectName"] = request.OssObjectName
	}

	if !tea.BoolValue(util.IsUnset(request.OuterOrderNo)) {
		body["OuterOrderNo"] = request.OuterOrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		body["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ContrastFaceVerify"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ContrastFaceVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ContrastFaceVerify(request *ContrastFaceVerifyRequest) (_result *ContrastFaceVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ContrastFaceVerifyResponse{}
	_body, _err := client.ContrastFaceVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ContrastFaceVerifyAdvance(request *ContrastFaceVerifyAdvanceRequest, runtime *util.RuntimeOptions) (_result *ContrastFaceVerifyResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("Cloudauth"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	contrastFaceVerifyReq := &ContrastFaceVerifyRequest{}
	openapiutil.Convert(request, contrastFaceVerifyReq)
	if !tea.BoolValue(util.IsUnset(request.FaceContrastFileObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.FaceContrastFileObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		contrastFaceVerifyReq.FaceContrastFile = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	contrastFaceVerifyResp, _err := client.ContrastFaceVerifyWithOptions(contrastFaceVerifyReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = contrastFaceVerifyResp
	return _result, _err
}

func (client *Client) CreateAuthKeyWithOptions(request *CreateAuthKeyRequest, runtime *util.RuntimeOptions) (_result *CreateAuthKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthYears)) {
		query["AuthYears"] = request.AuthYears
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.Test)) {
		query["Test"] = request.Test
	}

	if !tea.BoolValue(util.IsUnset(request.UserDeviceId)) {
		query["UserDeviceId"] = request.UserDeviceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAuthKey"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAuthKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAuthKey(request *CreateAuthKeyRequest) (_result *CreateAuthKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAuthKeyResponse{}
	_body, _err := client.CreateAuthKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVerifySettingWithOptions(request *CreateVerifySettingRequest, runtime *util.RuntimeOptions) (_result *CreateVerifySettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizName)) {
		query["BizName"] = request.BizName
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.GuideStep)) {
		query["GuideStep"] = request.GuideStep
	}

	if !tea.BoolValue(util.IsUnset(request.PrivacyStep)) {
		query["PrivacyStep"] = request.PrivacyStep
	}

	if !tea.BoolValue(util.IsUnset(request.ResultStep)) {
		query["ResultStep"] = request.ResultStep
	}

	if !tea.BoolValue(util.IsUnset(request.Solution)) {
		query["Solution"] = request.Solution
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVerifySetting"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVerifySettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVerifySetting(request *CreateVerifySettingRequest) (_result *CreateVerifySettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVerifySettingResponse{}
	_body, _err := client.CreateVerifySettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDeviceInfoWithOptions(request *DescribeDeviceInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeDeviceInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.ExpiredEndDay)) {
		query["ExpiredEndDay"] = request.ExpiredEndDay
	}

	if !tea.BoolValue(util.IsUnset(request.ExpiredStartDay)) {
		query["ExpiredStartDay"] = request.ExpiredStartDay
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.UserDeviceId)) {
		query["UserDeviceId"] = request.UserDeviceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDeviceInfo"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDeviceInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDeviceInfo(request *DescribeDeviceInfoRequest) (_result *DescribeDeviceInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDeviceInfoResponse{}
	_body, _err := client.DescribeDeviceInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFaceVerifyWithOptions(request *DescribeFaceVerifyRequest, runtime *util.RuntimeOptions) (_result *DescribeFaceVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertifyId)) {
		query["CertifyId"] = request.CertifyId
	}

	if !tea.BoolValue(util.IsUnset(request.PictureReturnType)) {
		query["PictureReturnType"] = request.PictureReturnType
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFaceVerify"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFaceVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFaceVerify(request *DescribeFaceVerifyRequest) (_result *DescribeFaceVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFaceVerifyResponse{}
	_body, _err := client.DescribeFaceVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOssUploadTokenWithOptions(runtime *util.RuntimeOptions) (_result *DescribeOssUploadTokenResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeOssUploadToken"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOssUploadTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOssUploadToken() (_result *DescribeOssUploadTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOssUploadTokenResponse{}
	_body, _err := client.DescribeOssUploadTokenWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVerifyResultWithOptions(request *DescribeVerifyResultRequest, runtime *util.RuntimeOptions) (_result *DescribeVerifyResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVerifyResult"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVerifyResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVerifyResult(request *DescribeVerifyResultRequest) (_result *DescribeVerifyResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVerifyResultResponse{}
	_body, _err := client.DescribeVerifyResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVerifySDKWithOptions(request *DescribeVerifySDKRequest, runtime *util.RuntimeOptions) (_result *DescribeVerifySDKResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVerifySDK"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVerifySDKResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVerifySDK(request *DescribeVerifySDKRequest) (_result *DescribeVerifySDKResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVerifySDKResponse{}
	_body, _err := client.DescribeVerifySDKWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVerifyTokenWithOptions(request *DescribeVerifyTokenRequest, runtime *util.RuntimeOptions) (_result *DescribeVerifyTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackSeed)) {
		query["CallbackSeed"] = request.CallbackSeed
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackUrl)) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	if !tea.BoolValue(util.IsUnset(request.FaceRetainedImageUrl)) {
		query["FaceRetainedImageUrl"] = request.FaceRetainedImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.FailedRedirectUrl)) {
		query["FailedRedirectUrl"] = request.FailedRedirectUrl
	}

	if !tea.BoolValue(util.IsUnset(request.IdCardBackImageUrl)) {
		query["IdCardBackImageUrl"] = request.IdCardBackImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.IdCardFrontImageUrl)) {
		query["IdCardFrontImageUrl"] = request.IdCardFrontImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.IdCardNumber)) {
		query["IdCardNumber"] = request.IdCardNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PassedRedirectUrl)) {
		query["PassedRedirectUrl"] = request.PassedRedirectUrl
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserIp)) {
		query["UserIp"] = request.UserIp
	}

	if !tea.BoolValue(util.IsUnset(request.UserPhoneNumber)) {
		query["UserPhoneNumber"] = request.UserPhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.UserRegistTime)) {
		query["UserRegistTime"] = request.UserRegistTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVerifyToken"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVerifyTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVerifyToken(request *DescribeVerifyTokenRequest) (_result *DescribeVerifyTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVerifyTokenResponse{}
	_body, _err := client.DescribeVerifyTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectFaceAttributesWithOptions(request *DetectFaceAttributesRequest, runtime *util.RuntimeOptions) (_result *DetectFaceAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		body["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.MaterialValue)) {
		body["MaterialValue"] = request.MaterialValue
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectFaceAttributes"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectFaceAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectFaceAttributes(request *DetectFaceAttributesRequest) (_result *DetectFaceAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectFaceAttributesResponse{}
	_body, _err := client.DetectFaceAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) Id2MetaVerifyWithOptions(request *Id2MetaVerifyRequest, runtime *util.RuntimeOptions) (_result *Id2MetaVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IdentifyNum)) {
		query["IdentifyNum"] = request.IdentifyNum
	}

	if !tea.BoolValue(util.IsUnset(request.ParamType)) {
		query["ParamType"] = request.ParamType
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Id2MetaVerify"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &Id2MetaVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Id2MetaVerify(request *Id2MetaVerifyRequest) (_result *Id2MetaVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &Id2MetaVerifyResponse{}
	_body, _err := client.Id2MetaVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InitFaceVerifyWithOptions(request *InitFaceVerifyRequest, runtime *util.RuntimeOptions) (_result *InitFaceVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Birthday)) {
		query["Birthday"] = request.Birthday
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackToken)) {
		query["CallbackToken"] = request.CallbackToken
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackUrl)) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	if !tea.BoolValue(util.IsUnset(request.CertName)) {
		query["CertName"] = request.CertName
	}

	if !tea.BoolValue(util.IsUnset(request.CertNo)) {
		query["CertNo"] = request.CertNo
	}

	if !tea.BoolValue(util.IsUnset(request.CertType)) {
		query["CertType"] = request.CertType
	}

	if !tea.BoolValue(util.IsUnset(request.CertifyId)) {
		query["CertifyId"] = request.CertifyId
	}

	if !tea.BoolValue(util.IsUnset(request.CertifyUrlStyle)) {
		query["CertifyUrlStyle"] = request.CertifyUrlStyle
	}

	if !tea.BoolValue(util.IsUnset(request.CertifyUrlType)) {
		query["CertifyUrlType"] = request.CertifyUrlType
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptType)) {
		query["EncryptType"] = request.EncryptType
	}

	if !tea.BoolValue(util.IsUnset(request.FaceContrastPictureUrl)) {
		query["FaceContrastPictureUrl"] = request.FaceContrastPictureUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.MetaInfo)) {
		query["MetaInfo"] = request.MetaInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		query["Mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.OssBucketName)) {
		query["OssBucketName"] = request.OssBucketName
	}

	if !tea.BoolValue(util.IsUnset(request.OssObjectName)) {
		query["OssObjectName"] = request.OssObjectName
	}

	if !tea.BoolValue(util.IsUnset(request.OuterOrderNo)) {
		query["OuterOrderNo"] = request.OuterOrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.ProcedurePriority)) {
		query["ProcedurePriority"] = request.ProcedurePriority
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.ReadImg)) {
		query["ReadImg"] = request.ReadImg
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnUrl)) {
		query["ReturnUrl"] = request.ReturnUrl
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.SuitableType)) {
		query["SuitableType"] = request.SuitableType
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.ValidityDate)) {
		query["ValidityDate"] = request.ValidityDate
	}

	if !tea.BoolValue(util.IsUnset(request.VoluntaryCustomizedContent)) {
		query["VoluntaryCustomizedContent"] = request.VoluntaryCustomizedContent
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthId)) {
		body["AuthId"] = request.AuthId
	}

	if !tea.BoolValue(util.IsUnset(request.Crop)) {
		body["Crop"] = request.Crop
	}

	if !tea.BoolValue(util.IsUnset(request.FaceContrastPicture)) {
		body["FaceContrastPicture"] = request.FaceContrastPicture
	}

	if !tea.BoolValue(util.IsUnset(request.Model)) {
		body["Model"] = request.Model
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InitFaceVerify"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InitFaceVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InitFaceVerify(request *InitFaceVerifyRequest) (_result *InitFaceVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InitFaceVerifyResponse{}
	_body, _err := client.InitFaceVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LivenessFaceVerifyWithOptions(request *LivenessFaceVerifyRequest, runtime *util.RuntimeOptions) (_result *LivenessFaceVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Model)) {
		query["Model"] = request.Model
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertifyId)) {
		body["CertifyId"] = request.CertifyId
	}

	if !tea.BoolValue(util.IsUnset(request.Crop)) {
		body["Crop"] = request.Crop
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceToken)) {
		body["DeviceToken"] = request.DeviceToken
	}

	if !tea.BoolValue(util.IsUnset(request.FaceContrastPicture)) {
		body["FaceContrastPicture"] = request.FaceContrastPicture
	}

	if !tea.BoolValue(util.IsUnset(request.FaceContrastPictureUrl)) {
		body["FaceContrastPictureUrl"] = request.FaceContrastPictureUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		body["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		body["Mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.OssBucketName)) {
		body["OssBucketName"] = request.OssBucketName
	}

	if !tea.BoolValue(util.IsUnset(request.OssObjectName)) {
		body["OssObjectName"] = request.OssObjectName
	}

	if !tea.BoolValue(util.IsUnset(request.OuterOrderNo)) {
		body["OuterOrderNo"] = request.OuterOrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		body["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("LivenessFaceVerify"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LivenessFaceVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LivenessFaceVerify(request *LivenessFaceVerifyRequest) (_result *LivenessFaceVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LivenessFaceVerifyResponse{}
	_body, _err := client.LivenessFaceVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) Mobile3MetaDetailVerifyWithOptions(request *Mobile3MetaDetailVerifyRequest, runtime *util.RuntimeOptions) (_result *Mobile3MetaDetailVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IdentifyNum)) {
		query["IdentifyNum"] = request.IdentifyNum
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		query["Mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.ParamType)) {
		query["ParamType"] = request.ParamType
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Mobile3MetaDetailVerify"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &Mobile3MetaDetailVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Mobile3MetaDetailVerify(request *Mobile3MetaDetailVerifyRequest) (_result *Mobile3MetaDetailVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &Mobile3MetaDetailVerifyResponse{}
	_body, _err := client.Mobile3MetaDetailVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) Mobile3MetaSimpleVerifyWithOptions(request *Mobile3MetaSimpleVerifyRequest, runtime *util.RuntimeOptions) (_result *Mobile3MetaSimpleVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IdentifyNum)) {
		query["IdentifyNum"] = request.IdentifyNum
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		query["Mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.ParamType)) {
		query["ParamType"] = request.ParamType
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Mobile3MetaSimpleVerify"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &Mobile3MetaSimpleVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Mobile3MetaSimpleVerify(request *Mobile3MetaSimpleVerifyRequest) (_result *Mobile3MetaSimpleVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &Mobile3MetaSimpleVerifyResponse{}
	_body, _err := client.Mobile3MetaSimpleVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDeviceInfoWithOptions(request *ModifyDeviceInfoRequest, runtime *util.RuntimeOptions) (_result *ModifyDeviceInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.ExpiredDay)) {
		query["ExpiredDay"] = request.ExpiredDay
	}

	if !tea.BoolValue(util.IsUnset(request.UserDeviceId)) {
		query["UserDeviceId"] = request.UserDeviceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDeviceInfo"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDeviceInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDeviceInfo(request *ModifyDeviceInfoRequest) (_result *ModifyDeviceInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDeviceInfoResponse{}
	_body, _err := client.ModifyDeviceInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyMaterialWithOptions(request *VerifyMaterialRequest, runtime *util.RuntimeOptions) (_result *VerifyMaterialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.FaceImageUrl)) {
		query["FaceImageUrl"] = request.FaceImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.IdCardBackImageUrl)) {
		query["IdCardBackImageUrl"] = request.IdCardBackImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.IdCardFrontImageUrl)) {
		query["IdCardFrontImageUrl"] = request.IdCardFrontImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.IdCardNumber)) {
		query["IdCardNumber"] = request.IdCardNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyMaterial"),
		Version:     tea.String("2019-03-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyMaterialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyMaterial(request *VerifyMaterialRequest) (_result *VerifyMaterialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyMaterialResponse{}
	_body, _err := client.VerifyMaterialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
