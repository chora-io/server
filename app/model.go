package app

type GetContentResponse struct {
	// Id is the unique identifier of the content.
	Id int64 `json:"id"`

	// Body is body of the content.
	Body string `json:"body"`
}

func NewGetContentResponse(id int64, body string) GetContentResponse {
	return GetContentResponse{
		Id:   id,
		Body: body,
	}
}

type PostContentResponse struct {
	// Id is the unique identifier of the content.
	Id int64 `json:"id"`

	// Body is body of the content.
	Body string `json:"body"`
}

func NewPostContentResponse(id int64, body string) PostContentResponse {
	return PostContentResponse{
		Id:   id,
		Body: body,
	}
}
