// Package regen includes the IRI defined within Regen Ledger:
// https://github.com/regen-network/regen-ledger/tree/v5.1.2/x/data
//
// This version modifies the original to support custom prefixes and
// limits the amount of dependencies required in chora server.
package regen

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/cosmos/btcutil/base58"
)

const (
	IriVersion0 byte = 0
)

const (
	IriRaw   byte = 0
	IriGraph byte = 1
)

// ToIRI converts the ContentHash to an IRI (internationalized URI) using the regen IRI scheme.
// A ContentHash IRI will look something like regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.rdf
// which is some base58check encoded data followed by a file extension or pseudo-extension.
// See ContentHash_Raw.ToIRI and ContentHash_Graph.ToIRI for more details on specific formatting.
func (r ContentHash) ToIRI(prefix string) (string, error) {
	if chr := r.Raw; chr != nil {
		return chr.ToIRI(prefix)
	} else if chg := r.Graph; chg != nil {
		return chg.ToIRI(prefix)
	}
	return "", fmt.Errorf("invalid %T", r)
}

// ToIRI converts the ContentHash_Raw to an IRI (internationalized URI) based on the following
// pattern: regen:{base58check(concat( byte(0x0), byte(digest_algorithm), hash))}.{media_type extension}
func (chr ContentHash_Raw) ToIRI(prefix string) (string, error) {
	err := chr.Validate()
	if err != nil {
		return "", err
	}

	bz := make([]byte, len(chr.Hash)+2)
	bz[0] = IriRaw
	bz[1] = byte(chr.DigestAlgorithm)
	copy(bz[2:], chr.Hash)

	// only one version for now
	hashStr := base58.CheckEncode(bz, IriVersion0)

	ext, err := chr.MediaType.ToExtension()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s:%s.%s", prefix, hashStr, ext), nil
}

// ToIRI converts the ContentHash_Graph to an IRI (internationalized URI) based on the following
// pattern: regen:{base58check(concat(byte(0x1), byte(canonicalization_algorithm),
// byte(merkle_tree), byte(digest_algorithm), hash))}.rdf
func (chg ContentHash_Graph) ToIRI(prefix string) (string, error) {
	err := chg.Validate()
	if err != nil {
		return "", err
	}

	bz := make([]byte, len(chg.Hash)+4)
	bz[0] = IriGraph
	bz[1] = byte(chg.CanonicalizationAlgorithm)
	bz[2] = byte(chg.MerkleTree)
	bz[3] = byte(chg.DigestAlgorithm)
	copy(bz[4:], chg.Hash)

	// only one version for now
	hashStr := base58.CheckEncode(bz, IriVersion0)

	return fmt.Sprintf("%s:%s.rdf", prefix, hashStr), nil
}

// ToExtension converts the media type to a file extension based on the mediaTypeExtensions map.
func (rmt RawMediaType) ToExtension() (string, error) {
	ext, ok := mediaExtensionTypeToString[rmt]
	if !ok {
		return "", fmt.Errorf("missing extension for %T %d", rmt, rmt)
	}

	return ext, nil
}

// ParseIRI parses an IRI string representation of a ContentHash into a ContentHash struct
// IRIs must have a prefix (e.g. "regen:"), and only ContentHash_Graph and ContentHash_Raw
// are supported.
func ParseIRI(iri string) (*ContentHash, error) {
	if iri == "" {
		return nil, fmt.Errorf("failed to parse IRI: empty string is not allowed")
	}

	splitPre := strings.Split(iri, ":")
	if len(splitPre) < 2 {
		return nil, fmt.Errorf(
			"failed to parse IRI %s: IRI without a prefix is not allowed", iri,
		)
	}
	if len(splitPre) > 2 {
		return nil, fmt.Errorf(
			"failed to parse IRI %s: IRI with multiple prefixes is not allowed", iri,
		)
	}
	if len(splitPre[0]) == 0 {
		return nil, fmt.Errorf(
			"failed to parse IRI %s: IRI with empty prefix not allowed", iri,
		)
	}

	splitExt := strings.Split(iri, ".")
	if len(splitExt) < 2 {
		return nil, fmt.Errorf(
			"failed to parse IRI %s: IRI without an extension is not allowed", iri,
		)
	}
	if len(splitExt) > 2 {
		return nil, fmt.Errorf(
			"failed to parse IRI %s: IRI with multiple extensions is not allowed", iri,
		)
	}
	if len(splitExt[1]) == 0 {
		return nil, fmt.Errorf(
			"failed to parse IRI %s: IRI with empty extension not allowed", iri,
		)
	}

	splitExtWithoutPre := strings.Split(splitPre[1], ".")
	hashPart, ext := splitExtWithoutPre[0], splitExtWithoutPre[1]

	res, version, err := base58.CheckDecode(hashPart)
	if err != nil {
		return nil, fmt.Errorf("failed to parse IRI %s: %s", iri, err)
	}

	// only one version supported at this time
	if version != IriVersion0 {
		return nil, fmt.Errorf("failed to parse IRI %s: invalid version", iri)
	}

	rdr := bytes.NewBuffer(res)

	// read first byte
	typ, err := rdr.ReadByte()
	if err != nil {
		return nil, err
	}

	// switch on first byte which represents the type prefix
	switch typ {
	case IriRaw:
		// read next byte
		b0, err := rdr.ReadByte()
		if err != nil {
			return nil, err
		}

		// look up extension as media type
		mediaType, ok := stringToMediaExtensionType[ext]
		if !ok {
			return nil, fmt.Errorf("failed to resolve media type for extension %s, expected %s", ext, mediaExtensionTypeToString[mediaType])
		}

		// interpret next byte as digest algorithm
		digestAlg := DigestAlgorithm(b0)
		hash := rdr.Bytes()
		err = digestAlg.Validate(hash)
		if err != nil {
			return nil, err
		}

		return &ContentHash{Raw: &ContentHash_Raw{
			Hash:            hash,
			DigestAlgorithm: digestAlg,
			MediaType:       mediaType,
		}}, nil

	case IriGraph:
		// rdf extension is expected for graph data
		if ext != "rdf" {
			return nil, fmt.Errorf("invalid extension .%s for graph data, expected .rdf", ext)
		}

		// read next byte
		b0, err := rdr.ReadByte()
		if err != nil {
			return nil, err
		}

		// interpret next byte as canonicalization algorithm
		c14Alg := GraphCanonicalizationAlgorithm(b0)
		err = c14Alg.Validate()
		if err != nil {
			return nil, err
		}

		// read next byte
		b0, err = rdr.ReadByte()
		if err != nil {
			return nil, err
		}

		// interpret next byte as merklization algorithm
		mtAlg := GraphMerkleTree(b0)
		err = mtAlg.Validate()
		if err != nil {
			return nil, err
		}

		// read next byte
		b0, err = rdr.ReadByte()
		if err != nil {
			return nil, err
		}

		// interpret next byte as digest algorithm
		digestAlg := DigestAlgorithm(b0)
		hash := rdr.Bytes()
		err = digestAlg.Validate(hash)
		if err != nil {
			return nil, err
		}

		return &ContentHash{Graph: &ContentHash_Graph{
			Hash:                      hash,
			DigestAlgorithm:           digestAlg,
			CanonicalizationAlgorithm: c14Alg,
			MerkleTree:                mtAlg,
		}}, nil
	}

	return nil, fmt.Errorf("unable to parse IRI %s", iri)
}
