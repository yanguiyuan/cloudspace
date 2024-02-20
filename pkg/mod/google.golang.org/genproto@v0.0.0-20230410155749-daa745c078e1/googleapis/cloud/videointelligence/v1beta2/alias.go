// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by aliasgen. DO NOT EDIT.

// Package videointelligence aliases all exported identifiers in package
// "cloud.google.com/go/videointelligence/apiv1beta2/videointelligencepb".
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1beta2/videointelligencepb.
// Please read https://github.com/googleapis/google-cloud-go/blob/main/migration.md
// for more details.
package videointelligence

import (
	src "cloud.google.com/go/videointelligence/apiv1beta2/videointelligencepb"
	grpc "google.golang.org/grpc"
)

// Deprecated: Please use consts in: cloud.google.com/go/videointelligence/apiv1beta2/videointelligencepb
const (
	Feature_EXPLICIT_CONTENT_DETECTION                  = src.Feature_EXPLICIT_CONTENT_DETECTION
	Feature_FACE_DETECTION                              = src.Feature_FACE_DETECTION
	Feature_FEATURE_UNSPECIFIED                         = src.Feature_FEATURE_UNSPECIFIED
	Feature_LABEL_DETECTION                             = src.Feature_LABEL_DETECTION
	Feature_SHOT_CHANGE_DETECTION                       = src.Feature_SHOT_CHANGE_DETECTION
	LabelDetectionMode_FRAME_MODE                       = src.LabelDetectionMode_FRAME_MODE
	LabelDetectionMode_LABEL_DETECTION_MODE_UNSPECIFIED = src.LabelDetectionMode_LABEL_DETECTION_MODE_UNSPECIFIED
	LabelDetectionMode_SHOT_AND_FRAME_MODE              = src.LabelDetectionMode_SHOT_AND_FRAME_MODE
	LabelDetectionMode_SHOT_MODE                        = src.LabelDetectionMode_SHOT_MODE
	Likelihood_LIKELIHOOD_UNSPECIFIED                   = src.Likelihood_LIKELIHOOD_UNSPECIFIED
	Likelihood_LIKELY                                   = src.Likelihood_LIKELY
	Likelihood_POSSIBLE                                 = src.Likelihood_POSSIBLE
	Likelihood_UNLIKELY                                 = src.Likelihood_UNLIKELY
	Likelihood_VERY_LIKELY                              = src.Likelihood_VERY_LIKELY
	Likelihood_VERY_UNLIKELY                            = src.Likelihood_VERY_UNLIKELY
)

// Deprecated: Please use vars in: cloud.google.com/go/videointelligence/apiv1beta2/videointelligencepb
var (
	Feature_name                                                         = src.Feature_name
	Feature_value                                                        = src.Feature_value
	File_google_cloud_videointelligence_v1beta2_video_intelligence_proto = src.File_google_cloud_videointelligence_v1beta2_video_intelligence_proto
	LabelDetectionMode_name                                              = src.LabelDetectionMode_name
	LabelDetectionMode_value                                             = src.LabelDetectionMode_value
	Likelihood_name                                                      = src.Likelihood_name
	Likelihood_value                                                     = src.Likelihood_value
)

// Video annotation progress. Included in the `metadata` field of the
// `Operation` returned by the `GetOperation` call of the
// `google::longrunning::Operations` service.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1beta2/videointelligencepb
type AnnotateVideoProgress = src.AnnotateVideoProgress

// Video annotation request.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1beta2/videointelligencepb
type AnnotateVideoRequest = src.AnnotateVideoRequest

// Video annotation response. Included in the `response` field of the
// `Operation` returned by the `GetOperation` call of the
// `google::longrunning::Operations` service.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1beta2/videointelligencepb
type AnnotateVideoResponse = src.AnnotateVideoResponse

// Detected entity from video analysis.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1beta2/videointelligencepb
type Entity = src.Entity

