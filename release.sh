## Compile Go
cd go
go get gioui.org/cmd/gogio
go run gioui.org/cmd/gogio -target js -o $TEMP/satellites .
mkdir -p ../_html/go
mv $TEMP/satellites/wasm.js ../_html/go/wasm.js
mv $TEMP/satellites/main.wasm ../_html/go/main.wasm
cp static/* ../_html/go

cd ..

## Compile CSharp
cd csharp
dotnet publish -c Release
mkdir -p ../_html/csharp/
mkdir -p ../_html/csharp/_framework/
mv bin/Release/net6.0/publish/wwwroot/_framework/* ../_html/csharp/_framework/
cp -f index.html ../_html/csharp/