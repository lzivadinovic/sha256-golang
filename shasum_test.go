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

func TestRoundDeclare(t *testing.T) {
	W := roundDeclare(roundTests.state, roundTests.roundConst, roundTests.scheduleWord)
	for index, value := range W {
		if roundTests.expect[index] != value {
			t.Errorf("Got: %d, expected: %d", W, roundTests.expect)
		}
	}
}

func TestRoundMutate(t *testing.T) {
	// CAREFULL! It will mutate roundTests variable!!!
	roundMutate(roundTests.state, roundTests.roundConst, roundTests.scheduleWord)
	for index, value := range roundTests.state {
		if roundTests.expect[index] != value {
			t.Errorf("Got: %d, expected: %d", roundTests.state, roundTests.expect)
		}
	}
}
