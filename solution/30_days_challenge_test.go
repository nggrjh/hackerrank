package solution

import (
	"reflect"
	"testing"
)

func Test_solveOperators(t *testing.T) {
	t.Parallel()
	type args struct {
		mealCost   float64
		tipPercent int32
		taxPercent int32
	}
	tests := map[string]struct {
		args args
		want int32
	}{
		"case_0": {
			args: args{
				mealCost:   12,
				tipPercent: 20,
				taxPercent: 8,
			},
			want: 15,
		},
		"case_1": {
			args: args{
				mealCost:   15.5,
				tipPercent: 15,
				taxPercent: 10,
			},
			want: 19,
		},
		"case_2": {
			args: args{
				mealCost:   10.25,
				tipPercent: 17,
				taxPercent: 5,
			},
			want: 13,
		},
	}
	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			if got := solveOperators(tt.args.mealCost, tt.args.tipPercent, tt.args.taxPercent); got != tt.want {
				t.Errorf("solveOperators() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxHourGlass(t *testing.T) {
	t.Parallel()
	type args struct {
		arr [][]int32
	}
	tests := map[string]struct {
		args args
		want int32
	}{
		"case_0": {
			args: args{
				arr: [][]int32{
					{1, 1, 1, 0, 0, 0},
					{0, 1, 0, 0, 0, 0},
					{1, 1, 1, 0, 0, 0},
					{0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0},
				},
			},
			want: 7,
		},
		"case_1": {
			args: args{
				arr: [][]int32{
					{1, 1, 1, 0, 0, 0},
					{0, 1, 0, 0, 0, 0},
					{1, 1, 1, 0, 0, 0},
					{0, 0, 2, 4, 4, 0},
					{0, 0, 0, 2, 0, 0},
					{0, 0, 1, 2, 4, 0},
				},
			},
			want: 19,
		},
		"case_3": {
			args: args{
				arr: [][]int32{
					{0, -4, -6, 0, -7, -6},
					{-1, -2, -6, -8, -3, -1},
					{-8, -4, -2, -8, -8, -6},
					{-3, -1, -2, -5, -7, -4},
					{-3, -5, -3, -6, -6, -6},
					{-3, -6, 0, -8, -6, -7},
				},
			},
			want: -19,
		},
	}
	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := maxHourGlass(tt.args.arr); got != tt.want {
				t.Errorf("maxHourGlass() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countSwaps(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := map[string]struct {
		args  args
		want  []int32
		want1 int32
	}{
		"case": {
			args: args{
				arr: []int32{4, 3, 1, 2},
			},
			want:  []int32{1, 2, 3, 4},
			want1: 5,
		},
		"case1": {
			args: args{
				arr: []int32{1, 2, 3, 4},
			},
			want:  []int32{1, 2, 3, 4},
			want1: 0,
		},
	}
	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			got, got1 := countSwaps(tt.args.arr)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countSwaps() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("countSwaps() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_calculateFine(t *testing.T) {
	t.Parallel()
	type args struct {
		returnedDate string
		dueDate      string
	}
	tests := map[string]struct {
		args args
		want int32
	}{
		"case_0.0": {
			args: args{
				returnedDate: "31 12 2014",
				dueDate:      "1 1 2015",
			},
			want: 0,
		},
		"case_0.1": {
			args: args{
				returnedDate: "1 1 2015",
				dueDate:      "31 12 2014",
			},
			want: 10000,
		},
		"case_2": {
			args: args{
				returnedDate: "9 6 2015",
				dueDate:      "6 6 2015",
			},
			want: 45,
		},
		"case_6": {
			args: args{
				returnedDate: "8 4 12",
				dueDate:      "1 1 1",
			},
			want: 10000,
		},
	}
	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := calculateFine(tt.args.returnedDate, tt.args.dueDate); got != tt.want {
				t.Errorf("calculateFine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkPrime(t *testing.T) {
	t.Parallel()
	type args struct {
		arr []int64
	}
	tests := map[string]struct {
		args args
		want []string
	}{
		"case_0": {
			args: args{
				arr: []int64{12, 5, 7},
			},
			want: []string{"Not prime", "Prime", "Prime"},
		},
		"case_1": {
			args: args{
				arr: []int64{31, 33},
			},
			want: []string{"Prime", "Not prime"},
		},
	}
	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := checkPrime(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checkPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bitwiseAnd(t *testing.T) {
	t.Parallel()
	type args struct {
		N int32
		K int32
	}
	tests := map[string]struct {
		args args
		want int32
	}{
		"case_0.0": {
			args: args{
				N: 5,
				K: 2,
			},
			want: 1,
		},
		"case_0.2": {
			args: args{
				N: 8,
				K: 5,
			},
			want: 4,
		},
		"case_0.3": {
			args: args{
				N: 2,
				K: 2,
			},
			want: 0,
		},
		"case_1.1": {
			args: args{
				N: 9,
				K: 2,
			},
			want: 1,
		},
		"case_1.2": {
			args: args{
				N: 8,
				K: 3,
			},
			want: 2,
		},
		"case_4.1": {
			args: args{
				N: 519,
				K: 177,
			},
			want: 176,
		},
		"case_4.2": {
			args: args{
				N: 410,
				K: 388,
			},
			want: 387,
		},
		"case_4.3": {
			args: args{
				N: 234,
				K: 172,
			},
			want: 171,
		},
		"case_4.4": {
			args: args{
				N: 449,
				K: 8,
			},
			want: 7,
		},
		"case_4.5": {
			args: args{
				N: 179,
				K: 157,
			},
			want: 156,
		},
		"case_4.6": {
			args: args{
				N: 582,
				K: 257,
			},
			want: 256,
		},
		"case_4.7": {
			args: args{
				N: 913,
				K: 363,
			},
			want: 362,
		},
		"case_4.8": {
			args: args{
				N: 116,
				K: 112,
			},
			want: 110,
		},
		"case_4.9": {
			args: args{
				N: 116,
				K: 50,
			},
			want: 49,
		},
		"case_4.10": {
			args: args{
				N: 276,
				K: 250,
			},
			want: 249,
		},
		"case_4.11": {
			args: args{
				N: 381,
				K: 199,
			},
			want: 198,
		},
		"case_4.12": {
			args: args{
				N: 655,
				K: 225,
			},
			want: 224,
		},
		"case_4.13": {
			args: args{
				N: 225,
				K: 5,
			},
			want: 4,
		},
		"case_4.14": {
			args: args{
				N: 424,
				K: 141,
			},
			want: 140,
		},
		"case_4.15": {
			args: args{
				N: 247,
				K: 97,
			},
			want: 96,
		},
		"case_4.16": {
			args: args{
				N: 118,
				K: 24,
			},
			want: 23,
		},
		"case_4.17": {
			args: args{
				N: 156,
				K: 52,
			},
			want: 51,
		},
		"case_4.18": {
			args: args{
				N: 614,
				K: 327,
			},
			want: 326,
		},
		"case_4.19": {
			args: args{
				N: 614,
				K: 327,
			},
			want: 326,
		},
		"case_4.20": {
			args: args{
				N: 160,
				K: 38,
			},
			want: 37,
		},
		"case_4.21": {
			args: args{
				N: 884,
				K: 855,
			},
			want: 854,
		},
		"case_4.22": {
			args: args{
				N: 336,
				K: 201,
			},
			want: 200,
		},
		"case_4.23": {
			args: args{
				N: 127,
				K: 76,
			},
			want: 75,
		},
		"case_4.24": {
			args: args{
				N: 785,
				K: 635,
			},
			want: 634,
		},
		"case_4.25": {
			args: args{
				N: 529,
				K: 52,
			},
			want: 51,
		},
		"case_4.26": {
			args: args{
				N: 417,
				K: 27,
			},
			want: 26,
		},
		"case_4.27": {
			args: args{
				N: 934,
				K: 502,
			},
			want: 501,
		},
		"case_4.28": {
			args: args{
				N: 157,
				K: 128,
			},
			want: 126,
		},
		"case_4.29": {
			args: args{
				N: 670,
				K: 172,
			},
			want: 171,
		},
		"case_4.30": {
			args: args{
				N: 412,
				K: 108,
			},
			want: 107,
		},
		"case_4.31": {
			args: args{
				N: 239,
				K: 173,
			},
			want: 172,
		},
		"case_4.32": {
			args: args{
				N: 457,
				K: 100,
			},
			want: 99,
		},
		"case_4.33": {
			args: args{
				N: 879,
				K: 588,
			},
			want: 587,
		},
		"case_4.34": {
			args: args{
				N: 117,
				K: 90,
			},
			want: 89,
		},
		"case_4.35": {
			args: args{
				N: 974,
				K: 331,
			},
			want: 330,
		},
		"case_4.36": {
			args: args{
				N: 195,
				K: 62,
			},
			want: 61,
		},
		"case_4.37": {
			args: args{
				N: 188,
				K: 125,
			},
			want: 124,
		},
		"case_4.38": {
			args: args{
				N: 773,
				K: 510,
			},
			want: 509,
		},
		"case_4.39": {
			args: args{
				N: 843,
				K: 531,
			},
			want: 530,
		},
		"case_4.40": {
			args: args{
				N: 661,
				K: 204,
			},
			want: 203,
		},
		"case_4.41": {
			args: args{
				N: 169,
				K: 14,
			},
			want: 13,
		},
		"case_4.42": {
			args: args{
				N: 253,
				K: 78,
			},
			want: 77,
		},
		"case_4.43": {
			args: args{
				N: 990,
				K: 39,
			},
			want: 38,
		},
		"case_4.44": {
			args: args{
				N: 731,
				K: 715,
			},
			want: 714,
		},
		"case_4.45": {
			args: args{
				N: 523,
				K: 328,
			},
			want: 327,
		},
		"case_4.46": {
			args: args{
				N: 260,
				K: 53,
			},
			want: 52,
		},
		"case_4.47": {
			args: args{
				N: 979,
				K: 401,
			},
			want: 400,
		},
		"case_4.48": {
			args: args{
				N: 393,
				K: 108,
			},
			want: 107,
		},
		"case_4.49": {
			args: args{
				N: 266,
				K: 161,
			},
			want: 160,
		},
		"case_4.50": {
			args: args{
				N: 348,
				K: 2,
			},
			want: 1,
		},
		"case_4.51": {
			args: args{
				N: 824,
				K: 135,
			},
			want: 134,
		},
		"case_4.52": {
			args: args{
				N: 466,
				K: 325,
			},
			want: 324,
		},
		"case_4.53": {
			args: args{
				N: 117,
				K: 9,
			},
			want: 8,
		},
		"case_4.54": {
			args: args{
				N: 385,
				K: 47,
			},
			want: 46,
		},
		"case_4.55": {
			args: args{
				N: 285,
				K: 41,
			},
			want: 40,
		},
		"case_4.56": {
			args: args{
				N: 541,
				K: 106,
			},
			want: 105,
		},
		"case_4.57": {
			args: args{
				N: 175,
				K: 125,
			},
			want: 124,
		},
		"case_4.58": {
			args: args{
				N: 223,
				K: 48,
			},
			want: 47,
		},
		"case_4.59": {
			args: args{
				N: 877,
				K: 853,
			},
			want: 852,
		},
		"case_4.60": {
			args: args{
				N: 711,
				K: 274,
			},
			want: 273,
		},
		"case_4.61": {
			args: args{
				N: 434,
				K: 189,
			},
			want: 188,
		},
		"case_4.62": {
			args: args{
				N: 157,
				K: 42,
			},
			want: 41,
		},
		"case_4.63": {
			args: args{
				N: 785,
				K: 305,
			},
			want: 304,
		},
		"case_4.64": {
			args: args{
				N: 515,
				K: 61,
			},
			want: 60,
		},
		"case_4.65": {
			args: args{
				N: 260,
				K: 220,
			},
			want: 219,
		},
		"case_4.66": {
			args: args{
				N: 961,
				K: 406,
			},
			want: 405,
		},
		"case_4.67": {
			args: args{
				N: 102,
				K: 59,
			},
			want: 58,
		},
		"case_4.68": {
			args: args{
				N: 774,
				K: 154,
			},
			want: 153,
		},
		"case_4.69": {
			args: args{
				N: 555,
				K: 374,
			},
			want: 373,
		},
		"case_4.70": {
			args: args{
				N: 103,
				K: 88,
			},
			want: 87,
		},
		"case_4.71": {
			args: args{
				N: 915,
				K: 722,
			},
			want: 721,
		},
		"case_4.72": {
			args: args{
				N: 415,
				K: 40,
			},
			want: 39,
		},
		"case_4.73": {
			args: args{
				N: 117,
				K: 32,
			},
			want: 31,
		},
		"case_4.74": {
			args: args{
				N: 146,
				K: 123,
			},
			want: 122,
		},
		"case_4.75": {
			args: args{
				N: 338,
				K: 49,
			},
			want: 48,
		},
		"case_4.76": {
			args: args{
				N: 871,
				K: 108,
			},
			want: 107,
		},
		"case_4.77": {
			args: args{
				N: 167,
				K: 127,
			},
			want: 126,
		},
		"case_4.78": {
			args: args{
				N: 140,
				K: 106,
			},
			want: 105,
		},
		"case_4.79": {
			args: args{
				N: 226,
				K: 109,
			},
			want: 108,
		},
		"case_4.80": {
			args: args{
				N: 693,
				K: 637,
			},
			want: 636,
		},
		"case_4.81": {
			args: args{
				N: 519,
				K: 286,
			},
			want: 285,
		},
		"case_4.82": {
			args: args{
				N: 431,
				K: 300,
			},
			want: 299,
		},
		"case_4.83": {
			args: args{
				N: 174,
				K: 49,
			},
			want: 48,
		},
		"case_4.84": {
			args: args{
				N: 101,
				K: 41,
			},
			want: 40,
		},
		"case_4.85": {
			args: args{
				N: 306,
				K: 23,
			},
			want: 22,
		},
		"case_4.86": {
			args: args{
				N: 397,
				K: 126,
			},
			want: 125,
		},
		"case_4.87": {
			args: args{
				N: 144,
				K: 17,
			},
			want: 16,
		},
		"case_4.88": {
			args: args{
				N: 547,
				K: 334,
			},
			want: 333,
		},
		"case_4.89": {
			args: args{
				N: 768,
				K: 204,
			},
			want: 203,
		},
		"case_4.90": {
			args: args{
				N: 983,
				K: 753,
			},
			want: 752,
		},
		"case_4.91": {
			args: args{
				N: 520,
				K: 87,
			},
			want: 86,
		},
		"case_4.92": {
			args: args{
				N: 281,
				K: 237,
			},
			want: 236,
		},
		"case_4.93": {
			args: args{
				N: 671,
				K: 269,
			},
			want: 268,
		},
		"case_4.94": {
			args: args{
				N: 544,
				K: 129,
			},
			want: 128,
		},
		"case_4.95": {
			args: args{
				N: 306,
				K: 179,
			},
			want: 178,
		},
		"case_4.96": {
			args: args{
				N: 757,
				K: 724,
			},
			want: 723,
		},
		"case_4.97": {
			args: args{
				N: 493,
				K: 285,
			},
			want: 284,
		},
		"case_4.98": {
			args: args{
				N: 700,
				K: 596,
			},
			want: 595,
		},
		"case_4.99": {
			args: args{
				N: 189,
				K: 103,
			},
			want: 102,
		},
		"case_4.100": {
			args: args{
				N: 291,
				K: 58,
			},
			want: 57,
		},
		"case_4.101": {
			args: args{
				N: 331,
				K: 284,
			},
			want: 283,
		},
		"case_4.102": {
			args: args{
				N: 347,
				K: 304,
			},
			want: 303,
		},
		"case_4.103": {
			args: args{
				N: 992,
				K: 856,
			},
			want: 855,
		},
		"case_4.104": {
			args: args{
				N: 944,
				K: 487,
			},
			want: 486,
		},
		"case_4.105": {
			args: args{
				N: 725,
				K: 325,
			},
			want: 324,
		},
		"case_4.106": {
			args: args{
				N: 163,
				K: 140,
			},
			want: 139,
		},
		"case_4.107": {
			args: args{
				N: 288,
				K: 216,
			},
			want: 215,
		},
		"case_4.108": {
			args: args{
				N: 885,
				K: 57,
			},
			want: 56,
		},
		"case_4.109": {
			args: args{
				N: 913,
				K: 42,
			},
			want: 41,
		},
		"case_4.110": {
			args: args{
				N: 401,
				K: 154,
			},
			want: 153,
		},
		"case_4.111": {
			args: args{
				N: 160,
				K: 18,
			},
			want: 17,
		},
		"case_4.112": {
			args: args{
				N: 575,
				K: 193,
			},
			want: 192,
		},
		"case_4.113": {
			args: args{
				N: 772,
				K: 595,
			},
			want: 594,
		},
		"case_4.114": {
			args: args{
				N: 123,
				K: 100,
			},
			want: 99,
		},
		"case_4.115": {
			args: args{
				N: 562,
				K: 132,
			},
			want: 131,
		},
		"case_4.116": {
			args: args{
				N: 523,
				K: 69,
			},
			want: 68,
		},
		"case_4.117": {
			args: args{
				N: 381,
				K: 68,
			},
			want: 67,
		},
		"case_4.118": {
			args: args{
				N: 512,
				K: 109,
			},
			want: 108,
		},
		"case_4.119": {
			args: args{
				N: 440,
				K: 209,
			},
			want: 208,
		},
		"case_4.120": {
			args: args{
				N: 479,
				K: 423,
			},
			want: 422,
		},
		"case_4.121": {
			args: args{
				N: 818,
				K: 572,
			},
			want: 571,
		},
		"case_4.122": {
			args: args{
				N: 917,
				K: 26,
			},
			want: 25,
		},
		"case_4.123": {
			args: args{
				N: 681,
				K: 435,
			},
			want: 434,
		},
		"case_4.124": {
			args: args{
				N: 645,
				K: 419,
			},
			want: 418,
		},
		"case_4.125": {
			args: args{
				N: 312,
				K: 241,
			},
			want: 240,
		},
		"case_4.126": {
			args: args{
				N: 526,
				K: 446,
			},
			want: 445,
		},
		"case_4.127": {
			args: args{
				N: 201,
				K: 155,
			},
			want: 154,
		},
		"case_4.128": {
			args: args{
				N: 738,
				K: 440,
			},
			want: 439,
		},
		"case_4.129": {
			args: args{
				N: 127,
				K: 64,
			},
			want: 63,
		},
		"case_4.130": {
			args: args{
				N: 794,
				K: 395,
			},
			want: 394,
		},
		"case_4.131": {
			args: args{
				N: 438,
				K: 373,
			},
			want: 372,
		},
		"case_4.132": {
			args: args{
				N: 761,
				K: 278,
			},
			want: 277,
		},
		"case_4.133": {
			args: args{
				N: 139,
				K: 82,
			},
			want: 81,
		},
		"case_4.134": {
			args: args{
				N: 568,
				K: 527,
			},
			want: 526,
		},
		"case_4.135": {
			args: args{
				N: 735,
				K: 454,
			},
			want: 453,
		},
		"case_4.136": {
			args: args{
				N: 527,
				K: 428,
			},
			want: 427,
		},
		"case_4.137": {
			args: args{
				N: 318,
				K: 311,
			},
			want: 310,
		},
		"case_4.138": {
			args: args{
				N: 739,
				K: 377,
			},
			want: 376,
		},
		"case_4.139": {
			args: args{
				N: 179,
				K: 98,
			},
			want: 97,
		},
		"case_4.140": {
			args: args{
				N: 951,
				K: 88,
			},
			want: 87,
		},
		"case_4.141": {
			args: args{
				N: 451,
				K: 182,
			},
			want: 181,
		},
		"case_4.142": {
			args: args{
				N: 858,
				K: 154,
			},
			want: 153,
		},
		"case_4.143": {
			args: args{
				N: 124,
				K: 77,
			},
			want: 76,
		},
		"case_4.144": {
			args: args{
				N: 180,
				K: 107,
			},
			want: 106,
		},
		"case_4.145": {
			args: args{
				N: 472,
				K: 458,
			},
			want: 457,
		},
		"case_4.146": {
			args: args{
				N: 899,
				K: 254,
			},
			want: 253,
		},
		"case_4.147": {
			args: args{
				N: 527,
				K: 476,
			},
			want: 475,
		},
		"case_4.148": {
			args: args{
				N: 553,
				K: 335,
			},
			want: 334,
		},
		"case_4.149": {
			args: args{
				N: 149,
				K: 78,
			},
			want: 77,
		},
		"case_4.150": {
			args: args{
				N: 982,
				K: 666,
			},
			want: 665,
		},
		"case_4.151": {
			args: args{
				N: 359,
				K: 249,
			},
			want: 248,
		},
		"case_4.152": {
			args: args{
				N: 191,
				K: 13,
			},
			want: 12,
		},
		"case_4.153": {
			args: args{
				N: 667,
				K: 237,
			},
			want: 236,
		},
		"case_4.154": {
			args: args{
				N: 449,
				K: 439,
			},
			want: 438,
		},
		"case_4.155": {
			args: args{
				N: 139,
				K: 33,
			},
			want: 32,
		},
		"case_4.156": {
			args: args{
				N: 528,
				K: 183,
			},
			want: 182,
		},
		"case_4.157": {
			args: args{
				N: 827,
				K: 107,
			},
			want: 106,
		},
		"case_4.158": {
			args: args{
				N: 112,
				K: 91,
			},
			want: 90,
		},
		"case_4.159": {
			args: args{
				N: 298,
				K: 213,
			},
			want: 212,
		},
		"case_4.160": {
			args: args{
				N: 278,
				K: 262,
			},
			want: 261,
		},
		"case_4.161": {
			args: args{
				N: 729,
				K: 96,
			},
			want: 95,
		},
		"case_4.162": {
			args: args{
				N: 271,
				K: 75,
			},
			want: 74,
		},
		"case_4.163": {
			args: args{
				N: 894,
				K: 540,
			},
			want: 539,
		},
		"case_4.164": {
			args: args{
				N: 683,
				K: 664,
			},
			want: 663,
		},
		"case_4.165": {
			args: args{
				N: 824,
				K: 417,
			},
			want: 416,
		},
		"case_4.166": {
			args: args{
				N: 231,
				K: 98,
			},
			want: 97,
		},
		"case_4.167": {
			args: args{
				N: 206,
				K: 25,
			},
			want: 24,
		},
		"case_4.168": {
			args: args{
				N: 178,
				K: 89,
			},
			want: 88,
		},
		"case_4.169": {
			args: args{
				N: 605,
				K: 112,
			},
			want: 111,
		},
		"case_4.170": {
			args: args{
				N: 486,
				K: 121,
			},
			want: 120,
		},
		"case_4.171": {
			args: args{
				N: 475,
				K: 192,
			},
			want: 191,
		},
		"case_4.172": {
			args: args{
				N: 753,
				K: 89,
			},
			want: 88,
		},
		"case_4.173": {
			args: args{
				N: 898,
				K: 820,
			},
			want: 819,
		},
		"case_4.174": {
			args: args{
				N: 450,
				K: 218,
			},
			want: 217,
		},
		"case_4.175": {
			args: args{
				N: 711,
				K: 371,
			},
			want: 370,
		},
		"case_4.176": {
			args: args{
				N: 684,
				K: 251,
			},
			want: 250,
		},
		"case_4.177": {
			args: args{
				N: 134,
				K: 5,
			},
			want: 4,
		},
		"case_4.178": {
			args: args{
				N: 674,
				K: 127,
			},
			want: 126,
		},
		"case_4.179": {
			args: args{
				N: 164,
				K: 123,
			},
			want: 122,
		},
		"case_4.180": {
			args: args{
				N: 754,
				K: 108,
			},
			want: 107,
		},
		"case_4.181": {
			args: args{
				N: 410,
				K: 398,
			},
			want: 397,
		},
		"case_4.182": {
			args: args{
				N: 163,
				K: 53,
			},
			want: 52,
		},
		"case_4.183": {
			args: args{
				N: 631,
				K: 33,
			},
			want: 32,
		},
		"case_4.184": {
			args: args{
				N: 440,
				K: 170,
			},
			want: 169,
		},
		"case_4.185": {
			args: args{
				N: 658,
				K: 600,
			},
			want: 599,
		},
		"case_4.186": {
			args: args{
				N: 184,
				K: 3,
			},
			want: 2,
		},
		"case_4.187": {
			args: args{
				N: 479,
				K: 357,
			},
			want: 356,
		},
		"case_4.188": {
			args: args{
				N: 346,
				K: 287,
			},
			want: 286,
		},
		"case_4.189": {
			args: args{
				N: 559,
				K: 261,
			},
			want: 260,
		},
		"case_4.190": {
			args: args{
				N: 576,
				K: 214,
			},
			want: 213,
		},
		"case_4.191": {
			args: args{
				N: 214,
				K: 55,
			},
			want: 54,
		},
		"case_4.192": {
			args: args{
				N: 952,
				K: 426,
			},
			want: 425,
		},
		"case_4.193": {
			args: args{
				N: 844,
				K: 288,
			},
			want: 287,
		},
		"case_4.194": {
			args: args{
				N: 938,
				K: 366,
			},
			want: 365,
		},
		"case_4.195": {
			args: args{
				N: 743,
				K: 547,
			},
			want: 546,
		},
		"case_4.196": {
			args: args{
				N: 360,
				K: 310,
			},
			want: 309,
		},
		"case_4.197": {
			args: args{
				N: 355,
				K: 84,
			},
			want: 83,
		},
		"case_4.198": {
			args: args{
				N: 195,
				K: 29,
			},
			want: 28,
		},
		"case_4.199": {
			args: args{
				N: 631,
				K: 488,
			},
			want: 487,
		},
		"case_4.200": {
			args: args{
				N: 700,
				K: 15,
			},
			want: 14,
		},
		"case_4.201": {
			args: args{
				N: 565,
				K: 28,
			},
			want: 27,
		},
		"case_4.202": {
			args: args{
				N: 641,
				K: 600,
			},
			want: 599,
		},
		"case_4.203": {
			args: args{
				N: 797,
				K: 717,
			},
			want: 716,
		},
		"case_4.204": {
			args: args{
				N: 551,
				K: 92,
			},
			want: 91,
		},
		"case_4.205": {
			args: args{
				N: 628,
				K: 64,
			},
			want: 63,
		},
		"case_4.206": {
			args: args{
				N: 400,
				K: 213,
			},
			want: 212,
		},
		"case_4.207": {
			args: args{
				N: 321,
				K: 194,
			},
			want: 193,
		},
		"case_4.208": {
			args: args{
				N: 228,
				K: 191,
			},
			want: 190,
		},
		"case_4.209": {
			args: args{
				N: 132,
				K: 24,
			},
			want: 23,
		},
		"case_4.210": {
			args: args{
				N: 520,
				K: 92,
			},
			want: 91,
		},
		"case_4.211": {
			args: args{
				N: 132,
				K: 126,
			},
			want: 125,
		},
		"case_4.212": {
			args: args{
				N: 152,
				K: 31,
			},
			want: 30,
		},
		"case_4.213": {
			args: args{
				N: 906,
				K: 219,
			},
			want: 218,
		},
		"case_4.214": {
			args: args{
				N: 272,
				K: 165,
			},
			want: 164,
		},
		"case_4.215": {
			args: args{
				N: 410,
				K: 385,
			},
			want: 384,
		},
		"case_4.216": {
			args: args{
				N: 865,
				K: 190,
			},
			want: 189,
		},
		"case_4.217": {
			args: args{
				N: 341,
				K: 239,
			},
			want: 238,
		},
		"case_4.218": {
			args: args{
				N: 933,
				K: 900,
			},
			want: 899,
		},
		"case_4.219": {
			args: args{
				N: 387,
				K: 264,
			},
			want: 263,
		},
		"case_4.220": {
			args: args{
				N: 743,
				K: 350,
			},
			want: 349,
		},
		"case_4.221": {
			args: args{
				N: 929,
				K: 216,
			},
			want: 215,
		},
		"case_4.222": {
			args: args{
				N: 686,
				K: 94,
			},
			want: 93,
		},
		"case_4.223": {
			args: args{
				N: 227,
				K: 168,
			},
			want: 167,
		},
		"case_4.224": {
			args: args{
				N: 554,
				K: 168,
			},
			want: 167,
		},
		"case_4.225": {
			args: args{
				N: 168,
				K: 74,
			},
			want: 73,
		},
		"case_4.226": {
			args: args{
				N: 778,
				K: 302,
			},
			want: 301,
		},
		"case_4.227": {
			args: args{
				N: 706,
				K: 39,
			},
			want: 38,
		},
		"case_4.228": {
			args: args{
				N: 328,
				K: 138,
			},
			want: 137,
		},
		"case_4.229": {
			args: args{
				N: 792,
				K: 432,
			},
			want: 431,
		},
		"case_4.230": {
			args: args{
				N: 372,
				K: 2,
			},
			want: 1,
		},
		"case_4.231": {
			args: args{
				N: 589,
				K: 477,
			},
			want: 476,
		},
		"case_4.232": {
			args: args{
				N: 436,
				K: 241,
			},
			want: 240,
		},
		"case_4.233": {
			args: args{
				N: 419,
				K: 346,
			},
			want: 345,
		},
		"case_4.234": {
			args: args{
				N: 184,
				K: 105,
			},
			want: 104,
		},
		"case_4.235": {
			args: args{
				N: 946,
				K: 214,
			},
			want: 213,
		},
		"case_4.236": {
			args: args{
				N: 769,
				K: 267,
			},
			want: 266,
		},
		"case_4.237": {
			args: args{
				N: 745,
				K: 411,
			},
			want: 410,
		},
		"case_4.238": {
			args: args{
				N: 108,
				K: 82,
			},
			want: 81,
		},
		"case_4.239": {
			args: args{
				N: 506,
				K: 228,
			},
			want: 227,
		},
		"case_4.240": {
			args: args{
				N: 203,
				K: 191,
			},
			want: 190,
		},
		"case_4.241": {
			args: args{
				N: 104,
				K: 81,
			},
			want: 80,
		},
		"case_4.242": {
			args: args{
				N: 391,
				K: 44,
			},
			want: 43,
		},
		"case_4.243": {
			args: args{
				N: 343,
				K: 196,
			},
			want: 195,
		},
		"case_4.244": {
			args: args{
				N: 824,
				K: 775,
			},
			want: 774,
		},
		"case_4.245": {
			args: args{
				N: 253,
				K: 175,
			},
			want: 174,
		},
		"case_4.246": {
			args: args{
				N: 912,
				K: 214,
			},
			want: 213,
		},
		"case_4.247": {
			args: args{
				N: 911,
				K: 274,
			},
			want: 273,
		},
		"case_4.248": {
			args: args{
				N: 924,
				K: 458,
			},
			want: 457,
		},
		"case_4.249": {
			args: args{
				N: 250,
				K: 143,
			},
			want: 142,
		},
		"case_4.250": {
			args: args{
				N: 875,
				K: 494,
			},
			want: 493,
		},
		"case_4.251": {
			args: args{
				N: 966,
				K: 336,
			},
			want: 335,
		},
		"case_4.252": {
			args: args{
				N: 698,
				K: 249,
			},
			want: 248,
		},
		"case_4.253": {
			args: args{
				N: 370,
				K: 128,
			},
			want: 127,
		},
		"case_4.254": {
			args: args{
				N: 578,
				K: 20,
			},
			want: 19,
		},
		"case_4.255": {
			args: args{
				N: 517,
				K: 89,
			},
			want: 88,
		},
		"case_4.256": {
			args: args{
				N: 625,
				K: 67,
			},
			want: 66,
		},
		"case_4.257": {
			args: args{
				N: 180,
				K: 62,
			},
			want: 61,
		},
		"case_4.258": {
			args: args{
				N: 561,
				K: 97,
			},
			want: 96,
		},
		"case_4.259": {
			args: args{
				N: 113,
				K: 70,
			},
			want: 69,
		},
		"case_4.260": {
			args: args{
				N: 837,
				K: 61,
			},
			want: 60,
		},
		"case_4.261": {
			args: args{
				N: 390,
				K: 99,
			},
			want: 98,
		},
		"case_4.262": {
			args: args{
				N: 970,
				K: 687,
			},
			want: 686,
		},
		"case_4.263": {
			args: args{
				N: 584,
				K: 284,
			},
			want: 283,
		},
		"case_4.264": {
			args: args{
				N: 476,
				K: 198,
			},
			want: 197,
		},
		"case_4.265": {
			args: args{
				N: 819,
				K: 284,
			},
			want: 283,
		},
		"case_4.266": {
			args: args{
				N: 106,
				K: 44,
			},
			want: 43,
		},
		"case_4.267": {
			args: args{
				N: 966,
				K: 889,
			},
			want: 888,
		},
		"case_4.268": {
			args: args{
				N: 375,
				K: 239,
			},
			want: 238,
		},
		"case_4.269": {
			args: args{
				N: 635,
				K: 118,
			},
			want: 117,
		},
		"case_4.270": {
			args: args{
				N: 174,
				K: 96,
			},
			want: 95,
		},
		"case_4.271": {
			args: args{
				N: 340,
				K: 185,
			},
			want: 184,
		},
		"case_4.272": {
			args: args{
				N: 127,
				K: 67,
			},
			want: 66,
		},
		"case_4.273": {
			args: args{
				N: 859,
				K: 784,
			},
			want: 783,
		},
		"case_4.274": {
			args: args{
				N: 865,
				K: 811,
			},
			want: 810,
		},
		"case_4.275": {
			args: args{
				N: 334,
				K: 266,
			},
			want: 265,
		},
		"case_4.276": {
			args: args{
				N: 698,
				K: 438,
			},
			want: 437,
		},
		"case_4.277": {
			args: args{
				N: 339,
				K: 330,
			},
			want: 329,
		},
		"case_4.278": {
			args: args{
				N: 281,
				K: 234,
			},
			want: 233,
		},
		"case_4.279": {
			args: args{
				N: 826,
				K: 275,
			},
			want: 274,
		},
		"case_4.280": {
			args: args{
				N: 968,
				K: 774,
			},
			want: 773,
		},
		"case_4.281": {
			args: args{
				N: 750,
				K: 87,
			},
			want: 86,
		},
		"case_4.282": {
			args: args{
				N: 277,
				K: 2,
			},
			want: 1,
		},
		"case_4.283": {
			args: args{
				N: 237,
				K: 209,
			},
			want: 208,
		},
		"case_4.284": {
			args: args{
				N: 750,
				K: 266,
			},
			want: 265,
		},
		"case_4.285": {
			args: args{
				N: 196,
				K: 56,
			},
			want: 55,
		},
		"case_4.286": {
			args: args{
				N: 590,
				K: 179,
			},
			want: 178,
		},
		"case_4.287": {
			args: args{
				N: 524,
				K: 384,
			},
			want: 383,
		},
		"case_4.288": {
			args: args{
				N: 240,
				K: 24,
			},
			want: 23,
		},
		"case_4.289": {
			args: args{
				N: 918,
				K: 155,
			},
			want: 154,
		},
		"case_4.290": {
			args: args{
				N: 666,
				K: 256,
			},
			want: 255,
		},
		"case_4.291": {
			args: args{
				N: 470,
				K: 383,
			},
			want: 382,
		},
		"case_4.292": {
			args: args{
				N: 106,
				K: 73,
			},
			want: 72,
		},
		"case_4.293": {
			args: args{
				N: 908,
				K: 536,
			},
			want: 535,
		},
		"case_4.294": {
			args: args{
				N: 147,
				K: 116,
			},
			want: 115,
		},
		"case_4.295": {
			args: args{
				N: 595,
				K: 482,
			},
			want: 481,
		},
		"case_4.296": {
			args: args{
				N: 744,
				K: 100,
			},
			want: 99,
		},
		"case_4.297": {
			args: args{
				N: 178,
				K: 28,
			},
			want: 27,
		},
		"case_4.298": {
			args: args{
				N: 699,
				K: 128,
			},
			want: 127,
		},
		"case_4.299": {
			args: args{
				N: 189,
				K: 53,
			},
			want: 52,
		},
		"case_4.300": {
			args: args{
				N: 872,
				K: 349,
			},
			want: 348,
		},
		"case_4.301": {
			args: args{
				N: 988,
				K: 575,
			},
			want: 574,
		},
		"case_4.302": {
			args: args{
				N: 989,
				K: 970,
			},
			want: 969,
		},
		"case_4.303": {
			args: args{
				N: 123,
				K: 28,
			},
			want: 27,
		},
		"case_4.304": {
			args: args{
				N: 924,
				K: 222,
			},
			want: 221,
		},
		"case_4.305": {
			args: args{
				N: 422,
				K: 136,
			},
			want: 135,
		},
		"case_4.306": {
			args: args{
				N: 347,
				K: 131,
			},
			want: 130,
		},
		"case_4.307": {
			args: args{
				N: 648,
				K: 293,
			},
			want: 292,
		},
		"case_4.308": {
			args: args{
				N: 476,
				K: 386,
			},
			want: 385,
		},
		"case_4.309": {
			args: args{
				N: 777,
				K: 380,
			},
			want: 379,
		},
		"case_4.310": {
			args: args{
				N: 633,
				K: 129,
			},
			want: 128,
		},
		"case_4.311": {
			args: args{
				N: 445,
				K: 53,
			},
			want: 52,
		},
		"case_4.312": {
			args: args{
				N: 202,
				K: 63,
			},
			want: 62,
		},
		"case_4.313": {
			args: args{
				N: 188,
				K: 75,
			},
			want: 74,
		},
		"case_4.314": {
			args: args{
				N: 225,
				K: 147,
			},
			want: 146,
		},
		"case_4.315": {
			args: args{
				N: 961,
				K: 896,
			},
			want: 894,
		},
		"case_4.316": {
			args: args{
				N: 123,
				K: 122,
			},
			want: 121,
		},
		"case_4.317": {
			args: args{
				N: 187,
				K: 178,
			},
			want: 177,
		},
		"case_4.318": {
			args: args{
				N: 838,
				K: 650,
			},
			want: 649,
		},
		"case_4.319": {
			args: args{
				N: 607,
				K: 306,
			},
			want: 305,
		},
		"case_4.320": {
			args: args{
				N: 244,
				K: 99,
			},
			want: 98,
		},
		"case_4.321": {
			args: args{
				N: 191,
				K: 55,
			},
			want: 54,
		},
		"case_4.322": {
			args: args{
				N: 634,
				K: 349,
			},
			want: 348,
		},
		"case_4.323": {
			args: args{
				N: 787,
				K: 599,
			},
			want: 598,
		},
		"case_4.324": {
			args: args{
				N: 303,
				K: 221,
			},
			want: 220,
		},
		"case_4.325": {
			args: args{
				N: 862,
				K: 543,
			},
			want: 542,
		},
		"case_4.326": {
			args: args{
				N: 426,
				K: 162,
			},
			want: 161,
		},
		"case_4.327": {
			args: args{
				N: 908,
				K: 173,
			},
			want: 172,
		},
		"case_4.328": {
			args: args{
				N: 559,
				K: 525,
			},
			want: 524,
		},
		"case_4.329": {
			args: args{
				N: 707,
				K: 406,
			},
			want: 405,
		},
		"case_4.330": {
			args: args{
				N: 130,
				K: 61,
			},
			want: 60,
		},
		"case_4.331": {
			args: args{
				N: 548,
				K: 533,
			},
			want: 532,
		},
		"case_4.332": {
			args: args{
				N: 459,
				K: 155,
			},
			want: 154,
		},
		"case_4.333": {
			args: args{
				N: 435,
				K: 378,
			},
			want: 377,
		},
		"case_4.334": {
			args: args{
				N: 653,
				K: 279,
			},
			want: 278,
		},
		"case_4.335": {
			args: args{
				N: 114,
				K: 9,
			},
			want: 8,
		},
		"case_4.336": {
			args: args{
				N: 131,
				K: 7,
			},
			want: 6,
		},
		"case_4.337": {
			args: args{
				N: 393,
				K: 347,
			},
			want: 346,
		},
		"case_4.338": {
			args: args{
				N: 162,
				K: 45,
			},
			want: 44,
		},
		"case_4.339": {
			args: args{
				N: 844,
				K: 813,
			},
			want: 812,
		},
		"case_4.340": {
			args: args{
				N: 562,
				K: 522,
			},
			want: 521,
		},
		"case_4.341": {
			args: args{
				N: 731,
				K: 328,
			},
			want: 327,
		},
		"case_4.342": {
			args: args{
				N: 630,
				K: 612,
			},
			want: 611,
		},
		"case_4.343": {
			args: args{
				N: 487,
				K: 434,
			},
			want: 433,
		},
		"case_4.344": {
			args: args{
				N: 398,
				K: 304,
			},
			want: 303,
		},
		"case_4.345": {
			args: args{
				N: 157,
				K: 89,
			},
			want: 88,
		},
		"case_4.346": {
			args: args{
				N: 583,
				K: 329,
			},
			want: 328,
		},
		"case_4.347": {
			args: args{
				N: 947,
				K: 566,
			},
			want: 565,
		},
		"case_4.348": {
			args: args{
				N: 979,
				K: 594,
			},
			want: 593,
		},
		"case_4.349": {
			args: args{
				N: 190,
				K: 27,
			},
			want: 26,
		},
		"case_4.350": {
			args: args{
				N: 359,
				K: 198,
			},
			want: 197,
		},
		"case_4.351": {
			args: args{
				N: 969,
				K: 231,
			},
			want: 230,
		},
		"case_4.352": {
			args: args{
				N: 468,
				K: 422,
			},
			want: 421,
		},
		"case_4.353": {
			args: args{
				N: 413,
				K: 76,
			},
			want: 75,
		},
		"case_4.354": {
			args: args{
				N: 998,
				K: 80,
			},
			want: 79,
		},
		"case_4.355": {
			args: args{
				N: 443,
				K: 153,
			},
			want: 152,
		},
		"case_4.356": {
			args: args{
				N: 635,
				K: 282,
			},
			want: 281,
		},
		"case_4.357": {
			args: args{
				N: 843,
				K: 744,
			},
			want: 743,
		},
		"case_4.358": {
			args: args{
				N: 560,
				K: 349,
			},
			want: 348,
		},
		"case_4.359": {
			args: args{
				N: 272,
				K: 150,
			},
			want: 149,
		},
		"case_4.360": {
			args: args{
				N: 905,
				K: 714,
			},
			want: 713,
		},
		"case_4.361": {
			args: args{
				N: 176,
				K: 126,
			},
			want: 125,
		},
		"case_4.362": {
			args: args{
				N: 323,
				K: 147,
			},
			want: 146,
		},
		"case_4.363": {
			args: args{
				N: 291,
				K: 124,
			},
			want: 123,
		},
		"case_4.364": {
			args: args{
				N: 799,
				K: 665,
			},
			want: 664,
		},
		"case_4.365": {
			args: args{
				N: 488,
				K: 392,
			},
			want: 391,
		},
		"case_4.366": {
			args: args{
				N: 953,
				K: 409,
			},
			want: 408,
		},
		"case_4.367": {
			args: args{
				N: 533,
				K: 365,
			},
			want: 364,
		},
		"case_4.368": {
			args: args{
				N: 680,
				K: 259,
			},
			want: 258,
		},
		"case_4.369": {
			args: args{
				N: 815,
				K: 556,
			},
			want: 555,
		},
		"case_4.370": {
			args: args{
				N: 598,
				K: 209,
			},
			want: 208,
		},
		"case_4.371": {
			args: args{
				N: 411,
				K: 258,
			},
			want: 257,
		},
		"case_4.372": {
			args: args{
				N: 755,
				K: 72,
			},
			want: 71,
		},
		"case_4.373": {
			args: args{
				N: 422,
				K: 244,
			},
			want: 243,
		},
		"case_4.374": {
			args: args{
				N: 845,
				K: 687,
			},
			want: 686,
		},
		"case_4.375": {
			args: args{
				N: 846,
				K: 21,
			},
			want: 20,
		},
		"case_4.376": {
			args: args{
				N: 272,
				K: 29,
			},
			want: 28,
		},
		"case_4.377": {
			args: args{
				N: 693,
				K: 537,
			},
			want: 536,
		},
		"case_4.378": {
			args: args{
				N: 881,
				K: 236,
			},
			want: 235,
		},
		"case_4.379": {
			args: args{
				N: 154,
				K: 35,
			},
			want: 34,
		},
		"case_4.380": {
			args: args{
				N: 197,
				K: 3,
			},
			want: 2,
		},
		"case_4.381": {
			args: args{
				N: 151,
				K: 55,
			},
			want: 54,
		},
		"case_4.382": {
			args: args{
				N: 663,
				K: 346,
			},
			want: 345,
		},
		"case_4.383": {
			args: args{
				N: 685,
				K: 252,
			},
			want: 251,
		},
		"case_4.384": {
			args: args{
				N: 603,
				K: 287,
			},
			want: 286,
		},
		"case_4.385": {
			args: args{
				N: 114,
				K: 29,
			},
			want: 28,
		},
		"case_4.386": {
			args: args{
				N: 884,
				K: 391,
			},
			want: 390,
		},
		"case_4.387": {
			args: args{
				N: 110,
				K: 8,
			},
			want: 7,
		},
		"case_4.388": {
			args: args{
				N: 107,
				K: 54,
			},
			want: 53,
		},
		"case_4.389": {
			args: args{
				N: 159,
				K: 101,
			},
			want: 100,
		},
		"case_4.390": {
			args: args{
				N: 669,
				K: 517,
			},
			want: 516,
		},
		"case_4.391": {
			args: args{
				N: 842,
				K: 109,
			},
			want: 108,
		},
		"case_4.392": {
			args: args{
				N: 568,
				K: 332,
			},
			want: 331,
		},
		"case_4.393": {
			args: args{
				N: 503,
				K: 51,
			},
			want: 50,
		},
		"case_4.394": {
			args: args{
				N: 114,
				K: 7,
			},
			want: 6,
		},
		"case_4.395": {
			args: args{
				N: 605,
				K: 35,
			},
			want: 34,
		},
		"case_4.396": {
			args: args{
				N: 387,
				K: 246,
			},
			want: 245,
		},
		"case_4.397": {
			args: args{
				N: 387,
				K: 130,
			},
			want: 129,
		},
		"case_4.398": {
			args: args{
				N: 167,
				K: 47,
			},
			want: 46,
		},
		"case_4.399": {
			args: args{
				N: 486,
				K: 307,
			},
			want: 306,
		},
		"case_4.400": {
			args: args{
				N: 538,
				K: 345,
			},
			want: 344,
		},
		"case_4.401": {
			args: args{
				N: 177,
				K: 176,
			},
			want: 174,
		},
		"case_4.402": {
			args: args{
				N: 732,
				K: 309,
			},
			want: 308,
		},
		"case_4.403": {
			args: args{
				N: 771,
				K: 221,
			},
			want: 220,
		},
		"case_4.404": {
			args: args{
				N: 146,
				K: 30,
			},
			want: 29,
		},
		"case_4.405": {
			args: args{
				N: 264,
				K: 3,
			},
			want: 2,
		},
		"case_4.406": {
			args: args{
				N: 114,
				K: 79,
			},
			want: 78,
		},
		"case_4.407": {
			args: args{
				N: 818,
				K: 633,
			},
			want: 632,
		},
		"case_4.408": {
			args: args{
				N: 215,
				K: 133,
			},
			want: 132,
		},
		"case_4.409": {
			args: args{
				N: 191,
				K: 151,
			},
			want: 150,
		},
		"case_4.410": {
			args: args{
				N: 180,
				K: 170,
			},
			want: 169,
		},
		"case_4.411": {
			args: args{
				N: 936,
				K: 814,
			},
			want: 813,
		},
		"case_4.412": {
			args: args{
				N: 396,
				K: 7,
			},
			want: 6,
		},
		"case_4.413": {
			args: args{
				N: 311,
				K: 3,
			},
			want: 2,
		},
		"case_4.414": {
			args: args{
				N: 737,
				K: 210,
			},
			want: 209,
		},
		"case_4.415": {
			args: args{
				N: 512,
				K: 440,
			},
			want: 439,
		},
		"case_4.416": {
			args: args{
				N: 296,
				K: 25,
			},
			want: 24,
		},
		"case_4.417": {
			args: args{
				N: 723,
				K: 356,
			},
			want: 355,
		},
		"case_4.418": {
			args: args{
				N: 353,
				K: 171,
			},
			want: 170,
		},
		"case_4.419": {
			args: args{
				N: 610,
				K: 316,
			},
			want: 315,
		},
		"case_4.420": {
			args: args{
				N: 383,
				K: 2,
			},
			want: 1,
		},
		"case_4.421": {
			args: args{
				N: 132,
				K: 61,
			},
			want: 60,
		},
		"case_4.422": {
			args: args{
				N: 457,
				K: 210,
			},
			want: 209,
		},
		"case_4.423": {
			args: args{
				N: 752,
				K: 678,
			},
			want: 677,
		},
		"case_4.424": {
			args: args{
				N: 766,
				K: 301,
			},
			want: 300,
		},
		"case_4.425": {
			args: args{
				N: 168,
				K: 53,
			},
			want: 52,
		},
		"case_4.426": {
			args: args{
				N: 793,
				K: 753,
			},
			want: 752,
		},
		"case_4.427": {
			args: args{
				N: 487,
				K: 199,
			},
			want: 198,
		},
		"case_4.428": {
			args: args{
				N: 422,
				K: 41,
			},
			want: 40,
		},
		"case_4.429": {
			args: args{
				N: 370,
				K: 306,
			},
			want: 305,
		},
		"case_4.430": {
			args: args{
				N: 857,
				K: 435,
			},
			want: 434,
		},
		"case_4.431": {
			args: args{
				N: 362,
				K: 31,
			},
			want: 30,
		},
		"case_4.432": {
			args: args{
				N: 462,
				K: 384,
			},
			want: 382,
		},
		"case_4.433": {
			args: args{
				N: 727,
				K: 648,
			},
			want: 647,
		},
		"case_4.434": {
			args: args{
				N: 963,
				K: 812,
			},
			want: 811,
		},
		"case_4.435": {
			args: args{
				N: 103,
				K: 6,
			},
			want: 5,
		},
		"case_4.436": {
			args: args{
				N: 336,
				K: 183,
			},
			want: 182,
		},
		"case_4.437": {
			args: args{
				N: 188,
				K: 91,
			},
			want: 90,
		},
		"case_4.438": {
			args: args{
				N: 136,
				K: 16,
			},
			want: 15,
		},
		"case_4.439": {
			args: args{
				N: 137,
				K: 9,
			},
			want: 8,
		},
		"case_4.440": {
			args: args{
				N: 343,
				K: 257,
			},
			want: 256,
		},
		"case_4.441": {
			args: args{
				N: 138,
				K: 119,
			},
			want: 118,
		},
		"case_4.442": {
			args: args{
				N: 683,
				K: 570,
			},
			want: 569,
		},
		"case_4.443": {
			args: args{
				N: 706,
				K: 55,
			},
			want: 54,
		},
		"case_4.444": {
			args: args{
				N: 877,
				K: 700,
			},
			want: 699,
		},
		"case_4.445": {
			args: args{
				N: 175,
				K: 102,
			},
			want: 101,
		},
		"case_4.446": {
			args: args{
				N: 280,
				K: 2,
			},
			want: 1,
		},
		"case_4.447": {
			args: args{
				N: 976,
				K: 794,
			},
			want: 793,
		},
		"case_4.448": {
			args: args{
				N: 613,
				K: 473,
			},
			want: 472,
		},
		"case_4.449": {
			args: args{
				N: 189,
				K: 119,
			},
			want: 118,
		},
		"case_4.450": {
			args: args{
				N: 376,
				K: 26,
			},
			want: 25,
		},
		"case_4.451": {
			args: args{
				N: 497,
				K: 69,
			},
			want: 68,
		},
		"case_4.452": {
			args: args{
				N: 976,
				K: 899,
			},
			want: 898,
		},
		"case_4.453": {
			args: args{
				N: 890,
				K: 197,
			},
			want: 196,
		},
		"case_4.454": {
			args: args{
				N: 925,
				K: 388,
			},
			want: 387,
		},
		"case_4.455": {
			args: args{
				N: 165,
				K: 22,
			},
			want: 21,
		},
		"case_4.456": {
			args: args{
				N: 142,
				K: 109,
			},
			want: 108,
		},
		"case_4.457": {
			args: args{
				N: 626,
				K: 2,
			},
			want: 1,
		},
		"case_4.458": {
			args: args{
				N: 316,
				K: 127,
			},
			want: 126,
		},
		"case_4.459": {
			args: args{
				N: 101,
				K: 29,
			},
			want: 28,
		},
		"case_4.460": {
			args: args{
				N: 473,
				K: 361,
			},
			want: 360,
		},
		"case_4.461": {
			args: args{
				N: 779,
				K: 540,
			},
			want: 539,
		},
		"case_4.462": {
			args: args{
				N: 693,
				K: 174,
			},
			want: 173,
		},
		"case_4.463": {
			args: args{
				N: 566,
				K: 44,
			},
			want: 43,
		},
		"case_4.464": {
			args: args{
				N: 182,
				K: 8,
			},
			want: 7,
		},
		"case_4.465": {
			args: args{
				N: 829,
				K: 191,
			},
			want: 190,
		},
		"case_4.466": {
			args: args{
				N: 261,
				K: 98,
			},
			want: 97,
		},
		"case_4.467": {
			args: args{
				N: 670,
				K: 529,
			},
			want: 528,
		},
		"case_4.468": {
			args: args{
				N: 603,
				K: 190,
			},
			want: 189,
		},
		"case_4.469": {
			args: args{
				N: 128,
				K: 60,
			},
			want: 59,
		},
		"case_4.470": {
			args: args{
				N: 663,
				K: 176,
			},
			want: 175,
		},
		"case_4.471": {
			args: args{
				N: 717,
				K: 438,
			},
			want: 437,
		},
		"case_4.472": {
			args: args{
				N: 439,
				K: 139,
			},
			want: 138,
		},
		"case_4.473": {
			args: args{
				N: 666,
				K: 98,
			},
			want: 97,
		},
		"case_4.474": {
			args: args{
				N: 664,
				K: 337,
			},
			want: 336,
		},
		"case_4.475": {
			args: args{
				N: 166,
				K: 16,
			},
			want: 15,
		},
		"case_4.476": {
			args: args{
				N: 738,
				K: 411,
			},
			want: 410,
		},
		"case_4.477": {
			args: args{
				N: 287,
				K: 75,
			},
			want: 74,
		},
		"case_4.478": {
			args: args{
				N: 773,
				K: 523,
			},
			want: 522,
		},
		"case_4.479": {
			args: args{
				N: 547,
				K: 523,
			},
			want: 522,
		},
		"case_4.480": {
			args: args{
				N: 996,
				K: 621,
			},
			want: 620,
		},
		"case_4.481": {
			args: args{
				N: 524,
				K: 12,
			},
			want: 11,
		},
		"case_4.482": {
			args: args{
				N: 406,
				K: 30,
			},
			want: 29,
		},
		"case_4.483": {
			args: args{
				N: 999,
				K: 737,
			},
			want: 736,
		},
		"case_4.484": {
			args: args{
				N: 841,
				K: 91,
			},
			want: 90,
		},
		"case_4.485": {
			args: args{
				N: 600,
				K: 323,
			},
			want: 322,
		},
		"case_4.486": {
			args: args{
				N: 799,
				K: 499,
			},
			want: 498,
		},
		"case_4.487": {
			args: args{
				N: 376,
				K: 61,
			},
			want: 60,
		},
		"case_4.488": {
			args: args{
				N: 268,
				K: 148,
			},
			want: 147,
		},
		"case_4.489": {
			args: args{
				N: 358,
				K: 162,
			},
			want: 161,
		},
		"case_4.490": {
			args: args{
				N: 838,
				K: 565,
			},
			want: 564,
		},
		"case_4.491": {
			args: args{
				N: 178,
				K: 39,
			},
			want: 38,
		},
		"case_4.492": {
			args: args{
				N: 149,
				K: 38,
			},
			want: 37,
		},
		"case_4.493": {
			args: args{
				N: 164,
				K: 76,
			},
			want: 75,
		},
		"case_4.494": {
			args: args{
				N: 933,
				K: 309,
			},
			want: 308,
		},
		"case_4.495": {
			args: args{
				N: 636,
				K: 134,
			},
			want: 133,
		},
		"case_4.496": {
			args: args{
				N: 977,
				K: 482,
			},
			want: 481,
		},
		"case_4.497": {
			args: args{
				N: 943,
				K: 556,
			},
			want: 555,
		},
		"case_4.498": {
			args: args{
				N: 325,
				K: 148,
			},
			want: 147,
		},
		"case_4.499": {
			args: args{
				N: 571,
				K: 125,
			},
			want: 124,
		},
		"case_4.500": {
			args: args{
				N: 876,
				K: 501,
			},
			want: 500,
		},
		"case_4.501": {
			args: args{
				N: 530,
				K: 490,
			},
			want: 489,
		},
		"case_4.502": {
			args: args{
				N: 192,
				K: 70,
			},
			want: 69,
		},
		"case_4.503": {
			args: args{
				N: 724,
				K: 604,
			},
			want: 603,
		},
		"case_4.504": {
			args: args{
				N: 212,
				K: 37,
			},
			want: 36,
		},
		"case_4.505": {
			args: args{
				N: 586,
				K: 466,
			},
			want: 465,
		},
		"case_4.506": {
			args: args{
				N: 300,
				K: 148,
			},
			want: 147,
		},
		"case_4.507": {
			args: args{
				N: 179,
				K: 120,
			},
			want: 119,
		},
		"case_4.508": {
			args: args{
				N: 116,
				K: 114,
			},
			want: 113,
		},
		"case_4.509": {
			args: args{
				N: 934,
				K: 786,
			},
			want: 785,
		},
		"case_4.510": {
			args: args{
				N: 509,
				K: 15,
			},
			want: 14,
		},
		"case_4.511": {
			args: args{
				N: 120,
				K: 92,
			},
			want: 91,
		},
		"case_4.512": {
			args: args{
				N: 670,
				K: 617,
			},
			want: 616,
		},
		"case_4.513": {
			args: args{
				N: 575,
				K: 435,
			},
			want: 434,
		},
		"case_4.514": {
			args: args{
				N: 747,
				K: 645,
			},
			want: 644,
		},
		"case_4.515": {
			args: args{
				N: 949,
				K: 890,
			},
			want: 889,
		},
		"case_4.516": {
			args: args{
				N: 179,
				K: 77,
			},
			want: 76,
		},
		"case_4.517": {
			args: args{
				N: 959,
				K: 77,
			},
			want: 76,
		},
		"case_4.518": {
			args: args{
				N: 904,
				K: 291,
			},
			want: 290,
		},
		"case_4.519": {
			args: args{
				N: 867,
				K: 14,
			},
			want: 13,
		},
		"case_4.520": {
			args: args{
				N: 722,
				K: 370,
			},
			want: 369,
		},
		"case_4.521": {
			args: args{
				N: 170,
				K: 112,
			},
			want: 111,
		},
		"case_4.522": {
			args: args{
				N: 151,
				K: 114,
			},
			want: 113,
		},
		"case_4.523": {
			args: args{
				N: 654,
				K: 274,
			},
			want: 273,
		},
		"case_4.524": {
			args: args{
				N: 521,
				K: 228,
			},
			want: 227,
		},
		"case_4.525": {
			args: args{
				N: 679,
				K: 341,
			},
			want: 340,
		},
		"case_4.526": {
			args: args{
				N: 714,
				K: 13,
			},
			want: 12,
		},
		"case_4.527": {
			args: args{
				N: 487,
				K: 399,
			},
			want: 398,
		},
		"case_4.528": {
			args: args{
				N: 436,
				K: 233,
			},
			want: 232,
		},
		"case_4.529": {
			args: args{
				N: 659,
				K: 533,
			},
			want: 532,
		},
		"case_4.530": {
			args: args{
				N: 284,
				K: 136,
			},
			want: 135,
		},
		"case_4.531": {
			args: args{
				N: 618,
				K: 19,
			},
			want: 18,
		},
		"case_4.532": {
			args: args{
				N: 503,
				K: 433,
			},
			want: 432,
		},
		"case_4.533": {
			args: args{
				N: 876,
				K: 231,
			},
			want: 230,
		},
		"case_4.534": {
			args: args{
				N: 391,
				K: 104,
			},
			want: 103,
		},
		"case_4.535": {
			args: args{
				N: 234,
				K: 24,
			},
			want: 23,
		},
		"case_4.536": {
			args: args{
				N: 871,
				K: 387,
			},
			want: 386,
		},
		"case_4.537": {
			args: args{
				N: 373,
				K: 31,
			},
			want: 30,
		},
		"case_4.538": {
			args: args{
				N: 596,
				K: 512,
			},
			want: 510,
		},
		"case_4.539": {
			args: args{
				N: 867,
				K: 79,
			},
			want: 78,
		},
		"case_4.540": {
			args: args{
				N: 654,
				K: 323,
			},
			want: 322,
		},
		"case_4.541": {
			args: args{
				N: 711,
				K: 184,
			},
			want: 183,
		},
		"case_4.542": {
			args: args{
				N: 261,
				K: 142,
			},
			want: 141,
		},
		"case_4.543": {
			args: args{
				N: 273,
				K: 95,
			},
			want: 94,
		},
		"case_4.544": {
			args: args{
				N: 874,
				K: 131,
			},
			want: 130,
		},
		"case_4.545": {
			args: args{
				N: 140,
				K: 100,
			},
			want: 99,
		},
		"case_4.546": {
			args: args{
				N: 123,
				K: 30,
			},
			want: 29,
		},
		"case_4.547": {
			args: args{
				N: 527,
				K: 198,
			},
			want: 197,
		},
		"case_4.548": {
			args: args{
				N: 740,
				K: 573,
			},
			want: 572,
		},
		"case_4.549": {
			args: args{
				N: 688,
				K: 389,
			},
			want: 388,
		},
		"case_4.550": {
			args: args{
				N: 119,
				K: 87,
			},
			want: 86,
		},
		"case_4.551": {
			args: args{
				N: 881,
				K: 612,
			},
			want: 611,
		},
		"case_4.552": {
			args: args{
				N: 594,
				K: 85,
			},
			want: 84,
		},
		"case_4.553": {
			args: args{
				N: 984,
				K: 692,
			},
			want: 691,
		},
		"case_4.554": {
			args: args{
				N: 110,
				K: 18,
			},
			want: 17,
		},
		"case_4.555": {
			args: args{
				N: 839,
				K: 78,
			},
			want: 77,
		},
		"case_4.556": {
			args: args{
				N: 584,
				K: 169,
			},
			want: 168,
		},
		"case_4.557": {
			args: args{
				N: 152,
				K: 118,
			},
			want: 117,
		},
		"case_4.558": {
			args: args{
				N: 149,
				K: 140,
			},
			want: 139,
		},
		"case_4.559": {
			args: args{
				N: 426,
				K: 232,
			},
			want: 231,
		},
		"case_4.560": {
			args: args{
				N: 402,
				K: 136,
			},
			want: 135,
		},
		"case_4.561": {
			args: args{
				N: 752,
				K: 538,
			},
			want: 537,
		},
		"case_4.562": {
			args: args{
				N: 729,
				K: 577,
			},
			want: 576,
		},
		"case_4.563": {
			args: args{
				N: 355,
				K: 225,
			},
			want: 224,
		},
		"case_4.564": {
			args: args{
				N: 871,
				K: 498,
			},
			want: 497,
		},
		"case_4.565": {
			args: args{
				N: 652,
				K: 651,
			},
			want: 650,
		},
		"case_4.566": {
			args: args{
				N: 362,
				K: 158,
			},
			want: 157,
		},
		"case_4.567": {
			args: args{
				N: 417,
				K: 235,
			},
			want: 234,
		},
		"case_4.568": {
			args: args{
				N: 474,
				K: 11,
			},
			want: 10,
		},
		"case_4.569": {
			args: args{
				N: 292,
				K: 103,
			},
			want: 102,
		},
		"case_4.570": {
			args: args{
				N: 409,
				K: 257,
			},
			want: 256,
		},
		"case_4.571": {
			args: args{
				N: 256,
				K: 53,
			},
			want: 52,
		},
		"case_4.572": {
			args: args{
				N: 780,
				K: 204,
			},
			want: 203,
		},
		"case_4.573": {
			args: args{
				N: 160,
				K: 141,
			},
			want: 140,
		},
		"case_4.574": {
			args: args{
				N: 795,
				K: 352,
			},
			want: 351,
		},
		"case_4.575": {
			args: args{
				N: 326,
				K: 177,
			},
			want: 176,
		},
		"case_4.576": {
			args: args{
				N: 173,
				K: 2,
			},
			want: 1,
		},
		"case_4.577": {
			args: args{
				N: 125,
				K: 121,
			},
			want: 120,
		},
		"case_4.578": {
			args: args{
				N: 798,
				K: 548,
			},
			want: 547,
		},
		"case_4.579": {
			args: args{
				N: 567,
				K: 492,
			},
			want: 491,
		},
		"case_4.580": {
			args: args{
				N: 788,
				K: 637,
			},
			want: 636,
		},
		"case_4.581": {
			args: args{
				N: 176,
				K: 145,
			},
			want: 144,
		},
		"case_4.582": {
			args: args{
				N: 645,
				K: 181,
			},
			want: 180,
		},
		"case_4.583": {
			args: args{
				N: 971,
				K: 141,
			},
			want: 140,
		},
		"case_4.584": {
			args: args{
				N: 309,
				K: 181,
			},
			want: 180,
		},
		"case_4.585": {
			args: args{
				N: 605,
				K: 83,
			},
			want: 82,
		},
		"case_4.586": {
			args: args{
				N: 177,
				K: 131,
			},
			want: 130,
		},
		"case_4.587": {
			args: args{
				N: 576,
				K: 424,
			},
			want: 423,
		},
		"case_4.588": {
			args: args{
				N: 522,
				K: 134,
			},
			want: 133,
		},
		"case_4.589": {
			args: args{
				N: 455,
				K: 449,
			},
			want: 448,
		},
		"case_4.590": {
			args: args{
				N: 134,
				K: 18,
			},
			want: 17,
		},
		"case_4.591": {
			args: args{
				N: 244,
				K: 201,
			},
			want: 200,
		},
		"case_4.592": {
			args: args{
				N: 173,
				K: 5,
			},
			want: 4,
		},
		"case_4.593": {
			args: args{
				N: 149,
				K: 68,
			},
			want: 67,
		},
		"case_4.594": {
			args: args{
				N: 876,
				K: 225,
			},
			want: 224,
		},
		"case_4.595": {
			args: args{
				N: 671,
				K: 650,
			},
			want: 649,
		},
		"case_4.596": {
			args: args{
				N: 360,
				K: 38,
			},
			want: 37,
		},
		"case_4.597": {
			args: args{
				N: 876,
				K: 453,
			},
			want: 452,
		},
		"case_4.598": {
			args: args{
				N: 465,
				K: 77,
			},
			want: 76,
		},
		"case_4.599": {
			args: args{
				N: 123,
				K: 87,
			},
			want: 86,
		},
		"case_4.600": {
			args: args{
				N: 594,
				K: 432,
			},
			want: 431,
		},
		"case_4.601": {
			args: args{
				N: 552,
				K: 34,
			},
			want: 33,
		},
		"case_4.602": {
			args: args{
				N: 617,
				K: 330,
			},
			want: 329,
		},
		"case_4.603": {
			args: args{
				N: 650,
				K: 144,
			},
			want: 143,
		},
		"case_4.604": {
			args: args{
				N: 760,
				K: 145,
			},
			want: 144,
		},
		"case_4.605": {
			args: args{
				N: 291,
				K: 75,
			},
			want: 74,
		},
		"case_4.606": {
			args: args{
				N: 150,
				K: 71,
			},
			want: 70,
		},
		"case_4.607": {
			args: args{
				N: 804,
				K: 90,
			},
			want: 89,
		},
		"case_4.608": {
			args: args{
				N: 561,
				K: 187,
			},
			want: 186,
		},
		"case_4.609": {
			args: args{
				N: 917,
				K: 646,
			},
			want: 645,
		},
		"case_4.610": {
			args: args{
				N: 960,
				K: 333,
			},
			want: 332,
		},
		"case_4.611": {
			args: args{
				N: 601,
				K: 108,
			},
			want: 107,
		},
		"case_4.612": {
			args: args{
				N: 809,
				K: 17,
			},
			want: 16,
		},
		"case_4.613": {
			args: args{
				N: 478,
				K: 83,
			},
			want: 82,
		},
		"case_4.614": {
			args: args{
				N: 135,
				K: 41,
			},
			want: 40,
		},
		"case_4.615": {
			args: args{
				N: 181,
				K: 29,
			},
			want: 28,
		},
		"case_4.616": {
			args: args{
				N: 823,
				K: 377,
			},
			want: 376,
		},
		"case_4.617": {
			args: args{
				N: 270,
				K: 164,
			},
			want: 163,
		},
		"case_4.618": {
			args: args{
				N: 302,
				K: 147,
			},
			want: 146,
		},
		"case_4.619": {
			args: args{
				N: 845,
				K: 613,
			},
			want: 612,
		},
		"case_4.620": {
			args: args{
				N: 864,
				K: 290,
			},
			want: 289,
		},
		"case_4.621": {
			args: args{
				N: 565,
				K: 194,
			},
			want: 193,
		},
		"case_4.622": {
			args: args{
				N: 290,
				K: 5,
			},
			want: 4,
		},
		"case_4.623": {
			args: args{
				N: 956,
				K: 896,
			},
			want: 894,
		},
		"case_4.624": {
			args: args{
				N: 344,
				K: 211,
			},
			want: 210,
		},
		"case_4.625": {
			args: args{
				N: 155,
				K: 102,
			},
			want: 101,
		},
		"case_4.626": {
			args: args{
				N: 406,
				K: 71,
			},
			want: 70,
		},
		"case_4.627": {
			args: args{
				N: 352,
				K: 151,
			},
			want: 150,
		},
		"case_4.628": {
			args: args{
				N: 378,
				K: 281,
			},
			want: 280,
		},
		"case_4.629": {
			args: args{
				N: 317,
				K: 268,
			},
			want: 267,
		},
		"case_4.630": {
			args: args{
				N: 759,
				K: 17,
			},
			want: 16,
		},
		"case_4.631": {
			args: args{
				N: 259,
				K: 184,
			},
			want: 183,
		},
		"case_4.632": {
			args: args{
				N: 570,
				K: 226,
			},
			want: 225,
		},
		"case_4.633": {
			args: args{
				N: 326,
				K: 39,
			},
			want: 38,
		},
		"case_4.634": {
			args: args{
				N: 797,
				K: 604,
			},
			want: 603,
		},
		"case_4.635": {
			args: args{
				N: 446,
				K: 352,
			},
			want: 351,
		},
		"case_4.636": {
			args: args{
				N: 367,
				K: 94,
			},
			want: 93,
		},
		"case_4.637": {
			args: args{
				N: 842,
				K: 725,
			},
			want: 724,
		},
		"case_4.638": {
			args: args{
				N: 827,
				K: 548,
			},
			want: 547,
		},
		"case_4.639": {
			args: args{
				N: 116,
				K: 91,
			},
			want: 90,
		},
		"case_4.640": {
			args: args{
				N: 758,
				K: 297,
			},
			want: 296,
		},
		"case_4.641": {
			args: args{
				N: 545,
				K: 141,
			},
			want: 140,
		},
		"case_4.642": {
			args: args{
				N: 180,
				K: 168,
			},
			want: 167,
		},
		"case_4.643": {
			args: args{
				N: 885,
				K: 533,
			},
			want: 532,
		},
		"case_4.644": {
			args: args{
				N: 940,
				K: 810,
			},
			want: 809,
		},
		"case_4.645": {
			args: args{
				N: 735,
				K: 150,
			},
			want: 149,
		},
		"case_4.646": {
			args: args{
				N: 775,
				K: 294,
			},
			want: 293,
		},
		"case_4.647": {
			args: args{
				N: 855,
				K: 666,
			},
			want: 665,
		},
		"case_4.648": {
			args: args{
				N: 854,
				K: 749,
			},
			want: 748,
		},
		"case_4.649": {
			args: args{
				N: 872,
				K: 735,
			},
			want: 734,
		},
		"case_4.650": {
			args: args{
				N: 787,
				K: 748,
			},
			want: 747,
		},
		"case_4.651": {
			args: args{
				N: 228,
				K: 164,
			},
			want: 163,
		},
		"case_4.652": {
			args: args{
				N: 608,
				K: 197,
			},
			want: 196,
		},
		"case_4.653": {
			args: args{
				N: 114,
				K: 11,
			},
			want: 10,
		},
		"case_4.654": {
			args: args{
				N: 645,
				K: 84,
			},
			want: 83,
		},
		"case_4.655": {
			args: args{
				N: 233,
				K: 150,
			},
			want: 149,
		},
		"case_4.656": {
			args: args{
				N: 723,
				K: 64,
			},
			want: 63,
		},
		"case_4.657": {
			args: args{
				N: 532,
				K: 293,
			},
			want: 292,
		},
		"case_4.658": {
			args: args{
				N: 626,
				K: 603,
			},
			want: 602,
		},
		"case_4.659": {
			args: args{
				N: 778,
				K: 180,
			},
			want: 179,
		},
		"case_4.660": {
			args: args{
				N: 392,
				K: 291,
			},
			want: 290,
		},
		"case_4.661": {
			args: args{
				N: 151,
				K: 79,
			},
			want: 78,
		},
		"case_4.662": {
			args: args{
				N: 827,
				K: 303,
			},
			want: 302,
		},
		"case_4.663": {
			args: args{
				N: 192,
				K: 6,
			},
			want: 5,
		},
		"case_4.664": {
			args: args{
				N: 152,
				K: 34,
			},
			want: 33,
		},
		"case_4.665": {
			args: args{
				N: 319,
				K: 44,
			},
			want: 43,
		},
		"case_4.666": {
			args: args{
				N: 717,
				K: 309,
			},
			want: 308,
		},
		"case_4.667": {
			args: args{
				N: 863,
				K: 675,
			},
			want: 674,
		},
	}
	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := bitwiseAnd(tt.args.N, tt.args.K); got != tt.want {
				t.Errorf("bitwiseAnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
