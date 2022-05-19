// Copyright 2022 Klaytn
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Generated by: OpenAPI Generator (https://openapi-generator.tech)

package types

// BlockResponse A BlockResponse includes a fully-populated block or a partially-populated block
// with a list of other transactions to fetch (other_transactions). As a result of the consensus
// algorithm of some blockchains, blocks can be omitted (i.e. certain block indices can be skipped).
// If a query for one of these omitted indices is made, the response should not include a `Block`
// object. It is VERY important to note that blocks MUST still form a canonical, connected chain of
// blocks where each block has a unique index. In other words, the `PartialBlockIdentifier` of a
// block after an omitted block should reference the last non-omitted block.
type BlockResponse struct {
	Block *Block `json:"block,omitempty"`
	// Some blockchains may require additional transactions to be fetched that weren't returned in
	// the block response (ex: block only returns transaction hashes). For blockchains with a lot of
	// transactions in each block, this can be very useful as consumers can concurrently fetch all
	// transactions returned.
	OtherTransactions []*TransactionIdentifier `json:"other_transactions,omitempty"`
}
