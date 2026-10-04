package model

type CommunityDetail struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Images      *string `json:"images"`
	Description *string `json:"description"`
	UserID      *int    `json:"user_id"`
	MemberCount int64   `json:"member_count"`
	CategoryIDs []int32 `json:"category_ids"`
}
