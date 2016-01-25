package api


type LoginData struct {
  Username string, `json:"username"`
  Password string, `json:"password"`
}

type CallData struct {
  Number string, `json:"number"`
}

type EndData struct {
  Lock string, `json:"lock"`
}
