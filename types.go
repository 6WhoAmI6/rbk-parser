package main

type RBCNews struct {
    MainNews    News   `json:"main_news"`
    TopNews     []News `json:"top_news"`
    CentralNews []News `json:"central_news"`
}

type News struct {
    URL   string `json:"url"`
    Title string `json:"title"`
}
