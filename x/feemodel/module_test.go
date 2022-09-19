package feemodel_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/CoreumFoundation/coreum/app"
	"github.com/CoreumFoundation/coreum/x/feemodel"
	"github.com/CoreumFoundation/coreum/x/feemodel/types"
)

func newKeeperMock(genesisState types.GenesisState) *keeperMock {
	return &keeperMock{
		state: genesisState,
	}
}

type keeperMock struct {
	state types.GenesisState
}

func (k *keeperMock) TrackedGas(ctx sdk.Context) int64 {
	return 1
}

func (k *keeperMock) SetParams(ctx sdk.Context, params types.Params) {
	k.state.Params = params
}

func (k *keeperMock) GetParams(ctx sdk.Context) types.Params {
	return k.state.Params
}

func (k *keeperMock) GetShortEMAGas(ctx sdk.Context) int64 {
	return 0
}

func (k *keeperMock) SetShortEMAGas(ctx sdk.Context, emaGas int64) {}

func (k *keeperMock) GetLongEMAGas(ctx sdk.Context) int64 {
	return 0
}

func (k *keeperMock) SetLongEMAGas(ctx sdk.Context, emaGas int64) {}

func (k *keeperMock) GetMinGasPrice(ctx sdk.Context) sdk.DecCoin {
	return k.state.MinGasPrice
}

func (k *keeperMock) SetMinGasPrice(ctx sdk.Context, minGasPrice sdk.DecCoin) {
	k.state.MinGasPrice = minGasPrice
}

func setup() (feemodel.AppModule, feemodel.Keeper, types.GenesisState, codec.Codec) {
	genesisState := types.GenesisState{
		Params: types.Params{
			InitialGasPrice:         sdk.NewDec(15),
			MaxGasPrice:             sdk.NewDec(150),
			MaxDiscount:             sdk.MustNewDecFromStr("0.1"),
			EscalationStartBlockGas: 7,
			MaxBlockGas:             10,
			ShortEmaBlockLength:     1,
			LongEmaBlockLength:      3,
		},
		MinGasPrice: sdk.NewDecCoin("coin", sdk.NewInt(155)),
	}
	cdc := app.NewEncodingConfig().Codec
	keeper := newKeeperMock(genesisState)
	module := feemodel.NewAppModule(keeper)

	return module, keeper, genesisState, cdc
}

func TestInitGenesis(t *testing.T) {
	module, keeper, state, cdc := setup()

	genesisState := state
	genesisState.Params.InitialGasPrice.Add(sdk.OneDec())
	genesisState.Params.MaxGasPrice.Add(sdk.OneDec())
	genesisState.Params.MaxDiscount.Add(sdk.MustNewDecFromStr("0.2"))
	genesisState.Params.EscalationStartBlockGas++
	genesisState.Params.MaxBlockGas++
	genesisState.Params.ShortEmaBlockLength++
	genesisState.Params.LongEmaBlockLength++
	genesisState.MinGasPrice.Denom = "coin2"
	genesisState.MinGasPrice.Amount.Add(sdk.OneDec())

	module.InitGenesis(sdk.Context{}, cdc, cdc.MustMarshalJSON(&genesisState))

	params := keeper.GetParams(sdk.Context{})
	minGasPrice := keeper.GetMinGasPrice(sdk.Context{})
	assert.True(t, genesisState.Params.InitialGasPrice.Equal(params.InitialGasPrice))
	assert.True(t, genesisState.Params.MaxGasPrice.Equal(params.MaxGasPrice))
	assert.True(t, genesisState.Params.MaxDiscount.Equal(params.MaxDiscount))
	assert.Equal(t, genesisState.Params.EscalationStartBlockGas, params.EscalationStartBlockGas)
	assert.Equal(t, genesisState.Params.MaxBlockGas, params.MaxBlockGas)
	assert.Equal(t, genesisState.Params.ShortEmaBlockLength, params.ShortEmaBlockLength)
	assert.Equal(t, genesisState.Params.LongEmaBlockLength, params.LongEmaBlockLength)
	assert.Equal(t, genesisState.MinGasPrice.Denom, minGasPrice.Denom)
	assert.True(t, genesisState.MinGasPrice.Amount.Equal(minGasPrice.Amount))
}

func TestExport(t *testing.T) {
	module, _, state, cdc := setup()

	var decodedGenesis types.GenesisState
	require.NoError(t, cdc.UnmarshalJSON(module.ExportGenesis(sdk.Context{}, cdc), &decodedGenesis))
	assert.EqualValues(t, state, decodedGenesis)
}

func TestEndBlock(t *testing.T) {
	module, keeper, state, _ := setup()

	module.EndBlock(sdk.Context{}, abci.RequestEndBlock{})

	model := types.NewModel(state.Params)
	minGasPrice := keeper.GetMinGasPrice(sdk.Context{})
	assert.True(t, minGasPrice.Amount.Equal(model.CalculateGasPriceWithMaxDiscount()))
	assert.Equal(t, minGasPrice.Denom, state.MinGasPrice.Denom)
}
