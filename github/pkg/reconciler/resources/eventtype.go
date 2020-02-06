/*
Copyright 2019 The Knative Authors

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

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	eventingv1alpha1 "knative.dev/eventing/pkg/apis/eventing/v1alpha1"
	"knative.dev/eventing/pkg/utils"

	"knative.dev/eventing-contrib/github/pkg/apis/sources/v1alpha1"
)

// EventTypeArgs are the arguments needed to create an EventType for an Couch DB source.
type EventTypeArgs struct {
	Src    *v1alpha1.GitHubSource
	Type   string
	Source string
}

// MakeEventType creates the in-memory representation of the EventType for the specified GitHubSource.
func MakeEventType(args *EventTypeArgs) eventingv1alpha1.EventType {
	return eventingv1alpha1.EventType{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: fmt.Sprintf("%s-", utils.ToDNS1123Subdomain(args.Type)),
			Labels:       Labels(args.Src.Name),
			Namespace:    args.Src.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(args.Src, schema.GroupVersionKind{
					Group:   v1alpha1.SchemeGroupVersion.Group,
					Version: v1alpha1.SchemeGroupVersion.Version,
					Kind:    "GitHubSource",
				}),
			},
		},
		Spec: eventingv1alpha1.EventTypeSpec{
			Type:   args.Type,
			Source: args.Source,
			Broker: args.Src.Spec.Sink.GetRef().Name,
		},
	}
}
