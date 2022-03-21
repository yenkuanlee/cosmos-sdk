<!--
order: 14
-->

# Transaction Tips

Transaction tips are a mechanism to pay for transaction fees using another denom than the native fee denom of the chain. {synopsis}

## Context

In a Cosmos ecosystem where more and more chains are connected via [IBC](https://ibc.cosmos.network/), it happens that users want to perform actions on chains where they don't have native tokens yet. An example would be an Osmosis user who wants to vote on a proposal on the Cosmos Hub, but they don't have ATOMs in their wallet. A solution would be to swap OSMO for ATOM just for voting on this proposal, but that is cumbersome.

Transaction tips is a new solution for cross-chain transaction fees payment. It is enabled by adding a new `Tip` field in the transaction itself.

We define the following parties:

- the tipper: this is the author of the transaction, who wants to execute a `Msg` on chain B, but doesn't have any native chain B tokens, only chain A tokens. In the example above, the tipper is the Osmosis user wanting to vote on a Cosmos Hub proposal.
- the feepayer: this is the party that will broadcast the final transaction on chain B. It doesn't need to be

We also assume that the tipper has already sent the fee they were willing to pay on chain B, using IBC. This means that chain A's bank module holds some IBC tokens under the tipper's address.

## Tipper and Feepayer Sign Modes
