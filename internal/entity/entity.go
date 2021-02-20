package entity

import "time"

type Slot struct {
	ID   int
	Name string
}

type Banner struct {
	ID   int
	Name string
}

type BannerStat struct {
	Banner     Banner
	ShowCount  int
	ClickCount int
}

type ClickEvent struct {
	BannerID    int
	SlotID      int
	UserGroupID int
	Created     time.Time
}

type ShowEvent struct {
	BannerID    int
	SlotID      int
	UserGroupID int
	Created     time.Time
}

type UserGroup struct {
	ID   int
	Name string
}
