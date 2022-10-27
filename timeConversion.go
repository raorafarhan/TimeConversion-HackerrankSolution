func getAMPM(time string) string {
	var result string

	for _, value := range time {
		switch {
		case string(value) == "P":
			result += string(value)
		case string(value) == "A":
			result += string(value)
		case string(value) == "M":
			result += string(value)
		}
	}

	return result
}

func timeConversion(s string) string {
	var oclock string
	ampm := getAMPM(s)
	mapPM := map[string]string{
		"12": "12",
		"01": "13",
		"02": "14",
		"03": "15",
		"04": "16",
		"05": "17",
		"06": "18",
		"07": "19",
		"08": "20",
		"09": "21",
		"10": "22",
		"11": "23",
	}
	mapAM := map[string]string{
		"12": "00",
		"01": "01",
		"02": "02",
		"03": "03",
		"04": "04",
		"05": "05",
		"06": "06",
		"07": "07",
		"08": "08",
		"09": "09",
		"10": "10",
		"11": "11",
	}

	arrayTime := strings.Split(s, ":")
	splitAMPM := strings.Split(arrayTime[2], ampm)
	if ampm == "PM" {
		oclock = fmt.Sprintf("%v:%v:%v", mapPM[arrayTime[0]], arrayTime[1], splitAMPM[0])
	} else {
		oclock = fmt.Sprintf("%v:%v:%v", mapAM[arrayTime[0]], arrayTime[1], splitAMPM[0])
	}

	return oclock
}
