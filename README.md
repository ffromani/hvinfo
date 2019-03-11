# hvinfo

hvinfo dumps the `HyperV` feature support bits as documented in the [HyperV Top Level Functional Specifications](https://github.com/MicrosoftDocs/Virtualization-Documentation/blob/live/tlfs/Hypervisor%20Top%20Level%20Functional%20Specification%20v5.0C.pdf).
The spec is [mirrored here verbatim  for convenience](docs/hyperv-tlfs-5.0c.pdf)
If you don't know what this does mean, you don't need `hvinfo` :)
hvinfo is (C) 2019 Red Hat Inc and licensed under the MIT license.

## compile

Cross compile on linux for windows:
```bash
GOOS=windows GOARCH=amd64 go build -o binaries/hvinfo.exe hvinfo.go
```

Should you need to compile for linux (negative test?):
```bash
go build -o binaries/hvinfo hvinfo.go
```
