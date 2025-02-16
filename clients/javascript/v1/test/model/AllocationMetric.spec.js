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
    instance = new nomad-client.AllocationMetric();
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

  describe('AllocationMetric', function() {
    it('should create an instance of AllocationMetric', function() {
      // uncomment below and update the code to test AllocationMetric
      //var instane = new nomad-client.AllocationMetric();
      //expect(instance).to.be.a(nomad-client.AllocationMetric);
    });

    it('should have the property allocationTime (base name: "AllocationTime")', function() {
      // uncomment below and update the code to test the property allocationTime
      //var instance = new nomad-client.AllocationMetric();
      //expect(instance).to.be();
    });

    it('should have the property classExhausted (base name: "ClassExhausted")', function() {
      // uncomment below and update the code to test the property classExhausted
      //var instance = new nomad-client.AllocationMetric();
      //expect(instance).to.be();
    });

    it('should have the property classFiltered (base name: "ClassFiltered")', function() {
      // uncomment below and update the code to test the property classFiltered
      //var instance = new nomad-client.AllocationMetric();
      //expect(instance).to.be();
    });

    it('should have the property coalescedFailures (base name: "CoalescedFailures")', function() {
      // uncomment below and update the code to test the property coalescedFailures
      //var instance = new nomad-client.AllocationMetric();
      //expect(instance).to.be();
    });

    it('should have the property constraintFiltered (base name: "ConstraintFiltered")', function() {
      // uncomment below and update the code to test the property constraintFiltered
      //var instance = new nomad-client.AllocationMetric();
      //expect(instance).to.be();
    });

    it('should have the property dimensionExhausted (base name: "DimensionExhausted")', function() {
      // uncomment below and update the code to test the property dimensionExhausted
      //var instance = new nomad-client.AllocationMetric();
      //expect(instance).to.be();
    });

    it('should have the property nodesAvailable (base name: "NodesAvailable")', function() {
      // uncomment below and update the code to test the property nodesAvailable
      //var instance = new nomad-client.AllocationMetric();
      //expect(instance).to.be();
    });

    it('should have the property nodesEvaluated (base name: "NodesEvaluated")', function() {
      // uncomment below and update the code to test the property nodesEvaluated
      //var instance = new nomad-client.AllocationMetric();
      //expect(instance).to.be();
    });

    it('should have the property nodesExhausted (base name: "NodesExhausted")', function() {
      // uncomment below and update the code to test the property nodesExhausted
      //var instance = new nomad-client.AllocationMetric();
      //expect(instance).to.be();
    });

    it('should have the property nodesFiltered (base name: "NodesFiltered")', function() {
      // uncomment below and update the code to test the property nodesFiltered
      //var instance = new nomad-client.AllocationMetric();
      //expect(instance).to.be();
    });

    it('should have the property quotaExhausted (base name: "QuotaExhausted")', function() {
      // uncomment below and update the code to test the property quotaExhausted
      //var instance = new nomad-client.AllocationMetric();
      //expect(instance).to.be();
    });

    it('should have the property resourcesExhausted (base name: "ResourcesExhausted")', function() {
      // uncomment below and update the code to test the property resourcesExhausted
      //var instance = new nomad-client.AllocationMetric();
      //expect(instance).to.be();
    });

    it('should have the property scoreMetaData (base name: "ScoreMetaData")', function() {
      // uncomment below and update the code to test the property scoreMetaData
      //var instance = new nomad-client.AllocationMetric();
      //expect(instance).to.be();
    });

    it('should have the property scores (base name: "Scores")', function() {
      // uncomment below and update the code to test the property scores
      //var instance = new nomad-client.AllocationMetric();
      //expect(instance).to.be();
    });

  });

}));
