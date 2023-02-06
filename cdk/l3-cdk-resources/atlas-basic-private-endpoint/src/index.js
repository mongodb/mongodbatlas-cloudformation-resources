"use strict";
var __extends = (this && this.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (Object.prototype.hasOwnProperty.call(b, p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        if (typeof b !== "function" && b !== null)
            throw new TypeError("Class extends value " + String(b) + " is not a constructor or null");
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();
exports.__esModule = true;
var privateEndpoint = require("@mongodbatlas-awscdk/private-endpoint");
var cdk = require("aws-cdk-lib");
var aws_cdk_lib_1 = require("aws-cdk-lib");
var AtlasBasePrivateEndpoint = /** @class */ (function (_super) {
    __extends(AtlasBasePrivateEndpoint, _super);
    function AtlasBasePrivateEndpoint(scope, id, props) {
        var _this = _super.call(this, scope, id, props) || this;
        var privEnd = new privateEndpoint.CfnPrivateEndpoint(_this, 'MyProject', {
            groupId: props.projectId,
            region: props.region,
            apiKeys: { privateKey: props.apiKeys.privateKey, publicKey: props.apiKeys.publicKey },
            privateEndpoints: [{ vpcId: props.vpcId, subnetIds: [props.subnetId] }]
        });
        return _this;
    }
    return AtlasBasePrivateEndpoint;
}(cdk.Stack));
var modal = {
    projectId: '',
    vpcId: 'vpc-',
    subnetId: 'subnet-',
    region: 'us-east-1',
    apiKeys: { privateKey: '', publicKey: '' }
};
var app = new aws_cdk_lib_1.App();
var greeter = new AtlasBasePrivateEndpoint(app, 'test-something', modal);
app.synth();
