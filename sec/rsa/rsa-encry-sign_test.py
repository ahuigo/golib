import rsa
from typing import NamedTuple
import json
class EncryData(NamedTuple):
    data: bytes
    sign: bytes

(public_key,private_key) = rsa.newkeys(512)

'''
任何人可公钥加密，所以要加上私钥签名，用来验证身份: 
- client 要用自己的私钥签名，用server的公解加密(在非对称加密中，每个参与者都应该有一对公钥和私钥)

但是这段代码有一个缺点：
- 密钥长度是n位，那么你可以加密的最大消息长度是n/8 - 11字节. 最多加密　512/8 - 11 = 53 字节
'''
def encrypt_with_sign(data: any, public_key: rsa.PublicKey, private_key: rsa.PrivateKey):
    def encrypt(message: bytes, public_key: rsa.PublicKey) -> bytes:
        return rsa.encrypt(message, public_key)
    def decrypt(encrypted_message: bytes, private_key: rsa.PrivateKey) -> bytes:
        return rsa.decrypt(encrypted_message, private_key)

    # 待加密的明文binary消息
    dataBin:bytes = json.dumps(data).encode('utf-8')
    # 加密消息
    encry_bytes: bytes = rsa.encrypt(dataBin, public_key)
    encry_sign: bytes = rsa.sign(encry_bytes, private_key, 'SHA-1')
    res = {
        "data": encry_bytes.hex(),
        "sign": encry_sign.hex(),
    }
    return json.dumps(res)

# 解密
def decrypt_with_sign(edataJson: str, public_key: rsa.PublicKey, private_key: rsa.PrivateKey):
    # 解析数据结构
    edataHex = json.loads(edataJson)
    edata: EncryData = EncryData(
        data=bytes.fromhex(edataHex["data"]),
        sign=bytes.fromhex(edataHex["sign"]),
    )

    # 使用公钥验证签名
    def verify_with_pubkey(message: bytes, signature: bytes):
        try:
            rsa.verify(message, signature, public_key)
        except rsa.VerificationError:
            raise("Signature verification failed.")
    # 验证签名
    verify_with_pubkey(edata.data, edata.sign)
    # 使用私钥解密消息
    databin = rsa.decrypt(edata.data, private_key)
    return json.loads(databin)

# 示例
(public_key,private_key) = rsa.newkeys(512)
data = "a"*50
sendData = encrypt_with_sign(data, public_key, private_key)
print("发送数据：", sendData)
oridata = decrypt_with_sign(sendData, public_key, private_key)
print("接收数据(解密)：", oridata)