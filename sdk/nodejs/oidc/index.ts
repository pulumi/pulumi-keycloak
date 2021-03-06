// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./googleIdentityProvider";
export * from "./identityProvider";

// Import resources to register:
import { GoogleIdentityProvider } from "./googleIdentityProvider";
import { IdentityProvider } from "./identityProvider";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "keycloak:oidc/googleIdentityProvider:GoogleIdentityProvider":
                return new GoogleIdentityProvider(name, <any>undefined, { urn })
            case "keycloak:oidc/identityProvider:IdentityProvider":
                return new IdentityProvider(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("keycloak", "oidc/googleIdentityProvider", _module)
pulumi.runtime.registerResourceModule("keycloak", "oidc/identityProvider", _module)
