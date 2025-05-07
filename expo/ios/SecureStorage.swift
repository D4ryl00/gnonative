//
//  SecureStorage.swift
//  Pods
//
//  Created by Rémi BARBERO on 02/05/2025.
//

import GnoCore

public class SecureStorage: NSObject, GnoGnonativeSecureNativeStorageProtocol {
    public func get(_ p0: Data?) -> Data? {
        return nil
    }
    
    public func has(_ key: Data?) -> Bool {
        NSLog("secureStorage: has: key:" + (String(data: key ?? Data(), encoding: .utf8) ?? ""))
        return false
    }
    
    public func set(_ p0: Data?, p1: Data?) {
        NSLog("secureStorage: set: key:" + (String(data: p0 ?? Data(), encoding: .utf8) ?? "") + " value:" + (String(data: p1 ?? Data(), encoding: .utf8) ?? ""))
    }
    
    public func setSync(_ p0: Data?, p1: Data?) {
        NSLog("secureStorage: setSync: key:" + (String(data: p0 ?? Data(), encoding: .utf8) ?? "") + " value:" + (String(data: p1 ?? Data(), encoding: .utf8) ?? ""))
    }
    
    public func delete(_ p0: Data?) {
        NSLog("secureStorage: delete: key:" + (String(data: p0 ?? Data(), encoding: .utf8) ?? ""))
    }
    
    public func deleteSync(_ p0: Data?) {
        NSLog("secureStorage: deleteSync: key:" + (String(data: p0 ?? Data(), encoding: .utf8) ?? ""))
    }
    
    public func write() {
        NSLog("secureStorage: write")
    }
    
    public func writeSync() {
        NSLog("secureStorage: writeSync")
    }
    
    public func close() {
        NSLog("secureStorage: close")
    }
}
