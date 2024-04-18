// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.saml;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.saml.inputs.GetClientArgs;
import com.pulumi.keycloak.saml.inputs.GetClientInstallationProviderArgs;
import com.pulumi.keycloak.saml.inputs.GetClientInstallationProviderPlainArgs;
import com.pulumi.keycloak.saml.inputs.GetClientPlainArgs;
import com.pulumi.keycloak.saml.outputs.GetClientInstallationProviderResult;
import com.pulumi.keycloak.saml.outputs.GetClientResult;
import java.util.concurrent.CompletableFuture;

public final class SamlFunctions {
    /**
     * This data source can be used to fetch properties of a Keycloak client that uses the SAML protocol.
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
     * import com.pulumi.keycloak.saml.SamlFunctions;
     * import com.pulumi.keycloak.saml.inputs.GetClientArgs;
     * import com.pulumi.keycloak.KeycloakFunctions;
     * import com.pulumi.keycloak.inputs.GetRoleArgs;
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
     *         final var realmManagement = SamlFunctions.getClient(GetClientArgs.builder()
     *             .realmId(&#34;my-realm&#34;)
     *             .clientId(&#34;realm-management&#34;)
     *             .build());
     * 
     *         // use the data source
     *         final var admin = KeycloakFunctions.getRole(GetRoleArgs.builder()
     *             .realmId(&#34;my-realm&#34;)
     *             .clientId(realmManagement.applyValue(getClientResult -&gt; getClientResult.id()))
     *             .name(&#34;realm-admin&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetClientResult> getClient(GetClientArgs args) {
        return getClient(args, InvokeOptions.Empty);
    }
    /**
     * This data source can be used to fetch properties of a Keycloak client that uses the SAML protocol.
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
     * import com.pulumi.keycloak.saml.SamlFunctions;
     * import com.pulumi.keycloak.saml.inputs.GetClientArgs;
     * import com.pulumi.keycloak.KeycloakFunctions;
     * import com.pulumi.keycloak.inputs.GetRoleArgs;
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
     *         final var realmManagement = SamlFunctions.getClient(GetClientArgs.builder()
     *             .realmId(&#34;my-realm&#34;)
     *             .clientId(&#34;realm-management&#34;)
     *             .build());
     * 
     *         // use the data source
     *         final var admin = KeycloakFunctions.getRole(GetRoleArgs.builder()
     *             .realmId(&#34;my-realm&#34;)
     *             .clientId(realmManagement.applyValue(getClientResult -&gt; getClientResult.id()))
     *             .name(&#34;realm-admin&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetClientResult> getClientPlain(GetClientPlainArgs args) {
        return getClientPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source can be used to fetch properties of a Keycloak client that uses the SAML protocol.
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
     * import com.pulumi.keycloak.saml.SamlFunctions;
     * import com.pulumi.keycloak.saml.inputs.GetClientArgs;
     * import com.pulumi.keycloak.KeycloakFunctions;
     * import com.pulumi.keycloak.inputs.GetRoleArgs;
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
     *         final var realmManagement = SamlFunctions.getClient(GetClientArgs.builder()
     *             .realmId(&#34;my-realm&#34;)
     *             .clientId(&#34;realm-management&#34;)
     *             .build());
     * 
     *         // use the data source
     *         final var admin = KeycloakFunctions.getRole(GetRoleArgs.builder()
     *             .realmId(&#34;my-realm&#34;)
     *             .clientId(realmManagement.applyValue(getClientResult -&gt; getClientResult.id()))
     *             .name(&#34;realm-admin&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetClientResult> getClient(GetClientArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("keycloak:saml/getClient:getClient", TypeShape.of(GetClientResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source can be used to fetch properties of a Keycloak client that uses the SAML protocol.
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
     * import com.pulumi.keycloak.saml.SamlFunctions;
     * import com.pulumi.keycloak.saml.inputs.GetClientArgs;
     * import com.pulumi.keycloak.KeycloakFunctions;
     * import com.pulumi.keycloak.inputs.GetRoleArgs;
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
     *         final var realmManagement = SamlFunctions.getClient(GetClientArgs.builder()
     *             .realmId(&#34;my-realm&#34;)
     *             .clientId(&#34;realm-management&#34;)
     *             .build());
     * 
     *         // use the data source
     *         final var admin = KeycloakFunctions.getRole(GetRoleArgs.builder()
     *             .realmId(&#34;my-realm&#34;)
     *             .clientId(realmManagement.applyValue(getClientResult -&gt; getClientResult.id()))
     *             .name(&#34;realm-admin&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetClientResult> getClientPlain(GetClientPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("keycloak:saml/getClient:getClient", TypeShape.of(GetClientResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source can be used to retrieve Installation Provider of a SAML Client.
     * 
     * ## Example Usage
     * 
     * In the example below, we extract the SAML metadata IDPSSODescriptor to pass it to the AWS IAM SAML Provider.
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
     * import com.pulumi.keycloak.saml.Client;
     * import com.pulumi.keycloak.saml.ClientArgs;
     * import com.pulumi.keycloak.saml.SamlFunctions;
     * import com.pulumi.keycloak.saml.inputs.GetClientInstallationProviderArgs;
     * import com.pulumi.aws.iamSamlProvider;
     * import com.pulumi.aws.IamSamlProviderArgs;
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
     *         var samlClient = new Client(&#34;samlClient&#34;, ClientArgs.builder()        
     *             .realmId(realm.id())
     *             .clientId(&#34;test-saml-client&#34;)
     *             .name(&#34;test-saml-client&#34;)
     *             .signDocuments(false)
     *             .signAssertions(true)
     *             .includeAuthnStatement(true)
     *             .signingCertificate(StdFunctions.file(FileArgs.builder()
     *                 .input(&#34;saml-cert.pem&#34;)
     *                 .build()).result())
     *             .signingPrivateKey(StdFunctions.file(FileArgs.builder()
     *                 .input(&#34;saml-key.pem&#34;)
     *                 .build()).result())
     *             .build());
     * 
     *         final var samlIdpDescriptor = SamlFunctions.getClientInstallationProvider(GetClientInstallationProviderArgs.builder()
     *             .realmId(realm.id())
     *             .clientId(samlClient.id())
     *             .providerId(&#34;saml-idp-descriptor&#34;)
     *             .build());
     * 
     *         var default_ = new IamSamlProvider(&#34;default&#34;, IamSamlProviderArgs.builder()        
     *             .name(&#34;myprovider&#34;)
     *             .samlMetadataDocument(samlIdpDescriptor.applyValue(getClientInstallationProviderResult -&gt; getClientInstallationProviderResult).applyValue(samlIdpDescriptor -&gt; samlIdpDescriptor.applyValue(getClientInstallationProviderResult -&gt; getClientInstallationProviderResult.value())))
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetClientInstallationProviderResult> getClientInstallationProvider(GetClientInstallationProviderArgs args) {
        return getClientInstallationProvider(args, InvokeOptions.Empty);
    }
    /**
     * This data source can be used to retrieve Installation Provider of a SAML Client.
     * 
     * ## Example Usage
     * 
     * In the example below, we extract the SAML metadata IDPSSODescriptor to pass it to the AWS IAM SAML Provider.
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
     * import com.pulumi.keycloak.saml.Client;
     * import com.pulumi.keycloak.saml.ClientArgs;
     * import com.pulumi.keycloak.saml.SamlFunctions;
     * import com.pulumi.keycloak.saml.inputs.GetClientInstallationProviderArgs;
     * import com.pulumi.aws.iamSamlProvider;
     * import com.pulumi.aws.IamSamlProviderArgs;
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
     *         var samlClient = new Client(&#34;samlClient&#34;, ClientArgs.builder()        
     *             .realmId(realm.id())
     *             .clientId(&#34;test-saml-client&#34;)
     *             .name(&#34;test-saml-client&#34;)
     *             .signDocuments(false)
     *             .signAssertions(true)
     *             .includeAuthnStatement(true)
     *             .signingCertificate(StdFunctions.file(FileArgs.builder()
     *                 .input(&#34;saml-cert.pem&#34;)
     *                 .build()).result())
     *             .signingPrivateKey(StdFunctions.file(FileArgs.builder()
     *                 .input(&#34;saml-key.pem&#34;)
     *                 .build()).result())
     *             .build());
     * 
     *         final var samlIdpDescriptor = SamlFunctions.getClientInstallationProvider(GetClientInstallationProviderArgs.builder()
     *             .realmId(realm.id())
     *             .clientId(samlClient.id())
     *             .providerId(&#34;saml-idp-descriptor&#34;)
     *             .build());
     * 
     *         var default_ = new IamSamlProvider(&#34;default&#34;, IamSamlProviderArgs.builder()        
     *             .name(&#34;myprovider&#34;)
     *             .samlMetadataDocument(samlIdpDescriptor.applyValue(getClientInstallationProviderResult -&gt; getClientInstallationProviderResult).applyValue(samlIdpDescriptor -&gt; samlIdpDescriptor.applyValue(getClientInstallationProviderResult -&gt; getClientInstallationProviderResult.value())))
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetClientInstallationProviderResult> getClientInstallationProviderPlain(GetClientInstallationProviderPlainArgs args) {
        return getClientInstallationProviderPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source can be used to retrieve Installation Provider of a SAML Client.
     * 
     * ## Example Usage
     * 
     * In the example below, we extract the SAML metadata IDPSSODescriptor to pass it to the AWS IAM SAML Provider.
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
     * import com.pulumi.keycloak.saml.Client;
     * import com.pulumi.keycloak.saml.ClientArgs;
     * import com.pulumi.keycloak.saml.SamlFunctions;
     * import com.pulumi.keycloak.saml.inputs.GetClientInstallationProviderArgs;
     * import com.pulumi.aws.iamSamlProvider;
     * import com.pulumi.aws.IamSamlProviderArgs;
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
     *         var samlClient = new Client(&#34;samlClient&#34;, ClientArgs.builder()        
     *             .realmId(realm.id())
     *             .clientId(&#34;test-saml-client&#34;)
     *             .name(&#34;test-saml-client&#34;)
     *             .signDocuments(false)
     *             .signAssertions(true)
     *             .includeAuthnStatement(true)
     *             .signingCertificate(StdFunctions.file(FileArgs.builder()
     *                 .input(&#34;saml-cert.pem&#34;)
     *                 .build()).result())
     *             .signingPrivateKey(StdFunctions.file(FileArgs.builder()
     *                 .input(&#34;saml-key.pem&#34;)
     *                 .build()).result())
     *             .build());
     * 
     *         final var samlIdpDescriptor = SamlFunctions.getClientInstallationProvider(GetClientInstallationProviderArgs.builder()
     *             .realmId(realm.id())
     *             .clientId(samlClient.id())
     *             .providerId(&#34;saml-idp-descriptor&#34;)
     *             .build());
     * 
     *         var default_ = new IamSamlProvider(&#34;default&#34;, IamSamlProviderArgs.builder()        
     *             .name(&#34;myprovider&#34;)
     *             .samlMetadataDocument(samlIdpDescriptor.applyValue(getClientInstallationProviderResult -&gt; getClientInstallationProviderResult).applyValue(samlIdpDescriptor -&gt; samlIdpDescriptor.applyValue(getClientInstallationProviderResult -&gt; getClientInstallationProviderResult.value())))
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetClientInstallationProviderResult> getClientInstallationProvider(GetClientInstallationProviderArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("keycloak:saml/getClientInstallationProvider:getClientInstallationProvider", TypeShape.of(GetClientInstallationProviderResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source can be used to retrieve Installation Provider of a SAML Client.
     * 
     * ## Example Usage
     * 
     * In the example below, we extract the SAML metadata IDPSSODescriptor to pass it to the AWS IAM SAML Provider.
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
     * import com.pulumi.keycloak.saml.Client;
     * import com.pulumi.keycloak.saml.ClientArgs;
     * import com.pulumi.keycloak.saml.SamlFunctions;
     * import com.pulumi.keycloak.saml.inputs.GetClientInstallationProviderArgs;
     * import com.pulumi.aws.iamSamlProvider;
     * import com.pulumi.aws.IamSamlProviderArgs;
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
     *         var samlClient = new Client(&#34;samlClient&#34;, ClientArgs.builder()        
     *             .realmId(realm.id())
     *             .clientId(&#34;test-saml-client&#34;)
     *             .name(&#34;test-saml-client&#34;)
     *             .signDocuments(false)
     *             .signAssertions(true)
     *             .includeAuthnStatement(true)
     *             .signingCertificate(StdFunctions.file(FileArgs.builder()
     *                 .input(&#34;saml-cert.pem&#34;)
     *                 .build()).result())
     *             .signingPrivateKey(StdFunctions.file(FileArgs.builder()
     *                 .input(&#34;saml-key.pem&#34;)
     *                 .build()).result())
     *             .build());
     * 
     *         final var samlIdpDescriptor = SamlFunctions.getClientInstallationProvider(GetClientInstallationProviderArgs.builder()
     *             .realmId(realm.id())
     *             .clientId(samlClient.id())
     *             .providerId(&#34;saml-idp-descriptor&#34;)
     *             .build());
     * 
     *         var default_ = new IamSamlProvider(&#34;default&#34;, IamSamlProviderArgs.builder()        
     *             .name(&#34;myprovider&#34;)
     *             .samlMetadataDocument(samlIdpDescriptor.applyValue(getClientInstallationProviderResult -&gt; getClientInstallationProviderResult).applyValue(samlIdpDescriptor -&gt; samlIdpDescriptor.applyValue(getClientInstallationProviderResult -&gt; getClientInstallationProviderResult.value())))
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetClientInstallationProviderResult> getClientInstallationProviderPlain(GetClientInstallationProviderPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("keycloak:saml/getClientInstallationProvider:getClientInstallationProvider", TypeShape.of(GetClientInstallationProviderResult.class), args, Utilities.withVersion(options));
    }
}
