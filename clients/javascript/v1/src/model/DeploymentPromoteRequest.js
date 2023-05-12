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
 * The DeploymentPromoteRequest model module.
 * @module model/DeploymentPromoteRequest
 * @version 1.1.4
 */
class DeploymentPromoteRequest {
    /**
     * Constructs a new <code>DeploymentPromoteRequest</code>.
     * @alias module:model/DeploymentPromoteRequest
     */
    constructor() { 
        
        DeploymentPromoteRequest.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>DeploymentPromoteRequest</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/DeploymentPromoteRequest} obj Optional instance to populate.
     * @return {module:model/DeploymentPromoteRequest} The populated <code>DeploymentPromoteRequest</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new DeploymentPromoteRequest();

            if (data.hasOwnProperty('All')) {
                obj['All'] = ApiClient.convertToType(data['All'], 'Boolean');
            }
            if (data.hasOwnProperty('DeploymentID')) {
                obj['DeploymentID'] = ApiClient.convertToType(data['DeploymentID'], 'String');
            }
            if (data.hasOwnProperty('Groups')) {
                obj['Groups'] = ApiClient.convertToType(data['Groups'], ['String']);
            }
            if (data.hasOwnProperty('Namespace')) {
                obj['Namespace'] = ApiClient.convertToType(data['Namespace'], 'String');
            }
            if (data.hasOwnProperty('Region')) {
                obj['Region'] = ApiClient.convertToType(data['Region'], 'String');
            }
            if (data.hasOwnProperty('SecretID')) {
                obj['SecretID'] = ApiClient.convertToType(data['SecretID'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {Boolean} All
 */
DeploymentPromoteRequest.prototype['All'] = undefined;

/**
 * @member {String} DeploymentID
 */
DeploymentPromoteRequest.prototype['DeploymentID'] = undefined;

/**
 * @member {Array.<String>} Groups
 */
DeploymentPromoteRequest.prototype['Groups'] = undefined;

/**
 * @member {String} Namespace
 */
DeploymentPromoteRequest.prototype['Namespace'] = undefined;

/**
 * @member {String} Region
 */
DeploymentPromoteRequest.prototype['Region'] = undefined;

/**
 * @member {String} SecretID
 */
DeploymentPromoteRequest.prototype['SecretID'] = undefined;






export default DeploymentPromoteRequest;

