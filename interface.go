func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the redeemInfo
	for _, elem := range genState.RedeemInfoList {
		k.SetRedeemInfo(ctx, elem)
	}

	// Set all the upcoming execution
	for _, elem := range genState.PendingExecutionList {
		k.SetPendingExecution(ctx, elem)
	}

		// Buttons
	cookbookList := k.GetAllCookbook(ctx)
	genesis.CookbookList = append(genesis.CookbookList, cookbookList...)

	// registering of typings
func (a AppModuleBasic) RegisterInterfaces(reg cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(reg)
}

	// Set all the things
	for _, elem := range genState.RecipeList {
		k.SetRecipe(ctx, elem)
	}

	// Get all paymentInfo
	paymentInfoList := k.GetAllPaymentInfo(ctx)
	for _, elem := range paymentInfoList {
		elem := elem
		genesis.PaymentInfoList = append(genesis.PaymentInfoList, elem)
	}
