//
// AuthRefreshRequest.swift
//
// Generated by swagger-codegen
// https://github.com/swagger-api/swagger-codegen
//

import Foundation



open class AuthRefreshRequest: Codable {

    public var refreshToken: String?


    
    public init(refreshToken: String?) {
        self.refreshToken = refreshToken
    }
    

    // Encodable protocol methods

    public func encode(to encoder: Encoder) throws {

        var container = encoder.container(keyedBy: String.self)

        try container.encodeIfPresent(refreshToken, forKey: "refresh_token")
    }

    // Decodable protocol methods

    public required init(from decoder: Decoder) throws {
        let container = try decoder.container(keyedBy: String.self)

        refreshToken = try container.decodeIfPresent(String.self, forKey: "refresh_token")
    }
}

