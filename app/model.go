package app

type GetDataResponse struct {
	// Iri is the unique identifier of the data.
	Iri string `json:"iri"`

	// Context is the schema context of the data.
	Context string `json:"context"`

	// Jsonld is the JSON-LD representation of the data.
	Jsonld string `json:"jsonld"`
}

func NewGetDataResponse(iri, context, jsonld string) GetDataResponse {
	return GetDataResponse{
		Iri:     iri,
		Context: context,
		Jsonld:  jsonld,
	}
}

type PostDataRequest struct {
	// Digest is the digest algorithm used to generate an IRI.
	Digest string `json:"digest"`

	// Canon is the canonicalization algorithm used to generate an IRI.
	Canon string `json:"canon"`

	// Merkle is the merkle tree type used to generate an IRI.
	Merkle string `json:"merkle"`

	// Context is the schema context of the data.
	Context string `json:"context"`

	// Jsonld is the JSON-LD representation of the data.
	Jsonld string `json:"jsonld"`
}

type PostDataResponse struct {
	// Iri is the unique identifier of the data.
	Iri string `json:"iri"`

	// Context is the schema context of the data.
	Context string `json:"context"`

	// Jsonld is the JSON-LD representation of the data.
	Jsonld string `json:"jsonld"`
}

func NewPostDataResponse(iri, context, jsonld string) PostDataResponse {
	return PostDataResponse{
		Iri:     iri,
		Context: context,
		Jsonld:  jsonld,
	}
}
