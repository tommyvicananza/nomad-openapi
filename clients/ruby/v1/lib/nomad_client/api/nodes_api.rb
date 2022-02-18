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
  class NodesApi
    attr_accessor :api_client

    def initialize(api_client = ApiClient.default)
      @api_client = api_client
    end
    # @param node_id [String] The ID of the node.
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
    # @return [Node]
    def get_node(node_id, opts = {})
      data, _status_code, _headers = get_node_with_http_info(node_id, opts)
      data
    end

    # @param node_id [String] The ID of the node.
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
    # @return [Array<(Node, Integer, Hash)>] Node data, response status code and response headers
    def get_node_with_http_info(node_id, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: NodesApi.get_node ...'
      end
      # verify the required parameter 'node_id' is set
      if @api_client.config.client_side_validation && node_id.nil?
        fail ArgumentError, "Missing the required parameter 'node_id' when calling NodesApi.get_node"
      end
      # resource path
      local_var_path = '/node/{nodeId}'.sub('{' + 'nodeId' + '}', CGI.escape(node_id.to_s))

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
      return_type = opts[:debug_return_type] || 'Node'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['X-Nomad-Token']

      new_options = opts.merge(
        :operation => :"NodesApi.get_node",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: NodesApi#get_node\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # @param node_id [String] The ID of the node.
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
    # @return [Array<AllocationListStub>]
    def get_node_allocations(node_id, opts = {})
      data, _status_code, _headers = get_node_allocations_with_http_info(node_id, opts)
      data
    end

    # @param node_id [String] The ID of the node.
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
    # @return [Array<(Array<AllocationListStub>, Integer, Hash)>] Array<AllocationListStub> data, response status code and response headers
    def get_node_allocations_with_http_info(node_id, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: NodesApi.get_node_allocations ...'
      end
      # verify the required parameter 'node_id' is set
      if @api_client.config.client_side_validation && node_id.nil?
        fail ArgumentError, "Missing the required parameter 'node_id' when calling NodesApi.get_node_allocations"
      end
      # resource path
      local_var_path = '/node/{nodeId}/allocations'.sub('{' + 'nodeId' + '}', CGI.escape(node_id.to_s))

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
      return_type = opts[:debug_return_type] || 'Array<AllocationListStub>'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['X-Nomad-Token']

      new_options = opts.merge(
        :operation => :"NodesApi.get_node_allocations",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: NodesApi#get_node_allocations\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
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
    # @option opts [Boolean] :resources Whether or not to include the NodeResources and ReservedResources fields in the response.
    # @return [Array<NodeListStub>]
    def get_nodes(opts = {})
      data, _status_code, _headers = get_nodes_with_http_info(opts)
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
    # @option opts [Boolean] :resources Whether or not to include the NodeResources and ReservedResources fields in the response.
    # @return [Array<(Array<NodeListStub>, Integer, Hash)>] Array<NodeListStub> data, response status code and response headers
    def get_nodes_with_http_info(opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: NodesApi.get_nodes ...'
      end
      # resource path
      local_var_path = '/nodes'

      # query parameters
      query_params = opts[:query_params] || {}
      query_params[:'region'] = opts[:'region'] if !opts[:'region'].nil?
      query_params[:'namespace'] = opts[:'namespace'] if !opts[:'namespace'].nil?
      query_params[:'wait'] = opts[:'wait'] if !opts[:'wait'].nil?
      query_params[:'stale'] = opts[:'stale'] if !opts[:'stale'].nil?
      query_params[:'prefix'] = opts[:'prefix'] if !opts[:'prefix'].nil?
      query_params[:'per_page'] = opts[:'per_page'] if !opts[:'per_page'].nil?
      query_params[:'next_token'] = opts[:'next_token'] if !opts[:'next_token'].nil?
      query_params[:'resources'] = opts[:'resources'] if !opts[:'resources'].nil?

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
      return_type = opts[:debug_return_type] || 'Array<NodeListStub>'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['X-Nomad-Token']

      new_options = opts.merge(
        :operation => :"NodesApi.get_nodes",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: NodesApi#get_nodes\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # @param node_id [String] The ID of the node.
    # @param node_update_drain_request [NodeUpdateDrainRequest] 
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
    # @return [NodeDrainUpdateResponse]
    def update_node_drain(node_id, node_update_drain_request, opts = {})
      data, _status_code, _headers = update_node_drain_with_http_info(node_id, node_update_drain_request, opts)
      data
    end

    # @param node_id [String] The ID of the node.
    # @param node_update_drain_request [NodeUpdateDrainRequest] 
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
    # @return [Array<(NodeDrainUpdateResponse, Integer, Hash)>] NodeDrainUpdateResponse data, response status code and response headers
    def update_node_drain_with_http_info(node_id, node_update_drain_request, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: NodesApi.update_node_drain ...'
      end
      # verify the required parameter 'node_id' is set
      if @api_client.config.client_side_validation && node_id.nil?
        fail ArgumentError, "Missing the required parameter 'node_id' when calling NodesApi.update_node_drain"
      end
      # verify the required parameter 'node_update_drain_request' is set
      if @api_client.config.client_side_validation && node_update_drain_request.nil?
        fail ArgumentError, "Missing the required parameter 'node_update_drain_request' when calling NodesApi.update_node_drain"
      end
      # resource path
      local_var_path = '/node/{nodeId}/drain'.sub('{' + 'nodeId' + '}', CGI.escape(node_id.to_s))

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
      # HTTP header 'Content-Type'
      header_params['Content-Type'] = @api_client.select_header_content_type(['application/json'])
      header_params[:'index'] = opts[:'index'] if !opts[:'index'].nil?
      header_params[:'X-Nomad-Token'] = opts[:'x_nomad_token'] if !opts[:'x_nomad_token'].nil?

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body] || @api_client.object_to_http_body(node_update_drain_request)

      # return_type
      return_type = opts[:debug_return_type] || 'NodeDrainUpdateResponse'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['X-Nomad-Token']

      new_options = opts.merge(
        :operation => :"NodesApi.update_node_drain",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: NodesApi#update_node_drain\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # @param node_id [String] The ID of the node.
    # @param node_update_eligibility_request [NodeUpdateEligibilityRequest] 
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
    # @return [NodeEligibilityUpdateResponse]
    def update_node_eligibility(node_id, node_update_eligibility_request, opts = {})
      data, _status_code, _headers = update_node_eligibility_with_http_info(node_id, node_update_eligibility_request, opts)
      data
    end

    # @param node_id [String] The ID of the node.
    # @param node_update_eligibility_request [NodeUpdateEligibilityRequest] 
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
    # @return [Array<(NodeEligibilityUpdateResponse, Integer, Hash)>] NodeEligibilityUpdateResponse data, response status code and response headers
    def update_node_eligibility_with_http_info(node_id, node_update_eligibility_request, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: NodesApi.update_node_eligibility ...'
      end
      # verify the required parameter 'node_id' is set
      if @api_client.config.client_side_validation && node_id.nil?
        fail ArgumentError, "Missing the required parameter 'node_id' when calling NodesApi.update_node_eligibility"
      end
      # verify the required parameter 'node_update_eligibility_request' is set
      if @api_client.config.client_side_validation && node_update_eligibility_request.nil?
        fail ArgumentError, "Missing the required parameter 'node_update_eligibility_request' when calling NodesApi.update_node_eligibility"
      end
      # resource path
      local_var_path = '/node/{nodeId}/eligibility'.sub('{' + 'nodeId' + '}', CGI.escape(node_id.to_s))

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
      # HTTP header 'Content-Type'
      header_params['Content-Type'] = @api_client.select_header_content_type(['application/json'])
      header_params[:'index'] = opts[:'index'] if !opts[:'index'].nil?
      header_params[:'X-Nomad-Token'] = opts[:'x_nomad_token'] if !opts[:'x_nomad_token'].nil?

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body] || @api_client.object_to_http_body(node_update_eligibility_request)

      # return_type
      return_type = opts[:debug_return_type] || 'NodeEligibilityUpdateResponse'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['X-Nomad-Token']

      new_options = opts.merge(
        :operation => :"NodesApi.update_node_eligibility",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: NodesApi#update_node_eligibility\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # @param node_id [String] The ID of the node.
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
    # @return [NodePurgeResponse]
    def update_node_purge(node_id, opts = {})
      data, _status_code, _headers = update_node_purge_with_http_info(node_id, opts)
      data
    end

    # @param node_id [String] The ID of the node.
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
    # @return [Array<(NodePurgeResponse, Integer, Hash)>] NodePurgeResponse data, response status code and response headers
    def update_node_purge_with_http_info(node_id, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: NodesApi.update_node_purge ...'
      end
      # verify the required parameter 'node_id' is set
      if @api_client.config.client_side_validation && node_id.nil?
        fail ArgumentError, "Missing the required parameter 'node_id' when calling NodesApi.update_node_purge"
      end
      # resource path
      local_var_path = '/node/{nodeId}/purge'.sub('{' + 'nodeId' + '}', CGI.escape(node_id.to_s))

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
      return_type = opts[:debug_return_type] || 'NodePurgeResponse'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['X-Nomad-Token']

      new_options = opts.merge(
        :operation => :"NodesApi.update_node_purge",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: NodesApi#update_node_purge\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end
  end
end