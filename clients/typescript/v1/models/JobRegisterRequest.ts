/**
 * Nomad
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * OpenAPI spec version: 1.1.4
 * Contact: support@hashicorp.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { Job } from './Job';
import { JobSubmission } from './JobSubmission';
import { HttpFile } from '../http/http';

export class JobRegisterRequest {
    'enforceIndex'?: boolean;
    'evalPriority'?: number;
    'job'?: Job;
    'jobModifyIndex'?: number;
    'namespace'?: string;
    'policyOverride'?: boolean;
    'preserveCounts'?: boolean;
    'region'?: string;
    'secretID'?: string;
    'submission'?: JobSubmission;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "enforceIndex",
            "baseName": "EnforceIndex",
            "type": "boolean",
            "format": ""
        },
        {
            "name": "evalPriority",
            "baseName": "EvalPriority",
            "type": "number",
            "format": ""
        },
        {
            "name": "job",
            "baseName": "Job",
            "type": "Job",
            "format": ""
        },
        {
            "name": "jobModifyIndex",
            "baseName": "JobModifyIndex",
            "type": "number",
            "format": ""
        },
        {
            "name": "namespace",
            "baseName": "Namespace",
            "type": "string",
            "format": ""
        },
        {
            "name": "policyOverride",
            "baseName": "PolicyOverride",
            "type": "boolean",
            "format": ""
        },
        {
            "name": "preserveCounts",
            "baseName": "PreserveCounts",
            "type": "boolean",
            "format": ""
        },
        {
            "name": "region",
            "baseName": "Region",
            "type": "string",
            "format": ""
        },
        {
            "name": "secretID",
            "baseName": "SecretID",
            "type": "string",
            "format": ""
        },
        {
            "name": "submission",
            "baseName": "Submission",
            "type": "JobSubmission",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return JobRegisterRequest.attributeTypeMap;
    }
    
    public constructor() {
    }
}

