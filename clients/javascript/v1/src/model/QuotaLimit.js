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
import Resources from './Resources';

/**
 * The QuotaLimit model module.
 * @module model/QuotaLimit
 * @version 1.1.4
 */
class QuotaLimit {
    /**
     * Constructs a new <code>QuotaLimit</code>.
     * @alias module:model/QuotaLimit
     */
    constructor() { 
        
        QuotaLimit.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>QuotaLimit</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/QuotaLimit} obj Optional instance to populate.
     * @return {module:model/QuotaLimit} The populated <code>QuotaLimit</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new QuotaLimit();

            if (data.hasOwnProperty('Hash')) {
                obj['Hash'] = ApiClient.convertToType(data['Hash'], 'Blob');
            }
            if (data.hasOwnProperty('Region')) {
                obj['Region'] = ApiClient.convertToType(data['Region'], 'String');
            }
            if (data.hasOwnProperty('RegionLimit')) {
                obj['RegionLimit'] = Resources.constructFromObject(data['RegionLimit']);
            }
        }
        return obj;
    }


}

/**
 * @member {Blob} Hash
 */
QuotaLimit.prototype['Hash'] = undefined;

/**
 * @member {String} Region
 */
QuotaLimit.prototype['Region'] = undefined;

/**
 * @member {module:model/Resources} RegionLimit
 */
QuotaLimit.prototype['RegionLimit'] = undefined;






export default QuotaLimit;
