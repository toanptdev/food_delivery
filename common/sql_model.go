package common

import "time"

type SqlModel struct {
	ID        int        `json:"id" form:"id"`
	Status    int        `json:"status" form:"status"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
