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
import ConsulExposeConfig from './ConsulExposeConfig';
import ConsulUpstream from './ConsulUpstream';

/**
 * The ConsulProxy model module.
 * @module model/ConsulProxy
 * @version 1.1.4
 */
class ConsulProxy {
    /**
     * Constructs a new <code>ConsulProxy</code>.
     * @alias module:model/ConsulProxy
     */
    constructor() { 
        
        ConsulProxy.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ConsulProxy</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ConsulProxy} obj Optional instance to populate.
     * @return {module:model/ConsulProxy} The populated <code>ConsulProxy</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ConsulProxy();

            if (data.hasOwnProperty('Config')) {
                obj['Config'] = ApiClient.convertToType(data['Config'], {'String': Object});
            }
            if (data.hasOwnProperty('ExposeConfig')) {
                obj['ExposeConfig'] = ConsulExposeConfig.constructFromObject(data['ExposeConfig']);
            }
            if (data.hasOwnProperty('LocalServiceAddress')) {
                obj['LocalServiceAddress'] = ApiClient.convertToType(data['LocalServiceAddress'], 'String');
            }
            if (data.hasOwnProperty('LocalServicePort')) {
                obj['LocalServicePort'] = ApiClient.convertToType(data['LocalServicePort'], 'Number');
            }
            if (data.hasOwnProperty('Upstreams')) {
                obj['Upstreams'] = ApiClient.convertToType(data['Upstreams'], [ConsulUpstream]);
            }
        }
        return obj;
    }


}

/**
 * @member {Object.<String, Object>} Config
 */
ConsulProxy.prototype['Config'] = undefined;

/**
 * @member {module:model/ConsulExposeConfig} ExposeConfig
 */
ConsulProxy.prototype['ExposeConfig'] = undefined;

/**
 * @member {String} LocalServiceAddress
 */
ConsulProxy.prototype['LocalServiceAddress'] = undefined;

/**
 * @member {Number} LocalServicePort
 */
ConsulProxy.prototype['LocalServicePort'] = undefined;

/**
 * @member {Array.<module:model/ConsulUpstream>} Upstreams
 */
ConsulProxy.prototype['Upstreams'] = undefined;






export default ConsulProxy;
