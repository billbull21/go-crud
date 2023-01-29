// models/post.go
package models

import "time"

type Post struct {
    ID        		uint      `json:"id" gorm:"primaryKey"`
    Title     		string    `json:"title"`
    Description		string    `json:"description"`
    Published		bool      `json:"published"`
    CreatedAt 		time.Time `json:"created_at"`
    UpdatedAt 		time.Time `json:"updated_at"`
}