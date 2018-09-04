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
    // AMD. Register as an anonymous module.
    define(['ApiClient', 'model/TodosTodo'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./TodosTodo'));
  } else {
    // Browser globals (root is window)
    if (!root.Api) {
      root.Api = {};
    }
    root.Api.TodosCreateTodoRequest = factory(root.Api.ApiClient, root.Api.TodosTodo);
  }
}(this, function(ApiClient, TodosTodo) {
  'use strict';




  /**
   * The TodosCreateTodoRequest model module.
   * @module model/TodosCreateTodoRequest
   * @version 0.0.1
   */

  /**
   * Constructs a new <code>TodosCreateTodoRequest</code>.
   * @alias module:model/TodosCreateTodoRequest
   * @class
   */
  var exports = function() {
    var _this = this;


  };

  /**
   * Constructs a <code>TodosCreateTodoRequest</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/TodosCreateTodoRequest} obj Optional instance to populate.
   * @return {module:model/TodosCreateTodoRequest} The populated <code>TodosCreateTodoRequest</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('todo')) {
        obj['todo'] = TodosTodo.constructFromObject(data['todo']);
      }
    }
    return obj;
  }

  /**
   * @member {module:model/TodosTodo} todo
   */
  exports.prototype['todo'] = undefined;



  return exports;
}));


