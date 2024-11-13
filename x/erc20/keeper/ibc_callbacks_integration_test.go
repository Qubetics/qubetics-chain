package keeper_test

import (
	"math/big"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/qubetics/qubetics-blockchain/v2/contracts"
	teststypes "github.com/qubetics/qubetics-blockchain/v2/types/tests"
	"github.com/qubetics/qubetics-blockchain/v2/utils"
	"github.com/qubetics/qubetics-blockchain/v2/x/erc20/types"

	//nolint:revive // dot imports are fine for Ginkgo
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Convert native ERC20 receiving from IBC back to Erc20", Ordered, func() {
	var (
		sender, receiver string
		receiverAcc      sdk.AccAddress
		senderAcc        sdk.AccAddress
		amount           int64 = 10
		pair             *types.TokenPair
		erc20Denomtrace  transfertypes.DenomTrace
	)

	BeforeEach(func() {
		s.suiteIBCTesting = true
		s.SetupTest()
		s.suiteIBCTesting = false
	})

	Describe("registered native erc20", func() {
		BeforeEach(func() {
			receiver = s.IBCOsmosisChain.SenderAccount.GetAddress().String()
			sender = s.QubeticsChain.SenderAccount.GetAddress().String()
			receiverAcc = sdk.MustAccAddressFromBech32(receiver)
			senderAcc = sdk.MustAccAddressFromBech32(sender)

			// Register ERC20 pair
			addr, err := s.DeployContractToChain("testcoin", "tt", 18)
			s.Require().NoError(err)
			pair, err = s.app.Erc20Keeper.RegisterERC20(s.QubeticsChain.GetContext(), addr)
			s.Require().NoError(err)

			erc20Denomtrace = transfertypes.DenomTrace{
				Path:      "transfer/channel-0",
				BaseDenom: pair.Denom,
			}

			s.QubeticsChain.SenderAccount.SetSequence(s.QubeticsChain.SenderAccount.GetSequence() + 1) //nolint:errcheck
		})
		It("should convert erc20 ibc voucher to original erc20", func() {
			// Mint tokens and send to receiver
			_, err := s.app.EvmKeeper.CallEVM(s.QubeticsChain.GetContext(), contracts.ERC20MinterBurnerDecimalsContract.ABI, common.BytesToAddress(senderAcc.Bytes()), pair.GetERC20Contract(), true, "mint", common.BytesToAddress(senderAcc.Bytes()), big.NewInt(amount))
			s.Require().NoError(err)
			// Check Balance
			balanceToken := s.app.Erc20Keeper.BalanceOf(s.QubeticsChain.GetContext(), contracts.ERC20MinterBurnerDecimalsContract.ABI, pair.GetERC20Contract(), common.BytesToAddress(senderAcc.Bytes()))
			s.Require().Equal(amount, balanceToken.Int64())

			// Convert half of the available tokens
			msgConvertERC20 := types.NewMsgConvertERC20(
				math.NewInt(amount),
				senderAcc,
				pair.GetERC20Contract(),
				common.BytesToAddress(senderAcc.Bytes()),
			)

			err = msgConvertERC20.ValidateBasic()
			s.Require().NoError(err)
			// Use MsgConvertERC20 to convert the ERC20 to a Cosmos IBC Coin
			_, err = s.app.Erc20Keeper.ConvertERC20(sdk.WrapSDKContext(s.QubeticsChain.GetContext()), msgConvertERC20)
			s.Require().NoError(err)

			// Check Balance
			balanceToken = s.app.Erc20Keeper.BalanceOf(s.QubeticsChain.GetContext(), contracts.ERC20MinterBurnerDecimalsContract.ABI, pair.GetERC20Contract(), common.BytesToAddress(senderAcc.Bytes()))
			s.Require().Equal(int64(0), balanceToken.Int64())

			// IBC coin balance should be amount
			erc20CoinsBalance := s.app.BankKeeper.GetBalance(s.QubeticsChain.GetContext(), senderAcc, pair.Denom)
			s.Require().Equal(amount, erc20CoinsBalance.Amount.Int64())

			s.QubeticsChain.Coordinator.CommitBlock()

			// Attempt to send erc20 into ibc, should send without conversion
			s.SendBackCoins(s.pathOsmosisQubetics, s.QubeticsChain, pair.Denom, amount, sender, receiver, 1, pair.Denom)
			s.IBCOsmosisChain.Coordinator.CommitBlock()

			// Check balance on the Osmosis chain
			erc20IBCBalance := s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), receiverAcc, erc20Denomtrace.IBCDenom())
			s.Require().Equal(amount, erc20IBCBalance.Amount.Int64())

			s.SendAndReceiveMessage(s.pathOsmosisQubetics, s.IBCOsmosisChain, erc20Denomtrace.IBCDenom(), amount, receiver, sender, 1, erc20Denomtrace.GetFullDenomPath())
			// Check Balance
			balanceToken = s.app.Erc20Keeper.BalanceOf(s.QubeticsChain.GetContext(), contracts.ERC20MinterBurnerDecimalsContract.ABI, pair.GetERC20Contract(), common.BytesToAddress(senderAcc.Bytes()))
			s.Require().Equal(amount, balanceToken.Int64())
		})

		It("should convert full available balance of erc20 coin to original erc20 token", func() {
			// Mint tokens and send to receiver
			_, err := s.app.EvmKeeper.CallEVM(s.QubeticsChain.GetContext(), contracts.ERC20MinterBurnerDecimalsContract.ABI, common.BytesToAddress(senderAcc.Bytes()), pair.GetERC20Contract(), true, "mint", common.BytesToAddress(senderAcc.Bytes()), big.NewInt(amount))
			s.Require().NoError(err)
			// Check Balance
			balanceToken := s.app.Erc20Keeper.BalanceOf(s.QubeticsChain.GetContext(), contracts.ERC20MinterBurnerDecimalsContract.ABI, pair.GetERC20Contract(), common.BytesToAddress(senderAcc.Bytes()))
			s.Require().Equal(amount, balanceToken.Int64())

			// Convert half of the available tokens
			msgConvertERC20 := types.NewMsgConvertERC20(
				math.NewInt(amount),
				senderAcc,
				pair.GetERC20Contract(),
				common.BytesToAddress(senderAcc.Bytes()),
			)

			err = msgConvertERC20.ValidateBasic()
			s.Require().NoError(err)
			// Use MsgConvertERC20 to convert the ERC20 to a Cosmos IBC Coin
			_, err = s.app.Erc20Keeper.ConvertERC20(sdk.WrapSDKContext(s.QubeticsChain.GetContext()), msgConvertERC20)
			s.Require().NoError(err)

			// Check Balance
			balanceToken = s.app.Erc20Keeper.BalanceOf(s.QubeticsChain.GetContext(), contracts.ERC20MinterBurnerDecimalsContract.ABI, pair.GetERC20Contract(), common.BytesToAddress(senderAcc.Bytes()))
			s.Require().Equal(int64(0), balanceToken.Int64())

			// erc20 coin balance should be amount
			erc20CoinsBalance := s.app.BankKeeper.GetBalance(s.QubeticsChain.GetContext(), senderAcc, pair.Denom)
			s.Require().Equal(amount, erc20CoinsBalance.Amount.Int64())

			s.QubeticsChain.Coordinator.CommitBlock()

			// Attempt to send erc20 into ibc, should send without conversion
			s.SendBackCoins(s.pathOsmosisQubetics, s.QubeticsChain, pair.Denom, amount/2, sender, receiver, 1, pair.Denom)
			s.IBCOsmosisChain.Coordinator.CommitBlock()

			// Check balance on the Osmosis chain
			erc20IBCBalance := s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), receiverAcc, erc20Denomtrace.IBCDenom())
			s.Require().Equal(amount/2, erc20IBCBalance.Amount.Int64())

			s.SendAndReceiveMessage(s.pathOsmosisQubetics, s.IBCOsmosisChain, erc20Denomtrace.IBCDenom(), amount/2, receiver, sender, 1, erc20Denomtrace.GetFullDenomPath())
			// Check Balance
			balanceToken = s.app.Erc20Keeper.BalanceOf(s.QubeticsChain.GetContext(), contracts.ERC20MinterBurnerDecimalsContract.ABI, pair.GetERC20Contract(), common.BytesToAddress(senderAcc.Bytes()))
			s.Require().Equal(amount, balanceToken.Int64())

			// IBC coin balance should be zero
			erc20CoinsBalance = s.app.BankKeeper.GetBalance(s.QubeticsChain.GetContext(), senderAcc, pair.Denom)
			s.Require().Equal(int64(0), erc20CoinsBalance.Amount.Int64())
		})

		It("send native ERC-20 to osmosis, when sending back IBC coins should convert full balance back to erc20 token", func() {
			// Mint tokens and send to receiver
			_, err := s.app.EvmKeeper.CallEVM(s.QubeticsChain.GetContext(), contracts.ERC20MinterBurnerDecimalsContract.ABI, common.BytesToAddress(senderAcc.Bytes()), pair.GetERC20Contract(), true, "mint", common.BytesToAddress(senderAcc.Bytes()), big.NewInt(amount))
			s.Require().NoError(err)
			// Check Balance
			balanceToken := s.app.Erc20Keeper.BalanceOf(s.QubeticsChain.GetContext(), contracts.ERC20MinterBurnerDecimalsContract.ABI, pair.GetERC20Contract(), common.BytesToAddress(senderAcc.Bytes()))
			s.Require().Equal(amount, balanceToken.Int64())

			s.QubeticsChain.Coordinator.CommitBlock()

			// Attempt to send 1/2 of erc20 balance via ibc, should convert erc20 tokens to ibc coins and send the converted balance via IBC
			s.SendBackCoins(s.pathOsmosisQubetics, s.QubeticsChain, types.ModuleName+"/"+pair.GetERC20Contract().String(), amount/2, sender, receiver, 1, "")
			s.IBCOsmosisChain.Coordinator.CommitBlock()

			// IBC coin balance should be zero
			erc20CoinsBalance := s.app.BankKeeper.GetBalance(s.QubeticsChain.GetContext(), senderAcc, pair.Denom)
			s.Require().Equal(int64(0), erc20CoinsBalance.Amount.Int64())

			// Check updated token Balance
			balanceToken = s.app.Erc20Keeper.BalanceOf(s.QubeticsChain.GetContext(), contracts.ERC20MinterBurnerDecimalsContract.ABI, pair.GetERC20Contract(), common.BytesToAddress(senderAcc.Bytes()))
			s.Require().Equal(amount/2, balanceToken.Int64())

			// Check balance on the Osmosis chain
			erc20IBCBalance := s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), receiverAcc, erc20Denomtrace.IBCDenom())
			s.Require().Equal(amount/2, erc20IBCBalance.Amount.Int64())

			// send back the IBC coins from Osmosis to Qubetics
			s.SendAndReceiveMessage(s.pathOsmosisQubetics, s.IBCOsmosisChain, erc20Denomtrace.IBCDenom(), amount/2, receiver, sender, 1, erc20Denomtrace.GetFullDenomPath())
			// Check Balance
			balanceToken = s.app.Erc20Keeper.BalanceOf(s.QubeticsChain.GetContext(), contracts.ERC20MinterBurnerDecimalsContract.ABI, pair.GetERC20Contract(), common.BytesToAddress(senderAcc.Bytes()))
			s.Require().Equal(amount, balanceToken.Int64())

			// IBC coin balance should be zero
			erc20CoinsBalance = s.app.BankKeeper.GetBalance(s.QubeticsChain.GetContext(), senderAcc, pair.Denom)
			s.Require().Equal(int64(0), erc20CoinsBalance.Amount.Int64())
		})
	})
})

