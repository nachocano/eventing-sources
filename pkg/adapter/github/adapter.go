/*
Copyright 2018 The Knative Authors

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

package github

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/cloudevents/sdk-go/pkg/cloudevents"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/client"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/types"
	sourcesv1alpha1 "github.com/knative/eventing-sources/pkg/apis/sources/v1alpha1"
	"github.com/knative/eventing-sources/pkg/kncloudevents"
	gh "gopkg.in/go-playground/webhooks.v5/github"
)

const (
	GHHeaderEvent    = "GitHub-Event"
	GHHeaderDelivery = "GitHub-Delivery"
)

// Adapter converts incoming GitHub webhook events to CloudEvents
type Adapter struct {
	client client.Client
	source *types.URLRef
}

// New creates an adapter to convert incoming GitHub webhook events to CloudEvents and
// then sends them to the specified Sink
func New(sinkURI, ownerRepo string) (*Adapter, error) {
	a := new(Adapter)
	var err error
	a.client, err = kncloudevents.NewDefaultClient(sinkURI)
	if err != nil {
		return nil, err
	}
	// set the CloudEvent source.
	source := types.ParseURLRef(fmt.Sprintf("/%s", ownerRepo))
	if source == nil {
		return nil, fmt.Errorf("invalid source %q", ownerRepo)
	}
	a.source = source
	return a, nil
}

// HandleEvent is invoked whenever an event comes in from GitHub
func (a *Adapter) HandleEvent(payload interface{}, header http.Header) {
	hdr := http.Header(header)
	err := a.handleEvent(payload, hdr)
	if err != nil {
		log.Printf("unexpected error handling GitHub event: %s", err)
	}
}

func (a *Adapter) handleEvent(payload interface{}, hdr http.Header) error {
	gitHubEventType := hdr.Get("X-" + GHHeaderEvent)
	eventID := hdr.Get("X-" + GHHeaderDelivery)
	extensions := map[string]interface{}{
		GHHeaderEvent:    gitHubEventType,
		GHHeaderDelivery: eventID,
	}

	log.Printf("Handling %s", gitHubEventType)

	cloudEventType := fmt.Sprintf("%s.%s", sourcesv1alpha1.GitHubSourceEventPrefix, gitHubEventType)
	subject := subjectFromGitHubEvent(gh.Event(gitHubEventType), payload)

	eventContext := cloudevents.EventContextV02{
		ID:         eventID,
		Type:       cloudEventType,
		Source:     *a.source,
		Extensions: extensions,
	}.AsV02()
	eventContext.SetSubject(subject)

	event := cloudevents.Event{
		Context: eventContext,
		Data:    payload,
	}
	_, err := a.client.Send(context.TODO(), event)
	return err
}

func subjectFromGitHubEvent(gitHubEvent gh.Event, payload interface{}) string {
	var subject string
	switch gitHubEvent {
	case gh.CheckSuiteEvent:
		if cs, ok := payload.(gh.CheckSuitePayload); ok {
			subject = string(cs.CheckSuite.ID)
		}
	case gh.CommitCommentEvent:
		if cc, ok := payload.(gh.CommitCommentPayload); ok {
			// E.g., https://github.com/Codertocat/Hello-World/commit/a10867b14bb761a232cd80139fbd4c0d33264240#commitcomment-29186860
			// and we keep with a10867b14bb761a232cd80139fbd4c0d33264240#commitcomment-29186860
			subject = lastPathPortion(cc.Comment.HTMLURL)
		}
	case gh.CreateEvent:
		if c, ok := payload.(gh.CreatePayload); ok {
			// The object that was created, can be repository, branch, or tag.
			subject = c.RefType
		}
	case gh.DeleteEvent:
		if d, ok := payload.(gh.DeletePayload); ok {
			// The object that was deleted, can be branch or tag.
			subject = d.RefType
		}
	case gh.DeploymentEvent:
		if d, ok := payload.(gh.DeploymentPayload); ok {
			subject = string(d.Deployment.ID)
		}
	case gh.DeploymentStatusEvent:
		if d, ok := payload.(gh.DeploymentStatusPayload); ok {
			subject = string(d.Deployment.ID)
		}
	case gh.ForkEvent:
		if f, ok := payload.(gh.ForkPayload); ok {
			subject = string(f.Forkee.ID)
		}
	case gh.GollumEvent:
		if g, ok := payload.(gh.GollumPayload); ok {
			pages := make([]string, 0, len(g.Pages))
			for _, page := range g.Pages {
				pages = append(pages, page.PageName)
			}
			subject = strings.Join(pages, ",")
		}
	case gh.InstallationEvent, gh.IntegrationInstallationEvent:
		if i, ok := payload.(gh.InstallationPayload); ok {
			subject = string(i.Installation.ID)
		}
	case gh.IssueCommentEvent:
		if i, ok := payload.(gh.IssueCommentPayload); ok {
			// E.g., https://github.com/Codertocat/Hello-World/issues/2#issuecomment-393304133
			// and we keep with 2#issuecomment-393304133
			subject = lastPathPortion(i.Comment.HTMLURL)
		}
	case gh.IssuesEvent:
		if i, ok := payload.(gh.IssuesPayload); ok {
			subject = string(i.Issue.Number)
		}
	case gh.LabelEvent:
		if l, ok := payload.(gh.LabelPayload); ok {
			// E.g., :bug: Bugfix
			subject = l.Label.Name
		}
	case gh.MemberEvent:
		if m, ok := payload.(gh.MemberPayload); ok {
			subject = string(m.Member.ID)
		}
	case gh.MembershipEvent:
		if m, ok := payload.(gh.MembershipPayload); ok {
			subject = string(m.Member.ID)
		}
	case gh.MilestoneEvent:
		if m, ok := payload.(gh.MilestonePayload); ok {
			subject = string(m.Milestone.Number)
		}
	case gh.OrganizationEvent:
		if o, ok := payload.(gh.OrganizationPayload); ok {
			// The action that was performed, can be member_added, member_removed, or member_invited.
			subject = o.Action
		}
	case gh.OrgBlockEvent:
		if o, ok := payload.(gh.OrgBlockPayload); ok {
			// The action performed, can be blocked or unblocked.
			subject = o.Action
		}
	case gh.PageBuildEvent:
		if p, ok := payload.(gh.PageBuildPayload); ok {
			subject = string(p.ID)
		}
	case gh.PingEvent:
		if p, ok := payload.(gh.PingPayload); ok {
			subject = string(p.HookID)
		}
	case gh.ProjectCardEvent:
		if p, ok := payload.(gh.ProjectCardPayload); ok {
			// The action performed on the project card, can be created, edited, moved, converted, or deleted.
			subject = p.Action
		}
	case gh.ProjectColumnEvent:
		if p, ok := payload.(gh.ProjectColumnPayload); ok {
			// The action performed on the project column, can be created, edited, moved, converted, or deleted.
			subject = p.Action
		}
	case gh.ProjectEvent:
		if p, ok := payload.(gh.ProjectPayload); ok {
			// The action that was performed on the project, can be created, edited, closed, reopened, or deleted.
			subject = p.Action
		}
	case gh.PublicEvent:
		if p, ok := payload.(gh.PublicPayload); ok {
			subject = string(p.Repository.ID)
		}
	case gh.PullRequestEvent:
		if p, ok := payload.(gh.PullRequestPayload); ok {
			subject = string(p.PullRequest.Number)
		}
	case gh.PullRequestReviewEvent:
		if p, ok := payload.(gh.PullRequestReviewPayload); ok {
			subject = string(p.Review.ID)
		}
	case gh.PullRequestReviewCommentEvent:
		if p, ok := payload.(gh.PullRequestReviewCommentPayload); ok {
			subject = string(p.Comment.ID)
		}
	case gh.PushEvent:
		if p, ok := payload.(gh.PushPayload); ok {
			subject = fmt.Sprintf("%s...%s", p.Before, p.After)
		}
	case gh.ReleaseEvent:
		if r, ok := payload.(gh.ReleasePayload); ok {
			subject = r.Release.TagName
		}
	case gh.RepositoryEvent:
		if r, ok := payload.(gh.RepositoryPayload); ok {
			subject = string(r.Repository.ID)
		}
	case gh.StatusEvent:
		if s, ok := payload.(gh.StatusPayload); ok {
			subject = s.Sha
		}
	case gh.TeamEvent:
		if t, ok := payload.(gh.TeamPayload); ok {
			subject = string(t.Team.ID)
		}
	case gh.TeamAddEvent:
		if t, ok := payload.(gh.TeamAddPayload); ok {
			subject = string(t.Repository.ID)
		}
	case gh.WatchEvent:
		if w, ok := payload.(gh.WatchPayload); ok {
			subject = string(w.Repository.ID)
		}
	}
	return subject
}

func lastPathPortion(url string) string {
	var subject string
	index := strings.LastIndex(url, "/")
	if index != -1 {
		// Keep the last part.
		subject = url[index+1:]
	}
	return subject
}
