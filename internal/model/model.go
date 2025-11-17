package model

import "time"

type KinkList struct {
    ID        string      `json:"id"`
    Nickname  string      `json:"nickname"`
    Ratings   []Rating    `json:"ratings"`
    Blocks    []Block     `json:"blocks"`
    CreatedAt time.Time   `json:"created_at"`
    UpdatedAt time.Time   `json:"updated_at"`
}

type Rating struct {
    ID    string `json:"id"`
    Label string `json:"label"`
    Color string `json:"color"`
}

type Block struct {
    ID        string     `json:"id"`
    Title     string     `json:"title"`
    Comment   string     `json:"comment"`
    Groups    []Group    `json:"groups"`
    Questions []Question `json:"questions"`
}

type Group struct {
    ID    string `json:"id"`
    Label string `json:"label"`
}

type Question struct {
    ID       string   `json:"id"`
    Title    string   `json:"title"`
    Comment  string   `json:"comment"`
    Answers  []Answer `json:"answers"`
}

type Answer struct {
    GroupID  string `json:"groupId"`
    RatingID string `json:"ratingId"`
}
