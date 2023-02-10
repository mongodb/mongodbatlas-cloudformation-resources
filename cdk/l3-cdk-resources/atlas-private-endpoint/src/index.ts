import {ApiKeyDefinition, AtlasBasic, AtlasBasicProps, ProjectProps} from "@mongodbatlas-awscdk/atlas-basic";
import {Construct} from "constructs";
import {
    ApiKey,
    CfnPrivateEndpoint,
    CfnPrivateEndpointProps, PrivateEndpoint,
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

    readonly atlas : AtlasBasic;
    readonly private_endpoint : CfnPrivateEndpoint;

    /**
     * Creates an instance of AtlasBasicPrivateEndpoint.
     * @param {Construct} scope
     * @param {string} id
     * @param {AtlasPrivateEndpointProps} props
     * @memberof AtlasBasicPrivateEndpoint
     */
    constructor(scope: Construct, id: string, props: AtlasPrivateEndpointProps) {
        super(scope, id);
        
        this.atlas = new AtlasBasic(this, 'atlas-basic-'.concat(id),
            {
                apiKeys: props.apiKeys,
                ...props.atlasBasicProps,
            });

        this.private_endpoint = new CfnPrivateEndpoint(this, 'private-endpoint-'.concat(id),
            {
                apiKeys: props.apiKeys,
                groupId: 'this.atlas.',  //TODO
                region: props.region || privateEndpointDefaults.region,
                ...props.privateEndpointProps
            });
    }
}

/**
 * @description
 * @export
 * @interface AtlasPrivateEndpointProps
 */
export interface AtlasPrivateEndpointProps {
    /**
     * @description
     * @type {string}
     * @memberof AtlasPrivateEndpointProps
     */
    readonly groupId?: string;
    /**
     * @description
     * @type {ApiKeyDefinition}
     * @memberof AtlasPrivateEndpointProps
     */
    readonly apiKeys: ApiKeyDefinition;
    /**
     * @description
     * @type {string}
     * @default us-east-1
     * @memberof AtlasPrivateEndpointProps
     */
    readonly region: string;
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
     * @type {ApiKey}
     * @memberof PrivateEndpointProps
     */
    readonly apiKeys?: ApiKey;
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
