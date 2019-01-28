package face

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/satori/go.uuid"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v1.0/face"

// AccessoryType enumerates the values for accessory type.
type AccessoryType string

const (
	// Glasses ...
	Glasses AccessoryType = "glasses"
	// HeadWear ...
	HeadWear AccessoryType = "headWear"
	// Mask ...
	Mask AccessoryType = "mask"
)

// PossibleAccessoryTypeValues returns an array of possible values for the AccessoryType const type.
func PossibleAccessoryTypeValues() []AccessoryType {
	return []AccessoryType{Glasses, HeadWear, Mask}
}

// AttributeType enumerates the values for attribute type.
type AttributeType string

const (
	// AttributeTypeAccessories ...
	AttributeTypeAccessories AttributeType = "accessories"
	// AttributeTypeAge ...
	AttributeTypeAge AttributeType = "age"
	// AttributeTypeBlur ...
	AttributeTypeBlur AttributeType = "blur"
	// AttributeTypeEmotion ...
	AttributeTypeEmotion AttributeType = "emotion"
	// AttributeTypeExposure ...
	AttributeTypeExposure AttributeType = "exposure"
	// AttributeTypeFacialHair ...
	AttributeTypeFacialHair AttributeType = "facialHair"
	// AttributeTypeGender ...
	AttributeTypeGender AttributeType = "gender"
	// AttributeTypeGlasses ...
	AttributeTypeGlasses AttributeType = "glasses"
	// AttributeTypeHair ...
	AttributeTypeHair AttributeType = "hair"
	// AttributeTypeHeadPose ...
	AttributeTypeHeadPose AttributeType = "headPose"
	// AttributeTypeMakeup ...
	AttributeTypeMakeup AttributeType = "makeup"
	// AttributeTypeNoise ...
	AttributeTypeNoise AttributeType = "noise"
	// AttributeTypeOcclusion ...
	AttributeTypeOcclusion AttributeType = "occlusion"
	// AttributeTypeSmile ...
	AttributeTypeSmile AttributeType = "smile"
)

// PossibleAttributeTypeValues returns an array of possible values for the AttributeType const type.
func PossibleAttributeTypeValues() []AttributeType {
	return []AttributeType{AttributeTypeAccessories, AttributeTypeAge, AttributeTypeBlur, AttributeTypeEmotion, AttributeTypeExposure, AttributeTypeFacialHair, AttributeTypeGender, AttributeTypeGlasses, AttributeTypeHair, AttributeTypeHeadPose, AttributeTypeMakeup, AttributeTypeNoise, AttributeTypeOcclusion, AttributeTypeSmile}
}

// BlurLevel enumerates the values for blur level.
type BlurLevel string

const (
	// High ...
	High BlurLevel = "High"
	// Low ...
	Low BlurLevel = "Low"
	// Medium ...
	Medium BlurLevel = "Medium"
)

// PossibleBlurLevelValues returns an array of possible values for the BlurLevel const type.
func PossibleBlurLevelValues() []BlurLevel {
	return []BlurLevel{High, Low, Medium}
}

// ExposureLevel enumerates the values for exposure level.
type ExposureLevel string

const (
	// GoodExposure ...
	GoodExposure ExposureLevel = "GoodExposure"
	// OverExposure ...
	OverExposure ExposureLevel = "OverExposure"
	// UnderExposure ...
	UnderExposure ExposureLevel = "UnderExposure"
)

// PossibleExposureLevelValues returns an array of possible values for the ExposureLevel const type.
func PossibleExposureLevelValues() []ExposureLevel {
	return []ExposureLevel{GoodExposure, OverExposure, UnderExposure}
}

// FindSimilarMatchMode enumerates the values for find similar match mode.
type FindSimilarMatchMode string

const (
	// MatchFace ...
	MatchFace FindSimilarMatchMode = "matchFace"
	// MatchPerson ...
	MatchPerson FindSimilarMatchMode = "matchPerson"
)

