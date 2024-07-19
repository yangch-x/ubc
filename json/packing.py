from json import tool
from openai import OpenAI
import json
import os
from termcolor import colored
client = OpenAI(api_key="sk-proj-t8eNeNsIoceuNnPgC8AJT3BlbkFJB7D5zYJv00KhqBYZjH1I")
ASSISTANT_ID = "asst_luJAZiqZnO92u5r39BsZ7Nxz"
INPUT_FILE = r'C:\Users\85431\OneDrive\桌面\Booking&packing\packing_input.txt'
JSON_OUTPUT_FILE = r'C:\Users\85431\OneDrive\桌面\Booking&packing\packing_output_data.json'
GPT_MODEL = "gpt-4o"
# Upload a file with an "assistants" purpose
with open(INPUT_FILE, 'r', encoding='utf-8') as file:
    file_content = file.read()
tools = [
    {
    "type": "function",
    "function": {
        "name": "export_simple_order_details",
        "description": "Export simplified order details to a file including PO#, color, style number, sizes, and quantities.",
        "parameters": {
            "type": "object",
            "properties": {
                "orders": {
                    "type": "array",
                    "description": "An array of simplified order objects",
                    "items": {
                        "type": "object",
                        "properties": {
                            "po": {
                                "type": "string",
                                "description": "Purchase order number"
                            },
                            "color": {
                                "type": "string",
                                "description": "Color of the item"
                            },
                            "style_number": {
                                "type": "string",
                                "description": "Style number of the item"
                            },
                            "sizes": {
                                "type": "array",
                                "description": "List of sizes and their quantities",
                                "items": {
                                    "type": "object",
                                    "properties": {
                                        "size": {
                                            "type": "string",
                                            "description": "Size of the item"
                                        },
                                        "quantity": {
                                            "type": "number",
                                            "description": "Quantity of the item in this size"
                                        }
                                    }
                                }
                            },
                            "total_quantity": {
                                "type": "number",
                                "description": "Total quantity of all sizes combined"
                            },
                            "CTNS": {
                                "type": "number",
                                "description": "Number of cartons for the order"
                            }
                        },
                        "required": ["po", "color", "style_number", "sizes", "total_quantity","CTNS"]
                    }
                }
            },
            "required": ["orders"]
        }
    }
}

]

def chat_completion_request(messages, tools=None, tool_choice=None, model=GPT_MODEL):
    
    try:
        response = client.chat.completions.create(
            model=model,
            messages=messages,
            tools=tools,
            tool_choice=tool_choice,
        )
        return response
    except Exception as e:
        print("Unable to generate ChatCompletion response")
        print(f"Exception: {e}")
        return e
def pretty_print_conversation(messages):
    role_to_color = {
        "system": "red",
        "user": "green",
        "assistant": "blue",
        "function": "magenta",
    }
    
    for message in messages:
        if message["role"] == "system":
            print(colored(f"system: {message['content']}\n", role_to_color[message["role"]]))
        elif message["role"] == "user":
            print(colored(f"user: {message['content']}\n", role_to_color[message["role"]]))
        elif message["role"] == "assistant" and message.get("function_call"):
            print(colored(f"assistant: {message['function_call']}\n", role_to_color[message["role"]]))
        elif message["role"] == "assistant" and not message.get("function_call"):
            print(colored(f"assistant: {message['content']}\n", role_to_color[message["role"]]))
        elif message["role"] == "function":
            print(colored(f"function ({message['name']}): {message['content']}\n", role_to_color[message["role"]]))
# in this cell we force the model to use get_n_day_weather_forecast
SEED = 123
messages = []
messages.append({"role": "system", "content": "Don't make assumptions about what values to plug into functions."})
commend = "Export order details to a file including various attributes like PO#, color, style#,size,quanty and more."
seed = SEED
temperature=0.00001
messages.append({"role": "user", "content": f"{commend}\n{file_content}"})
chat_response = chat_completion_request(
    messages, tools=tools, tool_choice={"type": "function", "function": {"name": "export_simple_order_details"}}
)
chat_response.choices[0].message
arguments = chat_response.choices[0].message.tool_calls
list_with_tool_call = arguments

# 访问并解析arguments
tool_call = list_with_tool_call[0]  # 假设您想处理列表中的第一个元素
arguments_str = tool_call.function.arguments  # 获取arguments属性
orders_data = json.loads(arguments_str)  # 解析JSON字符串为Python字典

# 将数据写入JSON文件
with open(JSON_OUTPUT_FILE, 'w', encoding='utf-8') as json_file:
    json.dump(orders_data, json_file, ensure_ascii=False, indent=4)

print("数据已经成功保存到'orders_output.json'文件中。")