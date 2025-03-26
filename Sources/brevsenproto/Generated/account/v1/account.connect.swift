// Code generated by protoc-gen-connect-swift. DO NOT EDIT.
//
// Source: account/v1/account.proto
//

import Connect
import Foundation
import SwiftProtobuf

public protocol Account_V1_AccountServiceClientInterface: Sendable {

    @discardableResult
    func `getAccount`(request: Account_V1_GetAccountRequest, headers: Connect.Headers, completion: @escaping @Sendable (ResponseMessage<Account_V1_GetAccountResponse>) -> Void) -> Connect.Cancelable

    @available(iOS 13, *)
    func `getAccount`(request: Account_V1_GetAccountRequest, headers: Connect.Headers) async -> ResponseMessage<Account_V1_GetAccountResponse>

    @discardableResult
    func `getAccountWithID`(request: Account_V1_GetAccountWithIdRequest, headers: Connect.Headers, completion: @escaping @Sendable (ResponseMessage<Account_V1_GetAccountWithIdResponse>) -> Void) -> Connect.Cancelable

    @available(iOS 13, *)
    func `getAccountWithID`(request: Account_V1_GetAccountWithIdRequest, headers: Connect.Headers) async -> ResponseMessage<Account_V1_GetAccountWithIdResponse>

    @discardableResult
    func `generateAccountLinkPath`(request: Account_V1_GenerateAccountLinkPathRequest, headers: Connect.Headers, completion: @escaping @Sendable (ResponseMessage<Account_V1_GenerateAccountLinkPathResponse>) -> Void) -> Connect.Cancelable

    @available(iOS 13, *)
    func `generateAccountLinkPath`(request: Account_V1_GenerateAccountLinkPathRequest, headers: Connect.Headers) async -> ResponseMessage<Account_V1_GenerateAccountLinkPathResponse>

    @discardableResult
    func `resolveAccountLinkPath`(request: Account_V1_ResolveAccountLinkPathRequest, headers: Connect.Headers, completion: @escaping @Sendable (ResponseMessage<Account_V1_ResolveAccountLinkPathResponse>) -> Void) -> Connect.Cancelable

    @available(iOS 13, *)
    func `resolveAccountLinkPath`(request: Account_V1_ResolveAccountLinkPathRequest, headers: Connect.Headers) async -> ResponseMessage<Account_V1_ResolveAccountLinkPathResponse>

    @discardableResult
    func `deleteAccount`(request: Account_V1_DeleteAccountRequest, headers: Connect.Headers, completion: @escaping @Sendable (ResponseMessage<Account_V1_DeleteAccountResponse>) -> Void) -> Connect.Cancelable

    @available(iOS 13, *)
    func `deleteAccount`(request: Account_V1_DeleteAccountRequest, headers: Connect.Headers) async -> ResponseMessage<Account_V1_DeleteAccountResponse>

    @discardableResult
    func `addFriendWithID`(request: Account_V1_AddFriendWithIdRequest, headers: Connect.Headers, completion: @escaping @Sendable (ResponseMessage<Account_V1_AddFriendWithIdResponse>) -> Void) -> Connect.Cancelable

    @available(iOS 13, *)
    func `addFriendWithID`(request: Account_V1_AddFriendWithIdRequest, headers: Connect.Headers) async -> ResponseMessage<Account_V1_AddFriendWithIdResponse>

    @discardableResult
    func `blockAccountWithID`(request: Account_V1_BlockAccountWithIdRequest, headers: Connect.Headers, completion: @escaping @Sendable (ResponseMessage<Account_V1_BlockAccountWithIdResponse>) -> Void) -> Connect.Cancelable

    @available(iOS 13, *)
    func `blockAccountWithID`(request: Account_V1_BlockAccountWithIdRequest, headers: Connect.Headers) async -> ResponseMessage<Account_V1_BlockAccountWithIdResponse>

    @discardableResult
    func `sendMessage`(request: Account_V1_SendMessageRequest, headers: Connect.Headers, completion: @escaping @Sendable (ResponseMessage<Account_V1_SendMessageResponse>) -> Void) -> Connect.Cancelable

