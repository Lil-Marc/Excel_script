package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
	"strings"
)
func main() {
	f, err := excelize.OpenFile("【0902新增独家书籍数据外放】-南泽29本.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	NANZE_BOOK_IDS := []int{46996, 46649, 46302, 46846, 47110, 46994, 47019, 47054, 46527, 72971, 72944, 46777, 72980, 47121, 129376, 47161, 112705, 74147, 47094, 47120, 46372, 46980, 73039, 73096, 46906, 47132, 129368, 47040, 47048, 137328, 47034, 47119, 136588,
		128206, 47113, 46453, 46794, 46712, 47059, 72967, 72937, 47025, 72981, 46603, 46722, 72979, 46791, 73063, 47073, 46738, 138293, 47003, 47062, 46518, 47020, 46774, 46688, 73078, 47135, 46885, 46948, 46962, 47053, 47042, 47066, 131092, 72959, 46779, 47096, 47174,
		47095, 46952, 72976, 46771, 73066, 137331, 46269, 46246, 46768, 46323, 73006, 112719, 46920, 73100, 73018, 47069, 72988, 112704, 73075, 72942, 46784, 73115, 47164, 112726, 47022, 112730, 47151, 72987, 129369, 138225, 73114, 72936, 47051, 72874, 72886, 72872,
		73055, 73043, 73082, 46594, 73105, 46704, 73059, 46868, 46999, 46854, 47086, 47083, 46931, 73019, 47165, 73112, 72993, 72985, 47039, 72965, 47156, 46472, 46891, 73056, 73072, 73081, 47044, 47127, 47134, 73041, 47070, 46800, 112715, 47002, 47140, 47084, 73040,
		46841, 73033, 47041, 73108, 47063, 46995, 47050, 138224, 138292, 137332, 72940, 47104, 47045, 46492, 47077, 72926, 47159, 46998, 46202, 112729, 47131, 47130, 46997, 72966, 46898, 47006, 47012, 73061, 46969, 72948, 47076, 46411, 47074, 72984, 47028, 137330, 47056,
		136587, 72970, 47043, 72933, 73051, 47112, 46986, 72996, 47032, 73011, 137334, 46627, 47021, 46993, 72974, 47162, 47146, 72934, 47049, 47142, 46803, 47036, 47035, 73113, 72931, 47139, 46825, 72994, 47166, 73003, 73002, 47126, 46830, 73067, 47004, 47065, 72982, 47079,
		73119, 47160, 112720, 47097, 47171, 73111, 73110, 73060, 47124, 73065, 47000, 73086, 47123, 47103, 46746, 73109, 46983, 46615, 47015, 46835, 46759, 47037, 73070, 73117, 47157, 46965, 73028, 73050, 46217, 73104, 46918, 46621, 72884, 73106, 73077, 112721, 73068, 72861,
		72885, 73007, 47109, 73107, 73047, 72975, 73090, 47005, 72992, 73032, 73057, 47008, 46820, 47061, 46810, 47089, 72995, 46935, 73045, 72991, 72956, 73021, 46481, 72864, 72863, 72862, 47106, 73062, 73031, 47024, 47064, 47011, 47099, 73036, 73034, 73054, 72888, 72880,
		72879, 72878, 72881, 46955, 47198, 112723, 72882, 72875, 72883, 73022, 72887, 73073, 73074, 73042, 47007, 47115, 73001, 72968, 73071, 73084, 73122, 47026, 47117, 72871, 72877, 47116, 47108, 73052, 73079, 72873, 73091, 73064, 47085, 73030, 47137, 72929, 73027, 47055,
		73058, 73035, 47038, 73038, 72962, 112725, 73026, 113155, 72889, 47013, 47138, 73049, 73103, 73046, 72997, 73048, 139988, 139985, 139976, 139986, 139990, 139981, 139984, 139989, 139974, 139983, 139982, 139978, 139980, 139979, 139987, 139977, 139975, 141830, 141831,
		141832, 141833, 141834, 141835, 141836, 141837, 141838, 141839, 141840, 141841, 141842, 141843, 141844, 141845, 141846, 141847, 141848, 141849, 141850, 141851, 141852, 141853, 141854, 141855, 141856, 141857, 141858, 141859, 141860, 141861, 141862, 141863, 141864,
		141865, 141866, 141867, 141868, 141869, 141870, 141871, 141872, 141873, 141874, 141875, 141876, 141877, 141878, 141879, 141880, 141881, 141882, 141883, 141884, 141885, 141886, 141887, 141888, 141889, 141890, 141891, 141892, 141893, 141894, 141895, 141896, 141897,
		144083, 144084, 144085, 144086, 144087, 144088, 144089, 144090, 144091, 144092, 144953, 144954, 144955, 144956, 144959, 144960, 144962, 144963, 144965, 144967, 144969, 144970, 144971, 144972, 144973, 144975, 144977, 144979, 144982, 144983, 144985, 144989,
		146388, 146387, 146391, 146393, 146386, 146389, 146401, 146397, 146398, 146402, 146396, 146394, 146400, 146392, 146390, 146404, 146399, 146385, 146395, 146410, 144957, 144961, 146471, 146472, 147335, 145737,
		147866, 147867, 148026, 148027, 148028, 148029, 148030, 148031, 148032, 148033, 148034, 148035, 148036, 148037, 148038, 148039, 148040, 148041, 148042, 148043, 148044, 148045, 148046,
		148047, 148612, 148620, 148621, 148622, 148623, 148624, 148625, 148626, 148627, 149085, 149086, 149093, 149094, 149095, 149096, 149097, 149098, 149099, 149100, 149101, 149102, 149103, 149104, 149105, 149353, 149354, 149355, 149356, 149357, 149358, 149359, 149360,
		149361, 149362, 149363, 149364, 149365, 149366, 149526, 149527, 149528, 149529, 150470, 150675,
		151112, 151108, 151104, 151116, 151113, 151114, 151110, 151109, 151107, 152129,

		152379, 152386, 152387, 152388, 152441, 152442,

		154270, 154271, 154278, 154275, 154284, 154669, 154804, 154288, 154267, 154265, 154276, 154266, 154274, 154268, 154286, 154273, 154285, 154269, 154889, 154277, 154954, 154264, 154281, 154737, 154806, 154272, 154282, 154280, 154279, 154805, 154283,
	}

	newBookIdMap := make(map[string]int, 0)
	// Get all the rows in the Sheet1.
	rows := f.GetRows("Sheet1")

	for i, row := range rows {
		if i == 0{continue}

		newBookIdMap[row[0]] = 0
	}
	newBookId := make([]string, 0)
	for _, id := range NANZE_BOOK_IDS {
		tmp := strconv.Itoa(id)
		if _, ok := newBookIdMap[tmp]; ok {
			newBookIdMap[tmp] = 1
		}
	}
	for k, val := range newBookIdMap {
		if val == 1 {
			fmt.Println("注意有重复书籍!")
			fmt.Println(k, val)
		} else if val == 0 {
			newBookId = append(newBookId, k)
		}
	}

	fmt.Println(strings.Join(newBookId, ","))
	fmt.Println("新增书籍有",len(newBookId))
}