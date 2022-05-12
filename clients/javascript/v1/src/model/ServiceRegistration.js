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

/**
 * The ServiceRegistration model module.
 * @module model/ServiceRegistration
 * @version 1.1.4
 */
class ServiceRegistration {
    /**
     * Constructs a new <code>ServiceRegistration</code>.
     * @alias module:model/ServiceRegistration
     */
    constructor() { 
        
        ServiceRegistration.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ServiceRegistration</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ServiceRegistration} obj Optional instance to populate.
     * @return {module:model/ServiceRegistration} The populated <code>ServiceRegistration</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ServiceRegistration();

            if (data.hasOwnProperty('Address')) {
                obj['Address'] = ApiClient.convertToType(data['Address'], 'String');
            }
            if (data.hasOwnProperty('AllocID')) {
                obj['AllocID'] = ApiClient.convertToType(data['AllocID'], 'String');
            }
            if (data.hasOwnProperty('CreateIndex')) {
                obj['CreateIndex'] = ApiClient.convertToType(data['CreateIndex'], 'Number');
            }
            if (data.hasOwnProperty('Datacenter')) {
                obj['Datacenter'] = ApiClient.convertToType(data['Datacenter'], 'String');
            }
            if (data.hasOwnProperty('ID')) {
                obj['ID'] = ApiClient.convertToType(data['ID'], 'String');
            }
            if (data.hasOwnProperty('JobID')) {
                obj['JobID'] = ApiClient.convertToType(data['JobID'], 'String');
            }
            if (data.hasOwnProperty('ModifyIndex')) {
                obj['ModifyIndex'] = ApiClient.convertToType(data['ModifyIndex'], 'Number');
            }
            if (data.hasOwnProperty('Namespace')) {
                obj['Namespace'] = ApiClient.convertToType(data['Namespace'], 'String');
            }
            if (data.hasOwnProperty('NodeID')) {
                obj['NodeID'] = ApiClient.convertToType(data['NodeID'], 'String');
            }
            if (data.hasOwnProperty('Port')) {
                obj['Port'] = ApiClient.convertToType(data['Port'], 'Number');
            }
            if (data.hasOwnProperty('ServiceName')) {
                obj['ServiceName'] = ApiClient.convertToType(data['ServiceName'], 'String');
            }
            if (data.hasOwnProperty('Tags')) {
                obj['Tags'] = ApiClient.convertToType(data['Tags'], ['String']);
            }
        }
        return obj;
    }


}

/**
 * @member {String} Address
 */
ServiceRegistration.prototype['Address'] = undefined;

/**
 * @member {String} AllocID
 */
ServiceRegistration.prototype['AllocID'] = undefined;

/**
 * @member {Number} CreateIndex
 */
ServiceRegistration.prototype['CreateIndex'] = undefined;

/**
 * @member {String} Datacenter
 */
ServiceRegistration.prototype['Datacenter'] = undefined;

/**
 * @member {String} ID
 */
ServiceRegistration.prototype['ID'] = undefined;

/**
 * @member {String} JobID
 */
ServiceRegistration.prototype['JobID'] = undefined;

/**
 * @member {Number} ModifyIndex
 */
ServiceRegistration.prototype['ModifyIndex'] = undefined;

/**
 * @member {String} Namespace
 */
ServiceRegistration.prototype['Namespace'] = undefined;

/**
 * @member {String} NodeID
 */
ServiceRegistration.prototype['NodeID'] = undefined;

/**
 * @member {Number} Port
 */
ServiceRegistration.prototype['Port'] = undefined;

/**
 * @member {String} ServiceName
 */
ServiceRegistration.prototype['ServiceName'] = undefined;

/**
 * @member {Array.<String>} Tags
 */
ServiceRegistration.prototype['Tags'] = undefined;






export default ServiceRegistration;
