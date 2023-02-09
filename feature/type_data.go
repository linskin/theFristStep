package feature

type Response struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

//	type S_Author struct {
//		ID            int    `json:"id"`
//		Name          string `json:"name"`
//		FollowCount   int    `json:"follow_count" gorm:"default:0"`
//		FollowerCount int    `json:"follower_count" gorm:"default:0"`
//		IsFollow      bool   `json:"is_follow"`
//	}
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
	Token         string `json:"-"`
	V_key         int    `json:"-" gorm:"ForeignKey"`
}

type Video struct {
	ID            int    `json:"id"`
	Author        User   `json:"author"`
	PlayURL       string `json:"play_url"`
	CoverURL      string `json:"cover_url"`
	FavoriteCount int    `json:"favorite_count"`
	CommentCount  int    `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
	UID           int
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

// 关于评论
type CommentResponse struct {
	StatusCode  int       `json:"status_code"`
	StatusMsg   string    `json:"status_msg"`
	CommentList []Comment `json:"comment_list"`
}

type Comment struct {
	ID         int    `json:"id"`
	User       User   `json:"user"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}

type Commentaction struct {
	StatusCode int     `json:"status_code"`
	StatusMsg  string  `json:"status_msg"`
	Comment    Comment `json:"comment"`
}
