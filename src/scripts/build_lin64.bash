go env -w GOOS=windows

go build -o F:/projects/SweetMint/sm-box/box/bin/box.exe ./cmd/app

go build -o F:/projects/SweetMint/sm-box/box/sbin/init.exe ./cmd/system/init_script