    @available(iOS 13, *)
    func `sendMessage`(request: Account_V1_SendMessageRequest, headers: Connect.Headers) async -> ResponseMessage<Account_V1_SendMessageResponse>
}

/// Concrete implementation of `Account_V1_AccountServiceClientInterface`.
public final class Account_V1_AccountServiceClient: Account_V1_AccountServiceClientInterface, Sendable {
    private let client: Connect.ProtocolClientInterface

    public init(client: Connect.ProtocolClientInterface) {
        self.client = client
    }

    @discardableResult
    public func `getAccount`(request: Account_V1_GetAccountRequest, headers: Connect.Headers = [:], completion: @escaping @Sendable (ResponseMessage<Account_V1_GetAccountResponse>) -> Void) -> Connect.Cancelable {
        return self.client.unary(path: "/account.v1.AccountService/GetAccount", idempotencyLevel: .unknown, request: request, headers: headers, completion: completion)
    }

    @available(iOS 13, *)
    public func `getAccount`(request: Account_V1_GetAccountRequest, headers: Connect.Headers = [:]) async -> ResponseMessage<Account_V1_GetAccountResponse> {
        return await self.client.unary(path: "/account.v1.AccountService/GetAccount", idempotencyLevel: .unknown, request: request, headers: headers)
    }

    @discardableResult
    public func `getAccountWithID`(request: Account_V1_GetAccountWithIdRequest, headers: Connect.Headers = [:], completion: @escaping @Sendable (ResponseMessage<Account_V1_GetAccountWithIdResponse>) -> Void) -> Connect.Cancelable {
        return self.client.unary(path: "/account.v1.AccountService/GetAccountWithId", idempotencyLevel: .unknown, request: request, headers: headers, completion: completion)
    }

    @available(iOS 13, *)
    public func `getAccountWithID`(request: Account_V1_GetAccountWithIdRequest, headers: Connect.Headers = [:]) async -> ResponseMessage<Account_V1_GetAccountWithIdResponse> {
        return await self.client.unary(path: "/account.v1.AccountService/GetAccountWithId", idempotencyLevel: .unknown, request: request, headers: headers)
    }

    @discardableResult
    public func `generateAccountLinkPath`(request: Account_V1_GenerateAccountLinkPathRequest, headers: Connect.Headers = [:], completion: @escaping @Sendable (ResponseMessage<Account_V1_GenerateAccountLinkPathResponse>) -> Void) -> Connect.Cancelable {
        return self.client.unary(path: "/account.v1.AccountService/GenerateAccountLinkPath", idempotencyLevel: .unknown, request: request, headers: headers, completion: completion)
    }

    @available(iOS 13, *)
    public func `generateAccountLinkPath`(request: Account_V1_GenerateAccountLinkPathRequest, headers: Connect.Headers = [:]) async -> ResponseMessage<Account_V1_GenerateAccountLinkPathResponse> {
        return await self.client.unary(path: "/account.v1.AccountService/GenerateAccountLinkPath", idempotencyLevel: .unknown, request: request, headers: headers)
    }

    @discardableResult
    public func `resolveAccountLinkPath`(request: Account_V1_ResolveAccountLinkPathRequest, headers: Connect.Headers = [:], completion: @escaping @Sendable (ResponseMessage<Account_V1_ResolveAccountLinkPathResponse>) -> Void) -> Connect.Cancelable {
        return self.client.unary(path: "/account.v1.AccountService/ResolveAccountLinkPath", idempotencyLevel: .unknown, request: request, headers: headers, completion: completion)
    }

    @available(iOS 13, *)
    public func `resolveAccountLinkPath`(request: Account_V1_ResolveAccountLinkPathRequest, headers: Connect.Headers = [:]) async -> ResponseMessage<Account_V1_ResolveAccountLinkPathResponse> {
        return await self.client.unary(path: "/account.v1.AccountService/ResolveAccountLinkPath", idempotencyLevel: .unknown, request: request, headers: headers)
    }

    @discardableResult
    public func `deleteAccount`(request: Account_V1_DeleteAccountRequest, headers: Connect.Headers = [:], completion: @escaping @Sendable (ResponseMessage<Account_V1_DeleteAccountResponse>) -> Void) -> Connect.Cancelable {
        return self.client.unary(path: "/account.v1.AccountService/DeleteAccount", idempotencyLevel: .unknown, request: request, headers: headers, completion: completion)
    }

