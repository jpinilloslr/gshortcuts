package gsettings

/*
#cgo pkg-config: gio-2.0 glib-2.0
#include <gio/gio.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type GSettings struct {
	ptr *C.GSettings
}

func New(schema string) (*GSettings, error) {
	cSchema := C.CString(schema)
	defer C.free(unsafe.Pointer(cSchema))

	ptr := cNewSettings(cSchema)
	if ptr == nil {
		return nil, fmt.Errorf(
			`failed to create GSettings object with schema "%s"`,
			schema,
		)
	}

	return &GSettings{
		ptr: ptr,
	}, nil
}

func NewWithPath(schema, path string) (*GSettings, error) {
	cSchema := C.CString(schema)
	defer C.free(unsafe.Pointer(cSchema))
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	ptr := cNewSettingsWithPath(cSchema, cPath)
	if ptr == nil {
		return nil, fmt.Errorf(
			`failed to create GSettings object with schema "%s" and path "%s"`,
			schema,
			path,
		)
	}

	return &GSettings{
		ptr: ptr,
	}, nil
}

func (gs *GSettings) Close() {
	if gs.ptr != nil {
		cUnref(unsafe.Pointer(gs.ptr))
		gs.ptr = nil
	}
}

func (gs *GSettings) GetStringArray(key string) []string {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	arr := cGetStringArray(gs.ptr, cKey)
	if arr == nil {
		return []string{}
	}

	defer cFreeStrv(arr)

	var result []string

	for i := 0; ; i++ {
		ptr := *(**C.char)(
			unsafe.Pointer(
				uintptr(unsafe.Pointer(arr)) + uintptr(i)*unsafe.Sizeof(uintptr(0)),
			),
		)
		if ptr == nil {
			break
		}
		result = append(result, C.GoString(ptr))
	}

	return result
}

func (gs *GSettings) GetString(key string) string {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	value := cGetString(gs.ptr, cKey)
	defer cFree(unsafe.Pointer(value))

	return C.GoString(value)
}

func (gs *GSettings) SetStringArray(key string, values []string) error {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	cArr := C.malloc(C.size_t(len(values)+1) * C.size_t(unsafe.Sizeof(uintptr(0))))
	strv := (**C.char)(cArr)

	for i, val := range values {
		cstr := C.CString(val)
		ptr := (**C.char)(
			unsafe.Pointer(
				uintptr(cArr) + uintptr(i)*unsafe.Sizeof(uintptr(0)),
			),
		)
		*ptr = cstr
	}

	last := (**C.char)(
		unsafe.Pointer(
			uintptr(cArr) + uintptr(len(values))*unsafe.Sizeof(uintptr(0)),
		),
	)
	*last = nil

	ok := cSetStringArray(gs.ptr, cKey, strv) == C.TRUE

	for i := range values {
		ptr := *(**C.char)(
			unsafe.Pointer(
				uintptr(cArr) + uintptr(i)*unsafe.Sizeof(uintptr(0)),
			),
		)
		C.free(unsafe.Pointer(ptr))
	}
	C.free(cArr)

	if !ok {
		return fmt.Errorf(`failed to set string array for key "%s"`, key)
	}
	return nil
}

func (gs *GSettings) SetString(key string, value string) error {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	cVal := C.CString(value)
	defer C.free(unsafe.Pointer(cVal))

	ok := cSetString(gs.ptr, cKey, cVal) == C.TRUE

	if !ok {
		return fmt.Errorf(`failed to set string for key "%s"`, key)
	}
	return nil
}

func (gs *GSettings) Reset(key string) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	C.g_settings_reset(gs.ptr, cKey)
}

func (gs *GSettings) Sync() {
	cSync()
}
