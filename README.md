# golang-united-school-homework-7

The goal of this homework is to write tests in `students_test.go` file and achieve as much coverage as possible.
Цель этого домашнего задания — написать тесты в файле `students_test.go` и добиться максимального охвата.

Inside this repo you can find 2 files and one directory:
1. File `toBeTested.go` contains functions which should be tested. Do not change this file.
2. File `students_test.go` is a boilerplate of file with your tests. At the start of this file you can find `init()` function, which is required because of peculiarities of Autocode platform. This function will copy your file to `/autocode` folder, do not touch this function.
3. Folder `/autocode` to which your test will be automatically copied when you locally trigger your tests.
Внутри этого репо вы можете найти 2 файла и один каталог:
1. Файл `toBeTested.go` содержит функции, которые необходимо протестировать. Не изменяйте этот файл.
2. Файл `students_test.go` — это шаблон файла с вашими тестами. В начале этого файла вы можете найти функцию `init()`, которая необходима из-за особенностей платформы Autocode. Эта функция скопирует ваш файл в папку `/autocode`, не трогайте эту функцию.
3. Папка `/autocode`, в которую ваш тест будет автоматически скопирован, когда вы локально запускаете свои тесты.

Workflow to pass a homework: 
1. Write tests in `students_test.go`
2. Run your tests locally to check it status and resulted coverage. Each time you run tests your code will be copied to `/autocode` folder. The name of copied file (without an extension) is `students_test` - it's correct. Please, do not modify this file manually, do it every time by triggering `go test`.
3. Repeat steps 1 and 2 before you get desired result.
4. When you will be ready to submit and test you solution on Autocode, you will need to commit and push `autocode/students_test` file to git. **Make sure that this file was automatically updated.** 
Порядок выполнения домашнего задания:
1. Пишите тесты в `students_test.go`
2. Запустите свои тесты локально, чтобы проверить их статус и полученное покрытие. Каждый раз, когда вы запускаете тесты, ваш код будет скопирован в папку `/autocode`. Имя скопированного файла (без расширения) `students_test` - правильное. Пожалуйста, не изменяйте этот файл вручную, делайте это каждый раз, запуская `go test`.
3. Повторите шаги 1 и 2 до получения желаемого результата.
4. Когда вы будете готовы отправить и протестировать свое решение на Autocode, вам нужно будет зафиксировать и отправить файл `autocode/students_test` в git. **Убедитесь, что этот файл был автоматически обновлен.**

After submitting of your solution on Autocode you will get some logs. Keep in mind that you will only get the result of your tests if they fail. Otherwise, you will only get the coverage value. If you faced with some error or did not achieve desired coverage - modify your code and try again.
После отправки вашего решения на Autocode вы получите несколько журналов. Имейте в виду, что вы получите результат своих тестов только в том случае, если они не пройдут. В противном случае вы получите только значение покрытия. Если вы столкнулись с какой-то ошибкой или не добились желаемого покрытия - измените свой код и повторите попытку.
