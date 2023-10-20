//
//  Generated code. Do not modify.
//  source: api/prelude.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use errCodeDescriptor instead')
const ErrCode$json = {
  '1': 'ErrCode',
  '2': [
    {'1': 'ErrOk', '2': 0},
    {'1': 'ErrUnknown', '2': 101},
    {'1': 'ErrArgsInvalid', '2': 102},
    {'1': 'ErrArgsEmpty', '2': 103},
    {'1': 'ErrSystem', '2': 104},
    {'1': 'ErrDB', '2': 105},
    {'1': 'ErrNoServe', '2': 106},
  ],
};

/// Descriptor for `ErrCode`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List errCodeDescriptor = $convert.base64Decode(
    'CgdFcnJDb2RlEgkKBUVyck9rEAASDgoKRXJyVW5rbm93bhBlEhIKDkVyckFyZ3NJbnZhbGlkEG'
    'YSEAoMRXJyQXJnc0VtcHR5EGcSDQoJRXJyU3lzdGVtEGgSCQoFRXJyREIQaRIOCgpFcnJOb1Nl'
    'cnZlEGo=');

@$core.Deprecated('Use sTPreludePingReqDescriptor instead')
const STPreludePingReq$json = {
  '1': 'STPreludePingReq',
  '2': [
    {'1': 'name', '3': 1, '4': 1, '5': 9, '10': 'name'},
  ],
};

/// Descriptor for `STPreludePingReq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List sTPreludePingReqDescriptor = $convert.base64Decode(
    'ChBTVFByZWx1ZGVQaW5nUmVxEhIKBG5hbWUYASABKAlSBG5hbWU=');

@$core.Deprecated('Use sTPreludePingRspDescriptor instead')
const STPreludePingRsp$json = {
  '1': 'STPreludePingRsp',
  '2': [
    {'1': 'traceId', '3': 1, '4': 1, '5': 9, '10': 'traceId'},
    {'1': 'message', '3': 2, '4': 1, '5': 9, '10': 'message'},
  ],
};

/// Descriptor for `STPreludePingRsp`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List sTPreludePingRspDescriptor = $convert.base64Decode(
    'ChBTVFByZWx1ZGVQaW5nUnNwEhgKB3RyYWNlSWQYASABKAlSB3RyYWNlSWQSGAoHbWVzc2FnZR'
    'gCIAEoCVIHbWVzc2FnZQ==');

