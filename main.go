package main

import (
	// "errors"
	"encoding/json"
	"fmt"
	"math"

	// "strconv"

	// "strconv"

	// "reflect"

	"strings"
	// "strings"
	// "golang.org/x/text/cases"
	// "math/rand"
	// "strings"
	// "time"
)

// func main() {
// 	var firstName string = "john"
// 	lastName := "wick"

// 	fmt.Printf("Hello %s %s !!\n" , firstName, lastName)
// // fmt.Println("Hellow World!!!")
// }

// func main() {
// 	fruits := [4]string{"banana", "mango", "apple", "grape"}

// 	fmt.Println(len(fruits))
// 	fmt.Println(fruits)
// }

// func main() {
// 	numbers := [...]int{1, 12, 43, 14, 98, 101}

// 	fmt.Println(len(numbers))
// 	fmt.Println(numbers)
// }

// func main() {
// 	var char [...]string

// 	char[0] = "test"

// 	fmt.Println(char)
// }

// func main() {
// 	chars := []string{"lorem", "ipsum", "dolor", "sit", "amet"}

// 	chars = append(chars, "bogor")

// 	for _, char := range chars {
// 		fmt.Println("char : ", char)
// 	}
// }

// func main() {
// 	var chicken map[string]int
// 	chicken = map[string]int{}

// 	chicken["januari"] = 10
// 	chicken["februari"] = 18

// 	delete(chicken, "februari")

// 	fmt.Println("Januari : ", chicken["januari"])
// 	fmt.Println(chicken)

// 	var chickens = []map[string]string{
// 		{"name": "blue chick", "age": "1"},
// 		{"name": "red chick", "age": "1"},
// 		{"name": "green chick", "age": "1"},
// 	}

// 	for _, chicken := range chickens {
// 		fmt.Println(chicken["name"], chicken["age"])
// 	}
// }

// func main() {
// 	var names = []string{"John", "Wick"}

// 	printMessage("Hi", names)
// }

// func printMessage(message string, arr []string) {
// 	var nameString = strings.Join(arr, " ")
// 	fmt.Println(nameString)
// }

// func main() {
// 	rand.Seed(time.Now().Unix())
// 	var randomValue int

// 	randomValue = randomWithRange(2, 10)
// 	fmt.Println("random number: ", randomValue)
// 	randomValue = randomWithRange(2, 10)
// 	fmt.Println("random number: ", randomValue)
// 	randomValue = randomWithRange(2, 10)
// 	fmt.Println("random number: ", randomValue)
// }

// func randomWithRange(min, max int) int {
// 	var value = rand.Int() % (max - min + 1) + min
// 	return value
// }

// func main() {
// 	var diameter float64 = 15
// 	var area, circumference = calculate(diameter)

// 	fmt.Printf("Luas: %.2f \n", area)
// 	fmt.Printf("Keliling: %2.f \n", circumference)
// }

// func calculate(d float64) (float64, float64) {
// 	// hitung luas
// 	var area = math.Pi * math.Pow(d / 2, 2)
// 	// hitung keliling
// 	var circumference = math.Pi * d

// 	return area, circumference
// }

// predefined return value
// func calculate(d float64) (area float64, circumference float64) {
// 	area = math.Pi * math.Pow(d/2, 2)

// 	circumference = math.Pi * d

// 	return
// }

// variadic func
// func main() {
// 	var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
// 	var msg = fmt.Sprintf("Rata - rata : %.2f", avg)

// 	fmt.Printf(msg)
// }

// func calculate(numbers ...int) float64 {
// 	var total int = 0

// 	for _, number := range numbers {
// 		total += number
// 	}

// 	var avg = float64(total) / float64(len(numbers))

// 	return avg
// }

// variadic func with normal param
// func main() {
// 	yourHobbies("agam", "basket", "travel")
// }

// func yourHobbies(name string, hobbies ...string) {
// 	var hobbiesAsString = strings.Join(hobbies, ",")

// 	fmt.Printf("Hello my name is %s\n", name)
// 	fmt.Printf("My hobbies are : %s\n", hobbiesAsString)
// }

// func inside variable
// func main() {
// 	var getMinMax = func(n []int) (int, int) {
// 		var min, max int

// 		for i, e := range n {
// 			switch {
// 				case i == 0:
// 					max, min = e, e
// 				case e > max:
// 					max = e
// 				case e < min:
// 					min = e
// 			}
// 		}

// 		return min, max
// 	}

// 	var numbers = []int{2,3,4,5,6,7,8}
// 	var min, max = getMinMax(numbers)
// 	fmt.Printf("data: %v\nmin : %v\nmax : %v\n", numbers, min, max)
// }

