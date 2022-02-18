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
import RaftServer from './RaftServer';

/**
 * The RaftConfiguration model module.
 * @module model/RaftConfiguration
 * @version 1.1.4
 */
class RaftConfiguration {
    /**
     * Constructs a new <code>RaftConfiguration</code>.
     * @alias module:model/RaftConfiguration
     */
    constructor() { 
        
        RaftConfiguration.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>RaftConfiguration</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/RaftConfiguration} obj Optional instance to populate.
     * @return {module:model/RaftConfiguration} The populated <code>RaftConfiguration</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new RaftConfiguration();

            if (data.hasOwnProperty('Index')) {
                obj['Index'] = ApiClient.convertToType(data['Index'], 'Number');
            }
            if (data.hasOwnProperty('Servers')) {
                obj['Servers'] = ApiClient.convertToType(data['Servers'], [RaftServer]);
            }
        }
        return obj;
    }


}

/**
 * @member {Number} Index
 */
RaftConfiguration.prototype['Index'] = undefined;

/**
 * @member {Array.<module:model/RaftServer>} Servers
 */
RaftConfiguration.prototype['Servers'] = undefined;






export default RaftConfiguration;
