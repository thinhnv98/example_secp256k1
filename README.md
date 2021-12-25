# example_secp256k1
This is an example for a digital signature based on secp256k1

### Run project
Clone the project

Go to project directory
```
make run
```

### Generate private key & public key
```http request
GET  /generate-keys
```

* Response body:
```json
{
    "private_key": "c209bd0b0fd2317a8e4e269a83d150165adbe3dbc13dca6abadb10d01f87ef75",
    "public_key_compressed": "022c592450b4274268f92a3c2ad5a7448e3f75bc890889fdd60f0f2c47593ade15",
    "succeed": true
}
```

### Someone creates message to send to above address's information
```http request
POST /someone-create-mess
```

* Request body:
```json
{
    "public_key": "022c592450b4274268f92a3c2ad5a7448e3f75bc890889fdd60f0f2c47593ade15",
    "text": "message text"
}
```

* Response body:
```json
{
    "encrypted_message": "b8c6265e4f64b0d7177666bb73e99a9f02ca00208c6e17a3716e5c04d598757cce6ff9bdaa5cad9687776cb236bf8f024535f44a00200510e883af403f6744a42ad86370dc1441b6593b5c8f322a7c37326f1a9dad2e3df5e6fdb766f9edfc5a7720a717296c75733242a39652f505bc85bf0253eebdc63f681efb94851624700c3f49f9b60f",
    "succeed": true
}
```

### Owner decrypt the message with private key
```http request
POST /decrypt-mess
```

* Request body:
```json
{
    "private_key": "c209bd0b0fd2317a8e4e269a83d150165adbe3dbc13dca6abadb10d01f87ef75",
    "encrypted_message":"b8c6265e4f64b0d7177666bb73e99a9f02ca00208c6e17a3716e5c04d598757cce6ff9bdaa5cad9687776cb236bf8f024535f44a00200510e883af403f6744a42ad86370dc1441b6593b5c8f322a7c37326f1a9dad2e3df5e6fdb766f9edfc5a7720a717296c75733242a39652f505bc85bf0253eebdc63f681efb94851624700c3f49f9b60f"
}
```

* Response body:
```json
{
    "message": "message text",
    "succeed": true
}
```

### Owner sign a message
```http request
POST /i-create-mess
```

* Request body:
```json
{
    "private_key": "c209bd0b0fd2317a8e4e269a83d150165adbe3dbc13dca6abadb10d01f87ef75",
    "text": "This is good message"
}
```

* Response body:
```json
{
    "message_hash": "5468697320697320676f6f64206d657373616765",
    "signature": "30440220765c36158da3a69158588cd430acebbd167c81c32c9a1b2541f04443f4886c57022030a855bbd1cd3a23ed3415336ab9777343f1c569a5d191b18921eb5ecd3e134f",
    "succeed": true
}
```

### Everyone who have above address's public key can verify the signature
```http request
POST /verify-mess
```

* Request body:
```json
{
    "encrypted_message": "5468697320697320676f6f64206d657373616765",
    "signature": "3045022100c63be66cc4f5706688a5d856488d39cc3a328864d4e479ddfab50331d8c183d002202fd455b893378036066c7f0a6afebd27d07e468e2ec637977547c45565c74d07",
    "public_key": "035c1ae93b43e476f4741c3a1fe5581429ee9a20d278d4346a5e1e3817fda272f7"
}
```

* Response body:
```json
{
    "message": "This is good message",
    "succeed": true,
    "verified": true
}
```