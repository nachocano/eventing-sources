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

package resources

import (
	"fmt"
	"knative.dev/eventing-contrib/container/pkg/apis/sources/v1alpha1"
	"knative.dev/eventing/pkg/utils"
)

func GenerateName(src *v1alpha1.ContainerSource) string {
	return utils.GenerateFixedName(src, fmt.Sprintf("containersource-%s", src.Name))
}
