"""
    Nomad

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)  # noqa: E501

    The version of the OpenAPI document: 1.1.3
    Contact: support@hashicorp.com
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import nomad_client
from nomad_client.model.csi_mount_options import CSIMountOptions
globals()['CSIMountOptions'] = CSIMountOptions
from nomad_client.model.volume_request import VolumeRequest


class TestVolumeRequest(unittest.TestCase):
    """VolumeRequest unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testVolumeRequest(self):
        """Test VolumeRequest"""
        # FIXME: construct object with mandatory attributes with example values
        # model = VolumeRequest()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()