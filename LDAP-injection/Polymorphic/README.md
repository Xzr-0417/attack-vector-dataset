```markdown
# Документация по инструменту генерации запутанных载荷

Данный инструмент предназначен для кодирования и запутывания вводимого текста (обычно атакующих载荷), чтобы повысить его скрытность и способность обходить системы обнаружения. Ниже приведен подробный анализ各个功能模块的:

## 核心编码函数模块

### URL编码
```python
def url_encode(line):
    return urllib.parse.quote(line)
```

**Принцип запутывания**:
- Преобразует специальные символы (например, пробелы, знаки препинания) в соответствующие Unicode-коды
- Использует формат кодирования с процентами (%XX), где XX — шестнадцатеричное значение символа
- Например, пробел преобразуется в `%20`, а символ новой строки в `%0A`
- Такой тип кодирования主要用于网络传输中的URL参数，以确保特殊字符不会破坏协议格式

### Преобразование регистра букв
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

**Принцип запутывания**:
- Постепенно обрабатывает буквенные символы в тексте
- Поочередно преобразует регистр букв в зависимости от их позиции (четная или нечетная)
- Например, "Payload" будет преобразован в "PaYlOaD" и类似的 формы
- С помощью нарушения визуальной формы слов原文的增加了文本识别的难度
- 保留非字母字符的原始形态，确保编码后的文本语法结构仍然合法

### Base64编码
```python
def base64_encode(line):
    return base64.b64encode(line.encode()).decode()
```

**Принцип запутывания**:
- Разбивает каждый байт на единицы 6 бит, а затем重新组合这些单元
- 使用64个可打印字符（A-Z、a-z、0-9、+、/）来表示这些单元
- 最后添加等号（=）进行填充，确保编码后数据长度符合要求
- 例如，"Hello"会被编码为"SGVsbG8="
- 这种编码将原始数据转换为纯文本格式，可方便地嵌入到各种协议中

### Unicode转义
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

**Принцип запутывания**:
- Converts each character to its corresponding Unicode code point
- 对于基本多文种平面（BMP）内的字符，使用`\uXXXX`格式
- 对于BMP外的字符，使用`\UXXXXXXXX`格式
- 例如，"A"（U+0041）转换为`\u0041`
- 这种编码将字符表示转换为纯文本形式，便于跨平台传输和存储

### HTML实体编码
```python
def html_encode(line):
    return html.escape(line, quote=True)
```

**Принцип запутывания**:
- Converts special characters to their corresponding HTML entities
- 常见转换包括：
  - `<` → `&lt;`
  - `>` → `&gt;`
  - `&` → `&amp;`
  - `"` → `&quot;`
  - `'` → `&#x27;`
- 这种编码确保文本在HTML环境中被正确解析，避免被误认为HTML标签
- 例如，`<script>`会被转换为`&lt;script&gt;`

## 编码器列表模块
```python
ENCODERS = [
    ("url", url_encode),
    ("case", case_convert),
    ("base64", base64_encode),
    ("unicode", unicode_escape),
    ("html", html_encode)
]
```

Этот модуль определяет поддерживаемые типы кодирования и соответствующие им функции, что позволяет быстро вызывать нужные функции обработки кодирования по имени.

## 使用示例

### Режим полного кодирования
```bash
python3 unified_encoder.py input.txt output.txt
```

### Режим отдельного кодирования
```bash
python3 unified_encoder.py input.txt output.txt url
```
