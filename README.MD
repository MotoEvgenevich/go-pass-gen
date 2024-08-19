# Random Password Generator

## Description:

A simple CLI tool that generates random passwords with specified parameters, such as length, the inclusion of special characters, numbers, etc.

## Key Features:

	•	Generation of a password with a specified length.
	•	Option to include/exclude numbers, uppercase letters, and special characters.
	•	Saving the generated password to a file.

## Flags:

### The following forms are permited:

    -flag
    --flag   // double dashes are also permitted
    -flag=x
    -flag x  // non-boolean flags only

### flags Description:

- len=10 - leight of generating password
- digit=false - add digits to password (true is default)
- upletter=false - add uppercase letters to password(true is default)
- spec=false - add spec Symbols to password(true is default)
- save=true - save password to paswords.txt new string(true is default) 

## Usage Example:

go run main.go -len=10 -digit=false -upletter=false -spec=false -save=false   

# Генератор случайных паролей
## Описание:

Простой CLI-инструмент, который генерирует случайные пароли с заданными параметрами, такими как длина, наличие специальных символов, цифр и т.д.

## Основные функции:

    • Генерация пароля заданной длины.
    • Возможность включения/исключения цифр, заглавных букв и специальных символов.
    • Сохранение сгенерированного пароля в файл.

## Флаги(аргументы):

### Допускаются следующие формы записи:

-flag
--flag   // допускается также двойное тире
-flag=x
-flag x  // только для не булевых флагов

### Описание флагов(аргументов):

- `len=10` - длина генерируемого пароля
- `digit=false` - добавить цифры в пароль (по умолчанию `true`)
- `upletter=false` - добавить заглавные буквы в пароль (по умолчанию `true`)
- `spec=false` - добавить специальные символы в пароль (по умолчанию `true`)
- `save=true` - сохранить пароль в файл `passwords.txt` с новой строки (по умолчанию `true`)


## Пример использования:
```
go run main.go -len=10 -digit=false -upletter=false -spec=false -save=false   

```