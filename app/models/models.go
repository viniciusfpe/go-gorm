package models

import (
    "time"
)

type Reminder struct {
    Id        int64     `json:"id"`
    Message   string    `sql:"size:1024" json:"message"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}