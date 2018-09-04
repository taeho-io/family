//
// NotesUpdateNoteRequest.swift
//
// Generated by swagger-codegen
// https://github.com/swagger-api/swagger-codegen
//

import Foundation



open class NotesUpdateNoteRequest: Codable {

    public var note: NotesNote?


    
    public init(note: NotesNote?) {
        self.note = note
    }
    

    // Encodable protocol methods

    public func encode(to encoder: Encoder) throws {

        var container = encoder.container(keyedBy: String.self)

        try container.encodeIfPresent(note, forKey: "note")
    }

    // Decodable protocol methods

    public required init(from decoder: Decoder) throws {
        let container = try decoder.container(keyedBy: String.self)

        note = try container.decodeIfPresent(NotesNote.self, forKey: "note")
    }
}

