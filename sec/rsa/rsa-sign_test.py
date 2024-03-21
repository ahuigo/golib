import rsa

# Generate a RSA key pair
(pubkey, privkey) = rsa.newkeys(512)

# Message to be signed
message = b'This is the message to be signed. ' * 1000
print("len:", len(message))

# Sign the message with the private key
signature = rsa.sign(message, privkey, 'SHA-1')

# Verify the signature with the public key
try:
    rsa.verify(message, signature, pubkey)
    print('The signature is valid.')
except rsa.VerificationError:
    print('The signature is not valid.')