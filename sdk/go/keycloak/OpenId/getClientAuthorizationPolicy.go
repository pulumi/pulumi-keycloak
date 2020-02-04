// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package OpenId

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

func LookupClientAuthorizationPolicy(ctx *pulumi.Context, args *GetClientAuthorizationPolicyArgs) (*GetClientAuthorizationPolicyResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
		inputs["realmId"] = args.RealmId
		inputs["resourceServerId"] = args.ResourceServerId
	}
	outputs, err := ctx.Invoke("keycloak:OpenId/getClientAuthorizationPolicy:getClientAuthorizationPolicy", inputs)
	if err != nil {
		return nil, err
	}
	return &GetClientAuthorizationPolicyResult{
		DecisionStrategy: outputs["decisionStrategy"],
		Logic: outputs["logic"],
		Name: outputs["name"],
		Owner: outputs["owner"],
		Policies: outputs["policies"],
		RealmId: outputs["realmId"],
		ResourceServerId: outputs["resourceServerId"],
		Resources: outputs["resources"],
		Scopes: outputs["scopes"],
		Type: outputs["type"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getClientAuthorizationPolicy.
type GetClientAuthorizationPolicyArgs struct {
	Name interface{}
	RealmId interface{}
	ResourceServerId interface{}
}

// A collection of values returned by getClientAuthorizationPolicy.
type GetClientAuthorizationPolicyResult struct {
	DecisionStrategy interface{}
	Logic interface{}
	Name interface{}
	Owner interface{}
	Policies interface{}
	RealmId interface{}
	ResourceServerId interface{}
	Resources interface{}
	Scopes interface{}
	Type interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