// PossibleFindSimilarMatchModeValues returns an array of possible values for the FindSimilarMatchMode const type.
func PossibleFindSimilarMatchModeValues() []FindSimilarMatchMode {
	return []FindSimilarMatchMode{MatchFace, MatchPerson}
}

// Gender enumerates the values for gender.
type Gender string

const (
	// Female ...
	Female Gender = "female"
	// Genderless ...
	Genderless Gender = "genderless"
	// Male ...
	Male Gender = "male"
)

// PossibleGenderValues returns an array of possible values for the Gender const type.
func PossibleGenderValues() []Gender {
	return []Gender{Female, Genderless, Male}
}

// GlassesType enumerates the values for glasses type.
type GlassesType string

const (
	// NoGlasses ...
	NoGlasses GlassesType = "noGlasses"
	// ReadingGlasses ...
	ReadingGlasses GlassesType = "readingGlasses"
	// Sunglasses ...
	Sunglasses GlassesType = "sunglasses"
	// SwimmingGoggles ...
	SwimmingGoggles GlassesType = "swimmingGoggles"
)

// PossibleGlassesTypeValues returns an array of possible values for the GlassesType const type.
func PossibleGlassesTypeValues() []GlassesType {
	return []GlassesType{NoGlasses, ReadingGlasses, Sunglasses, SwimmingGoggles}
}

// HairColorType enumerates the values for hair color type.
type HairColorType string

const (
	// Black ...
	Black HairColorType = "black"
	// Blond ...
	Blond HairColorType = "blond"
	// Brown ...
	Brown HairColorType = "brown"
	// Gray ...
	Gray HairColorType = "gray"
	// Other ...
	Other HairColorType = "other"
	// Red ...
	Red HairColorType = "red"
	// Unknown ...
	Unknown HairColorType = "unknown"
	// White ...
	White HairColorType = "white"
)

// PossibleHairColorTypeValues returns an array of possible values for the HairColorType const type.
func PossibleHairColorTypeValues() []HairColorType {
	return []HairColorType{Black, Blond, Brown, Gray, Other, Red, Unknown, White}
}

// NoiseLevel enumerates the values for noise level.
type NoiseLevel string

const (
	// NoiseLevelHigh ...
	NoiseLevelHigh NoiseLevel = "High"
	// NoiseLevelLow ...
	NoiseLevelLow NoiseLevel = "Low"
	// NoiseLevelMedium ...
	NoiseLevelMedium NoiseLevel = "Medium"
)

// PossibleNoiseLevelValues returns an array of possible values for the NoiseLevel const type.
func PossibleNoiseLevelValues() []NoiseLevel {
	return []NoiseLevel{NoiseLevelHigh, NoiseLevelLow, NoiseLevelMedium}
}

// TrainingStatusType enumerates the values for training status type.
type TrainingStatusType string

const (
	// Failed ...
	Failed TrainingStatusType = "failed"
	// Nonstarted ...
	Nonstarted TrainingStatusType = "nonstarted"
	// Running ...
	Running TrainingStatusType = "running"
	// Succeeded ...
	Succeeded TrainingStatusType = "succeeded"
)

// PossibleTrainingStatusTypeValues returns an array of possible values for the TrainingStatusType const type.
func PossibleTrainingStatusTypeValues() []TrainingStatusType {
	return []TrainingStatusType{Failed, Nonstarted, Running, Succeeded}
}

// Accessory accessory item and corresponding confidence level.
type Accessory struct {
	// Type - Type of an accessory. Possible values include: 'HeadWear', 'Glasses', 'Mask'
	Type AccessoryType `json:"type,omitempty"`
	// Confidence - Confidence level of an accessory
	Confidence *float64 `json:"confidence,omitempty"`
}

// APIError error information returned by the API
type APIError struct {
	Error *Error `json:"error,omitempty"`
}

