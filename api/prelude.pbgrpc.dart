//
//  Generated code. Do not modify.
//  source: api/prelude.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'package:protobuf/protobuf.dart' as $pb;

import 'prelude.pb.dart' as $0;

export 'prelude.pb.dart';

@$pb.GrpcServiceName('prelude.Prelude')
class PreludeClient extends $grpc.Client {
  static final _$ping = $grpc.ClientMethod<$0.STPreludePingReq, $0.STPreludePingRsp>(
      '/prelude.Prelude/Ping',
      ($0.STPreludePingReq value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.STPreludePingRsp.fromBuffer(value));

  PreludeClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options,
        interceptors: interceptors);

  $grpc.ResponseFuture<$0.STPreludePingRsp> ping($0.STPreludePingReq request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$ping, request, options: options);
  }
}

@$pb.GrpcServiceName('prelude.Prelude')
abstract class PreludeServiceBase extends $grpc.Service {
  $core.String get $name => 'prelude.Prelude';

  PreludeServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.STPreludePingReq, $0.STPreludePingRsp>(
        'Ping',
        ping_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.STPreludePingReq.fromBuffer(value),
        ($0.STPreludePingRsp value) => value.writeToBuffer()));
  }

  $async.Future<$0.STPreludePingRsp> ping_Pre($grpc.ServiceCall call, $async.Future<$0.STPreludePingReq> request) async {
    return ping(call, await request);
  }

  $async.Future<$0.STPreludePingRsp> ping($grpc.ServiceCall call, $0.STPreludePingReq request);
}
