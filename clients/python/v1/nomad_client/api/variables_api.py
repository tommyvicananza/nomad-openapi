"""
    Nomad

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)  # noqa: E501

    The version of the OpenAPI document: 1.1.4
    Contact: support@hashicorp.com
    Generated by: https://openapi-generator.tech
"""


import re  # noqa: F401
import sys  # noqa: F401

from nomad_client.api_client import ApiClient, Endpoint as _Endpoint
from nomad_client.model_utils import (  # noqa: F401
    check_allowed_values,
    check_validations,
    date,
    datetime,
    file_type,
    none_type,
    validate_and_convert_types
)
from nomad_client.model.variable import Variable
from nomad_client.model.variable_metadata import VariableMetadata


class VariablesApi(object):
    """NOTE: This class is auto generated by OpenAPI Generator
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    def __init__(self, api_client=None):
        if api_client is None:
            api_client = ApiClient()
        self.api_client = api_client

        def __delete_variable(
            self,
            path,
            variable,
            **kwargs
        ):
            """delete_variable  # noqa: E501

            This method makes a synchronous HTTP request by default. To make an
            asynchronous HTTP request, please pass async_req=True

            >>> thread = api.delete_variable(path, variable, async_req=True)
            >>> result = thread.get()

            Args:
                path (str): A path to a Nomad Variable
                variable (Variable):

            Keyword Args:
                region (str): Filters results based on the specified region.. [optional]
                namespace (str): Filters results based on the specified namespace.. [optional]
                x_nomad_token (str): A Nomad ACL token.. [optional]
                idempotency_token (str): Can be used to ensure operations are only run once.. [optional]
                cas (int): A compare-and-set parameter for Nomad Variables. [optional]
                _return_http_data_only (bool): response data without head status
                    code and headers. Default is True.
                _preload_content (bool): if False, the urllib3.HTTPResponse object
                    will be returned without reading/decoding response data.
                    Default is True.
                _request_timeout (int/float/tuple): timeout setting for this request. If
                    one number provided, it will be total request timeout. It can also
                    be a pair (tuple) of (connection, read) timeouts.
                    Default is None.
                _check_input_type (bool): specifies if type checking
                    should be done one the data sent to the server.
                    Default is True.
                _check_return_type (bool): specifies if type checking
                    should be done one the data received from the server.
                    Default is True.
                _host_index (int/None): specifies the index of the server
                    that we want to use.
                    Default is read from the configuration.
                async_req (bool): execute request asynchronously

            Returns:
                None
                    If the method is called asynchronously, returns the request
                    thread.
            """
            kwargs['async_req'] = kwargs.get(
                'async_req', False
            )
            kwargs['_return_http_data_only'] = kwargs.get(
                '_return_http_data_only', True
            )
            kwargs['_preload_content'] = kwargs.get(
                '_preload_content', True
            )
            kwargs['_request_timeout'] = kwargs.get(
                '_request_timeout', None
            )
            kwargs['_check_input_type'] = kwargs.get(
                '_check_input_type', True
            )
            kwargs['_check_return_type'] = kwargs.get(
                '_check_return_type', True
            )
            kwargs['_host_index'] = kwargs.get('_host_index')
            kwargs['path'] = \
                path
            kwargs['variable'] = \
                variable
            return self.call_with_http_info(**kwargs)

        self.delete_variable = _Endpoint(
            settings={
                'response_type': None,
                'auth': [
                    'X-Nomad-Token'
                ],
                'endpoint_path': '/var/{path}',
                'operation_id': 'delete_variable',
                'http_method': 'DELETE',
                'servers': None,
            },
            params_map={
                'all': [
                    'path',
                    'variable',
                    'region',
                    'namespace',
                    'x_nomad_token',
                    'idempotency_token',
                    'cas',
                ],
                'required': [
                    'path',
                    'variable',
                ],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                ]
            },
            root_map={
                'validations': {
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'path':
                        (str,),
                    'variable':
                        (Variable,),
                    'region':
                        (str,),
                    'namespace':
                        (str,),
                    'x_nomad_token':
                        (str,),
                    'idempotency_token':
                        (str,),
                    'cas':
                        (int,),
                },
                'attribute_map': {
                    'path': 'path',
                    'region': 'region',
                    'namespace': 'namespace',
                    'x_nomad_token': 'X-Nomad-Token',
                    'idempotency_token': 'idempotency_token',
                    'cas': 'cas',
                },
                'location_map': {
                    'path': 'path',
                    'variable': 'body',
                    'region': 'query',
                    'namespace': 'query',
                    'x_nomad_token': 'header',
                    'idempotency_token': 'query',
                    'cas': 'query',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [],
                'content_type': [
                    'application/json'
                ]
            },
            api_client=api_client,
            callable=__delete_variable
        )

        def __get_variable_query(
            self,
            path,
            **kwargs
        ):
            """get_variable_query  # noqa: E501

            This method makes a synchronous HTTP request by default. To make an
            asynchronous HTTP request, please pass async_req=True

            >>> thread = api.get_variable_query(path, async_req=True)
            >>> result = thread.get()

            Args:
                path (str): A path to a Nomad Variable

            Keyword Args:
                region (str): Filters results based on the specified region.. [optional]
                namespace (str): Filters results based on the specified namespace.. [optional]
                index (int): If set, wait until query exceeds given index. Must be provided with WaitParam.. [optional]
                wait (str): Provided with IndexParam to wait for change.. [optional]
                stale (str): If present, results will include stale reads.. [optional]
                prefix (str): Constrains results to jobs that start with the defined prefix. [optional]
                x_nomad_token (str): A Nomad ACL token.. [optional]
                per_page (int): Maximum number of results to return.. [optional]
                next_token (str): Indicates where to start paging for queries that support pagination.. [optional]
                _return_http_data_only (bool): response data without head status
                    code and headers. Default is True.
                _preload_content (bool): if False, the urllib3.HTTPResponse object
                    will be returned without reading/decoding response data.
                    Default is True.
                _request_timeout (int/float/tuple): timeout setting for this request. If
                    one number provided, it will be total request timeout. It can also
                    be a pair (tuple) of (connection, read) timeouts.
                    Default is None.
                _check_input_type (bool): specifies if type checking
                    should be done one the data sent to the server.
                    Default is True.
                _check_return_type (bool): specifies if type checking
                    should be done one the data received from the server.
                    Default is True.
                _host_index (int/None): specifies the index of the server
                    that we want to use.
                    Default is read from the configuration.
                async_req (bool): execute request asynchronously

            Returns:
                Variable
                    If the method is called asynchronously, returns the request
                    thread.
            """
            kwargs['async_req'] = kwargs.get(
                'async_req', False
            )
            kwargs['_return_http_data_only'] = kwargs.get(
                '_return_http_data_only', True
            )
            kwargs['_preload_content'] = kwargs.get(
                '_preload_content', True
            )
            kwargs['_request_timeout'] = kwargs.get(
                '_request_timeout', None
            )
            kwargs['_check_input_type'] = kwargs.get(
                '_check_input_type', True
            )
            kwargs['_check_return_type'] = kwargs.get(
                '_check_return_type', True
            )
            kwargs['_host_index'] = kwargs.get('_host_index')
            kwargs['path'] = \
                path
            return self.call_with_http_info(**kwargs)

        self.get_variable_query = _Endpoint(
            settings={
                'response_type': (Variable,),
                'auth': [
                    'X-Nomad-Token'
                ],
                'endpoint_path': '/var/{path}',
                'operation_id': 'get_variable_query',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'path',
                    'region',
                    'namespace',
                    'index',
                    'wait',
                    'stale',
                    'prefix',
                    'x_nomad_token',
                    'per_page',
                    'next_token',
                ],
                'required': [
                    'path',
                ],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                ]
            },
            root_map={
                'validations': {
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'path':
                        (str,),
                    'region':
                        (str,),
                    'namespace':
                        (str,),
                    'index':
                        (int,),
                    'wait':
                        (str,),
                    'stale':
                        (str,),
                    'prefix':
                        (str,),
                    'x_nomad_token':
                        (str,),
                    'per_page':
                        (int,),
                    'next_token':
                        (str,),
                },
                'attribute_map': {
                    'path': 'path',
                    'region': 'region',
                    'namespace': 'namespace',
                    'index': 'index',
                    'wait': 'wait',
                    'stale': 'stale',
                    'prefix': 'prefix',
                    'x_nomad_token': 'X-Nomad-Token',
                    'per_page': 'per_page',
                    'next_token': 'next_token',
                },
                'location_map': {
                    'path': 'path',
                    'region': 'query',
                    'namespace': 'query',
                    'index': 'header',
                    'wait': 'query',
                    'stale': 'query',
                    'prefix': 'query',
                    'x_nomad_token': 'header',
                    'per_page': 'query',
                    'next_token': 'query',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/json'
                ],
                'content_type': [],
            },
            api_client=api_client,
            callable=__get_variable_query
        )

        def __get_variables_list_request(
            self,
            **kwargs
        ):
            """get_variables_list_request  # noqa: E501

            This method makes a synchronous HTTP request by default. To make an
            asynchronous HTTP request, please pass async_req=True

            >>> thread = api.get_variables_list_request(async_req=True)
            >>> result = thread.get()


            Keyword Args:
                region (str): Filters results based on the specified region.. [optional]
                namespace (str): Filters results based on the specified namespace.. [optional]
                index (int): If set, wait until query exceeds given index. Must be provided with WaitParam.. [optional]
                wait (str): Provided with IndexParam to wait for change.. [optional]
                stale (str): If present, results will include stale reads.. [optional]
                prefix (str): Constrains results to jobs that start with the defined prefix. [optional]
                x_nomad_token (str): A Nomad ACL token.. [optional]
                per_page (int): Maximum number of results to return.. [optional]
                next_token (str): Indicates where to start paging for queries that support pagination.. [optional]
                _return_http_data_only (bool): response data without head status
                    code and headers. Default is True.
                _preload_content (bool): if False, the urllib3.HTTPResponse object
                    will be returned without reading/decoding response data.
                    Default is True.
                _request_timeout (int/float/tuple): timeout setting for this request. If
                    one number provided, it will be total request timeout. It can also
                    be a pair (tuple) of (connection, read) timeouts.
                    Default is None.
                _check_input_type (bool): specifies if type checking
                    should be done one the data sent to the server.
                    Default is True.
                _check_return_type (bool): specifies if type checking
                    should be done one the data received from the server.
                    Default is True.
                _host_index (int/None): specifies the index of the server
                    that we want to use.
                    Default is read from the configuration.
                async_req (bool): execute request asynchronously

            Returns:
                [VariableMetadata]
                    If the method is called asynchronously, returns the request
                    thread.
            """
            kwargs['async_req'] = kwargs.get(
                'async_req', False
            )
            kwargs['_return_http_data_only'] = kwargs.get(
                '_return_http_data_only', True
            )
            kwargs['_preload_content'] = kwargs.get(
                '_preload_content', True
            )
            kwargs['_request_timeout'] = kwargs.get(
                '_request_timeout', None
            )
            kwargs['_check_input_type'] = kwargs.get(
                '_check_input_type', True
            )
            kwargs['_check_return_type'] = kwargs.get(
                '_check_return_type', True
            )
            kwargs['_host_index'] = kwargs.get('_host_index')
            return self.call_with_http_info(**kwargs)

        self.get_variables_list_request = _Endpoint(
            settings={
                'response_type': ([VariableMetadata],),
                'auth': [
                    'X-Nomad-Token'
                ],
                'endpoint_path': '/vars',
                'operation_id': 'get_variables_list_request',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'region',
                    'namespace',
                    'index',
                    'wait',
                    'stale',
                    'prefix',
                    'x_nomad_token',
                    'per_page',
                    'next_token',
                ],
                'required': [],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                ]
            },
            root_map={
                'validations': {
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'region':
                        (str,),
                    'namespace':
                        (str,),
                    'index':
                        (int,),
                    'wait':
                        (str,),
                    'stale':
                        (str,),
                    'prefix':
                        (str,),
                    'x_nomad_token':
                        (str,),
                    'per_page':
                        (int,),
                    'next_token':
                        (str,),
                },
                'attribute_map': {
                    'region': 'region',
                    'namespace': 'namespace',
                    'index': 'index',
                    'wait': 'wait',
                    'stale': 'stale',
                    'prefix': 'prefix',
                    'x_nomad_token': 'X-Nomad-Token',
                    'per_page': 'per_page',
                    'next_token': 'next_token',
                },
                'location_map': {
                    'region': 'query',
                    'namespace': 'query',
                    'index': 'header',
                    'wait': 'query',
                    'stale': 'query',
                    'prefix': 'query',
                    'x_nomad_token': 'header',
                    'per_page': 'query',
                    'next_token': 'query',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/json'
                ],
                'content_type': [],
            },
            api_client=api_client,
            callable=__get_variables_list_request
        )

        def __post_variable(
            self,
            path,
            variable,
            **kwargs
        ):
            """post_variable  # noqa: E501

            This method makes a synchronous HTTP request by default. To make an
            asynchronous HTTP request, please pass async_req=True

            >>> thread = api.post_variable(path, variable, async_req=True)
            >>> result = thread.get()

            Args:
                path (str): A path to a Nomad Variable
                variable (Variable):

            Keyword Args:
                region (str): Filters results based on the specified region.. [optional]
                namespace (str): Filters results based on the specified namespace.. [optional]
                x_nomad_token (str): A Nomad ACL token.. [optional]
                idempotency_token (str): Can be used to ensure operations are only run once.. [optional]
                cas (int): A compare-and-set parameter for Nomad Variables. [optional]
                _return_http_data_only (bool): response data without head status
                    code and headers. Default is True.
                _preload_content (bool): if False, the urllib3.HTTPResponse object
                    will be returned without reading/decoding response data.
                    Default is True.
                _request_timeout (int/float/tuple): timeout setting for this request. If
                    one number provided, it will be total request timeout. It can also
                    be a pair (tuple) of (connection, read) timeouts.
                    Default is None.
                _check_input_type (bool): specifies if type checking
                    should be done one the data sent to the server.
                    Default is True.
                _check_return_type (bool): specifies if type checking
                    should be done one the data received from the server.
                    Default is True.
                _host_index (int/None): specifies the index of the server
                    that we want to use.
                    Default is read from the configuration.
                async_req (bool): execute request asynchronously

            Returns:
                Variable
                    If the method is called asynchronously, returns the request
                    thread.
            """
            kwargs['async_req'] = kwargs.get(
                'async_req', False
            )
            kwargs['_return_http_data_only'] = kwargs.get(
                '_return_http_data_only', True
            )
            kwargs['_preload_content'] = kwargs.get(
                '_preload_content', True
            )
            kwargs['_request_timeout'] = kwargs.get(
                '_request_timeout', None
            )
            kwargs['_check_input_type'] = kwargs.get(
                '_check_input_type', True
            )
            kwargs['_check_return_type'] = kwargs.get(
                '_check_return_type', True
            )
            kwargs['_host_index'] = kwargs.get('_host_index')
            kwargs['path'] = \
                path
            kwargs['variable'] = \
                variable
            return self.call_with_http_info(**kwargs)

        self.post_variable = _Endpoint(
            settings={
                'response_type': (Variable,),
                'auth': [
                    'X-Nomad-Token'
                ],
                'endpoint_path': '/var/{path}',
                'operation_id': 'post_variable',
                'http_method': 'POST',
                'servers': None,
            },
            params_map={
                'all': [
                    'path',
                    'variable',
                    'region',
                    'namespace',
                    'x_nomad_token',
                    'idempotency_token',
                    'cas',
                ],
                'required': [
                    'path',
                    'variable',
                ],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                ]
            },
            root_map={
                'validations': {
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'path':
                        (str,),
                    'variable':
                        (Variable,),
                    'region':
                        (str,),
                    'namespace':
                        (str,),
                    'x_nomad_token':
                        (str,),
                    'idempotency_token':
                        (str,),
                    'cas':
                        (int,),
                },
                'attribute_map': {
                    'path': 'path',
                    'region': 'region',
                    'namespace': 'namespace',
                    'x_nomad_token': 'X-Nomad-Token',
                    'idempotency_token': 'idempotency_token',
                    'cas': 'cas',
                },
                'location_map': {
                    'path': 'path',
                    'variable': 'body',
                    'region': 'query',
                    'namespace': 'query',
                    'x_nomad_token': 'header',
                    'idempotency_token': 'query',
                    'cas': 'query',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/json'
                ],
                'content_type': [
                    'application/json'
                ]
            },
            api_client=api_client,
            callable=__post_variable
        )

        def __put_variable(
            self,
            path,
            variable,
            **kwargs
        ):
            """put_variable  # noqa: E501

            This method makes a synchronous HTTP request by default. To make an
            asynchronous HTTP request, please pass async_req=True

            >>> thread = api.put_variable(path, variable, async_req=True)
            >>> result = thread.get()

            Args:
                path (str): A path to a Nomad Variable
                variable (Variable):

            Keyword Args:
                region (str): Filters results based on the specified region.. [optional]
                namespace (str): Filters results based on the specified namespace.. [optional]
                x_nomad_token (str): A Nomad ACL token.. [optional]
                idempotency_token (str): Can be used to ensure operations are only run once.. [optional]
                cas (int): A compare-and-set parameter for Nomad Variables. [optional]
                _return_http_data_only (bool): response data without head status
                    code and headers. Default is True.
                _preload_content (bool): if False, the urllib3.HTTPResponse object
                    will be returned without reading/decoding response data.
                    Default is True.
                _request_timeout (int/float/tuple): timeout setting for this request. If
                    one number provided, it will be total request timeout. It can also
                    be a pair (tuple) of (connection, read) timeouts.
                    Default is None.
                _check_input_type (bool): specifies if type checking
                    should be done one the data sent to the server.
                    Default is True.
                _check_return_type (bool): specifies if type checking
                    should be done one the data received from the server.
                    Default is True.
                _host_index (int/None): specifies the index of the server
                    that we want to use.
                    Default is read from the configuration.
                async_req (bool): execute request asynchronously

            Returns:
                Variable
                    If the method is called asynchronously, returns the request
                    thread.
            """
            kwargs['async_req'] = kwargs.get(
                'async_req', False
            )
            kwargs['_return_http_data_only'] = kwargs.get(
                '_return_http_data_only', True
            )
            kwargs['_preload_content'] = kwargs.get(
                '_preload_content', True
            )
            kwargs['_request_timeout'] = kwargs.get(
                '_request_timeout', None
            )
            kwargs['_check_input_type'] = kwargs.get(
                '_check_input_type', True
            )
            kwargs['_check_return_type'] = kwargs.get(
                '_check_return_type', True
            )
            kwargs['_host_index'] = kwargs.get('_host_index')
            kwargs['path'] = \
                path
            kwargs['variable'] = \
                variable
            return self.call_with_http_info(**kwargs)

        self.put_variable = _Endpoint(
            settings={
                'response_type': (Variable,),
                'auth': [
                    'X-Nomad-Token'
                ],
                'endpoint_path': '/var/{path}',
                'operation_id': 'put_variable',
                'http_method': 'PUT',
                'servers': None,
            },
            params_map={
                'all': [
                    'path',
                    'variable',
                    'region',
                    'namespace',
                    'x_nomad_token',
                    'idempotency_token',
                    'cas',
                ],
                'required': [
                    'path',
                    'variable',
                ],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                ]
            },
            root_map={
                'validations': {
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'path':
                        (str,),
                    'variable':
                        (Variable,),
                    'region':
                        (str,),
                    'namespace':
                        (str,),
                    'x_nomad_token':
                        (str,),
                    'idempotency_token':
                        (str,),
                    'cas':
                        (int,),
                },
                'attribute_map': {
                    'path': 'path',
                    'region': 'region',
                    'namespace': 'namespace',
                    'x_nomad_token': 'X-Nomad-Token',
                    'idempotency_token': 'idempotency_token',
                    'cas': 'cas',
                },
                'location_map': {
                    'path': 'path',
                    'variable': 'body',
                    'region': 'query',
                    'namespace': 'query',
                    'x_nomad_token': 'header',
                    'idempotency_token': 'query',
                    'cas': 'query',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/json'
                ],
                'content_type': [
                    'application/json'
                ]
            },
            api_client=api_client,
            callable=__put_variable
        )
