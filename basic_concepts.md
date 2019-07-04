# Coding
## Smell code
berikut adalah salah satu contoh code smell

```swift
func sortBasedMethod(data:[Int], method:String){

    var result:[Int]

    switch method {
        case "bubbleSort":
        // perform bubble sort
        result = sorted_data
        case "heapSort":
        // perform heap sort
        result = sorted_data
        case "mergeSort":
        // perform merge sort
        result = sorted_data
        case "insertionSort":
        // perform merge sort
        result = sorted_data     
    }
}
```

terlalu sering menggunakan ```if``` atau ```switch``` tidak baik dalam OOP dan dapat melanggar prinsip open clode dalam OOP.

Untuk mengatasinya, dapat dilakuakn dengan cara melakukan Polymorphism, seperti berikut

```swift
protocol SortingInterface {
    sortData(data:[Int]) -> [Int]
}
```

```swift
class HeapSort:SortingInterface {
    sortData(data:[Int]) -> [Int]{
        // perform hear sort
        return data
    }
}
```

```swift
class MergeSort:SortingInterface {
    sortData(data:[Int]) -> [Int]{
        // perform merge sort
        return data
    }
}
```

```swift
class BubbleSort:SortingInterface {
    sortData(data:[Int]) -> [Int]{
        // perform bubble sort
        return data
    }
}
```

```swift
class InsertionSort:SortingInterface {
    sortData(data:[Int]) -> [Int]{
        // perform insertsion sort
        return data
    }
}
```

jika ada metode sorting baru yang ingin ditambahkan, dapat dilakukan dengan meng-inheritance interface sorting tanpa harus mengubah method di luar class

## Depedency Injection (DI)

Depedency Injection adalah pola desain yang memasukkan beberapa depdencies ke dalam beberapa objects yang menggunakannya.

Misalkan terdapat dua kelas yang menyediakan gateway pembayaran ke dalam dua layanan pembayaran berbeda. Pertamatama, kedua kelas tersebut harus dibuat interface-nya terlebih dahulu.

```swift
protocol PaymentGatewayInterface{
    func pay(amount: Int) -> Result<Success, Error>
}
```

Lalu interface tersebut diimplementasikan ke dalam dua kelas yang berbeda sesuai dengan layanan pembayarannya.

```swift
class GoPayGateway:PaymentGatewayInterface {
    func pay(amount: Int) -> Result<Success, Error> {
        // Lakukan pembayaran ke GoPay
    }
}

class OVOGateway:PaymentGatewayInterface {
    func pay(amount: Int) -> Result<Success, Error> {
        // Lakukan pembayaran ke OVO
    }
}
```

Lalu terdapat kelas yang butuh menggunakan melakukan proses pembayaran, maka interface dapat di-inject atau dimasukkan ke dalam kelas tersebut

```swift
class CheckoutRoute {
    var paymentGateway:PaymentGatewayInterface

    func setPaymentGateway(paymentGateway:PaymentGatewayInterface){
        self.paymentGateway = paymentGateway
    }

    func doPayment(){
        var amountToPay: Int
        // calculate amount to pay
        self.paymentGateway.pay(amountToPay)
    }
}
```

Penggunaan DI diperlukan agar: 

1. Mempermudah unit test. Dalam contoh code di atas payment gateway palsu dapat dibuat untuk tujuan testing.

2. Code boilerplate dapat dikurangi karena tidak harus meng-instansiasi kelas untuk paymeny gateway yang berbeda

3. Dapat melakukan penambahan kelas untuk metode pembayaran lain, tanpa harus merubah kelas `CheckoutRoute`

4. Mengurangi kebergantungan ke banyak kelas lain

# REST

## GET

Method `GET` seharusnya hanya digunakan untuk mendapatkan data dari server ke client

Method `GET` tidak boleh digunakan untuk mengirimkan informasi yang sensitf, karena data yang dikirimkan melalui method ini akan tampil langsun di url.

## POST

Method `POST` digunakan untuk mengirimkan data dari client ke server, data yang dikirimkan tidak terbatas, dapat berupa karakter ASCII ataupun binary data

Method `POST` seharusnya tidak digunakan untuk meminta data dari server, karena tidak dapat disimpan ke dalam cache


