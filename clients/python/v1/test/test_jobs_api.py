# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

"""
    Nomad

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)  # noqa: E501

    The version of the OpenAPI document: 1.1.3
    Contact: support@hashicorp.com
    Generated by: https://openapi-generator.tech
"""


import unittest

import nomad_client
from nomad_client.api.jobs_api import JobsApi  # noqa: E501


class TestJobsApi(unittest.TestCase):
    """JobsApi unit test stubs"""

    def setUp(self):
        self.api = JobsApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_delete_job(self):
        """Test case for delete_job

        """
        pass

    def test_get_job(self):
        """Test case for get_job

        """
        pass

    def test_get_job_allocations(self):
        """Test case for get_job_allocations

        """
        pass

    def test_get_job_deployment(self):
        """Test case for get_job_deployment

        """
        pass

    def test_get_job_deployments(self):
        """Test case for get_job_deployments

        """
        pass

    def test_get_job_evaluations(self):
        """Test case for get_job_evaluations

        """
        pass

    def test_get_job_scale_status(self):
        """Test case for get_job_scale_status

        """
        pass

    def test_get_job_summary(self):
        """Test case for get_job_summary

        """
        pass

    def test_get_job_versions(self):
        """Test case for get_job_versions

        """
        pass

    def test_get_jobs(self):
        """Test case for get_jobs

        """
        pass

    def test_post_job(self):
        """Test case for post_job

        """
        pass

    def test_post_job_dispatch(self):
        """Test case for post_job_dispatch

        """
        pass

    def test_post_job_evaluate(self):
        """Test case for post_job_evaluate

        """
        pass

    def test_post_job_parse(self):
        """Test case for post_job_parse

        """
        pass

    def test_post_job_periodic_force(self):
        """Test case for post_job_periodic_force

        """
        pass

    def test_post_job_plan(self):
        """Test case for post_job_plan

        """
        pass

    def test_post_job_revert(self):
        """Test case for post_job_revert

        """
        pass

    def test_post_job_scaling_request(self):
        """Test case for post_job_scaling_request

        """
        pass

    def test_post_job_stability(self):
        """Test case for post_job_stability

        """
        pass

    def test_post_job_validate_request(self):
        """Test case for post_job_validate_request

        """
        pass

    def test_register_job(self):
        """Test case for register_job

        """
        pass


if __name__ == '__main__':
    unittest.main()
