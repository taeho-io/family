//
// TodosServiceAPI.swift
//
// Generated by swagger-codegen
// https://github.com/swagger-api/swagger-codegen
//

import Foundation
import Alamofire
import RxSwift



open class TodosServiceAPI {
    /**

     - parameter body: (body)  
     - parameter completion: completion handler to receive the data and the error objects
     */
    open class func createTodo(body: TodosCreateTodoRequest, completion: @escaping ((_ data: TodosCreateTodoResponse?,_ error: Error?) -> Void)) {
        createTodoWithRequestBuilder(body: body).execute { (response, error) -> Void in
            completion(response?.body, error);
        }
    }

    /**

     - parameter body: (body)  
     - returns: Observable<TodosCreateTodoResponse>
     */
    open class func createTodo(body: TodosCreateTodoRequest) -> Observable<TodosCreateTodoResponse> {
        return Observable.create { observer -> Disposable in
            createTodo(body: body) { data, error in
                if let error = error {
                    observer.on(.error(error))
                } else {
                    observer.on(.next(data!))
                }
                observer.on(.completed)
            }
            return Disposables.create()
        }
    }

    /**
     - POST /v1/todos
     - examples: [{contentType=application/json, example={
  "todo" : {
    "todo_id" : "todo_id",
    "done_at" : "done_at",
    "created_at" : "created_at",
    "description" : "description",
    "parent_type" : { },
    "priority" : { },
    "title" : "title",
    "updated_at" : "updated_at",
    "parent_id" : "parent_id",
    "sub_tasks" : [ null, null ],
    "due_at" : "due_at",
    "assigned_to" : "assigned_to",
    "order" : "order",
    "status" : { }
  }
}}]
     
     - parameter body: (body)  

     - returns: RequestBuilder<TodosCreateTodoResponse> 
     */
    open class func createTodoWithRequestBuilder(body: TodosCreateTodoRequest) -> RequestBuilder<TodosCreateTodoResponse> {
        let path = "/v1/todos"
        let URLString = SwaggerClientAPI.basePath + path
        let parameters = JSONEncodingHelper.encodingParameters(forEncodableObject: body)

        let url = NSURLComponents(string: URLString)


        let requestBuilder: RequestBuilder<TodosCreateTodoResponse>.Type = SwaggerClientAPI.requestBuilderFactory.getBuilder()

        return requestBuilder.init(method: "POST", URLString: (url?.string ?? URLString), parameters: parameters, isBody: true)
    }

    /**

     - parameter todoId: (path)  
     - parameter completion: completion handler to receive the data and the error objects
     */
    open class func deleteTodo(todoId: String, completion: @escaping ((_ data: TodosDeleteTodoResponse?,_ error: Error?) -> Void)) {
        deleteTodoWithRequestBuilder(todoId: todoId).execute { (response, error) -> Void in
            completion(response?.body, error);
        }
    }

    /**

     - parameter todoId: (path)  
     - returns: Observable<TodosDeleteTodoResponse>
     */
    open class func deleteTodo(todoId: String) -> Observable<TodosDeleteTodoResponse> {
        return Observable.create { observer -> Disposable in
            deleteTodo(todoId: todoId) { data, error in
                if let error = error {
                    observer.on(.error(error))
                } else {
                    observer.on(.next(data!))
                }
                observer.on(.completed)
            }
            return Disposables.create()
        }
    }

    /**
     - DELETE /v1/todos/{todo_id}
     - examples: [{contentType=application/json, example={ }}]
     
     - parameter todoId: (path)  

     - returns: RequestBuilder<TodosDeleteTodoResponse> 
     */
    open class func deleteTodoWithRequestBuilder(todoId: String) -> RequestBuilder<TodosDeleteTodoResponse> {
        var path = "/v1/todos/{todo_id}"
        path = path.replacingOccurrences(of: "{todo_id}", with: "\(todoId)", options: .literal, range: nil)
        let URLString = SwaggerClientAPI.basePath + path
        let parameters: [String:Any]? = nil

        let url = NSURLComponents(string: URLString)


        let requestBuilder: RequestBuilder<TodosDeleteTodoResponse>.Type = SwaggerClientAPI.requestBuilderFactory.getBuilder()

        return requestBuilder.init(method: "DELETE", URLString: (url?.string ?? URLString), parameters: parameters, isBody: false)
    }

