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

// AccountBalanceRequest An AccountBalanceRequest is utilized to make a balance request on the
// /account/balance endpoint. If the block_identifier is populated, a historical balance query
// should be performed.
type AccountBalanceRequest struct {
	NetworkIdentifier *NetworkIdentifier      `json:"network_identifier"`
	AccountIdentifier *AccountIdentifier      `json:"account_identifier"`
	BlockIdentifier   *PartialBlockIdentifier `json:"block_identifier,omitempty"`
	// In some cases, the caller may not want to retrieve all available balances for an
	// AccountIdentifier. If the currencies field is populated, only balances for the specified
	// currencies will be returned. If not populated, all available balances will be returned.
	Currencies []*Currency `json:"currencies,omitempty"`
}
