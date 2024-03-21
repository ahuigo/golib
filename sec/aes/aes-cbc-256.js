engine = 'aes-256-cbc'; //aes-256-ecb, aes192, ...
crypto = require('crypto')
log= console.log
ivx = true
class Aes{
    static encryptAes256Cbc(text, _key){
        var cipher;
        let m = crypto.createHash('sha256');
        m.update(_key);
        let key = m.digest();
        let iv = Buffer.from([...Array(16).keys()])
        cipher = crypto.createCipheriv(engine, key, iv);
        //cipher = crypto.createCipher(engine, key); // deprecated: iv 复用风险
        cipher.setAutoPadding(true)
        let ciph = cipher.update(text, 'utf8', 'base64'); //base64, hex ...
        ciph += cipher.final('base64');
        return ciph;
    }

    static decryptAes256Cbc(text, _key){
        var decipher;
        let m = crypto.createHash('sha256');
        m.update(_key);
        let key = m.digest();
        let iv = Buffer.from([...Array(16).keys()])
        decipher = crypto.createDecipheriv(engine, key, iv);
        // decipher.setAutoPadding(false)
        let txt = decipher.update(text,'base64', 'utf8');
        txt  = txt + decipher.final('utf8');
        return txt;
    } 
}

text = '数据'
pass = 'password'
enc = Aes.encryptAes256Cbc(text, pass)
console.log(enc)
dec = Aes.decryptAes256Cbc(enc, pass)
console.log(enc, dec)