// Attributes face Attributes
type Attributes struct {
	// Age - Age in years
	Age *float64 `json:"age,omitempty"`
	// Gender - Possible gender of the face. Possible values include: 'Male', 'Female', 'Genderless'
	Gender Gender `json:"gender,omitempty"`
	// Smile - Smile intensity, a number between [0,1]
	Smile *float64 `json:"smile,omitempty"`
	// FacialHair - Properties describing facial hair attributes.
	FacialHair *FacialHair `json:"facialHair,omitempty"`
	// Glasses - Glasses type if any of the face. Possible values include: 'NoGlasses', 'ReadingGlasses', 'Sunglasses', 'SwimmingGoggles'
	Glasses GlassesType `json:"glasses,omitempty"`
	// HeadPose - Properties indicating head pose of the face.
	HeadPose *HeadPose `json:"headPose,omitempty"`
	// Emotion - Properties describing facial emotion in form of confidence ranging from 0 to 1.
	Emotion *Emotion `json:"emotion,omitempty"`
	// Hair - Properties describing hair attributes.
	Hair *Hair `json:"hair,omitempty"`
	// Makeup - Properties describing present makeups on a given face.
	Makeup *Makeup `json:"makeup,omitempty"`
	// Occlusion - Properties describing occlusions on a given face.
	Occlusion *Occlusion `json:"occlusion,omitempty"`
	// Accessories - Properties describing any accessories on a given face.
	Accessories *[]Accessory `json:"accessories,omitempty"`
	// Blur - Properties describing any presence of blur within the image.
	Blur *Blur `json:"blur,omitempty"`
	// Exposure - Properties describing exposure level of the image.
	Exposure *Exposure `json:"exposure,omitempty"`
	// Noise - Properties describing noise level of the image.
	Noise *Noise `json:"noise,omitempty"`
}

// Blur properties describing any presence of blur within the image.
type Blur struct {
	// BlurLevel - An enum value indicating level of blurriness. Possible values include: 'Low', 'Medium', 'High'
	BlurLevel BlurLevel `json:"blurLevel,omitempty"`
	// Value - A number indicating level of blurriness ranging from 0 to 1.
	Value *float64 `json:"value,omitempty"`
}

// Coordinate coordinates within an image
type Coordinate struct {
	// X - The horizontal component, in pixels.
	X *float64 `json:"x,omitempty"`
	// Y - The vertical component, in pixels.
	Y *float64 `json:"y,omitempty"`
}

// DetectedFace detected Face object.
type DetectedFace struct {
	FaceID         *uuid.UUID  `json:"faceId,omitempty"`
	FaceRectangle  *Rectangle  `json:"faceRectangle,omitempty"`
	FaceLandmarks  *Landmarks  `json:"faceLandmarks,omitempty"`
	FaceAttributes *Attributes `json:"faceAttributes,omitempty"`
}

// Emotion properties describing facial emotion in form of confidence ranging from 0 to 1.
type Emotion struct {
	Anger     *float64 `json:"anger,omitempty"`
	Contempt  *float64 `json:"contempt,omitempty"`
	Disgust   *float64 `json:"disgust,omitempty"`
	Fear      *float64 `json:"fear,omitempty"`
	Happiness *float64 `json:"happiness,omitempty"`
	Neutral   *float64 `json:"neutral,omitempty"`
	Sadness   *float64 `json:"sadness,omitempty"`
	Surprise  *float64 `json:"surprise,omitempty"`
}

// Error error body.
type Error struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// Exposure properties describing exposure level of the image.
type Exposure struct {
	// ExposureLevel - An enum value indicating level of exposure. Possible values include: 'UnderExposure', 'GoodExposure', 'OverExposure'
	ExposureLevel ExposureLevel `json:"exposureLevel,omitempty"`
	// Value - A number indicating level of exposure level ranging from 0 to 1. [0, 0.25) is under exposure. [0.25, 0.75) is good exposure. [0.75, 1] is over exposure.
	Value *float64 `json:"value,omitempty"`
}

