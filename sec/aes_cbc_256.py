# py-crypt-.md
# pip3 install pycryptodome
import base64
from Crypto.Cipher import AES
from Crypto.Hash import SHA256
from Crypto import Random

def encrypt(key, source):
    source = source.encode()
    key = SHA256.new(key.encode()).digest()  
    IV = bytes(range(16))
    encryptor = AES.new(key, AES.MODE_CBC, IV)
    padding = AES.block_size - len(source) % AES.block_size  
    source += bytes([padding]) * padding 
    data = encryptor.encrypt(source)  
    return base64.b64encode(data).decode("utf8")

def decrypt(key, source):
    source = base64.b64decode(source.encode("utf8"))
    key = SHA256.new(key.encode()).digest()  
    IV = bytes(range(16))
    decryptor = AES.new(key, AES.MODE_CBC, IV)
    data = decryptor.decrypt(source)  
    padding = data[-1]  # pick the padding value from the end;
    if data[-padding:] != bytes([padding]) * padding:
        raise ValueError("Invalid padding...")
    return data[:-padding].decode()  # remove the padding


my_password = "password"
my_data = "数据"

encrypted = encrypt(my_password, my_data)    
decrypted = decrypt(my_password, encrypted)    

print("key:  {}".format(my_password))    
print("data: {}".format(my_data))    
print("enc:  {}".format(encrypted))    
print("dec:  {}".format(decrypted))    
print("data match: {}".format(my_data == decrypted))
