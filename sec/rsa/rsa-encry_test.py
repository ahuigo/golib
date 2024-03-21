import rsa
import binascii

data: str = "12345678910"

(public_key,private_key) = rsa.newkeys(2024)

# 任何人都可公解加密，不能验证身份(如果要验证身份，需要使用私钥签名)
def encrypt_with_pubkey():
    def encrypt(message: bytes, public_key: rsa.PublicKey) -> bytes:
        return rsa.encrypt(message, public_key)

    def decrypt(encrypted_message: bytes, private_key: rsa.PrivateKey) -> bytes:
        return rsa.decrypt(encrypted_message, private_key)
    # 待加密的明文消息
    message: bytes = data.encode('utf-8')
    # 加密消息
    encrypted_message: bytes = encrypt(message, public_key)
    # 解密消息
    decrypted_message = rsa.decrypt(encrypted_message, private_key)
    print(f'Decrypted message: {decrypted_message.decode("utf-8")}')

encrypt_with_pubkey()

