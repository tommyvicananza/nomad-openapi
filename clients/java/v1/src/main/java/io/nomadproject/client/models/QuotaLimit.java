/*
 * Nomad
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.1.4
 * Contact: support@hashicorp.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package io.nomadproject.client.models;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.nomadproject.client.models.Resources;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;

/**
 * QuotaLimit
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class QuotaLimit {
  public static final String SERIALIZED_NAME_HASH = "Hash";
  @SerializedName(SERIALIZED_NAME_HASH)
  private byte[] hash;

  public static final String SERIALIZED_NAME_REGION = "Region";
  @SerializedName(SERIALIZED_NAME_REGION)
  private String region;

  public static final String SERIALIZED_NAME_REGION_LIMIT = "RegionLimit";
  @SerializedName(SERIALIZED_NAME_REGION_LIMIT)
  private Resources regionLimit;

  public static final String SERIALIZED_NAME_VARIABLES_LIMIT = "VariablesLimit";
  @SerializedName(SERIALIZED_NAME_VARIABLES_LIMIT)
  private Integer variablesLimit;


  public QuotaLimit hash(byte[] hash) {
    
    this.hash = hash;
    return this;
  }

   /**
   * Get hash
   * @return hash
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public byte[] getHash() {
    return hash;
  }


  public void setHash(byte[] hash) {
    this.hash = hash;
  }


  public QuotaLimit region(String region) {
    
    this.region = region;
    return this;
  }

   /**
   * Get region
   * @return region
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRegion() {
    return region;
  }


  public void setRegion(String region) {
    this.region = region;
  }


  public QuotaLimit regionLimit(Resources regionLimit) {
    
    this.regionLimit = regionLimit;
    return this;
  }

   /**
   * Get regionLimit
   * @return regionLimit
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Resources getRegionLimit() {
    return regionLimit;
  }


  public void setRegionLimit(Resources regionLimit) {
    this.regionLimit = regionLimit;
  }


  public QuotaLimit variablesLimit(Integer variablesLimit) {
    
    this.variablesLimit = variablesLimit;
    return this;
  }

   /**
   * Get variablesLimit
   * @return variablesLimit
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getVariablesLimit() {
    return variablesLimit;
  }


  public void setVariablesLimit(Integer variablesLimit) {
    this.variablesLimit = variablesLimit;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    QuotaLimit quotaLimit = (QuotaLimit) o;
    return Arrays.equals(this.hash, quotaLimit.hash) &&
        Objects.equals(this.region, quotaLimit.region) &&
        Objects.equals(this.regionLimit, quotaLimit.regionLimit) &&
        Objects.equals(this.variablesLimit, quotaLimit.variablesLimit);
  }

  @Override
  public int hashCode() {
    return Objects.hash(Arrays.hashCode(hash), region, regionLimit, variablesLimit);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class QuotaLimit {\n");
    sb.append("    hash: ").append(toIndentedString(hash)).append("\n");
    sb.append("    region: ").append(toIndentedString(region)).append("\n");
    sb.append("    regionLimit: ").append(toIndentedString(regionLimit)).append("\n");
    sb.append("    variablesLimit: ").append(toIndentedString(variablesLimit)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

