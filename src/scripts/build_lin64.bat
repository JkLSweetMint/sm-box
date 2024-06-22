go env -w GOOS=linux

go build -o F:/projects/SweetMint/sm-box/box/bin/box ./cmd/app

go build -o F:/projects/SweetMint/sm-box/box/sbin/tools/init ./cmd/tools/init_cli
