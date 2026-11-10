const donutvpn = require("./donutvpn_grpc_web_pb.js");
const extension = require("./extension_grpc_web_pb.js");

const grpcServerAddress = '/';
const extensionClient = new extension.ExtensionHostServicePromiseClient(grpcServerAddress, null, null);
const donutvpnClient = new donutvpn.CorePromiseClient(grpcServerAddress, null, null);

module.exports = { extensionClient ,donutvpnClient};