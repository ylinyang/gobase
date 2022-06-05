package models

type Image struct {
    Id    int    `uri:"id"`
    Title string `json:"title"`
    Desc  string `json:"desc"`
}
