#A Golang Product Hunt API Client

This is a wrapper client for the Product Hunt API. It implements the most recent version, v1.

Currently, it can request the top stories for the current day as well as previous days.

##Usage
From a command line, run `$ go fetch github.com/munrocape/ph/client`
Next, you will need to configure two environment variables - PH_CLIENT_ID and PH_CLIENT_SECRET. You can determine what their values should be via the [Product Hunt Api Dashboard.](http://www.producthunt.com/v1/oauth/applications) Be sure to be logged into Product Hunt before accessing that URL.

You can then include it in any go file. 
```Go
import "github.com/munrocape/ph/client"
```

Below is a quick overview of the currently implemented methods.

```Go
c, _ := ph.NewClient()
c.GetPostsToday() // Get the posts from today
c.GetPostsOffset(10) // Get posts from ten days ago
```

##Structs
The three main structs are Posts, Users, and Collections. 

###Posts
A GetPosts* call returns a PostsResponse which is an array of Post structs.

```Go
type PostsResponse struct {
	Posts []Post
}
```

```Go
type Post struct {
	CommentsCount int                 `json:"comments_count"`
	CreatedAt     string              `json:"created_at"`
	CurrentUser   UserPostInteraction `json:"current_user"`
	Day           string              `json:"day"`
	DiscussionUrl string              `json:"discussion_url"`
	Id            int                 `json:"id"`
	MakerInside   bool                `json:"maker_inside"`
	Makers        []User              `json:"makers"`
	Name          string              `json:"name"`
	RedirectUrl   string              `json:"redirect_url"`
	ScreenshotUrl ScreenshotUrl       `json:"screenshot_url"`
	Tagline       string              `json:"tagline"`
	User          User                `json:"user"`
	VotesCount    int                 `json:"votes_count"`
}
```

A Post struct includes the UserPostInteraction and ScreenshotUrl structs.

```Go
type UserPostInteraction struct {
	CommentedOnPost bool `json:"commented_on_post"`
	VotedForPost    bool `json:"voted_for_post"`
}
```

```Go
type ScreenshotUrl struct {
	Px300 string `json:"300px"`
	Px850 string `json:"850px"`
}
```

###User
```Go
type User struct {
	CreatedAt  string            `json:"created_at"`
	Headline   string            `json:"headline"`
	Id         int               `json:"id"`
	Images     map[string]string `json:"images"`
	Name       string            `json:"name"`
	ProfileUrl string            `json:"profile_url"`
	Username   string            `json:"username"`
	WebsiteUrl string            `json:"website_url"`
}

```