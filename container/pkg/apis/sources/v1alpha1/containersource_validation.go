/*
Copyright 2020 The Knative Authors

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

package v1alpha1

import (
	"context"

	"knative.dev/pkg/apis"
)

func (c *ContainerSource) Validate(ctx context.Context) *apis.FieldError {
	return c.Spec.Validate(ctx).ViaField("spec")
}

func (cs *ContainerSourceSpec) Validate(ctx context.Context) *apis.FieldError {
	var errs *apis.FieldError
	if fe := cs.Sink.Validate(ctx); fe != nil {
		errs = errs.Also(fe.ViaField("sink"))
	}
	return errs
}
