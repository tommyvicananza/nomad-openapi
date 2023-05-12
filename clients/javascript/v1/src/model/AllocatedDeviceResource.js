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
 * The AllocatedDeviceResource model module.
 * @module model/AllocatedDeviceResource
 * @version 1.1.4
 */
class AllocatedDeviceResource {
    /**
     * Constructs a new <code>AllocatedDeviceResource</code>.
     * @alias module:model/AllocatedDeviceResource
     */
    constructor() { 
        
        AllocatedDeviceResource.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>AllocatedDeviceResource</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/AllocatedDeviceResource} obj Optional instance to populate.
     * @return {module:model/AllocatedDeviceResource} The populated <code>AllocatedDeviceResource</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new AllocatedDeviceResource();

            if (data.hasOwnProperty('DeviceIDs')) {
                obj['DeviceIDs'] = ApiClient.convertToType(data['DeviceIDs'], ['String']);
            }
            if (data.hasOwnProperty('Name')) {
                obj['Name'] = ApiClient.convertToType(data['Name'], 'String');
            }
            if (data.hasOwnProperty('Type')) {
                obj['Type'] = ApiClient.convertToType(data['Type'], 'String');
            }
            if (data.hasOwnProperty('Vendor')) {
                obj['Vendor'] = ApiClient.convertToType(data['Vendor'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {Array.<String>} DeviceIDs
 */
AllocatedDeviceResource.prototype['DeviceIDs'] = undefined;

/**
 * @member {String} Name
 */
AllocatedDeviceResource.prototype['Name'] = undefined;

/**
 * @member {String} Type
 */
AllocatedDeviceResource.prototype['Type'] = undefined;

/**
 * @member {String} Vendor
 */
AllocatedDeviceResource.prototype['Vendor'] = undefined;






export default AllocatedDeviceResource;

