go run main.go -std="serhat" -ozel="." -std="cicek" -gun=0/7 -ozel="." -yil=2 -ay=true -ozel="." -rpad="./30"


-std -> Standart string değerler belirlmek için kullanılır 

-ozel -> Ozel karakter belirlemek için kullanılır şimdilik std den bir farkı yok ilerleyen zamanlarda güncellenebilir olması için farklı bir argüman ile kullanmayı tercih ettim

-gun -> verilen değeri "/" ifadesine göre parse eder ve başlangıç bitiş günlerini alır örneğin "0/7" değeri için haftanın 0. gününden 7. gününe kadar olan günleri alır

-yil -> verilen değere göre şuanki yıldan geriye doğru kaç yıl gideceğini belirler 

-ay -> liste oluşturma işlemine ayların dahil edilip edilmeyeceğini belirler

-rpad -> verilen değeri "/" ifadesine göre parse eder ve ona göre stringi sağa doğru doldurur. Örneğin "./30" ifadesi string değerini 30 a kadar "." ile doldurur

-lpad -> rpad ile aynı işlemi yapar fakat sağa doğru doldurur


Verilen argümanların sırasına göre dinamik olarak brute force listesi oluşturur. verilen ilk argüman ile bir liste oluşturur, ondan sonraki argümanlara göre listenin matrisi alınır. Tüm argümanlar dönülerek brute force listesi oluşturulmuş olur.