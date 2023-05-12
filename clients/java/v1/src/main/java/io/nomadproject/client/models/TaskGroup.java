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
import io.nomadproject.client.models.Affinity;
import io.nomadproject.client.models.Constraint;
import io.nomadproject.client.models.Consul;
import io.nomadproject.client.models.EphemeralDisk;
import io.nomadproject.client.models.MigrateStrategy;
import io.nomadproject.client.models.NetworkResource;
import io.nomadproject.client.models.ReschedulePolicy;
import io.nomadproject.client.models.RestartPolicy;
import io.nomadproject.client.models.ScalingPolicy;
import io.nomadproject.client.models.Service;
import io.nomadproject.client.models.Spread;
import io.nomadproject.client.models.Task;
import io.nomadproject.client.models.UpdateStrategy;
import io.nomadproject.client.models.VolumeRequest;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * TaskGroup
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class TaskGroup {
  public static final String SERIALIZED_NAME_AFFINITIES = "Affinities";
  @SerializedName(SERIALIZED_NAME_AFFINITIES)
  private List<Affinity> affinities = null;

  public static final String SERIALIZED_NAME_CONSTRAINTS = "Constraints";
  @SerializedName(SERIALIZED_NAME_CONSTRAINTS)
  private List<Constraint> constraints = null;

  public static final String SERIALIZED_NAME_CONSUL = "Consul";
  @SerializedName(SERIALIZED_NAME_CONSUL)
  private Consul consul;

  public static final String SERIALIZED_NAME_COUNT = "Count";
  @SerializedName(SERIALIZED_NAME_COUNT)
  private Integer count;

  public static final String SERIALIZED_NAME_EPHEMERAL_DISK = "EphemeralDisk";
  @SerializedName(SERIALIZED_NAME_EPHEMERAL_DISK)
  private EphemeralDisk ephemeralDisk;

  public static final String SERIALIZED_NAME_MAX_CLIENT_DISCONNECT = "MaxClientDisconnect";
  @SerializedName(SERIALIZED_NAME_MAX_CLIENT_DISCONNECT)
  private Long maxClientDisconnect;

  public static final String SERIALIZED_NAME_META = "Meta";
  @SerializedName(SERIALIZED_NAME_META)
  private Map<String, String> meta = null;

  public static final String SERIALIZED_NAME_MIGRATE = "Migrate";
  @SerializedName(SERIALIZED_NAME_MIGRATE)
  private MigrateStrategy migrate;

  public static final String SERIALIZED_NAME_NAME = "Name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_NETWORKS = "Networks";
  @SerializedName(SERIALIZED_NAME_NETWORKS)
  private List<NetworkResource> networks = null;

  public static final String SERIALIZED_NAME_RESCHEDULE_POLICY = "ReschedulePolicy";
  @SerializedName(SERIALIZED_NAME_RESCHEDULE_POLICY)
  private ReschedulePolicy reschedulePolicy;

  public static final String SERIALIZED_NAME_RESTART_POLICY = "RestartPolicy";
  @SerializedName(SERIALIZED_NAME_RESTART_POLICY)
  private RestartPolicy restartPolicy;

  public static final String SERIALIZED_NAME_SCALING = "Scaling";
  @SerializedName(SERIALIZED_NAME_SCALING)
  private ScalingPolicy scaling;

  public static final String SERIALIZED_NAME_SERVICES = "Services";
  @SerializedName(SERIALIZED_NAME_SERVICES)
  private List<Service> services = null;

  public static final String SERIALIZED_NAME_SHUTDOWN_DELAY = "ShutdownDelay";
  @SerializedName(SERIALIZED_NAME_SHUTDOWN_DELAY)
  private Long shutdownDelay;

  public static final String SERIALIZED_NAME_SPREADS = "Spreads";
  @SerializedName(SERIALIZED_NAME_SPREADS)
  private List<Spread> spreads = null;

  public static final String SERIALIZED_NAME_STOP_AFTER_CLIENT_DISCONNECT = "StopAfterClientDisconnect";
  @SerializedName(SERIALIZED_NAME_STOP_AFTER_CLIENT_DISCONNECT)
  private Long stopAfterClientDisconnect;

  public static final String SERIALIZED_NAME_TASKS = "Tasks";
  @SerializedName(SERIALIZED_NAME_TASKS)
  private List<Task> tasks = null;

  public static final String SERIALIZED_NAME_UPDATE = "Update";
  @SerializedName(SERIALIZED_NAME_UPDATE)
  private UpdateStrategy update;

  public static final String SERIALIZED_NAME_VOLUMES = "Volumes";
  @SerializedName(SERIALIZED_NAME_VOLUMES)
  private Map<String, VolumeRequest> volumes = null;


  public TaskGroup affinities(List<Affinity> affinities) {
    
    this.affinities = affinities;
    return this;
  }

  public TaskGroup addAffinitiesItem(Affinity affinitiesItem) {
    if (this.affinities == null) {
      this.affinities = new ArrayList<Affinity>();
    }
    this.affinities.add(affinitiesItem);
    return this;
  }

   /**
   * Get affinities
   * @return affinities
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Affinity> getAffinities() {
    return affinities;
  }


  public void setAffinities(List<Affinity> affinities) {
    this.affinities = affinities;
  }


  public TaskGroup constraints(List<Constraint> constraints) {
    
    this.constraints = constraints;
    return this;
  }

  public TaskGroup addConstraintsItem(Constraint constraintsItem) {
    if (this.constraints == null) {
      this.constraints = new ArrayList<Constraint>();
    }
    this.constraints.add(constraintsItem);
    return this;
  }

   /**
   * Get constraints
   * @return constraints
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Constraint> getConstraints() {
    return constraints;
  }


  public void setConstraints(List<Constraint> constraints) {
    this.constraints = constraints;
  }


  public TaskGroup consul(Consul consul) {
    
    this.consul = consul;
    return this;
  }

   /**
   * Get consul
   * @return consul
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Consul getConsul() {
    return consul;
  }


  public void setConsul(Consul consul) {
    this.consul = consul;
  }


  public TaskGroup count(Integer count) {
    
    this.count = count;
    return this;
  }

   /**
   * Get count
   * @return count
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getCount() {
    return count;
  }


  public void setCount(Integer count) {
    this.count = count;
  }


  public TaskGroup ephemeralDisk(EphemeralDisk ephemeralDisk) {
    
    this.ephemeralDisk = ephemeralDisk;
    return this;
  }

   /**
   * Get ephemeralDisk
   * @return ephemeralDisk
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public EphemeralDisk getEphemeralDisk() {
    return ephemeralDisk;
  }


  public void setEphemeralDisk(EphemeralDisk ephemeralDisk) {
    this.ephemeralDisk = ephemeralDisk;
  }


  public TaskGroup maxClientDisconnect(Long maxClientDisconnect) {
    
    this.maxClientDisconnect = maxClientDisconnect;
    return this;
  }

   /**
   * Get maxClientDisconnect
   * @return maxClientDisconnect
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getMaxClientDisconnect() {
    return maxClientDisconnect;
  }


  public void setMaxClientDisconnect(Long maxClientDisconnect) {
    this.maxClientDisconnect = maxClientDisconnect;
  }


  public TaskGroup meta(Map<String, String> meta) {
    
    this.meta = meta;
    return this;
  }

  public TaskGroup putMetaItem(String key, String metaItem) {
    if (this.meta == null) {
      this.meta = new HashMap<String, String>();
    }
    this.meta.put(key, metaItem);
    return this;
  }

   /**
   * Get meta
   * @return meta
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Map<String, String> getMeta() {
    return meta;
  }


  public void setMeta(Map<String, String> meta) {
    this.meta = meta;
  }


  public TaskGroup migrate(MigrateStrategy migrate) {
    
    this.migrate = migrate;
    return this;
  }

   /**
   * Get migrate
   * @return migrate
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public MigrateStrategy getMigrate() {
    return migrate;
  }


  public void setMigrate(MigrateStrategy migrate) {
    this.migrate = migrate;
  }


  public TaskGroup name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Get name
   * @return name
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public TaskGroup networks(List<NetworkResource> networks) {
    
    this.networks = networks;
    return this;
  }

  public TaskGroup addNetworksItem(NetworkResource networksItem) {
    if (this.networks == null) {
      this.networks = new ArrayList<NetworkResource>();
    }
    this.networks.add(networksItem);
    return this;
  }

   /**
   * Get networks
   * @return networks
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<NetworkResource> getNetworks() {
    return networks;
  }


  public void setNetworks(List<NetworkResource> networks) {
    this.networks = networks;
  }


  public TaskGroup reschedulePolicy(ReschedulePolicy reschedulePolicy) {
    
    this.reschedulePolicy = reschedulePolicy;
    return this;
  }

   /**
   * Get reschedulePolicy
   * @return reschedulePolicy
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ReschedulePolicy getReschedulePolicy() {
    return reschedulePolicy;
  }


  public void setReschedulePolicy(ReschedulePolicy reschedulePolicy) {
    this.reschedulePolicy = reschedulePolicy;
  }


  public TaskGroup restartPolicy(RestartPolicy restartPolicy) {
    
    this.restartPolicy = restartPolicy;
    return this;
  }

   /**
   * Get restartPolicy
   * @return restartPolicy
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public RestartPolicy getRestartPolicy() {
    return restartPolicy;
  }


  public void setRestartPolicy(RestartPolicy restartPolicy) {
    this.restartPolicy = restartPolicy;
  }


  public TaskGroup scaling(ScalingPolicy scaling) {
    
    this.scaling = scaling;
    return this;
  }

   /**
   * Get scaling
   * @return scaling
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ScalingPolicy getScaling() {
    return scaling;
  }


  public void setScaling(ScalingPolicy scaling) {
    this.scaling = scaling;
  }


  public TaskGroup services(List<Service> services) {
    
    this.services = services;
    return this;
  }

  public TaskGroup addServicesItem(Service servicesItem) {
    if (this.services == null) {
      this.services = new ArrayList<Service>();
    }
    this.services.add(servicesItem);
    return this;
  }

   /**
   * Get services
   * @return services
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Service> getServices() {
    return services;
  }


  public void setServices(List<Service> services) {
    this.services = services;
  }


  public TaskGroup shutdownDelay(Long shutdownDelay) {
    
    this.shutdownDelay = shutdownDelay;
    return this;
  }

   /**
   * Get shutdownDelay
   * @return shutdownDelay
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getShutdownDelay() {
    return shutdownDelay;
  }


  public void setShutdownDelay(Long shutdownDelay) {
    this.shutdownDelay = shutdownDelay;
  }


  public TaskGroup spreads(List<Spread> spreads) {
    
    this.spreads = spreads;
    return this;
  }

  public TaskGroup addSpreadsItem(Spread spreadsItem) {
    if (this.spreads == null) {
      this.spreads = new ArrayList<Spread>();
    }
    this.spreads.add(spreadsItem);
    return this;
  }

   /**
   * Get spreads
   * @return spreads
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Spread> getSpreads() {
    return spreads;
  }


  public void setSpreads(List<Spread> spreads) {
    this.spreads = spreads;
  }


  public TaskGroup stopAfterClientDisconnect(Long stopAfterClientDisconnect) {
    
    this.stopAfterClientDisconnect = stopAfterClientDisconnect;
    return this;
  }

   /**
   * Get stopAfterClientDisconnect
   * @return stopAfterClientDisconnect
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getStopAfterClientDisconnect() {
    return stopAfterClientDisconnect;
  }


  public void setStopAfterClientDisconnect(Long stopAfterClientDisconnect) {
    this.stopAfterClientDisconnect = stopAfterClientDisconnect;
  }


  public TaskGroup tasks(List<Task> tasks) {
    
    this.tasks = tasks;
    return this;
  }

  public TaskGroup addTasksItem(Task tasksItem) {
    if (this.tasks == null) {
      this.tasks = new ArrayList<Task>();
    }
    this.tasks.add(tasksItem);
    return this;
  }

   /**
   * Get tasks
   * @return tasks
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Task> getTasks() {
    return tasks;
  }


  public void setTasks(List<Task> tasks) {
    this.tasks = tasks;
  }


  public TaskGroup update(UpdateStrategy update) {
    
    this.update = update;
    return this;
  }

   /**
   * Get update
   * @return update
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public UpdateStrategy getUpdate() {
    return update;
  }


  public void setUpdate(UpdateStrategy update) {
    this.update = update;
  }


  public TaskGroup volumes(Map<String, VolumeRequest> volumes) {
    
    this.volumes = volumes;
    return this;
  }

  public TaskGroup putVolumesItem(String key, VolumeRequest volumesItem) {
    if (this.volumes == null) {
      this.volumes = new HashMap<String, VolumeRequest>();
    }
    this.volumes.put(key, volumesItem);
    return this;
  }

   /**
   * Get volumes
   * @return volumes
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Map<String, VolumeRequest> getVolumes() {
    return volumes;
  }


  public void setVolumes(Map<String, VolumeRequest> volumes) {
    this.volumes = volumes;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    TaskGroup taskGroup = (TaskGroup) o;
    return Objects.equals(this.affinities, taskGroup.affinities) &&
        Objects.equals(this.constraints, taskGroup.constraints) &&
        Objects.equals(this.consul, taskGroup.consul) &&
        Objects.equals(this.count, taskGroup.count) &&
        Objects.equals(this.ephemeralDisk, taskGroup.ephemeralDisk) &&
        Objects.equals(this.maxClientDisconnect, taskGroup.maxClientDisconnect) &&
        Objects.equals(this.meta, taskGroup.meta) &&
        Objects.equals(this.migrate, taskGroup.migrate) &&
        Objects.equals(this.name, taskGroup.name) &&
        Objects.equals(this.networks, taskGroup.networks) &&
        Objects.equals(this.reschedulePolicy, taskGroup.reschedulePolicy) &&
        Objects.equals(this.restartPolicy, taskGroup.restartPolicy) &&
        Objects.equals(this.scaling, taskGroup.scaling) &&
        Objects.equals(this.services, taskGroup.services) &&
        Objects.equals(this.shutdownDelay, taskGroup.shutdownDelay) &&
        Objects.equals(this.spreads, taskGroup.spreads) &&
        Objects.equals(this.stopAfterClientDisconnect, taskGroup.stopAfterClientDisconnect) &&
        Objects.equals(this.tasks, taskGroup.tasks) &&
        Objects.equals(this.update, taskGroup.update) &&
        Objects.equals(this.volumes, taskGroup.volumes);
  }

  @Override
  public int hashCode() {
    return Objects.hash(affinities, constraints, consul, count, ephemeralDisk, maxClientDisconnect, meta, migrate, name, networks, reschedulePolicy, restartPolicy, scaling, services, shutdownDelay, spreads, stopAfterClientDisconnect, tasks, update, volumes);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class TaskGroup {\n");
    sb.append("    affinities: ").append(toIndentedString(affinities)).append("\n");
    sb.append("    constraints: ").append(toIndentedString(constraints)).append("\n");
    sb.append("    consul: ").append(toIndentedString(consul)).append("\n");
    sb.append("    count: ").append(toIndentedString(count)).append("\n");
    sb.append("    ephemeralDisk: ").append(toIndentedString(ephemeralDisk)).append("\n");
    sb.append("    maxClientDisconnect: ").append(toIndentedString(maxClientDisconnect)).append("\n");
    sb.append("    meta: ").append(toIndentedString(meta)).append("\n");
    sb.append("    migrate: ").append(toIndentedString(migrate)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    networks: ").append(toIndentedString(networks)).append("\n");
    sb.append("    reschedulePolicy: ").append(toIndentedString(reschedulePolicy)).append("\n");
    sb.append("    restartPolicy: ").append(toIndentedString(restartPolicy)).append("\n");
    sb.append("    scaling: ").append(toIndentedString(scaling)).append("\n");
    sb.append("    services: ").append(toIndentedString(services)).append("\n");
    sb.append("    shutdownDelay: ").append(toIndentedString(shutdownDelay)).append("\n");
    sb.append("    spreads: ").append(toIndentedString(spreads)).append("\n");
    sb.append("    stopAfterClientDisconnect: ").append(toIndentedString(stopAfterClientDisconnect)).append("\n");
    sb.append("    tasks: ").append(toIndentedString(tasks)).append("\n");
    sb.append("    update: ").append(toIndentedString(update)).append("\n");
    sb.append("    volumes: ").append(toIndentedString(volumes)).append("\n");
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

