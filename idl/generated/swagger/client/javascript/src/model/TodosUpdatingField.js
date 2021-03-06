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
    define(['ApiClient'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'));
  } else {
    // Browser globals (root is window)
    if (!root.Api) {
      root.Api = {};
    }
    root.Api.TodosUpdatingField = factory(root.Api.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';


  /**
   * Enum class TodosUpdatingField.
   * @enum {}
   * @readonly
   */
  var exports = {
    /**
     * value: "UPDATING_FIELD_PARENT"
     * @const
     */
    "PARENT": "UPDATING_FIELD_PARENT",
    /**
     * value: "UPDATING_FIELD_TITLE"
     * @const
     */
    "TITLE": "UPDATING_FIELD_TITLE",
    /**
     * value: "UPDATING_FIELD_DESCRIPTION"
     * @const
     */
    "DESCRIPTION": "UPDATING_FIELD_DESCRIPTION",
    /**
     * value: "UPDATING_FIELD_STATUS"
     * @const
     */
    "STATUS": "UPDATING_FIELD_STATUS",
    /**
     * value: "UPDATING_FIELD_ORDER"
     * @const
     */
    "ORDER": "UPDATING_FIELD_ORDER",
    /**
     * value: "UPDATING_FIELD_ASSIGNED_TO"
     * @const
     */
    "ASSIGNED_TO": "UPDATING_FIELD_ASSIGNED_TO",
    /**
     * value: "UPDATING_FIELD_PRIORITY"
     * @const
     */
    "PRIORITY": "UPDATING_FIELD_PRIORITY",
    /**
     * value: "UPDATING_FIELD_DUE_AT"
     * @const
     */
    "DUE_AT": "UPDATING_FIELD_DUE_AT"  };

  /**
   * Returns a <code>TodosUpdatingField</code> enum value from a Javascript object name.
   * @param {Object} data The plain JavaScript object containing the name of the enum value.
   * @return {module:model/TodosUpdatingField} The enum <code>TodosUpdatingField</code> value.
   */
  exports.constructFromObject = function(object) {
    return object;
  }

  return exports;
}));


