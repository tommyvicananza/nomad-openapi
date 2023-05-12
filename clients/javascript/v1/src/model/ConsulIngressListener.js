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
import ConsulIngressService from './ConsulIngressService';

/**
 * The ConsulIngressListener model module.
 * @module model/ConsulIngressListener
 * @version 1.1.4
 */
class ConsulIngressListener {
    /**
     * Constructs a new <code>ConsulIngressListener</code>.
     * @alias module:model/ConsulIngressListener
     */
    constructor() { 
        
        ConsulIngressListener.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ConsulIngressListener</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ConsulIngressListener} obj Optional instance to populate.
     * @return {module:model/ConsulIngressListener} The populated <code>ConsulIngressListener</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ConsulIngressListener();

            if (data.hasOwnProperty('Port')) {
                obj['Port'] = ApiClient.convertToType(data['Port'], 'Number');
            }
            if (data.hasOwnProperty('Protocol')) {
                obj['Protocol'] = ApiClient.convertToType(data['Protocol'], 'String');
            }
            if (data.hasOwnProperty('Services')) {
                obj['Services'] = ApiClient.convertToType(data['Services'], [ConsulIngressService]);
            }
        }
        return obj;
    }


}

/**
 * @member {Number} Port
 */
ConsulIngressListener.prototype['Port'] = undefined;

/**
 * @member {String} Protocol
 */
ConsulIngressListener.prototype['Protocol'] = undefined;

/**
 * @member {Array.<module:model/ConsulIngressService>} Services
 */
ConsulIngressListener.prototype['Services'] = undefined;






export default ConsulIngressListener;

