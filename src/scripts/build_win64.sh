go env -w GOOS=windows

go build -o F:/projects/SweetMint/sm-box/box/bin/box.exe ./cmd/app

go build -o F:/projects/SweetMint/sm-box/box/sbin/tools/init.exe ./cmd/tools/init_cli