# Inspired from https://medium.com/@ismailakkila/black-hat-python-encrypt-and-decrypt-with-rsa-cryptography-bd6df84d65bc
# Updated to use python3 bytes and pathlib

import zlib
import base64
from Crypto.PublicKey import RSA
from Crypto.Cipher import PKCS1_OAEP
from pathlib import Path

def generate_new_key_pair():
    new_key = RSA.generate(1024, e=65537)
    private_key = new_key.exportKey("PEM")
    public_key = new_key.publickey().exportKey("PEM")
    return public_key, private_key

#Our Encryption Function(pem key)
def encrypt_blob(blob:bytes, public_key: bytes):
    rsa_key = RSA.importKey(public_key)
    rsa_key = PKCS1_OAEP.new(rsa_key)

    encrypted = rsa_key.encrypt(blob)
    return base64.b64encode(encrypted)

#Our Decryption Function
def decrypt_blob(encrypted_blob:bytes, private_key:bytes):
    rsakey = RSA.importKey(private_key)
    rsakey = PKCS1_OAEP.new(rsakey)

    encrypted_blob = base64.b64decode(encrypted_blob)
    decrypted = rsakey.decrypt(encrypted_blob)
    return decrypted

if __name__ == "__main__":
    public_key,private_key = generate_new_key_pair()
    encrypted_msg = encrypt_blob(b'hello', public_key)
    decrypted_msg = decrypt_blob(encrypted_msg, private_key)
    print('cipher message :\n'+str(encrypted_msg))
    print('\nplain  message :\n'+str(decrypted_msg))