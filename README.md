# 🚨 REPOSITORY MOVED

This project has migrated to **Codeberg** for better alignment with open-source values.

## 👉 New Location: [codeberg.org/ilius/bip39-coder](https://codeberg.org/ilius/bip39-coder)

# bip39-coder

Command line tool for encoding/decoding BIP39 cryptocurrency seeds to/from binary data (which is much smaller)

# Install

```
go get github.com/ilius/bip39-coder
```

# Encode and encrypt your seed

Run this command:

```
bip39-coder | gpg -c | base64 
```

And it will ask for a password (for encryption with `gpg`), and you enter the password twice.
Then you need to type/paste your seed, and press Control+D
Then it will print something like this:

```
jA0EAwMCGXIUpj5/mw3uyScGjXLarDTGMna/Vo2z+fh1PHZX8FaMFg+oBP3/go+NimPL2u8JxBw=
```

# Decrypt and decode your seed

Run this command:

```
base64 -d | gpg -d | bip39-coder -d
```

Then type/paste the encoded encrypted seed (like the output of last command), then press Control+D
Then it will ask for password (unless it's stored by gpg-agent) and then print out your seed

## WARNING: it is strongly advisied against stroring your seed electrically (in a file, on your hard drive or flash drive), even in encrypted form!

# Generate a 12-word random seed

```
$ dd if=/dev/urandom bs=16 count=1 | bip39-coder -d
1+0 records in
1+0 records out
16 bytes copied, 8.5288e-05 s, 188 kB/s
corn cupboard often kid wonder example dignity process file author alien above
```
