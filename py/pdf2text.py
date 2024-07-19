import PyPDF2
import os


def extract_pdf_text(file_path):
    text = ""
    with open(file_path, 'rb') as file:
        reader = PyPDF2.PdfReader(file)
        number_of_pages = len(reader.pages)

        for page_number in range(number_of_pages):
            page = reader.pages[page_number]
            text += page.extract_text()
    return text


def main():
    global text
    file_path = r'PACKING LIST for MIDNIGHT to US 17CTNS  ETD 6-12.pdf'
    output_file = 'ccc.txt'

    if os.path.exists(file_path):
        text = extract_pdf_text(file_path)
    print(text)
    #     # 将提取的文本写入到输出文本文件
    #     with open(output_file, 'w', encoding='utf-8') as file:
    #         file.write(text)
    #     print(f"Data exported successfully to {output_file}")
    # else:
    #     print("File does not exist at the specified path. Please check the file path.")


if __name__ == "__main__":
    main()
