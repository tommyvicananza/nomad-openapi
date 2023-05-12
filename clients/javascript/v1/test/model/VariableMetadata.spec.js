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
    instance = new nomad-client.VariableMetadata();
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

  describe('VariableMetadata', function() {
    it('should create an instance of VariableMetadata', function() {
      // uncomment below and update the code to test VariableMetadata
      //var instane = new nomad-client.VariableMetadata();
      //expect(instance).to.be.a(nomad-client.VariableMetadata);
    });

    it('should have the property createIndex (base name: "CreateIndex")', function() {
      // uncomment below and update the code to test the property createIndex
      //var instance = new nomad-client.VariableMetadata();
      //expect(instance).to.be();
    });

    it('should have the property createTime (base name: "CreateTime")', function() {
      // uncomment below and update the code to test the property createTime
      //var instance = new nomad-client.VariableMetadata();
      //expect(instance).to.be();
    });

    it('should have the property modifyIndex (base name: "ModifyIndex")', function() {
      // uncomment below and update the code to test the property modifyIndex
      //var instance = new nomad-client.VariableMetadata();
      //expect(instance).to.be();
    });

    it('should have the property modifyTime (base name: "ModifyTime")', function() {
      // uncomment below and update the code to test the property modifyTime
      //var instance = new nomad-client.VariableMetadata();
      //expect(instance).to.be();
    });

    it('should have the property namespace (base name: "Namespace")', function() {
      // uncomment below and update the code to test the property namespace
      //var instance = new nomad-client.VariableMetadata();
      //expect(instance).to.be();
    });

    it('should have the property path (base name: "Path")', function() {
      // uncomment below and update the code to test the property path
      //var instance = new nomad-client.VariableMetadata();
      //expect(instance).to.be();
    });

  });

}));
