package main

import (
	"testing"
)

type addTest struct {
	arg1, arg2, expected uint32
}

var addTests = []addTest{
	addTest{1, 2, 3},
	addTest{4294967295, 1, 0},
	addTest{4294967295, 10, 9},
	addTest{3050487260, 3710144918, 2465664882},
}

type rotateTest struct {
	n, k, expect uint32
}

var rotateTests = []rotateTest{
	rotateTest{2, 1, 1},
	rotateTest{1, 1, 2147483648},
	rotateTest{2919882184, 31, 1544797073},
	rotateTest{2919882184, 64, 2919882184},
}

func TestAddUInt(t *testing.T) {
	for _, test := range addTests {
		output := addUInt(test.arg1, test.arg2)
		if output != test.expected {
			t.Errorf("Got: %d, expected: %d", output, test.expected)
		}
	}
}

func TestAddDirect(t *testing.T) {
	for _, test := range addTests {
		output := addDirect(test.arg1, test.arg2)
		if output != test.expected {
			t.Errorf("Got: %d, expected: %d", output, test.expected)
		}
	}
}

func TestRightRotate(t *testing.T) {
	for _, test := range rotateTests {
		output := rightRotate(test.n, test.k)
		if output != test.expect {
			t.Errorf("Got: %d, expected: %d", output, test.expect)
		}
	}
}
func TestGoRightRotate(t *testing.T) {
	for _, test := range rotateTests {
		output := goRightRotate(test.n, test.k)
		if output != test.expect {
			t.Errorf("Got: %d, expected: %d", output, test.expect)
		}
	}
}

type sigmaTest struct {
	x, expect uint32
}

var littleSigma0Tests = sigmaTest{1114723206, 1345017931}
var littleSigma1Tests = sigmaTest{1232674167, 2902922196}
var bigSigma0Tests = sigmaTest{1114723206, 3386924642}
var bigSigma1Tests = sigmaTest{1114723206, 1365788741}

func TestLittleSigma0(t *testing.T) {
	output := littleSigma0(littleSigma0Tests.x)
	if output != littleSigma0Tests.expect {
		t.Errorf("Got: %d, expected: %d", output, littleSigma0Tests.expect)
	}
}

func TestLittleSigma1(t *testing.T) {
	output := littleSigma1(littleSigma1Tests.x)
	if output != littleSigma1Tests.expect {
		t.Errorf("Got: %d, expected: %d", output, littleSigma1Tests.expect)
	}
}

func TestBigSigma0(t *testing.T) {
	output := bigSigma0(bigSigma0Tests.x)
	if output != bigSigma0Tests.expect {
		t.Errorf("Got: %d, expected: %d", output, bigSigma0Tests.expect)
	}
}
func TestBigSigma1(t *testing.T) {
	output := bigSigma1(bigSigma1Tests.x)
	if output != bigSigma1Tests.expect {
		t.Errorf("Got: %d, expected: %d", output, bigSigma1Tests.expect)
	}
}

type messageScheduleTest struct {
	x      []byte
	expect []uint32
}

var messageScheduleTests = messageScheduleTest{[]byte("iguana wombat dog kangaroo llama turkey yak unicorn sheep xenoce"),
	[]uint32{1768387937, 1851859063, 1869439585, 1948279919, 1730177889, 1852268914, 1869553772, 1818324321,
		544503154, 1801812256, 2036427552, 1970170211, 1869770272, 1936221541, 1881176165, 1852793701,
		3002878561, 3711121932, 1520676164, 3002441970, 2935068969, 1610329529, 1904580351, 3219988740,
		2337695268, 263015313, 2120931855, 131203777, 3818546915, 19163115, 3479924161, 2154860703,
		1790169326, 516580487, 2414737634, 909025701, 2241053595, 1237268359, 3797503938, 1773623028,
		2840671725, 2299292186, 1933596460, 2279513616, 514132674, 3245155609, 1753922983, 2241450350,
		2449659630, 262239956, 773552098, 3253131632, 3863807927, 879696536, 3143654396, 3973063648,
		509015903, 270850193, 1893431553, 719566283, 2310657204, 365781698, 3761063438, 1007484868,
	}}

func TestMessageSchedule(t *testing.T) {
	output := messageSchedule(messageScheduleTests.x)
	for index, value := range messageScheduleTests.expect {
		if output[index] != value {
			t.Errorf("Got: %d, expected: %d, on expect array index %d", output[index], value, index)
		}
	}
}

