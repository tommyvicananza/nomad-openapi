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
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

/**
 * DeploymentAllocHealthRequest
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class DeploymentAllocHealthRequest {
  public static final String SERIALIZED_NAME_DEPLOYMENT_I_D = "DeploymentID";
  @SerializedName(SERIALIZED_NAME_DEPLOYMENT_I_D)
  private String deploymentID;

  public static final String SERIALIZED_NAME_HEALTHY_ALLOCATION_I_DS = "HealthyAllocationIDs";
  @SerializedName(SERIALIZED_NAME_HEALTHY_ALLOCATION_I_DS)
  private List<String> healthyAllocationIDs = null;

  public static final String SERIALIZED_NAME_NAMESPACE = "Namespace";
  @SerializedName(SERIALIZED_NAME_NAMESPACE)
  private String namespace;

  public static final String SERIALIZED_NAME_REGION = "Region";
  @SerializedName(SERIALIZED_NAME_REGION)
  private String region;

  public static final String SERIALIZED_NAME_SECRET_I_D = "SecretID";
  @SerializedName(SERIALIZED_NAME_SECRET_I_D)
  private String secretID;

  public static final String SERIALIZED_NAME_UNHEALTHY_ALLOCATION_I_DS = "UnhealthyAllocationIDs";
  @SerializedName(SERIALIZED_NAME_UNHEALTHY_ALLOCATION_I_DS)
  private List<String> unhealthyAllocationIDs = null;


  public DeploymentAllocHealthRequest deploymentID(String deploymentID) {
    
    this.deploymentID = deploymentID;
    return this;
  }

   /**
   * Get deploymentID
   * @return deploymentID
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDeploymentID() {
    return deploymentID;
  }


  public void setDeploymentID(String deploymentID) {
    this.deploymentID = deploymentID;
  }


  public DeploymentAllocHealthRequest healthyAllocationIDs(List<String> healthyAllocationIDs) {
    
    this.healthyAllocationIDs = healthyAllocationIDs;
    return this;
  }

  public DeploymentAllocHealthRequest addHealthyAllocationIDsItem(String healthyAllocationIDsItem) {
    if (this.healthyAllocationIDs == null) {
      this.healthyAllocationIDs = new ArrayList<String>();
    }
    this.healthyAllocationIDs.add(healthyAllocationIDsItem);
    return this;
  }

   /**
   * Get healthyAllocationIDs
   * @return healthyAllocationIDs
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<String> getHealthyAllocationIDs() {
    return healthyAllocationIDs;
  }


  public void setHealthyAllocationIDs(List<String> healthyAllocationIDs) {
    this.healthyAllocationIDs = healthyAllocationIDs;
  }


  public DeploymentAllocHealthRequest namespace(String namespace) {
    
    this.namespace = namespace;
    return this;
  }

   /**
   * Get namespace
   * @return namespace
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getNamespace() {
    return namespace;
  }


  public void setNamespace(String namespace) {
    this.namespace = namespace;
  }


  public DeploymentAllocHealthRequest region(String region) {
    
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


  public DeploymentAllocHealthRequest secretID(String secretID) {
    
    this.secretID = secretID;
    return this;
  }

   /**
   * Get secretID
   * @return secretID
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSecretID() {
    return secretID;
  }


  public void setSecretID(String secretID) {
    this.secretID = secretID;
  }


  public DeploymentAllocHealthRequest unhealthyAllocationIDs(List<String> unhealthyAllocationIDs) {
    
    this.unhealthyAllocationIDs = unhealthyAllocationIDs;
    return this;
  }

  public DeploymentAllocHealthRequest addUnhealthyAllocationIDsItem(String unhealthyAllocationIDsItem) {
    if (this.unhealthyAllocationIDs == null) {
      this.unhealthyAllocationIDs = new ArrayList<String>();
    }
    this.unhealthyAllocationIDs.add(unhealthyAllocationIDsItem);
    return this;
  }

   /**
   * Get unhealthyAllocationIDs
   * @return unhealthyAllocationIDs
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<String> getUnhealthyAllocationIDs() {
    return unhealthyAllocationIDs;
  }


  public void setUnhealthyAllocationIDs(List<String> unhealthyAllocationIDs) {
    this.unhealthyAllocationIDs = unhealthyAllocationIDs;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    DeploymentAllocHealthRequest deploymentAllocHealthRequest = (DeploymentAllocHealthRequest) o;
    return Objects.equals(this.deploymentID, deploymentAllocHealthRequest.deploymentID) &&
        Objects.equals(this.healthyAllocationIDs, deploymentAllocHealthRequest.healthyAllocationIDs) &&
        Objects.equals(this.namespace, deploymentAllocHealthRequest.namespace) &&
        Objects.equals(this.region, deploymentAllocHealthRequest.region) &&
        Objects.equals(this.secretID, deploymentAllocHealthRequest.secretID) &&
        Objects.equals(this.unhealthyAllocationIDs, deploymentAllocHealthRequest.unhealthyAllocationIDs);
  }

  @Override
  public int hashCode() {
    return Objects.hash(deploymentID, healthyAllocationIDs, namespace, region, secretID, unhealthyAllocationIDs);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class DeploymentAllocHealthRequest {\n");
    sb.append("    deploymentID: ").append(toIndentedString(deploymentID)).append("\n");
    sb.append("    healthyAllocationIDs: ").append(toIndentedString(healthyAllocationIDs)).append("\n");
    sb.append("    namespace: ").append(toIndentedString(namespace)).append("\n");
    sb.append("    region: ").append(toIndentedString(region)).append("\n");
    sb.append("    secretID: ").append(toIndentedString(secretID)).append("\n");
    sb.append("    unhealthyAllocationIDs: ").append(toIndentedString(unhealthyAllocationIDs)).append("\n");
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

