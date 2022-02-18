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

/**
 * AutopilotConfiguration
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class AutopilotConfiguration {
  public static final String SERIALIZED_NAME_CLEANUP_DEAD_SERVERS = "CleanupDeadServers";
  @SerializedName(SERIALIZED_NAME_CLEANUP_DEAD_SERVERS)
  private Boolean cleanupDeadServers;

  public static final String SERIALIZED_NAME_CREATE_INDEX = "CreateIndex";
  @SerializedName(SERIALIZED_NAME_CREATE_INDEX)
  private Integer createIndex;

  public static final String SERIALIZED_NAME_DISABLE_UPGRADE_MIGRATION = "DisableUpgradeMigration";
  @SerializedName(SERIALIZED_NAME_DISABLE_UPGRADE_MIGRATION)
  private Boolean disableUpgradeMigration;

  public static final String SERIALIZED_NAME_ENABLE_CUSTOM_UPGRADES = "EnableCustomUpgrades";
  @SerializedName(SERIALIZED_NAME_ENABLE_CUSTOM_UPGRADES)
  private Boolean enableCustomUpgrades;

  public static final String SERIALIZED_NAME_ENABLE_REDUNDANCY_ZONES = "EnableRedundancyZones";
  @SerializedName(SERIALIZED_NAME_ENABLE_REDUNDANCY_ZONES)
  private Boolean enableRedundancyZones;

  public static final String SERIALIZED_NAME_LAST_CONTACT_THRESHOLD = "LastContactThreshold";
  @SerializedName(SERIALIZED_NAME_LAST_CONTACT_THRESHOLD)
  private String lastContactThreshold;

  public static final String SERIALIZED_NAME_MAX_TRAILING_LOGS = "MaxTrailingLogs";
  @SerializedName(SERIALIZED_NAME_MAX_TRAILING_LOGS)
  private Integer maxTrailingLogs;

  public static final String SERIALIZED_NAME_MIN_QUORUM = "MinQuorum";
  @SerializedName(SERIALIZED_NAME_MIN_QUORUM)
  private Integer minQuorum;

  public static final String SERIALIZED_NAME_MODIFY_INDEX = "ModifyIndex";
  @SerializedName(SERIALIZED_NAME_MODIFY_INDEX)
  private Integer modifyIndex;

  public static final String SERIALIZED_NAME_SERVER_STABILIZATION_TIME = "ServerStabilizationTime";
  @SerializedName(SERIALIZED_NAME_SERVER_STABILIZATION_TIME)
  private String serverStabilizationTime;


  public AutopilotConfiguration cleanupDeadServers(Boolean cleanupDeadServers) {
    
    this.cleanupDeadServers = cleanupDeadServers;
    return this;
  }

   /**
   * Get cleanupDeadServers
   * @return cleanupDeadServers
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getCleanupDeadServers() {
    return cleanupDeadServers;
  }


  public void setCleanupDeadServers(Boolean cleanupDeadServers) {
    this.cleanupDeadServers = cleanupDeadServers;
  }


  public AutopilotConfiguration createIndex(Integer createIndex) {
    
    this.createIndex = createIndex;
    return this;
  }

   /**
   * Get createIndex
   * minimum: 0
   * maximum: 384
   * @return createIndex
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getCreateIndex() {
    return createIndex;
  }


  public void setCreateIndex(Integer createIndex) {
    this.createIndex = createIndex;
  }


  public AutopilotConfiguration disableUpgradeMigration(Boolean disableUpgradeMigration) {
    
    this.disableUpgradeMigration = disableUpgradeMigration;
    return this;
  }

   /**
   * Get disableUpgradeMigration
   * @return disableUpgradeMigration
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getDisableUpgradeMigration() {
    return disableUpgradeMigration;
  }


  public void setDisableUpgradeMigration(Boolean disableUpgradeMigration) {
    this.disableUpgradeMigration = disableUpgradeMigration;
  }


  public AutopilotConfiguration enableCustomUpgrades(Boolean enableCustomUpgrades) {
    
    this.enableCustomUpgrades = enableCustomUpgrades;
    return this;
  }

   /**
   * Get enableCustomUpgrades
   * @return enableCustomUpgrades
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getEnableCustomUpgrades() {
    return enableCustomUpgrades;
  }


  public void setEnableCustomUpgrades(Boolean enableCustomUpgrades) {
    this.enableCustomUpgrades = enableCustomUpgrades;
  }


  public AutopilotConfiguration enableRedundancyZones(Boolean enableRedundancyZones) {
    
    this.enableRedundancyZones = enableRedundancyZones;
    return this;
  }

   /**
   * Get enableRedundancyZones
   * @return enableRedundancyZones
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getEnableRedundancyZones() {
    return enableRedundancyZones;
  }


  public void setEnableRedundancyZones(Boolean enableRedundancyZones) {
    this.enableRedundancyZones = enableRedundancyZones;
  }


  public AutopilotConfiguration lastContactThreshold(String lastContactThreshold) {
    
    this.lastContactThreshold = lastContactThreshold;
    return this;
  }

   /**
   * Get lastContactThreshold
   * @return lastContactThreshold
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getLastContactThreshold() {
    return lastContactThreshold;
  }


  public void setLastContactThreshold(String lastContactThreshold) {
    this.lastContactThreshold = lastContactThreshold;
  }


  public AutopilotConfiguration maxTrailingLogs(Integer maxTrailingLogs) {
    
    this.maxTrailingLogs = maxTrailingLogs;
    return this;
  }

   /**
   * Get maxTrailingLogs
   * minimum: 0
   * maximum: 384
   * @return maxTrailingLogs
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getMaxTrailingLogs() {
    return maxTrailingLogs;
  }


  public void setMaxTrailingLogs(Integer maxTrailingLogs) {
    this.maxTrailingLogs = maxTrailingLogs;
  }


  public AutopilotConfiguration minQuorum(Integer minQuorum) {
    
    this.minQuorum = minQuorum;
    return this;
  }

   /**
   * Get minQuorum
   * minimum: 0
   * @return minQuorum
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getMinQuorum() {
    return minQuorum;
  }


  public void setMinQuorum(Integer minQuorum) {
    this.minQuorum = minQuorum;
  }


  public AutopilotConfiguration modifyIndex(Integer modifyIndex) {
    
    this.modifyIndex = modifyIndex;
    return this;
  }

   /**
   * Get modifyIndex
   * minimum: 0
   * maximum: 384
   * @return modifyIndex
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getModifyIndex() {
    return modifyIndex;
  }


  public void setModifyIndex(Integer modifyIndex) {
    this.modifyIndex = modifyIndex;
  }


  public AutopilotConfiguration serverStabilizationTime(String serverStabilizationTime) {
    
    this.serverStabilizationTime = serverStabilizationTime;
    return this;
  }

   /**
   * Get serverStabilizationTime
   * @return serverStabilizationTime
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getServerStabilizationTime() {
    return serverStabilizationTime;
  }


  public void setServerStabilizationTime(String serverStabilizationTime) {
    this.serverStabilizationTime = serverStabilizationTime;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    AutopilotConfiguration autopilotConfiguration = (AutopilotConfiguration) o;
    return Objects.equals(this.cleanupDeadServers, autopilotConfiguration.cleanupDeadServers) &&
        Objects.equals(this.createIndex, autopilotConfiguration.createIndex) &&
        Objects.equals(this.disableUpgradeMigration, autopilotConfiguration.disableUpgradeMigration) &&
        Objects.equals(this.enableCustomUpgrades, autopilotConfiguration.enableCustomUpgrades) &&
        Objects.equals(this.enableRedundancyZones, autopilotConfiguration.enableRedundancyZones) &&
        Objects.equals(this.lastContactThreshold, autopilotConfiguration.lastContactThreshold) &&
        Objects.equals(this.maxTrailingLogs, autopilotConfiguration.maxTrailingLogs) &&
        Objects.equals(this.minQuorum, autopilotConfiguration.minQuorum) &&
        Objects.equals(this.modifyIndex, autopilotConfiguration.modifyIndex) &&
        Objects.equals(this.serverStabilizationTime, autopilotConfiguration.serverStabilizationTime);
  }

  @Override
  public int hashCode() {
    return Objects.hash(cleanupDeadServers, createIndex, disableUpgradeMigration, enableCustomUpgrades, enableRedundancyZones, lastContactThreshold, maxTrailingLogs, minQuorum, modifyIndex, serverStabilizationTime);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class AutopilotConfiguration {\n");
    sb.append("    cleanupDeadServers: ").append(toIndentedString(cleanupDeadServers)).append("\n");
    sb.append("    createIndex: ").append(toIndentedString(createIndex)).append("\n");
    sb.append("    disableUpgradeMigration: ").append(toIndentedString(disableUpgradeMigration)).append("\n");
    sb.append("    enableCustomUpgrades: ").append(toIndentedString(enableCustomUpgrades)).append("\n");
    sb.append("    enableRedundancyZones: ").append(toIndentedString(enableRedundancyZones)).append("\n");
    sb.append("    lastContactThreshold: ").append(toIndentedString(lastContactThreshold)).append("\n");
    sb.append("    maxTrailingLogs: ").append(toIndentedString(maxTrailingLogs)).append("\n");
    sb.append("    minQuorum: ").append(toIndentedString(minQuorum)).append("\n");
    sb.append("    modifyIndex: ").append(toIndentedString(modifyIndex)).append("\n");
    sb.append("    serverStabilizationTime: ").append(toIndentedString(serverStabilizationTime)).append("\n");
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
