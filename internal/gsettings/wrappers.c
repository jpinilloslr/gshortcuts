#include <stdlib.h>
#include <gio/gio.h>

// TODO: let's avoid redundant allocations of GSettings by passing it
// as a parameter instead of creating it inside the function.
int gsettings_key_is_modified(const char* schema, const char* key, char** error_out) {
    GSettings *s = g_settings_new(schema);
    if (!s) {
        *error_out = g_strdup("g_settings_new() failed");
        return -1;
    }

    GSettingsSchemaSource *src = g_settings_schema_source_get_default();
    if (!src) {
        *error_out = g_strdup("schema source not found");
        g_object_unref(s);
        return -1;
    }

    GSettingsSchema *sch = g_settings_schema_source_lookup(src, schema, TRUE);
    if (!sch) {
        *error_out = g_strdup("schema not found");
        g_object_unref(s);
        return -1;
    }

    GSettingsSchemaKey *k = g_settings_schema_get_key(sch, key);
    if (!k) {
        *error_out = g_strdup("schema key not found");
        g_object_unref(s);
        return -1;
    }

    GVariant *def = g_settings_schema_key_get_default_value(k);
    GVariant *cur = g_settings_get_value(s, key);
    if (!def || !cur) {
        *error_out = g_strdup("failed to read current or default value");
        g_object_unref(s);
        return -1;
    }

    gboolean changed = !g_variant_equal(def, cur);
    g_variant_unref(def);
    g_variant_unref(cur);
    g_object_unref(s);
    return changed ? 1 : 0;
}


char **gsettings_list_schema_keys_by_schema(const char *schema_id, char **error_out) {
    if (!schema_id) {
        *error_out = g_strdup("schema_id is NULL");
        return NULL;
    }

    GSettingsSchemaSource *src = g_settings_schema_source_get_default();
    if (!src) {
        *error_out = g_strdup("no default GSettingsSchemaSource");
        return NULL;
    }

    GSettingsSchema *schema = 
        g_settings_schema_source_lookup(src, schema_id, TRUE);
    if (!schema) {
        *error_out = g_strdup("schema not found: ");
        *error_out = g_strconcat(*error_out, schema_id, NULL);
        return NULL;
    }

    return g_settings_schema_list_keys(schema);
}
