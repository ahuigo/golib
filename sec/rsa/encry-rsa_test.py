import rsa
import binascii

def generate_key_from_pem():
    # 1. Convert the keys to PEM format
    keys = rsa.newkeys(1024) #　最长加密：1024/8 - 11 = 117 bytes
    public_key: rsa.PublicKey = keys[0]
    private_key: rsa.PrivateKey = keys[1]
    public_key_pem: str = public_key.save_pkcs1().decode()
    private_key_pem: str = private_key.save_pkcs1().decode()
    print("Public Key (PEM):\n", public_key_pem)
    print("Private Key (PEM):\n", private_key_pem)

    # 2. convert pem to rsa keys
    public_key: rsa.PrivateKey = rsa.PublicKey.load_pkcs1(public_key_pem)
    private_key: rsa.PrivateKey = rsa.PrivateKey.load_pkcs1(private_key_pem)
    return public_key, private_key

# (public_key,private_key) = rsa.newkeys(1024)
(public_key,private_key) = generate_key_from_pem()


# 任何人都可公解加密，不能验证身份(如果要验证身份，需要使用私钥签名)
def encrypt_with_pubkey():
    # 待加密的明文消息
    data: str = "h"*117
    message: bytes = data.encode('utf-8')
    # 加密消息
    encrypted_message: bytes = rsa.encrypt(message, public_key)
    # 解密消息
    decrypted_message = rsa.decrypt(encrypted_message, private_key)
    print(f'Decrypted message: {decrypted_message.decode("utf-8")}')

encrypt_with_pubkey()