// FacialHair properties describing facial hair attributes.
type FacialHair struct {
	Moustache *float64 `json:"moustache,omitempty"`
	Beard     *float64 `json:"beard,omitempty"`
	Sideburns *float64 `json:"sideburns,omitempty"`
}

// FindSimilarRequest request body for find similar operation.
type FindSimilarRequest struct {
	// FaceID - FaceId of the query face. User needs to call Face - Detect first to get a valid faceId. Note that this faceId is not persisted and will expire 24 hours after the detection call
	FaceID *uuid.UUID `json:"faceId,omitempty"`
	// FaceListID - An existing user-specified unique candidate face list, created in Face List - Create a Face List. Face list contains a set of persistedFaceIds which are persisted and will never expire. Parameter faceListId, largeFaceListId and faceIds should not be provided at the same time。
	FaceListID *string `json:"faceListId,omitempty"`
	// LargeFaceListID - An existing user-specified unique candidate large face list, created in LargeFaceList - Create. Large face list contains a set of persistedFaceIds which are persisted and will never expire. Parameter faceListId, largeFaceListId and faceIds should not be provided at the same time.
	LargeFaceListID *string `json:"largeFaceListId,omitempty"`
	// FaceIds - An array of candidate faceIds. All of them are created by Face - Detect and the faceIds will expire 24 hours after the detection call. The number of faceIds is limited to 1000. Parameter faceListId, largeFaceListId and faceIds should not be provided at the same time.
	FaceIds *[]uuid.UUID `json:"faceIds,omitempty"`
	// MaxNumOfCandidatesReturned - The number of top similar faces returned. The valid range is [1, 1000].
	MaxNumOfCandidatesReturned *int32 `json:"maxNumOfCandidatesReturned,omitempty"`
	// Mode - Similar face searching mode. It can be "matchPerson" or "matchFace". Possible values include: 'MatchPerson', 'MatchFace'
	Mode FindSimilarMatchMode `json:"mode,omitempty"`
}

// GroupRequest request body for group request.
type GroupRequest struct {
	// FaceIds - Array of candidate faceId created by Face - Detect. The maximum is 1000 faces
	FaceIds *[]uuid.UUID `json:"faceIds,omitempty"`
}

// GroupResult an array of face groups based on face similarity.
type GroupResult struct {
	autorest.Response `json:"-"`
	// Groups - A partition of the original faces based on face similarity. Groups are ranked by number of faces
	Groups *[][]uuid.UUID `json:"groups,omitempty"`
	// MessyGroup - Face ids array of faces that cannot find any similar faces from original faces.
	MessyGroup *[]uuid.UUID `json:"messyGroup,omitempty"`
}

// Hair properties describing hair attributes.
type Hair struct {
	// Bald - A number describing confidence level of whether the person is bald.
	Bald *float64 `json:"bald,omitempty"`
	// Invisible - A boolean value describing whether the hair is visible in the image.
	Invisible *bool `json:"invisible,omitempty"`
	// HairColor - An array of candidate colors and confidence level in the presence of each.
	HairColor *[]HairColor `json:"hairColor,omitempty"`
}

// HairColor hair color and associated confidence
type HairColor struct {
	// Color - Name of the hair color. Possible values include: 'Unknown', 'White', 'Gray', 'Blond', 'Brown', 'Red', 'Black', 'Other'
	Color HairColorType `json:"color,omitempty"`
	// Confidence - Confidence level of the color
	Confidence *float64 `json:"confidence,omitempty"`
}

// HeadPose properties indicating head pose of the face.
type HeadPose struct {
	Roll  *float64 `json:"roll,omitempty"`
	Yaw   *float64 `json:"yaw,omitempty"`
	Pitch *float64 `json:"pitch,omitempty"`
}

// IdentifyCandidate all possible faces that may qualify.
type IdentifyCandidate struct {
	// PersonID - Id of candidate
	PersonID *uuid.UUID `json:"personId,omitempty"`
	// Confidence - Confidence threshold of identification, used to judge whether one face belong to one person. The range of confidenceThreshold is [0, 1] (default specified by algorithm).
	Confidence *float64 `json:"confidence,omitempty"`
}

