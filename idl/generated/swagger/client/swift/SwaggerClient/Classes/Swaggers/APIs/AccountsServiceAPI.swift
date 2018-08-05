//
// AccountsServiceAPI.swift
//
// Generated by swagger-codegen
// https://github.com/swagger-api/swagger-codegen
//

import Foundation
import Alamofire
import RxSwift



open class AccountsServiceAPI {
    /**

     - parameter body: (body)  
     - parameter completion: completion handler to receive the data and the error objects
     */
    open class func logIn(body: AccountsLogInRequest, completion: @escaping ((_ data: AccountsLogInResponse?,_ error: Error?) -> Void)) {
        logInWithRequestBuilder(body: body).execute { (response, error) -> Void in
            completion(response?.body, error);
        }
    }

    /**

     - parameter body: (body)  
     - returns: Observable<AccountsLogInResponse>
     */
    open class func logIn(body: AccountsLogInRequest) -> Observable<AccountsLogInResponse> {
        return Observable.create { observer -> Disposable in
            logIn(body: body) { data, error in
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
     - POST /v1/accounts/logIn
     - examples: [{contentType=application/json, example={
  "access_token" : "access_token",
  "refresh_token" : "refresh_token",
  "account_id" : "account_id",
  "token_type" : "token_type",
  "expires_in" : "expires_in"
}}]
     
     - parameter body: (body)  

     - returns: RequestBuilder<AccountsLogInResponse> 
     */
    open class func logInWithRequestBuilder(body: AccountsLogInRequest) -> RequestBuilder<AccountsLogInResponse> {
        let path = "/v1/accounts/logIn"
        let URLString = SwaggerClientAPI.basePath + path
        let parameters = JSONEncodingHelper.encodingParameters(forEncodableObject: body)

        let url = NSURLComponents(string: URLString)


        let requestBuilder: RequestBuilder<AccountsLogInResponse>.Type = SwaggerClientAPI.requestBuilderFactory.getBuilder()

        return requestBuilder.init(method: "POST", URLString: (url?.string ?? URLString), parameters: parameters, isBody: true)
    }

    /**

     - parameter body: (body)  
     - parameter completion: completion handler to receive the data and the error objects
     */
    open class func register(body: AccountsRegisterRequest, completion: @escaping ((_ data: AccountsRegisterResponse?,_ error: Error?) -> Void)) {
        registerWithRequestBuilder(body: body).execute { (response, error) -> Void in
            completion(response?.body, error);
        }
    }

    /**

     - parameter body: (body)  
     - returns: Observable<AccountsRegisterResponse>
     */
    open class func register(body: AccountsRegisterRequest) -> Observable<AccountsRegisterResponse> {
        return Observable.create { observer -> Disposable in
            register(body: body) { data, error in
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
     - POST /v1/accounts/register
     - examples: [{contentType=application/json, example={
  "auth_type" : { },
  "account_id" : "account_id"
}}]
     
     - parameter body: (body)  

     - returns: RequestBuilder<AccountsRegisterResponse> 
     */
    open class func registerWithRequestBuilder(body: AccountsRegisterRequest) -> RequestBuilder<AccountsRegisterResponse> {
        let path = "/v1/accounts/register"
        let URLString = SwaggerClientAPI.basePath + path
        let parameters = JSONEncodingHelper.encodingParameters(forEncodableObject: body)

        let url = NSURLComponents(string: URLString)


        let requestBuilder: RequestBuilder<AccountsRegisterResponse>.Type = SwaggerClientAPI.requestBuilderFactory.getBuilder()

        return requestBuilder.init(method: "POST", URLString: (url?.string ?? URLString), parameters: parameters, isBody: true)
    }

}