type choiceTest struct {
	x, y, z, expect uint32
}

var choiceTests = choiceTest{2749825547, 776049372, 1213590135, 1783753340}
var majorityTests = choiceTest{3758166654, 2821345890, 1850678816, 3893039714}

func TestChoice(t *testing.T) {
	output := choice(choiceTests.x, choiceTests.y, choiceTests.z)
	if output != choiceTests.expect {
		t.Errorf("Got: %d, expected: %d", output, choiceTests.expect)
	}
}

func TestMajority(t *testing.T) {
	output := majority(majorityTests.x, majorityTests.y, majorityTests.z)
	if output != majorityTests.expect {
		t.Errorf("Got: %d, expected: %d", output, majorityTests.expect)
	}
}

type roundTest struct {
	state, expect            []uint32
	roundConst, scheduleWord uint32
}

var roundTests = roundTest{
	[]uint32{2739944672, 3126690193, 4191866847, 1163785745, 3714074692, 1172792371, 283469062, 826169706},
	[]uint32{1724514418, 2739944672, 3126690193, 4191866847, 1638715774, 3714074692, 1172792371, 283469062},
	961987163,
	3221900128,
}

// we are using mutate in prod code!
// func TestRoundDeclare(t *testing.T) {
// 	W := roundDeclare(roundTests.state, roundTests.roundConst, roundTests.scheduleWord)
// 	for index, value := range W {
// 		if roundTests.expect[index] != value {
// 			t.Errorf("Got: %d, expected: %d", W, roundTests.expect)
// 		}
// 	}
// }

func TestRoundMutate(t *testing.T) {
	// CAREFULL! It will mutate roundTests variable!!!
	roundMutate(roundTests.state, roundTests.roundConst, roundTests.scheduleWord)
	for index, value := range roundTests.state {
		if roundTests.expect[index] != value {
			t.Errorf("Got: %d, expected: %d", value, roundTests.expect)
		}
	}
}

type compressBlockTest struct {
	state  []uint32
	block  []byte
	expect []uint32
}

var compressBlockTests = compressBlockTest{
	[]uint32{2918946378, 1679978889, 1678006433, 650957219, 379281712, 2112907926, 1775216060, 2152648190},
	[]byte("manatee fox unicorn octopus dog fox fox llama vulture jaguar xen"),
	[]uint32{1251501988, 1663226031, 2877128394, 4050467288, 2375501075, 1434687977, 2625842981, 650253644},
}

//var K = []uint32{0x428a2f98, 0x71374491, 0xb5c0fbcf, 0xe9b5dba5, 0x3956c25b, 0x59f111f1, 0x923f82a4, 0xab1c5ed5,
//	0xd807aa98, 0x12835b01, 0x243185be, 0x550c7dc3, 0x72be5d74, 0x80deb1fe, 0x9bdc06a7, 0xc19bf174,
//	0xe49b69c1, 0xefbe4786, 0x0fc19dc6, 0x240ca1cc, 0x2de92c6f, 0x4a7484aa, 0x5cb0a9dc, 0x76f988da,
//	0x983e5152, 0xa831c66d, 0xb00327c8, 0xbf597fc7, 0xc6e00bf3, 0xd5a79147, 0x06ca6351, 0x14292967,
//	0x27b70a85, 0x2e1b2138, 0x4d2c6dfc, 0x53380d13, 0x650a7354, 0x766a0abb, 0x81c2c92e, 0x92722c85,
//	0xa2bfe8a1, 0xa81a664b, 0xc24b8b70, 0xc76c51a3, 0xd192e819, 0xd6990624, 0xf40e3585, 0x106aa070,
//	0x19a4c116, 0x1e376c08, 0x2748774c, 0x34b0bcb5, 0x391c0cb3, 0x4ed8aa4a, 0x5b9cca4f, 0x682e6ff3,
//	0x748f82ee, 0x78a5636f, 0x84c87814, 0x8cc70208, 0x90befffa, 0xa4506ceb, 0xbef9a3f7, 0xc67178f2}

func TestCompressBlock(t *testing.T) {
	newState := compressBlock(compressBlockTests.state, compressBlockTests.block)
	for index, value := range newState {
		if compressBlockTests.expect[index] != value {
			t.Errorf("Got: %d, expected: %d", value, compressBlockTests.expect[index])
		}
	}
}