// IdentifyRequest request body for identify face operation.
type IdentifyRequest struct {
	// FaceIds - Array of query faces faceIds, created by the Face - Detect. Each of the faces are identified independently. The valid number of faceIds is between [1, 10].
	FaceIds *[]uuid.UUID `json:"faceIds,omitempty"`
	// PersonGroupID - PersonGroupId of the target person group, created by PersonGroup - Create. Parameter personGroupId and largePersonGroupId should not be provided at the same time.
	PersonGroupID *string `json:"personGroupId,omitempty"`
	// LargePersonGroupID - LargePersonGroupId of the target large person group, created by LargePersonGroup - Create. Parameter personGroupId and largePersonGroupId should not be provided at the same time.
	LargePersonGroupID *string `json:"largePersonGroupId,omitempty"`
	// MaxNumOfCandidatesReturned - The range of maxNumOfCandidatesReturned is between 1 and 5 (default is 1).
	MaxNumOfCandidatesReturned *int32 `json:"maxNumOfCandidatesReturned,omitempty"`
	// ConfidenceThreshold - Confidence threshold of identification, used to judge whether one face belong to one person. The range of confidenceThreshold is [0, 1] (default specified by algorithm).
	ConfidenceThreshold *float64 `json:"confidenceThreshold,omitempty"`
}

// IdentifyResult response body for identify face operation.
type IdentifyResult struct {
	// FaceID - FaceId of the query face
	FaceID *uuid.UUID `json:"faceId,omitempty"`
	// Candidates - Identified person candidates for that face (ranked by confidence). Array size should be no larger than input maxNumOfCandidatesReturned. If no person is identified, will return an empty array.
	Candidates *[]IdentifyCandidate `json:"candidates,omitempty"`
}

// ImageURL ...
type ImageURL struct {
	// URL - Publicly reachable URL of an image
	URL *string `json:"url,omitempty"`
}

// Landmarks a collection of 27-point face landmarks pointing to the important positions of face
// components.
type Landmarks struct {
	PupilLeft           *Coordinate `json:"pupilLeft,omitempty"`
	PupilRight          *Coordinate `json:"pupilRight,omitempty"`
	NoseTip             *Coordinate `json:"noseTip,omitempty"`
	MouthLeft           *Coordinate `json:"mouthLeft,omitempty"`
	MouthRight          *Coordinate `json:"mouthRight,omitempty"`
	EyebrowLeftOuter    *Coordinate `json:"eyebrowLeftOuter,omitempty"`
	EyebrowLeftInner    *Coordinate `json:"eyebrowLeftInner,omitempty"`
	EyeLeftOuter        *Coordinate `json:"eyeLeftOuter,omitempty"`
	EyeLeftTop          *Coordinate `json:"eyeLeftTop,omitempty"`
	EyeLeftBottom       *Coordinate `json:"eyeLeftBottom,omitempty"`
	EyeLeftInner        *Coordinate `json:"eyeLeftInner,omitempty"`
	EyebrowRightInner   *Coordinate `json:"eyebrowRightInner,omitempty"`
	EyebrowRightOuter   *Coordinate `json:"eyebrowRightOuter,omitempty"`
	EyeRightInner       *Coordinate `json:"eyeRightInner,omitempty"`
	EyeRightTop         *Coordinate `json:"eyeRightTop,omitempty"`
	EyeRightBottom      *Coordinate `json:"eyeRightBottom,omitempty"`
	EyeRightOuter       *Coordinate `json:"eyeRightOuter,omitempty"`
	NoseRootLeft        *Coordinate `json:"noseRootLeft,omitempty"`
	NoseRootRight       *Coordinate `json:"noseRootRight,omitempty"`
	NoseLeftAlarTop     *Coordinate `json:"noseLeftAlarTop,omitempty"`
	NoseRightAlarTop    *Coordinate `json:"noseRightAlarTop,omitempty"`
	NoseLeftAlarOutTip  *Coordinate `json:"noseLeftAlarOutTip,omitempty"`
	NoseRightAlarOutTip *Coordinate `json:"noseRightAlarOutTip,omitempty"`
	UpperLipTop         *Coordinate `json:"upperLipTop,omitempty"`
	UpperLipBottom      *Coordinate `json:"upperLipBottom,omitempty"`
	UnderLipTop         *Coordinate `json:"underLipTop,omitempty"`
	UnderLipBottom      *Coordinate `json:"underLipBottom,omitempty"`
}

