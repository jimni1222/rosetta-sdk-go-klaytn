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

// CoinChange CoinChange is used to represent a change in state of a some coin identified by a
// coin_identifier. This object is part of the Operation model and must be populated for UTXO-based
// blockchains. Coincidentally, this abstraction of UTXOs allows for supporting both account-based
// transfers and UTXO-based transfers on the same blockchain (when a transfer is account-based,
// don't populate this model).
type CoinChange struct {
	CoinIdentifier *CoinIdentifier `json:"coin_identifier"`
	CoinAction     CoinAction      `json:"coin_action"`
}
