package end


func EndsWith(original string, end string)bool {
	var index = len(original) - len(end)
	return original[index:] == end 
}
