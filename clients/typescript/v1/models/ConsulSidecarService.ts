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

import { ConsulProxy } from './ConsulProxy';
import { HttpFile } from '../http/http';

export class ConsulSidecarService {
    'disableDefaultTCPCheck'?: boolean;
    'meta'?: { [key: string]: string; };
    'port'?: string;
    'proxy'?: ConsulProxy;
    'tags'?: Array<string>;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "disableDefaultTCPCheck",
            "baseName": "DisableDefaultTCPCheck",
            "type": "boolean",
            "format": ""
        },
        {
            "name": "meta",
            "baseName": "Meta",
            "type": "{ [key: string]: string; }",
            "format": ""
        },
        {
            "name": "port",
            "baseName": "Port",
            "type": "string",
            "format": ""
        },
        {
            "name": "proxy",
            "baseName": "Proxy",
            "type": "ConsulProxy",
            "format": ""
        },
        {
            "name": "tags",
            "baseName": "Tags",
            "type": "Array<string>",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return ConsulSidecarService.attributeTypeMap;
    }
    
    public constructor() {
    }
}

