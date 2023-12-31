// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.openid.AudienceResolveProtocolMappterArgs;
import com.pulumi.keycloak.openid.inputs.AudienceResolveProtocolMappterState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * @deprecated
 * keycloak.openid/audienceresolveprotocolmappter.AudienceResolveProtocolMappter has been deprecated in favor of keycloak.openid/audienceresolveprotocolmapper.AudienceResolveProtocolMapper
 * 
 */
@Deprecated /* keycloak.openid/audienceresolveprotocolmappter.AudienceResolveProtocolMappter has been deprecated in favor of keycloak.openid/audienceresolveprotocolmapper.AudienceResolveProtocolMapper */
@ResourceType(type="keycloak:openid/audienceResolveProtocolMappter:AudienceResolveProtocolMappter")
public class AudienceResolveProtocolMappter extends com.pulumi.resources.CustomResource {
    /**
     * The mapper&#39;s associated client. Cannot be used at the same time as client_scope_id.
     * 
     */
    @Export(name="clientId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientId;

    /**
     * @return The mapper&#39;s associated client. Cannot be used at the same time as client_scope_id.
     * 
     */
    public Output<Optional<String>> clientId() {
        return Codegen.optional(this.clientId);
    }
    /**
     * The mapper&#39;s associated client scope. Cannot be used at the same time as client_id.
     * 
     */
    @Export(name="clientScopeId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientScopeId;

    /**
     * @return The mapper&#39;s associated client scope. Cannot be used at the same time as client_id.
     * 
     */
    public Output<Optional<String>> clientScopeId() {
        return Codegen.optional(this.clientScopeId);
    }
    /**
     * A human-friendly name that will appear in the Keycloak console.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A human-friendly name that will appear in the Keycloak console.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The realm id where the associated client or client scope exists.
     * 
     */
    @Export(name="realmId", refs={String.class}, tree="[0]")
    private Output<String> realmId;

    /**
     * @return The realm id where the associated client or client scope exists.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AudienceResolveProtocolMappter(String name) {
        this(name, AudienceResolveProtocolMappterArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AudienceResolveProtocolMappter(String name, AudienceResolveProtocolMappterArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AudienceResolveProtocolMappter(String name, AudienceResolveProtocolMappterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:openid/audienceResolveProtocolMappter:AudienceResolveProtocolMappter", name, args == null ? AudienceResolveProtocolMappterArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AudienceResolveProtocolMappter(String name, Output<String> id, @Nullable AudienceResolveProtocolMappterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:openid/audienceResolveProtocolMappter:AudienceResolveProtocolMappter", name, state, makeResourceOptions(options, id));
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
    public static AudienceResolveProtocolMappter get(String name, Output<String> id, @Nullable AudienceResolveProtocolMappterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AudienceResolveProtocolMappter(name, id, state, options);
    }
}
