package models

type Task struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    IsCompleted bool   `json:"isCompleted"`
}