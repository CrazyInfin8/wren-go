package internals

func (c *Context) WrenGetUserData(p0 int32) int32 {
	return f681(c, p0)
}

func (c *Context) WrenInitConfiguration(p0 int32) {
	f642(c, p0)
}

func (c *Context) Realloc(p0 int32, p1 int32) int32 {
	return f644(c, p0, p1)
}

func (c *Context) WrenSetSlotNewList(p0 int32, p1 int32) {
	f663(c, p0, p1)
}

func (c *Context) WrenGetMapCount(p0 int32, p1 int32) int32 {
	return f672(c, p0, p1)
}

func (c *Context) WrenGetVariable(p0 int32, p1 int32, p2 int32, p3 int32) {
	f677(c, p0, p1, p2, p3)
}

func (c *Context) WrenHasModule(p0 int32, p1 int32) int32 {
	return f679(c, p0, p1)
}

func (c *Context) Malloc(p0 int32) int32 {
	return f684(c, p0)
}

func (c *Context) WrenInterpret(p0 int32, p1 int32, p2 int32) int32 {
	return f639(c, p0, p1, p2)
}

func (c *Context) WrenCollectGarbage(p0 int32) {
	f640(c, p0)
}

func (c *Context) WrenReleaseHandle(p0 int32, p1 int32) {
	f649(c, p0, p1)
}

func (c *Context) WrenSetSlotDouble(p0 int32, p1 int32, p2 float64) {
	f661(c, p0, p1, p2)
}

func (c *Context) WrenSetSlotHandle(p0 int32, p1 int32, p2 int32) {
	f667(c, p0, p1, p2)
}

func (c *Context) WrenRemoveMapValue(p0 int32, p1 int32, p2 int32, p3 int32) {
	f676(c, p0, p1, p2, p3)
}

func (c *Context) WrenMakeCallHandle(p0 int32, p1 int32) int32 {
	return f647(c, p0, p1)
}

func (c *Context) WrenCall(p0 int32, p1 int32) int32 {
	return f648(c, p0, p1)
}

func (c *Context) WrenGetSlotForeign(p0 int32, p1 int32) int32 {
	return f656(c, p0, p1)
}

func (c *Context) Calloc(p0 int32, p1 int32) int32 {
	return f685(c, p0, p1)
}

func (c *Context) GetVM() int32 {
	return f683(c)
}

func (c *Context) Start() {
	f638(c)
}

func (c *Context) WrenEnsureSlots(p0 int32, p1 int32) {
	f651(c, p0, p1)
}

func (c *Context) WrenGetSlotBool(p0 int32, p1 int32) int32 {
	return f653(c, p0, p1)
}

func (c *Context) WrenGetSlotString(p0 int32, p1 int32) int32 {
	return f657(c, p0, p1)
}

func (c *Context) WrenGetListElement(p0 int32, p1 int32, p2 int32, p3 int32) {
	f669(c, p0, p1, p2, p3)
}

func (c *Context) WrenSetListElement(p0 int32, p1 int32, p2 int32, p3 int32) {
	f670(c, p0, p1, p2, p3)
}

func (c *Context) WrenGetSlotBytes(p0 int32, p1 int32, p2 int32) int32 {
	return f654(c, p0, p1, p2)
}

func (c *Context) WrenSetSlotNull(p0 int32, p1 int32) {
	f665(c, p0, p1)
}

func (c *Context) WrenGetMapContainsKey(p0 int32, p1 int32, p2 int32) int32 {
	return f673(c, p0, p1, p2)
}

func (c *Context) WrenGetMapValue(p0 int32, p1 int32, p2 int32, p3 int32) {
	f674(c, p0, p1, p2, p3)
}

func (c *Context) WrenSetMapValue(p0 int32, p1 int32, p2 int32, p3 int32) {
	f675(c, p0, p1, p2, p3)
}

func (c *Context) Free(p0 int32) {
	f643(c, p0)
}

func (c *Context) WrenSetSlotBool(p0 int32, p1 int32, p2 int32) {
	f659(c, p0, p1, p2)
}

func (c *Context) WrenSetSlotNewForeign(p0 int32, p1 int32, p2 int32, p3 int32) int32 {
	return f662(c, p0, p1, p2, p3)
}

func (c *Context) WrenGetSlotType(p0 int32, p1 int32) int32 {
	return f652(c, p0, p1)
}

func (c *Context) WrenSetSlotBytes(p0 int32, p1 int32, p2 int32, p3 int32) {
	f660(c, p0, p1, p2, p3)
}

func (c *Context) WrenSetSlotNewMap(p0 int32, p1 int32) {
	f664(c, p0, p1)
}

func (c *Context) WrenInsertInList(p0 int32, p1 int32, p2 int32, p3 int32) {
	f671(c, p0, p1, p2, p3)
}

func (c *Context) WrenHasVariable(p0 int32, p1 int32, p2 int32) int32 {
	return f678(c, p0, p1, p2)
}

func (c *Context) WrenAbortFiber(p0 int32, p1 int32) {
	f680(c, p0, p1)
}

func (c *Context) WrenGetListCount(p0 int32, p1 int32) int32 {
	return f668(c, p0, p1)
}

func (c *Context) WrenGetVersionNumber() int32 {
	return f641(c)
}

func (c *Context) WrenNewVM(p0 int32) int32 {
	return f645(c, p0)
}

func (c *Context) WrenFreeVM(p0 int32) {
	f646(c, p0)
}

func (c *Context) WrenGetSlotCount(p0 int32) int32 {
	return f650(c, p0)
}

func (c *Context) WrenGetSlotDouble(p0 int32, p1 int32) float64 {
	return f655(c, p0, p1)
}

func (c *Context) WrenGetSlotHandle(p0 int32, p1 int32) int32 {
	return f658(c, p0, p1)
}

func (c *Context) WrenSetSlotString(p0 int32, p1 int32, p2 int32) {
	f666(c, p0, p1, p2)
}

func (c *Context) WrenSetUserData(p0 int32, p1 int32) {
	f682(c, p0, p1)
}
