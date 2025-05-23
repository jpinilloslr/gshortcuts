package gsettings

/*
#cgo pkg-config: gio-2.0 glib-2.0
#include <gio/gio.h>
*/
import "C"
import "unsafe"

func cNewSettings(schema *C.char) *C.GSettings {
	return C.g_settings_new(schema)
}

func cNewSettingsWithPath(schema *C.char, path *C.char) *C.GSettings {
	return C.g_settings_new_with_path(schema, path)
}

func cGetStringArray(settings *C.GSettings, key *C.char) **C.char {
	return C.g_settings_get_strv(settings, key)
}

func cGetString(s *C.GSettings, key *C.char) *C.char {
	return C.g_settings_get_string(s, key)
}

func cSetStringArray(s *C.GSettings, key *C.char, val **C.char) C.gboolean {
	return C.g_settings_set_strv(s, key, val)
}

func cSetString(s *C.GSettings, key, val *C.char) C.gboolean {
	return C.g_settings_set_string(s, key, val)
}

func cSync() {
	C.g_settings_sync()
}

func cFree(ptr unsafe.Pointer) {
	C.g_free(C.gpointer(ptr))
}

func cFreeStrv(arr **C.char) {
	C.g_strfreev(arr)
}

func cUnref(obj unsafe.Pointer) {
	C.g_object_unref(C.gpointer(obj))
}
