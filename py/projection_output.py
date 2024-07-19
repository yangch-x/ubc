import sys
from json import tool
from openai import OpenAI
import json
import os
from termcolor import colored

# projection  2 excel 2 txt output (excel)
# packing excel 2 pdf 2 txt 2   output (pdf)
# pdf 2 output py
client = OpenAI(api_key="sk-proj-t8eNeNsIoceuNnPgC8AJT3BlbkFJB7D5zYJv00KhqBYZjH1I")
ASSISTANT_ID = "asst_luJAZiqZnO92u5r39BsZ7Nxz"
JSON_OUTPUT_FILE = r'projection.json'
GPT_MODEL = "gpt-4o"

tools = [{
    "type": "function",
    "function": {
        "name": "export_order_details",
        "description": "Export order details to a file including various attributes like purchase order number, color, price, customer details, and more.",
        "parameters": {
            "type": "object",
            "properties": {
                "orders": {
                    "type": "array",
                    "description": "An array of order objects",
                    "items": {
                        "type": "object",
                        "properties": {
                            "ex_fty_in_house": {
                                "type": "string",
                                "description": "Ex-factory and in-house details"
                            },
                            "customer": {
                                "type": "string",
                                "description": "Customer details"
                            },
                            "customer_po": {
                                "type": "string",
                                "description": "Customer purchase order number"
                            },
                            "style_no": {
                                "type": "string",
                                "description": "Style number"
                            },
                            "desc_style_name": {
                                "type": "string",
                                "description": "Description or name of the style"
                            },
                            "color": {
                                "type": "string",
                                "description": "Color of the item"
                            },
                            "fabrication": {
                                "type": "string",
                                "description": "Fabrication details"
                            },
                            "qty_pc": {
                                "type": "number",
                                "description": "Quantity per piece"
                            },
                            "buy": {
                                "type": "number",
                                "description": "Buy price"
                            },
                            "ttl_buy": {
                                "type": "number",
                                "description": "Total buy amount"
                            },
                            "sell": {
                                "type": "number",
                                "description": "Sell price"
                            },
                            "ttl_sell": {
                                "type": "number",
                                "description": "Total sell amount"
                            },
                            "vendor": {
                                "type": "string",
                                "description": "Vendor details"
                            },
                            "water_resistant": {
                                "type": "string",
                                "description": "Water resistant status (Yes/No)"
                            },
                            "note": {
                                "type": "string",
                                "description": "Additional notes"
                            },
                            "country_brand_id": {
                                "type": "string",
                                "description": "Country and Brand ID associated with the order"
                            }
                        },
                        "required": [
                            "customer",
                            "customer_po",
                            "style_no",
                            "desc_style_name",
                            "color",
                            "fabrication",
                            "qty_pc",
                            "buy",
                            "sell"
                        ]
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


def _main(file_content):
    # in this cell we force the model to use get_n_day_weather_forecast
    SEED = 123
    messages = []
    messages.append({"role": "system", "content": "Don't make assumptions about what values to plug into functions."})
    commend = "Export order details to a file including various attributes like PO#, color, price, customer details, and more."
    seed = SEED
    temperature = 0.00001
    messages.append({"role": "user", "content": f"{commend}\n{file_content}"})
    chat_response = chat_completion_request(
        messages, tools=tools, tool_choice={"type": "function", "function": {"name": "export_order_details"}}
    )
    chat_response.choices[0].message
    arguments = chat_response.choices[0].message.tool_calls
    list_with_tool_call = arguments

    # 访问并解析arguments
    tool_call = list_with_tool_call[0]  # 假设您想处理列表中的第一个元素
    arguments_str = tool_call.function.arguments  # 获取arguments属性
    orders_data = json.loads(arguments_str)  # 解析JSON字符串为Python字典

    # Output orders_data to standard output with markers
    print("START_ORDERS_DATA")
    print(json.dumps(orders_data))
    print("END_ORDERS_DATA")


if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python script.py <file_content>")
        sys.exit(1)

    file_content = sys.argv[1]
    _main(file_content)
