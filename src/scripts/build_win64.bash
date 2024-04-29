go env -w GOOS=linux

go build -o F:/projects/SweetMint/sm-box/box/bin/box ./cmd/app

go build -o F:/projects/SweetMint/sm-box/box/sbin/init ./cmd/system/init_script