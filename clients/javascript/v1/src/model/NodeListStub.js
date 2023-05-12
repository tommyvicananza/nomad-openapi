/**
 * Nomad
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.1.4
 * Contact: support@hashicorp.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
import DrainMetadata from './DrainMetadata';
import DriverInfo from './DriverInfo';
import NodeReservedResources from './NodeReservedResources';
import NodeResources from './NodeResources';

/**
 * The NodeListStub model module.
 * @module model/NodeListStub
 * @version 1.1.4
 */
class NodeListStub {
    /**
     * Constructs a new <code>NodeListStub</code>.
     * @alias module:model/NodeListStub
     */
    constructor() { 
        
        NodeListStub.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>NodeListStub</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/NodeListStub} obj Optional instance to populate.
     * @return {module:model/NodeListStub} The populated <code>NodeListStub</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new NodeListStub();

            if (data.hasOwnProperty('Address')) {
                obj['Address'] = ApiClient.convertToType(data['Address'], 'String');
            }
            if (data.hasOwnProperty('Attributes')) {
                obj['Attributes'] = ApiClient.convertToType(data['Attributes'], {'String': 'String'});
            }
            if (data.hasOwnProperty('CreateIndex')) {
                obj['CreateIndex'] = ApiClient.convertToType(data['CreateIndex'], 'Number');
            }
            if (data.hasOwnProperty('Datacenter')) {
                obj['Datacenter'] = ApiClient.convertToType(data['Datacenter'], 'String');
            }
            if (data.hasOwnProperty('Drain')) {
                obj['Drain'] = ApiClient.convertToType(data['Drain'], 'Boolean');
            }
            if (data.hasOwnProperty('Drivers')) {
                obj['Drivers'] = ApiClient.convertToType(data['Drivers'], {'String': DriverInfo});
            }
            if (data.hasOwnProperty('ID')) {
                obj['ID'] = ApiClient.convertToType(data['ID'], 'String');
            }
            if (data.hasOwnProperty('LastDrain')) {
                obj['LastDrain'] = DrainMetadata.constructFromObject(data['LastDrain']);
            }
            if (data.hasOwnProperty('ModifyIndex')) {
                obj['ModifyIndex'] = ApiClient.convertToType(data['ModifyIndex'], 'Number');
            }
            if (data.hasOwnProperty('Name')) {
                obj['Name'] = ApiClient.convertToType(data['Name'], 'String');
            }
            if (data.hasOwnProperty('NodeClass')) {
                obj['NodeClass'] = ApiClient.convertToType(data['NodeClass'], 'String');
            }
            if (data.hasOwnProperty('NodeResources')) {
                obj['NodeResources'] = NodeResources.constructFromObject(data['NodeResources']);
            }
            if (data.hasOwnProperty('ReservedResources')) {
                obj['ReservedResources'] = NodeReservedResources.constructFromObject(data['ReservedResources']);
            }
            if (data.hasOwnProperty('SchedulingEligibility')) {
                obj['SchedulingEligibility'] = ApiClient.convertToType(data['SchedulingEligibility'], 'String');
            }
            if (data.hasOwnProperty('Status')) {
                obj['Status'] = ApiClient.convertToType(data['Status'], 'String');
            }
            if (data.hasOwnProperty('StatusDescription')) {
                obj['StatusDescription'] = ApiClient.convertToType(data['StatusDescription'], 'String');
            }
            if (data.hasOwnProperty('Version')) {
                obj['Version'] = ApiClient.convertToType(data['Version'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} Address
 */
NodeListStub.prototype['Address'] = undefined;

/**
 * @member {Object.<String, String>} Attributes
 */
NodeListStub.prototype['Attributes'] = undefined;

/**
 * @member {Number} CreateIndex
 */
NodeListStub.prototype['CreateIndex'] = undefined;

/**
 * @member {String} Datacenter
 */
NodeListStub.prototype['Datacenter'] = undefined;

/**
 * @member {Boolean} Drain
 */
NodeListStub.prototype['Drain'] = undefined;

/**
 * @member {Object.<String, module:model/DriverInfo>} Drivers
 */
NodeListStub.prototype['Drivers'] = undefined;

/**
 * @member {String} ID
 */
NodeListStub.prototype['ID'] = undefined;

/**
 * @member {module:model/DrainMetadata} LastDrain
 */
NodeListStub.prototype['LastDrain'] = undefined;

/**
 * @member {Number} ModifyIndex
 */
NodeListStub.prototype['ModifyIndex'] = undefined;

/**
 * @member {String} Name
 */
NodeListStub.prototype['Name'] = undefined;

/**
 * @member {String} NodeClass
 */
NodeListStub.prototype['NodeClass'] = undefined;

/**
 * @member {module:model/NodeResources} NodeResources
 */
NodeListStub.prototype['NodeResources'] = undefined;

/**
 * @member {module:model/NodeReservedResources} ReservedResources
 */
NodeListStub.prototype['ReservedResources'] = undefined;

/**
 * @member {String} SchedulingEligibility
 */
NodeListStub.prototype['SchedulingEligibility'] = undefined;

/**
 * @member {String} Status
 */
NodeListStub.prototype['Status'] = undefined;

/**
 * @member {String} StatusDescription
 */
NodeListStub.prototype['StatusDescription'] = undefined;

/**
 * @member {String} Version
 */
NodeListStub.prototype['Version'] = undefined;






export default NodeListStub;

