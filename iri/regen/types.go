// Package regen implements the IRI defined within Regen Ledger:
// https://github.com/regen-network/regen-ledger/tree/v5.1.2/x/data
//
// This version modifies the original to support additional configuration
// options including custom prefixes and IRI versioning.
package regen

// ContentHash specifies a hash-based content identifier for a piece of data.
type ContentHash struct {
	// Raw specifies "raw" data which does not specify a deterministic, canonical
	// encoding. Users of these hashes MUST maintain a copy of the hashed data
	// which is preserved bit by bit. All other content encodings specify a
	// deterministic, canonical encoding allowing implementations to choose from a
	// variety of alternative formats for transport and encoding while maintaining
	// the guarantee that the canonical hash will not change. The media type for
	// "raw" data is defined by the MediaType enum.
	Raw *ContentHash_Raw `protobuf:"bytes,1,opt,name=raw,proto3" json:"raw,omitempty"`
	// Graph specifies graph data that conforms to the RDF data model.
	// The canonicalization algorithm used for an RDF graph is specified by
	// GraphCanonicalizationAlgorithm.
	Graph *ContentHash_Graph `protobuf:"bytes,2,opt,name=graph,proto3" json:"graph,omitempty"`
}

// ContentHash_Raw is a raw content hash.
type ContentHash_Raw struct {
	// hash represents the hash of the data based on the specified
	// digest_algorithm.
	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	// digest_algorithm represents the hash digest algorithm.
	DigestAlgorithm DigestAlgorithm `protobuf:"varint,2,opt,name=digest_algorithm,json=digestAlgorithm,proto3,enum=regen.data.v1.DigestAlgorithm" json:"digest_algorithm,omitempty"`
	// media_type represents the media type for raw data.
	MediaType RawMediaType `protobuf:"varint,3,opt,name=media_type,json=mediaType,proto3,enum=regen.data.v1.RawMediaType" json:"media_type,omitempty"`
}

// Validate validates ContentHash_Raw.
func (chr ContentHash_Raw) Validate() error {
	return nil
}

// ContentHash_Graph is a graph content hash.
type ContentHash_Graph struct {
	// hash represents the hash of the data based on the specified
	// digest_algorithm.
	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	// digest_algorithm represents the hash digest algorithm.
	DigestAlgorithm DigestAlgorithm `protobuf:"varint,2,opt,name=digest_algorithm,json=digestAlgorithm,proto3,enum=regen.data.v1.DigestAlgorithm" json:"digest_algorithm,omitempty"`
	// graph_canonicalization_algorithm represents the RDF graph
	// canonicalization algorithm.
	CanonicalizationAlgorithm GraphCanonicalizationAlgorithm `protobuf:"varint,3,opt,name=canonicalization_algorithm,json=canonicalizationAlgorithm,proto3,enum=regen.data.v1.GraphCanonicalizationAlgorithm" json:"canonicalization_algorithm,omitempty"`
	// merkle_tree is the merkle tree type used for the graph hash, if any.
	MerkleTree GraphMerkleTree `protobuf:"varint,4,opt,name=merkle_tree,json=merkleTree,proto3,enum=regen.data.v1.GraphMerkleTree" json:"merkle_tree,omitempty"`
}

// Validate validates ContentHash_Graph.
func (chg ContentHash_Graph) Validate() error {
	return nil
}

// DigestAlgorithm is the digest algorithm.
type DigestAlgorithm int32

const (
	DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256 DigestAlgorithm = 1
)

// Validate validates DigestAlgorithm.
func (a DigestAlgorithm) Validate(hash []byte) error {
	return nil
}

// GraphCanonicalizationAlgorithm is the graph canonicalization algorithm.
type GraphCanonicalizationAlgorithm int32

const (
	GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015 GraphCanonicalizationAlgorithm = 1
)

// Validate validates GraphCanonicalizationAlgorithm.
func (a GraphCanonicalizationAlgorithm) Validate() error {
	return nil
}

// GraphMerkleTree is the graph merkle tree type.
type GraphMerkleTree int32

const (
	GraphMerkleTree_GRAPH_MERKLE_TREE_NONE_UNSPECIFIED GraphMerkleTree = 0
)

// Validate validates GraphMerkleTree.
func (t GraphMerkleTree) Validate() error {
	return nil
}

// RawMediaType is the raw media type.
type RawMediaType int32

const (
	RawMediaType_RAW_MEDIA_TYPE_UNSPECIFIED RawMediaType = 0
	RawMediaType_RAW_MEDIA_TYPE_TEXT_PLAIN  RawMediaType = 1
	RawMediaType_RAW_MEDIA_TYPE_JSON        RawMediaType = 2
	RawMediaType_RAW_MEDIA_TYPE_CSV         RawMediaType = 3
	RawMediaType_RAW_MEDIA_TYPE_XML         RawMediaType = 4
	RawMediaType_RAW_MEDIA_TYPE_PDF         RawMediaType = 5
	RawMediaType_RAW_MEDIA_TYPE_TIFF        RawMediaType = 16
	RawMediaType_RAW_MEDIA_TYPE_JPG         RawMediaType = 17
	RawMediaType_RAW_MEDIA_TYPE_PNG         RawMediaType = 18
	RawMediaType_RAW_MEDIA_TYPE_SVG         RawMediaType = 19
	RawMediaType_RAW_MEDIA_TYPE_WEBP        RawMediaType = 20
	RawMediaType_RAW_MEDIA_TYPE_AVIF        RawMediaType = 21
	RawMediaType_RAW_MEDIA_TYPE_GIF         RawMediaType = 22
	RawMediaType_RAW_MEDIA_TYPE_APNG        RawMediaType = 23
	RawMediaType_RAW_MEDIA_TYPE_MPEG        RawMediaType = 32
	RawMediaType_RAW_MEDIA_TYPE_MP4         RawMediaType = 33
	RawMediaType_RAW_MEDIA_TYPE_WEBM        RawMediaType = 34
	RawMediaType_RAW_MEDIA_TYPE_OGG         RawMediaType = 35
)
