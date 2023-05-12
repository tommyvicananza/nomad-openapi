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
import CSITopology from './CSITopology';

/**
 * The CSIVolumeListStub model module.
 * @module model/CSIVolumeListStub
 * @version 1.1.4
 */
class CSIVolumeListStub {
    /**
     * Constructs a new <code>CSIVolumeListStub</code>.
     * @alias module:model/CSIVolumeListStub
     */
    constructor() { 
        
        CSIVolumeListStub.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>CSIVolumeListStub</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/CSIVolumeListStub} obj Optional instance to populate.
     * @return {module:model/CSIVolumeListStub} The populated <code>CSIVolumeListStub</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new CSIVolumeListStub();

            if (data.hasOwnProperty('AccessMode')) {
                obj['AccessMode'] = ApiClient.convertToType(data['AccessMode'], 'String');
            }
            if (data.hasOwnProperty('AttachmentMode')) {
                obj['AttachmentMode'] = ApiClient.convertToType(data['AttachmentMode'], 'String');
            }
            if (data.hasOwnProperty('ControllerRequired')) {
                obj['ControllerRequired'] = ApiClient.convertToType(data['ControllerRequired'], 'Boolean');
            }
            if (data.hasOwnProperty('ControllersExpected')) {
                obj['ControllersExpected'] = ApiClient.convertToType(data['ControllersExpected'], 'Number');
            }
            if (data.hasOwnProperty('ControllersHealthy')) {
                obj['ControllersHealthy'] = ApiClient.convertToType(data['ControllersHealthy'], 'Number');
            }
            if (data.hasOwnProperty('CreateIndex')) {
                obj['CreateIndex'] = ApiClient.convertToType(data['CreateIndex'], 'Number');
            }
            if (data.hasOwnProperty('CurrentReaders')) {
                obj['CurrentReaders'] = ApiClient.convertToType(data['CurrentReaders'], 'Number');
            }
            if (data.hasOwnProperty('CurrentWriters')) {
                obj['CurrentWriters'] = ApiClient.convertToType(data['CurrentWriters'], 'Number');
            }
            if (data.hasOwnProperty('ExternalID')) {
                obj['ExternalID'] = ApiClient.convertToType(data['ExternalID'], 'String');
            }
            if (data.hasOwnProperty('ID')) {
                obj['ID'] = ApiClient.convertToType(data['ID'], 'String');
            }
            if (data.hasOwnProperty('ModifyIndex')) {
                obj['ModifyIndex'] = ApiClient.convertToType(data['ModifyIndex'], 'Number');
            }
            if (data.hasOwnProperty('Name')) {
                obj['Name'] = ApiClient.convertToType(data['Name'], 'String');
            }
            if (data.hasOwnProperty('Namespace')) {
                obj['Namespace'] = ApiClient.convertToType(data['Namespace'], 'String');
            }
            if (data.hasOwnProperty('NodesExpected')) {
                obj['NodesExpected'] = ApiClient.convertToType(data['NodesExpected'], 'Number');
            }
            if (data.hasOwnProperty('NodesHealthy')) {
                obj['NodesHealthy'] = ApiClient.convertToType(data['NodesHealthy'], 'Number');
            }
            if (data.hasOwnProperty('PluginID')) {
                obj['PluginID'] = ApiClient.convertToType(data['PluginID'], 'String');
            }
            if (data.hasOwnProperty('Provider')) {
                obj['Provider'] = ApiClient.convertToType(data['Provider'], 'String');
            }
            if (data.hasOwnProperty('ResourceExhausted')) {
                obj['ResourceExhausted'] = ApiClient.convertToType(data['ResourceExhausted'], 'Date');
            }
            if (data.hasOwnProperty('Schedulable')) {
                obj['Schedulable'] = ApiClient.convertToType(data['Schedulable'], 'Boolean');
            }
            if (data.hasOwnProperty('Topologies')) {
                obj['Topologies'] = ApiClient.convertToType(data['Topologies'], [CSITopology]);
            }
        }
        return obj;
    }


}

/**
 * @member {String} AccessMode
 */
CSIVolumeListStub.prototype['AccessMode'] = undefined;

/**
 * @member {String} AttachmentMode
 */
CSIVolumeListStub.prototype['AttachmentMode'] = undefined;

/**
 * @member {Boolean} ControllerRequired
 */
CSIVolumeListStub.prototype['ControllerRequired'] = undefined;

/**
 * @member {Number} ControllersExpected
 */
CSIVolumeListStub.prototype['ControllersExpected'] = undefined;

/**
 * @member {Number} ControllersHealthy
 */
CSIVolumeListStub.prototype['ControllersHealthy'] = undefined;

/**
 * @member {Number} CreateIndex
 */
CSIVolumeListStub.prototype['CreateIndex'] = undefined;

/**
 * @member {Number} CurrentReaders
 */
CSIVolumeListStub.prototype['CurrentReaders'] = undefined;

/**
 * @member {Number} CurrentWriters
 */
CSIVolumeListStub.prototype['CurrentWriters'] = undefined;

/**
 * @member {String} ExternalID
 */
CSIVolumeListStub.prototype['ExternalID'] = undefined;

/**
 * @member {String} ID
 */
CSIVolumeListStub.prototype['ID'] = undefined;

/**
 * @member {Number} ModifyIndex
 */
CSIVolumeListStub.prototype['ModifyIndex'] = undefined;

/**
 * @member {String} Name
 */
CSIVolumeListStub.prototype['Name'] = undefined;

/**
 * @member {String} Namespace
 */
CSIVolumeListStub.prototype['Namespace'] = undefined;

/**
 * @member {Number} NodesExpected
 */
CSIVolumeListStub.prototype['NodesExpected'] = undefined;

/**
 * @member {Number} NodesHealthy
 */
CSIVolumeListStub.prototype['NodesHealthy'] = undefined;

/**
 * @member {String} PluginID
 */
CSIVolumeListStub.prototype['PluginID'] = undefined;

/**
 * @member {String} Provider
 */
CSIVolumeListStub.prototype['Provider'] = undefined;

/**
 * @member {Date} ResourceExhausted
 */
CSIVolumeListStub.prototype['ResourceExhausted'] = undefined;

/**
 * @member {Boolean} Schedulable
 */
CSIVolumeListStub.prototype['Schedulable'] = undefined;

/**
 * @member {Array.<module:model/CSITopology>} Topologies
 */
CSIVolumeListStub.prototype['Topologies'] = undefined;






export default CSIVolumeListStub;

