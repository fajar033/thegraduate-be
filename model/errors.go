package model

type NotFoundError struct {
	Message string `json:"message"`
}

func (n *NotFoundError) Error() string {

	return n.Message
}

type ValidationError struct {
	ErrMessage any    `json:"errMessage"`
	Message    string `json:"message"`
}

func (h *ValidationError) Error() string {
	return h.Message
}

type ConflictError struct {
	Message string `json:"message"`
}
type BadRequestError struct {
	Message string `json:"message"`
}

func (c *BadRequestError) Error() string {
	return c.Message
}

func (c *ConflictError) Error() string {
	return c.Message
}
