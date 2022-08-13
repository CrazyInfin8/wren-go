package internals

import (
	"errors"
	"math"
)

type Context struct {
	Mem []byte
	f   ImportedFuncs
	G0  int32
	G1  int32
	G2  int32
	G3  int32
}

func NewContext(f ImportedFuncs) *Context {
	c := &Context{Mem: newMemory(), f: f}
	c.G0 = 109792
	c.G1 = 22844
	c.G2 = 39368
	c.G3 = 21124
	return c
}
func (c *Context) Copy() *Context {
	d := *c
	d.Mem = make([]byte, len(c.Mem))
	copy(d.Mem, c.Mem)
	return &d
}

func (c *Context) growMem(size int) {
	c.Mem = append(c.Mem, make([]byte, size)...)
}

var (
	ErrDivByZero   = errors.New("div by zero")
	ErrIntOverflow = errors.New("int overflow")
)

func i32DivS(x, y int32) int32 {
	if y == 0 {
		panic(ErrDivByZero)
	}
	if x == math.MinInt32 && y == -1 {
		panic(ErrIntOverflow)
	}
	return x / y
}

func i64DivS(x, y int64) int64 {
	if y == 0 {
		panic(ErrDivByZero)
	}
	if x == math.MinInt64 && y == -1 {
		panic(ErrIntOverflow)
	}
	return x / y
}

func i32RemS(x, y int32) int32 {
	if y == 0 {
		panic(ErrDivByZero)
	}
	if x == math.MinInt32 && y == -1 {
		return 0
	}
	return x % y
}

func i64RemS(x, y int64) int64 {
	if y == 0 {
		panic(ErrDivByZero)
	}
	if x == math.MinInt64 && y == -1 {
		return 0
	}
	return x % y
}

func i32DivU(x, y int32) int32 {
	if y == 0 {
		panic(ErrDivByZero)
	}
	return int32(uint32(x) / uint32(y))
}

func i64DivU(x, y int64) int64 {
	if y == 0 {
		panic(ErrDivByZero)
	}
	return int64(uint64(x) / uint64(y))
}

func i32RemU(x, y int32) int32 {
	if y == 0 {
		panic(ErrDivByZero)
	}
	return int32(uint32(x) % uint32(y))
}

func i64RemU(x, y int64) int64 {
	if y == 0 {
		panic(ErrDivByZero)
	}
	return int64(uint64(x) % uint64(y))
}

type tableEntry struct {
	f         interface{}
	numParams int
}

var table [209]tableEntry

