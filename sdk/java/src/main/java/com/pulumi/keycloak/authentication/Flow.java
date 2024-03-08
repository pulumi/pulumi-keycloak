// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.authentication;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.authentication.FlowArgs;
import com.pulumi.keycloak.authentication.inputs.FlowState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows for creating and managing an authentication flow within Keycloak.
 * 
 * [Authentication flows](https://www.keycloak.org/docs/11.0/server_admin/index.html#_authentication-flows) describe a sequence
 * of actions that a user or service must perform in order to be authenticated to Keycloak. The authentication flow itself
 * is a container for these actions, which are otherwise known as executions.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
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
 *         var execution = new Execution(&#34;execution&#34;, ExecutionArgs.builder()        
 *             .realmId(realm.id())
 *             .parentFlowAlias(flow.alias())
 *             .authenticator(&#34;identity-provider-redirector&#34;)
 *             .requirement(&#34;REQUIRED&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Authentication flows can be imported using the format `{{realmId}}/{{authenticationFlowId}}`. The authentication flow ID is
 * 
 * typically a GUID which is autogenerated when the flow is created via Keycloak.
 * 
 * Unfortunately, it is not trivial to retrieve the authentication flow ID from the UI. The best way to do this is to visit the
 * 
 * &#34;Authentication&#34; page in Keycloak, and use the network tab of your browser to view the response of the API call to `/auth/admin/realms/${realm}/authentication/flows`,
 * 
 * which will be a list of authentication flows.
 * 
 * Example:
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import keycloak:authentication/flow:Flow flow my-realm/e9a5641e-778c-4daf-89c0-f4ef617987d1
 * ```
 * 
 */
@ResourceType(type="keycloak:authentication/flow:Flow")
public class Flow extends com.pulumi.resources.CustomResource {
    /**
     * The alias for this authentication flow.
     * 
     */
    @Export(name="alias", refs={String.class}, tree="[0]")
    private Output<String> alias;

    /**
     * @return The alias for this authentication flow.
     * 
     */
    public Output<String> alias() {
        return this.alias;
    }
    /**
     * A description for the authentication flow.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A description for the authentication flow.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The type of authentication flow to create. Valid choices include `basic-flow` and `client-flow`. Defaults to `basic-flow`.
     * 
     */
    @Export(name="providerId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> providerId;

    /**
     * @return The type of authentication flow to create. Valid choices include `basic-flow` and `client-flow`. Defaults to `basic-flow`.
     * 
     */
    public Output<Optional<String>> providerId() {
        return Codegen.optional(this.providerId);
    }
    /**
     * The realm that the authentication flow exists in.
     * 
     */
    @Export(name="realmId", refs={String.class}, tree="[0]")
    private Output<String> realmId;

    /**
     * @return The realm that the authentication flow exists in.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Flow(String name) {
        this(name, FlowArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Flow(String name, FlowArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Flow(String name, FlowArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:authentication/flow:Flow", name, args == null ? FlowArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Flow(String name, Output<String> id, @Nullable FlowState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:authentication/flow:Flow", name, state, makeResourceOptions(options, id));
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
    public static Flow get(String name, Output<String> id, @Nullable FlowState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Flow(name, id, state, options);
    }
}
