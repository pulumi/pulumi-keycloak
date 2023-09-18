// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.ldap;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.ldap.MsadLdsUserAccountControlMapperArgs;
import com.pulumi.keycloak.ldap.inputs.MsadLdsUserAccountControlMapperState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Allows for creating and managing MSAD-LDS user account control mappers for Keycloak
 * users federated via LDAP.
 * 
 * The MSAD-LDS (Microsoft Active Directory Lightweight Directory Service) user account control mapper is specific
 * to LDAP user federation providers that are pulling from AD-LDS, and it can propagate
 * AD-LDS user state to Keycloak in order to enforce settings like expired passwords
 * or disabled accounts.
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
 * import com.pulumi.keycloak.ldap.UserFederation;
 * import com.pulumi.keycloak.ldap.UserFederationArgs;
 * import com.pulumi.keycloak.ldap.MsadLdsUserAccountControlMapper;
 * import com.pulumi.keycloak.ldap.MsadLdsUserAccountControlMapperArgs;
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
 *         var ldapUserFederation = new UserFederation(&#34;ldapUserFederation&#34;, UserFederationArgs.builder()        
 *             .realmId(realm.id())
 *             .usernameLdapAttribute(&#34;cn&#34;)
 *             .rdnLdapAttribute(&#34;cn&#34;)
 *             .uuidLdapAttribute(&#34;objectGUID&#34;)
 *             .userObjectClasses(            
 *                 &#34;person&#34;,
 *                 &#34;organizationalPerson&#34;,
 *                 &#34;user&#34;)
 *             .connectionUrl(&#34;ldap://my-ad-server&#34;)
 *             .usersDn(&#34;dc=example,dc=org&#34;)
 *             .bindDn(&#34;cn=admin,dc=example,dc=org&#34;)
 *             .bindCredential(&#34;admin&#34;)
 *             .build());
 * 
 *         var msadLdsUserAccountControlMapper = new MsadLdsUserAccountControlMapper(&#34;msadLdsUserAccountControlMapper&#34;, MsadLdsUserAccountControlMapperArgs.builder()        
 *             .realmId(realm.id())
 *             .ldapUserFederationId(ldapUserFederation.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * LDAP mappers can be imported using the format `{{realm_id}}/{{ldap_user_federation_id}}/{{ldap_mapper_id}}`. The ID of the LDAP user federation provider and the mapper can be found within the Keycloak GUI, and they are typically GUIDs. Examplebash
 * 
 * ```sh
 *  $ pulumi import keycloak:ldap/msadLdsUserAccountControlMapper:MsadLdsUserAccountControlMapper msad_lds_user_account_control_mapper my-realm/af2a6ca3-e4d7-49c3-b08b-1b3c70b4b860/3d923ece-1a91-4bf7-adaf-3b82f2a12b67
 * ```
 * 
 */
@ResourceType(type="keycloak:ldap/msadLdsUserAccountControlMapper:MsadLdsUserAccountControlMapper")
public class MsadLdsUserAccountControlMapper extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the LDAP user federation provider to attach this mapper to.
     * 
     */
    @Export(name="ldapUserFederationId", refs={String.class}, tree="[0]")
    private Output<String> ldapUserFederationId;

    /**
     * @return The ID of the LDAP user federation provider to attach this mapper to.
     * 
     */
    public Output<String> ldapUserFederationId() {
        return this.ldapUserFederationId;
    }
    /**
     * Display name of this mapper when displayed in the console.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Display name of this mapper when displayed in the console.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The realm that this LDAP mapper will exist in.
     * 
     */
    @Export(name="realmId", refs={String.class}, tree="[0]")
    private Output<String> realmId;

    /**
     * @return The realm that this LDAP mapper will exist in.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public MsadLdsUserAccountControlMapper(String name) {
        this(name, MsadLdsUserAccountControlMapperArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public MsadLdsUserAccountControlMapper(String name, MsadLdsUserAccountControlMapperArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public MsadLdsUserAccountControlMapper(String name, MsadLdsUserAccountControlMapperArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:ldap/msadLdsUserAccountControlMapper:MsadLdsUserAccountControlMapper", name, args == null ? MsadLdsUserAccountControlMapperArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private MsadLdsUserAccountControlMapper(String name, Output<String> id, @Nullable MsadLdsUserAccountControlMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:ldap/msadLdsUserAccountControlMapper:MsadLdsUserAccountControlMapper", name, state, makeResourceOptions(options, id));
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
    public static MsadLdsUserAccountControlMapper get(String name, Output<String> id, @Nullable MsadLdsUserAccountControlMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new MsadLdsUserAccountControlMapper(name, id, state, options);
    }
}
