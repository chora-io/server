// Package regen includes the IRI defined within Regen Ledger:
// https://github.com/regen-network/regen-ledger/tree/v5.1.2/x/data
//
// This version modifies the original to support custom prefixes and
// to limit the amount of dependencies.
package regen

import (
	"fmt"
	"reflect"
)

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
	err := chr.DigestAlgorithm.Validate(chr.Hash)
	if err != nil {
		return err
	}

	return chr.MediaType.Validate()
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
	err := chg.DigestAlgorithm.Validate(chg.Hash)
	if err != nil {
		return err
	}

	err = chg.CanonicalizationAlgorithm.Validate()
	if err != nil {
		return err
	}

	return chg.MerkleTree.Validate()
}

// DigestAlgorithm is the digest algorithm.
type DigestAlgorithm int32

const (
	DigestAlgorithm_DIGEST_ALGORITHM_UNSPECIFIED DigestAlgorithm = 0
	DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256 DigestAlgorithm = 1
)

var DigestAlgorithmLength = map[DigestAlgorithm]int{
	DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256: 256,
}

func (m DigestAlgorithm) String() string { return "DigestAlgorithm" }

// Validate validates DigestAlgorithm.
func (da DigestAlgorithm) Validate(hash []byte) error {
	if reflect.DeepEqual(hash, []byte(nil)) {
		return fmt.Errorf("hash cannot be empty")
	}

	if da == DigestAlgorithm_DIGEST_ALGORITHM_UNSPECIFIED {
		return fmt.Errorf("invalid %T %s", da, da)
	}

	nBits, ok := DigestAlgorithmLength[da]
	if !ok {
		return fmt.Errorf("unknown %T %s", da, da)
	}

	nBytes := nBits / 8
	if len(hash) != nBytes {
		return fmt.Errorf("expected %d bytes for %s, got %d", nBytes, da, len(hash))
	}

	return nil
}

// GraphCanonicalizationAlgorithm is the graph canonicalization algorithm.
type GraphCanonicalizationAlgorithm int32

const (
	GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_UNSPECIFIED GraphCanonicalizationAlgorithm = 0
	GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015   GraphCanonicalizationAlgorithm = 1
)

var GraphCanonicalizationAlgorithm_name = map[int32]string{
	0: "GRAPH_CANONICALIZATION_ALGORITHM_UNSPECIFIED",
	1: "GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015",
}

func (gca GraphCanonicalizationAlgorithm) String() string { return "GraphCanonicalizationAlgorithm" }

// Validate validates GraphCanonicalizationAlgorithm.
func (gca GraphCanonicalizationAlgorithm) Validate() error {
	if _, ok := GraphCanonicalizationAlgorithm_name[int32(gca)]; !ok {
		return fmt.Errorf("unknown %T %d", gca, gca)
	}

	if gca == GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_UNSPECIFIED {
		return fmt.Errorf("invalid %T %s", gca, gca)
	}

	return nil
}

// GraphMerkleTree is the graph merkle tree type.
type GraphMerkleTree int32

const (
	GraphMerkleTree_GRAPH_MERKLE_TREE_NONE_UNSPECIFIED GraphMerkleTree = 0
)

var GraphMerkleTree_name = map[int32]string{
	0: "GRAPH_MERKLE_TREE_NONE_UNSPECIFIED",
}

// Validate validates GraphMerkleTree.
func (gmt GraphMerkleTree) Validate() error {
	if _, ok := GraphMerkleTree_name[int32(gmt)]; !ok {
		return fmt.Errorf("unknown %T %d", gmt, gmt)
	}

	return nil
}

// RawMediaType is the raw media type.
type RawMediaType int32

var RawMediaType_name = map[int32]string{
	0:  "RAW_MEDIA_TYPE_UNSPECIFIED",
	1:  "RAW_MEDIA_TYPE_TEXT_PLAIN",
	2:  "RAW_MEDIA_TYPE_JSON",
	3:  "RAW_MEDIA_TYPE_CSV",
	4:  "RAW_MEDIA_TYPE_XML",
	5:  "RAW_MEDIA_TYPE_PDF",
	16: "RAW_MEDIA_TYPE_TIFF",
	17: "RAW_MEDIA_TYPE_JPG",
	18: "RAW_MEDIA_TYPE_PNG",
	19: "RAW_MEDIA_TYPE_SVG",
	20: "RAW_MEDIA_TYPE_WEBP",
	21: "RAW_MEDIA_TYPE_AVIF",
	22: "RAW_MEDIA_TYPE_GIF",
	23: "RAW_MEDIA_TYPE_APNG",
	32: "RAW_MEDIA_TYPE_MPEG",
	33: "RAW_MEDIA_TYPE_MP4",
	34: "RAW_MEDIA_TYPE_WEBM",
	35: "RAW_MEDIA_TYPE_OGG",
}

func (rmt RawMediaType) Validate() error {
	if _, ok := RawMediaType_name[int32(rmt)]; !ok {
		return fmt.Errorf("unknown %T %d", rmt, rmt)
	}

	return nil
}

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

var mediaExtensionTypeToString = map[RawMediaType]string{
	RawMediaType_RAW_MEDIA_TYPE_UNSPECIFIED: "bin",
	RawMediaType_RAW_MEDIA_TYPE_TEXT_PLAIN:  "txt",
	RawMediaType_RAW_MEDIA_TYPE_CSV:         "csv",
	RawMediaType_RAW_MEDIA_TYPE_JSON:        "json",
	RawMediaType_RAW_MEDIA_TYPE_XML:         "xml",
	RawMediaType_RAW_MEDIA_TYPE_PDF:         "pdf",
	RawMediaType_RAW_MEDIA_TYPE_TIFF:        "tiff",
	RawMediaType_RAW_MEDIA_TYPE_JPG:         "jpg",
	RawMediaType_RAW_MEDIA_TYPE_PNG:         "png",
	RawMediaType_RAW_MEDIA_TYPE_SVG:         "svg",
	RawMediaType_RAW_MEDIA_TYPE_WEBP:        "webp",
	RawMediaType_RAW_MEDIA_TYPE_AVIF:        "avif",
	RawMediaType_RAW_MEDIA_TYPE_GIF:         "gif",
	RawMediaType_RAW_MEDIA_TYPE_APNG:        "apng",
	RawMediaType_RAW_MEDIA_TYPE_MPEG:        "mpeg",
	RawMediaType_RAW_MEDIA_TYPE_MP4:         "mp4",
	RawMediaType_RAW_MEDIA_TYPE_WEBM:        "webm",
	RawMediaType_RAW_MEDIA_TYPE_OGG:         "ogg",
}

var stringToMediaExtensionType = map[string]RawMediaType{}

func init() {
	for mt, ext := range mediaExtensionTypeToString {
		stringToMediaExtensionType[ext] = mt
	}
}
