=begin
#Nomad

#No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

The version of the OpenAPI document: 1.1.4
Contact: support@hashicorp.com
Generated by: https://openapi-generator.tech
OpenAPI Generator version: 5.2.0

=end

require 'spec_helper'
require 'json'

# Unit tests for NomadClient::VariablesApi
# Automatically generated by openapi-generator (https://openapi-generator.tech)
# Please update as you see appropriate
describe 'VariablesApi' do
  before do
    # run before each test
    @api_instance = NomadClient::VariablesApi.new
  end

  after do
    # run after each test
  end

  describe 'test an instance of VariablesApi' do
    it 'should create an instance of VariablesApi' do
      expect(@api_instance).to be_instance_of(NomadClient::VariablesApi)
    end
  end

  # unit tests for delete_variable
  # @param path A path to a Nomad Variable
  # @param variable 
  # @param [Hash] opts the optional parameters
  # @option opts [String] :region Filters results based on the specified region.
  # @option opts [String] :namespace Filters results based on the specified namespace.
  # @option opts [String] :x_nomad_token A Nomad ACL token.
  # @option opts [String] :idempotency_token Can be used to ensure operations are only run once.
  # @option opts [Integer] :cas A compare-and-set parameter for Nomad Variables
  # @return [nil]
  describe 'delete_variable test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for get_variable_query
  # @param path A path to a Nomad Variable
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
  describe 'get_variable_query test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for get_variables_list_request
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
  describe 'get_variables_list_request test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for post_variable
  # @param path A path to a Nomad Variable
  # @param variable 
  # @param [Hash] opts the optional parameters
  # @option opts [String] :region Filters results based on the specified region.
  # @option opts [String] :namespace Filters results based on the specified namespace.
  # @option opts [String] :x_nomad_token A Nomad ACL token.
  # @option opts [String] :idempotency_token Can be used to ensure operations are only run once.
  # @option opts [Integer] :cas A compare-and-set parameter for Nomad Variables
  # @return [Variable]
  describe 'post_variable test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for put_variable
  # @param path A path to a Nomad Variable
  # @param variable 
  # @param [Hash] opts the optional parameters
  # @option opts [String] :region Filters results based on the specified region.
  # @option opts [String] :namespace Filters results based on the specified namespace.
  # @option opts [String] :x_nomad_token A Nomad ACL token.
  # @option opts [String] :idempotency_token Can be used to ensure operations are only run once.
  # @option opts [Integer] :cas A compare-and-set parameter for Nomad Variables
  # @return [Variable]
  describe 'put_variable test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end
