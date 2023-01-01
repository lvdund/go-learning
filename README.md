## Khả năng truy xuất
### Viết hoa vs thường

## Array
```php
    arr := [5]int{7, 10, 5, 11, 18}
    arr := [...]int{7, 10, 5, 11, 18}
    var arr [5]int
    arr2 := arr         // copy array
    arr2 := &arr        // arr2 -> arr
```
## Slice
```php
    slc := []int{7, 10, 5, 11, 18}
    slc2 := slc                     // slc2  -> slc
    slc3 := slc[1:3]                // 10, 5 -> len = 2, cap = 4
    slc3 := make([]int, 2, 3)       // 0, 0  -> len = 2, cap = 3
    slc3 := append(slc3, 1, 2, 3)   // 0, 0, 1, 2, 3 -> len = 5, cap = ??
    slc3 := append(slc3, []int{4, 5, 6}...)
```
## Map
```php
    m := make(map[string]int)       // map['key']'value' -> key != slice
    m = map[string]int {
        "vu"    : 18,
        "dung"  : 22,
        "giang" : 45
    }
    age, isExist := m["linh"]
    // delete, len
    m2 := m                         // m2 -> m
```
## Struct
```php
    type St struct {
        id          int
        name        string  `Length must > 3`   // -> tag
    }
    type St2 struct {
        St
        score       []int
    }

    // khai báo
    a := St {
        id:     12,
        name    "linh",
    }
    // tag : must import "reflect"
    field, _ := reflect.TypeOf(a).FieldByName("id")
```
## If else
```php
    if khai báo; true/false {}
    else{}
```
## Switch case
```php
    switch number {
        case 1:
        default:
    }                       // ko cần break
    switch {
        case true:
        default:
    }
    switch var i interface{} = 1; i.(type){
        case int:
        default:
    }
    // từ khóa fallthrough
```
## Loop
```php
    for i:=0; i<10; i++ {}
    for i < 10 { i++ }
    for {}                      // ko có điểm dừng
    Loop:
        for {
            if i > 10 { break Loop }
        }
    for index, valaue := range a {}     // a: map, string, array
```
## defer, recover, panic
```php
    defer ...   -> lệnh vào ngân xếp, thực hiện trước khi hàm trả về kết quả
    panic("...")-> lỗi, dừng chương trình
    recover ... -> check panic
        defer func() {
            if err := recover(); err != nil { ... }
        }
        panic("this is error")
```
## Pointer
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Con trỏ: & *

## Function
```php
    func foo(a, b int) (result int) {
        ...
        return
    }               // có thể truyền vào slide hoặc multiple var: a ...int
    func devide(a, b int) (int, error){
        ...
        return 1, fmt.Errorf("error")
    }

    // struct func
    type St struct { ... }
    func (a St) foo() { ... }
    ss := St{ ... }
    ss.foo()        // copy ss rồi dùng hàm foo() -> phải dùng pointer
```
## Interface
```php
    type Writer interface {
        Write([]byte)   (int, error)
    }
    type ConsoleWriter struct {}
    func (cw ConsoleWriter) Write(data []byte) (int, error){
        n, err := fmt.Println(string(data))
        return n, err
    }
    var w Writer = ConsoleWriter{}
    w.Write([]byte("hello world"))
```
## Goroutine - Channel - fan in fan out - work pool - 
### thực thi song song vận dụng sức mạnh của vi xử lý - đa tiến trình

## Creational Pattern 
1. **Singleton** – Đảm bảo rằng chỉ có duy nhất 1 instance của object tồn tại suốt ứng dụng.
2. **Factory Method** – Cho phép bạn tạo các họ đối tượng liên quan mà không chỉ định các lớp cụ thể của chúng.
3. **Abstract Factory** – Tạo ra các đối tượng phụ thuộc liên quan giống nhau.
4. **Builder** – Cho phép bạn từng bước xây dựng các đối tượng phức tạp. Mẫu này cho phép bạn tạo ra các loại và biểu diễn khác nhau của một đối tượng bằng cách sử dụng cùng một mã xây dựng.
5. **Prototype** - Cho phép bạn sao chép các đối tượng hiện có mà không làm cho mã của bạn phụ thuộc vào các lớp của chúng.
* Link tham khảo [tại đây](https://refactoring.guru/design-patterns/go)

## Important
* Encapsulation --> Packages
* Inheritance --> Composition
* Polymorphism --> Interfaces
* Abstraction ---> Embedding

## Tham khảo [Youtube](https://www.youtube.com/playlist?list=PLw-L1SGSvTEco7QvKTEd39wrMoTCPNUuN)