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
 * NodeReservedNetworkResources
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class NodeReservedNetworkResources {
  public static final String SERIALIZED_NAME_RESERVED_HOST_PORTS = "ReservedHostPorts";
  @SerializedName(SERIALIZED_NAME_RESERVED_HOST_PORTS)
  private String reservedHostPorts;


  public NodeReservedNetworkResources reservedHostPorts(String reservedHostPorts) {
    
    this.reservedHostPorts = reservedHostPorts;
    return this;
  }

   /**
   * Get reservedHostPorts
   * @return reservedHostPorts
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getReservedHostPorts() {
    return reservedHostPorts;
  }


  public void setReservedHostPorts(String reservedHostPorts) {
    this.reservedHostPorts = reservedHostPorts;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    NodeReservedNetworkResources nodeReservedNetworkResources = (NodeReservedNetworkResources) o;
    return Objects.equals(this.reservedHostPorts, nodeReservedNetworkResources.reservedHostPorts);
  }

  @Override
  public int hashCode() {
    return Objects.hash(reservedHostPorts);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class NodeReservedNetworkResources {\n");
    sb.append("    reservedHostPorts: ").append(toIndentedString(reservedHostPorts)).append("\n");
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

