///
//  Generated code. Do not modify.
//  source: api/prelude.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

// ignore_for_file: UNDEFINED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class ErrCode extends $pb.ProtobufEnum {
  static const ErrCode ErrOk = ErrCode._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'ErrOk');
  static const ErrCode ErrUnknown = ErrCode._(101, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'ErrUnknown');
  static const ErrCode ErrArgsInvalid = ErrCode._(102, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'ErrArgsInvalid');
  static const ErrCode ErrArgsEmpty = ErrCode._(103, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'ErrArgsEmpty');
  static const ErrCode ErrSystem = ErrCode._(104, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'ErrSystem');
  static const ErrCode ErrDB = ErrCode._(105, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'ErrDB');
  static const ErrCode ErrNoServe = ErrCode._(106, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'ErrNoServe');

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

