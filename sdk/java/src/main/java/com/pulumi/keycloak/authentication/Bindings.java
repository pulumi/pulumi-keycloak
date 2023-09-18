// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.authentication;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.authentication.BindingsArgs;
import com.pulumi.keycloak.authentication.inputs.BindingsState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Allows for creating and managing realm authentication flow bindings within Keycloak.
 * 
 * [Authentication flows](https://www.keycloak.org/docs/latest/server_admin/index.html#_authentication-flows) describe a sequence
 * of actions that a user or service must perform in order to be authenticated to Keycloak. The authentication flow itself
 * is a container for these actions, which are otherwise known as executions.
 * 
 * Realms assign authentication flows to supported user flows such as `registration` and `browser`. This resource allows the
 * updating of realm authentication flow bindings to custom authentication flows created by `keycloak.authentication.Flow`.
 * 
 * Note that you can also use the `keycloak.Realm` resource to assign authentication flow bindings at the realm level. This
 * resource is useful if you would like to create a realm and an authentication flow, and assign this flow to the realm within
 * a single run of `pulumi up`. In any case, do not attempt to use both the arguments within the `keycloak.Realm` resource
 * and this resource to manage authentication flow bindings, you should choose one or the other.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.keycloak.Realm;
 * import com.pulumi.keycloak.RealmArgs;
 * import com.pulumi.keycloak.authentication.Flow;
 * import com.pulumi.keycloak.authentication.FlowArgs;
 * import com.pulumi.keycloak.authentication.Execution;
 * import com.pulumi.keycloak.authentication.ExecutionArgs;
 * import com.pulumi.keycloak.authentication.Bindings;
 * import com.pulumi.keycloak.authentication.BindingsArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var realm = new Realm(&#34;realm&#34;, RealmArgs.builder()        
 *             .realm(&#34;my-realm&#34;)
 *             .enabled(true)
 *             .build());
 * 
 *         var flow = new Flow(&#34;flow&#34;, FlowArgs.builder()        
 *             .realmId(realm.id())
 *             .alias(&#34;my-flow-alias&#34;)
 *             .build());
 * 
 *         var executionOne = new Execution(&#34;executionOne&#34;, ExecutionArgs.builder()        
 *             .realmId(realm.id())
 *             .parentFlowAlias(flow.alias())
 *             .authenticator(&#34;auth-cookie&#34;)
 *             .requirement(&#34;ALTERNATIVE&#34;)
 *             .build());
 * 
 *         var executionTwo = new Execution(&#34;executionTwo&#34;, ExecutionArgs.builder()        
 *             .realmId(realm.id())
 *             .parentFlowAlias(flow.alias())
 *             .authenticator(&#34;identity-provider-redirector&#34;)
 *             .requirement(&#34;ALTERNATIVE&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(executionOne)
 *                 .build());
 * 
 *         var browserAuthenticationBinding = new Bindings(&#34;browserAuthenticationBinding&#34;, BindingsArgs.builder()        
 *             .realmId(realm.id())
 *             .browserFlow(flow.alias())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="keycloak:authentication/bindings:Bindings")
public class Bindings extends com.pulumi.resources.CustomResource {
    /**
     * The alias of the flow to assign to the realm BrowserFlow.
     * 
     */
    @Export(name="browserFlow", refs={String.class}, tree="[0]")
    private Output<String> browserFlow;

    /**
     * @return The alias of the flow to assign to the realm BrowserFlow.
     * 
     */
    public Output<String> browserFlow() {
        return this.browserFlow;
    }
    /**
     * The alias of the flow to assign to the realm ClientAuthenticationFlow.
     * 
     */
    @Export(name="clientAuthenticationFlow", refs={String.class}, tree="[0]")
    private Output<String> clientAuthenticationFlow;

    /**
     * @return The alias of the flow to assign to the realm ClientAuthenticationFlow.
     * 
     */
    public Output<String> clientAuthenticationFlow() {
        return this.clientAuthenticationFlow;
    }
    /**
     * The alias of the flow to assign to the realm DirectGrantFlow.
     * 
     */
    @Export(name="directGrantFlow", refs={String.class}, tree="[0]")
    private Output<String> directGrantFlow;

    /**
     * @return The alias of the flow to assign to the realm DirectGrantFlow.
     * 
     */
    public Output<String> directGrantFlow() {
        return this.directGrantFlow;
    }
    /**
     * The alias of the flow to assign to the realm DockerAuthenticationFlow.
     * 
     */
    @Export(name="dockerAuthenticationFlow", refs={String.class}, tree="[0]")
    private Output<String> dockerAuthenticationFlow;

    /**
     * @return The alias of the flow to assign to the realm DockerAuthenticationFlow.
     * 
     */
    public Output<String> dockerAuthenticationFlow() {
        return this.dockerAuthenticationFlow;
    }
    /**
     * The realm the authentication flow binding exists in.
     * 
     */
    @Export(name="realmId", refs={String.class}, tree="[0]")
    private Output<String> realmId;

    /**
     * @return The realm the authentication flow binding exists in.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }
    /**
     * The alias of the flow to assign to the realm RegistrationFlow.
     * 
     */
    @Export(name="registrationFlow", refs={String.class}, tree="[0]")
    private Output<String> registrationFlow;

    /**
     * @return The alias of the flow to assign to the realm RegistrationFlow.
     * 
     */
    public Output<String> registrationFlow() {
        return this.registrationFlow;
    }
    /**
     * The alias of the flow to assign to the realm ResetCredentialsFlow.
     * 
     */
    @Export(name="resetCredentialsFlow", refs={String.class}, tree="[0]")
    private Output<String> resetCredentialsFlow;

    /**
     * @return The alias of the flow to assign to the realm ResetCredentialsFlow.
     * 
     */
    public Output<String> resetCredentialsFlow() {
        return this.resetCredentialsFlow;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Bindings(String name) {
        this(name, BindingsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Bindings(String name, BindingsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Bindings(String name, BindingsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:authentication/bindings:Bindings", name, args == null ? BindingsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Bindings(String name, Output<String> id, @Nullable BindingsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:authentication/bindings:Bindings", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static Bindings get(String name, Output<String> id, @Nullable BindingsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Bindings(name, id, state, options);
    }
}
