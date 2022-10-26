package types

import "time"

type CodeCountResponse struct {
	Count       int         `json:"count"`
	Mode        int         `json:"mode"`
	QueryErrors interface{} `json:"query_errors"`
}

type MetricResponse struct {
	Message  string            `json:"message"`
	Count    int               `json:"count"`
	Metadata map[string]string `json:"metadata"`
}

type IssuesSearch struct {
	TotalCount        int     `json:"total_count"`
	IncompleteResults bool    `json:"incomplete_results"`
	Items             []Issue `json:"items"`
}

type Issue struct {
	URL                   string        `json:"url"`
	RepositoryURL         string        `json:"repository_url"`
	LabelsURL             string        `json:"labels_url"`
	CommentsURL           string        `json:"comments_url"`
	EventsURL             string        `json:"events_url"`
	HTMLURL               string        `json:"html_url"`
	ID                    int           `json:"id"`
	NodeID                string        `json:"node_id"`
	Number                int           `json:"number"`
	Title                 string        `json:"title"`
	User                  User          `json:"user"`
	Labels                []Label       `json:"labels"`
	State                 string        `json:"state"`
	Locked                bool          `json:"locked"`
	Assignee              interface{}   `json:"assignee"`
	Assignees             []interface{} `json:"assignees"`
	Milestone             interface{}   `json:"milestone"`
	Comments              int           `json:"comments"`
	CreatedAt             time.Time     `json:"created_at"`
	UpdatedAt             time.Time     `json:"updated_at"`
	ClosedAt              interface{}   `json:"closed_at"`
	AuthorAssociation     string        `json:"author_association"`
	ActiveLockReason      interface{}   `json:"active_lock_reason"`
	Draft                 bool          `json:"draft,omitempty"`
	PullRequest           PullRequest   `json:"pull_request,omitempty"`
	Body                  string        `json:"body"`
	Reactions             Reactions     `json:"reactions"`
	TimelineURL           string        `json:"timeline_url"`
	PerformedViaGithubApp interface{}   `json:"performed_via_github_app"`
	StateReason           interface{}   `json:"state_reason"`
	Score                 float64       `json:"score"`
}

type Label struct {
	ID          int64  `json:"id"`
	NodeID      string `json:"node_id"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	Default     bool   `json:"default"`
	Description string `json:"description"`
}

type PullRequest struct {
	URL      string      `json:"url"`
	HTMLURL  string      `json:"html_url"`
	DiffURL  string      `json:"diff_url"`
	PatchURL string      `json:"patch_url"`
	MergedAt interface{} `json:"merged_at"`
}

type Reactions struct {
	URL        string `json:"url"`
	TotalCount int    `json:"total_count"`
	Num1       int    `json:"+1"`
	Num10      int    `json:"-1"`
	Laugh      int    `json:"laugh"`
	Hooray     int    `json:"hooray"`
	Confused   int    `json:"confused"`
	Heart      int    `json:"heart"`
	Rocket     int    `json:"rocket"`
	Eyes       int    `json:"eyes"`
}

type User struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}
