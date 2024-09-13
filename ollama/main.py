import os
import time

from docx import Document
from typing import Union

from fastapi import FastAPI, Request, Form, HTTPException
from fastapi.templating import Jinja2Templates
from starlette.responses import FileResponse

from patent_disclosure_generator import generate_patent_disclosure

# 注：要想启动服务器，需要使用终端或者命令提示符，将路径cd到当前文件所在的文件夹，然后输入运行：uvicorn main:app --reload
#    这条命令里的 main 是当前文件的Python文件名（不包含.py扩展名），app是FastAPI实例的名称（不用改）。--reload参数会在代码更改时自动重新加载服务器。


# 创建
app = FastAPI()
templates = Jinja2Templates(directory="templates")


# 业务路由：
# post方式的表单提交到这个路由，能接收表单提交的值，经过一些业务处理后，能将一些值返回到新的html页面。
@app.post("/submit-form")
# 通过函数形参获取值
async def submit_form(
        request: Request,
        str1: str = Form(...),  # 发明名称
        str2: str = Form(...),  # 技术方案
        str3: str = Form(...),  # 产生效果
        str4: str = Form(...)):  # 具体实施例

    print("已经接收到了表单发送的数据：", str1, str2, str3, str4)

    # message = "发明名称：" + str1 + "\n技术方案：" + str2 + "\n产生效果：" + str3 + "\n具体实施例：" + str4
    # print("##最终组合的message：", message)
    #
    # # 数据处理，准备数据
    # doc_bytes = generate_patent_disclosure(message)  # 调用LLM业务函数，对输入的文本进行处理，并返回docx文件
    # print("调用完成，返回值已经取得。")  # 判断是否正常返回。
    # print("调用generate_patent_disclosure函数的返回值类型为：", type(doc_bytes))  # 判断返回值类型。
    #
    # # 将返回的东西转化回word文档
    # title_time = time.time()
    # title_time = str(title_time)  # 转为字符串
    # print("获取到的时间：", title_time, "   数据类型：", type(title_time))
    # with open(f'./word_file/{str1}_{title_time}_hd.docx', 'wb') as file:
    #     file.write(doc_bytes)
    #
    #
    # # 返回值区域
    # # * 这个位置第一和第三个参数需要修改 *
    # return FileResponse(path=f"./word_file/{str1}_{title_time}_hd.docx",
    #                     media_type='application/vnd.openxmlformats-officedocument.wordprocessingml.document',
    #                     filename=f"{str1}_{title_time}_qd.docx")
    return str1

# 测试用的api 1
#   可以用来测试get方式下的前到后和后到前的参数传递和模版渲染，可以通过路由对应的url携带参数访问该路由，并在控制台输出验证后返回到greet.html页面并发送参数。
#   当你访问 http://127.0.0.1:8000/greet/John 时，你应该会看到页面显示“Hello, John!”
@app.get("/greet/{name}")
async def greet(request: Request, name: str):
    # 准备数据
    greeting = f"Hello, {name}!"
    print(greeting, "name数据类型：", type(name))

    # 渲染模板并传递数据
    return templates.TemplateResponse("greet.html", {"request": request, "greeting": greeting})


# 测试用的api 2
#   可以通过该路由来测试docx文档的创建和返回，该文档被返回后会被浏览器直接下载。
#   当你访问 http://127.0.0.1:8000/greet2/John 时，应该会返回并下载一个word文档，里面只有一个段落，一个词，就是John
# @app.get("/greet2/{name}")
# async def greet(request: Request, name: str):
#     # 准备数据
#     greeting = f"Hello, {name}!"
#     # 创建一个新的Word文档
#     doc = Document()  # 创建一个新的Word文档
#     doc.add_paragraph(name)  # 添加一个段落
#
#     # 保存文档到服务器上的某个位置
#     # 注意：这里需要确保你的应用有权限写入这个目录
#     file_path = f"/document/document_{title_time}.docx"  # 文件路径
#     print("文件路径和名称：", file_path)
#     os.makedirs(os.path.dirname(file_path), exist_ok=True)  # 确保目录存在
#     doc.save(file_path)  # 保存文档
#
#     # 返回文件下载链接（这里简化处理，直接返回文件）
#     return FileResponse(path=file_path,
#                         media_type='application/vnd.openxmlformats-officedocument.wordprocessingml.document',
#                         filename=f"document_{title_time}.docx")  # 直接返回并下载。路径是浏览器下载路径。


# 根路由，用来测试基本的连接和服务器有没有正常启动。
# 当你访问 http://127.0.0.1:8000 时，跳转到一个html页面，上面显示：你好里世界:
@app.get("/")
def read_root(request: Request):
    print("接收到了")
    # return {"Hello": "World"}
    return templates.TemplateResponse("index.html", {"request": request})
