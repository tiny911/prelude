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

/// 服务错误码
/// 下列错误码会同步至commont/error/error.go文件中
class ErrCode extends $pb.ProtobufEnum {
  static const ErrCode ErrOk = ErrCode._(0, _omitEnumNames ? '' : 'ErrOk');
  static const ErrCode ErrUnknown = ErrCode._(101, _omitEnumNames ? '' : 'ErrUnknown');
  static const ErrCode ErrArgsInvalid = ErrCode._(102, _omitEnumNames ? '' : 'ErrArgsInvalid');
  static const ErrCode ErrArgsEmpty = ErrCode._(103, _omitEnumNames ? '' : 'ErrArgsEmpty');
  static const ErrCode ErrSystem = ErrCode._(104, _omitEnumNames ? '' : 'ErrSystem');
  static const ErrCode ErrDB = ErrCode._(105, _omitEnumNames ? '' : 'ErrDB');
  static const ErrCode ErrNoServe = ErrCode._(106, _omitEnumNames ? '' : 'ErrNoServe');

  static const $core.List<ErrCode> values = <ErrCode> [
    ErrOk,
    ErrUnknown,
    ErrArgsInvalid,
    ErrArgsEmpty,
    ErrSystem,
    ErrDB,
    ErrNoServe,
  ];

  static final $core.Map<$core.int, ErrCode> _byValue = $pb.ProtobufEnum.initByValue(values);
  static ErrCode? valueOf($core.int value) => _byValue[value];

  const ErrCode._($core.int v, $core.String n) : super(v, n);
}


const _omitEnumNames = $core.bool.fromEnvironment('protobuf.omit_enum_names');
