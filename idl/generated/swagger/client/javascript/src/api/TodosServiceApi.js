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
    define(['ApiClient', 'model/TodosCreateTodoRequest', 'model/TodosCreateTodoResponse', 'model/TodosDeleteTodoResponse', 'model/TodosGetTodoResponse', 'model/TodosListTodosResponse', 'model/TodosUpdateTodoRequest', 'model/TodosUpdateTodoResponse'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('../model/TodosCreateTodoRequest'), require('../model/TodosCreateTodoResponse'), require('../model/TodosDeleteTodoResponse'), require('../model/TodosGetTodoResponse'), require('../model/TodosListTodosResponse'), require('../model/TodosUpdateTodoRequest'), require('../model/TodosUpdateTodoResponse'));
  } else {
    // Browser globals (root is window)
    if (!root.Api) {
      root.Api = {};
    }
    root.Api.TodosServiceApi = factory(root.Api.ApiClient, root.Api.TodosCreateTodoRequest, root.Api.TodosCreateTodoResponse, root.Api.TodosDeleteTodoResponse, root.Api.TodosGetTodoResponse, root.Api.TodosListTodosResponse, root.Api.TodosUpdateTodoRequest, root.Api.TodosUpdateTodoResponse);
  }
}(this, function(ApiClient, TodosCreateTodoRequest, TodosCreateTodoResponse, TodosDeleteTodoResponse, TodosGetTodoResponse, TodosListTodosResponse, TodosUpdateTodoRequest, TodosUpdateTodoResponse) {
  'use strict';

  /**
   * TodosService service.
   * @module api/TodosServiceApi
   * @version 0.0.1
   */

  /**
   * Constructs a new TodosServiceApi. 
   * @alias module:api/TodosServiceApi
   * @class
   * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
   * default to {@link module:ApiClient#instance} if unspecified.
   */
  var exports = function(apiClient) {
    this.apiClient = apiClient || ApiClient.instance;



    /**
     * @param {module:model/TodosCreateTodoRequest} body 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/TodosCreateTodoResponse} and HTTP response
     */
    this.createTodoWithHttpInfo = function(body) {
      var postBody = body;

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling createTodo");
      }


      var pathParams = {
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = [];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = TodosCreateTodoResponse;

      return this.apiClient.callApi(
        '/v1/todos', 'POST',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType
      );
    }

    /**
     * @param {module:model/TodosCreateTodoRequest} body 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/TodosCreateTodoResponse}
     */
    this.createTodo = function(body) {
      return this.createTodoWithHttpInfo(body)
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


    /**
     * @param {String} todoId 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/TodosDeleteTodoResponse} and HTTP response
     */
    this.deleteTodoWithHttpInfo = function(todoId) {
      var postBody = null;

      // verify the required parameter 'todoId' is set
      if (todoId === undefined || todoId === null) {
        throw new Error("Missing the required parameter 'todoId' when calling deleteTodo");
      }


      var pathParams = {
        'todo_id': todoId
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = [];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = TodosDeleteTodoResponse;

      return this.apiClient.callApi(
        '/v1/todos/{todo_id}', 'DELETE',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType
      );
    }

    /**
     * @param {String} todoId 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/TodosDeleteTodoResponse}
     */
    this.deleteTodo = function(todoId) {
      return this.deleteTodoWithHttpInfo(todoId)
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


    /**
     * @param {String} todoId 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/TodosGetTodoResponse} and HTTP response
     */
    this.getTodoWithHttpInfo = function(todoId) {
      var postBody = null;

      // verify the required parameter 'todoId' is set
      if (todoId === undefined || todoId === null) {
        throw new Error("Missing the required parameter 'todoId' when calling getTodo");
      }


      var pathParams = {
        'todo_id': todoId
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = [];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = TodosGetTodoResponse;

      return this.apiClient.callApi(
        '/v1/todos/{todo_id}', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType
      );
    }

    /**
     * @param {String} todoId 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/TodosGetTodoResponse}
     */
    this.getTodo = function(todoId) {
      return this.getTodoWithHttpInfo(todoId)
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


    /**
     * @param {Object} opts Optional parameters
     * @param {module:model/String} opts.parentType  (default to PARENT_TYPE_TODO_GROUP)
     * @param {String} opts.parentId 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/TodosListTodosResponse} and HTTP response
     */
    this.listTodosWithHttpInfo = function(opts) {
      opts = opts || {};
      var postBody = null;


      var pathParams = {
      };
      var queryParams = {
        'parent_type': opts['parentType'],
        'parent_id': opts['parentId'],
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = [];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = TodosListTodosResponse;

      return this.apiClient.callApi(
        '/v1/todos', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType
      );
    }

    /**
     * @param {Object} opts Optional parameters
     * @param {module:model/String} opts.parentType  (default to PARENT_TYPE_TODO_GROUP)
     * @param {String} opts.parentId 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/TodosListTodosResponse}
     */
    this.listTodos = function(opts) {
      return this.listTodosWithHttpInfo(opts)
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


    /**
     * @param {String} todoTodoId 
     * @param {module:model/TodosUpdateTodoRequest} body 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/TodosUpdateTodoResponse} and HTTP response
     */
    this.updateTodoWithHttpInfo = function(todoTodoId, body) {
      var postBody = body;

      // verify the required parameter 'todoTodoId' is set
      if (todoTodoId === undefined || todoTodoId === null) {
        throw new Error("Missing the required parameter 'todoTodoId' when calling updateTodo");
      }

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling updateTodo");
      }


      var pathParams = {
        'todo.todo_id': todoTodoId
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = [];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = TodosUpdateTodoResponse;

      return this.apiClient.callApi(
        '/v1/todos/{todo.todo_id}', 'PUT',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType
      );
    }

    /**
     * @param {String} todoTodoId 
     * @param {module:model/TodosUpdateTodoRequest} body 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/TodosUpdateTodoResponse}
     */
    this.updateTodo = function(todoTodoId, body) {
      return this.updateTodoWithHttpInfo(todoTodoId, body)
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }
  };

  return exports;
}));
