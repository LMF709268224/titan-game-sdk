package gamevrf

import (
	"encoding/binary"

	"github.com/zscboy/titan-game-sdk/vrf/filrpc"

	"github.com/filecoin-project/go-address"
	"github.com/minio/blake2b-simd"
	"golang.org/x/xerrors"
)

func drawRandomness(rbase []byte, pers DomainSeparationTag, height uint64, entropy []byte) ([]byte, error) {
	h := blake2b.New256()
	if err := binary.Write(h, binary.BigEndian, int64(pers)); err != nil {
		return nil, xerrors.Errorf("drawRandomness deriving randomness: %w", err)
	}
	VRFDigest := blake2b.Sum256(rbase)
	_, err := h.Write(VRFDigest[:])
	if err != nil {
		return nil, xerrors.Errorf("drawRandomness hashing VRFDigest: %w", err)
	}
	if err := binary.Write(h, binary.BigEndian, height); err != nil {
		return nil, xerrors.Errorf("drawRandomness deriving randomness: %w", err)
	}
	_, err = h.Write(entropy)
	if err != nil {
		return nil, xerrors.Errorf("drawRandomness hashing entropy: %w", err)
	}

	return h.Sum(nil), nil
}

func VerifyVRF(pubkey []byte,
	pers DomainSeparationTag, rbase []byte, entropy []byte, vrf *VRFOut) error {

	// draw randomness
	randomness, err := drawRandomness(rbase, pers, vrf.Height, entropy)
	if err != nil {
		return xerrors.Errorf("VerifyVRF drawRandomness failed: %w", err)
	}

	return blsVerify(pubkey, randomness, vrf.Proof)
}

func GenerateVRF(pers DomainSeparationTag,
	privateKey []byte, rbase []byte, height uint64, entropy []byte) (*VRFOut, error) {

	// draw randomness
	randomness, err := drawRandomness(rbase, pers, height, entropy)
	if err != nil {
		return nil, xerrors.Errorf("GenerateVRF drawRandomness failed: %w", err)
	}

	// compute vrf
	vrf, err := blsSign(privateKey, randomness)
	if err != nil {
		return nil, xerrors.Errorf("GenerateVRF blsSign failed: %w", err)
	}

	return &VRFOut{
		Height: height,
		Proof:  vrf,
	}, nil
}

func FilVerifyVRFByTipSet(pers DomainSeparationTag, worker address.Address,
	ts *filrpc.TipSet, entropy []byte, vrf *VRFOut) error {
	if ts.Height() != vrf.Height {
		return xerrors.Errorf("FilVerifyVRFByTipSet tipset height %d != %d(vrf)", ts.Height(), vrf.Height)
	}

	if len(ts.Blocks()) == 0 {
		return xerrors.Errorf("FilVerifyVRFByTipSet no block in tipset(height:%d)", ts.Height())
	}

	// use min ticket
	minTicket := ts.MinTicket()
	return VerifyVRF(worker.Payload(), pers, minTicket.VRFProof, entropy, vrf)
}

func FilGenerateVRFByTipSet(pers DomainSeparationTag,
	privateKey []byte, ts *filrpc.TipSet, entropy []byte) (*VRFOut, error) {
	if len(ts.Blocks()) == 0 {
		return nil, xerrors.Errorf("FilGenerateVRFByTipSet no block in tipset(height:%d)", ts.Height())
	}

	privateKey = FilBlsKey2KyberBlsKey(privateKey)

	// use min ticket
	minTicket := ts.MinTicket()
	return GenerateVRF(pers, privateKey, minTicket.VRFProof, ts.Height(), entropy)
}
