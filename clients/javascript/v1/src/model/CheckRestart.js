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
 * The CheckRestart model module.
 * @module model/CheckRestart
 * @version 1.1.4
 */
class CheckRestart {
    /**
     * Constructs a new <code>CheckRestart</code>.
     * @alias module:model/CheckRestart
     */
    constructor() { 
        
        CheckRestart.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>CheckRestart</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/CheckRestart} obj Optional instance to populate.
     * @return {module:model/CheckRestart} The populated <code>CheckRestart</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new CheckRestart();

            if (data.hasOwnProperty('Grace')) {
                obj['Grace'] = ApiClient.convertToType(data['Grace'], 'Number');
            }
            if (data.hasOwnProperty('IgnoreWarnings')) {
                obj['IgnoreWarnings'] = ApiClient.convertToType(data['IgnoreWarnings'], 'Boolean');
            }
            if (data.hasOwnProperty('Limit')) {
                obj['Limit'] = ApiClient.convertToType(data['Limit'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} Grace
 */
CheckRestart.prototype['Grace'] = undefined;

/**
 * @member {Boolean} IgnoreWarnings
 */
CheckRestart.prototype['IgnoreWarnings'] = undefined;

/**
 * @member {Number} Limit
 */
CheckRestart.prototype['Limit'] = undefined;






export default CheckRestart;
