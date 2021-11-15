# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/sphinx-agent.proto](#npool/sphinx-agent.proto)
    - [AccountBalance](#sphinx.agent.v1.AccountBalance)
    - [GetBalanceRequest](#sphinx.agent.v1.GetBalanceRequest)
    - [GetTxStatusRequest](#sphinx.agent.v1.GetTxStatusRequest)
    - [GetTxStatusResponse](#sphinx.agent.v1.GetTxStatusResponse)
    - [VersionResponse](#sphinx.agent.v1.VersionResponse)
  
    - [SphinxAgent](#sphinx.agent.v1.SphinxAgent)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/sphinx-agent.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/sphinx-agent.proto



<a name="sphinx.agent.v1.AccountBalance"></a>

### AccountBalance
GetBalance 返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| coin_id | [int32](#int32) |  |  |
| address | [string](#string) |  | 查询的钱包地址 |
| timestamp_utc | [int64](#int64) |  | 长整型时间戳 |
| amount_float64 | [double](#double) |  | 不入库的参考金额 |
| amount_uint64 | [uint64](#uint64) |  | 内部交互标准金额格式 |






<a name="sphinx.agent.v1.GetBalanceRequest"></a>

### GetBalanceRequest
GetBalance 参数


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| coin_id | [int32](#int32) |  |  |
| address | [string](#string) |  | 查询的钱包地址 |
| timestamp_utc | [int64](#int64) |  | 长整型时间戳 |






<a name="sphinx.agent.v1.GetTxStatusRequest"></a>

### GetTxStatusRequest
GetTxStatus 参数


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transaction_id_chain | [string](#string) |  |  |






<a name="sphinx.agent.v1.GetTxStatusResponse"></a>

### GetTxStatusResponse
GetTxStatus 返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| amount_float64 | [double](#double) |  | 不入库的参考金额 |
| amount_uint64 | [uint64](#uint64) |  | 内部交互标准金额格式 |
| address_from | [string](#string) |  | 发送方 |
| address_to | [string](#string) |  | 接收方 |
| transaction_id_chain | [string](#string) |  | 公链交易ID |
| createtime_utc | [int64](#int64) |  | 创建时间 |
| updatetime_utc | [int64](#int64) |  | 上次更新时间 |
| is_success | [bool](#bool) |  | 便于调用方判断 |
| is_failed | [bool](#bool) |  | 不success不fail就是pending了 |
| num_blocks_confirmed | [int32](#int32) |  | 已确认区块数 |






<a name="sphinx.agent.v1.VersionResponse"></a>

### VersionResponse
request body and response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |





 

 

 


<a name="sphinx.agent.v1.SphinxAgent"></a>

### SphinxAgent
钱包代理服务

从RabbitMQ获取消息通知
从MySQL拉取交易，以及交易的详细信息
根据不同交易状态进行不同操作
每步操作完成后同步更新数据库状态
审核通过交易，调用离线签名
离线签名后，交由钱包节点进行广播
广播成功后获取到CID，更新数据库标记为完成

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetTxStatus | [GetTxStatusRequest](#sphinx.agent.v1.GetTxStatusRequest) | [GetTxStatusResponse](#sphinx.agent.v1.GetTxStatusResponse) | 交易状态查询 |
| GetBalance | [GetBalanceRequest](#sphinx.agent.v1.GetBalanceRequest) | [AccountBalance](#sphinx.agent.v1.AccountBalance) | 余额查询 |
| Version | [.google.protobuf.Empty](#google.protobuf.Empty) | [VersionResponse](#sphinx.agent.v1.VersionResponse) | Method Version |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

