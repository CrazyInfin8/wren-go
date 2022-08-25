#define IMPORT(mod, name) __attribute__((import_module(#mod), import_name(#name)))
#define EXPORT(name) __attribute__((export_name(#name)))
#include <string.h>
#include <wren.h>
WrenVM* vm;

WrenVM* getVM() {
    return vm;
}

void IMPORT(wren, writeFn) writeFn(WrenVM* vm, const char* text);

void IMPORT(wren, errorFn) errorFn(WrenVM* vm, WrenErrorType type, const char* module, int line, const char* message);

const char *IMPORT(wren, resolveModuleFn) resolveModuleFn(WrenVM *vm, const char *importer, const char *name);

char* IMPORT(wren, loadModuleFn) loadModuleFn(WrenVM *vm, const char *name);

static void _loadModuleCompleteFn(WrenVM* vm, const char* name, struct WrenLoadModuleResult result) {
    if (result.userData) {
        free(result.userData);
    }
}

static WrenLoadModuleResult _loadModuleFn(WrenVM* vm, const char* name) {
    WrenLoadModuleResult result;
    char* source = loadModuleFn(vm, name);
    if (source) {
        result.source = source;
        result.onComplete = _loadModuleCompleteFn;
        result.userData = source;
    }
    return result;
}

void IMPORT(wren, dispatchForeignMethodFn) dispatchForeignMethodFn(WrenVM *vm, void *userData);

void* IMPORT(wren, bindForeignMethodFn) bindForeignMethodFn(WrenVM* vm,
    const char* module, const char* className, bool isStatic,
    const char* signature);

static WrenBindForeignMethodResult _bindForeignMethodFn(WrenVM* vm,
    const char* module, const char* className, bool isStatic,
    const char* signature) {
    WrenBindForeignMethodResult result;
    result.userData = bindForeignMethodFn(vm, module, className, isStatic, signature);
    if (result.userData) {
        result.executeFn = dispatchForeignMethodFn;
    }
    return result;
}

void IMPORT(wren, dispatchFinalizerFn) dispatchFinalizerFn(WrenVM *vm, void *data, void *userData);

void IMPORT(wren, bindForeignClassFn) bindForeignClassFn(WrenVM* vm, const char* module, const char* className, void** allocate, void** finalize);

// See: https://github.com/wren-lang/wren/issues/811
static void invalidConstructor(WrenVM *vm, void *userData) {
    const char err[] = "Foreign class does not have a constructor.";
    wrenEnsureSlots(vm, 1);
    wrenSetSlotBytes(vm, 0, err, sizeof(err));
    wrenAbortFiber(vm, 0);
}

static WrenForeignClassMethods _bindForeignClassFn(WrenVM* vm, const char* module, const char* className) {
    WrenForeignClassMethods result = { NULL };
    bindForeignClassFn(vm, module, className, &(result.allocateUserData), &(result.finalizeUserData));
    if (result.allocateUserData) {
        result.allocate = dispatchForeignMethodFn;
        if (result.finalizeUserData) {
            result.finalize = dispatchFinalizerFn;
        }
    } else {
        if (strcmp(module, "random") == 0 || strcmp(module, "meta") == 0) {
            return result;
        }
        result.allocate = invalidConstructor;
    }
    return result;
}

typedef struct HeapSettings {
    size_t initSize, minSize;
    int growthPercent;
} HeapSettings;

bool IMPORT(wren, heapSettings) heapSettings(HeapSettings* settings);

const int versionTuple[3] = {WREN_VERSION_MAJOR, WREN_VERSION_MINOR, WREN_VERSION_PATCH};
#define __VERSION_STRING(major, minor, patch) #major"."#minor"."#patch
#define _VERSION_STRING(major, minor, patch) __VERSION_STRING(major, minor, patch)
#define VERSION_STRING _VERSION_STRING(WREN_VERSION_MAJOR, WREN_VERSION_MINOR, WREN_VERSION_PATCH)
const char versionString[sizeof(VERSION_STRING)] = VERSION_STRING;

int main()
{
    WrenConfiguration config;
    wrenInitConfiguration(&config);
    HeapSettings settings;
    config.resolveModuleFn = resolveModuleFn;
    config.loadModuleFn = _loadModuleFn;
    config.bindForeignMethodFn = _bindForeignMethodFn;
    config.bindForeignClassFn = _bindForeignClassFn;
    config.writeFn = writeFn;
    config.errorFn = errorFn;
    settings.initSize = config.initialHeapSize;
    settings.minSize = config.minHeapSize;
    settings.growthPercent = config.heapGrowthPercent;
    if (heapSettings(&settings)) {
        config.initialHeapSize = settings.initSize;
        config.minHeapSize = settings.minSize;
        config.heapGrowthPercent = settings.growthPercent;
    }
    vm = wrenNewVM(&config);
    
    return 0;
}