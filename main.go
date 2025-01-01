package main

/*
#cgo CFLAGS: -I./pkg/lua-5.4.7/src
#cgo LDFLAGS: ./pkg/lua-5.4.7/src/liblua.a -lm

#include "pkg/lua-5.4.7/src/lua.h"
#include "pkg/lua-5.4.7/src/lauxlib.h"
#include "pkg/lua-5.4.7/src/lualib.h"
#include <stdlib.h>

static lua_State* get_lua_state() {
    lua_State *L = luaL_newstate();
    luaL_openlibs(L);
    return L;
}

static int run_lua(lua_State *L, const char *s) {
    return luaL_loadstring(L, s) || lua_pcall(L, 0, LUA_MULTRET, 0);
}

static int run_lua_file(lua_State *L, const char *filename) {
    return (luaL_loadfile(L, filename) || lua_pcall(L, 0, LUA_MULTRET, 0));
}

static const char* get_lua_error(lua_State *L) {
    return lua_tostring(L, -1);
}
*/
import "C"

import (
	"fmt"
	"os"
	"unsafe"
)

func main() {
	L := C.get_lua_state()
	defer C.lua_close(L)

	filename := os.Args[1]

	filenameC := C.CString(filename)
	defer C.free(unsafe.Pointer(filenameC))

	if C.run_lua_file(L, filenameC) != 0 {
		errMsg := C.GoString(C.get_lua_error(L))
		fmt.Println(errMsg)
	}
}
