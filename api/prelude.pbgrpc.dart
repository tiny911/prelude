///
//  Generated code. Do not modify.
//  source: api/prelude.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'prelude.pb.dart' as $0;
export 'prelude.pb.dart';

class PreludeClient extends $grpc.Client {
  static final _$ping =
      $grpc.ClientMethod<$0.STPreludePingReq, $0.STPreludePingRsp>(
          '/prelude.Prelude/Ping',
          ($0.STPreludePingReq value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.STPreludePingRsp.fromBuffer(value));

  PreludeClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$0.STPreludePingRsp> ping($0.STPreludePingReq request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$ping, request, options: options);
  }
}

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

  $async.Future<$0.STPreludePingRsp> ping_Pre($grpc.ServiceCall call,
      $async.Future<$0.STPreludePingReq> request) async {
    return ping(call, await request);
  }

  $async.Future<$0.STPreludePingRsp> ping(
      $grpc.ServiceCall call, $0.STPreludePingReq request);
}
