package validator

import (
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"time"

	secp256k1Curve "github.com/ethereum/go-ethereum/crypto/secp256k1"
	torusCommon "github.com/torusresearch/torus-common/common"
	"github.com/torusresearch/torus-common/crypto"
	"github.com/torusresearch/torus-common/secp256k1"
)

type CommonSignatureData struct {
	PubKeyX      string `json:"pub_key_x"`
	PubKeyY      string `json:"pub_key_y"`
	TempKeyX     string `json:"temp_key_x"`
	TempKeyY     string `json:"temp_key_y"`
	VerifierName string `json:"verifier_name"`
	VerifierID   string `json:"verifier_id"`
}

type Signature struct {
	Sig        string                 `json:"sig"`
	Data       string                 `json:"data"`
	ParsedData map[string]interface{} `json:"parsedData"`
}

type SessionValidator struct {
	NodePubKeysXArray []string
	NodePubKeysYArray []string
}

type ValidationOptions struct {
	SigBase64Encoded  bool
	SkipExpValidation bool
}

type VerifyResult struct {
	IsValid bool
}

func (sv *SessionValidator) Authenticate(signatures []string, options *ValidationOptions) (*CommonSignatureData, error) {
	if options == nil {
		options = &ValidationOptions{}
	}
	verifiedData, err := sv.validateSignatures(signatures, options)
	if err != nil {
		return nil, err
	}

	return deserializeSignatureData(verifiedData)
}

func (sv *SessionValidator) validateSignatures(signatures []string, options *ValidationOptions) (string, error) {
	if len(signatures) == 0 {
		return "", errors.New("invalid signature count: 0")
	}

	if len(sv.NodePubKeysXArray) != len(sv.NodePubKeysYArray) {
		return "", errors.New("invalid node public key count")
	}

	if len(signatures) > len(sv.NodePubKeysXArray) {
		return "", errors.New("invalid signature count: more than number of nodes")
	}

	parsedSignatures := decodeSigData(signatures)

	threshold := (len(sv.NodePubKeysXArray) / 2) + 1
	data := sv.validateThreshold(parsedSignatures, threshold, options)
	if data == "" {
		return "", errors.New("insufficient number of valid signatures")
	}

	return data, nil
}

func (sv *SessionValidator) validateThreshold(signatures []Signature, thresholdCount int, options *ValidationOptions) string {
	validSigData := make([]string, 0)
	dataMap := make(map[string]int)

	usedKeys := make(map[int]bool)

	for _, sig := range signatures {
		result := sv.verifySignature(sig.Sig, sig.Data, usedKeys, options)

		if !result.IsValid {
			continue
		}

		expValue, ok := sig.ParsedData["exp"]
		var expTime int64
		if ok {
			switch v := expValue.(type) {
			case float64:
				expTime = int64(v)
			case int64:
				expTime = v
			default:
				ok = false
			}
			sigNotExpired := options.SkipExpValidation || (ok && expTime > time.Now().Unix())
			if !sigNotExpired {
				continue
			}

			serializedData, err := serializeSignatureData(sig.ParsedData)
			if err != nil {
				continue
			}
			dataMap[serializedData]++
			validSigData = append(validSigData, sig.Data)

			if len(validSigData) >= thresholdCount {
				for k, v := range dataMap {
					if v >= thresholdCount {
						return k
					}
				}
			}
		}
	}
	return ""
}

func (sv *SessionValidator) verifySignature(signature, message string, usedKeys map[int]bool, options *ValidationOptions) VerifyResult {
	result := VerifyResult{IsValid: false}

	if len(sv.NodePubKeysXArray) != len(sv.NodePubKeysYArray) {
		return result
	}

	for i := 0; i < len(sv.NodePubKeysXArray); i++ {
		if usedKeys[i] {
			continue
		}

		var hash32 [32]byte
		copy(hash32[:], secp256k1.Keccak256([]byte(message))[:32])

		// Note: The actual EC key verification would need to be implemented
		// using a suitable elliptic curve library (e.g., secp256k1)
		// This is a placeholder for the actual verification logic
		pubKey := createPublicKey(sv.NodePubKeysXArray[i], sv.NodePubKeysYArray[i])

		sig := signature
		if options.SigBase64Encoded {
			decoded, err := base64.StdEncoding.DecodeString(signature)
			if err != nil {
				continue
			}
			sig = string(decoded)
		}

		// Placeholder for actual signature verification
		isValid := crypto.VerifySignature(pubKey, hash32, sig)

		if isValid {
			result.IsValid = true
			// Remove the used public key
			usedKeys[i] = true
			break
		}
	}

	return result
}

// Helper functions that would need to be implemented
func serializeSignatureData(data map[string]interface{}) (string, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func deserializeSignatureData(data string) (*CommonSignatureData, error) {
	var parsedData CommonSignatureData
	if err := json.Unmarshal([]byte(data), &parsedData); err != nil {
		return nil, err
	}
	return &parsedData, nil
}

func decodeSigData(signatures []string) []Signature {
	var parsedSignatures []Signature

	for _, s := range signatures {
		if s == "" {
			continue
		}

		var sig Signature
		if err := json.Unmarshal([]byte(s), &sig); err != nil {
			continue
		}

		decodedData, err := base64.StdEncoding.DecodeString(sig.Data)
		if err != nil {
			continue
		}

		var parsedData map[string]interface{}
		if err := json.Unmarshal(decodedData, &parsedData); err != nil {
			continue
		}

		sig.ParsedData = parsedData
		parsedSignatures = append(parsedSignatures, sig)
	}

	return parsedSignatures
}

func createPublicKey(x, y string) ecdsa.PublicKey {
	return ecdsa.PublicKey{
		Curve: secp256k1Curve.S256(),
		X:     torusCommon.HexToBigInt(x),
		Y:     torusCommon.HexToBigInt(y),
	}
}
