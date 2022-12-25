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

type Response struct {
    Status string  `json:"status"`
    Error  string  `json:"error"`
    Data   RBCNews `json:"data"`
}
