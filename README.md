#Duplicate-file-finder - программа для поиска дубликатов файла

### Running

####Запуск с параметрами 
В программе предусмотрено два параметра:
1. -dir - директория для поиска дублей
2. -delete - необходимость удаления дублей

Пример запуска: _go run main.go -dir="E:\temp" -delete=false_

Будут найдены и напечатаны в консоль все дубли в директории "E:\temp", 
однако они не будут удалены поскольку параметр -delete принимает значение false.

####Запуск без параметров
При запуска приложения без параметров, найдены и напечатаны в консоль все дубли из директории запуска приложения.

Пример запуска: go run main.go

### TODO:

1. Покрыть программу тестами
2. Дополнить документацию и сделать example
3. Добавить -h/--help
