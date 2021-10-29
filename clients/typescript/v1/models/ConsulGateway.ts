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

import { AnyType } from './AnyType';
import { ConsulGatewayProxy } from './ConsulGatewayProxy';
import { ConsulIngressConfigEntry } from './ConsulIngressConfigEntry';
import { ConsulTerminatingConfigEntry } from './ConsulTerminatingConfigEntry';
import { HttpFile } from '../http/http';

export class ConsulGateway {
    'ingress'?: ConsulIngressConfigEntry;
    'mesh'?: AnyType;
    'proxy'?: ConsulGatewayProxy;
    'terminating'?: ConsulTerminatingConfigEntry;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "ingress",
            "baseName": "Ingress",
            "type": "ConsulIngressConfigEntry",
            "format": ""
        },
        {
            "name": "mesh",
            "baseName": "Mesh",
            "type": "AnyType",
            "format": ""
        },
        {
            "name": "proxy",
            "baseName": "Proxy",
            "type": "ConsulGatewayProxy",
            "format": ""
        },
        {
            "name": "terminating",
            "baseName": "Terminating",
            "type": "ConsulTerminatingConfigEntry",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return ConsulGateway.attributeTypeMap;
    }
    
    public constructor() {
    }
}
