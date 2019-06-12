package main

type Multimap map[string][]string

func (multimap Multimap) Add(key, value string) {
       if len(multimap[key]) == 0 {
              multimap[key] = []string{value}
       } else {
              multimap[key] = append(multimap[key], value)
       }
}

func (multimap Multimap) Get(key string) []string {
       if multimap == nil {
              return nil
       }
       values := multimap[key]
       return values
}

func main() {
       var myMap Multimap
       myMap = make(Multimap);
       myMap.Add("黑龙江", "齐齐哈尔")
       myMap.Add("黑龙江", "哈尔滨")
       myMap.Add("黑龙江", "大庆")
       myMap.Add("辽宁", "大连")
       myMap.Add("辽宁", "沈阳")

       for key := range myMap {
              println("key:", key)
              for j := 0; j < len(myMap[key]); j++ {
                     println(myMap[key][j])
              }
       }

}