    @available(iOS 13, *)
    public func `deleteAccount`(request: Account_V1_DeleteAccountRequest, headers: Connect.Headers = [:]) async -> ResponseMessage<Account_V1_DeleteAccountResponse> {
        return await self.client.unary(path: "/account.v1.AccountService/DeleteAccount", idempotencyLevel: .unknown, request: request, headers: headers)
    }

    @discardableResult
    public func `addFriendWithID`(request: Account_V1_AddFriendWithIdRequest, headers: Connect.Headers = [:], completion: @escaping @Sendable (ResponseMessage<Account_V1_AddFriendWithIdResponse>) -> Void) -> Connect.Cancelable {
        return self.client.unary(path: "/account.v1.AccountService/AddFriendWithId", idempotencyLevel: .unknown, request: request, headers: headers, completion: completion)
    }

    @available(iOS 13, *)
    public func `addFriendWithID`(request: Account_V1_AddFriendWithIdRequest, headers: Connect.Headers = [:]) async -> ResponseMessage<Account_V1_AddFriendWithIdResponse> {
        return await self.client.unary(path: "/account.v1.AccountService/AddFriendWithId", idempotencyLevel: .unknown, request: request, headers: headers)
    }

    @discardableResult
    public func `blockAccountWithID`(request: Account_V1_BlockAccountWithIdRequest, headers: Connect.Headers = [:], completion: @escaping @Sendable (ResponseMessage<Account_V1_BlockAccountWithIdResponse>) -> Void) -> Connect.Cancelable {
        return self.client.unary(path: "/account.v1.AccountService/BlockAccountWithId", idempotencyLevel: .unknown, request: request, headers: headers, completion: completion)
    }

    @available(iOS 13, *)
    public func `blockAccountWithID`(request: Account_V1_BlockAccountWithIdRequest, headers: Connect.Headers = [:]) async -> ResponseMessage<Account_V1_BlockAccountWithIdResponse> {
        return await self.client.unary(path: "/account.v1.AccountService/BlockAccountWithId", idempotencyLevel: .unknown, request: request, headers: headers)
    }

    @discardableResult
    public func `sendMessage`(request: Account_V1_SendMessageRequest, headers: Connect.Headers = [:], completion: @escaping @Sendable (ResponseMessage<Account_V1_SendMessageResponse>) -> Void) -> Connect.Cancelable {
        return self.client.unary(path: "/account.v1.AccountService/SendMessage", idempotencyLevel: .unknown, request: request, headers: headers, completion: completion)
    }

    @available(iOS 13, *)
    public func `sendMessage`(request: Account_V1_SendMessageRequest, headers: Connect.Headers = [:]) async -> ResponseMessage<Account_V1_SendMessageResponse> {
        return await self.client.unary(path: "/account.v1.AccountService/SendMessage", idempotencyLevel: .unknown, request: request, headers: headers)
    }

    public enum Metadata {
        public enum Methods {
            public static let getAccount = Connect.MethodSpec(name: "GetAccount", service: "account.v1.AccountService", type: .unary)
            public static let getAccountWithID = Connect.MethodSpec(name: "GetAccountWithId", service: "account.v1.AccountService", type: .unary)
            public static let generateAccountLinkPath = Connect.MethodSpec(name: "GenerateAccountLinkPath", service: "account.v1.AccountService", type: .unary)
            public static let resolveAccountLinkPath = Connect.MethodSpec(name: "ResolveAccountLinkPath", service: "account.v1.AccountService", type: .unary)
            public static let deleteAccount = Connect.MethodSpec(name: "DeleteAccount", service: "account.v1.AccountService", type: .unary)
            public static let addFriendWithID = Connect.MethodSpec(name: "AddFriendWithId", service: "account.v1.AccountService", type: .unary)
            public static let blockAccountWithID = Connect.MethodSpec(name: "BlockAccountWithId", service: "account.v1.AccountService", type: .unary)
            public static let sendMessage = Connect.MethodSpec(name: "SendMessage", service: "account.v1.AccountService", type: .unary)
        }
    }
}
