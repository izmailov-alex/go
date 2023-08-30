package models

type Segment struct {
	SegmentID   string `json:"segment_id"`
	SegmentName string `json:"segment_name"`
}

type UserSegments struct {
	UserID    string `json:"user_id"`
	SegmentID string `json:"segment_id"`
}

type Put struct {
	SegmentAdd    []string `json:"segment_add"`
	SegmentDelete []string `json:"segment_delete"`
}