// Explicit content annotation (based on per-frame visual signals only). If no
// explicit content has been detected in a frame, no annotations are present
// for that frame.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1beta2/videointelligencepb
type ExplicitContentAnnotation = src.ExplicitContentAnnotation

// Config for EXPLICIT_CONTENT_DETECTION.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1beta2/videointelligencepb
type ExplicitContentDetectionConfig = src.ExplicitContentDetectionConfig

// Video frame level annotation results for explicit content.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1beta2/videointelligencepb
type ExplicitContentFrame = src.ExplicitContentFrame

// Face annotation.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1beta2/videointelligencepb
type FaceAnnotation = src.FaceAnnotation

// Config for FACE_DETECTION.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1beta2/videointelligencepb
type FaceDetectionConfig = src.FaceDetectionConfig

// Video frame level annotation results for face detection.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1beta2/videointelligencepb
type FaceFrame = src.FaceFrame

// Video segment level annotation results for face detection.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1beta2/videointelligencepb
type FaceSegment = src.FaceSegment

// Video annotation feature.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1beta2/videointelligencepb
type Feature = src.Feature

// Label annotation.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1beta2/videointelligencepb
type LabelAnnotation = src.LabelAnnotation

// Config for LABEL_DETECTION.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1beta2/videointelligencepb
type LabelDetectionConfig = src.LabelDetectionConfig

// Label detection mode.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1beta2/videointelligencepb
type LabelDetectionMode = src.LabelDetectionMode

// Video frame level annotation results for label detection.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1beta2/videointelligencepb
type LabelFrame = src.LabelFrame

// Video segment level annotation results for label detection.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1beta2/videointelligencepb
type LabelSegment = src.LabelSegment

// Bucketized representation of likelihood.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1beta2/videointelligencepb
type Likelihood = src.Likelihood

// Normalized bounding box. The normalized vertex coordinates are relative to
// the original image. Range: [0, 1].
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1beta2/videointelligencepb
type NormalizedBoundingBox = src.NormalizedBoundingBox

// Config for SHOT_CHANGE_DETECTION.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1beta2/videointelligencepb
type ShotChangeDetectionConfig = src.ShotChangeDetectionConfig

// UnimplementedVideoIntelligenceServiceServer can be embedded to have forward
// compatible implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1beta2/videointelligencepb
type UnimplementedVideoIntelligenceServiceServer = src.UnimplementedVideoIntelligenceServiceServer

// Annotation progress for a single video.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1beta2/videointelligencepb
type VideoAnnotationProgress = src.VideoAnnotationProgress

// Annotation results for a single video.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1beta2/videointelligencepb
type VideoAnnotationResults = src.VideoAnnotationResults

// Video context and/or feature-specific parameters.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1beta2/videointelligencepb
type VideoContext = src.VideoContext

// VideoIntelligenceServiceClient is the client API for
// VideoIntelligenceService service. For semantics around ctx use and
// closing/ending streaming RPCs, please refer to
// https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1beta2/videointelligencepb
type VideoIntelligenceServiceClient = src.VideoIntelligenceServiceClient

// VideoIntelligenceServiceServer is the server API for
// VideoIntelligenceService service.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1beta2/videointelligencepb
type VideoIntelligenceServiceServer = src.VideoIntelligenceServiceServer

// Video segment.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1beta2/videointelligencepb
type VideoSegment = src.VideoSegment

// Deprecated: Please use funcs in: cloud.google.com/go/videointelligence/apiv1beta2/videointelligencepb
func NewVideoIntelligenceServiceClient(cc grpc.ClientConnInterface) VideoIntelligenceServiceClient {
	return src.NewVideoIntelligenceServiceClient(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/videointelligence/apiv1beta2/videointelligencepb
func RegisterVideoIntelligenceServiceServer(s *grpc.Server, srv VideoIntelligenceServiceServer) {
	src.RegisterVideoIntelligenceServiceServer(s, srv)
}
