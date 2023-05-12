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

import { HttpFile } from '../http/http';

export class JobACL {
    'group'?: string;
    'jobID'?: string;
    'namespace'?: string;
    'task'?: string;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "group",
            "baseName": "Group",
            "type": "string",
            "format": ""
        },
        {
            "name": "jobID",
            "baseName": "JobID",
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
            "name": "task",
            "baseName": "Task",
            "type": "string",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return JobACL.attributeTypeMap;
    }
    
    public constructor() {
    }
}

