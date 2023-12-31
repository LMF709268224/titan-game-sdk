package gamevrf

import (
	"encoding/hex"
	"encoding/json"

	bls "github.com/drand/kyber-bls12381"
	sign "github.com/drand/kyber/sign/bls"
	rand2 "github.com/drand/kyber/util/random"
	"golang.org/x/xerrors"
)

// Mark: filecoin blst rust implementation use follow DST
const DST = string("BLS_SIG_BLS12381G2_XMD:SHA-256_SSWU_RO_NUL_")

type KeyInfo struct {
	Type       string
	PrivateKey []byte
}

func KyberBlsGenPrivateKey() ([]byte, []byte, error) {
	suite := bls.NewBLS12381Suite()
	scheme := sign.NewSchemeOnG2(suite)
	priv, pub := scheme.NewKeyPair(rand2.New())
	privb, err := priv.MarshalBinary()
	if err != nil {
		return nil, nil, xerrors.Errorf("KyberBlsGenPrivateKey MarshalBinary failed: %w", err)
	}

	pubb, err := pub.MarshalBinary()
	if err != nil {
		return nil, nil, xerrors.Errorf("KyberBlsGenPrivateKey MarshalBinary failed: %w", err)
	}

	return privb, pubb, nil
}

func FilBlsKey2KyberBlsKey(filKey []byte) []byte {
	kyberKey := make([]byte, len(filKey))
	reverse(kyberKey, filKey)

	return kyberKey
}

func KyberBlsKey2FilBlsKey(kyberKey []byte) []byte {
	filKey := make([]byte, len(kyberKey))
	reverse(filKey, kyberKey)

	return filKey
}

func FilBlsKey2PublicKey(filKey []byte) ([]byte, error) {
	privateKey := FilBlsKey2KyberBlsKey(filKey)

	suite := bls.NewBLS12381Suite()
	sc := suite.G1().Scalar()
	err := sc.UnmarshalBinary(privateKey)
	if err != nil {
		return nil, xerrors.Errorf("FilBlsKey2PublicKey UnmarshalBinary failed: %w", err)
	}

	pub := suite.G1().Point().Mul(sc, nil)
	pubb, err := pub.MarshalBinary()
	if err != nil {
		return nil, xerrors.Errorf("FilBlsKey2PublicKey MarshalBinary failed: %w", err)
	}

	return pubb, nil
}

// the privateKey is export from filecoin wallet
func FilBlsKeyFromString(privateKey string) ([]byte, error) {
	priv, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, err
	}

	var keyInfo KeyInfo
	err = json.Unmarshal(priv, &keyInfo)
	if err != nil {
		return nil, err
	}

	return keyInfo.PrivateKey, nil
}

func reverse(dst, src []byte) []byte {
	if dst == nil {
		dst = make([]byte, len(src))
	}
	l := len(dst)
	for i, j := 0, l-1; i < (l+1)/2; {
		dst[i], dst[j] = src[j], src[i]
		i++
		j--
	}
	return dst
}

func blsVerify(pubKey []byte, vrfBase, vrfproof []byte) error {
	suite := bls.NewBLS12381Suite()
	scheme := sign.NewSchemeOnG2(suite)
	sp := suite.G1().Point()
	err := sp.UnmarshalBinary(pubKey)
	if err != nil {
		return xerrors.Errorf("blsVerify UnmarshalBinary failed: %w", err)
	}

	err = scheme.Verify(sp, vrfBase, vrfproof)

	return err
}

func blsSign(privateKey, sigInput []byte) ([]byte, error) {
	suite := bls.NewBLS12381Suite()
	scheme := sign.NewSchemeOnG2(suite)
	sc := suite.G1().Scalar()
	err := sc.UnmarshalBinary(privateKey)
	if err != nil {
		return nil, xerrors.Errorf("blsSign UnmarshalBinary failed: %w", err)
	}

	sig, err := scheme.Sign(sc, sigInput)
	if err != nil {
		return nil, xerrors.Errorf("blsSign Sign failed: %w", err)
	}

	return sig, nil
}
