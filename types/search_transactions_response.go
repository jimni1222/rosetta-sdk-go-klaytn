// Copyright 2022 Coinbase, Inc.
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
//
// Modifications Copyright © 2022 Klaytn
// Modified and improved for the Klaytn development.

// Generated by: OpenAPI Generator (https://openapi-generator.tech)

package types

// SearchTransactionsResponse SearchTransactionsResponse contains an ordered collection of
// BlockTransactions that match the query in SearchTransactionsRequest. These BlockTransactions are
// sorted from most recent block to oldest block.
type SearchTransactionsResponse struct {
	// transactions is an array of BlockTransactions sorted by most recent BlockIdentifier (meaning
	// that transactions in recent blocks appear first). If there are many transactions for a
	// particular search, transactions may not contain all matching transactions. It is up to the
	// caller to paginate these transactions using the max_block field.
	Transactions []*BlockTransaction `json:"transactions"`
	// total_count is the number of results for a given search. Callers typically use this value to
	// concurrently fetch results by offset or to display a virtual page number associated with
	// results.
	TotalCount int64 `json:"total_count"`
	// next_offset is the next offset to use when paginating through transaction results. If this
	// field is not populated, there are no more transactions to query.
	NextOffset *int64 `json:"next_offset,omitempty"`
}
