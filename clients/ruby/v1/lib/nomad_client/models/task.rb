=begin
#Nomad

#No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

The version of the OpenAPI document: 1.1.4
Contact: support@hashicorp.com
Generated by: https://openapi-generator.tech
OpenAPI Generator version: 5.2.0

=end

require 'date'
require 'time'

module NomadClient
  class Task
    attr_accessor :affinities

    attr_accessor :artifacts

    attr_accessor :csi_plugin_config

    attr_accessor :config

    attr_accessor :constraints

    attr_accessor :dispatch_payload

    attr_accessor :driver

    attr_accessor :env

    attr_accessor :identity

    attr_accessor :kill_signal

    attr_accessor :kill_timeout

    attr_accessor :kind

    attr_accessor :leader

    attr_accessor :lifecycle

    attr_accessor :log_config

    attr_accessor :meta

    attr_accessor :name

    attr_accessor :resources

    attr_accessor :restart_policy

    attr_accessor :scaling_policies

    attr_accessor :services

    attr_accessor :shutdown_delay

    attr_accessor :templates

    attr_accessor :user

    attr_accessor :vault

    attr_accessor :volume_mounts

    # Attribute mapping from ruby-style variable name to JSON key.
    def self.attribute_map
      {
        :'affinities' => :'Affinities',
        :'artifacts' => :'Artifacts',
        :'csi_plugin_config' => :'CSIPluginConfig',
        :'config' => :'Config',
        :'constraints' => :'Constraints',
        :'dispatch_payload' => :'DispatchPayload',
        :'driver' => :'Driver',
        :'env' => :'Env',
        :'identity' => :'Identity',
        :'kill_signal' => :'KillSignal',
        :'kill_timeout' => :'KillTimeout',
        :'kind' => :'Kind',
        :'leader' => :'Leader',
        :'lifecycle' => :'Lifecycle',
        :'log_config' => :'LogConfig',
        :'meta' => :'Meta',
        :'name' => :'Name',
        :'resources' => :'Resources',
        :'restart_policy' => :'RestartPolicy',
        :'scaling_policies' => :'ScalingPolicies',
        :'services' => :'Services',
        :'shutdown_delay' => :'ShutdownDelay',
        :'templates' => :'Templates',
        :'user' => :'User',
        :'vault' => :'Vault',
        :'volume_mounts' => :'VolumeMounts'
      }
    end

    # Returns all the JSON keys this model knows about
    def self.acceptable_attributes
      attribute_map.values
    end

    # Attribute type mapping.
    def self.openapi_types
      {
        :'affinities' => :'Array<Affinity>',
        :'artifacts' => :'Array<TaskArtifact>',
        :'csi_plugin_config' => :'TaskCSIPluginConfig',
        :'config' => :'Hash<String, AnyType>',
        :'constraints' => :'Array<Constraint>',
        :'dispatch_payload' => :'DispatchPayloadConfig',
        :'driver' => :'String',
        :'env' => :'Hash<String, String>',
        :'identity' => :'WorkloadIdentity',
        :'kill_signal' => :'String',
        :'kill_timeout' => :'Integer',
        :'kind' => :'String',
        :'leader' => :'Boolean',
        :'lifecycle' => :'TaskLifecycle',
        :'log_config' => :'LogConfig',
        :'meta' => :'Hash<String, String>',
        :'name' => :'String',
        :'resources' => :'Resources',
        :'restart_policy' => :'RestartPolicy',
        :'scaling_policies' => :'Array<ScalingPolicy>',
        :'services' => :'Array<Service>',
        :'shutdown_delay' => :'Integer',
        :'templates' => :'Array<Template>',
        :'user' => :'String',
        :'vault' => :'Vault',
        :'volume_mounts' => :'Array<VolumeMount>'
      }
    end

    # List of attributes with nullable: true
    def self.openapi_nullable
      Set.new([
      ])
    end

    # Initializes the object
    # @param [Hash] attributes Model attributes in the form of hash
    def initialize(attributes = {})
      if (!attributes.is_a?(Hash))
        fail ArgumentError, "The input argument (attributes) must be a hash in `NomadClient::Task` initialize method"
      end

      # check to see if the attribute exists and convert string to symbol for hash key
      attributes = attributes.each_with_object({}) { |(k, v), h|
        if (!self.class.attribute_map.key?(k.to_sym))
          fail ArgumentError, "`#{k}` is not a valid attribute in `NomadClient::Task`. Please check the name to make sure it's valid. List of attributes: " + self.class.attribute_map.keys.inspect
        end
        h[k.to_sym] = v
      }

      if attributes.key?(:'affinities')
        if (value = attributes[:'affinities']).is_a?(Array)
          self.affinities = value
        end
      end

      if attributes.key?(:'artifacts')
        if (value = attributes[:'artifacts']).is_a?(Array)
          self.artifacts = value
        end
      end

      if attributes.key?(:'csi_plugin_config')
        self.csi_plugin_config = attributes[:'csi_plugin_config']
      end

      if attributes.key?(:'config')
        if (value = attributes[:'config']).is_a?(Hash)
          self.config = value
        end
      end

      if attributes.key?(:'constraints')
        if (value = attributes[:'constraints']).is_a?(Array)
          self.constraints = value
        end
      end

      if attributes.key?(:'dispatch_payload')
        self.dispatch_payload = attributes[:'dispatch_payload']
      end

      if attributes.key?(:'driver')
        self.driver = attributes[:'driver']
      end

      if attributes.key?(:'env')
        if (value = attributes[:'env']).is_a?(Hash)
          self.env = value
        end
      end

      if attributes.key?(:'identity')
        self.identity = attributes[:'identity']
      end

      if attributes.key?(:'kill_signal')
        self.kill_signal = attributes[:'kill_signal']
      end

      if attributes.key?(:'kill_timeout')
        self.kill_timeout = attributes[:'kill_timeout']
      end

      if attributes.key?(:'kind')
        self.kind = attributes[:'kind']
      end

      if attributes.key?(:'leader')
        self.leader = attributes[:'leader']
      end

      if attributes.key?(:'lifecycle')
        self.lifecycle = attributes[:'lifecycle']
      end

      if attributes.key?(:'log_config')
        self.log_config = attributes[:'log_config']
      end

      if attributes.key?(:'meta')
        if (value = attributes[:'meta']).is_a?(Hash)
          self.meta = value
        end
      end

      if attributes.key?(:'name')
        self.name = attributes[:'name']
      end

      if attributes.key?(:'resources')
        self.resources = attributes[:'resources']
      end

      if attributes.key?(:'restart_policy')
        self.restart_policy = attributes[:'restart_policy']
      end

      if attributes.key?(:'scaling_policies')
        if (value = attributes[:'scaling_policies']).is_a?(Array)
          self.scaling_policies = value
        end
      end

      if attributes.key?(:'services')
        if (value = attributes[:'services']).is_a?(Array)
          self.services = value
        end
      end

      if attributes.key?(:'shutdown_delay')
        self.shutdown_delay = attributes[:'shutdown_delay']
      end

      if attributes.key?(:'templates')
        if (value = attributes[:'templates']).is_a?(Array)
          self.templates = value
        end
      end

      if attributes.key?(:'user')
        self.user = attributes[:'user']
      end

      if attributes.key?(:'vault')
        self.vault = attributes[:'vault']
      end

      if attributes.key?(:'volume_mounts')
        if (value = attributes[:'volume_mounts']).is_a?(Array)
          self.volume_mounts = value
        end
      end
    end

    # Show invalid properties with the reasons. Usually used together with valid?
    # @return Array for valid properties with the reasons
    def list_invalid_properties
      invalid_properties = Array.new
      invalid_properties
    end

    # Check to see if the all the properties in the model are valid
    # @return true if the model is valid
    def valid?
      true
    end

    # Checks equality by comparing each attribute.
    # @param [Object] Object to be compared
    def ==(o)
      return true if self.equal?(o)
      self.class == o.class &&
          affinities == o.affinities &&
          artifacts == o.artifacts &&
          csi_plugin_config == o.csi_plugin_config &&
          config == o.config &&
          constraints == o.constraints &&
          dispatch_payload == o.dispatch_payload &&
          driver == o.driver &&
          env == o.env &&
          identity == o.identity &&
          kill_signal == o.kill_signal &&
          kill_timeout == o.kill_timeout &&
          kind == o.kind &&
          leader == o.leader &&
          lifecycle == o.lifecycle &&
          log_config == o.log_config &&
          meta == o.meta &&
          name == o.name &&
          resources == o.resources &&
          restart_policy == o.restart_policy &&
          scaling_policies == o.scaling_policies &&
          services == o.services &&
          shutdown_delay == o.shutdown_delay &&
          templates == o.templates &&
          user == o.user &&
          vault == o.vault &&
          volume_mounts == o.volume_mounts
    end

    # @see the `==` method
    # @param [Object] Object to be compared
    def eql?(o)
      self == o
    end

    # Calculates hash code according to all attributes.
    # @return [Integer] Hash code
    def hash
      [affinities, artifacts, csi_plugin_config, config, constraints, dispatch_payload, driver, env, identity, kill_signal, kill_timeout, kind, leader, lifecycle, log_config, meta, name, resources, restart_policy, scaling_policies, services, shutdown_delay, templates, user, vault, volume_mounts].hash
    end

    # Builds the object from hash
    # @param [Hash] attributes Model attributes in the form of hash
    # @return [Object] Returns the model itself
    def self.build_from_hash(attributes)
      new.build_from_hash(attributes)
    end

    # Builds the object from hash
    # @param [Hash] attributes Model attributes in the form of hash
    # @return [Object] Returns the model itself
    def build_from_hash(attributes)
      return nil unless attributes.is_a?(Hash)
      self.class.openapi_types.each_pair do |key, type|
        if attributes[self.class.attribute_map[key]].nil? && self.class.openapi_nullable.include?(key)
          self.send("#{key}=", nil)
        elsif type =~ /\AArray<(.*)>/i
          # check to ensure the input is an array given that the attribute
          # is documented as an array but the input is not
          if attributes[self.class.attribute_map[key]].is_a?(Array)
            self.send("#{key}=", attributes[self.class.attribute_map[key]].map { |v| _deserialize($1, v) })
          end
        elsif !attributes[self.class.attribute_map[key]].nil?
          self.send("#{key}=", _deserialize(type, attributes[self.class.attribute_map[key]]))
        end
      end

      self
    end

    # Deserializes the data based on type
    # @param string type Data type
    # @param string value Value to be deserialized
    # @return [Object] Deserialized data
    def _deserialize(type, value)
      case type.to_sym
      when :Time
        Time.parse(value)
      when :Date
        Date.parse(value)
      when :String
        value.to_s
      when :Integer
        value.to_i
      when :Float
        value.to_f
      when :Boolean
        if value.to_s =~ /\A(true|t|yes|y|1)\z/i
          true
        else
          false
        end
      when :Object
        # generic object (usually a Hash), return directly
        value
      when /\AArray<(?<inner_type>.+)>\z/
        inner_type = Regexp.last_match[:inner_type]
        value.map { |v| _deserialize(inner_type, v) }
      when /\AHash<(?<k_type>.+?), (?<v_type>.+)>\z/
        k_type = Regexp.last_match[:k_type]
        v_type = Regexp.last_match[:v_type]
        {}.tap do |hash|
          value.each do |k, v|
            hash[_deserialize(k_type, k)] = _deserialize(v_type, v)
          end
        end
      else # model
        # models (e.g. Pet) or oneOf
        klass = NomadClient.const_get(type)
        klass.respond_to?(:openapi_one_of) ? klass.build(value) : klass.build_from_hash(value)
      end
    end

    # Returns the string representation of the object
    # @return [String] String presentation of the object
    def to_s
      to_hash.to_s
    end

    # to_body is an alias to to_hash (backward compatibility)
    # @return [Hash] Returns the object in the form of hash
    def to_body
      to_hash
    end

    # Returns the object in the form of hash
    # @return [Hash] Returns the object in the form of hash
    def to_hash
      hash = {}
      self.class.attribute_map.each_pair do |attr, param|
        value = self.send(attr)
        if value.nil?
          is_nullable = self.class.openapi_nullable.include?(attr)
          next if !is_nullable || (is_nullable && !instance_variable_defined?(:"@#{attr}"))
        end

        hash[param] = _to_hash(value)
      end
      hash
    end

    # Outputs non-array value in the form of hash
    # For object, use to_hash. Otherwise, just return the value
    # @param [Object] value Any valid value
    # @return [Hash] Returns the value in the form of hash
    def _to_hash(value)
      if value.is_a?(Array)
        value.compact.map { |v| _to_hash(v) }
      elsif value.is_a?(Hash)
        {}.tap do |hash|
          value.each { |k, v| hash[k] = _to_hash(v) }
        end
      elsif value.respond_to? :to_hash
        value.to_hash
      else
        value
      end
    end

  end

end
