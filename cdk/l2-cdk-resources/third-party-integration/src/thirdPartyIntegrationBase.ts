// Copyright 2023 MongoDB Inc
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

export interface ApiKeyDefinition {
  /**
   * Atlas private API key.
   */
  readonly privateKey: string;

  /**
   * Atlas public API key.
   */
  readonly publicKey: string;
}

export interface ThirdPartyIntegrationProps {

  /**
   * Unique 24-hexadecimal digit string that identifies your project.
   */
  readonly projectId: string;

  /**
   * Atlas API keys.
   */
  readonly apiKeys: ApiKeyDefinition;
}

