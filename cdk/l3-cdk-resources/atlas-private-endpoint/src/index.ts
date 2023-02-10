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


import {AtlasBasic, AtlasBasicProps} from "@mongodbatlas-awscdk/atlas-basic";
import {Construct} from "constructs";
import {
    CfnPrivateEndpoint, PrivateEndpoint,
} from "@mongodbatlas-awscdk/private-endpoint";

/** @type {*} */
const privateEndpointDefaults = {
    region: 'us-east-1'
};

/**
 * @description
 * @export
 * @class AtlasBasicPrivateEndpoint
 * @extends {Construct}
 */
export class AtlasBasicPrivateEndpoint extends Construct {

    readonly atlasBasic : AtlasBasic;
    readonly privateEndpoint : CfnPrivateEndpoint;

    /**
     * Creates an instance of AtlasBasicPrivateEndpoint.
     * @param {Construct} scope
     * @param {string} id
     * @param {AtlasPrivateEndpointProps} props
     * @memberof AtlasBasicPrivateEndpoint
     */
    constructor(scope: Construct, id: string, props: AtlasPrivateEndpointProps) {
        super(scope, id);
        validate(props)
        // Validation for atlas basic would be delegated to the library
        this.atlasBasic = new AtlasBasic(this, 'atlas-basic-'.concat(id),
            {
                ...props.atlasBasicProps,
            });
        this.privateEndpoint = new CfnPrivateEndpoint(this, 'private-endpoint-'.concat(id),
            {
                apiKeys: props.apiKeys,
                groupId: this.atlasBasic.mProject.ref,
                region: props.region || privateEndpointDefaults.region,
                ...props.privateEndpointProps
            });
    }
}

const validate = (props: AtlasPrivateEndpointProps) => {
    if (!props.atlasBasicProps.projectProps.orgId) {throw Error('Validation error: orgId is not defined');}

    if (!props.apiKeys.publicKey) {throw Error('Validation error: publicKey is not defined');}

    if (!props.apiKeys.privateKey) {throw Error('Validation error: privateKey is not defined');}
};


/**
 * @description
 * @export
 * @interface AtlasPrivateEndpointProps
 */
export interface AtlasPrivateEndpointProps {

    readonly apiKeys: ApiKeys;

    /**
     * @description
     * @type {string}
     * @default us-east-1
     * @memberof AtlasPrivateEndpointProps
     */
    readonly region?: string;
    /**
     * @description
     * @type {AtlasBasicProps}
     * @memberof AtlasPrivateEndpointProps
     */
    readonly atlasBasicProps : AtlasBasicProps;
    /**
     * @description
     * @type {CfnPrivateEndpointProps}
     * @memberof AtlasPrivateEndpointProps
     */
    readonly privateEndpointProps: PrivateEndpointProps;
}

export interface ApiKeys {
    /**
     * @description
     * @type {string}
     * @memberof AtlasPrivateEndpointProps
     */
    readonly privateKey: string;
    /**
     * @description
     * @type {string}
     * @memberof AtlasPrivateEndpointProps
     */
    readonly publicKey: string;
}

/**
 * @description
 * @export
 * @interface PrivateEndpointProps
 */
export interface PrivateEndpointProps {
    /**
     * @description
     * @type {string}
     * @memberof PrivateEndpointProps
     */
    readonly endpointServiceName?: string;
    /**
     * @description
     * @type {string}
     * @memberof PrivateEndpointProps
     */
    readonly errorMessage?: string;
    /**
     * @description
     * @type {string}
     * @memberof PrivateEndpointProps
     */
    readonly status?: string;
    /**
     * @description
     * @type {string}
     * @memberof PrivateEndpointProps
     */
    readonly groupId?: string;
    /**
     * @description
     * @type {string}
     * @memberof PrivateEndpointProps
     */
    readonly region?: string;
    /**
     * @description
     * @type {PrivateEndpoint[]}
     * @memberof PrivateEndpointProps
     */
    readonly privateEndpoints?: PrivateEndpoint[];

}