    /**

     - parameter todoId: (path)  
     - parameter completion: completion handler to receive the data and the error objects
     */
    open class func getTodo(todoId: String, completion: @escaping ((_ data: TodosGetTodoResponse?,_ error: Error?) -> Void)) {
        getTodoWithRequestBuilder(todoId: todoId).execute { (response, error) -> Void in
            completion(response?.body, error);
        }
    }

    /**

     - parameter todoId: (path)  
     - returns: Observable<TodosGetTodoResponse>
     */
    open class func getTodo(todoId: String) -> Observable<TodosGetTodoResponse> {
        return Observable.create { observer -> Disposable in
            getTodo(todoId: todoId) { data, error in
                if let error = error {
                    observer.on(.error(error))
                } else {
                    observer.on(.next(data!))
                }
                observer.on(.completed)
            }
            return Disposables.create()
        }
    }

    /**
     - GET /v1/todos/{todo_id}
     - examples: [{contentType=application/json, example={
  "todo" : {
    "todo_id" : "todo_id",
    "done_at" : "done_at",
    "created_at" : "created_at",
    "description" : "description",
    "parent_type" : { },
    "priority" : { },
    "title" : "title",
    "updated_at" : "updated_at",
    "parent_id" : "parent_id",
    "sub_tasks" : [ null, null ],
    "due_at" : "due_at",
    "assigned_to" : "assigned_to",
    "order" : "order",
    "status" : { }
  }
}}]
     
     - parameter todoId: (path)  

     - returns: RequestBuilder<TodosGetTodoResponse> 
     */
    open class func getTodoWithRequestBuilder(todoId: String) -> RequestBuilder<TodosGetTodoResponse> {
        var path = "/v1/todos/{todo_id}"
        path = path.replacingOccurrences(of: "{todo_id}", with: "\(todoId)", options: .literal, range: nil)
        let URLString = SwaggerClientAPI.basePath + path
        let parameters: [String:Any]? = nil

        let url = NSURLComponents(string: URLString)


        let requestBuilder: RequestBuilder<TodosGetTodoResponse>.Type = SwaggerClientAPI.requestBuilderFactory.getBuilder()

        return requestBuilder.init(method: "GET", URLString: (url?.string ?? URLString), parameters: parameters, isBody: false)
    }

    /**
     * enum for parameter parentType
     */
    public enum ParentType_listTodos: String { 
        case todoGroup = "PARENT_TYPE_TODO_GROUP"
        case todo = "PARENT_TYPE_TODO"
    }

    /**

     - parameter parentType: (query)  (optional, default to PARENT_TYPE_TODO_GROUP)
     - parameter parentId: (query)  (optional)
     - parameter completion: completion handler to receive the data and the error objects
     */
    open class func listTodos(parentType: ParentType_listTodos? = nil, parentId: String? = nil, completion: @escaping ((_ data: TodosListTodosResponse?,_ error: Error?) -> Void)) {
        listTodosWithRequestBuilder(parentType: parentType, parentId: parentId).execute { (response, error) -> Void in
            completion(response?.body, error);
        }
    }

    /**

     - parameter parentType: (query)  (optional, default to PARENT_TYPE_TODO_GROUP)
     - parameter parentId: (query)  (optional)
     - returns: Observable<TodosListTodosResponse>
     */
    open class func listTodos(parentType: ParentType_listTodos? = nil, parentId: String? = nil) -> Observable<TodosListTodosResponse> {
        return Observable.create { observer -> Disposable in
            listTodos(parentType: parentType, parentId: parentId) { data, error in
                if let error = error {
                    observer.on(.error(error))
                } else {
                    observer.on(.next(data!))
                }
                observer.on(.completed)
            }
            return Disposables.create()
        }
    }

