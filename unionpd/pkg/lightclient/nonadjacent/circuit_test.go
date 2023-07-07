package nonadjacent

import (
	"encoding/base64"
	"encoding/hex"
	"log"
	"math/big"
	"unionp/pkg/lightclient"

	cometbn254 "github.com/cometbft/cometbft/crypto/bn254"
	"github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/consensys/gnark/frontend"
	sdk "github.com/cosmos/cosmos-sdk/types"

	cometbft_bn254 "github.com/cometbft/cometbft/crypto/bn254"
	ce "github.com/cometbft/cometbft/crypto/encoding"
	"github.com/cometbft/cometbft/crypto/merkle"
	"github.com/cometbft/cometbft/libs/protoio"
	"github.com/consensys/gnark-crypto/ecc"
	curve "github.com/consensys/gnark-crypto/ecc/bn254"
	back "github.com/consensys/gnark/backend"

	gadget "github.com/consensys/gnark/std/algebra/emulated/sw_bn254"

	"testing"

	"github.com/consensys/gnark/test"
)

func TestNonAdjacent(t *testing.T) {
	// cSpell:disable
	/*
		{
		  "block_id": {
		    "hash": "B105399C580D97B7644231473EE35852E033AD60F674BDF2F90B4A85C8D9A6D0",
		    "parts": {
		      "total": 1,
		      "hash": "88E1090871BE13B2A576AE63425AF99D46382F3E9AE6ECA4CB39B08B78AC4E52"
		    }
		  },
		  "block": {
		    "header": {
		      "version": {
		        "block": "11"
		      },
		      "chain_id": "union-devnet-1",
		      "height": "575",
		      "time": "2023-06-17T11:11:53.693111971Z",
		      "last_block_id": {
		        "hash": "7022627E60ED78120D2FE8DC7ACDB58A2321B0304F8912D2DFB86CE038E23CA8",
		        "parts": {
		          "total": 1,
		          "hash": "41B8793236EE0980E2EAF1A2FAD268C4A3D8979A0C432F06E284EEC5E74DD69C"
		        }
		      },
		      "last_commit_hash": "A571E99454431CBF4C22377388160BFF7DB20C1BBA9FC4C9CD7C09CFA9BB93A3",
		      "data_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
		      "validators_hash": "F9CE4237F8D95BDF686A7FBE68D66593F612FF2695DC27A72DAEFFDB1557B304",
		      "next_validators_hash": "F9CE4237F8D95BDF686A7FBE68D66593F612FF2695DC27A72DAEFFDB1557B304",
		      "consensus_hash": "048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F",
		      "app_hash": "7CDFF8464A7166DAE92EC9AB4103A978ED284D426C97944FC2AB25D820A2B681",
		      "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
		      "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
		      "proposer_address": "7041D470CFFB5F210003289D3EBC0713B30C279A"
		    },
		    "data": {
		      "txs": null
		    },
		    "evidence": {
		      "evidence": null
		    },
		    "last_commit": {
		      "height": "574",
		      "round": 0,
		      "block_id": {
		        "hash": "7022627E60ED78120D2FE8DC7ACDB58A2321B0304F8912D2DFB86CE038E23CA8",
		        "parts": {
		          "total": 1,
		          "hash": "41B8793236EE0980E2EAF1A2FAD268C4A3D8979A0C432F06E284EEC5E74DD69C"
		        }
		      },
		      "signatures": [
		        {
		          "block_id_flag": 2,
		          "validator_address": "135ECF966844EDDA4DDA4E04402A1FBB6F67CA2F",
		          "timestamp": "2023-06-17T11:11:53.771049219Z",
		          "signature": "ot0fBmXr8EoA0aDoeqlsjl+eKHmBAsAnw4uUSGgtPOcFQ9amhgwGsS3RQ4LsDyvDpHSHY+wJXehDOcelWRlwhw=="
		        },
		        {
		          "block_id_flag": 2,
		          "validator_address": "629ADA7B1B463B0E16C189F0264CB18E8E83C940",
		          "timestamp": "2023-06-17T11:11:53.869030133Z",
		          "signature": "0SU+F1fVo0tIpm8x8ZxSjAFB+T65tV39hbr5VPAskRcQZD68Bjo/ahFn48Qe9oEofxoMQKteZZZ5ZzVN7N1w6Q=="
		        },
		        {
		          "block_id_flag": 2,
		          "validator_address": "7041D470CFFB5F210003289D3EBC0713B30C279A",
		          "timestamp": "2023-06-17T11:11:53.693111971Z",
		          "signature": "3E1nbXPJ7UFsGslA3Asm1WqJvFNkseC3wiARerflVH8RkJEGFqJ1TLVdzB9+8943Ay+PzqHR8AnVDbkGp/ANJw=="
		        },
		        {
		          "block_id_flag": 2,
		          "validator_address": "C6B5C8EAAE0B8DA907AFDCD4BFEA555949FC9E9B",
		          "timestamp": "2023-06-17T11:11:53.677919338Z",
		          "signature": "i2cPkcfsBRAehCVbYHx3vtAaIfD+JXhCniv3aRfiAvgvO0uDC0iA3rQrKm8Qp3ddIZavmB2airwXjCBwHvy4kw=="
		        }
		      ]
		    }
		  }
		}


		{
		  "pagination": {
		    "next_key": null,
		    "total": "0"
		  },
		  "validators": [
		    {
		      "commission": {
		        "commission_rates": {
		          "max_change_rate": "0.010000000000000000",
		          "max_rate": "0.200000000000000000",
		          "rate": "0.100000000000000000"
		        },
		        "update_time": "2023-06-17T10:18:00.424685336Z"
		      },
		      "consensus_pubkey": {
		        "@type": "/cosmos.crypto.bn254.PubKey",
		        "key": "w9svXa/PETX5x8PcovHun9rZSehY89LLP6T9Bhc+5Uk="
		      },
		      "delegator_shares": "1000000000000000000000.000000000000000000",
		      "description": {
		        "details": "",
		        "identity": "",
		        "moniker": "validator-3",
		        "security_contact": "",
		        "website": ""
		      },
		      "jailed": false,
		      "min_self_delegation": "1",
		      "operator_address": "unionvaloper1sw27dhxh32dz4klrllygy24g7tlse7latavr30",
		      "status": "BOND_STATUS_BONDED",
		      "tokens": "1000000000000000000000",
		      "unbonding_height": "0",
		      "unbonding_ids": [],
		      "unbonding_on_hold_ref_count": "0",
		      "unbonding_time": "1970-01-01T00:00:00Z"
		    },
		    {
		      "commission": {
		        "commission_rates": {
		          "max_change_rate": "0.010000000000000000",
		          "max_rate": "0.200000000000000000",
		          "rate": "0.100000000000000000"
		        },
		        "update_time": "2023-06-17T10:18:00.424685336Z"
		      },
		      "consensus_pubkey": {
		        "@type": "/cosmos.crypto.bn254.PubKey",
		        "key": "1BWtpD0d5nXD5mFX24/hQAuYCGNRKs7wAw8KPxmFdQM="
		      },
		      "delegator_shares": "1000000000000000000000.000000000000000000",
		      "description": {
		        "details": "",
		        "identity": "",
		        "moniker": "validator-1",
		        "security_contact": "",
		        "website": ""
		      },
		      "jailed": false,
		      "min_self_delegation": "1",
		      "operator_address": "unionvaloper1s36k93gu7x58zns0g4mrep8qgnr5fm4jrqddkv",
		      "status": "BOND_STATUS_BONDED",
		      "tokens": "1000000000000000000000",
		      "unbonding_height": "0",
		      "unbonding_ids": [],
		      "unbonding_on_hold_ref_count": "0",
		      "unbonding_time": "1970-01-01T00:00:00Z"
		    },
		    {
		      "commission": {
		        "commission_rates": {
		          "max_change_rate": "0.010000000000000000",
		          "max_rate": "0.200000000000000000",
		          "rate": "0.100000000000000000"
		        },
		        "update_time": "2023-06-17T10:18:00.424685336Z"
		      },
		      "consensus_pubkey": {
		        "@type": "/cosmos.crypto.bn254.PubKey",
		        "key": "r88zHnnu0zDYC9KBKpnXP9K5V1brl8UnVG/PjaPqCWE="
		      },
		      "delegator_shares": "1000000000000000000000.000000000000000000",
		      "description": {
		        "details": "",
		        "identity": "",
		        "moniker": "validator-2",
		        "security_contact": "",
		        "website": ""
		      },
		      "jailed": false,
		      "min_self_delegation": "1",
		      "operator_address": "unionvaloper1ndm3ljzqwvc60uvxwkhczgrczc4jv2ll9etcn6",
		      "status": "BOND_STATUS_BONDED",
		      "tokens": "1000000000000000000000",
		      "unbonding_height": "0",
		      "unbonding_ids": [],
		      "unbonding_on_hold_ref_count": "0",
		      "unbonding_time": "1970-01-01T00:00:00Z"
		    },
		    {
		      "commission": {
		        "commission_rates": {
		          "max_change_rate": "0.010000000000000000",
		          "max_rate": "0.200000000000000000",
		          "rate": "0.100000000000000000"
		        },
		        "update_time": "2023-06-17T10:18:00.424685336Z"
		      },
		      "consensus_pubkey": {
		        "@type": "/cosmos.crypto.bn254.PubKey",
		        "key": "nuGoo7R0UOvKTWSo1oHQyUPmHwwzkaFLv7I9ClrmW5Y="
		      },
		      "delegator_shares": "1000000000000000000000.000000000000000000",
		      "description": {
		        "details": "",
		        "identity": "",
		        "moniker": "validator-0",
		        "security_contact": "",
		        "website": ""
		      },
		      "jailed": false,
		      "min_self_delegation": "1",
		      "operator_address": "unionvaloper14fldwd959h7glh2e3k45veuqfszvgm693pv868",
		      "status": "BOND_STATUS_BONDED",
		      "tokens": "1000000000000000000000",
		      "unbonding_height": "0",
		      "unbonding_ids": [],
		      "unbonding_on_hold_ref_count": "0",
		      "unbonding_time": "1970-01-01T00:00:00Z"
		    }
		  ]
		}
	*/
	// cSpell:enable

	decodeB64 := func(s string) []byte {
		bz, err := base64.StdEncoding.DecodeString(s)
		if err != nil {
			log.Fatal(err)
		}
		return bz
	}

	// Nb of tokens for each val in devnet
	tokens, success := new(big.Int).SetString("1000000000000000000000", 10)
	if !success {
		log.Fatal("Impossible; qed;")
	}

	toValidator := func(pubKey []byte) *types.SimpleValidator {
		protoPK, err := ce.PubKeyToProto(cometbn254.PubKey(pubKey))
		if err != nil {
			log.Fatal(err)
		}
		return &types.SimpleValidator{
			PubKey:      &protoPK,
			VotingPower: sdk.TokensToConsensusPower(sdk.NewIntFromBigInt(tokens), sdk.DefaultPowerReduction),
		}
	}

	blockHash, err := hex.DecodeString("7022627E60ED78120D2FE8DC7ACDB58A2321B0304F8912D2DFB86CE038E23CA8")
	if err != nil {
		log.Fatal(err)
	}

	partSetHeaderHash, err := hex.DecodeString("41B8793236EE0980E2EAF1A2FAD268C4A3D8979A0C432F06E284EEC5E74DD69C")
	if err != nil {
		log.Fatal(err)
	}

	vote := types.CanonicalVote{
		Type:   types.PrecommitType,
		Height: 574,
		Round:  0,
		BlockID: &types.CanonicalBlockID{
			Hash: blockHash,
			PartSetHeader: types.CanonicalPartSetHeader{
				Total: 1,
				Hash:  partSetHeaderHash,
			},
		},
		ChainID: "union-devnet-1",
	}

	validators := []*types.SimpleValidator{
		toValidator(decodeB64("nuGoo7R0UOvKTWSo1oHQyUPmHwwzkaFLv7I9ClrmW5Y=")), // cspell:disable-line
		toValidator(decodeB64("1BWtpD0d5nXD5mFX24/hQAuYCGNRKs7wAw8KPxmFdQM=")), // cspell:disable-line
		toValidator(decodeB64("r88zHnnu0zDYC9KBKpnXP9K5V1brl8UnVG/PjaPqCWE=")), // cspell:disable-line
		toValidator(decodeB64("w9svXa/PETX5x8PcovHun9rZSehY89LLP6T9Bhc+5Uk=")), // cspell:disable-line
	}

	trustedValidators := validators
	untrustedValidators := validators

	signatures := [][]byte{
		decodeB64("ot0fBmXr8EoA0aDoeqlsjl+eKHmBAsAnw4uUSGgtPOcFQ9amhgwGsS3RQ4LsDyvDpHSHY+wJXehDOcelWRlwhw=="), // cspell:disable-line
		decodeB64("0SU+F1fVo0tIpm8x8ZxSjAFB+T65tV39hbr5VPAskRcQZD68Bjo/ahFn48Qe9oEofxoMQKteZZZ5ZzVN7N1w6Q=="), // cspell:disable-line
		decodeB64("3E1nbXPJ7UFsGslA3Asm1WqJvFNkseC3wiARerflVH8RkJEGFqJ1TLVdzB9+8943Ay+PzqHR8AnVDbkGp/ANJw=="), // cspell:disable-line
		decodeB64("i2cPkcfsBRAehCVbYHx3vtAaIfD+JXhCniv3aRfiAvgvO0uDC0iA3rQrKm8Qp3ddIZavmB2airwXjCBwHvy4kw=="), // cspell:disable-line
	}

	trustedSignatures := signatures
	untrustedSignatures := signatures

	var bitmap big.Int
	bitmap.SetBit(&bitmap, 0, 1)
	bitmap.SetBit(&bitmap, 1, 1)
	bitmap.SetBit(&bitmap, 2, 1)
	bitmap.SetBit(&bitmap, 3, 1)

	trustedBitmap := bitmap
	untrustedBitmap := bitmap

	if err != nil {
		log.Fatal(err)
	}

	reverseBytes := func(numbers []byte) []byte {
		newNumbers := make([]byte, 0, len(numbers))
		for i := len(numbers) - 1; i >= 0; i-- {
			newNumbers = append(newNumbers, numbers[i])
		}
		return newNumbers
	}

	marshalValidators := func(validators []*types.SimpleValidator) ([lightclient.MaxVal][4]frontend.Variable, []byte, error) {
		validatorsProto := [lightclient.MaxVal][4]frontend.Variable{}
		// Make sure we zero initialize
		for i := 0; i < lightclient.MaxVal; i++ {
			validatorsProto[i][0] = 0
			validatorsProto[i][1] = 0
			validatorsProto[i][2] = 0
			validatorsProto[i][3] = 0
		}
		merkleTree := make([][]byte, len(validators))
		for i, val := range validators {
			protoEncoding, err := val.Marshal()
			if err != nil {
				return validatorsProto, nil, err
			}

			merkleTree[i] = protoEncoding

			tmPK, err := ce.PubKeyFromProto(*val.PubKey)
			if err != nil {
				return validatorsProto, nil, err
			}

			compressedPK := tmPK.Bytes()

			var PK curve.G1Affine
			_, err = PK.SetBytes(compressedPK)
			if err != nil {
				return validatorsProto, nil, err
			}

			PKX := PK.X.Bytes()
			PKY := PK.Y.Bytes()
			// Need to reverse to simplify circuit computation
			power := reverseBytes(protoEncoding[lightclient.ValProtoPower:])
			mask := compressedPK[0] >> 6
			validatorsProto[i][0] = PKX[:]
			validatorsProto[i][1] = PKY[:]
			validatorsProto[i][2] = power
			validatorsProto[i][3] = mask
		}
		return validatorsProto, merkle.HashFromByteSlices(merkleTree), nil
	}

	aggregateSignatures := func(signatures [][]byte) (curve.G2Affine, error) {
		var aggregatedSignature curve.G2Affine
		var decompressedSignature curve.G2Affine
		for _, signature := range signatures {
			_, err := decompressedSignature.SetBytes(signature)
			if err != nil {
				return curve.G2Affine{}, err
			}
			aggregatedSignature.Add(&aggregatedSignature, &decompressedSignature)
		}
		return aggregatedSignature, nil
	}

	trustedValidatorsProto, trustedValidatorsRoot, err := marshalValidators(trustedValidators)
	if err != nil {
		log.Fatal(err)
	}
	trustedAggregatedSignature, err := aggregateSignatures(trustedSignatures)
	if err != nil {
		log.Fatal(err)
	}

	untrustedValidatorsProto, untrustedValidatorsRoot, err := marshalValidators(untrustedValidators)
	if err != nil {
		log.Fatal(err)
	}
	untrustedAggregatedSignature, err := aggregateSignatures(untrustedSignatures)
	if err != nil {
		log.Fatal(err)
	}

	trustedInput := TendermintNonAdjacentLightClientInput{
		Sig:             gadget.NewG2Affine(trustedAggregatedSignature),
		ProtoValidators: trustedValidatorsProto,
		NbOfVal:         len(trustedValidators),
		NbOfSignature:   len(trustedSignatures),
		Bitmap:          trustedBitmap,
	}

	untrustedInput := TendermintNonAdjacentLightClientInput{
		Sig:             gadget.NewG2Affine(untrustedAggregatedSignature),
		ProtoValidators: untrustedValidatorsProto,
		NbOfVal:         len(untrustedValidators),
		NbOfSignature:   len(untrustedSignatures),
		Bitmap:          untrustedBitmap,
	}

	signedBytes, err := protoio.MarshalDelimited(&vote)
	if err != nil {
		log.Fatal(err)
	}

	hmX, hmY := cometbft_bn254.HashToField2(signedBytes)

	witness := Circuit{
		TrustedInput:   trustedInput,
		UntrustedInput: untrustedInput,
		ExpectedTrustedValRoot: [2]frontend.Variable{
			trustedValidatorsRoot[0:16],
			trustedValidatorsRoot[16:32],
		},
		ExpectedUntrustedValRoot: [2]frontend.Variable{
			untrustedValidatorsRoot[0:16],
			untrustedValidatorsRoot[16:32],
		},
		Message: [2]frontend.Variable{hmX, hmY},
	}

	assert := test.NewAssert(t)

	var circuit Circuit

	assert.ProverSucceeded(&circuit, &witness, test.WithCurves(ecc.BN254), test.WithBackends(back.GROTH16), test.NoSerialization(), test.NoFuzzing())
}