var _ = Describe("Native coins from IBC", Ordered, func() {
	// amount to be transferred
	var amount int64 = 10

	BeforeEach(func() {
		s.suiteIBCTesting = true
		s.SetupTest()
		s.suiteIBCTesting = false
	})
	It("Is native from source chain - should transfer and register pair and deploy a precompile", func() {
		osmosisAddress := s.IBCOsmosisChain.SenderAccount.GetAddress().String()
		qubeticsAddress := s.QubeticsChain.SenderAccount.GetAddress().String()
		qubeticsAccount := sdk.MustAccAddressFromBech32(qubeticsAddress)

		// Precompile should not be available before IBC
		uosmoContractAddr, err := utils.GetIBCDenomAddress(teststypes.UosmoIbcdenom)
		s.Require().NoError(err)

		params := s.app.EvmKeeper.GetParams(s.QubeticsChain.GetContext())
		s.Require().False(s.app.EvmKeeper.IsAvailableStaticPrecompile(&params, uosmoContractAddr))
		// Check receiver's balance for IBC before transfer. Should be zero
		ibcOsmoBalanceBefore := s.app.BankKeeper.GetBalance(s.QubeticsChain.GetContext(), qubeticsAccount, teststypes.UosmoIbcdenom)
		s.Require().Equal(int64(0), ibcOsmoBalanceBefore.Amount.Int64())
		s.QubeticsChain.Coordinator.CommitBlock()

		// Send uosmo from osmosis to qubetics
		s.SendAndReceiveMessage(s.pathOsmosisQubetics, s.IBCOsmosisChain, "uosmo", amount, osmosisAddress, qubeticsAddress, 1, "")
		s.QubeticsChain.Coordinator.CommitBlock()

		// Check IBC uosmo coin balance - should be equals to amount sended
		ibcOsmoBalanceAfter := s.app.BankKeeper.GetBalance(s.QubeticsChain.GetContext(), qubeticsAccount, teststypes.UosmoIbcdenom)
		s.Require().Equal(amount, ibcOsmoBalanceAfter.Amount.Int64())

		// Pair should be registered now and precompile available
		pairID := s.app.Erc20Keeper.GetTokenPairID(s.QubeticsChain.GetContext(), teststypes.UosmoIbcdenom)
		_, found := s.app.Erc20Keeper.GetTokenPair(s.QubeticsChain.GetContext(), pairID)
		s.Require().True(found)
		activeDynamicPrecompiles := s.app.Erc20Keeper.GetParams(s.QubeticsChain.GetContext()).DynamicPrecompiles
		s.Require().Contains(activeDynamicPrecompiles, uosmoContractAddr.String())
	})
	It("Not native from source chain - should transfer and not register pair or deploy precompile", func() {
		// Send from Cosmos to Osmosis
		sender := s.IBCCosmosChain.SenderAccount.GetAddress().String()
		receiver := s.IBCOsmosisChain.SenderAccount.GetAddress().String()
		receiverAcc := sdk.MustAccAddressFromBech32(receiver)

		UatomInOsmosisDenomtrace := transfertypes.DenomTrace{
			Path:      "transfer/channel-1",
			BaseDenom: "uatom",
		}
		UatomInOsmosisIbcdenom := UatomInOsmosisDenomtrace.IBCDenom()
		uosmoContractAddr, err := utils.GetIBCDenomAddress(UatomInOsmosisIbcdenom)
		s.Require().NoError(err)
		params := s.app.EvmKeeper.GetParams(s.QubeticsChain.GetContext())
		s.Require().False(s.app.EvmKeeper.IsAvailableStaticPrecompile(&params, uosmoContractAddr))

		// check balance before transfer is 0
		ibcAtomBalanceBefore := s.app.BankKeeper.GetBalance(s.QubeticsChain.GetContext(), receiverAcc, teststypes.UatomOsmoIbcdenom)
		s.Require().Equal(int64(0), ibcAtomBalanceBefore.Amount.Int64())

		s.QubeticsChain.Coordinator.CommitBlock()
		s.SendBackCoins(s.pathOsmosisCosmos, s.IBCCosmosChain, "uatom", amount, sender, receiver, 1, "")

		// Balance of atom in Osmosis
		ibcOsmoBalanceAfter := s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), receiverAcc, UatomInOsmosisIbcdenom)
		s.Require().Equal(amount, ibcOsmoBalanceAfter.Amount.Int64())

		// Send ibc atom from osmosis account to our Qubetics account
		sender = s.IBCOsmosisChain.SenderAccount.GetAddress().String()
		receiver = s.QubeticsChain.SenderAccount.GetAddress().String()
		receiverAcc = sdk.MustAccAddressFromBech32(receiver)
		s.SendBackCoins(s.pathOsmosisQubetics, s.IBCOsmosisChain, UatomInOsmosisIbcdenom, amount, sender, receiver, 1, "transfer/channel-1/uatom")

		// check balance of ibc atom on qubetics after transfer
		ibcAtomBalanceAfter := s.app.BankKeeper.GetBalance(s.QubeticsChain.GetContext(), receiverAcc, teststypes.UatomOsmoIbcdenom)
		s.Require().Equal(amount, ibcAtomBalanceAfter.Amount.Int64())

		// Pair should not have been registered since atom is not native from osmosis
		// Precompile shouldnt be deployed
		pairID := s.app.Erc20Keeper.GetTokenPairID(s.QubeticsChain.GetContext(), teststypes.UatomOsmoIbcdenom)
		_, found := s.app.Erc20Keeper.GetTokenPair(s.QubeticsChain.GetContext(), pairID)
		s.Require().False(found)
		activeDynamicPrecompiles := s.app.Erc20Keeper.GetParams(s.QubeticsChain.GetContext()).DynamicPrecompiles
		s.Require().NotContains(activeDynamicPrecompiles, uosmoContractAddr.String())
	})
})
