// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.RealmKeystoreJavaGeneratedArgs;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.inputs.RealmKeystoreJavaGeneratedState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows for creating and managing `java-keystore` Realm keystores within Keycloak.
 * 
 * A realm keystore manages generated key pairs that are used by Keycloak to perform cryptographic signatures and encryption.
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
 * import com.pulumi.keycloak.RealmKeystoreJavaGenerated;
 * import com.pulumi.keycloak.RealmKeystoreJavaGeneratedArgs;
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
 *             .build());
 * 
 *         var javaKeystore = new RealmKeystoreJavaGenerated(&#34;javaKeystore&#34;, RealmKeystoreJavaGeneratedArgs.builder()        
 *             .realmId(realm.id())
 *             .enabled(true)
 *             .active(true)
 *             .keystore(&#34;&lt;path to your keystore&gt;&#34;)
 *             .keystorePassword(&#34;&lt;password for keystore&gt;&#34;)
 *             .keyAlias(&#34;&lt;alias for the private key&gt;&#34;)
 *             .keyPassword(&#34;&lt;password for the private key&gt;&#34;)
 *             .priority(100)
 *             .algorithm(&#34;RS256&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Realm keys can be imported using realm name and keystore id, you can find it in web UI. Examplebash
 * 
 * ```sh
 *  $ pulumi import keycloak:index/realmKeystoreJavaGenerated:RealmKeystoreJavaGenerated java_keystore my-realm/618cfba7-49aa-4c09-9a19-2f699b576f0b
 * ```
 * 
 */
@ResourceType(type="keycloak:index/realmKeystoreJavaGenerated:RealmKeystoreJavaGenerated")
public class RealmKeystoreJavaGenerated extends com.pulumi.resources.CustomResource {
    /**
     * When `false`, key in not used for signing. Defaults to `true`.
     * 
     */
    @Export(name="active", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> active;

    /**
     * @return When `false`, key in not used for signing. Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> active() {
        return Codegen.optional(this.active);
    }
    /**
     * Intended algorithm for the key. Defaults to `RS256`
     * 
     */
    @Export(name="algorithm", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> algorithm;

    /**
     * @return Intended algorithm for the key. Defaults to `RS256`
     * 
     */
    public Output<Optional<String>> algorithm() {
        return Codegen.optional(this.algorithm);
    }
    /**
     * When `false`, key is not accessible in this realm. Defaults to `true`.
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return When `false`, key is not accessible in this realm. Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    /**
     * Alias for the private key.
     * 
     */
    @Export(name="keyAlias", refs={String.class}, tree="[0]")
    private Output<String> keyAlias;

    /**
     * @return Alias for the private key.
     * 
     */
    public Output<String> keyAlias() {
        return this.keyAlias;
    }
    /**
     * Password for the private key.
     * 
     */
    @Export(name="keyPassword", refs={String.class}, tree="[0]")
    private Output<String> keyPassword;

    /**
     * @return Password for the private key.
     * 
     */
    public Output<String> keyPassword() {
        return this.keyPassword;
    }
    /**
     * Path to keys file on keycloak instance.
     * 
     */
    @Export(name="keystore", refs={String.class}, tree="[0]")
    private Output<String> keystore;

    /**
     * @return Path to keys file on keycloak instance.
     * 
     */
    public Output<String> keystore() {
        return this.keystore;
    }
    /**
     * Password for the keys.
     * 
     */
    @Export(name="keystorePassword", refs={String.class}, tree="[0]")
    private Output<String> keystorePassword;

    /**
     * @return Password for the keys.
     * 
     */
    public Output<String> keystorePassword() {
        return this.keystorePassword;
    }
    /**
     * Display name of provider when linked in admin console.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Display name of provider when linked in admin console.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Priority for the provider. Defaults to `0`
     * 
     */
    @Export(name="priority", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> priority;

    /**
     * @return Priority for the provider. Defaults to `0`
     * 
     */
    public Output<Optional<Integer>> priority() {
        return Codegen.optional(this.priority);
    }
    /**
     * The realm this keystore exists in.
     * 
     */
    @Export(name="realmId", refs={String.class}, tree="[0]")
    private Output<String> realmId;

    /**
     * @return The realm this keystore exists in.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RealmKeystoreJavaGenerated(String name) {
        this(name, RealmKeystoreJavaGeneratedArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RealmKeystoreJavaGenerated(String name, RealmKeystoreJavaGeneratedArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RealmKeystoreJavaGenerated(String name, RealmKeystoreJavaGeneratedArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/realmKeystoreJavaGenerated:RealmKeystoreJavaGenerated", name, args == null ? RealmKeystoreJavaGeneratedArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RealmKeystoreJavaGenerated(String name, Output<String> id, @Nullable RealmKeystoreJavaGeneratedState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/realmKeystoreJavaGenerated:RealmKeystoreJavaGenerated", name, state, makeResourceOptions(options, id));
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
    public static RealmKeystoreJavaGenerated get(String name, Output<String> id, @Nullable RealmKeystoreJavaGeneratedState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RealmKeystoreJavaGenerated(name, id, state, options);
    }
}
