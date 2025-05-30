// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.openid.ClientAuthorizationClientScopePolicyArgs;
import com.pulumi.keycloak.openid.inputs.ClientAuthorizationClientScopePolicyState;
import com.pulumi.keycloak.openid.outputs.ClientAuthorizationClientScopePolicyScope;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows you to manage openid Client Authorization Client Scope type Policies.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.keycloak.Realm;
 * import com.pulumi.keycloak.RealmArgs;
 * import com.pulumi.keycloak.openid.Client;
 * import com.pulumi.keycloak.openid.ClientArgs;
 * import com.pulumi.keycloak.openid.inputs.ClientAuthorizationArgs;
 * import com.pulumi.keycloak.openid.ClientScope;
 * import com.pulumi.keycloak.openid.ClientScopeArgs;
 * import com.pulumi.keycloak.openid.ClientAuthorizationClientScopePolicy;
 * import com.pulumi.keycloak.openid.ClientAuthorizationClientScopePolicyArgs;
 * import com.pulumi.keycloak.openid.inputs.ClientAuthorizationClientScopePolicyScopeArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var realm = new Realm("realm", RealmArgs.builder()
 *             .realm("my-realm")
 *             .enabled(true)
 *             .build());
 * 
 *         var test = new Client("test", ClientArgs.builder()
 *             .clientId("client_id")
 *             .realmId(realm.id())
 *             .accessType("CONFIDENTIAL")
 *             .serviceAccountsEnabled(true)
 *             .authorization(ClientAuthorizationArgs.builder()
 *                 .policyEnforcementMode("ENFORCING")
 *                 .build())
 *             .build());
 * 
 *         var test1 = new ClientScope("test1", ClientScopeArgs.builder()
 *             .realmId(realm.id())
 *             .name("test1")
 *             .description("test1")
 *             .build());
 * 
 *         var test2 = new ClientScope("test2", ClientScopeArgs.builder()
 *             .realmId(realm.id())
 *             .name("test2")
 *             .description("test2")
 *             .build());
 * 
 *         var testClientAuthorizationClientScopePolicy = new ClientAuthorizationClientScopePolicy("testClientAuthorizationClientScopePolicy", ClientAuthorizationClientScopePolicyArgs.builder()
 *             .resourceServerId(test.resourceServerId())
 *             .realmId(realm.id())
 *             .name("test_policy_single")
 *             .description("test")
 *             .decisionStrategy("AFFIRMATIVE")
 *             .logic("POSITIVE")
 *             .scopes(ClientAuthorizationClientScopePolicyScopeArgs.builder()
 *                 .id(test1.id())
 *                 .required(false)
 *                 .build())
 *             .build());
 * 
 *         var testMultiple = new ClientAuthorizationClientScopePolicy("testMultiple", ClientAuthorizationClientScopePolicyArgs.builder()
 *             .resourceServerId(test.resourceServerId())
 *             .realmId(realm.id())
 *             .name("test_policy_multiple")
 *             .description("test")
 *             .decisionStrategy("AFFIRMATIVE")
 *             .logic("POSITIVE")
 *             .scopes(            
 *                 ClientAuthorizationClientScopePolicyScopeArgs.builder()
 *                     .id(test1.id())
 *                     .required(false)
 *                     .build(),
 *                 ClientAuthorizationClientScopePolicyScopeArgs.builder()
 *                     .id(test2.id())
 *                     .required(true)
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Argument Reference
 * 
 * The following arguments are supported:
 * 
 * - `realm_id` - (Required) The realm this group exists in.
 * - `resource_server_id` - (Required) The ID of the resource server.
 * - `name` - (Required) The name of the policy.
 * - `description` - (Optional) A description for the authorization policy.
 * - `decision_strategy` - (Optional) The decision strategy, can be one of `UNANIMOUS`, `AFFIRMATIVE`, or `CONSENSUS`. Defaults to `UNANIMOUS`.
 * - `logic` - (Optional) The logic, can be one of `POSITIVE` or `NEGATIVE`. Defaults to `POSITIVE`.
 * - `scope` - An client scope to add client scope. At least one should be defined.
 * 
 * ### Scope Arguments
 * 
 * - `id` - (Required) Id of client scope.
 * - `required` - (Optional) When `true`, then this client scope will be set as required. Defaults to `false`.
 * 
 * ### Attributes Reference
 * 
 * In addition to the arguments listed above, the following computed attributes are exported:
 * 
 * - `id` - Policy ID representing the policy.
 * 
 * ## Import
 * 
 * Client authorization policies can be imported using the format: `{{realmId}}/{{resourceServerId}}/{{policyId}}`.
 * 
 * Example:
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import keycloak:openid/clientAuthorizationClientScopePolicy:ClientAuthorizationClientScopePolicy test my-realm/3bd4a686-1062-4b59-97b8-e4e3f10b99da/63b3cde8-987d-4cd9-9306-1955579281d9
 * ```
 * 
 */
@ResourceType(type="keycloak:openid/clientAuthorizationClientScopePolicy:ClientAuthorizationClientScopePolicy")
public class ClientAuthorizationClientScopePolicy extends com.pulumi.resources.CustomResource {
    @Export(name="decisionStrategy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> decisionStrategy;

    public Output<Optional<String>> decisionStrategy() {
        return Codegen.optional(this.decisionStrategy);
    }
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    @Export(name="logic", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> logic;

    public Output<Optional<String>> logic() {
        return Codegen.optional(this.logic);
    }
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    public Output<String> name() {
        return this.name;
    }
    @Export(name="realmId", refs={String.class}, tree="[0]")
    private Output<String> realmId;

    public Output<String> realmId() {
        return this.realmId;
    }
    @Export(name="resourceServerId", refs={String.class}, tree="[0]")
    private Output<String> resourceServerId;

    public Output<String> resourceServerId() {
        return this.resourceServerId;
    }
    @Export(name="scopes", refs={List.class,ClientAuthorizationClientScopePolicyScope.class}, tree="[0,1]")
    private Output<List<ClientAuthorizationClientScopePolicyScope>> scopes;

    public Output<List<ClientAuthorizationClientScopePolicyScope>> scopes() {
        return this.scopes;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ClientAuthorizationClientScopePolicy(java.lang.String name) {
        this(name, ClientAuthorizationClientScopePolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ClientAuthorizationClientScopePolicy(java.lang.String name, ClientAuthorizationClientScopePolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ClientAuthorizationClientScopePolicy(java.lang.String name, ClientAuthorizationClientScopePolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:openid/clientAuthorizationClientScopePolicy:ClientAuthorizationClientScopePolicy", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ClientAuthorizationClientScopePolicy(java.lang.String name, Output<java.lang.String> id, @Nullable ClientAuthorizationClientScopePolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:openid/clientAuthorizationClientScopePolicy:ClientAuthorizationClientScopePolicy", name, state, makeResourceOptions(options, id), false);
    }

    private static ClientAuthorizationClientScopePolicyArgs makeArgs(ClientAuthorizationClientScopePolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ClientAuthorizationClientScopePolicyArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static ClientAuthorizationClientScopePolicy get(java.lang.String name, Output<java.lang.String> id, @Nullable ClientAuthorizationClientScopePolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ClientAuthorizationClientScopePolicy(name, id, state, options);
    }
}
