import urllib.parse
import sys
import base64
import html
import hashlib

#======================= Core Encoding Functions ========================
def url_encode(line):
    return urllib.parse.quote(line)

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

def base64_encode(line):
    return base64.b64encode(line.encode()).decode()

def base32_encode(line):
    return base64.b32encode(line.encode()).decode()

def unicode_escape(line):
    escaped = []
    for char in line:
        cp = ord(char)
        if cp <= 0xFFFF:
            escaped.append(f"\\u{cp:04X}")
        else:
            escaped.append(f"\\U{cp:08X}")
    return ''.join(escaped)

def html_encode(line):
    return html.escape(line, quote=True)

def sha256_encode(line):
    return hashlib.sha256(line.encode()).hexdigest()

def ascii_encode(line):  # 新增ASCII编码
    try:
        encoded_bytes = line.encode('ascii')
    except UnicodeEncodeError:
        encoded_bytes = line.encode('ascii', 'replace')
    ascii_str = ' '.join(f'{byte:02x}' for byte in encoded_bytes)
    return ascii_str

def base16_encode(line):  # 新增Base16编码
    encoded_bytes = base64.b16encode(line.encode('utf-8'))
    return encoded_bytes.decode('ascii')

#======================= Encoder List ========================
ENCODERS = [
    ("url", url_encode),
    ("case", case_convert),
    ("base64", base64_encode),
    ("base32", base32_encode),
    ("unicode", unicode_escape),
    ("html", html_encode),
    ("sha256", sha256_encode),
    ("ascii", ascii_encode),  # 新增条目
    ("base16", base16_encode),  # 新增条目
]

#======================= Main Processing Logic ========================
def process_file(input_file, output_file, encode_mode):
    try:
        with open(input_file, "rb") as src, \
             open(output_file, "w", encoding="utf-8") as dest:
            
            for line_num, byte_line in enumerate(src, 1):
                try:
                    raw_line = byte_line.decode('utf-8').rstrip('\n')
                except UnicodeDecodeError:
                    print(f"Skip line {line_num}: Invalid UTF-8 encoding")
                    continue
                
                if encode_mode == "all":
                    for enc_name, encoder in ENCODERS:
                        try:
                            encoded = encoder(raw_line)
                            dest.write(f"{encoded}\n")
                        except Exception as e:
                            print(f"Line {line_num} [{enc_name}] encoding failed - {str(e)}")
                else:
                    encoder = dict(ENCODERS).get(encode_mode)
                    if not encoder:
                        raise ValueError(f"Invalid encoding type: {encode_mode}")
                    encoded = encoder(raw_line)
                    dest.write(f"{encoded}\n")

        print(f"Processing completed! Results saved to {output_file}")

    except FileNotFoundError:
        print(f"Error: Input file {input_file} not found")
    except PermissionError:
        print(f"Error: No write permission for {output_file}")
    except Exception as e:
        print(f"Program terminated abnormally: {str(e)}")

#======================= Command Line Argument Processing ========================
if __name__ == "__main__":
    if len(sys.argv) not in [3, 4]:
        print("Usage: python3 unified_encoder.py <input file> <output file> [encoding type]")
        print("Encoding types (optional): url, case, base64, base32, unicode, html, sha256, ascii, base16")  # 更新提示
        print("Example1 (all encodings): python3 unified_encoder.py input.txt output.txt")
        print("Example2 (single encoding): python3 unified_encoder.py input.txt output.txt url")
        sys.exit(1)

    if len(sys.argv) == 3:
        _, input_file, output_file = sys.argv
        encode_mode = "all"
    else:
        _, input_file, output_file, encode_type = sys.argv
        encode_mode = encode_type.lower()

    try:
        process_file(input_file, output_file, encode_mode)
    except ValueError as ve:
        print(ve)
        sys.exit(1)
