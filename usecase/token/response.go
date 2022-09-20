package token

type ResultResponse struct {
	Name      string `json:"name"`
	Token     string `json:"token"`
	ExpiredAt string `json:"expired_at"`
}