    /**
     - GET /v1/todos
     - examples: [{contentType=application/json, example={
  "todos" : [ {
    "todo_id" : "todo_id",
    "done_at" : "done_at",
    "created_at" : "created_at",
    "description" : "description",
    "parent_type" : { },
    "priority" : { },
    "title" : "title",
    "updated_at" : "updated_at",
    "parent_id" : "parent_id",
    "sub_tasks" : [ null, null ],
    "due_at" : "due_at",
    "assigned_to" : "assigned_to",
    "order" : "order",
    "status" : { }
  }, {
    "todo_id" : "todo_id",
    "done_at" : "done_at",
    "created_at" : "created_at",
    "description" : "description",
    "parent_type" : { },
    "priority" : { },
    "title" : "title",
    "updated_at" : "updated_at",
    "parent_id" : "parent_id",
    "sub_tasks" : [ null, null ],
    "due_at" : "due_at",
    "assigned_to" : "assigned_to",
    "order" : "order",
    "status" : { }
  } ]
}}]
     
     - parameter parentType: (query)  (optional, default to PARENT_TYPE_TODO_GROUP)
     - parameter parentId: (query)  (optional)

     - returns: RequestBuilder<TodosListTodosResponse> 
     */
    open class func listTodosWithRequestBuilder(parentType: ParentType_listTodos? = nil, parentId: String? = nil) -> RequestBuilder<TodosListTodosResponse> {
        let path = "/v1/todos"
        let URLString = SwaggerClientAPI.basePath + path
        let parameters: [String:Any]? = nil

        let url = NSURLComponents(string: URLString)
        url?.queryItems = APIHelper.mapValuesToQueryItems(values:[
            "parent_type": parentType?.rawValue, 
            "parent_id": parentId
        ])
        

        let requestBuilder: RequestBuilder<TodosListTodosResponse>.Type = SwaggerClientAPI.requestBuilderFactory.getBuilder()

        return requestBuilder.init(method: "GET", URLString: (url?.string ?? URLString), parameters: parameters, isBody: false)
    }

    /**

     - parameter todoTodoId: (path)  
     - parameter body: (body)  
     - parameter completion: completion handler to receive the data and the error objects
     */
    open class func updateTodo(todoTodoId: String, body: TodosUpdateTodoRequest, completion: @escaping ((_ data: TodosUpdateTodoResponse?,_ error: Error?) -> Void)) {
        updateTodoWithRequestBuilder(todoTodoId: todoTodoId, body: body).execute { (response, error) -> Void in
            completion(response?.body, error);
        }
    }

    /**

     - parameter todoTodoId: (path)  
     - parameter body: (body)  
     - returns: Observable<TodosUpdateTodoResponse>
     */
    open class func updateTodo(todoTodoId: String, body: TodosUpdateTodoRequest) -> Observable<TodosUpdateTodoResponse> {
        return Observable.create { observer -> Disposable in
            updateTodo(todoTodoId: todoTodoId, body: body) { data, error in
                if let error = error {
                    observer.on(.error(error))
                } else {
                    observer.on(.next(data!))
                }
                observer.on(.completed)
            }
            return Disposables.create()
        }
    }

    /**
     - PUT /v1/todos/{todo.todo_id}
     - examples: [{contentType=application/json, example={
  "todo" : {
    "todo_id" : "todo_id",
    "done_at" : "done_at",
    "created_at" : "created_at",
    "description" : "description",
    "parent_type" : { },
    "priority" : { },
    "title" : "title",
    "updated_at" : "updated_at",
    "parent_id" : "parent_id",
    "sub_tasks" : [ null, null ],
    "due_at" : "due_at",
    "assigned_to" : "assigned_to",
    "order" : "order",
    "status" : { }
  }
}}]
     
     - parameter todoTodoId: (path)  
     - parameter body: (body)  

     - returns: RequestBuilder<TodosUpdateTodoResponse> 
     */
    open class func updateTodoWithRequestBuilder(todoTodoId: String, body: TodosUpdateTodoRequest) -> RequestBuilder<TodosUpdateTodoResponse> {
        var path = "/v1/todos/{todo.todo_id}"
        path = path.replacingOccurrences(of: "{todo.todo_id}", with: "\(todoTodoId)", options: .literal, range: nil)
        let URLString = SwaggerClientAPI.basePath + path
        let parameters = JSONEncodingHelper.encodingParameters(forEncodableObject: body)

        let url = NSURLComponents(string: URLString)


        let requestBuilder: RequestBuilder<TodosUpdateTodoResponse>.Type = SwaggerClientAPI.requestBuilderFactory.getBuilder()

        return requestBuilder.init(method: "PUT", URLString: (url?.string ?? URLString), parameters: parameters, isBody: true)
    }

}