// func as a param
// func main() {
// 	var data = []string{"wich", "john", "ethan"}
// 	var dataContainsO = filter(data, func(each string) bool {
// 		return strings.Contains(each, "o")
// 	})

// 	var dataLength5 = filter(data, func(each string) bool {
// 		return len(each) == 5
// 	})

// 	fmt.Println("data : ", data)
// 	fmt.Println("filter contain o : ", dataContainsO)
// 	fmt.Println("filter count char : ", dataLength5)
// }

// func filter(data []string, callback func(string) bool) []string {
// 	var result []string
// 	for _, each := range data {
// 		if filtered := callback(each); filtered {
// 			result = append(result, each)
// 		}
// 	}

// 	return result
// }

// pointer
// func main() {
// 	var numberA int = 4
// 	var numberB *int = &numberA

// 	fmt.Println("numberA (value)  :", numberA)
// 	fmt.Println("numberA (address)  :", &numberA)
// 	fmt.Println("numberB (value)  :", *numberB)
// 	fmt.Println("numberB (address)  :", numberB)

// 	fmt.Println("")

// 	numberA = 5
// 	fmt.Println("numberA (value)  :", numberA)
// 	fmt.Println("numberA (address)  :", &numberA)
// 	fmt.Println("numberB (value)  :", *numberB)
// 	fmt.Println("numberB (address)  :", numberB)

// }

// parameter pointer
// func main() {
// 	var number = 4
// 	fmt.Println("before : ", number)

// 	change(&number, 10)

// 	fmt.Println("after : ", number)
// }

// func change(original *int, value int) {
// 	*original = value
// }

// struct
// type student struct {
// 	grade int
// 	person
// 	age int
// }

type person struct {
	name string
	age  int
}

// func main() {
// 	var s1 student
// 	s1.name = "grinder wald"
// 	s1.grade = 2

// 	var s2 = student{"wise kid", 3}
// 	var s3 = student{name: "harry dend"}
// 	var s4 = student{grade: 3, name: "tony will"}
// 	// fmt.Println("Student Name : ", s1.name)
// 	// fmt.Println("Student Grade : ", s1.grade)
// 	fmt.Println("Student Name 1 : ", s1.name)
// 	fmt.Println("Student Name 2 : ", s2.name)
// 	fmt.Println("Student Name 3 : ", s3.name)
// 	fmt.Println("Student Name 4 : ", s4.name)
// }

// struct object pointer
// func main() {
// 	var s1 = student{name: "john", grade: 1}
// 	var s2 *student = &s1

// 	fmt.Println("name student 1 :", s1.name)
// 	fmt.Println("name student 4 :", s2.name)

// 	s2.name = "wick"
// 	fmt.Println("name student 1 :", s1.name)
// 	fmt.Println("name student 4 :", s2.name)
// }

// embedded struct
// func main() {
// var s1 = student{}
// s1.name = "john"
// s1.age = 12
// s1.person.age = 15
// s1.grade = 2

// fmt.Println("name : ", s1.person.name)
// fmt.Println("age : ", s1.person.age)
// fmt.Println("age : ", s1.age)
// fmt.Println("age : ", s1.grade)

// var s2 = struct {
// 	person
// 	grade int
// }{}

// s2.person = person{name: "thed", age: 44}
// s2.grade = 1

// fmt.Println("name : ", s2.person.name)
// fmt.Println("age : ", s2.person.age)
// fmt.Println("grade : ", s2.grade)

// var allStudents = []person{
// 	{name: "Wise", age: 19},
// 	{name: "Zedd", age: 21},
// 	{name: "Drek", age: 22},
// 	{name: "Pablo", age: 20},
// }

// for _, student := range allStudents {
// 	fmt.Println("Name : ", student.name)
// 	fmt.Println("Age : ", student.age)
// }

// // anonymous struct
// var lecturer = struct {
// 	name string
// 	age  int
// }{
// 	"Escobar",
// 	40,
// }

// fmt.Println("Lecturer Name : ", lecturer.name)
// fmt.Println("Lecturer Age : ", lecturer.age)

// nested struct
// var data = employee{}
// data.division = "Tech"
// data.hobbies = []string{
// 	"basket",
// 	"sport",
// 	"code",
// 	"gaming",
// }
// data.person.name = "Ricard"
// data.person.age = 23

// fmt.Println("Name : ", data.person.name)
// fmt.Println("Division : ", data.division)
// for _, hobby := range data.hobbies {
// 	fmt.Println("Hobbies : ", hobby)
// }

