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

import { ConsulGateway } from './ConsulGateway';
import { ConsulSidecarService } from './ConsulSidecarService';
import { SidecarTask } from './SidecarTask';
import { HttpFile } from '../http/http';

export class ConsulConnect {
    'gateway'?: ConsulGateway;
    '_native'?: boolean;
    'sidecarService'?: ConsulSidecarService;
    'sidecarTask'?: SidecarTask;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "gateway",
            "baseName": "Gateway",
            "type": "ConsulGateway",
            "format": ""
        },
        {
            "name": "_native",
            "baseName": "Native",
            "type": "boolean",
            "format": ""
        },
        {
            "name": "sidecarService",
            "baseName": "SidecarService",
            "type": "ConsulSidecarService",
            "format": ""
        },
        {
            "name": "sidecarTask",
            "baseName": "SidecarTask",
            "type": "SidecarTask",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return ConsulConnect.attributeTypeMap;
    }
    
    public constructor() {
    }
}

