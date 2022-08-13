#define IMPORT(mod, name) __attribute__((import_module(#mod), import_name(#name)))
#define EXPORT(name) __attribute__((export_name(#name)))
// #define WREN_API WASM_SYMBOL_EXPORTED 
#define DEBUG

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

static WrenForeignClassMethods _bindForeignClassFn(WrenVM* vm, const char* module, const char* className) {
    WrenForeignClassMethods result;
    bindForeignClassFn(vm, module, className, &(result.allocateUserData), &(result.finalizeUserData));
    if (result.allocateUserData) {
        result.allocate = dispatchForeignMethodFn;
        if (result.finalizeUserData) {
            result.finalize = dispatchFinalizerFn;
        }
    }
    return result;
}

int main()
{
    WrenConfiguration config;
    wrenInitConfiguration(&config);
    config.resolveModuleFn = resolveModuleFn;
    config.loadModuleFn = _loadModuleFn;
    config.bindForeignMethodFn = _bindForeignMethodFn;
    config.bindForeignClassFn = _bindForeignClassFn;
    config.writeFn = writeFn;
    config.errorFn = errorFn;
    vm = wrenNewVM(&config);
    
    return 0;
}