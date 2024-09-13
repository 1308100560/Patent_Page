from PIL import Image

def convert_transparent_to_white(input_image_path, output_image_path):
    # 打开图像文件
    with Image.open(input_image_path) as img:
        # 确保图像有 alpha 通道
        img = img.convert('RGBA')

        # 获取图像数据
        datas = img.getdata()

        # 创建新的图像数据列表
        new_data = []
        for item in datas:
            # 检查是否是透明像素
            if item[3] != 0:
                # 将透明像素替换为白色
                new_data.append((255, 255, 255, 255))
            else:
                new_data.append(item)

        # 更新图像数据
        img.putdata(new_data)

        # 保存新图像
        img.save(output_image_path)

# 使用函数
input_path = 'input_image.png'
output_path = 'path_to_output_image.png'
convert_transparent_to_white(input_path, output_path)
