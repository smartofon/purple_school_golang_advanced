package verify

type SendRequest struct {
	Email string `json:"email" validate:"email,required"`
}

type VerifyRequest struct{}

type SendResponce struct {
	Result string `json:"result"`
}

type VerifyResponce struct {
	Result string `json:"result"`
}
