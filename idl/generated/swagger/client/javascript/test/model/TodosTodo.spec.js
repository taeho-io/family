/**
 * API
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 0.0.1
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.3.1
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', '../../src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require('../../src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.Api);
  }
}(this, function(expect, Api) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new Api.TodosTodo();
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

  describe('TodosTodo', function() {
    it('should create an instance of TodosTodo', function() {
      // uncomment below and update the code to test TodosTodo
      //var instane = new Api.TodosTodo();
      //expect(instance).to.be.a(Api.TodosTodo);
    });

    it('should have the property assignedTo (base name: "assigned_to")', function() {
      // uncomment below and update the code to test the property assignedTo
      //var instane = new Api.TodosTodo();
      //expect(instance).to.be();
    });

    it('should have the property createdAt (base name: "created_at")', function() {
      // uncomment below and update the code to test the property createdAt
      //var instane = new Api.TodosTodo();
      //expect(instance).to.be();
    });

    it('should have the property description (base name: "description")', function() {
      // uncomment below and update the code to test the property description
      //var instane = new Api.TodosTodo();
      //expect(instance).to.be();
    });

    it('should have the property doneAt (base name: "done_at")', function() {
      // uncomment below and update the code to test the property doneAt
      //var instane = new Api.TodosTodo();
      //expect(instance).to.be();
    });

    it('should have the property dueAt (base name: "due_at")', function() {
      // uncomment below and update the code to test the property dueAt
      //var instane = new Api.TodosTodo();
      //expect(instance).to.be();
    });

    it('should have the property order (base name: "order")', function() {
      // uncomment below and update the code to test the property order
      //var instane = new Api.TodosTodo();
      //expect(instance).to.be();
    });

    it('should have the property parentId (base name: "parent_id")', function() {
      // uncomment below and update the code to test the property parentId
      //var instane = new Api.TodosTodo();
      //expect(instance).to.be();
    });

    it('should have the property parentType (base name: "parent_type")', function() {
      // uncomment below and update the code to test the property parentType
      //var instane = new Api.TodosTodo();
      //expect(instance).to.be();
    });

    it('should have the property priority (base name: "priority")', function() {
      // uncomment below and update the code to test the property priority
      //var instane = new Api.TodosTodo();
      //expect(instance).to.be();
    });

    it('should have the property status (base name: "status")', function() {
      // uncomment below and update the code to test the property status
      //var instane = new Api.TodosTodo();
      //expect(instance).to.be();
    });

    it('should have the property subTasks (base name: "sub_tasks")', function() {
      // uncomment below and update the code to test the property subTasks
      //var instane = new Api.TodosTodo();
      //expect(instance).to.be();
    });

    it('should have the property title (base name: "title")', function() {
      // uncomment below and update the code to test the property title
      //var instane = new Api.TodosTodo();
      //expect(instance).to.be();
    });

    it('should have the property todoId (base name: "todo_id")', function() {
      // uncomment below and update the code to test the property todoId
      //var instane = new Api.TodosTodo();
      //expect(instance).to.be();
    });

    it('should have the property updatedAt (base name: "updated_at")', function() {
      // uncomment below and update the code to test the property updatedAt
      //var instane = new Api.TodosTodo();
      //expect(instance).to.be();
    });

  });

}));