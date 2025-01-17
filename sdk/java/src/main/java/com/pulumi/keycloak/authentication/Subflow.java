// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.authentication;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.authentication.SubflowArgs;
import com.pulumi.keycloak.authentication.inputs.SubflowState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows for creating and managing an authentication subflow within Keycloak.
 * 
 * Like authentication flows, authentication subflows are containers for authentication executions.
 * As its name implies, an authentication subflow is contained in an authentication flow.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Authentication flows can be imported using the format `{{realmId}}/{{parentFlowAlias}}/{{authenticationSubflowId}}`.
 * 
 * The authentication subflow ID is typically a GUID which is autogenerated when the subflow is created via Keycloak.
 * 
 * Unfortunately, it is not trivial to retrieve the authentication subflow ID from the UI. The best way to do this is to visit the
 * 
 * &#34;Authentication&#34; page in Keycloak, and use the network tab of your browser to view the response of the API call to
 * 
 * `/auth/admin/realms/${realm}/authentication/flows/{flow}/executions`, which will be a list of executions, where the subflow will be.
 * 
 * __The subflow ID is contained in the `flowID` field__ (not, as one could guess, the `id` field).
 * 
 * Example:
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import keycloak:authentication/subflow:Subflow subflow my-realm/&#34;Parent Flow&#34;/3bad1172-bb5c-4a77-9615-c2606eb03081
 * ```
 * 
 */
@ResourceType(type="keycloak:authentication/subflow:Subflow")
public class Subflow extends com.pulumi.resources.CustomResource {
    /**
     * The alias for this authentication subflow.
     * 
     */
    @Export(name="alias", refs={String.class}, tree="[0]")
    private Output<String> alias;

    /**
     * @return The alias for this authentication subflow.
     * 
     */
    public Output<String> alias() {
        return this.alias;
    }
    /**
     * The name of the authenticator. Might be needed to be set with certain custom subflows with specific
     * authenticators. In general this will remain empty.
     * 
     */
    @Export(name="authenticator", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> authenticator;

    /**
     * @return The name of the authenticator. Might be needed to be set with certain custom subflows with specific
     * authenticators. In general this will remain empty.
     * 
     */
    public Output<Optional<String>> authenticator() {
        return Codegen.optional(this.authenticator);
    }
    /**
     * A description for the authentication subflow.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A description for the authentication subflow.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The alias for the parent authentication flow.
     * 
     */
    @Export(name="parentFlowAlias", refs={String.class}, tree="[0]")
    private Output<String> parentFlowAlias;

    /**
     * @return The alias for the parent authentication flow.
     * 
     */
    public Output<String> parentFlowAlias() {
        return this.parentFlowAlias;
    }
    /**
     * The type of authentication subflow to create. Valid choices include `basic-flow`, `form-flow`
     * and `client-flow`. Defaults to `basic-flow`.
     * 
     */
    @Export(name="providerId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> providerId;

    /**
     * @return The type of authentication subflow to create. Valid choices include `basic-flow`, `form-flow`
     * and `client-flow`. Defaults to `basic-flow`.
     * 
     */
    public Output<Optional<String>> providerId() {
        return Codegen.optional(this.providerId);
    }
    /**
     * The realm that the authentication subflow exists in.
     * 
     */
    @Export(name="realmId", refs={String.class}, tree="[0]")
    private Output<String> realmId;

    /**
     * @return The realm that the authentication subflow exists in.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }
    /**
     * The requirement setting, which can be one of `REQUIRED`, `ALTERNATIVE`, `OPTIONAL`, `CONDITIONAL`,
     * or `DISABLED`. Defaults to `DISABLED`.
     * 
     */
    @Export(name="requirement", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> requirement;

    /**
     * @return The requirement setting, which can be one of `REQUIRED`, `ALTERNATIVE`, `OPTIONAL`, `CONDITIONAL`,
     * or `DISABLED`. Defaults to `DISABLED`.
     * 
     */
    public Output<Optional<String>> requirement() {
        return Codegen.optional(this.requirement);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Subflow(java.lang.String name) {
        this(name, SubflowArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Subflow(java.lang.String name, SubflowArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Subflow(java.lang.String name, SubflowArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:authentication/subflow:Subflow", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Subflow(java.lang.String name, Output<java.lang.String> id, @Nullable SubflowState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:authentication/subflow:Subflow", name, state, makeResourceOptions(options, id), false);
    }

    private static SubflowArgs makeArgs(SubflowArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? SubflowArgs.Empty : args;
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
    public static Subflow get(java.lang.String name, Output<java.lang.String> id, @Nullable SubflowState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Subflow(name, id, state, options);
    }
}
