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

import { JobSummary } from './JobSummary';
import { HttpFile } from '../http/http';

export class JobListStub {
    'createIndex'?: number;
    'datacenters'?: Array<string>;
    'ID'?: string;
    'jobModifyIndex'?: number;
    'jobSummary'?: JobSummary;
    'meta'?: { [key: string]: string; };
    'modifyIndex'?: number;
    'name'?: string;
    'namespace'?: string;
    'parameterizedJob'?: boolean;
    'parentID'?: string;
    'periodic'?: boolean;
    'priority'?: number;
    'status'?: string;
    'statusDescription'?: string;
    'stop'?: boolean;
    'submitTime'?: number;
    'type'?: string;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "createIndex",
            "baseName": "CreateIndex",
            "type": "number",
            "format": ""
        },
        {
            "name": "datacenters",
            "baseName": "Datacenters",
            "type": "Array<string>",
            "format": ""
        },
        {
            "name": "ID",
            "baseName": "ID",
            "type": "string",
            "format": ""
        },
        {
            "name": "jobModifyIndex",
            "baseName": "JobModifyIndex",
            "type": "number",
            "format": ""
        },
        {
            "name": "jobSummary",
            "baseName": "JobSummary",
            "type": "JobSummary",
            "format": ""
        },
        {
            "name": "meta",
            "baseName": "Meta",
            "type": "{ [key: string]: string; }",
            "format": ""
        },
        {
            "name": "modifyIndex",
            "baseName": "ModifyIndex",
            "type": "number",
            "format": ""
        },
        {
            "name": "name",
            "baseName": "Name",
            "type": "string",
            "format": ""
        },
        {
            "name": "namespace",
            "baseName": "Namespace",
            "type": "string",
            "format": ""
        },
        {
            "name": "parameterizedJob",
            "baseName": "ParameterizedJob",
            "type": "boolean",
            "format": ""
        },
        {
            "name": "parentID",
            "baseName": "ParentID",
            "type": "string",
            "format": ""
        },
        {
            "name": "periodic",
            "baseName": "Periodic",
            "type": "boolean",
            "format": ""
        },
        {
            "name": "priority",
            "baseName": "Priority",
            "type": "number",
            "format": ""
        },
        {
            "name": "status",
            "baseName": "Status",
            "type": "string",
            "format": ""
        },
        {
            "name": "statusDescription",
            "baseName": "StatusDescription",
            "type": "string",
            "format": ""
        },
        {
            "name": "stop",
            "baseName": "Stop",
            "type": "boolean",
            "format": ""
        },
        {
            "name": "submitTime",
            "baseName": "SubmitTime",
            "type": "number",
            "format": "int64"
        },
        {
            "name": "type",
            "baseName": "Type",
            "type": "string",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return JobListStub.attributeTypeMap;
    }
    
    public constructor() {
    }
}

