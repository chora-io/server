package app

type GetDataResponse struct {
	// Id is the unique identifier of the data.
	Id int32 `json:"id"`

	// Canon is the canonicalization of the data.
	Canon string `json:"canon"`

	// Context is the schema context of the data.
	Context string `json:"context"`

	// Jsonld is the JSON-LD representation of the data.
	Jsonld string `json:"jsonld"`
}

func NewGetDataResponse(id int32, canon, context, jsonld string) GetDataResponse {
	return GetDataResponse{
		Id:      id,
		Canon:   canon,
		Context: context,
		Jsonld:  jsonld,
	}
}

type PostDataResponse struct {
	// Id is the unique identifier of the data.
	Id int32 `json:"id"`

	// Canon is the canonicalization of the data.
	Canon string `json:"canon"`

	// Context is the schema context of the data.
	Context string `json:"context"`

	// Jsonld is the JSON-LD representation of the data.
	Jsonld string `json:"jsonld"`
}

func NewPostDataResponse(id int32, canon, context, jsonld string) PostDataResponse {
	return PostDataResponse{
		Id:      id,
		Canon:   canon,
		Context: context,
		Jsonld:  jsonld,
	}
}
