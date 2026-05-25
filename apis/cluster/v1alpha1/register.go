/*
Copyright 2024 The Crossplane Authors.
*/

// Package v1alpha1 is a placeholder for resources promoted to v1beta1 (or
// not yet generated). Upjet's codegen emits an import of this package
// even when no v1alpha1 resources exist; keeping an empty SchemeBuilder
// keeps the scheme-registration code in zz_register.go satisfied.
package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// Package type metadata.
const (
	Group   = "listmonk.crossplane.io"
	Version = "v1alpha1"
)

var (
	// SchemeGroupVersion is the group/version this stub package would
	// register if it held any types. Kept exported so the codegen's
	// zz_register.go import compiles even when no resources exist at
	// this scope/version.
	SchemeGroupVersion = schema.GroupVersion{Group: Group, Version: Version}
	// SchemeBuilder is the scheme builder; empty by design (see file
	// doc comment).
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}
)
