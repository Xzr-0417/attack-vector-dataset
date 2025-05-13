# Polymorphic.py

### - Кодировка URL (URL Encoding)
```python
def url_encode(line):
    return urllib.parse.quote(line)
```
**Принцип работы**: Функция преобразует строку в формат, безопасный для передачи в URL. Небезопасные символы (например, пробелы, кириллица) заменяются на `%` и двузначное шестнадцатеричное число, соответствующее коду символа в таблице ASCII. Например:
```python
url_encode("Hello World!") → "Hello%20World%21"
```

### - Преобразование регистра символов (Case Conversion)
```python
def case_convert(line):
    converted = []
    letter_count = 0
    for char in line:
        if char.isalpha():
            letter_count += 1
            converted.append(char.swapcase() if letter_count % 2 == 1 else char)
        else:
            converted.append(char)
    return ''.join(converted)
```
**Принцип работы**: Функция меняет регистр букв через одну. Например:
```python
case_convert("Hello World!") → "hElLo wOrLd!"
```
Это делает строку менее узнаваемой для человека.

### - Кодирование Base64
```python
def base64_encode(line):
    return base64.b64encode(line.encode()).decode()
```
**Принцип работы**: Преобразует строку в представление из 64 печатаемых символов. Например:
```python
base64_encode("Hello") → "SGVsbG8="
```
Работает путем деления входных данных на 3-байтовые блоки и преобразования их в 4-символьные последовательности.

### - Unicode-эскейп (Unicode Escape)
```python
def unicode_escape(line):
    escaped = []
    for char in line:
        cp = ord(char)
        if cp <= 0xFFFF:
            escaped.append(f"\\u{cp:04X}")
        else:
            escaped.append(f"\\U{cp:08X}")
    return ''.join(escaped)
```
**Принцип работы**: Заменяет каждый символ на его Unicode-представление. Например:
```python
unicode_escape("A") → "\\u0041"
```
Символы представляются в формате `\uXXXX` (для BMP) или `\UXXXXXXXX` (для символов вне BMP).

### - HTML-эскейп (HTML Encoding)
```python
def html_encode(line):
    return html.escape(line, quote=True)
```
**Принцип работы**: Заменяет специальные символы HTML (`<`, `>`, `&` и др.) на соответствующие HTML-сущности. Например:
```python
html_encode("<script>alert('Hello')</script>") → "&lt;script&gt;alert(&#x27;Hello&#x27;)&lt;/script&gt;"
```
Это предотвращает интерпретацию символов как HTML-разметки.

### - Base32 Encoding
```python
def base32_encode(line):
    return base64.b32encode(line.encode()).decode()
```
**Принцип работы**: Преобразует данные в строку из 32 символов (A-Z, 2-7), где каждые 5 бит входных данных кодируются одним печатным символом. Например:
```python
base32_encode("Hello") → "JBSWY3DP"
```
Используется для создания компактного и человекочитаемого представления бинарных данных.

### - SHA-256 Hashing
```python
def sha256_encode(line):
    return hashlib.sha256(line.encode()).hexdigest()
```
**Принцип работы**: Генерирует фиксированный 256-битный хеш (64 шестнадцатеричных символа) с использованием криптографического алгоритма SHA-256. Например:
```python
sha256_encode("Hello") → "185f8db32271fe25f561a6fc938b2e264306ec304eda518007d1764826381969"
```
Любое изменение входной строки полностью меняет хеш. Обратное преобразование невозможно.
