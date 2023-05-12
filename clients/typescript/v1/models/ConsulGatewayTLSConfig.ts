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

export class ConsulGatewayTLSConfig {
    'cipherSuites'?: Array<string>;
    'enabled'?: boolean;
    'tLSMaxVersion'?: string;
    'tLSMinVersion'?: string;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "cipherSuites",
            "baseName": "CipherSuites",
            "type": "Array<string>",
            "format": ""
        },
        {
            "name": "enabled",
            "baseName": "Enabled",
            "type": "boolean",
            "format": ""
        },
        {
            "name": "tLSMaxVersion",
            "baseName": "TLSMaxVersion",
            "type": "string",
            "format": ""
        },
        {
            "name": "tLSMinVersion",
            "baseName": "TLSMinVersion",
            "type": "string",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return ConsulGatewayTLSConfig.attributeTypeMap;
    }
    
    public constructor() {
    }
}

