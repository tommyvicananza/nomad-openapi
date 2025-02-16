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
import ACLTokenRoleLink from './ACLTokenRoleLink';

/**
 * The ACLTokenListStub model module.
 * @module model/ACLTokenListStub
 * @version 1.1.4
 */
class ACLTokenListStub {
    /**
     * Constructs a new <code>ACLTokenListStub</code>.
     * @alias module:model/ACLTokenListStub
     */
    constructor() { 
        
        ACLTokenListStub.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ACLTokenListStub</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ACLTokenListStub} obj Optional instance to populate.
     * @return {module:model/ACLTokenListStub} The populated <code>ACLTokenListStub</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ACLTokenListStub();

            if (data.hasOwnProperty('AccessorID')) {
                obj['AccessorID'] = ApiClient.convertToType(data['AccessorID'], 'String');
            }
            if (data.hasOwnProperty('CreateIndex')) {
                obj['CreateIndex'] = ApiClient.convertToType(data['CreateIndex'], 'Number');
            }
            if (data.hasOwnProperty('CreateTime')) {
                obj['CreateTime'] = ApiClient.convertToType(data['CreateTime'], 'Date');
            }
            if (data.hasOwnProperty('ExpirationTime')) {
                obj['ExpirationTime'] = ApiClient.convertToType(data['ExpirationTime'], 'Date');
            }
            if (data.hasOwnProperty('Global')) {
                obj['Global'] = ApiClient.convertToType(data['Global'], 'Boolean');
            }
            if (data.hasOwnProperty('ModifyIndex')) {
                obj['ModifyIndex'] = ApiClient.convertToType(data['ModifyIndex'], 'Number');
            }
            if (data.hasOwnProperty('Name')) {
                obj['Name'] = ApiClient.convertToType(data['Name'], 'String');
            }
            if (data.hasOwnProperty('Policies')) {
                obj['Policies'] = ApiClient.convertToType(data['Policies'], ['String']);
            }
            if (data.hasOwnProperty('Roles')) {
                obj['Roles'] = ApiClient.convertToType(data['Roles'], [ACLTokenRoleLink]);
            }
            if (data.hasOwnProperty('Type')) {
                obj['Type'] = ApiClient.convertToType(data['Type'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} AccessorID
 */
ACLTokenListStub.prototype['AccessorID'] = undefined;

/**
 * @member {Number} CreateIndex
 */
ACLTokenListStub.prototype['CreateIndex'] = undefined;

/**
 * @member {Date} CreateTime
 */
ACLTokenListStub.prototype['CreateTime'] = undefined;

/**
 * @member {Date} ExpirationTime
 */
ACLTokenListStub.prototype['ExpirationTime'] = undefined;

/**
 * @member {Boolean} Global
 */
ACLTokenListStub.prototype['Global'] = undefined;

/**
 * @member {Number} ModifyIndex
 */
ACLTokenListStub.prototype['ModifyIndex'] = undefined;

/**
 * @member {String} Name
 */
ACLTokenListStub.prototype['Name'] = undefined;

/**
 * @member {Array.<String>} Policies
 */
ACLTokenListStub.prototype['Policies'] = undefined;

/**
 * @member {Array.<module:model/ACLTokenRoleLink>} Roles
 */
ACLTokenListStub.prototype['Roles'] = undefined;

/**
 * @member {String} Type
 */
ACLTokenListStub.prototype['Type'] = undefined;






export default ACLTokenListStub;

