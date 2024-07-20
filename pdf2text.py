import PyPDF2
import sys


def extract_pdf_text(file_stream):
    text = ""
    reader = PyPDF2.PdfReader(file_stream)
    number_of_pages = len(reader.pages)

    for page_number in range(number_of_pages):
        page = reader.pages[page_number]
        text += page.extract_text()
    return text


def main():
    # 读取文件流
    file_stream = sys.stdin.buffer
    text = extract_pdf_text(file_stream)
    print(text)


if __name__ == "__main__":
    main()
