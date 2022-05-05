# qrcode
CLI: Generate QR codes embedding a string such as Bitcoin address

## disclaimer
>The use of this tool does not guarantee security or suitability
for any particular use. Please review the code and use at your own risk.

## installation
This step assumes you have [Go compiler toolchain](https://go.dev/dl/)
installed on your system.

```bash
go install github.com/kubetrail/qrcode@latest
```
Install shell completion. For instance `bash` completion can be installed
by adding following line to your `.bashrc`:
```bash
source <(qrcode completion bash)
```

## usage
In it's simples form a string can be encoded as a qrcode as follows:
```bash
echo this is input | qrcode gen
```

Below is another example showing how a new crypto mnemonic sentence is generated
using [bip39](https://github.com/kubetrail/bip39)
that is first saved in Google secrets engine
using [mksecret](https://github.com/kubetrail/mksecret), 
then passed to 
[bip32](https://github.com/kubetrail/bip32) tool
to generate Bitcoin address and then passed to `qrcode` to generate
a qrcode.

qrcode is displayed on STDOUT. Use `--output-filename` to write to
a file as `png` file.
```bash
bip39 gen --length=12 \
  | mksecret set --name="my-secret-mnemonic" \
  | bip32 gen --derivation-path=m/0h/0/0 \
  | jq -r '.addr' \
  | qrcode gen
```
