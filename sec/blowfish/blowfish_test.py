from Crypto.Cipher import Blowfish
from Crypto.Random import get_random_bytes
from Crypto.Util.Padding import pad, unpad

# Generate a random 56-byte key
key: bytes = get_random_bytes(56) # 密钥长度应该在4字节（32位）到56字节（448位）之间
print(f'Blowfish key: {len(key)}') 

# Create a new Blowfish cipher object with the key
cipher = Blowfish.new(key, Blowfish.MODE_ECB)

# The message to be encrypted
message = b'This is the message to be encrypted. '*1

# Blowfish input strings must be a multiple of 8 in length, so we might need to pad the message
# 不要直接在原始字符串后面加空格：无法区分是否是空格或者数据。通常会使用更复杂的填充方案，比如PKCS#7。
message = pad(message, Blowfish.block_size)

# Encrypt the message
ciphertext = cipher.encrypt(message)

print(f'Ciphertext: {ciphertext}')

# To decrypt, we create a new Blowfish cipher object
decipher = Blowfish.new(key, Blowfish.MODE_ECB)

# Then we can decrypt the ciphertext
plaintext = decipher.decrypt(ciphertext)

# Remove the padding
plaintext = unpad(plaintext, Blowfish.block_size)

print(f'Plaintext: {plaintext}')