// LargeFaceList large face list object.
type LargeFaceList struct {
	autorest.Response `json:"-"`
	// LargeFaceListID - LargeFaceListId of the target large face list.
	LargeFaceListID *string `json:"largeFaceListId,omitempty"`
	// Name - User defined name, maximum length is 128.
	Name *string `json:"name,omitempty"`
	// UserData - User specified data. Length should not exceed 16KB.
	UserData *string `json:"userData,omitempty"`
}

// LargePersonGroup large person group object.
type LargePersonGroup struct {
	autorest.Response `json:"-"`
	// LargePersonGroupID - LargePersonGroupId of the target large person groups
	LargePersonGroupID *string `json:"largePersonGroupId,omitempty"`
	// Name - User defined name, maximum length is 128.
	Name *string `json:"name,omitempty"`
	// UserData - User specified data. Length should not exceed 16KB.
	UserData *string `json:"userData,omitempty"`
}

// List face list object.
type List struct {
	autorest.Response `json:"-"`
	// FaceListID - FaceListId of the target face list.
	FaceListID *string `json:"faceListId,omitempty"`
	// PersistedFaces - Persisted faces within the face list.
	PersistedFaces *[]PersistedFace `json:"persistedFaces,omitempty"`
	// Name - User defined name, maximum length is 128.
	Name *string `json:"name,omitempty"`
	// UserData - User specified data. Length should not exceed 16KB.
	UserData *string `json:"userData,omitempty"`
}

// ListDetectedFace ...
type ListDetectedFace struct {
	autorest.Response `json:"-"`
	Value             *[]DetectedFace `json:"value,omitempty"`
}

// ListIdentifyResult ...
type ListIdentifyResult struct {
	autorest.Response `json:"-"`
	Value             *[]IdentifyResult `json:"value,omitempty"`
}

// ListLargeFaceList ...
type ListLargeFaceList struct {
	autorest.Response `json:"-"`
	Value             *[]LargeFaceList `json:"value,omitempty"`
}

// ListLargePersonGroup ...
type ListLargePersonGroup struct {
	autorest.Response `json:"-"`
	Value             *[]LargePersonGroup `json:"value,omitempty"`
}

// ListList ...
type ListList struct {
	autorest.Response `json:"-"`
	Value             *[]List `json:"value,omitempty"`
}

// ListPersistedFace ...
type ListPersistedFace struct {
	autorest.Response `json:"-"`
	Value             *[]PersistedFace `json:"value,omitempty"`
}

// ListPerson ...
type ListPerson struct {
	autorest.Response `json:"-"`
	Value             *[]Person `json:"value,omitempty"`
}

// ListPersonGroup ...
type ListPersonGroup struct {
	autorest.Response `json:"-"`
	Value             *[]PersonGroup `json:"value,omitempty"`
}

// ListSimilarFace ...
type ListSimilarFace struct {
	autorest.Response `json:"-"`
	Value             *[]SimilarFace `json:"value,omitempty"`
}

// Makeup properties describing present makeups on a given face.
type Makeup struct {
	// EyeMakeup - A boolean value describing whether eye makeup is present on a face.
	EyeMakeup *bool `json:"eyeMakeup,omitempty"`
	// LipMakeup - A boolean value describing whether lip makeup is present on a face.
	LipMakeup *bool `json:"lipMakeup,omitempty"`
}

