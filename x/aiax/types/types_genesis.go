package types

func NewGenesisState(params Params) GenesisState {
  return GenesisState{
    Params: params,
  }
}

func (g GenesisState) Validate() error {
  return g.Params.Validate()
}

func DefaultGenesisState() *GenesisState {
  return &GenesisState{
    Params: DefaultParams(),
  }
}
