/**
 * Copyright (c) HashiCorp, Inc.
 * SPDX-License-Identifier: MPL-2.0
 */

/**
 * Nomad
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.1.3
 * Contact: support@hashicorp.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', process.cwd()+'/src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require(process.cwd()+'/src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.nomad-client);
  }
}(this, function(expect, nomad-client) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new nomad-client.VolumeRequest();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('VolumeRequest', function() {
    it('should create an instance of VolumeRequest', function() {
      // uncomment below and update the code to test VolumeRequest
      //var instane = new nomad-client.VolumeRequest();
      //expect(instance).to.be.a(nomad-client.VolumeRequest);
    });

    it('should have the property accessMode (base name: "AccessMode")', function() {
      // uncomment below and update the code to test the property accessMode
      //var instance = new nomad-client.VolumeRequest();
      //expect(instance).to.be();
    });

    it('should have the property attachmentMode (base name: "AttachmentMode")', function() {
      // uncomment below and update the code to test the property attachmentMode
      //var instance = new nomad-client.VolumeRequest();
      //expect(instance).to.be();
    });

    it('should have the property mountOptions (base name: "MountOptions")', function() {
      // uncomment below and update the code to test the property mountOptions
      //var instance = new nomad-client.VolumeRequest();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "Name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new nomad-client.VolumeRequest();
      //expect(instance).to.be();
    });

    it('should have the property perAlloc (base name: "PerAlloc")', function() {
      // uncomment below and update the code to test the property perAlloc
      //var instance = new nomad-client.VolumeRequest();
      //expect(instance).to.be();
    });

    it('should have the property readOnly (base name: "ReadOnly")', function() {
      // uncomment below and update the code to test the property readOnly
      //var instance = new nomad-client.VolumeRequest();
      //expect(instance).to.be();
    });

    it('should have the property source (base name: "Source")', function() {
      // uncomment below and update the code to test the property source
      //var instance = new nomad-client.VolumeRequest();
      //expect(instance).to.be();
    });

    it('should have the property type (base name: "Type")', function() {
      // uncomment below and update the code to test the property type
      //var instance = new nomad-client.VolumeRequest();
      //expect(instance).to.be();
    });

  });

}));
