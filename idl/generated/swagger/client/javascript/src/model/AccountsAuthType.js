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
    root.Api.AccountsAuthType = factory(root.Api.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';


  /**
   * Enum class AccountsAuthType.
   * @enum {}
   * @readonly
   */
  var exports = {
    /**
     * value: "NONE"
     * @const
     */
    "NONE": "NONE",
    /**
     * value: "EMAIL"
     * @const
     */
    "EMAIL": "EMAIL"  };

  /**
   * Returns a <code>AccountsAuthType</code> enum value from a Javascript object name.
   * @param {Object} data The plain JavaScript object containing the name of the enum value.
   * @return {module:model/AccountsAuthType} The enum <code>AccountsAuthType</code> value.
   */
  exports.constructFromObject = function(object) {
    return object;
  }

  return exports;
}));


