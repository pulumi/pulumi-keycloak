// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetRealmKeysKey {
    /**
     * @return Key algorithm (string)
     * 
     */
    private String algorithm;
    /**
     * @return Key certificate (string)
     * 
     */
    private String certificate;
    /**
     * @return Key ID (string)
     * 
     */
    private String kid;
    /**
     * @return Key provider ID (string)
     * 
     */
    private String providerId;
    /**
     * @return Key provider priority (int64)
     * 
     */
    private Integer providerPriority;
    /**
     * @return Key public key (string)
     * 
     */
    private String publicKey;
    /**
     * @return When specified, keys will be filtered by status. The statuses can be any of `ACTIVE`, `DISABLED` and `PASSIVE`.
     * 
     */
    private String status;
    /**
     * @return Key type (string)
     * 
     */
    private String type;

    private GetRealmKeysKey() {}
    /**
     * @return Key algorithm (string)
     * 
     */
    public String algorithm() {
        return this.algorithm;
    }
    /**
     * @return Key certificate (string)
     * 
     */
    public String certificate() {
        return this.certificate;
    }
    /**
     * @return Key ID (string)
     * 
     */
    public String kid() {
        return this.kid;
    }
    /**
     * @return Key provider ID (string)
     * 
     */
    public String providerId() {
        return this.providerId;
    }
    /**
     * @return Key provider priority (int64)
     * 
     */
    public Integer providerPriority() {
        return this.providerPriority;
    }
    /**
     * @return Key public key (string)
     * 
     */
    public String publicKey() {
        return this.publicKey;
    }
    /**
     * @return When specified, keys will be filtered by status. The statuses can be any of `ACTIVE`, `DISABLED` and `PASSIVE`.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return Key type (string)
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRealmKeysKey defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String algorithm;
        private String certificate;
        private String kid;
        private String providerId;
        private Integer providerPriority;
        private String publicKey;
        private String status;
        private String type;
        public Builder() {}
        public Builder(GetRealmKeysKey defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.algorithm = defaults.algorithm;
    	      this.certificate = defaults.certificate;
    	      this.kid = defaults.kid;
    	      this.providerId = defaults.providerId;
    	      this.providerPriority = defaults.providerPriority;
    	      this.publicKey = defaults.publicKey;
    	      this.status = defaults.status;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder algorithm(String algorithm) {
            if (algorithm == null) {
              throw new MissingRequiredPropertyException("GetRealmKeysKey", "algorithm");
            }
            this.algorithm = algorithm;
            return this;
        }
        @CustomType.Setter
        public Builder certificate(String certificate) {
            if (certificate == null) {
              throw new MissingRequiredPropertyException("GetRealmKeysKey", "certificate");
            }
            this.certificate = certificate;
            return this;
        }
        @CustomType.Setter
        public Builder kid(String kid) {
            if (kid == null) {
              throw new MissingRequiredPropertyException("GetRealmKeysKey", "kid");
            }
            this.kid = kid;
            return this;
        }
        @CustomType.Setter
        public Builder providerId(String providerId) {
            if (providerId == null) {
              throw new MissingRequiredPropertyException("GetRealmKeysKey", "providerId");
            }
            this.providerId = providerId;
            return this;
        }
        @CustomType.Setter
        public Builder providerPriority(Integer providerPriority) {
            if (providerPriority == null) {
              throw new MissingRequiredPropertyException("GetRealmKeysKey", "providerPriority");
            }
            this.providerPriority = providerPriority;
            return this;
        }
        @CustomType.Setter
        public Builder publicKey(String publicKey) {
            if (publicKey == null) {
              throw new MissingRequiredPropertyException("GetRealmKeysKey", "publicKey");
            }
            this.publicKey = publicKey;
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            if (status == null) {
              throw new MissingRequiredPropertyException("GetRealmKeysKey", "status");
            }
            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            if (type == null) {
              throw new MissingRequiredPropertyException("GetRealmKeysKey", "type");
            }
            this.type = type;
            return this;
        }
        public GetRealmKeysKey build() {
            final var _resultValue = new GetRealmKeysKey();
            _resultValue.algorithm = algorithm;
            _resultValue.certificate = certificate;
            _resultValue.kid = kid;
            _resultValue.providerId = providerId;
            _resultValue.providerPriority = providerPriority;
            _resultValue.publicKey = publicKey;
            _resultValue.status = status;
            _resultValue.type = type;
            return _resultValue;
        }
    }
}
