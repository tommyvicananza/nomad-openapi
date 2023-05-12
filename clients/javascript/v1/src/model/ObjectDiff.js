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
import FieldDiff from './FieldDiff';

/**
 * The ObjectDiff model module.
 * @module model/ObjectDiff
 * @version 1.1.4
 */
class ObjectDiff {
    /**
     * Constructs a new <code>ObjectDiff</code>.
     * @alias module:model/ObjectDiff
     */
    constructor() { 
        
        ObjectDiff.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ObjectDiff</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ObjectDiff} obj Optional instance to populate.
     * @return {module:model/ObjectDiff} The populated <code>ObjectDiff</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ObjectDiff();

            if (data.hasOwnProperty('Fields')) {
                obj['Fields'] = ApiClient.convertToType(data['Fields'], [FieldDiff]);
            }
            if (data.hasOwnProperty('Name')) {
                obj['Name'] = ApiClient.convertToType(data['Name'], 'String');
            }
            if (data.hasOwnProperty('Objects')) {
                obj['Objects'] = ApiClient.convertToType(data['Objects'], [ObjectDiff]);
            }
            if (data.hasOwnProperty('Type')) {
                obj['Type'] = ApiClient.convertToType(data['Type'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {Array.<module:model/FieldDiff>} Fields
 */
ObjectDiff.prototype['Fields'] = undefined;

/**
 * @member {String} Name
 */
ObjectDiff.prototype['Name'] = undefined;

/**
 * @member {Array.<module:model/ObjectDiff>} Objects
 */
ObjectDiff.prototype['Objects'] = undefined;

/**
 * @member {String} Type
 */
ObjectDiff.prototype['Type'] = undefined;






export default ObjectDiff;

