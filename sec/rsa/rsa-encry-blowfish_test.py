import rsa
from Crypto.Cipher import Blowfish
from Crypto.Random import get_random_bytes
from Crypto.Util.Padding import pad, unpad
import json

# 生成RSA密钥对
(public_key, private_key) = rsa.newkeys(1024) # 加密数据长度最大：1024/8 - 11 = 117 字节

def encrypt_with_sign(dataObj: any, public_key: rsa.PublicKey, private_key: rsa.PrivateKey):
    data = json.dumps(dataObj).encode('utf-8')
    # 生成随机的Blowfish密钥
    blowfish_key: bytes = get_random_bytes(55)  # Blowfish密钥长度限制4-56字节

    # 使用Blowfish密钥加密数据(pkcs7 对齐模式)
    cipher = Blowfish.new(blowfish_key, Blowfish.MODE_ECB)
    data = pad(data, 8, 'pkcs7')
    encry_data = cipher.encrypt(data)  

    # 使用RSA公钥加密Blowfish密钥
    encry_key = rsa.encrypt(blowfish_key, public_key)

    # 使用RSA私钥签名加密的Blowfish密钥
    sign = rsa.sign(encry_key, private_key, 'SHA-1')

    # 序列化加密的数据、加密的Blowfish密钥和签名
    res = {
        "data": encry_data.hex(),
        "key": encry_key.hex(),
        "sign": sign.hex(),
    }
    return json.dumps(res)

def decrypt_with_sign(edata_json: str, public_key: rsa.PublicKey, private_key: rsa.PrivateKey):
    # 解析数据结构
    edata = json.loads(edata_json)
    encry_data = bytes.fromhex(edata["data"])
    encry_key = bytes.fromhex(edata["key"])
    sign = bytes.fromhex(edata["sign"])

    # 使用公钥验证签名
    try:
        rsa.verify(encry_key, sign, public_key)
    except rsa.VerificationError:
        raise Exception("Signature verification failed.")

    # 使用私钥解密Blowfish密钥
    blowfish_key = rsa.decrypt(encry_key, private_key)

    # 使用Blowfish密钥解密数据
    cipher = Blowfish.new(blowfish_key, Blowfish.MODE_ECB)
    data = cipher.decrypt(encry_data).rstrip()  # 移除可能的填充

    return data.decode('utf-8')

# 示例
data = {
    "uuid":"xxx",
    "timestamp": 7234567890,
    "timeout": 3600,
}
send_data = encrypt_with_sign(data, public_key, private_key)
print("发送数据：", send_data)
oridata = decrypt_with_sign(send_data, public_key, private_key)
print("接收数据(解密)：", oridata)