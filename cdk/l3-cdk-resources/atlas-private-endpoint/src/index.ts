import {ApiKeyDefinition, AtlasBasic, AtlasBasicProps, ProjectProps} from "@mongodbatlas-awscdk/atlas-basic";
import {Construct} from "constructs";
import {
    CfnPrivateEndpoint,
    CfnPrivateEndpointProps,
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
    /**
     * Creates an instance of AtlasBasicPrivateEndpoint.
     * @param {Construct} scope
     * @param {string} id
     * @param {AtlasPrivateEndpointProps} props
     * @memberof AtlasBasicPrivateEndpoint
     */
    constructor(scope: Construct, id: string, props: AtlasPrivateEndpointProps) {
        super(scope, id);
        
        const atlas_basic = new AtlasBasic(this, 'atlas-basic-'.concat(id),
            {
                apiKeys: props.apiKeys,
                ...props.atlasBasicProps,
            }
            );

        new CfnPrivateEndpoint(this, 'private-endpoint-'.concat(id),
            {
                apiKeys: props.apiKeys,
                ...props.privateEndpointProps,
            }
           );
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
    readonly groupId: string;
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
    readonly privateEndpointProps: CfnPrivateEndpointProps;
}