func init() {
	table[0] = tableEntry{f: nil}
	table[1] = tableEntry{f: f135, numParams: 2}
	table[2] = tableEntry{f: f137, numParams: 2}
	table[3] = tableEntry{f: f138, numParams: 2}
	table[4] = tableEntry{f: f139, numParams: 2}
	table[5] = tableEntry{f: f140, numParams: 2}
	table[6] = tableEntry{f: f141, numParams: 2}
	table[7] = tableEntry{f: f143, numParams: 2}
	table[8] = tableEntry{f: f144, numParams: 2}
	table[9] = tableEntry{f: f145, numParams: 2}
	table[10] = tableEntry{f: f146, numParams: 2}
	table[11] = tableEntry{f: f147, numParams: 2}
	table[12] = tableEntry{f: f150, numParams: 2}
	table[13] = tableEntry{f: f151, numParams: 2}
	table[14] = tableEntry{f: f152, numParams: 2}
	table[15] = tableEntry{f: f153, numParams: 2}
	table[16] = tableEntry{f: f154, numParams: 2}
	table[17] = tableEntry{f: f155, numParams: 2}
	table[18] = tableEntry{f: f156, numParams: 2}
	table[19] = tableEntry{f: f157, numParams: 2}
	table[20] = tableEntry{f: f158, numParams: 2}
	table[21] = tableEntry{f: f159, numParams: 2}
	table[22] = tableEntry{f: f160, numParams: 2}
	table[23] = tableEntry{f: f161, numParams: 2}
	table[24] = tableEntry{f: f162, numParams: 2}
	table[25] = tableEntry{f: f163, numParams: 2}
	table[26] = tableEntry{f: f164, numParams: 2}
	table[27] = tableEntry{f: f165, numParams: 2}
	table[28] = tableEntry{f: f166, numParams: 2}
	table[29] = tableEntry{f: f167, numParams: 2}
	table[30] = tableEntry{f: f168, numParams: 2}
	table[31] = tableEntry{f: f169, numParams: 2}
	table[32] = tableEntry{f: f170, numParams: 2}
	table[33] = tableEntry{f: f171, numParams: 2}
	table[34] = tableEntry{f: f172, numParams: 2}
	table[35] = tableEntry{f: f173, numParams: 2}
	table[36] = tableEntry{f: f174, numParams: 2}
	table[37] = tableEntry{f: f175, numParams: 2}
	table[38] = tableEntry{f: f176, numParams: 2}
	table[39] = tableEntry{f: f177, numParams: 2}
	table[40] = tableEntry{f: f178, numParams: 2}
	table[41] = tableEntry{f: f179, numParams: 2}
	table[42] = tableEntry{f: f180, numParams: 2}
	table[43] = tableEntry{f: f181, numParams: 2}
	table[44] = tableEntry{f: f182, numParams: 2}
	table[45] = tableEntry{f: f183, numParams: 2}
	table[46] = tableEntry{f: f184, numParams: 2}
	table[47] = tableEntry{f: f185, numParams: 2}
	table[48] = tableEntry{f: f186, numParams: 2}
	table[49] = tableEntry{f: f187, numParams: 2}
	table[50] = tableEntry{f: f188, numParams: 2}
	table[51] = tableEntry{f: f189, numParams: 2}
	table[52] = tableEntry{f: f190, numParams: 2}
	table[53] = tableEntry{f: f191, numParams: 2}
	table[54] = tableEntry{f: f192, numParams: 2}
	table[55] = tableEntry{f: f193, numParams: 2}
	table[56] = tableEntry{f: f194, numParams: 2}
	table[57] = tableEntry{f: f195, numParams: 2}
	table[58] = tableEntry{f: f196, numParams: 2}
	table[59] = tableEntry{f: f197, numParams: 2}
	table[60] = tableEntry{f: f198, numParams: 2}
	table[61] = tableEntry{f: f199, numParams: 2}
	table[62] = tableEntry{f: f200, numParams: 2}
	table[63] = tableEntry{f: f201, numParams: 2}
	table[64] = tableEntry{f: f202, numParams: 2}
	table[65] = tableEntry{f: f203, numParams: 2}
	table[66] = tableEntry{f: f204, numParams: 2}
	table[67] = tableEntry{f: f205, numParams: 2}
	table[68] = tableEntry{f: f206, numParams: 2}
	table[69] = tableEntry{f: f207, numParams: 2}
	table[70] = tableEntry{f: f208, numParams: 2}
	table[71] = tableEntry{f: f209, numParams: 2}
	table[72] = tableEntry{f: f210, numParams: 2}
	table[73] = tableEntry{f: f211, numParams: 2}
	table[74] = tableEntry{f: f212, numParams: 2}
	table[75] = tableEntry{f: f213, numParams: 2}
	table[76] = tableEntry{f: f214, numParams: 2}
	table[77] = tableEntry{f: f215, numParams: 2}
	table[78] = tableEntry{f: f216, numParams: 2}
	table[79] = tableEntry{f: f217, numParams: 2}
	table[80] = tableEntry{f: f218, numParams: 2}
	table[81] = tableEntry{f: f219, numParams: 2}
	table[82] = tableEntry{f: f220, numParams: 2}
	table[83] = tableEntry{f: f221, numParams: 2}
	table[84] = tableEntry{f: f222, numParams: 2}
	table[85] = tableEntry{f: f223, numParams: 2}
	table[86] = tableEntry{f: f224, numParams: 2}
	table[87] = tableEntry{f: f225, numParams: 2}
	table[88] = tableEntry{f: f226, numParams: 2}
	table[89] = tableEntry{f: f227, numParams: 2}
	table[90] = tableEntry{f: f228, numParams: 2}
	table[91] = tableEntry{f: f229, numParams: 2}
	table[92] = tableEntry{f: f230, numParams: 2}
	table[93] = tableEntry{f: f231, numParams: 2}
	table[94] = tableEntry{f: f232, numParams: 2}
	table[95] = tableEntry{f: f233, numParams: 2}
	table[96] = tableEntry{f: f234, numParams: 2}
	table[97] = tableEntry{f: f235, numParams: 2}
	table[98] = tableEntry{f: f236, numParams: 2}
	table[99] = tableEntry{f: f237, numParams: 2}
	table[100] = tableEntry{f: f238, numParams: 2}
	table[101] = tableEntry{f: f239, numParams: 2}
	table[102] = tableEntry{f: f240, numParams: 2}
	table[103] = tableEntry{f: f241, numParams: 2}
	table[104] = tableEntry{f: f242, numParams: 2}
	table[105] = tableEntry{f: f243, numParams: 2}
	table[106] = tableEntry{f: f244, numParams: 2}
	table[107] = tableEntry{f: f245, numParams: 2}
	table[108] = tableEntry{f: f246, numParams: 2}
	table[109] = tableEntry{f: f247, numParams: 2}
	table[110] = tableEntry{f: f248, numParams: 2}
	table[111] = tableEntry{f: f249, numParams: 2}
	table[112] = tableEntry{f: f250, numParams: 2}
	table[113] = tableEntry{f: f251, numParams: 2}
	table[114] = tableEntry{f: f252, numParams: 2}
	table[115] = tableEntry{f: f253, numParams: 2}
	table[116] = tableEntry{f: f254, numParams: 2}
	table[117] = tableEntry{f: f255, numParams: 2}
	table[118] = tableEntry{f: f256, numParams: 2}
	table[119] = tableEntry{f: f257, numParams: 2}
	table[120] = tableEntry{f: f258, numParams: 2}
	table[121] = tableEntry{f: f259, numParams: 2}
	table[122] = tableEntry{f: f260, numParams: 2}
	table[123] = tableEntry{f: f261, numParams: 2}
	table[124] = tableEntry{f: f262, numParams: 2}
	table[125] = tableEntry{f: f263, numParams: 2}
	table[126] = tableEntry{f: f264, numParams: 2}
	table[127] = tableEntry{f: f265, numParams: 2}
	table[128] = tableEntry{f: f266, numParams: 2}
	table[129] = tableEntry{f: f267, numParams: 2}
	table[130] = tableEntry{f: f268, numParams: 2}
	table[131] = tableEntry{f: f269, numParams: 2}
	table[132] = tableEntry{f: f270, numParams: 2}
	table[133] = tableEntry{f: f271, numParams: 2}
	table[134] = tableEntry{f: f272, numParams: 2}
	table[135] = tableEntry{f: f273, numParams: 2}
	table[136] = tableEntry{f: f274, numParams: 2}
	table[137] = tableEntry{f: f275, numParams: 2}
	table[138] = tableEntry{f: f276, numParams: 2}
	table[139] = tableEntry{f: f277, numParams: 2}
	table[140] = tableEntry{f: f278, numParams: 2}
	table[141] = tableEntry{f: f279, numParams: 2}
	table[142] = tableEntry{f: f280, numParams: 2}
	table[143] = tableEntry{f: f281, numParams: 2}
	table[144] = tableEntry{f: f282, numParams: 2}
	table[145] = tableEntry{f: f283, numParams: 2}
	table[146] = tableEntry{f: f284, numParams: 2}
	table[147] = tableEntry{f: f285, numParams: 2}
	table[148] = tableEntry{f: f286, numParams: 2}
	table[149] = tableEntry{f: f287, numParams: 2}
	table[150] = tableEntry{f: f288, numParams: 2}
	table[151] = tableEntry{f: f289, numParams: 2}
	table[152] = tableEntry{f: f290, numParams: 2}
	table[153] = tableEntry{f: f291, numParams: 2}
	table[154] = tableEntry{f: f292, numParams: 2}
	table[155] = tableEntry{f: f293, numParams: 2}
	table[156] = tableEntry{f: f294, numParams: 2}
	table[157] = tableEntry{f: f295, numParams: 2}
	table[158] = tableEntry{f: f296, numParams: 2}
	table[159] = tableEntry{f: f297, numParams: 2}
	table[160] = tableEntry{f: f409, numParams: 3}
	table[161] = tableEntry{f: f468, numParams: 2}
	table[162] = tableEntry{f: f469, numParams: 2}
	table[163] = tableEntry{f: f474, numParams: 2}
	table[164] = tableEntry{f: f476, numParams: 2}
	table[165] = tableEntry{f: f477, numParams: 2}
	table[166] = tableEntry{f: f478, numParams: 2}
	table[167] = tableEntry{f: f479, numParams: 2}
	table[168] = tableEntry{f: f480, numParams: 2}
	table[169] = tableEntry{f: f48, numParams: 2}
	table[170] = tableEntry{f: f50, numParams: 2}
	table[171] = tableEntry{f: f53, numParams: 2}
	table[172] = tableEntry{f: f39, numParams: 2}
	table[173] = tableEntry{f: f56, numParams: 2}
	table[174] = tableEntry{f: f57, numParams: 2}
	table[175] = tableEntry{f: f19, numParams: 2}
	table[176] = tableEntry{f: f31, numParams: 2}
	table[177] = tableEntry{f: f59, numParams: 2}
	table[178] = tableEntry{f: f36, numParams: 2}
	table[179] = tableEntry{f: f60, numParams: 2}
	table[180] = tableEntry{f: f63, numParams: 2}
	table[181] = tableEntry{f: f35, numParams: 2}
	table[182] = tableEntry{f: f64, numParams: 2}
	table[183] = tableEntry{f: f45, numParams: 2}
	table[184] = tableEntry{f: f65, numParams: 2}
	table[185] = tableEntry{f: f67, numParams: 2}
	table[186] = tableEntry{f: f68, numParams: 2}
	table[187] = tableEntry{f: f72, numParams: 2}
	table[188] = tableEntry{f: f73, numParams: 2}
	table[189] = tableEntry{f: f76, numParams: 2}
	table[190] = tableEntry{f: f82, numParams: 2}
	table[191] = tableEntry{f: f43, numParams: 2}
	table[192] = tableEntry{f: f87, numParams: 2}
	table[193] = tableEntry{f: f89, numParams: 2}
	table[194] = tableEntry{f: f0, numParams: 3}
	table[195] = tableEntry{f: f528, numParams: 3}
	table[196] = tableEntry{f: f529, numParams: 6}
	table[197] = tableEntry{f: f530, numParams: 4}
	table[198] = tableEntry{f: f1, numParams: 2}
	table[199] = tableEntry{f: f2, numParams: 5}
	table[200] = tableEntry{f: f531, numParams: 3}
	table[201] = tableEntry{f: f5, numParams: 2}
	table[202] = tableEntry{f: f7, numParams: 3}
	table[203] = tableEntry{f: f559, numParams: 3}
	table[204] = tableEntry{f: f557, numParams: 1}
	table[205] = tableEntry{f: f561, numParams: 3}
	table[206] = tableEntry{f: f563, numParams: 3}
	table[207] = tableEntry{f: f577, numParams: 3}
	table[208] = tableEntry{f: f637, numParams: 3}
}
