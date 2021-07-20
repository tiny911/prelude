///
//  Generated code. Do not modify.
//  source: api/prelude.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields,deprecated_member_use_from_same_package

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use errCodeDescriptor instead')
const ErrCode$json = const {
  '1': 'ErrCode',
  '2': const [
    const {'1': 'ErrOk', '2': 0},
    const {'1': 'ErrUnknown', '2': 101},
    const {'1': 'ErrArgsInvalid', '2': 102},
    const {'1': 'ErrArgsEmpty', '2': 103},
    const {'1': 'ErrSystem', '2': 104},
    const {'1': 'ErrDB', '2': 105},
    const {'1': 'ErrNoServe', '2': 106},
  ],
};

/// Descriptor for `ErrCode`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List errCodeDescriptor = $convert.base64Decode('CgdFcnJDb2RlEgkKBUVyck9rEAASDgoKRXJyVW5rbm93bhBlEhIKDkVyckFyZ3NJbnZhbGlkEGYSEAoMRXJyQXJnc0VtcHR5EGcSDQoJRXJyU3lzdGVtEGgSCQoFRXJyREIQaRIOCgpFcnJOb1NlcnZlEGo=');
@$core.Deprecated('Use sTPreludePingReqDescriptor instead')
const STPreludePingReq$json = const {
  '1': 'STPreludePingReq',
  '2': const [
    const {'1': 'name', '3': 1, '4': 1, '5': 9, '10': 'name'},
  ],
};

/// Descriptor for `STPreludePingReq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List sTPreludePingReqDescriptor = $convert.base64Decode('ChBTVFByZWx1ZGVQaW5nUmVxEhIKBG5hbWUYASABKAlSBG5hbWU=');
@$core.Deprecated('Use sTPreludePingRspDescriptor instead')
const STPreludePingRsp$json = const {
  '1': 'STPreludePingRsp',
  '2': const [
    const {'1': 'traceId', '3': 1, '4': 1, '5': 9, '10': 'traceId'},
    const {'1': 'message', '3': 2, '4': 1, '5': 9, '10': 'message'},
  ],
};

/// Descriptor for `STPreludePingRsp`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List sTPreludePingRspDescriptor = $convert.base64Decode('ChBTVFByZWx1ZGVQaW5nUnNwEhgKB3RyYWNlSWQYASABKAlSB3RyYWNlSWQSGAoHbWVzc2FnZRgCIAEoCVIHbWVzc2FnZQ==');
