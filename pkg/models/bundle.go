package models

import (
	apimanifests "github.com/operator-framework/api/pkg/manifests"
)

type ExtractBundle struct {
	Bundle                  *apimanifests.Bundle
	OperatorBundleName      string
	OperatorBundleImagePath string
	Errors                  []string
}

func NewExtractBundle(operatorBundleName, operatorBundleImagePath string) *ExtractBundle {
	extractBundle := ExtractBundle{}
	extractBundle.OperatorBundleName = operatorBundleName
	extractBundle.OperatorBundleImagePath = operatorBundleImagePath
	return &extractBundle
}