// NameAndUserDataContract a combination of user defined name and user specified data for the person,
// largePersonGroup/personGroup, and largeFaceList/faceList.
type NameAndUserDataContract struct {
	// Name - User defined name, maximum length is 128.
	Name *string `json:"name,omitempty"`
	// UserData - User specified data. Length should not exceed 16KB.
	UserData *string `json:"userData,omitempty"`
}

// Noise properties describing noise level of the image.
type Noise struct {
	// NoiseLevel - An enum value indicating level of noise. Possible values include: 'NoiseLevelLow', 'NoiseLevelMedium', 'NoiseLevelHigh'
	NoiseLevel NoiseLevel `json:"noiseLevel,omitempty"`
	// Value - A number indicating level of noise level ranging from 0 to 1. [0, 0.25) is under exposure. [0.25, 0.75) is good exposure. [0.75, 1] is over exposure. [0, 0.3) is low noise level. [0.3, 0.7) is medium noise level. [0.7, 1] is high noise level.
	Value *float64 `json:"value,omitempty"`
}

// Occlusion properties describing occlusions on a given face.
type Occlusion struct {
	// ForeheadOccluded - A boolean value indicating whether forehead is occluded.
	ForeheadOccluded *bool `json:"foreheadOccluded,omitempty"`
	// EyeOccluded - A boolean value indicating whether eyes are occluded.
	EyeOccluded *bool `json:"eyeOccluded,omitempty"`
	// MouthOccluded - A boolean value indicating whether the mouth is occluded.
	MouthOccluded *bool `json:"mouthOccluded,omitempty"`
}

// PersistedFace personFace object.
type PersistedFace struct {
	autorest.Response `json:"-"`
	// PersistedFaceID - The persistedFaceId of the target face, which is persisted and will not expire. Different from faceId created by Face - Detect and will expire in 24 hours after the detection call.
	PersistedFaceID *uuid.UUID `json:"persistedFaceId,omitempty"`
	// UserData - User-provided data attached to the face. The size limit is 1KB.
	UserData *string `json:"userData,omitempty"`
}

// Person person object.
type Person struct {
	autorest.Response `json:"-"`
	// PersonID - PersonId of the target face list.
	PersonID *uuid.UUID `json:"personId,omitempty"`
	// PersistedFaceIds - PersistedFaceIds of registered faces in the person. These persistedFaceIds are returned from Person - Add a Person Face, and will not expire.
	PersistedFaceIds *[]uuid.UUID `json:"persistedFaceIds,omitempty"`
	// Name - User defined name, maximum length is 128.
	Name *string `json:"name,omitempty"`
	// UserData - User specified data. Length should not exceed 16KB.
	UserData *string `json:"userData,omitempty"`
}

// PersonGroup person group object.
type PersonGroup struct {
	autorest.Response `json:"-"`
	// PersonGroupID - PersonGroupId of the target person group.
	PersonGroupID *string `json:"personGroupId,omitempty"`
	// Name - User defined name, maximum length is 128.
	Name *string `json:"name,omitempty"`
	// UserData - User specified data. Length should not exceed 16KB.
	UserData *string `json:"userData,omitempty"`
}

// Rectangle a rectangle within which a face can be found
type Rectangle struct {
	// Width - The width of the rectangle, in pixels.
	Width *int32 `json:"width,omitempty"`
	// Height - The height of the rectangle, in pixels.
	Height *int32 `json:"height,omitempty"`
	// Left - The distance from the left edge if the image to the left edge of the rectangle, in pixels.
	Left *int32 `json:"left,omitempty"`
	// Top - The distance from the top edge if the image to the top edge of the rectangle, in pixels.
	Top *int32 `json:"top,omitempty"`
}

