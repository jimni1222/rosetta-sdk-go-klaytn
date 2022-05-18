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

// RelatedTransaction The related_transaction allows implementations to link together multiple
// transactions. An unpopulated network identifier indicates that the related transaction is on the
// same network.
type RelatedTransaction struct {
	NetworkIdentifier     *NetworkIdentifier     `json:"network_identifier,omitempty"`
	TransactionIdentifier *TransactionIdentifier `json:"transaction_identifier"`
	Direction             Direction              `json:"direction"`
}
