package types_test

import (
	"testing"

	"github.com/qubetics/qubetics-blockchain/v2/x/vesting/types"
	"github.com/stretchr/testify/suite"
)

type ProposalTestSuite struct {
	suite.Suite
}

func TestProposalTestSuite(t *testing.T) {
	suite.Run(t, new(ProposalTestSuite))
}

func (suite *ProposalTestSuite) TestKeysTypes() {
	suite.Require().Equal("vesting", (&types.ClawbackProposal{}).ProposalRoute())
	suite.Require().Equal("Clawback", (&types.ClawbackProposal{}).ProposalType())
}

func (suite *ProposalTestSuite) TestClawbackProposal() {
	// Valid address generated with valid checksum
	validAddr := "qubetics1l3vfmfnlhmg0mylvf2kv9t9eap40fyqyvz235u"

	testCases := []struct {
		msg                string
		title              string
		description        string
		address            string
		destinationAddress string
		expectPass         bool
	}{
		{
			msg:         "Clawback proposal - valid address",
			title:       "test",
			description: "test desc",
			address:     validAddr,
			expectPass:  true,
		},
		{
			msg:         "Clawback proposal - invalid missing title",
			title:       "",
			description: "test desc",
			address:     validAddr,
			expectPass:  false,
		},
		{
			msg:         "Clawback proposal - invalid missing description",
			title:       "test",
			description: "",
			address:     validAddr,
			expectPass:  false,
		},
		{
			msg:         "Clawback proposal - invalid address (bad format)",
			title:       "test",
			description: "test desc",
			address:     "qubetics1invalidaddressxxxxxxxxxxxxxxxxxxxxxxxxxx",
			expectPass:  false,
		},
		{
			msg:                "Clawback proposal - invalid destination addr",
			title:              "test",
			description:        "test desc",
			address:            validAddr,
			destinationAddress: "125182ujaisch8hsgs",
			expectPass:         false,
		},
	}

	for i, tc := range testCases {
		tx := types.NewClawbackProposal(tc.title, tc.description, tc.address, tc.destinationAddress)
		err := tx.ValidateBasic()

		if tc.expectPass {
			suite.Require().NoError(err, "valid test %d failed: %s", i, tc.msg)
		} else {
			suite.Require().Error(err, "invalid test %d passed: %s", i, tc.msg)
		}
	}
}