// SimilarFace response body for find similar face operation.
type SimilarFace struct {
	// FaceID - FaceId of candidate face when find by faceIds. faceId is created by Face - Detect and will expire 24 hours after the detection call
	FaceID *uuid.UUID `json:"faceId,omitempty"`
	// PersistedFaceID - PersistedFaceId of candidate face when find by faceListId. persistedFaceId in face list is persisted and will not expire. As showed in below response
	PersistedFaceID *uuid.UUID `json:"persistedFaceId,omitempty"`
	// Confidence - Similarity confidence of the candidate face. The higher confidence, the more similar. Range between [0,1].
	Confidence *float64 `json:"confidence,omitempty"`
}

// TrainingStatus training status object.
type TrainingStatus struct {
	autorest.Response `json:"-"`
	// Status - Training status: notstarted, running, succeeded, failed. If the training process is waiting to perform, the status is notstarted. If the training is ongoing, the status is running. Status succeed means this person group or large person group is ready for Face - Identify, or this large face list is ready for Face - Find Similar. Status failed is often caused by no person or no persisted face exist in the person group or large person group, or no persisted face exist in the large face list. Possible values include: 'Nonstarted', 'Running', 'Succeeded', 'Failed'
	Status TrainingStatusType `json:"status,omitempty"`
	// Created - A combined UTC date and time string that describes the created time of the person group, large person group or large face list.
	Created *date.Time `json:"createdDateTime,omitempty"`
	// LastAction - A combined UTC date and time string that describes the last modify time of the person group, large person group or large face list, could be null value when the group is not successfully trained.
	LastAction *date.Time `json:"lastActionDateTime,omitempty"`
	// LastSuccessfulTraining - A combined UTC date and time string that describes the last successful training time of the person group, large person group or large face list.
	LastSuccessfulTraining *date.Time `json:"lastSuccessfulTrainingDateTime,omitempty"`
	// Message - Show failure message when training failed (omitted when training succeed).
	Message *string `json:"message,omitempty"`
}

// UpdateFaceRequest request to update face data.
type UpdateFaceRequest struct {
	// UserData - User-provided data attached to the face. The size limit is 1KB.
	UserData *string `json:"userData,omitempty"`
}

// VerifyFaceToFaceRequest request body for face to face verification.
type VerifyFaceToFaceRequest struct {
	// FaceID1 - FaceId of the first face, comes from Face - Detect
	FaceID1 *uuid.UUID `json:"faceId1,omitempty"`
	// FaceID2 - FaceId of the second face, comes from Face - Detect
	FaceID2 *uuid.UUID `json:"faceId2,omitempty"`
}

// VerifyFaceToPersonRequest request body for face to person verification.
type VerifyFaceToPersonRequest struct {
	// FaceID - FaceId of the face, comes from Face - Detect
	FaceID *uuid.UUID `json:"faceId,omitempty"`
	// PersonGroupID - Using existing personGroupId and personId for fast loading a specified person. personGroupId is created in PersonGroup - Create. Parameter personGroupId and largePersonGroupId should not be provided at the same time.
	PersonGroupID *string `json:"personGroupId,omitempty"`
	// LargePersonGroupID - Using existing largePersonGroupId and personId for fast loading a specified person. largePersonGroupId is created in LargePersonGroup - Create. Parameter personGroupId and largePersonGroupId should not be provided at the same time.
	LargePersonGroupID *string `json:"largePersonGroupId,omitempty"`
	// PersonID - Specify a certain person in a person group or a large person group. personId is created in PersonGroup Person - Create or LargePersonGroup Person - Create.
	PersonID *uuid.UUID `json:"personId,omitempty"`
}

// VerifyResult result of the verify operation.
type VerifyResult struct {
	autorest.Response `json:"-"`
	// IsIdentical - True if the two faces belong to the same person or the face belongs to the person, otherwise false.
	IsIdentical *bool `json:"isIdentical,omitempty"`
	// Confidence - A number indicates the similarity confidence of whether two faces belong to the same person, or whether the face belongs to the person. By default, isIdentical is set to True if similarity confidence is greater than or equal to 0.5. This is useful for advanced users to override "isIdentical" and fine-tune the result on their own data.
	Confidence *float64 `json:"confidence,omitempty"`
}
