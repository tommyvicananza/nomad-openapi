"""
    Nomad

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)  # noqa: E501

    The version of the OpenAPI document: 1.1.4
    Contact: support@hashicorp.com
    Generated by: https://openapi-generator.tech
"""


import unittest

import nomad_client
from nomad_client.api.variables_api import VariablesApi  # noqa: E501


class TestVariablesApi(unittest.TestCase):
    """VariablesApi unit test stubs"""

    def setUp(self):
        self.api = VariablesApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_delete_variable(self):
        """Test case for delete_variable

        """
        pass

    def test_get_variable_query(self):
        """Test case for get_variable_query

        """
        pass

    def test_get_variables_list_request(self):
        """Test case for get_variables_list_request

        """
        pass

    def test_post_variable(self):
        """Test case for post_variable

        """
        pass

    def test_put_variable(self):
        """Test case for put_variable

        """
        pass


if __name__ == '__main__':
    unittest.main()