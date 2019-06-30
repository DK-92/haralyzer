package structs

// Data holds the file(s) and other information needed for this program
type Data struct {
	HarFiles []MainLog
}

var data *Data

func GetData() *Data {
	if data == nil {
		data = new(Data)
		data.HarFiles = make([]MainLog, 0)
	}

	return data
}
