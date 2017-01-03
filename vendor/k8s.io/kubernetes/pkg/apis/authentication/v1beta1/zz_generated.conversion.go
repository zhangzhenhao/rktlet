// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1beta1

import (
	authentication "k8s.io/kubernetes/pkg/apis/authentication"
	conversion "k8s.io/kubernetes/pkg/conversion"
	runtime "k8s.io/kubernetes/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	SchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1beta1_TokenReview_To_authentication_TokenReview,
		Convert_authentication_TokenReview_To_v1beta1_TokenReview,
		Convert_v1beta1_TokenReviewSpec_To_authentication_TokenReviewSpec,
		Convert_authentication_TokenReviewSpec_To_v1beta1_TokenReviewSpec,
		Convert_v1beta1_TokenReviewStatus_To_authentication_TokenReviewStatus,
		Convert_authentication_TokenReviewStatus_To_v1beta1_TokenReviewStatus,
		Convert_v1beta1_UserInfo_To_authentication_UserInfo,
		Convert_authentication_UserInfo_To_v1beta1_UserInfo,
	)
}

func autoConvert_v1beta1_TokenReview_To_authentication_TokenReview(in *TokenReview, out *authentication.TokenReview, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_v1beta1_TokenReviewSpec_To_authentication_TokenReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_TokenReviewStatus_To_authentication_TokenReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1beta1_TokenReview_To_authentication_TokenReview(in *TokenReview, out *authentication.TokenReview, s conversion.Scope) error {
	return autoConvert_v1beta1_TokenReview_To_authentication_TokenReview(in, out, s)
}

func autoConvert_authentication_TokenReview_To_v1beta1_TokenReview(in *authentication.TokenReview, out *TokenReview, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_authentication_TokenReviewSpec_To_v1beta1_TokenReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_authentication_TokenReviewStatus_To_v1beta1_TokenReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_authentication_TokenReview_To_v1beta1_TokenReview(in *authentication.TokenReview, out *TokenReview, s conversion.Scope) error {
	return autoConvert_authentication_TokenReview_To_v1beta1_TokenReview(in, out, s)
}

func autoConvert_v1beta1_TokenReviewSpec_To_authentication_TokenReviewSpec(in *TokenReviewSpec, out *authentication.TokenReviewSpec, s conversion.Scope) error {
	out.Token = in.Token
	return nil
}

func Convert_v1beta1_TokenReviewSpec_To_authentication_TokenReviewSpec(in *TokenReviewSpec, out *authentication.TokenReviewSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_TokenReviewSpec_To_authentication_TokenReviewSpec(in, out, s)
}

func autoConvert_authentication_TokenReviewSpec_To_v1beta1_TokenReviewSpec(in *authentication.TokenReviewSpec, out *TokenReviewSpec, s conversion.Scope) error {
	out.Token = in.Token
	return nil
}

func Convert_authentication_TokenReviewSpec_To_v1beta1_TokenReviewSpec(in *authentication.TokenReviewSpec, out *TokenReviewSpec, s conversion.Scope) error {
	return autoConvert_authentication_TokenReviewSpec_To_v1beta1_TokenReviewSpec(in, out, s)
}

func autoConvert_v1beta1_TokenReviewStatus_To_authentication_TokenReviewStatus(in *TokenReviewStatus, out *authentication.TokenReviewStatus, s conversion.Scope) error {
	out.Authenticated = in.Authenticated
	if err := Convert_v1beta1_UserInfo_To_authentication_UserInfo(&in.User, &out.User, s); err != nil {
		return err
	}
	out.Error = in.Error
	return nil
}

func Convert_v1beta1_TokenReviewStatus_To_authentication_TokenReviewStatus(in *TokenReviewStatus, out *authentication.TokenReviewStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_TokenReviewStatus_To_authentication_TokenReviewStatus(in, out, s)
}

func autoConvert_authentication_TokenReviewStatus_To_v1beta1_TokenReviewStatus(in *authentication.TokenReviewStatus, out *TokenReviewStatus, s conversion.Scope) error {
	out.Authenticated = in.Authenticated
	if err := Convert_authentication_UserInfo_To_v1beta1_UserInfo(&in.User, &out.User, s); err != nil {
		return err
	}
	out.Error = in.Error
	return nil
}

func Convert_authentication_TokenReviewStatus_To_v1beta1_TokenReviewStatus(in *authentication.TokenReviewStatus, out *TokenReviewStatus, s conversion.Scope) error {
	return autoConvert_authentication_TokenReviewStatus_To_v1beta1_TokenReviewStatus(in, out, s)
}

func autoConvert_v1beta1_UserInfo_To_authentication_UserInfo(in *UserInfo, out *authentication.UserInfo, s conversion.Scope) error {
	out.Username = in.Username
	out.UID = in.UID
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.Extra = *(*map[string]authentication.ExtraValue)(unsafe.Pointer(&in.Extra))
	return nil
}

func Convert_v1beta1_UserInfo_To_authentication_UserInfo(in *UserInfo, out *authentication.UserInfo, s conversion.Scope) error {
	return autoConvert_v1beta1_UserInfo_To_authentication_UserInfo(in, out, s)
}

func autoConvert_authentication_UserInfo_To_v1beta1_UserInfo(in *authentication.UserInfo, out *UserInfo, s conversion.Scope) error {
	out.Username = in.Username
	out.UID = in.UID
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.Extra = *(*map[string]ExtraValue)(unsafe.Pointer(&in.Extra))
	return nil
}

func Convert_authentication_UserInfo_To_v1beta1_UserInfo(in *authentication.UserInfo, out *UserInfo, s conversion.Scope) error {
	return autoConvert_authentication_UserInfo_To_v1beta1_UserInfo(in, out, s)
}
