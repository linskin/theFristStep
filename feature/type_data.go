package feature

type S_Author struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	FollowCount   int    `json:"follow_count"`
	FollowerCount int    `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type Video struct {
	ID            int      `json:"id"`
	Author        S_Author `json:"author"`
	PlayURL       string   `json:"play_url"`
	CoverURL      string   `json:"cover_url"`
	FavoriteCount int      `json:"favorite_count"`
	CommentCount  int      `json:"comment_count"`
	IsFavorite    bool     `json:"is_favorite"`
	Title         string   `json:"title"`
}

type Video_Feed struct {
	StatusCode int     `json:"status_code"`
	StatusMsg  string  `json:"status_msg"`
	NextTime   int64   `json:"next_time"`
	Vlist      []Video `json:"video_list"`
}

type LikeAction struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type UserLR struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserID     int    `json:"user_id" gorm:"primaryKey"`
	Token      string `json:"token"`
}

type UserInfo struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	User       User   `json:"user"`
}
type User struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	FollowCount   int    `json:"follow_count" gorm:"default:0"`
	FollowerCount int    `json:"follower_count" gorm:"default:0"`
	IsFollow      bool   `json:"is_follow"`
	Token         string
}
