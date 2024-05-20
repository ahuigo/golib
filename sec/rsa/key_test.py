# python两种加密库, 加密方式不同, 但是都是基于RSA的
# 1. pycrypto: rsa, blowfish, des, aes, etc. 更灵活强大
# 2. rsa: 纯python rsa库, 更简单易用
def generate_key_from_pem_with_rsa():
    # 纯python rsa库
    import rsa
    # 1. Convert the keys to PEM format
    keys = rsa.newkeys(1024) #　非pem格式
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

def generate_key_with_crypto():
    # Crypto 是 pycrypto 库中导入 RSA 模块。pycrypto 是一个包含了许多不同的加密算法的库，包括 RSA, DSA, AES, DES, Blowfish, ElGamal, etc.
    from Crypto.PublicKey import RSA
    from Crypto.Cipher import PKCS1_OAEP
    key = RSA.generate(1024)
    private_key_pem = key.exportKey('PEM')
    public_key_pem = key.publickey().exportKey('PEM')

    print("Public key:", public_key_pem.decode())
    print("Private key:", private_key_pem.decode())

    # 2. convert pem to rsa keys
    private_key = RSA.importKey(private_key_pem)
    public_key = RSA.importKey(public_key_pem)
    # 3. PCKCS1 OAEP(Optimal Asymmetric Encryption Padding)padding
    # 这种填充方式是RSAES-OAEP，有更强的安全性:它是随机填充，每次加密相同数据结果不同
    private_key = PKCS1_OAEP.new(private_key)
    public_key = PKCS1_OAEP.new(public_key)
    return private_key, public_key

generate_key_with_crypto()