// }

// nested struct
// type employee struct {
// 	person struct {
// 		name string
// 		age  int
// 	}
// 	division string
// 	hobbies  []string
// }
type employee struct {
	person struct {
		name string
		age  int
	}
	division string
}

func (e employee) say(message string) {
	fmt.Println(message, e.person.name)
}

func (e employee) getNameAt(i int) string {
	return strings.Split(e.person.name, " ")[i-1]
}

func (e employee) changeName1(name string) {
	fmt.Println("-> on changeName1, name changed to : ", name)
	e.person.name = name
}

func (e *employee) changeName2(name string) {
	fmt.Println("-> on changeName2, name changes to : ", name)
	e.person.name = name
}

// func main() {
// 	var people = employee{person: person{name: "Brath Wild", age: 48}, division: "Kitchen"}

// 	people.say("Halo my name is")
// 	var name = people.getNameAt(2)
// 	fmt.Println("You can call me", name)

// 	fmt.Println("Name before change : ", people.person.name)
// 	people.changeName1("Brandon Whelter")
// 	fmt.Println("Name after changeName1", people.person.name)

// 	people.changeName2("Karlt Philip")
// 	fmt.Println("Name after changeName3", people.person.name)
// }

// type hitung interface {
// 	luas() float64
// 	keliling() float64
// }

type lingkaran struct {
	diameter float64
}

