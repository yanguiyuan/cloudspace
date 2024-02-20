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

// Package addressvalidation aliases all exported identifiers in package
// "cloud.google.com/go/maps/addressvalidation/apiv1/addressvalidationpb".
//
// Deprecated: Please use types in: cloud.google.com/go/maps/addressvalidation/apiv1/addressvalidationpb.
// Please read https://github.com/googleapis/google-cloud-go/blob/main/migration.md
// for more details.
package addressvalidation

import (
	src "cloud.google.com/go/maps/addressvalidation/apiv1/addressvalidationpb"
	grpc "google.golang.org/grpc"
)

// Deprecated: Please use consts in: cloud.google.com/go/maps/addressvalidation/apiv1/addressvalidationpb
const (
	AddressComponent_CONFIRMATION_LEVEL_UNSPECIFIED                    = src.AddressComponent_CONFIRMATION_LEVEL_UNSPECIFIED
	AddressComponent_CONFIRMED                                         = src.AddressComponent_CONFIRMED
	AddressComponent_UNCONFIRMED_AND_SUSPICIOUS                        = src.AddressComponent_UNCONFIRMED_AND_SUSPICIOUS
	AddressComponent_UNCONFIRMED_BUT_PLAUSIBLE                         = src.AddressComponent_UNCONFIRMED_BUT_PLAUSIBLE
	ProvideValidationFeedbackRequest_UNUSED                            = src.ProvideValidationFeedbackRequest_UNUSED
	ProvideValidationFeedbackRequest_UNVALIDATED_VERSION_USED          = src.ProvideValidationFeedbackRequest_UNVALIDATED_VERSION_USED
	ProvideValidationFeedbackRequest_USER_VERSION_USED                 = src.ProvideValidationFeedbackRequest_USER_VERSION_USED
	ProvideValidationFeedbackRequest_VALIDATED_VERSION_USED            = src.ProvideValidationFeedbackRequest_VALIDATED_VERSION_USED
	ProvideValidationFeedbackRequest_VALIDATION_CONCLUSION_UNSPECIFIED = src.ProvideValidationFeedbackRequest_VALIDATION_CONCLUSION_UNSPECIFIED
	Verdict_BLOCK                                                      = src.Verdict_BLOCK
	Verdict_GRANULARITY_UNSPECIFIED                                    = src.Verdict_GRANULARITY_UNSPECIFIED
	Verdict_OTHER                                                      = src.Verdict_OTHER
	Verdict_PREMISE                                                    = src.Verdict_PREMISE
	Verdict_PREMISE_PROXIMITY                                          = src.Verdict_PREMISE_PROXIMITY
	Verdict_ROUTE                                                      = src.Verdict_ROUTE
	Verdict_SUB_PREMISE                                                = src.Verdict_SUB_PREMISE
)

// Deprecated: Please use vars in: cloud.google.com/go/maps/addressvalidation/apiv1/addressvalidationpb
var (
	AddressComponent_ConfirmationLevel_name                                = src.AddressComponent_ConfirmationLevel_name
	AddressComponent_ConfirmationLevel_value                               = src.AddressComponent_ConfirmationLevel_value
	File_google_maps_addressvalidation_v1_address_proto                    = src.File_google_maps_addressvalidation_v1_address_proto
	File_google_maps_addressvalidation_v1_address_validation_service_proto = src.File_google_maps_addressvalidation_v1_address_validation_service_proto
	File_google_maps_addressvalidation_v1_geocode_proto                    = src.File_google_maps_addressvalidation_v1_geocode_proto
	File_google_maps_addressvalidation_v1_metadata_proto                   = src.File_google_maps_addressvalidation_v1_metadata_proto
	File_google_maps_addressvalidation_v1_usps_data_proto                  = src.File_google_maps_addressvalidation_v1_usps_data_proto
	ProvideValidationFeedbackRequest_ValidationConclusion_name             = src.ProvideValidationFeedbackRequest_ValidationConclusion_name
	ProvideValidationFeedbackRequest_ValidationConclusion_value            = src.ProvideValidationFeedbackRequest_ValidationConclusion_value
	Verdict_Granularity_name                                               = src.Verdict_Granularity_name
	Verdict_Granularity_value                                              = src.Verdict_Granularity_value
)

// Details of the address parsed from the input.
//
// Deprecated: Please use types in: cloud.google.com/go/maps/addressvalidation/apiv1/addressvalidationpb
type Address = src.Address

// Represents an address component, such as a street, city, or state.
//
// Deprecated: Please use types in: cloud.google.com/go/maps/addressvalidation/apiv1/addressvalidationpb
type AddressComponent = src.AddressComponent

// The different possible values for confirmation levels.
//
// Deprecated: Please use types in: cloud.google.com/go/maps/addressvalidation/apiv1/addressvalidationpb
type AddressComponent_ConfirmationLevel = src.AddressComponent_ConfirmationLevel

