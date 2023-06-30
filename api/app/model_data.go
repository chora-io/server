package app

type AuthResponse struct {
	// Token is the authentication token.
	Token string `json:"token"`
}

func NewAuthResponse(token string) AuthResponse {
	return AuthResponse{
		Token: token,
	}
}

type GetDataResponse struct {
	// Iri is the unique identifier of the data.
	Iri string `json:"iri"`

	// Jsonld is the JSON-LD representation of the data.
	Jsonld string `json:"jsonld"`
}

func NewGetDataResponse(iri, jsonld string) GetDataResponse {
	return GetDataResponse{
		Iri:    iri,
		Jsonld: jsonld,
	}
}

type PostDataRequest struct {
	// Digest is the digest algorithm used to generate an IRI.
	Digest string `json:"digest"`

	// Canon is the canonicalization algorithm used to generate an IRI.
	Canon string `json:"canon"`

	// Merkle is the merkle tree type used to generate an IRI.
	Merkle string `json:"merkle"`

	// Jsonld is the JSON-LD representation of the data.
	Jsonld string `json:"jsonld"`
}

type PostDataResponse struct {
	// Iri is the unique identifier of the data.
	Iri string `json:"iri"`

	// Jsonld is the JSON-LD representation of the data.
	Jsonld string `json:"jsonld"`
}

func NewPostDataResponse(iri, jsonld string) PostDataResponse {
	return PostDataResponse{
		Iri:    iri,
		Jsonld: jsonld,
	}
}
