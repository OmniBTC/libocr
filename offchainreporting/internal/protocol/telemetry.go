package protocol

import (
	"github.com/OmniBTC/libocr/commontypes"
	"github.com/OmniBTC/libocr/offchainreporting/types"
)

type TelemetrySender interface {
	RoundStarted(
		configDigest types.ConfigDigest,
		epoch uint32,
		round uint8,
		leader commontypes.OracleID,
	)
}
