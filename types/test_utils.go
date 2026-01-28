package types

import (
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	TestTimeZero = time.Time{}
	TestTimeNow  = time.Now()

	TestAddrEmpty         = ""
	TestAddrInvalid       = "invalid"
	TestAddrInvalidPrefix = "invalid1qypqxpq9qcrsszgszyfpx9q4zct3sxfqe52gp4"

	TestBech32AccAddr10Bytes = "qubetics1qypqxpq9qcrsszgslawd5s"
	TestBech32AccAddr20Bytes = "qubetics1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj"
	TestBech32AccAddr30Bytes = "qubetics1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fszvfck8"

	TestBech32NodeAddr10Bytes = "qubeticsnode1qypqxpq9qcrsszgse4wwrm"
	TestBech32NodeAddr20Bytes = "qubeticsnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey"
	TestBech32NodeAddr30Bytes = "qubeticsnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsxqglcv"

	TestBech32ProvAddr10Bytes = "qubeticsprov1qypqxpq9qcrsszgsutj8xr"
	TestBech32ProvAddr20Bytes = "qubeticsprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82"
	TestBech32ProvAddr30Bytes = "qubeticsprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsh33zgx"

	TestDenomEmpty   = ""
	TestDenomInvalid = "i"
	TestDenomOne     = "one"
	TestDenomTwo     = "two"

	TestIntEmpty = sdkmath.Int{}
	TestIntNeg   = sdkmath.NewInt(-1000)
	TestIntZero  = sdkmath.NewInt(0)
	TestIntPos   = sdkmath.NewInt(1000)

	TestCoinEmpty      = sdk.Coin{}
	TestCoinEmptyZero  = sdk.Coin{Amount: TestIntZero, Denom: TestDenomEmpty}
	TestCoinEmptyPos   = sdk.Coin{Amount: TestIntPos, Denom: TestDenomEmpty}
	TestCoinInvalidPos = sdk.Coin{Amount: TestIntPos, Denom: TestDenomInvalid}
	TestCoinOneEmpty   = sdk.Coin{Amount: TestIntEmpty, Denom: TestDenomOne}
	TestCoinOneNeg     = sdk.Coin{Amount: TestIntNeg, Denom: TestDenomOne}
	TestCoinOneZero    = sdk.Coin{Amount: TestIntZero, Denom: TestDenomOne}
	TestCoinOnePos     = sdk.Coin{Amount: TestIntPos, Denom: TestDenomOne}

	TestCoinsNil        sdk.Coins = nil
	TestCoinsEmpty                = sdk.Coins{}
	TestCoinsEmptyPos             = sdk.Coins{TestCoinEmptyPos}
	TestCoinsInvalidPos           = sdk.Coins{TestCoinInvalidPos}
	TestCoinsOneEmpty             = sdk.Coins{TestCoinOneEmpty}
	TestCoinsOneNeg               = sdk.Coins{TestCoinOneNeg}
	TestCoinsOneZero              = sdk.Coins{TestCoinOneZero}
	TestCoinsOnePos               = sdk.Coins{TestCoinOnePos}
)
