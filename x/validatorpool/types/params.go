package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var KeyAnnualRoi = []byte("AnnualRoi")

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(annualRoi uint32) Params {
	return Params{
		AnnualRoi: annualRoi,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(30) // 30% in basis points
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyAnnualRoi, &p.AnnualRoi, validateAnnualRoi),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateAnnualRoi(p.AnnualRoi); err != nil {
		return err
	}
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validateAnnualRoi(i interface{}) error {
	_, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	// if v == 0 {
	// 	return fmt.Errorf("annual ROI cannot be zero")
	// }
	return nil
}
