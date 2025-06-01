package main

type Params struct {
    UserID     string   `json:"user_id"`
    MathParams []string `json:"math_params"`
}

type MathParam struct {
    UserID string `json:"user_id"`
    Param  string `json:"-"`
}
