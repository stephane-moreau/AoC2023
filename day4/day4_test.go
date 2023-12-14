package day4

import (
	"fmt"
	"strings"
	"testing"
)

var (
	testInput = []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}
	dataInput = []string{
		"Card   1: 36 15 12 91 47 98 59 46 83 86 | 86 34 88  7 36 82 90 32 83 56 27 45 49 69 91 47 98 59 13 15 68 12 17 11 46",
		"Card   2: 55 82  6 50 62  4 21 96 14 39 | 11 34 14 52  7 28 13 81 10 98 88 48 21 22 96  1 93 83 85 60 76 87 57 56 62",
		"Card   3: 11 88 40  9  7 37 92 15 70 53 | 55 37 97 22 77 34 83 98  1 92 38  5 69 54 25 73 13 94 62 96 78 93 75 23 31",
		"Card   4: 83 15  5  4  3 92 86  2 91 58 | 58  2 77 32 93 64 37 82 47 28 13 89 16 24 59 76 52 50 55 15 14 68 61 91 98",
		"Card   5: 59 25  2 83 13 26 39 45 28 74 | 29 46 37 81 25 39  4  1  2 64 52 70 57 32 62 95 74  7 33 78 59 92 56  3 26",
		"Card   6: 42 94 58 49 24 26 89 77 83 66 | 88  7 54 27 89 46 57 22 97 69 65 36 13 32 11 12 31 62 35 98  4 23 79 18 85",
		"Card   7: 25  9 48 40 69 82 75 78 73 18 | 54 48 73 91 30 65 42 85 88 18 40 69 23 82 58 49 75 43 35 47 46 80 78  9 32",
		"Card   8: 21 77 46 67 93 59 81 74 85 38 | 45 37 87 59 68 85 99  9 74 21 98 52 67 81  8 27 77 83 33 18 60 11 58 25 46",
		"Card   9:  9 57 53  7 22 40  8 96 38 37 |  7 52 53 49 51 69 45 80 71  1 39 62  9 93 27 38 43 57 79 25 77 17 48 65 98",
		"Card  10: 33 54 80 15 81 44 62 41 82 55 | 43 37 18 29 17  6 91 35 70 72 97 73 75 83 74 16 40 33 19 28 56 25 32 24 52",
		"Card  11: 17 84 52 51 47 16 18  1 82 55 | 74 55 91 48 29 94 32 54 20 43 23 81 14 65 98 38 27 46 62 87 56 66 40 34 88",
		"Card  12: 64 84 36 81 75  9 68 88 54 66 |  2 63 55 97 91 67 26 15 27  6 99 19 86 24 89 25 30 37 95 33 78 28 49 44  5",
		"Card  13: 24 25 76 80  4 34 22 11 30 31 | 60 93 13 33 82 19 97 26  5 87 49 59 15 74 62 92 78  1 85 14 48 27 39 71 32",
		"Card  14: 43 21  4 55 23 41 19 88 34 27 | 52 96 61 83 39 35 44 50 63 29 84 19 75 38  7 68 20 18 94 79 45 74 53 76 55",
		"Card  15: 94 13 62 51 49 75 37 38 26 96 | 92 89 12 34 40 30 47 85 29 91 98 10 59 26 42 93 71 95  4 83 87 11 80 60 82",
		"Card  16: 54 63 61 28 53 43 69 58 32  1 | 20 88 96 89 56 64 34 74 39 76  2 30 23 27 86 19 14 49 38 94 90 55 21 45 31",
		"Card  17:  9 88 53 97 82 50 64 57 39 87 | 44 53 40 39  9 96 60 32 27 50 57 94 87 20 64 88 82 97 21 75 79  8 48 30 72",
		"Card  18: 25 34 47 32  9 40 62 76 85 99 | 62  7 40 29 78 47 10 14 68 59 32 66 34 33 84 42 99 27 70 71 76 77 20 18 37",
		"Card  19: 43 88 34  7 48 23 59 37 13 49 | 89 80 90 35 43 37 59  7 61 88 15 50 48 72 28 23 13 34 73 38  1 95 19  5  4",
		"Card  20: 72 41 70 13  4  2 12 58 89 69 | 53 99 12  5 65  4 87 52 44 28 21 26 35 55 94 27 13  2 40 39 22 89 69 72 47",
		"Card  21: 86 58 95 15 50 24 87 35 81 36 | 35  3 82 32 24 91 28 77 63 86 90 54  6  7 50 95 58 15 36 87  1 60 83 40 81",
		"Card  22: 62 19 22 91 12 75 94 97  1 88 | 58 45 75  1 97 72 63 12 52 20 85 88 62 40  9 16 10 94 50 19 93 87 15 73 91",
		"Card  23: 73 68 87 82 99 23 29 69 85 45 | 27 55  6 33  8 86 66 79 48 72  1 67 57 88 46 53 81 97 15 92  5 20 25 52 91",
		"Card  24: 48 81 34 47 64 33 40 18  1 91 | 86 46 64 49 65 22 23 33 60 75 71 87 91 21 94 79 31  3 37 89 38 90 47 10  4",
		"Card  25: 19 14 68 11 46 80  9 92 59 43 | 78 43 79 74 93 85 24 80 27 44  9 92 19 14 71 84 46 11 40 83 70 51 25 72 68",
		"Card  26: 68 72 65 43 87 97 88 47 67 34 | 87  1 57 47 62 97 79 90  6 69 68 43 17 65 46 74 31 91 67 78 98 54 32 34 29",
		"Card  27: 29 48 88 95 84 11 80 43 31 60 | 64 98 28 52 35 88 71 76 13 86 31 29 11 48 83 78 84 65 43  6 87  9 14 47  5",
		"Card  28: 36 38 72 52 98 12 79 47 75 71 | 67 25 56 45 44  5 19 11 84 96 95 22 81 50 17 46 65 63  9 69 77 18 37 39 66",
		"Card  29: 23 84 27 89 78 76 70 91  9 44 | 87 66 31 50 33 76 32 47 41 83 28  7 55 53 18  5 38 91 92 19 10 37  6 82 97",
		"Card  30: 50  4 85 77  1 46 99 37 23 35 | 26 17 11 91  8 81 40 27 90 99 21 20 68 79  5 84 61 19 86 92 36 39 16 78 73",
		"Card  31: 34 83 24 58 37 13 41 76 99 22 | 20 95 66 99 36 78 50 15 97 68 39 12  3 57 56 51 18 28 29 74 37 85 63 33 60",
		"Card  32: 21  6 59 74 76 19 41 33 97 34 | 79 77 87 60 13 59 96 72 82 78 54 98 65 51 19 22 63 91 30 15 26  9 16 40  3",
		"Card  33: 46 27 59 82 62 66 34 94 79 99 | 85 68 12 53 51 75 71 43 78 96 39 73  2 70 26 92 64 61 63  9 60 23 16 87 22",
		"Card  34: 13 85 31 30 75  2 11  3 20 28 | 23 88 62 52 18 78 56 47 54  8 61 49  6 36 73 29 46 92 17 12 72 76 97 42 21",
		"Card  35:  9 14 47 76 35  8 59 48  6 22 | 90 68 15 11 44 46 91 33  6 37 47 59 14 87 70 45 42 31 22  3 65 21 63 27  2",
		"Card  36: 66 18 69 62 12 10 90 44 36 86 | 69 39 24 90 12 86  1 18 10 52 91 76  7 44 66 85 96 62 36 58 30  8 20 56 75",
		"Card  37: 23 69 40 57 80 87 78 89 58 59 | 40 80 89 58 49 11 63 81 65  6 19 87 66 82 70 59 16 23 64 61 78 90 57 69 95",
		"Card  38: 26 47 25 29 69 90 34 14 74  5 |  6 14 64  8  2 19 15 80 92 56 67 16 38 20 65 52 53 35 68 21 51 42 99 72 94",
		"Card  39: 59 45 21 52 28 25 77 88 61 49 |  6 81 16 31 45 61 99 35 32  8 26 89 63 69  7 37 64 84 14 11 80 21 67 59 44",
		"Card  40:  8 70 30 10 52 80 87 28 47 51 | 90 17 15 70 52 60 73 63  5 72 30 62 34 47  3 51 16  6 91 23 28 18 99 48 88",
		"Card  41: 19 15 95  8 91 84 92 70 47 69 | 95  7 80 69 30 91 68 94 36 12 70 14 64  5  8 84 92 71 97 19 99 15 47 29 86",
		"Card  42: 41 28 42 53 83 36 26 55 63 45 | 19 99 86 70 39  8 92 42 63 98 69 54 87 26 83 84 20 14 28 94 11 10 57 43 44",
		"Card  43: 49 52 29 81 12 48 44 62 84 43 | 37 80 18 20 26 64 21 41 56 90 15 70 84  6 83 30 67 72 19  1 95 13 24 58 54",
		"Card  44: 76 36 69 80  6 20 38 84 59 12 | 86  3 53 71 59 27 94 92 50 51 89 55 14 47 20 60 58 68 11  6 64 65 25 57 22",
		"Card  45:  5 74 14 24 90 28 35 78 92 12 | 25 80 93 84 97  8 71 32 91 15 73 72 58 38 95  1 56 52 81 79 45 96 49 65 34",
		"Card  46: 73 64 82 24 31 81 18 79  2 96 | 79 36 38 15 51 29 28 80 95 64  4 41 16 99  3 67 96 56  2 25 34 35 97 92 82",
		"Card  47: 65 26 54 44 96 58 63 83 45 76 | 13  9 48 22 15 64 56 71 26 65 27  8 70 50 85 66 98 82 54 87 95 33  2 16 53",
		"Card  48:  7  3  2 80 64 51 17  5 97 28 | 25 73 98 87 80 70  8  9 59 15 14 60 21 86 81 66 90 67 42 52  5 57 31 53 34",
		"Card  49: 67 63 26 16 71  1 29 38 90 93 | 83 62 60 15 87 84 48 10 96 70 34 36  2 45 25 31 52  8 14 64 56 88 95 33 38",
		"Card  50: 44 39 53 49 46 77 23 28  7 57 | 54 43 68 83 40 66  6 73 60  5 32 22 97 34  2 16 90 15 65 86 99 89 24 11  8",
		"Card  51: 33 17 53 74 49 79  6 26 38 83 | 54 63  1 82 99 58 78 52 22 35 29 86 25 42 43 36 24 95 55 96 56 45  9 89 77",
		"Card  52: 65 45 73 78 71 49 35 42 34 16 | 71 35 45  3 16 57 32 25 48 34 87 93 49 12 20 31 95  1 14 64 40 74 78 89 61",
		"Card  53: 52 87 86  2 70 78 10 33 67 74 | 70 33 49 29 67 55 54 65 80 14 16 52  7 59 66 91  5 74 31 85 84 75 48 28 87",
		"Card  54: 37 41 81 57  2 78 53 84  9 94 | 39 78 57  4 84  9 94 58 97 47 85 21 65 81 99 53 37 60  2 50 80 28 75 43 41",
		"Card  55: 22 38 85 65 19 82 21 66 32 50 | 39 75 30 21 17 98 50 80 81 28 58 43  7 76  3 66 82 78 63 11 37 64 79 51 27",
		"Card  56:  8  4 86 70 65 25 91 12 45 26 |  8 41 26 62 29 34 12 75 86 91 37 66 70 14  4 52 30 16 25 58 43 50  1 71 63",
		"Card  57: 59 68 40 75 58 37 82 77 67 66 | 86 75  2 40 78 62 28 95 77 10 50 54 73 64 89 37 20  5 61 66 68 82 58 29 43",
		"Card  58: 40 26 86 55 90 89 66 71 11  5 |  5 30 15 94 81 82 75 43 71 86 39 11 54 42 44 73 79 40 72 26 65 87 66 61 98",
		"Card  59: 15 97  7 25  9 19 84 14 49 58 | 59 76  9 25 52 15 18 49 97 19  8 58 41 57  7 38 14 91 10 65 84 50 33 40 64",
		"Card  60: 42 81 61 29 94 85 68 43 75 91 | 33 52 26 86 17 72 13 69 68 41 51 75 80 77 91 81 74 18 88 90  4 76 46 25 42",
		"Card  61: 69 64 72 32 43 35 86 81  3 27 | 33 18 22 68 99 69 66 59 83 72 44 31 84 78 49 81 58 74 88  3  6 53 92 55 42",
		"Card  62: 71 17 60  7 84 67 37 19 13 35 | 28 37 63 48 71  4 60 88 46 57 67 94 74 45 14 66 19 87 17 84 54  7 13 75 35",
		"Card  63: 96  5 62 72 74 25 80 85 16 95 | 69 73 79 47 87 36 16 51  4 53 93 43 19 72 27 78 56 21 77 61 46 76 89 60 88",
		"Card  64: 52  2 74 71 14 93 25 80 99 62 | 12 14 97 39 49 56 64 33 42 48 68 57 44  3 29 67  8 95  1 83 91 13 11 78 93",
		"Card  65: 97 19  8 47 32 17 28 64 11 85 |  6 12 32 22 97 28 49 95 55 62 85 78 72 36  4 71 33 47 57 87 39 19 53 64 54",
		"Card  66: 80 58  2 35 41 26 25  5  9 79 | 13 79 66 89 33 99 49 41 74  5 38 25 23 62 34  2 47 84 88 10 15 39 16 26 97",
		"Card  67: 54 13 75 14 18  7 21 60 50 28 |  7 79  2 90 86 44 28 13 76 77 97 36 11 20 80 29 38 52 98 19 60 74 51 43 91",
		"Card  68: 78 40  9 24 10 23  1 76  5 21 | 62 76  1  2 88 56 39 69  5 81 19 72 92  3 67 91 53 84 99  6 86 94 74 48 20",
		"Card  69: 34  7  4 96 69 61 67 57 92 22 | 93 20 78 95 73 29 50 26 86 42 89 35 98 65 54 13 39 34 91 40 14 79 45 47 57",
		"Card  70: 95 13 74 50 33 78 15 63  2  6 | 46 62 50 37  4 63 67 26 18 87 92 97 83 96  2 85 45 53 16 65 38 41 40 44 86",
		"Card  71: 89 65 75 55  1 26 45 72 16 12 | 24 85 73 66 70  8 25 93 95 79 84 11 19 18 30 68 45 37 21  4 38 12 28 87  1",
		"Card  72:  8 60 38 76 10 18 79 58 34 68 |  6 46  1 44 93 61 26 98 41  2 90 45 94 91  4 25  7 17 14 63 36 75 67 70 60",
		"Card  73: 13 65 45 41 74 67 60 36 72 42 | 66 76 77 14 38  7  2 50 81 56 78 48 79 69 30 83 34 43 54 87 28 10 82 86 47",
		"Card  74: 95 69 68 20 66 54 72 36 76 85 | 65 29 86 57 47 59 17 74 58 26 81 25 39 27 97 38 70 61 88 41 48 67 40 93  3",
		"Card  75: 38 67 97 12 45 34 93 20 75 71 | 26 93 84 73 12 97 22 63 59 75 33 62 69 68 80 41 64 45 38 55 95 67 34 35 76",
		"Card  76: 70 38 31 94 97 36 23 39  5 29 | 55 41 79 82 10 68 92 43 23  7 84 54 30  4 60 16 19 80 49 73 83  8 22 69 47",
		"Card  77: 87 97 43 12 57 15 64 52 58 39 | 70 26 56 99 44 15 12 85 48 74 80 31 76 23 16 35 67  7 79 52 11 49 54 28 69",
		"Card  78: 41 32 48 23 95 64 76 39 80 26 | 73 38 12 99 27 97 32 23 48 21 76 91 49 59 87 67 74  4 43 56 37 95 80 19 52",
		"Card  79: 86 77 79 59 22 10 62 50 18 26 | 50 62 26 71 24 25 22 36 81 31 79 87 76 18 77 86 10 53 95  6 40 43  9 59 64",
		"Card  80: 95 78 52 48 49 27 96 67 23 24 | 96  5 24 45 67 92 49  7  1 94 66 18 86 48 27  4  9 39 83 99 81 23 52 79 75",
		"Card  81: 61 98 74 65  6 64 68 25 90 56 | 35 76 81 65 44 98 37 99 64 67 50 73 39 32 74 20 25 14 30  2 59 96 79  1 36",
		"Card  82: 79 73 18 90 99 26 13 14 96 91 | 91 14 19 26 72  4 96 33 89  9 99 27 29 12 37 79 18 57 73 17 11 63 90 74 94",
		"Card  83: 24 17 35 53  7  1 20 52 21  2 | 83 13 35 28 98 40 53 20 62 86 17 63 22 19 10 89 97 71  2 24 29 33 95 85 72",
		"Card  84: 74 99 95 53 57 97 27 24 88 56 |  5 99 95 27 91 93 88 58  8 20 12 85 86 14 67 21 53 56 46 39 57 38 74 77 24",
		"Card  85: 58 69 43 81 37 89 16 49 36 26 | 58 43 16 80 12 17 62 49 23 68 90 51 22 84 40 44 52 10  1 81 11 26 30 36 79",
		"Card  86: 33 45  4 20 58 96 77 19 39 13 | 54 81 77  6 30 19 96 44 91 68 98 51 85 62  8 33 26 61 55 87 97 35 32  9 59",
		"Card  87: 95 74 24 89 96 94 85 38 10 12 |  2 69 48 19 66 47 20 27 67 91 95 53 68 32 29 10 31 11  4 59 26 92 97 99  3",
		"Card  88: 53 11 85 19 39 78 18  7 97 29 |  8 46 73 49 80 85 33 65 32 88 51 26 94 19 68 57 29 28 90 35 11 53 40 12 77",
		"Card  89: 77 93 40 33 81 50 30 51 65 61 | 87 34 58 29 66  7 76 54 84 30 37 33 81 48 26 68 90 62 22  4 24 67 23 80  1",
		"Card  90: 71 91 65 19 12 87  2 99 33 74 | 24  5 53 20 88 86 60 23  6 73 31 36  7 97 14  4 78 34 82 66 22 56 89 48 28",
		"Card  91: 55 84 61 56 88 35 90 49  2 89 | 21 39 46 81 96 85 55 77 71 64 34 78  5  9 18 49 82 87 93 41 27 72 50 75 88",
		"Card  92: 22 78  1 56 28 40 69 44 68 13 |  7 98 99 10  5 61 54 15 70 96 43 13  1 32 83 34 12 38 37 50 57 35 58 16  2",
		"Card  93: 31 62 64 13  6 40 51  8 37 96 | 90 88 11  9 35 70 30 76 61 73 81 50 98 93 82 26 99 41 34 25 59  8 18 21 54",
		"Card  94: 74 26 80 66 17 70 44 99 96 97 |  3 79 15 41 36 84 72 86 32 83 91 64 34 43 31 49 46 24 95 47 56 62 33 16 55",
		"Card  95: 19 66 59 91 25 72 53 14 17 44 | 51  1 37 43 91 82 55 57 33  2  3 53 66 52 96 19 49 75 31 71 32 83 59 74 61",
		"Card  96: 69 39 12 87 82 24 42  4 96 23 | 57 60 18 24 13 21 39 44 87 99 14 72 92 82 35 73 37 56 65 55 52 42 33 59 69",
		"Card  97: 79 98 61 35 78 28 64 34 48 99 | 22 99 34 78 79 33 38 94 89 61 26 20 91 64  9 69 75 28 48 35 98  2  4 10 37",
		"Card  98: 15 29 96 93  9 22 45 74 24 63 | 20 30 44 26 29  4 63 75 32  9  3 15 45 24 93 21 96 28 82 94  6 74 46 22 79",
		"Card  99:  8 62 79 47  4 13 46 93  2 52 | 91 55 81 37 70 49 44 94 22 84 74 80 97 18 36 69 17 14 23 57 86 12  1 30 64",
		"Card 100: 37 14 18 88 54  1 28 78 58 84 |  4 49  1 88 18 66 90 65 28 64 54 29  6 14 34 98 25 63 59  5 72 37 75 95 41",
		"Card 101: 51 14  5 84 38 49 69 29  7 52 | 31 96  8  2 55 66 41 79 45 93  5 26 42 85 27 61 91 78 49 87 84 40 39 60 44",
		"Card 102: 16 35 55 56 98 29 25 31 67 87 | 41 11 61 96 16 93 75 56 51 43 54 92 84 55 36 34 35 31 14 59  3 15 73 12 17",
		"Card 103: 78 58 45 13 80 16 72 20 36 67 | 84 67 72 46 77 14 27 35 50 45 47 21 78 28 16 24 36 55 13 33 22 62 80 52 11",
		"Card 104: 25 85 18 31 65 78 34 91 66  3 | 85 87 71 53 17 98 72  2 61 25 94 33 89 34  3 56 42 18 99 91 90  4 40 20 50",
		"Card 105: 54 30 50 19  8 90 56 37 58 29 |  4 60 17 15 36 95 42 13 92 77 53 61 57 94 98 93 40  5 68 33 23 47 72 79 83",
		"Card 106: 34 52 66 44 57 27 19  4 26 71 | 27 83 96 33 66 47 64 14 80 77 41 34 57  4 39 71 26 50 52 29 79  5 62 19 44",
		"Card 107: 23 97 80 95 62 99  6 52 60 53 | 70 19 99 51 95 72 96 94 11  5 93 21 16 67 73 47 68  1 17 77  2 30 56 74 83",
		"Card 108:  1  9 71 31 82 17 16 95 63 18 | 73 12 70  1 28 11 32 86 54 64 20  3  9 26 31 17 82 68 13 16 93 39 35 63 95",
		"Card 109:  6 90 24 69 70 77 47 58 28  9 | 72 58  2 25 60 80 33  6 30 85 40 11 41 90 82 28 62 42 36 95 77 51 55  9 70",
		"Card 110: 23 29 78 34 31 53 83 13 73 47 | 27 15  7  6 92 50  9 56 61 26 73  4 30  5 12 28 80 58  2 43 39 36 63 90 64",
		"Card 111:  4 50 17 67 60 30 22 29 73 96 | 88 73 89  5 53 92 21 36 60  1 27 67 50 91 74 30 11 38 96 18 83 99 20 13 82",
		"Card 112: 48 53 95 44 98 10 82 70 15  4 | 27 39 52 42 87 34 43 50 19  3 78 90 57 96 56 54 63 24 85 83 21 37 59 41 31",
		"Card 113: 14 39 90 84 80 13 37 50 55 18 | 73  1 71 30 13 22 98  5 44 41 68 62 35 61  7 15 23 29 45 63 25 50 16 26 83",
		"Card 114: 74 38 43 34 58 61 12 49 44 26 | 67 19 81 87  2  9 88 20 71 84 53 62 15 96 92 58 36 52  7 90 86 35 46 70 55",
		"Card 115: 55 25 96 21 46 34 59 13 78 23 | 63 76 29 75 14 26 53 52 88  7 83 93  1 27  4 64 37 31 65 79 15 18 38 62 19",
		"Card 116:  6 28 48 76 43 64  8 90 70 88 | 55 97 92 17 52 91 39 12 23 83 96 66 99 37 58 10 46 81 35 75 31 61 95 80 63",
		"Card 117: 60 59 75  9 26 50 90 87 84 15 | 54 56 76 36 94 57 11  6 18 66  4 32 91 82 97 88 46 16 78 84 14 22 61 24 33",
		"Card 118: 54 18 66 34 89 85 92 74 22 14 | 45 21 98  3 32 88 64 29  2 83 15 78 36 44 76 60 48 87 59 96 33 49  7 84 82",
		"Card 119: 58 95 64 60 30 23 83 46 92  4 |  1 59 34 11 41 42 51  3 50 96 24 86 71 20 62 90  7  2 27 53 29 77 37 44 40",
		"Card 120: 48 38 27 97 81  8  1 89 67 29 |  4  1  7  8  3 68 42 21 37 29 83 60 89 46 16 58 98 67 66 44 54 32 96 73 45",
		"Card 121: 39 27 64 60 48 75  2 56 23  4 | 22  4 38 60 37 39 28 54 58 16 26 95 51 12 57  2 23 56 59 64 63 88 73 27 49",
		"Card 122: 53 34 40 31 58 63 75 84 96  5 | 53 49 68 85 16 66 20 24 41 29 94 50 96 77 46 88 40 91 34 69 78  8 25  3 43",
		"Card 123: 33 84 11  4 53 66 10 97 92 48 | 23 55 65 86 90 84 17 15 48 64 33  4 11 87 53 10 93 51 60 61 97 52 27 35 36",
		"Card 124: 25 46 23 27 51 69 38  1  8 71 | 42 13 27 38 69 63 54 33  1 62 97 57  7 71 41  9 35 52 44 87 32 51 65 84 64",
		"Card 125: 59 40 61 60 47 81 90 84 54 75 | 49  8 46 57 56 39 34 65  9 12 77  4  1 15 42 78 55 38 20 75 79 96 35 13 88",
		"Card 126: 15 60 67 48 47 58 34 92 64 94 | 19 84 87 67 48 20 56 17 62 22 31 11 83 24 61 47 59 34 97 64 10 42 92 13 60",
		"Card 127: 29 18 32 86 89 45 13 44 85  3 |  7 89  5 32 69 83 86 49 13 29 94 44  9  3 48 41 96 74 85 33  4 45 18  2 81",
		"Card 128: 61 85 27 44 92 26 60 52 37 50 | 78  8 80 18 21 76 84 34 42 12 99 74 98 10 65 59 13 90 82 77 89 23 70 22 43",
		"Card 129: 69 45 65 82  1  7 95 37 10 21 | 45 19 50 34 70 31 18 84 67 53 77 85 38 42 63 35 97 54 58  4  9 93 30 86 56",
		"Card 130: 66 33 46 27 85 72 77  4  3 14 | 64 55 93 17 49 82 22 77 28 30 12 81  4 35 65 79 95 84 66 75 83 42 74 14 80",
		"Card 131: 95 39 83 84 90 44 91 22 63 55 |  2 51 62 82 41 96 10  9 29 28 91 86 44 59 43 26 13 81  1 79 20 92 15 27 87",
		"Card 132: 38 60 36 70 10 68 11 20 29  7 | 50 75 33 92  8 30 77 42 80 54 14 74 59 82 45 86 71 12 67 15 99 84 51 55 22",
		"Card 133: 47 90 12 18 72 88 79 67 15 13 | 37 21 50 73 60 83 84 63 38 45 78 90 95 66 54  6 12 15 48 36 64  7 43 34 75",
		"Card 134: 83 96 51 15 20 74 82 98 49 19 | 11 67 68 99 43 98 84 90 95 65 58 88 10 23 12 75 32 33 29 91 56 93 46 86 31",
		"Card 135: 29 23 40 35 24 88 19 90 31 21 |  8 98 28 75 13  6 71 61 42 38 53 33 30 36 22 58 93 78 80 68 55 90 81 48 45",
		"Card 136:  9 18 73 47 43 19  1 27 59 86 | 79 56  5 28 36  3 45 53 83 66 78 42 99 75 81 87 91 14 11 41 49 70 48 80 10",
		"Card 137: 30 83 66 15  6 94 82 99 42 20 | 88 84 96 92 28 38 81 64 33 58 70 72 73 90 24 46 16 50 79 97  2 19 23 26 34",
		"Card 138: 21 14 47 43 99 42 98 41 82 12 | 73 47 14 97 43 40 82 57 99 56 54 55 18 86  9  1 21 24 41 23 12 26 98 42 33",
		"Card 139: 25  9 96  2 57 78 94 71 68 74 | 78  9 96 88 80 76 53 25 67 23 68 97 40 94  2 41 39 10 28 24 37 57 27 74 71",
		"Card 140: 93 77 52 60 43 79 32  1 39 84 | 93 37 28 77 38 47 52  5 59 31 16 46 43 33 32 79 55  9  1 84 27  7 39 60 53",
		"Card 141: 93 95 23 22 60 18 48 58 85 73 | 39  2  6 67 85 75 22 58 86 59 40 23 88 38 93 70 18 55 48 97 68 28 94  3 60",
		"Card 142: 24 69 66  1 70 46 76 95 13 74 | 24 74  5 69 66  2 52 70 39 76 99 21 72 20 96 53 57 78 41 46 95 13  3  1 31",
		"Card 143: 73 42 95 58 34 97 23 67 83 47 | 34 21 23 48 58 10 67  8  1 19 84  9 36 45 28 70 17 83 60 66 44 50 69 27 73",
		"Card 144: 73 57 22 89 52 43 96 66 39 32 | 35 38 47 25 46 86 18 79 91 72 56 84  3 30 37 22 51 31  8 66 42 57 44 41 15",
		"Card 145: 62 59 49 34 20 58 53 36 23 77 | 82 56 59 17 70 54  3  8 91 63 45 77 47 60 96 20 26 16 12 83  2 18 57 38 19",
		"Card 146: 79 26  2 71 31 74 76 34 46 10 | 98  1 71  4 34 31 12 45 33 76 30 75 26 49 18 40 37 78 36 59 79 85  7 46 60",
		"Card 147: 81 64 13  5 61 35  4 46 32 53 | 90 83 62 94 74 12  9 85 80 61 45 34  8 77 30 95 64  1 13 38 35 23 42 25 81",
		"Card 148: 40 73 39 13 19  5 62 86 89 11 | 38 83 63 10 23 74 41 56 97 51 92 53  6 85 70 13 94 64 15 62 47 77 99  4 24",
		"Card 149: 25 82 20  5 69 17 92 37  8 35 | 90 48 17 42 60 24 46 92 82 10  8 25 15 84 49 69 20 37  5  3 45 80 31 35 36",
		"Card 150: 81 87  3 14  1 46 21 86 50 15 | 70 28 98 56 30  3 50 44 76 49 81 21 46 91 14 96 48 95 93 20 61 26 34 15 47",
		"Card 151: 86 73  7 35 71 23 61  4 47 15 |  2  5  4 34 29 71 12 31 81 36 80 35 27 23  8 42 88 47 52 99  7 67 61 91 16",
		"Card 152: 57 12 20 38 86 39 91 30 52 64 | 63 10  6 89 56 57 22 30 81 44 91 75 18 86 65 60 84 12 52 39  8 64 25 38 11",
		"Card 153: 82 41 45 17 65 61 36 92  2 22 | 57  3 75 39 56 62 24  8 81 12 79  7 53 18 65 37 72 68 92 45 42 84 27 73 64",
		"Card 154: 12 51 52 96 31 79 72 59 43 47 | 97 43 88 63 65 12 41 54 91 99 10  2 34 59 36 96 18  1 46 23 55 74 84 51 64",
		"Card 155: 36 53 25 77 50 34 87 59 69 16 | 36 30 50  5 82 34 71 91 62 16 65 85 58 31 42 56 40 25 24 61 73 21 97 59 20",
		"Card 156: 79 19 23 40 96 59 37 27 38 92 | 54 46 77 26 76 81 73 16  8 98 53 28 50 14 67 90 52 35 22 12 97 30 44 47 78",
		"Card 157: 32 91 99  5 29 98 81 31 20 95 | 14 48 69 79 87 29  2 62 41 11 68 95 72  5 94 92 36 81 33 64 97 55 43  3 89",
		"Card 158: 72 45 56 83 37 79 59 52  2 13 | 39 44 77 99 23  4 83 35 38 84 51 52 37 36 57 96 81 65 25 71 87 47 80  9 32",
		"Card 159: 73 65 12 70 55 25 83 67  6 74 | 91 34 70 21 42 46 83 47 77 13 19 82  2 40 89 81 27  3 11 80 69 76 79 99 54",
		"Card 160: 75 57 60 32 66 87 31 92 45 42 |  5  3 76 97 99  7 89 22  2 69 16 81 65 73 88 98 58 40 33 59 82 72 75 50 66",
		"Card 161: 18 36 95 99 57 20 33 84 21 63 | 61 58 27 38  3 55 69 68 48 86 24 42 63 71 87 23 17 30 52 59  1 12  6 28 19",
		"Card 162: 87 52 95 78 63 17 15 33 21 70 | 38  9 64 86  8 91 66 96 72 71 79 94 28 82 60 45 39 16 14 89 75 99 90 20 10",
		"Card 163: 99 57 56  2 37 85 29 60 88 64 | 77 11 47 80 67 76 62 98 50 70 21 45 95 18 68 79 14 90 46 43 19 27 25  7 71",
		"Card 164: 42  7 47 27 32 38 41 25 63 19 | 47  7 57 88 63 70 19 40 41 27 80 18 25 46 38 50 32 83 90 42 97 22 28 77  4",
		"Card 165: 61 24 38 72 20 92 80 70  5 49 | 79 24 61 25 82 38 20 43 33 64 72 51 49 91 12 90 70 86 27 56 80 88 97 92 17",
		"Card 166: 28 45 17 96 27 22 75 67  7 72 | 66 80 11 87 28  7 27 22 97 94 45 83  8 72 67 14 96 21 17 34 75 73 99 82 58",
		"Card 167: 95 81 44 93 30 26 75 68 19 43 | 15 99 66 57 16 76 68 95 25 61 92 44 97 14 72 91 93 64  6 58 51 78 60 63  2",
		"Card 168: 14 27 20 26 25 55 66 51 47  4 | 51 26 79 94  6 69 47 11 18  8 25 71  4 37 84 75 20 46 66 14 27 55 30 15 49",
		"Card 169: 40  7 89 73 71 57 81 93  2 34 | 46 50 12 44 40 79 54 82 95  3 15 43 78 48 32 26 51 53 76 58 38 62  1 17 28",
		"Card 170: 56 86 42 26 79 40  5 52 46 93 | 94 72  3 56 25 42 57 16  5 26 93 52 61 40 86 79 46 39 78 62 36  9 35 99 83",
		"Card 171: 56 17 25 24 70 33  8 83 14 35 | 68  7 50 54 74 58 42 47 46 87  3 48 62 78 61 91  1 21 36  5  6  4 38 65 92",
		"Card 172: 80 96 68 49 79 24 57 72 26 50 | 81 38 72 80 68 92 66 13 28 57 26 10 49 78 24  9 75 96 54  2 79 55 77 50 40",
		"Card 173: 42 65 14 73 74 27 95  3 61 39 | 74 99 13  2 19  9 94  4 22 61 86 95 18 57 53  6 90 77 79 10 82 44 67 27  5",
		"Card 174: 25 27 47 41 65 34 44 40 49 92 | 20 30 24 46 82 31 78 28 99 94  9 87 73 34 14 79 39 86 21 37 15 80 64  7 13",
		"Card 175: 87  1 57 33  2 42 29 63 54 17 | 55 69 72 15  6 80  5 38 91 37 20 46 33 49 57 11 88 89 94 87 35 54 63 17 28",
		"Card 176: 16  5 15 73 98 30 74  2 72 94 |  4 66 56 84 72  5 44 49 63 87 76 36 25 85 35 78 28 58 42 82 38 12 53 97 51",
		"Card 177: 75  1 50 87 26 58 93 13 24 80 | 88 89 62 42 63 12  4 73 76 68 81 58 95 43 78 18 35 83 57 91 44 22 38 14 37",
		"Card 178: 17  3 42  4  7 27 35  9 88 72 | 91 64 54 46  1 34 96 39 75 69 36 17 79 74 89 42 95  7 80 82  3 37  9 55 88",
		"Card 179: 13 70 65 40 45  8 62  1 19 64 | 79 19 82 67 56 26 61 74  5 94 37 71  2 16 28 22 58 24 15 83 21 68 12 59 96",
		"Card 180: 17 29 67 85 57 54 73 13 86 48 | 31 45 92 66 89 82 17 95 75  4 65 10 28 41  3 79 68 51 34 72 29 52 21  7 46",
		"Card 181: 54 49 76 95 15 40 77 70 86 11 | 40 35 75 87 70 42 41 66 67 77 45 68 25 32 98 38  6 92 73 89 28 74 99 21 59",
		"Card 182: 41  7 61 39 66  9 33 32 18 11 | 46 35 64 47 70 68 31 84  3 24 63  2 50 28 30 17 11 59 79 71 10 54 34  9 81",
		"Card 183: 76 49 64  3  8 95 15 99  6 28 | 63 19 16 71 21 40 66 67 79 94 85  5 54 17 57 87 11 36 92 30 72 33 73 55 64",
		"Card 184: 30 37 90 81 56 40 16 26 82 35 | 88 25 61 70  7 38 34 43  8 53 22 63 23 92 33 50 93 55 45 13  9 83 79 17 36",
		"Card 185: 99 59  9 50 18 69 96 76 29 55 | 25 98 46 13 55 72 66 52 73 18 19 47 41 76 96 36  9 99  2 29 87 59 21 69 50",
		"Card 186: 32 29 34 31 53 41 10 56 69 86 | 51 54 41 82 19 56 67 10 76 58 86 69 71 28 95 61 32 34 53 31 73 68 79 29 90",
		"Card 187: 93 21 13 65 14 49 30 43 44 69 | 43 46 13 69 49 65 66 14 93 82 81 48 42 21 50 24 62 30 44 63 88 91  5  8 34",
		"Card 188: 63 45 43 81 69  7 89 93 52 24 | 73 51 70  7 37 92  4 87 96 49 83 72 62 36 63 81 50 40 66 25 93 79 97 76 33",
		"Card 189: 78 65 81  3 44 31  2 74 40 87 | 91  3 68 74 31 64 44 75  2 24 56 84 54 71 40 33 19 14 76 65 62 78 66 87 81",
		"Card 190: 99 20 62 48 40 37 81 78 87 75 | 99 84 78 48 27 87 66 72 37 46 75 22 61 74 80 62 65 82 29  7 40 77 81 20 58",
		"Card 191: 90 11 94 98 46 17 45 66 85 22 | 15 90 13  2 10 48 53 39 65 86  5 85 43 66 60 24 46 18 32 29 19 64 14 17 27",
		"Card 192: 66 21 58 45 98 57 82 10 64 73 |  4 64 82 81 57 71  2 30 58 97 89 63 46 73 45 98 66 21 83 40 75 85 76 10 12",
		"Card 193: 74 98 16 40 18 48 56 41 37 71 | 78 89 71 93 26 57 32 18 43 19 42 56 30 11 92 50 58 40 28 46 98 37 65 90 48",
		"Card 194: 54 33  3 28 15 66 43 64 62 42 | 40 90 34 89 73 88 48 45 65 58 24 13 63 16 76 84 87 23 35 53  4  9 97 20 17",
		"Card 195: 91 43 85 65 25 10 38 29  7 97 | 76 83 74 98 45 84 90 78 64 28 59 62 49 94 86 36 71 15 35 37 12 46 29 25 44",
		"Card 196: 58  4 94 64 85 68 25 19 40 15 |  4 82 86 62 25 69 88 94 61 20 75 19 64 24 70 15 12 92 14  7 55 84 68 34 85",
		"Card 197: 37 78 70 73 54 25 29 59 61 99 | 22 86 69 26 97 57  5 45 56 14 18 89 65 55 40 30 77 62 36 19 63 21 60 85 44",
		"Card 198: 79 91 51 64 32 46 78  9 58 69 | 23 22 51 41 99 20 64 84  8 68 83 58 91 46 39 55  9 70 50 75 89 32 44 94 49",
		"Card 199: 69 84 47 50 43 62 35 73  2 94 | 35 92 75 72 96  6 23 94  2 32 16 55  9 58 56 17 71 50 39 47 87 34 98 88 86",
		"Card 200: 25 82 75 18 58 15 64 51 96 19 | 45 84 20 11 62 15 82 16 51 26 14 27 58 43 81  1  3 86 33  9 66 35 37  2 98",
		"Card 201:  4 60 21 68 93 95 13  2 73 83 |  1 73 74 46 89 50 15 42 10 29 45 21 41 90 78 70 27 33 58 99 81 38 83 25 31",
		"Card 202: 70 49 77 13 46 79 38 48 94 73 | 64 86 89 23 69 76 60 24 16 81 54 10 22 68 97  5 33 90 74 38 88 17 37 49 93",
		"Card 203: 25 91 76 63 64 34 50 83 56  1 | 33 19 13 89 61 21 22 55  5  4 27 37 84 42 75 48  2 71 25 35 23 67 46 65 72",
		"Card 204: 35 58 42 71 86 19 52 39 96 10 | 15 72  7  6  1 25 74 30 61 59  4 62 67 34 91 20 78 77 95 55 68 32 82 93 69",
		"Card 205:  3 11 85 48 68 96 94 25 78 44 | 87 54 73 32 23 58 97 52 79 74 43 62 57 14 99 34 66 61 64 75 33 93 10 19 92",
	}
)

