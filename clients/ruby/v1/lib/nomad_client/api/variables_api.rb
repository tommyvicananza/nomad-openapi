=begin
#Nomad

#No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

The version of the OpenAPI document: 1.1.4
Contact: support@hashicorp.com
Generated by: https://openapi-generator.tech
OpenAPI Generator version: 5.2.0

=end

require 'cgi'

module NomadClient
  class VariablesApi
    attr_accessor :api_client

    def initialize(api_client = ApiClient.default)
      @api_client = api_client
    end
    # @param path [String] A path to a Nomad Variable
    # @param variable [Variable] 
    # @param [Hash] opts the optional parameters
    # @option opts [String] :region Filters results based on the specified region.
    # @option opts [String] :namespace Filters results based on the specified namespace.
    # @option opts [String] :x_nomad_token A Nomad ACL token.
    # @option opts [String] :idempotency_token Can be used to ensure operations are only run once.
    # @option opts [Integer] :cas A compare-and-set parameter for Nomad Variables
    # @return [nil]
    def delete_variable(path, variable, opts = {})
      delete_variable_with_http_info(path, variable, opts)
      nil
    end

    # @param path [String] A path to a Nomad Variable
    # @param variable [Variable] 
    # @param [Hash] opts the optional parameters
    # @option opts [String] :region Filters results based on the specified region.
    # @option opts [String] :namespace Filters results based on the specified namespace.
    # @option opts [String] :x_nomad_token A Nomad ACL token.
    # @option opts [String] :idempotency_token Can be used to ensure operations are only run once.
    # @option opts [Integer] :cas A compare-and-set parameter for Nomad Variables
    # @return [Array<(nil, Integer, Hash)>] nil, response status code and response headers
    def delete_variable_with_http_info(path, variable, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: VariablesApi.delete_variable ...'
      end
      # verify the required parameter 'path' is set
      if @api_client.config.client_side_validation && path.nil?
        fail ArgumentError, "Missing the required parameter 'path' when calling VariablesApi.delete_variable"
      end
      # verify the required parameter 'variable' is set
      if @api_client.config.client_side_validation && variable.nil?
        fail ArgumentError, "Missing the required parameter 'variable' when calling VariablesApi.delete_variable"
      end
      # resource path
      local_var_path = '/var/{path}'.sub('{' + 'path' + '}', CGI.escape(path.to_s))

      # query parameters
      query_params = opts[:query_params] || {}
      query_params[:'region'] = opts[:'region'] if !opts[:'region'].nil?
      query_params[:'namespace'] = opts[:'namespace'] if !opts[:'namespace'].nil?
      query_params[:'idempotency_token'] = opts[:'idempotency_token'] if !opts[:'idempotency_token'].nil?
      query_params[:'cas'] = opts[:'cas'] if !opts[:'cas'].nil?

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Content-Type'
      header_params['Content-Type'] = @api_client.select_header_content_type(['application/json'])
      header_params[:'X-Nomad-Token'] = opts[:'x_nomad_token'] if !opts[:'x_nomad_token'].nil?

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body] || @api_client.object_to_http_body(variable)

      # return_type
      return_type = opts[:debug_return_type]

      # auth_names
      auth_names = opts[:debug_auth_names] || ['X-Nomad-Token']

      new_options = opts.merge(
        :operation => :"VariablesApi.delete_variable",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:DELETE, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: VariablesApi#delete_variable\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # @param path [String] A path to a Nomad Variable
    # @param [Hash] opts the optional parameters
    # @option opts [String] :region Filters results based on the specified region.
    # @option opts [String] :namespace Filters results based on the specified namespace.
    # @option opts [Integer] :index If set, wait until query exceeds given index. Must be provided with WaitParam.
    # @option opts [String] :wait Provided with IndexParam to wait for change.
    # @option opts [String] :stale If present, results will include stale reads.
    # @option opts [String] :prefix Constrains results to jobs that start with the defined prefix
    # @option opts [String] :x_nomad_token A Nomad ACL token.
    # @option opts [Integer] :per_page Maximum number of results to return.
    # @option opts [String] :next_token Indicates where to start paging for queries that support pagination.
    # @return [Variable]
    def get_variable_query(path, opts = {})
      data, _status_code, _headers = get_variable_query_with_http_info(path, opts)
      data
    end

    # @param path [String] A path to a Nomad Variable
    # @param [Hash] opts the optional parameters
    # @option opts [String] :region Filters results based on the specified region.
    # @option opts [String] :namespace Filters results based on the specified namespace.
    # @option opts [Integer] :index If set, wait until query exceeds given index. Must be provided with WaitParam.
    # @option opts [String] :wait Provided with IndexParam to wait for change.
    # @option opts [String] :stale If present, results will include stale reads.
    # @option opts [String] :prefix Constrains results to jobs that start with the defined prefix
    # @option opts [String] :x_nomad_token A Nomad ACL token.
    # @option opts [Integer] :per_page Maximum number of results to return.
    # @option opts [String] :next_token Indicates where to start paging for queries that support pagination.
    # @return [Array<(Variable, Integer, Hash)>] Variable data, response status code and response headers
    def get_variable_query_with_http_info(path, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: VariablesApi.get_variable_query ...'
      end
      # verify the required parameter 'path' is set
      if @api_client.config.client_side_validation && path.nil?
        fail ArgumentError, "Missing the required parameter 'path' when calling VariablesApi.get_variable_query"
      end
      # resource path
      local_var_path = '/var/{path}'.sub('{' + 'path' + '}', CGI.escape(path.to_s))

      # query parameters
      query_params = opts[:query_params] || {}
      query_params[:'region'] = opts[:'region'] if !opts[:'region'].nil?
      query_params[:'namespace'] = opts[:'namespace'] if !opts[:'namespace'].nil?
      query_params[:'wait'] = opts[:'wait'] if !opts[:'wait'].nil?
      query_params[:'stale'] = opts[:'stale'] if !opts[:'stale'].nil?
      query_params[:'prefix'] = opts[:'prefix'] if !opts[:'prefix'].nil?
      query_params[:'per_page'] = opts[:'per_page'] if !opts[:'per_page'].nil?
      query_params[:'next_token'] = opts[:'next_token'] if !opts[:'next_token'].nil?

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])
      header_params[:'index'] = opts[:'index'] if !opts[:'index'].nil?
      header_params[:'X-Nomad-Token'] = opts[:'x_nomad_token'] if !opts[:'x_nomad_token'].nil?

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'Variable'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['X-Nomad-Token']

      new_options = opts.merge(
        :operation => :"VariablesApi.get_variable_query",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: VariablesApi#get_variable_query\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # @param [Hash] opts the optional parameters
    # @option opts [String] :region Filters results based on the specified region.
    # @option opts [String] :namespace Filters results based on the specified namespace.
    # @option opts [Integer] :index If set, wait until query exceeds given index. Must be provided with WaitParam.
    # @option opts [String] :wait Provided with IndexParam to wait for change.
    # @option opts [String] :stale If present, results will include stale reads.
    # @option opts [String] :prefix Constrains results to jobs that start with the defined prefix
    # @option opts [String] :x_nomad_token A Nomad ACL token.
    # @option opts [Integer] :per_page Maximum number of results to return.
    # @option opts [String] :next_token Indicates where to start paging for queries that support pagination.
    # @return [Array<VariableMetadata>]
    def get_variables_list_request(opts = {})
      data, _status_code, _headers = get_variables_list_request_with_http_info(opts)
      data
    end

    # @param [Hash] opts the optional parameters
    # @option opts [String] :region Filters results based on the specified region.
    # @option opts [String] :namespace Filters results based on the specified namespace.
    # @option opts [Integer] :index If set, wait until query exceeds given index. Must be provided with WaitParam.
    # @option opts [String] :wait Provided with IndexParam to wait for change.
    # @option opts [String] :stale If present, results will include stale reads.
    # @option opts [String] :prefix Constrains results to jobs that start with the defined prefix
    # @option opts [String] :x_nomad_token A Nomad ACL token.
    # @option opts [Integer] :per_page Maximum number of results to return.
    # @option opts [String] :next_token Indicates where to start paging for queries that support pagination.
    # @return [Array<(Array<VariableMetadata>, Integer, Hash)>] Array<VariableMetadata> data, response status code and response headers
    def get_variables_list_request_with_http_info(opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: VariablesApi.get_variables_list_request ...'
      end
      # resource path
      local_var_path = '/vars'

      # query parameters
      query_params = opts[:query_params] || {}
      query_params[:'region'] = opts[:'region'] if !opts[:'region'].nil?
      query_params[:'namespace'] = opts[:'namespace'] if !opts[:'namespace'].nil?
      query_params[:'wait'] = opts[:'wait'] if !opts[:'wait'].nil?
      query_params[:'stale'] = opts[:'stale'] if !opts[:'stale'].nil?
      query_params[:'prefix'] = opts[:'prefix'] if !opts[:'prefix'].nil?
      query_params[:'per_page'] = opts[:'per_page'] if !opts[:'per_page'].nil?
      query_params[:'next_token'] = opts[:'next_token'] if !opts[:'next_token'].nil?

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])
      header_params[:'index'] = opts[:'index'] if !opts[:'index'].nil?
      header_params[:'X-Nomad-Token'] = opts[:'x_nomad_token'] if !opts[:'x_nomad_token'].nil?

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'Array<VariableMetadata>'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['X-Nomad-Token']

      new_options = opts.merge(
        :operation => :"VariablesApi.get_variables_list_request",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: VariablesApi#get_variables_list_request\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # @param path [String] A path to a Nomad Variable
    # @param variable [Variable] 
    # @param [Hash] opts the optional parameters
    # @option opts [String] :region Filters results based on the specified region.
    # @option opts [String] :namespace Filters results based on the specified namespace.
    # @option opts [String] :x_nomad_token A Nomad ACL token.
    # @option opts [String] :idempotency_token Can be used to ensure operations are only run once.
    # @option opts [Integer] :cas A compare-and-set parameter for Nomad Variables
    # @return [Variable]
    def post_variable(path, variable, opts = {})
      data, _status_code, _headers = post_variable_with_http_info(path, variable, opts)
      data
    end

    # @param path [String] A path to a Nomad Variable
    # @param variable [Variable] 
    # @param [Hash] opts the optional parameters
    # @option opts [String] :region Filters results based on the specified region.
    # @option opts [String] :namespace Filters results based on the specified namespace.
    # @option opts [String] :x_nomad_token A Nomad ACL token.
    # @option opts [String] :idempotency_token Can be used to ensure operations are only run once.
    # @option opts [Integer] :cas A compare-and-set parameter for Nomad Variables
    # @return [Array<(Variable, Integer, Hash)>] Variable data, response status code and response headers
    def post_variable_with_http_info(path, variable, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: VariablesApi.post_variable ...'
      end
      # verify the required parameter 'path' is set
      if @api_client.config.client_side_validation && path.nil?
        fail ArgumentError, "Missing the required parameter 'path' when calling VariablesApi.post_variable"
      end
      # verify the required parameter 'variable' is set
      if @api_client.config.client_side_validation && variable.nil?
        fail ArgumentError, "Missing the required parameter 'variable' when calling VariablesApi.post_variable"
      end
      # resource path
      local_var_path = '/var/{path}'.sub('{' + 'path' + '}', CGI.escape(path.to_s))

      # query parameters
      query_params = opts[:query_params] || {}
      query_params[:'region'] = opts[:'region'] if !opts[:'region'].nil?
      query_params[:'namespace'] = opts[:'namespace'] if !opts[:'namespace'].nil?
      query_params[:'idempotency_token'] = opts[:'idempotency_token'] if !opts[:'idempotency_token'].nil?
      query_params[:'cas'] = opts[:'cas'] if !opts[:'cas'].nil?

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])
      # HTTP header 'Content-Type'
      header_params['Content-Type'] = @api_client.select_header_content_type(['application/json'])
      header_params[:'X-Nomad-Token'] = opts[:'x_nomad_token'] if !opts[:'x_nomad_token'].nil?

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body] || @api_client.object_to_http_body(variable)

      # return_type
      return_type = opts[:debug_return_type] || 'Variable'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['X-Nomad-Token']

      new_options = opts.merge(
        :operation => :"VariablesApi.post_variable",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: VariablesApi#post_variable\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # @param path [String] A path to a Nomad Variable
    # @param variable [Variable] 
    # @param [Hash] opts the optional parameters
    # @option opts [String] :region Filters results based on the specified region.
    # @option opts [String] :namespace Filters results based on the specified namespace.
    # @option opts [String] :x_nomad_token A Nomad ACL token.
    # @option opts [String] :idempotency_token Can be used to ensure operations are only run once.
    # @option opts [Integer] :cas A compare-and-set parameter for Nomad Variables
    # @return [Variable]
    def put_variable(path, variable, opts = {})
      data, _status_code, _headers = put_variable_with_http_info(path, variable, opts)
      data
    end

    # @param path [String] A path to a Nomad Variable
    # @param variable [Variable] 
    # @param [Hash] opts the optional parameters
    # @option opts [String] :region Filters results based on the specified region.
    # @option opts [String] :namespace Filters results based on the specified namespace.
    # @option opts [String] :x_nomad_token A Nomad ACL token.
    # @option opts [String] :idempotency_token Can be used to ensure operations are only run once.
    # @option opts [Integer] :cas A compare-and-set parameter for Nomad Variables
    # @return [Array<(Variable, Integer, Hash)>] Variable data, response status code and response headers
    def put_variable_with_http_info(path, variable, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: VariablesApi.put_variable ...'
      end
      # verify the required parameter 'path' is set
      if @api_client.config.client_side_validation && path.nil?
        fail ArgumentError, "Missing the required parameter 'path' when calling VariablesApi.put_variable"
      end
      # verify the required parameter 'variable' is set
      if @api_client.config.client_side_validation && variable.nil?
        fail ArgumentError, "Missing the required parameter 'variable' when calling VariablesApi.put_variable"
      end
      # resource path
      local_var_path = '/var/{path}'.sub('{' + 'path' + '}', CGI.escape(path.to_s))

      # query parameters
      query_params = opts[:query_params] || {}
      query_params[:'region'] = opts[:'region'] if !opts[:'region'].nil?
      query_params[:'namespace'] = opts[:'namespace'] if !opts[:'namespace'].nil?
      query_params[:'idempotency_token'] = opts[:'idempotency_token'] if !opts[:'idempotency_token'].nil?
      query_params[:'cas'] = opts[:'cas'] if !opts[:'cas'].nil?

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])
      # HTTP header 'Content-Type'
      header_params['Content-Type'] = @api_client.select_header_content_type(['application/json'])
      header_params[:'X-Nomad-Token'] = opts[:'x_nomad_token'] if !opts[:'x_nomad_token'].nil?

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body] || @api_client.object_to_http_body(variable)

      # return_type
      return_type = opts[:debug_return_type] || 'Variable'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['X-Nomad-Token']

      new_options = opts.merge(
        :operation => :"VariablesApi.put_variable",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:PUT, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: VariablesApi#put_variable\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end
  end
end
