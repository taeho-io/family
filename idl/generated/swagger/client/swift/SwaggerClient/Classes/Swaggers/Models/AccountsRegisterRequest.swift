//
// AccountsRegisterRequest.swift
//
// Generated by swagger-codegen
// https://github.com/swagger-api/swagger-codegen
//

import Foundation



open class AccountsRegisterRequest: Codable {

    public var authType: AccountsAuthType?
    public var email: String?
    public var fullName: String?
    public var password: String?


    
    public init(authType: AccountsAuthType?, email: String?, fullName: String?, password: String?) {
        self.authType = authType
        self.email = email
        self.fullName = fullName
        self.password = password
    }
    

    // Encodable protocol methods

    public func encode(to encoder: Encoder) throws {

        var container = encoder.container(keyedBy: String.self)

        try container.encodeIfPresent(authType, forKey: "auth_type")
        try container.encodeIfPresent(email, forKey: "email")
        try container.encodeIfPresent(fullName, forKey: "full_name")
        try container.encodeIfPresent(password, forKey: "password")
    }

    // Decodable protocol methods

    public required init(from decoder: Decoder) throws {
        let container = try decoder.container(keyedBy: String.self)

        authType = try container.decodeIfPresent(AccountsAuthType.self, forKey: "auth_type")
        email = try container.decodeIfPresent(String.self, forKey: "email")
        fullName = try container.decodeIfPresent(String.self, forKey: "full_name")
        password = try container.decodeIfPresent(String.self, forKey: "password")
    }
}

