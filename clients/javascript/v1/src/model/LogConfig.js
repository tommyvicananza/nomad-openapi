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
 * The LogConfig model module.
 * @module model/LogConfig
 * @version 1.1.4
 */
class LogConfig {
    /**
     * Constructs a new <code>LogConfig</code>.
     * @alias module:model/LogConfig
     */
    constructor() { 
        
        LogConfig.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>LogConfig</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/LogConfig} obj Optional instance to populate.
     * @return {module:model/LogConfig} The populated <code>LogConfig</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new LogConfig();

            if (data.hasOwnProperty('Disabled')) {
                obj['Disabled'] = ApiClient.convertToType(data['Disabled'], 'Boolean');
            }
            if (data.hasOwnProperty('Enabled')) {
                obj['Enabled'] = ApiClient.convertToType(data['Enabled'], 'Boolean');
            }
            if (data.hasOwnProperty('MaxFileSizeMB')) {
                obj['MaxFileSizeMB'] = ApiClient.convertToType(data['MaxFileSizeMB'], 'Number');
            }
            if (data.hasOwnProperty('MaxFiles')) {
                obj['MaxFiles'] = ApiClient.convertToType(data['MaxFiles'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * @member {Boolean} Disabled
 */
LogConfig.prototype['Disabled'] = undefined;

/**
 * @member {Boolean} Enabled
 */
LogConfig.prototype['Enabled'] = undefined;

/**
 * @member {Number} MaxFileSizeMB
 */
LogConfig.prototype['MaxFileSizeMB'] = undefined;

/**
 * @member {Number} MaxFiles
 */
LogConfig.prototype['MaxFiles'] = undefined;






export default LogConfig;

