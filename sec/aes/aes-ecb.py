'''
1. ecb 不需要iv
2.
    DES.block_size = 8
    AES.block_size = 16 # 16*8 = 128
'''
from Crypto.Cipher import AES
class AesEcb():
    def __init__(self, key):
        key = key if isinstance(key, bytes) else key.encode()
        self.aes = AES.new(self.pad(key), AES.MODE_ECB)

    def pad(self, data):
        pad_len = (16 - len(data) % 16)
        return data + bytes([pad_len]) * pad_len

    def encrypt(self, data):
        enc = self.aes.encrypt(self.pad(data))
        return enc

    def decrypt(self, data):
        de = self.aes.decrypt(data)
        return (de[:-de[-1]]) # remove padding

from Crypto.Random import get_random_bytes
# key = get_random_bytes(32)
key = 'password'
text = '数据'
enc = AesEcb(key).encrypt(text.encode())
print(enc)
dec = AesEcb(key).decrypt(enc).decode()
print(dec)