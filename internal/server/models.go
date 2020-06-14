package server

type (
	Failure struct {
		Error error `json:"error"`
	}

	Success struct {
		Result string `json:"result"`
	}
)