type persegi struct {
	sisi float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

// func main() {
// 	var bangunDatar hitung

// 	bangunDatar = persegi{10.0}
// 	fmt.Println("===== persegi")
// 	fmt.Println("luas      :", bangunDatar.luas())
// 	fmt.Println("keliling  :", bangunDatar.keliling())

// 	bangunDatar = lingkaran{14.0}
// 	fmt.Println("===== lingkaran")
// 	fmt.Println("luas      :", bangunDatar.luas())
// 	fmt.Println("keliling  :", bangunDatar.keliling())
// 	fmt.Println("jari-jari :", bangunDatar.(lingkaran).jariJari())

// }

// embedded interface
type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type hitung interface {
	hitung2d
	hitung3d
}

type kubus struct {
	sisi float64
}

func (k *kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}

func (k *kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}

func (k *kubus) keliling() float64 {
	return k.sisi * 12
}

// func main() {
// 	var bangunRuang hitung = &kubus{4}

// 	fmt.Println("Luas : ", bangunRuang.luas())
// 	fmt.Println("Volume : ", bangunRuang.volume())
// 	fmt.Println("Keliling : ", bangunRuang.keliling())
// }

// func main() {
// 	// var secret interface{}

// 	// secret = "john ethan"
// 	// fmt.Println(secret)

// 	// secret = []string{"apple", "banana", "mango", "starfruit"}
// 	// fmt.Println(secret)

// 	// secret = 12.4
// 	// fmt.Println(secret)

// 	// var data map[string]interface{}

// 	// data = map[string]interface{}{
// 	// 	"name" : "john wick",
// 	// 	"grade" : 2,
// 	// 	"hobby" : []string{"otomotif", "sport", "code"},
// 	// }

// 	// var data map[string]any

// 	// data = map[string]any{
// 	// 	"name" : "john wick",
// 	// 	"grade" : 2,
// 	// 	"hobby" : []string{"otomotif", "sport", "code"},
// 	// }
// 	// interface kosong
// 	// var secret interface{}

// 	// secret = 2

// 	// casting variable interface
// 	// var number = secret.(int) * 10
// 	// fmt.Println(secret, "hasil perkalian 10nya adalah", number)

// 	// secret = []string{"apple", "banana", "mango", "watermelon"}
// 	// var fruits = strings.Join(secret.([]string), ", ")
// 	// fmt.Println(fruits, "is my favorite fruits")

// 	// casting variable interface ke pointer

// 	// var secret interface{} = &peops{name: "Thony", age: 24}
// 	// var name = secret.(*peops).name
// 	// fmt.Println(name)

// 	// // combine slice, map, interface
// 	// var person = []map[string]interface{}{
// 	// 	{"name": "John", "age": 48},
// 	// 	{"name": "Youre", "age": 34},
// 	// 	{"name": "Perdy", "age": 28},
// 	// }

// 	// for _, each := range person {
// 	// 	fmt.Println(each["name"], "age is", each["age"])
// 	// }

// 	var fruits = []interface{}{
// 		map[string]interface{}{"name": "banana", "total": 10},
// 		[]string{"mango", "grape", "apple"},
// 		"orange",
// 		10,
// 	}

// 	for _,each := range fruits {
// 		fmt.Println(each)
// 	}
// }

// type peops struct {
// 	name string
// 	age  int
// }

// reflect

// func main() {
// 	// var number = 23
// 	// var reflectValue = reflect.ValueOf(number)

// 	// fmt.Println("tipe variabelnya adalah", reflectValue.Type())

// 	// if reflectValue.Kind() == reflect.Int {
// 	// 	fmt.Println("nilai variabel : ", reflectValue.Int())
// 	// }

// 	// var number = 24
// 	// var reflectValue = reflect.ValueOf(number)

// 	// fmt.Println("tipe variabel : ", reflectValue.Type())
// 	// fmt.Println("nilai variabel : ", reflectValue.Interface())

// 	// property variabel objek
// 	var s1 = &student{Name: "Herry", Grade: 4}
// 	s1.getPropertyInfo()
// }

// func (s *student) getPropertyInfo() {
// 	var reflectValue = reflect.ValueOf(s)

// 	if reflectValue.Kind() == reflect.Ptr {
// 		reflectValue = reflectValue.Elem()
// 	}

// 	var reflectType = reflectValue.Type()

// 	for i := 0; i < reflectValue.NumField(); i++ {
// 		fmt.Println("name : ", reflectType.Field(i).Name)
// 		fmt.Println("tipe data : ", reflectType.Field(i).Type)
// 		fmt.Println("nilai : ", reflectValue.Field(i).Interface())
// 		fmt.Println("")
// 	}

// }

// type student struct {
// 	Name  string
// 	Grade int
// }

// func (s *student) SetName(name string) {
// 	s.Name = name
// }

// func main() {
// 	var s1 = &student{Name: "Karri Brat", Grade: 2}
// 	fmt.Println("nama : ", s1.Name)

// 	var reflectValue = reflect.ValueOf(s1)
// 	var method = reflectValue.MethodByName("SetName")
// 	method.Call([]reflect.Value{
// 		reflect.ValueOf("Mordi"),
// 	})

// 	fmt.Println("nama : ", s1.Name)
// }

// func main() {
// 	runtime.GOMAXPROCS(2)

// 	go print(5, "halo")

// 	print(5, "good morning")

// 	var input string
// 	fmt.Scanln(&input)
// }

// func print(till int, message string) {
// 	for i := 0; i < till; i++ {
// 		fmt.Println((i+1), message)
// 	}
// }

// error
// func main() {
// 	// var input string
// 	// fmt.Print("Type some number : ")
// 	// fmt.Scanln(&input)

// 	// var number int
// 	// var err error
// 	// number, err = strconv.Atoi(input)

// 	// if err == nil {
// 	// 	fmt.Println(number, "is number")
// 	// } else {
// 	// 	fmt.Println("is not number")
// 	// 	fmt.Println(err.Error())
// 	// }

// 	// validate
// 	// defer catch()
// 	// defer func() {
// 	// 	if r := recover(); r != nil {
// 	// 		fmt.Println("Error occured", r)
// 	// 	} else {
// 	// 		fmt.Println("Application running perfectly")
// 	// 	}
// 	// }()

// 	// var input string
// 	// fmt.Print("Input your name : ")
// 	// fmt.Scanln(&input)

// 	// if valid, err := validate(input); valid {
// 	// 	fmt.Println("Halo", input)
// 	// } else {
// 	// 	panic(err.Error())
// 	// 	fmt.Println(err.Error())
// 	// }

// 	data := []string{"mango", "car", "chair", "table", "ant", "whale"}

// 	for _, each := range data {
// 		var number int
// 		var err error
// 		number, err = strconv.Atoi(each)

// 		func() {

// 			defer func() {
// 				if r := recover(); r != nil {
// 					fmt.Println("Panic occured on looping", each, "| message : ", r)
// 				} else {
// 					fmt.Println("Application running...")
// 				}
// 			}()
// 			if err == nil {
// 				fmt.Println(number)
// 			} else {
// 				panic(err.Error())
// 			}
// 		}()

// 	}

// }

// func validate(input string) (bool, error) {
// 	if strings.TrimSpace(input) == "" {
// 		return false, errors.New("Input cannot be empty")
// 	}

// 	return true, nil
// }

// func catch() {
// 	if r := recover(); r != nil {
// 		fmt.Println("Error occured", r)
// 	} else {
// 		fmt.Println("Application running perfectly")
// 	}
// }

// layout format string
// type student struct {
// 	name        string
// 	height      float64
// 	age         int32
// 	isGraduated bool
// 	hobbies     []string
// }

// var data = student{
// 	name:        "Waze",
// 	height:      189.5,
// 	age:         19,
// 	isGraduated: false,
// 	hobbies:     []string{"eating", "sleeping"},
// }

// func main() {
// 	// data numerik to string numerik berbasis biner
// 	fmt.Printf("%b\n", data.age)

// 	// format data numerik yang berupa kode unicode, menjadi bentuk string karakter unicode-nya.
// 	fmt.Printf("%c\n", 1400)
// 	fmt.Printf("%c\n", 1235)

// 	// data numerik to string numerik berbasis oktal
// 	fmt.Printf("%d\n", data.age)

// 	// data numerik desimal ke dalam bentuk notasi numerik standar Scientific notation.
// 	fmt.Printf("%e\n", data.height)
// 	fmt.Printf("%E\n", data.height)

// 	// memformat data numerik desimal, dengan lebar desimal bisa ditentukan. Secara default lebar digit desimal adalah 6 digit.
// 	fmt.Printf("%f\n", data.height)
// 	fmt.Printf("%.9f\n", data.height)
// 	fmt.Printf("%.2f\n", data.height)
// 	fmt.Printf("%.f\n", data.height)

// 	// memformat data numerik desimal, dengan lebar desimal bisa ditentukan
// 	fmt.Printf("%g\n", 0.123123123123)

// 	// memformat data numerik, menjadi bentuk string numerik berbasis 8 (oktal)
// 	fmt.Printf("%o\n", data.age)

// 	// memformat data pointer, mengembalikan alamat pointer referensi variabel-nya
// 	fmt.Printf("%p\n", &data.name)

// 	// escape string. Meskipun string yang dipakai menggunakan literal \ akan tetap di-escape.
// 	fmt.Printf("%q\n", `"name \ height"`)

// 	// Digunakan untuk memformat data string.
// 	fmt.Printf("%s", data.name)

// 	// memformat data boolean, menampilkan nilai bool-nya.
// 	fmt.Printf("%t\n", data.isGraduated)

// 	// Berfungsi untuk mengambil tipe variabel yang akan diformat.

// 	fmt.Printf("%T\n", data.name)
// 	// string

// 	fmt.Printf("%T\n", data.height)
// 	// float64

// 	fmt.Printf("%T\n", data.age)
// 	// int32

// 	fmt.Printf("%T\n", data.isGraduated)
// 	// bool

// 	fmt.Printf("%T\n", data.hobbies)
// 	// []string

// 	// untuk memformat data apa saja (termasuk data bertipe interface{}).
// 	fmt.Printf("%v\n", data)

// 	// untuk memformat struct, mengembalikan nama tiap property dan nilainya berurutan
// 	// sesuai dengan struktur struct
// 	fmt.Printf("%v+\n", data)

// 	// untuk memformat struct, mengembalikan nama dan nilai tiap property sesuai dengan struktur struct dan juga bagaimana objek tersebut dideklarasikan
// 	fmt.Printf("%#v\n", data)

// 	// memformat data numerik, menjadi bentuk string numerik berbasis 16 (heksadesimal)
// 	fmt.Printf("%x\n", data.age)

// }

// bermain dengan json dan object
type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	var jsonString = `{"Name" : "Johny Will", "Age": 38}`
	var jsonData = []byte(jsonString)

	var data User

	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("User: ", data.FullName)
	fmt.Println("Age: ", data.Age)

	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)

	fmt.Println("User: ", data1["Name"])
	fmt.Println("Age: ", data1["Age"])

	var data2 interface{}
	json.Unmarshal(jsonData, &data2)

	var decodedData = data2.(map[string]interface{})
	fmt.Print(decodedData)
	fmt.Println("User: ", decodedData["Name"])
	fmt.Println("Age: ", decodedData["Age"])

	// array json to array object
	var jsonString2 = `[
		{"Name" : "Kally", "Age":47},
		{"Name" : "Irving", "Age":27}
	]`

	var dataArray []User

	var err2 = json.Unmarshal([]byte(jsonString2), &dataArray)
	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}

	fmt.Println("user 1 : ", dataArray[0].FullName)
	fmt.Println("user 2 : ", dataArray[1].FullName)

	// encode object to json string
	var object = []User{{"Larry Thing", 39}, {"Yurrin", 34}}
	var jsonData2, err3 = json.Marshal(object)

	if err3 != nil {
		fmt.Println(err3.Error())
		return
	}

	var jsonString3 = string(jsonData2)
	fmt.Println(jsonString3)

}
