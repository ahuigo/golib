'''
# refer: post/py/py-crypt.md
# 有时我们想随机产生一个IV,  此时IV 应该像pading 一样`保存`到加密结果的头或者尾部
# Warn: 
# 1. 随机IV 会导致 aes-cbc 加密结果(加IV前)是随机的
# 2. ECB 会忽略IV 

## 固定iv的风险 ###
# 很多示例是用的内部默认的iv，固定的iv 会有风险:
# 1. https://stackoverflow.com/questions/3008139/why-is-using-a-non-random-iv-with-cbc-mode-a-vulnerability
# 2. https://crypto.stackexchange.com/questions/5094/is-aes-in-cbc-mode-secure-if-a-known-and-or-fixed-iv-is-used
'''

import base64    
from Crypto.Cipher import AES    
from Crypto.Hash import SHA256    
from Crypto import Random

def encrypt(key, source):    
    source = source.encode()    
    key = SHA256.new(key.encode()).digest()      
    IV = Random.new().read(AES.block_size)      
    encryptor = AES.new(key, AES.MODE_CBC, IV)    
    padding = AES.block_size - len(source) % AES.block_size      
    source += bytes([padding]) * padding    
    enc = encryptor.encrypt(source) 
    data = IV + enc # store the IV at the beginning and encrypt    
    return base64.b64encode(data).decode("utf8")    

def decrypt(key, source):    
    source = base64.b64decode(source.encode("utf8"))    
    key = SHA256.new(key.encode()).digest()      
    IV = source[:AES.block_size]  # extract the IV from the beginning    
    decryptor = AES.new(key, AES.MODE_CBC, IV)    
    data = decryptor.decrypt(source[AES.block_size:])  # decrypt    
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
