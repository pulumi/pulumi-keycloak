// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class RealmInternationalization {
    private String defaultLocale;
    private List<String> supportedLocales;

    private RealmInternationalization() {}
    public String defaultLocale() {
        return this.defaultLocale;
    }
    public List<String> supportedLocales() {
        return this.supportedLocales;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RealmInternationalization defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String defaultLocale;
        private List<String> supportedLocales;
        public Builder() {}
        public Builder(RealmInternationalization defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.defaultLocale = defaults.defaultLocale;
    	      this.supportedLocales = defaults.supportedLocales;
        }

        @CustomType.Setter
        public Builder defaultLocale(String defaultLocale) {
            if (defaultLocale == null) {
              throw new MissingRequiredPropertyException("RealmInternationalization", "defaultLocale");
            }
            this.defaultLocale = defaultLocale;
            return this;
        }
        @CustomType.Setter
        public Builder supportedLocales(List<String> supportedLocales) {
            if (supportedLocales == null) {
              throw new MissingRequiredPropertyException("RealmInternationalization", "supportedLocales");
            }
            this.supportedLocales = supportedLocales;
            return this;
        }
        public Builder supportedLocales(String... supportedLocales) {
            return supportedLocales(List.of(supportedLocales));
        }
        public RealmInternationalization build() {
            final var _resultValue = new RealmInternationalization();
            _resultValue.defaultLocale = defaultLocale;
            _resultValue.supportedLocales = supportedLocales;
            return _resultValue;
        }
    }
}