func compute(lines []string) {
	score := 0
	for _, line := range lines {
		parts := strings.Split(line, "|")

		selection := parts[1]
		selected := make(map[string]bool)
		for _, val := range strings.Split(selection, " ") {
			value := strings.TrimSpace(val)
			if value != "" {
				selected[value] = true
			}
		}

		picks := strings.Split(parts[0], ":")[1]
		selScore := 0
		for _, val := range strings.Split(picks, " ") {
			value := strings.TrimSpace(val)
			if value != "" && selected[value] {
				if selScore == 0 {
					selScore = 1
				} else {
					selScore *= 2
				}
			}
		}

		score += selScore
	}
	fmt.Printf("Score card is %d\n", score)
}

func TestDay4Phase1(t *testing.T) {
	compute(testInput)
	compute(dataInput)
}

type cardNumAndValue struct {
	num       int
	picks     string
	selection string
}

func computeLineNumber(lines []string) {

	cards := make(map[int]*cardNumAndValue)
	for i, line := range lines {
		parts := strings.Split(line, "|")
		cardAndPicks := strings.Split(parts[0], ":")
		cards[i] = &cardNumAndValue{
			num:       1,
			picks:     cardAndPicks[1],
			selection: parts[1],
		}
	}

	for i := 0; i < len(cards); i++ {
		card := cards[i]

		selected := make(map[string]bool)
		for _, val := range strings.Split(card.selection, " ") {
			value := strings.TrimSpace(val)
			if value != "" {
				selected[value] = true
			}
		}

		score := 0
		for _, val := range strings.Split(card.picks, " ") {
			value := strings.TrimSpace(val)
			if value != "" && selected[value] {
				score++
			}
		}

		for j := 0; j < score; j++ {
			cards[i+j+1].num += cards[i].num
		}
	}

	cardNum := 0
	for i := 0; i < len(cards); i++ {
		cardNum += cards[i].num
	}
	fmt.Printf("Num card is %d\n", cardNum)
}

func TestDay4Phase2(t *testing.T) {
	computeLineNumber(testInput)
	computeLineNumber(dataInput)
}