// The metadata for the address.
//
// Deprecated: Please use types in: cloud.google.com/go/maps/addressvalidation/apiv1/addressvalidationpb
type AddressMetadata = src.AddressMetadata

// AddressValidationClient is the client API for AddressValidation service.
// For semantics around ctx use and closing/ending streaming RPCs, please refer
// to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/maps/addressvalidation/apiv1/addressvalidationpb
type AddressValidationClient = src.AddressValidationClient

// AddressValidationServer is the server API for AddressValidation service.
//
// Deprecated: Please use types in: cloud.google.com/go/maps/addressvalidation/apiv1/addressvalidationpb
type AddressValidationServer = src.AddressValidationServer

// A wrapper for the name of the component.
//
// Deprecated: Please use types in: cloud.google.com/go/maps/addressvalidation/apiv1/addressvalidationpb
type ComponentName = src.ComponentName

// Contains information about the place the input was geocoded to.
//
// Deprecated: Please use types in: cloud.google.com/go/maps/addressvalidation/apiv1/addressvalidationpb
type Geocode = src.Geocode

// Plus code (http://plus.codes) is a location reference with two formats:
// global code defining a 14mx14m (1/8000th of a degree) or smaller rectangle,
// and compound code, replacing the prefix with a reference location.
//
// Deprecated: Please use types in: cloud.google.com/go/maps/addressvalidation/apiv1/addressvalidationpb
type PlusCode = src.PlusCode

// The request for sending validation feedback.
//
// Deprecated: Please use types in: cloud.google.com/go/maps/addressvalidation/apiv1/addressvalidationpb
type ProvideValidationFeedbackRequest = src.ProvideValidationFeedbackRequest

// The possible final outcomes of the sequence of address validation requests
// needed to validate an address.
//
// Deprecated: Please use types in: cloud.google.com/go/maps/addressvalidation/apiv1/addressvalidationpb
type ProvideValidationFeedbackRequest_ValidationConclusion = src.ProvideValidationFeedbackRequest_ValidationConclusion

// The response for validation feedback. The response is empty if the feedback
// is sent successfully.
//
// Deprecated: Please use types in: cloud.google.com/go/maps/addressvalidation/apiv1/addressvalidationpb
type ProvideValidationFeedbackResponse = src.ProvideValidationFeedbackResponse

// UnimplementedAddressValidationServer can be embedded to have forward
// compatible implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/maps/addressvalidation/apiv1/addressvalidationpb
type UnimplementedAddressValidationServer = src.UnimplementedAddressValidationServer

// USPS representation of a US address.
//
// Deprecated: Please use types in: cloud.google.com/go/maps/addressvalidation/apiv1/addressvalidationpb
type UspsAddress = src.UspsAddress

// The USPS data for the address.
//
// Deprecated: Please use types in: cloud.google.com/go/maps/addressvalidation/apiv1/addressvalidationpb
type UspsData = src.UspsData

// The request for validating an address.
//
// Deprecated: Please use types in: cloud.google.com/go/maps/addressvalidation/apiv1/addressvalidationpb
type ValidateAddressRequest = src.ValidateAddressRequest

// The response to an address validation request.
//
// Deprecated: Please use types in: cloud.google.com/go/maps/addressvalidation/apiv1/addressvalidationpb
type ValidateAddressResponse = src.ValidateAddressResponse

// The result of validating an address.
//
// Deprecated: Please use types in: cloud.google.com/go/maps/addressvalidation/apiv1/addressvalidationpb
type ValidationResult = src.ValidationResult

// High level overview of the address validation result and geocode.
//
// Deprecated: Please use types in: cloud.google.com/go/maps/addressvalidation/apiv1/addressvalidationpb
type Verdict = src.Verdict

// The various granularities that an address or a geocode can have. When used
// to indicate granularity for an *address*, these values indicate with how
// fine a granularity the address identifies a mailing destination. For
// example, an address such as "123 Main Street, Redwood City, CA, 94061"
// identifies a `PREMISE` while something like "Redwood City, CA, 94061"
// identifies a `LOCALITY`. However, if we are unable to find a geocode for
// "123 Main Street" in Redwood City, the geocode returned might be of
// `LOCALITY` granularity even though the address is more granular.
//
// Deprecated: Please use types in: cloud.google.com/go/maps/addressvalidation/apiv1/addressvalidationpb
type Verdict_Granularity = src.Verdict_Granularity

// Deprecated: Please use funcs in: cloud.google.com/go/maps/addressvalidation/apiv1/addressvalidationpb
func NewAddressValidationClient(cc grpc.ClientConnInterface) AddressValidationClient {
	return src.NewAddressValidationClient(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/maps/addressvalidation/apiv1/addressvalidationpb
func RegisterAddressValidationServer(s *grpc.Server, srv AddressValidationServer) {
	src.RegisterAddressValidationServer(s, srv)
}
