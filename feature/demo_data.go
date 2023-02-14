package feature

var DemoAuthor = User{
	ID:            3,
	Name:          "SWAG",
	FollowCount:   19,
	FollowerCount: 9999,
	IsFollow:      true,
	Token:         "NSWAGP123456",
}

/*
ID            int      `json:"id"`
Author        S_Author `json:"author"`
PlayURL       string   `json:"play_url"`
CoverURL      string   `json:"cover_url"`
FavoriteCount int      `json:"favorite_count"`
CommentCount  int      `json:"comment_count"`
IsFavorite    bool     `json:"is_favorite"`
Title         string   `json:"title"`
*/
var DemoVideo = []Video{
	{
		ID:            1,
		Author:        DemoAuthor,
		PlayURL:       "http://192.168.1.102:9090/Dazhuojun.mp4",
		CoverURL:      "https://img0.baidu.com/it/u=3294539948,324399065&fm=253&fmt=auto&app=138&f=JPEG?w=822&h=500",
		FavoriteCount: 999,
		CommentCount:  2,
		IsFavorite:    false,
		Title:         "amazing movie!",
	},
}

var DemoComment = []Comment{
	{
		ID:         3,
		User:       DemoAuthor,
		Content:    "会场厉害!!!太像肖酱",
		CreateDate: "2023-2-9",
	},
}
