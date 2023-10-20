//
//  Generated code. Do not modify.
//  source: api/prelude.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

export 'prelude.pbenum.dart';

/// Ping请求
class STPreludePingReq extends $pb.GeneratedMessage {
  factory STPreludePingReq({
    $core.String? name,
  }) {
    final $result = create();
    if (name != null) {
      $result.name = name;
    }
    return $result;
  }
  STPreludePingReq._() : super();
  factory STPreludePingReq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory STPreludePingReq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'STPreludePingReq', package: const $pb.PackageName(_omitMessageNames ? '' : 'prelude'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'name')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  STPreludePingReq clone() => STPreludePingReq()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  STPreludePingReq copyWith(void Function(STPreludePingReq) updates) => super.copyWith((message) => updates(message as STPreludePingReq)) as STPreludePingReq;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static STPreludePingReq create() => STPreludePingReq._();
  STPreludePingReq createEmptyInstance() => create();
  static $pb.PbList<STPreludePingReq> createRepeated() => $pb.PbList<STPreludePingReq>();
  @$core.pragma('dart2js:noInline')
  static STPreludePingReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<STPreludePingReq>(create);
  static STPreludePingReq? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get name => $_getSZ(0);
  @$pb.TagNumber(1)
  set name($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasName() => $_has(0);
  @$pb.TagNumber(1)
  void clearName() => clearField(1);
}

/// Ping响应
class STPreludePingRsp extends $pb.GeneratedMessage {
  factory STPreludePingRsp({
    $core.String? traceId,
    $core.String? message,
  }) {
    final $result = create();
    if (traceId != null) {
      $result.traceId = traceId;
    }
    if (message != null) {
      $result.message = message;
    }
    return $result;
  }
  STPreludePingRsp._() : super();
  factory STPreludePingRsp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory STPreludePingRsp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'STPreludePingRsp', package: const $pb.PackageName(_omitMessageNames ? '' : 'prelude'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'traceId', protoName: 'traceId')
    ..aOS(2, _omitFieldNames ? '' : 'message')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  STPreludePingRsp clone() => STPreludePingRsp()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  STPreludePingRsp copyWith(void Function(STPreludePingRsp) updates) => super.copyWith((message) => updates(message as STPreludePingRsp)) as STPreludePingRsp;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static STPreludePingRsp create() => STPreludePingRsp._();
  STPreludePingRsp createEmptyInstance() => create();
  static $pb.PbList<STPreludePingRsp> createRepeated() => $pb.PbList<STPreludePingRsp>();
  @$core.pragma('dart2js:noInline')
  static STPreludePingRsp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<STPreludePingRsp>(create);
  static STPreludePingRsp? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get traceId => $_getSZ(0);
  @$pb.TagNumber(1)
  set traceId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasTraceId() => $_has(0);
  @$pb.TagNumber(1)
  void clearTraceId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get message => $_getSZ(1);
  @$pb.TagNumber(2)
  set message($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasMessage() => $_has(1);
  @$pb.TagNumber(2)
  void clearMessage() => clearField(2);
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
