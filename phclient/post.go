package phclient

type PostsResponse struct {
	Posts []Post
}

type UserPostInteraction struct {
	CommentedOnPost bool `json:"commented_on_post"`
	VotedForPost    bool `json:"voted_for_post"`
}

type Post struct {
	CommentsCount int                 `json:"comments_count"`
	CreatedAt     string              `json:"created_at"`
	CurrentUser   UserPostInteraction `json:"current_user"`
	Day           string              `json:"day"`
	DiscussionUrl string              `json:"discussion_url"`
	Id            int                 `json:"id"`
	MakerInside   bool                `json:"maker_inside"`
	Makers        []Maker             `json:"makers"`
	Name          string              `json:"name"`
	RedirectUrl   string              `json:"redirect_url"`
	ScreenshotUrl []string            `json:"screenshot_url"`
	Tagline       string              `json:"tagline"`
	User          string              `json:"user"`
	VotesCount    int                 `json:"votes_count"`
}
