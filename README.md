# Пример организации проекта на Go

Репозиторий отвечает на вопрос "Как организовать проект на Go?"
Он включает в себя как публичные реюзабельные библиотеки (`multiplylib` и `sumlib`), библиотеку исключительно для внутреннего использования (`commandlineutils`) и две программы для терминала (`sum` и `multiply`).

Для использования проекта как библиотеки, достаточно импортировать его:

```
// ...
import (
	sgpl "github.com/wcademy/simplegoprojectlayout"
	"github.com/wcademy/simplegoprojectlayout/sumlib"
)

// ...
ints := sgpl.Ints{1, 2, 3}
result := sumlib.Sum(ints)
// ...
```  

Для установки программ:
```shell script
go get github.com/wcademy/simple-go-project-layout/cmd/sum
# или
go get github.com/wcademy/simple-go-project-layout/cmd/multiply
```

Все пакеты в модуле могут быть импортированы снаружи, кроме расположенных внутри `internal`, которые могут использоваться только внутри самого модуля.

Запуск тестов:
```shell script
golangci-lint run --fix && go test ./...
```

Подробнее [здесь](https://wcademy.ru/%D0%9A%D0%B0%D0%BA%20%D0%BE%D1%80%D0%B3%D0%B0%D0%BD%D0%B8%D0%B7%D0%BE%D0%B2%D0%B0%D1%82%D1%8C%20%D0%BF%D1%80%D0%BE%D0%B5%D0%BA%D1%82%20%D0%BD%D0%B0%20go/).
