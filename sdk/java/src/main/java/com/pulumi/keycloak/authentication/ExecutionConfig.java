// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.authentication;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.authentication.ExecutionConfigArgs;
import com.pulumi.keycloak.authentication.inputs.ExecutionConfigState;
import java.lang.String;
import java.util.Map;
import javax.annotation.Nullable;

/**
 * Allows for managing an authentication execution&#39;s configuration. If a particular authentication execution supports additional
 * configuration (such as with the `identity-provider-redirector` execution), this can be managed with this resource.
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
 * import com.pulumi.keycloak.authentication.ExecutionConfig;
 * import com.pulumi.keycloak.authentication.ExecutionConfigArgs;
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
 *             .build());
 * 
 *         var config = new ExecutionConfig(&#34;config&#34;, ExecutionConfigArgs.builder()        
 *             .realmId(realm.id())
 *             .executionId(execution.id())
 *             .alias(&#34;my-config-alias&#34;)
 *             .config(Map.of(&#34;defaultProvider&#34;, &#34;my-config-default-idp&#34;))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Configurations can be imported using the format `{{realm}}/{{authenticationExecutionId}}/{{authenticationExecutionConfigId}}`. If the `authenticationExecutionId` is incorrect, the import will still be successful. A subsequent apply will change the `authenticationExecutionId` to the correct one, which causes the configuration to be replaced. Examplebash
 * 
 * ```sh
 *  $ pulumi import keycloak:authentication/executionConfig:ExecutionConfig config my-realm/be081463-ddbf-4b42-9eff-9c97886f24ff/30559fcf-6fb8-45ea-8c46-2b86f46ebc17
 * ```
 * 
 */
@ResourceType(type="keycloak:authentication/executionConfig:ExecutionConfig")
public class ExecutionConfig extends com.pulumi.resources.CustomResource {
    /**
     * The name of the configuration.
     * 
     */
    @Export(name="alias", type=String.class, parameters={})
    private Output<String> alias;

    /**
     * @return The name of the configuration.
     * 
     */
    public Output<String> alias() {
        return this.alias;
    }
    /**
     * The configuration. Keys are specific to each configurable authentication execution and not checked when applying.
     * 
     */
    @Export(name="config", type=Map.class, parameters={String.class, String.class})
    private Output<Map<String,String>> config;

    /**
     * @return The configuration. Keys are specific to each configurable authentication execution and not checked when applying.
     * 
     */
    public Output<Map<String,String>> config() {
        return this.config;
    }
    /**
     * The authentication execution this configuration is attached to.
     * 
     */
    @Export(name="executionId", type=String.class, parameters={})
    private Output<String> executionId;

    /**
     * @return The authentication execution this configuration is attached to.
     * 
     */
    public Output<String> executionId() {
        return this.executionId;
    }
    /**
     * The realm the authentication execution exists in.
     * 
     */
    @Export(name="realmId", type=String.class, parameters={})
    private Output<String> realmId;

    /**
     * @return The realm the authentication execution exists in.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ExecutionConfig(String name) {
        this(name, ExecutionConfigArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ExecutionConfig(String name, ExecutionConfigArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ExecutionConfig(String name, ExecutionConfigArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:authentication/executionConfig:ExecutionConfig", name, args == null ? ExecutionConfigArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ExecutionConfig(String name, Output<String> id, @Nullable ExecutionConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:authentication/executionConfig:ExecutionConfig", name, state, makeResourceOptions(options, id));
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
    public static ExecutionConfig get(String name, Output<String> id, @Nullable ExecutionConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ExecutionConfig(name, id, state, options);
    }
}