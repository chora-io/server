package app

type GetIpfsResponse struct {
	// Cid is the unique identifier of the content.
	Cid string `json:"cid"`

	// Content is the requested content.
	Content string `json:"content"`
}

func NewGetIpfsResponse(cid, content string) GetIpfsResponse {
	return GetIpfsResponse{
		Cid:     cid,
		Content: content,
	}
}

type PostIpfsRequest struct {
	// Content is the requested content.
	Content string `json:"content"`
}

type PostIpfsResponse struct {
	// Cid is the unique identifier of the content.
	Cid string `json:"cid"`

	// Content is the requested content.
	Content string `json:"content"`
}

func NewPostIpfsResponse(cid, content string) PostIpfsResponse {
	return PostIpfsResponse{
		Cid:     cid,
		Content: content,
	}
}
