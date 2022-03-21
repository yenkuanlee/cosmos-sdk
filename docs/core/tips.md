<!--
order: 14
-->

# Transaction Tips

Transaction tips are a mechanism to pay for transaction fees using another denom than the native fee denom of the chain. {synopsis}

## Context

In a Cosmos ecosystem where more and more chains are connected via [IBC](https://ibc.cosmos.network/), it happens that users want to perform actions on chains where they don't have native tokens yet. An example would be an Osmosis user who wants to vote on a proposal on the Cosmos Hub, but they don't have ATOMs in their wallet. A solution would be to swap OSMO for ATOM just for voting on this proposal, but that is cumbersome. Cross-chain DeFi project [Emeris](https://emeris.com/) is another use case.

Transaction tips is a new solution for cross-chain transaction fees payment. It is enabled by adding a new `Tip` field in the transaction itself.

We define the following parties:

- the tipper: this is the author of the transaction, who wants to execute a `Msg` on chain B, but doesn't have any native chain B tokens, only chain A tokens. In our example above, the tipper is the Osmosis user wanting to vote on a Cosmos Hub proposal.
- the feepayer: this is the party that will broadcast the final transaction on chain B. The tipper doesn't need to trust the feepayer.

We also assume that the tipper has already sent the fee they were willing to pay on chain B, using IBC. This means that chain A's bank module holds some IBC tokens under the tipper's address.

## Transaction Tips Flow

The transaction tips flow happens in multipe steps.

1. The tipper want to initiate a transaction on chain B (the target chain). They create the following `SignDocDirectAux` document:

+++ https://github.com/cosmos/cosmos-sdk/blob/79d23008502eee143ad7c8dd7f4061f860f6766f/proto/cosmos/tx/v1beta1/tx.proto#L67-L93

where `Tip` is defined as

+++ https://github.com/cosmos/cosmos-sdk/blob/79d23008502eee143ad7c8dd7f4061f860f6766f/proto/cosmos/tx/v1beta1/tx.proto#L219-L228

## Tipper and Feepayer Sign Modes
