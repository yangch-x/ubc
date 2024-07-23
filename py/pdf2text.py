import PyPDF2
import os
import argparse

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
    parser = argparse.ArgumentParser(description='Extract text from a PDF file.')
    parser.add_argument('file_path', type=str, help='The path to the PDF file.')
    args = parser.parse_args()

    file_path = args.file_path

    if os.path.exists(file_path):
        text = extract_pdf_text(file_path)
        print(text)
    else:
        print("File does not exist at the specified path. Please check the file path.")

if __name__ == "__main__":
    main()
