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
 * The AllocDeploymentStatus model module.
 * @module model/AllocDeploymentStatus
 * @version 1.1.4
 */
class AllocDeploymentStatus {
    /**
     * Constructs a new <code>AllocDeploymentStatus</code>.
     * @alias module:model/AllocDeploymentStatus
     */
    constructor() { 
        
        AllocDeploymentStatus.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>AllocDeploymentStatus</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/AllocDeploymentStatus} obj Optional instance to populate.
     * @return {module:model/AllocDeploymentStatus} The populated <code>AllocDeploymentStatus</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new AllocDeploymentStatus();

            if (data.hasOwnProperty('Canary')) {
                obj['Canary'] = ApiClient.convertToType(data['Canary'], 'Boolean');
            }
            if (data.hasOwnProperty('Healthy')) {
                obj['Healthy'] = ApiClient.convertToType(data['Healthy'], 'Boolean');
            }
            if (data.hasOwnProperty('ModifyIndex')) {
                obj['ModifyIndex'] = ApiClient.convertToType(data['ModifyIndex'], 'Number');
            }
            if (data.hasOwnProperty('Timestamp')) {
                obj['Timestamp'] = ApiClient.convertToType(data['Timestamp'], 'Date');
            }
        }
        return obj;
    }


}

/**
 * @member {Boolean} Canary
 */
AllocDeploymentStatus.prototype['Canary'] = undefined;

/**
 * @member {Boolean} Healthy
 */
AllocDeploymentStatus.prototype['Healthy'] = undefined;

/**
 * @member {Number} ModifyIndex
 */
AllocDeploymentStatus.prototype['ModifyIndex'] = undefined;

/**
 * @member {Date} Timestamp
 */
AllocDeploymentStatus.prototype['Timestamp'] = undefined;






export default AllocDeploymentStatus;

