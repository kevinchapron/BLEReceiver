package settings

type ArrayWrapper []string

func (wrap ArrayWrapper) String() string{
	var s string = "[";
	for _, elt := range([]string(wrap)){
		s+=elt+","
	}
	return s[:len(s)-1] + "]"
}
func (wrap ArrayWrapper) Has(s string) bool{
	for _, a := range([]string(wrap)){
		if(a==s){
			return true;
		}
	}
	return false;
}
