// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.RealmUserProfileArgs;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.inputs.RealmUserProfileState;
import com.pulumi.keycloak.outputs.RealmUserProfileAttribute;
import com.pulumi.keycloak.outputs.RealmUserProfileGroup;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows for managing Realm User Profiles within Keycloak.
 * 
 * A user profile defines a schema for representing user attributes and how they are managed within a realm.
 * This is a preview feature, hence not fully supported and disabled by default.
 * To enable it, start the server with one of the following flags:
 * - WildFly distribution: `-Dkeycloak.profile.feature.declarative_user_profile=enabled`
 * - Quarkus distribution: `--features=preview` or `--features=declarative-user-profile`
 * 
 * The realm linked to the `keycloak.RealmUserProfile` resource must have the user profile feature enabled.
 * It can be done via the administration UI, or by setting the `userProfileEnabled` realm attribute to `true`.
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
 * import com.pulumi.keycloak.RealmUserProfile;
 * import com.pulumi.keycloak.RealmUserProfileArgs;
 * import com.pulumi.keycloak.inputs.RealmUserProfileAttributeArgs;
 * import com.pulumi.keycloak.inputs.RealmUserProfileAttributePermissionsArgs;
 * import com.pulumi.keycloak.inputs.RealmUserProfileGroupArgs;
 * import static com.pulumi.codegen.internal.Serialization.*;
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
 *             .attributes(Map.of(&#34;userProfileEnabled&#34;, true))
 *             .build());
 * 
 *         var userprofile = new RealmUserProfile(&#34;userprofile&#34;, RealmUserProfileArgs.builder()        
 *             .realmId(keycloak_realm.my_realm().id())
 *             .attributes(            
 *                 RealmUserProfileAttributeArgs.builder()
 *                     .name(&#34;field1&#34;)
 *                     .displayName(&#34;Field 1&#34;)
 *                     .group(&#34;group1&#34;)
 *                     .enabledWhenScopes(&#34;offline_access&#34;)
 *                     .requiredForRoles(&#34;user&#34;)
 *                     .requiredForScopes(&#34;offline_access&#34;)
 *                     .permissions(RealmUserProfileAttributePermissionsArgs.builder()
 *                         .views(                        
 *                             &#34;admin&#34;,
 *                             &#34;user&#34;)
 *                         .edits(                        
 *                             &#34;admin&#34;,
 *                             &#34;user&#34;)
 *                         .build())
 *                     .validators(                    
 *                         RealmUserProfileAttributeValidatorArgs.builder()
 *                             .name(&#34;person-name-prohibited-characters&#34;)
 *                             .build(),
 *                         RealmUserProfileAttributeValidatorArgs.builder()
 *                             .name(&#34;pattern&#34;)
 *                             .config(Map.ofEntries(
 *                                 Map.entry(&#34;pattern&#34;, &#34;^[a-z]+$&#34;),
 *                                 Map.entry(&#34;error-message&#34;, &#34;Nope&#34;)
 *                             ))
 *                             .build())
 *                     .annotations(Map.of(&#34;foo&#34;, &#34;bar&#34;))
 *                     .build(),
 *                 RealmUserProfileAttributeArgs.builder()
 *                     .name(&#34;field2&#34;)
 *                     .validators(RealmUserProfileAttributeValidatorArgs.builder()
 *                         .name(&#34;options&#34;)
 *                         .config(Map.of(&#34;options&#34;, serializeJson(
 *                             jsonArray(&#34;opt1&#34;))))
 *                         .build())
 *                     .annotations(Map.of(&#34;foo&#34;, serializeJson(
 *                         jsonObject(
 *                             jsonProperty(&#34;key&#34;, &#34;val&#34;)
 *                         ))))
 *                     .build())
 *             .groups(            
 *                 RealmUserProfileGroupArgs.builder()
 *                     .name(&#34;group1&#34;)
 *                     .displayHeader(&#34;Group 1&#34;)
 *                     .displayDescription(&#34;A first group&#34;)
 *                     .annotations(Map.ofEntries(
 *                         Map.entry(&#34;foo&#34;, &#34;bar&#34;),
 *                         Map.entry(&#34;foo2&#34;, serializeJson(
 *                             jsonObject(
 *                                 jsonProperty(&#34;key&#34;, &#34;val&#34;)
 *                             )))
 *                     ))
 *                     .build(),
 *                 RealmUserProfileGroupArgs.builder()
 *                     .name(&#34;group2&#34;)
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * This resource currently does not support importing.
 * 
 */
@ResourceType(type="keycloak:index/realmUserProfile:RealmUserProfile")
public class RealmUserProfile extends com.pulumi.resources.CustomResource {
    /**
     * An ordered list of attributes.
     * 
     */
    @Export(name="attributes", refs={List.class,RealmUserProfileAttribute.class}, tree="[0,1]")
    private Output</* @Nullable */ List<RealmUserProfileAttribute>> attributes;

    /**
     * @return An ordered list of attributes.
     * 
     */
    public Output<Optional<List<RealmUserProfileAttribute>>> attributes() {
        return Codegen.optional(this.attributes);
    }
    /**
     * A list of groups.
     * 
     */
    @Export(name="groups", refs={List.class,RealmUserProfileGroup.class}, tree="[0,1]")
    private Output</* @Nullable */ List<RealmUserProfileGroup>> groups;

    /**
     * @return A list of groups.
     * 
     */
    public Output<Optional<List<RealmUserProfileGroup>>> groups() {
        return Codegen.optional(this.groups);
    }
    /**
     * The ID of the realm the user profile applies to.
     * 
     */
    @Export(name="realmId", refs={String.class}, tree="[0]")
    private Output<String> realmId;

    /**
     * @return The ID of the realm the user profile applies to.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RealmUserProfile(String name) {
        this(name, RealmUserProfileArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RealmUserProfile(String name, RealmUserProfileArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RealmUserProfile(String name, RealmUserProfileArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/realmUserProfile:RealmUserProfile", name, args == null ? RealmUserProfileArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RealmUserProfile(String name, Output<String> id, @Nullable RealmUserProfileState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/realmUserProfile:RealmUserProfile", name, state, makeResourceOptions(options, id));
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
    public static RealmUserProfile get(String name, Output<String> id, @Nullable RealmUserProfileState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RealmUserProfile(name, id, state, options);
    }
}
