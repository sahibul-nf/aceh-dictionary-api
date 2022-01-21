package advice

type QueryInput struct {
	Input string `json:"input" binding:"required"`
}
