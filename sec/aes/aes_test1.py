from Crypto.Cipher import AES
from Crypto.Random import get_random_bytes

def aes_ecb():
    # Generate a random 256-bit key
    key = get_random_bytes(32)

    # Create a new AES cipher object with the key and AES.MODE_ECB mode
    cipher = AES.new(key, AES.MODE_ECB)

    # The message to be encrypted
    message = b'This is the message to be encrypted' * 100

    # AES input strings must be a multiple of 16 in length, so we might need to pad the message
    message = message + b' ' * (16 - len(message) % 16)

    # Encrypt the message
    ciphertext = cipher.encrypt(message)

    print(f'Ciphertext: {ciphertext}')

    # To decrypt, we create a new AES cipher object
    decipher = AES.new(key, AES.MODE_ECB)

    # Then we can decrypt the ciphertext
    plaintext = decipher.decrypt(ciphertext)

    # Remove the padding
    plaintext = plaintext.rstrip(b' ')

    print(f'Plaintext: {plaintext}')

# Call the function
aes_example()