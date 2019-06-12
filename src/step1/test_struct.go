package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
	haha    *string
}

func main() {
	var Book1 Books /* Declare Book1 of type Book */
	var Book2 Books /* Declare Book2 of type Book */
	//var tmp_str_ptr *string
	tmp_str := string("haha")
	tmp_str_ptr := &tmp_str
	//tmp_str_ptr := &(string("haha"))

	Book3 := Books{"haha", "you", "sb", 0, tmp_str_ptr}
	var Book4 *Books = &Books{"haha", "you", "sb", 0, nil}
	Book5 := &Books{title: "haha", author: "you", book_id: 0}
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})
	fmt.Println("*Book5", *Book5)
	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	/* 打印 Book1 信息 */
	printBook(&Book1)

	/* 打印 Book2 信息 */
	printBook(&Book2)
	printBook(&Book3)
	printBook(Book4)

}
func printBook(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}
