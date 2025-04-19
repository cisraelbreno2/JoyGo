#!/bin/bash

set -e

APP_NAME="joyGoMacos"
DIST_DIR="dist/JoyGoApp"
LIBS_DIR="libs"
CONFIGS_DIR="config"

echo "🔨 Compilando o projeto com SDL2..."
CGO_ENABLED=1 \
CC=clang \
CGO_CFLAGS="-I$(pwd)/SDL2/include" \
CGO_LDFLAGS="-L$(pwd)/$LIBS_DIR -lSDL2 -Wl,-rpath,@executable_path/$LIBS_DIR" \
GOOS=darwin GOARCH=arm64 \
go build -o $APP_NAME

echo "📦 Preparando pasta final em $DIST_DIR..."
rm -rf "$DIST_DIR"
mkdir -p "$DIST_DIR/$LIBS_DIR"

cp "$APP_NAME" "$DIST_DIR/"
cp "$LIBS_DIR"/*.dylib "$DIST_DIR/$LIBS_DIR/"

if [ -d "$CONFIGS_DIR" ]; then
    echo "📁 Copiando configs..."
    cp -R "$CONFIGS_DIR" "$DIST_DIR/"
fi

echo "🔧 Corrigindo caminho da SDL2 com install_name_tool..."
install_name_tool -change \
/opt/homebrew/opt/sdl2/lib/libSDL2-2.0.0.dylib \
@executable_path/$LIBS_DIR/libSDL2-2.0.0.dylib \
"$DIST_DIR/$APP_NAME"

echo "🔍 Verificando dependências do binário final:"
otool -L "$DIST_DIR/$APP_NAME"

echo "✅ Pronto! App empacotado em: $DIST_DIR"
