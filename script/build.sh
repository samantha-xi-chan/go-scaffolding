set -e

current_dir="$(cd "$(dirname "$0")" && pwd)"
echo "$current_dir：$current_dir"
verFile=$current_dir"/version"
Ver=$(cat $verFile) ;
echo $Ver

BuildT=$(date -u +'%Y-%m-%dT%H:%M:%SZ')
GitBranch=$(git branch --show-current)
GitCommit=$(git rev-parse --short HEAD)
echo "\$Ver:    "     $Ver        ;
echo "\$BuildT: "     $BuildT     ;
echo "\$GitBranch: "  $GitBranch  ;
echo "\$GitCommit: "  $GitCommit  ;

GOOS="darwin"; GOARCH="arm64"; APP="app01"; APP_NAME=$APP"-"$GOOS"-"$GOARCH; echo $APP_NAME
CGO_ENABLED=0 GOOS=$GOOS  GOARCH=$GOARCH go build -o bin/$APP_NAME       -ldflags "-X main.Version=$Ver -X main.BuildTime=$BuildT -X main.GitBranch=$GitBranch -X main.GitCommit=$GitCommit"   cmd/version.go cmd/app01.go ;

GOOS="linux";  GOARCH="amd64"; APP="app01"; APP_NAME=$APP"-"$GOOS"-"$GOARCH; echo $APP_NAME
CGO_ENABLED=0 GOOS=$GOOS  GOARCH=$GOARCH go build -o bin/$APP_NAME        -ldflags "-X main.Version=$Ver -X main.BuildTime=$BuildT -X main.GitBranch=$GitBranch -X main.GitCommit=$GitCommit"   cmd/version.go cmd/app01.go ;

