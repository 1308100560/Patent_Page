import os
from bs4 import BeautifulSoup
import requests
import json
import time
from datetime import datetime

post_url = 'https://www.chonghus.com/hxapi/qc/zytxt'
headers = {
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0',
    'Content-Type': 'application/json'
}
data = {
    'type': 'search_zy_txt',
    'Context': '利用超声波对污泥进行预处理，可以破碎污泥的絮体结构，有利于污泥分散，增加污泥的比表面积，有利于后续的热水解处理效率',
    'total': '10',
    'title': '摘要全文查重'
}
try:
    response = requests.post(post_url, headers=headers, json=data, timeout=10)
    response.raise_for_status()  # 检查响应是否成功
except requests.exceptions.RequestException as e:
    print(f"请求失败: {e}")
    exit()

# 解析响应内容
soup = BeautifulSoup(response.content, 'lxml')
json_text = soup.p.text
data = json.loads(json_text)
unique_patents = {}
# Step 1: 直接合并所有数据
for index, item in enumerate(data['msg']):
    unique_patents[str(index)] = item

# Step 2: 根据 pat_name 进行去重
temp_dict = {}
for key, value in unique_patents.items():
    pat_name = value['pat_name']
    if pat_name not in temp_dict:
        temp_dict[pat_name] = value

# 将去重后的数据更新回 unique_patents
unique_patents = {str(i): v for i, v in enumerate(temp_dict.values())}

# Step 3: 根据 uuid 进行详细数据的合并
for key in unique_patents:
    uuid = unique_patents[key]['uuid']
    url = f'https://www.chonghus.com/hxapi2/pat/detail?id={uuid}'
    headers = {
        'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36'
    }
    response = requests.get(url, headers=headers)
    if response.status_code == 200:
        response.encoding = 'utf-8'
        json_data = response.json()
        msg_list = json_data.get('msg', {})
        response_id = msg_list.get('id')
        for key, patent in unique_patents.items():
            if patent.get('uuid') == response_id:
                unique_patents[key].update(msg_list)
                break
for key in unique_patents:
    patent = unique_patents[key]
    patent.pop('xh', None)  # ?
    patent.pop('flag', None)  # ?
    patent.pop('uuid', None)
    patent.pop('id', None)
    patent.pop('sq_date', None)  # 申请日
    patent.pop('pat_status', None)  #在审/有效
    patent.pop('pat_type', None)  # 发明公布/实用新型
    patent.pop('img_src', None)  # 示意图
    patent.pop('sqr_cn', None)  # 申请人？
    patent.pop('sqr_cn_first', None)  # 申请人？
    patent.pop('address', None)  # 申请人
    patent.pop('fmr_cn', None)  # 发明人
    patent.pop('main_fl_sec', None)  # 主分类号？
    patent.pop('main_fl', None)  # 主分类号？
    patent.pop('fl', None)  # 分类号
    patent.pop('flows', None)  # 法律状态变更历史
    patent.pop('open_date', None)  # 公开日
    patent.pop('dljg', None)  # 代理机构
    patent.pop('dlr', None)  # 代理人
    patent.pop('expireDate', None)  # 失效日
    patent.pop('yxq', None)  # ？
    patent.pop('gjsq', None)  # ？
    patent.pop('gjgb', None)  # ？
    patent.pop('img_list', None)  # 附图
    patent.pop('zqx', None)  # ？

json_data = json.dumps(unique_patents, ensure_ascii=False, indent=4)

# 将格式化后的HTML内容写入文件
output_file = 'output_file.html'
with open(output_file, "w", encoding='utf-8') as fd:
    fd.write(json_data)

print(f"内容已写入 {output_file}")
breakpoint()
# 如果需要进一步处理 soup 对象，可在此处编写代码
# 如 soup.find_all('tag') 等操作
