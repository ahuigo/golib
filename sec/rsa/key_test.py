# python两种加密库, 加密方式不同, 但是都是基于RSA的
# 1. pycrypto: rsa, blowfish, des, aes, etc. 更灵活强大
# 2. rsa: 纯python rsa库, 更简单易用
def generate_pcks1_pem_with_rsa():
    # 纯python rsa库
    import rsa
    # 1. Convert the keys to PEM format
    keys = rsa.newkeys(1024) #　非pem格式
    public_key: rsa.PublicKey = keys[0]
    private_key: rsa.PrivateKey = keys[1]
    public_key_pem: bytes= public_key.save_pkcs1()
    private_key_pem: bytes = private_key.save_pkcs1()
    print("PKCS1:")
    print("publicKeyPem = []byte(`"+public_key_pem.decode()+'`)')
    print("privateKeyPem = []byte(`"+private_key_pem.decode()+"`)")

    # public_key_pem, private_key_pem = generate_pki_pem_with_cryptography()
    # public_key_pem, private_key_pem = generate_pcks1_pem_with_crypto()

    # 2. convert pem to rsa keys
    public_key: rsa.PrivateKey = rsa.PublicKey.load_pkcs1(public_key_pem)
    private_key: rsa.PrivateKey = rsa.PrivateKey.load_pkcs1(private_key_pem)
    # return public_key, private_key
    return public_key_pem, private_key_pem

def generate_pki_pem_with_cryptography():
    from cryptography.hazmat.primitives import serialization
    from cryptography.hazmat.primitives.asymmetric import rsa
    from cryptography.hazmat.backends import default_backend

    # Generate a private/public RSA key
    private_key = rsa.generate_private_key(
        public_exponent=65537,
        key_size=2048,
        backend=default_backend()
    )
    public_key = private_key.public_key()

    # Save the private/public key in PEM format
    private_key_pem = private_key.private_bytes(
        encoding=serialization.Encoding.PEM,
        format=serialization.PrivateFormat.PKCS8,
        encryption_algorithm=serialization.NoEncryption()
    )
    public_key_pem = public_key.public_bytes(
        encoding=serialization.Encoding.PEM,
        format=serialization.PublicFormat.SubjectPublicKeyInfo
    )

    # public_key_pem, private_key_pem = generate_pcks1_pem_with_rsa()

    print("PKI:")
    print("publicKeyPem = []byte(`" + public_key_pem.decode() + '`)')
    print("privateKeyPem = []byte(`" + private_key_pem.decode() + "`)")

    # load public_key, private_key
    private_key = serialization.load_pem_private_key(
        private_key_pem,
        password=None,
        backend=default_backend()
    )
    public_key = serialization.load_pem_public_key(
        public_key_pem,
        backend=default_backend()
    )
    return public_key_pem, private_key_pem



def generate_pcks1_pem_with_crypto():
    # Crypto 是 pycrypto 库中导入 RSA 模块。pycrypto 是一个包含了许多不同的加密算法的库，包括 RSA, DSA, AES, DES, Blowfish, ElGamal, etc.
    from Crypto.PublicKey import RSA
    from Crypto.Cipher import PKCS1_OAEP
    key = RSA.generate(1024)
    private_key_pem = key.exportKey('PEM')
    public_key_pem = key.publickey().exportKey('PEM')

    print("PKI")
    print("publicKeyPem = []byte(`"+public_key_pem.decode()+'`)')
    print("privateKeyPem = []byte(`"+private_key_pem.decode()+"`)")

    # mock
    # public_key_pem, private_key_pem = generate_pcks1_pem_with_rsa()
    # public_key_pem, private_key_pem = generate_pki_pem_with_cryptography()
    # public_key_pem, private_key_pem = generate_pcks1_pem_with_crypto()

    # 2. convert pki pem to obj(PKCS1_OAEP 即接受PKCS1填充方式，也接受PKI OAEP 或PKI其它)
    private_key = RSA.importKey(private_key_pem)
    public_key = RSA.importKey(public_key_pem)
    # 3. PKCS1 OAEP(Optimal Asymmetric Encryption Padding)# 这种填充方式是证书PKI格式，有更强的安全性:它是随机填充，每次加密相同数据结果不同
    private_key = PKCS1_OAEP.new(private_key)
    public_key = PKCS1_OAEP.new(public_key)
    # return private_key, public_key
    return public_key_pem, private_key_pem

# 1. 只load　PKCS1 pem
generate_pcks1_pem_with_rsa()

# 2. load多种格式(PCKS1, PKI)
# generate_pki_pem_with_cryptography()
# generate_pcks1_pem_with_crypto()