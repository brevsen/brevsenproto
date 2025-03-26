// swift-tools-version: 6.0
// The swift-tools-version declares the minimum version of Swift required to build this package.

import PackageDescription

let package = Package(
    name: "brevsenproto",
    products: [
        // Products define the executables and libraries a package produces, making them visible to other packages.
        .library(
            name: "brevsenproto",
            targets: ["brevsenproto"]),
    ],
    dependencies: [
        .package(url: "https://github.com/connectrpc/connect-swift", "1.0.2"..<"2.0.0"),
    ],
    targets: [
        // Targets are the basic building blocks of a package, defining a module or a test suite.
        // Targets can depend on other targets in this package and products from dependencies.
        .target(
            name: "brevsenproto",
            dependencies: [
                .product(name: "Connect", package: "connect-swift")
            ]
        ),
        .testTarget(
            name: "brevsenprotoTests",
            dependencies: ["brevsenproto"]
        ),
    ]
)
