<template>
  <div class="container">
    <div class="description">
      <i class="el-icon-info-circle"></i>
      <p>本栏目所显示的技术领域是根据IPC分类表设置，根据具体技术方案归属哪个分类，可以选择具体类别，直接进入下一个栏目，特此说明。</p>
      <p>如果不能确定技术方案属于哪个分类，可以点击下面的模糊类别，直接进入下一个栏目，但准确性会有所差别。</p>
    </div>

    <div class="search-container">
      <el-button type="primary" class="search-button" @click="fuzzySearch">
        <i class="el-icon-search"></i> 模糊搜索
      </el-button>
    </div>

    <div class="category-section">
      <div class="category-grid">
        <el-button v-for="cat in 'ABCD'" :key="cat"
                   class="category-button" @click="selectedCategory = cat">
          <span class="category-letter">{{ cat }}</span>
          <span class="category-name">{{ getCategoryName(cat) }}</span>
        </el-button>
      </div>
      <div class="category-grid">
        <el-button v-for="cat in 'EFGH'" :key="cat"
                   class="category-button" @click="selectedCategory = cat">
          <span class="category-letter">{{ cat }}</span>
          <span class="category-name">{{ getCategoryName(cat) }}</span>
        </el-button>
      </div>
    </div>

    <transition name="fade">
      <div v-if="selectedCategory" class="subcategory-section">
        <h2>{{ selectedCategory }} 类子分类</h2>
        <el-scrollbar height="400px">
          <el-button v-for="item in filteredSubcategories" :key="item.code"
                     class="subcategory-button" @click="goToPage(item.code)">
            <span class="code">{{ item.code }}</span>
            <span class="name">{{ item.name }}</span>
          </el-button>
        </el-scrollbar>
      </div>
    </transition>
  </div>
</template>


<script>
export default {
  data() {
    return {
      selectedCategory: '',
      subcategories: {
        A: [
          { code: 'A01', name: '农业；林业；畜牧业；狩猎；诱捕；捕鱼' },
          { code: 'A21', name: '焙烤；制作或处理面团的设备；焙烤用面团[2006.01]' },
          { code: 'A22', name: '屠宰；肉品处理；家禽或鱼的加工' },
          { code: 'A23', name: '其他类不包含的食品或食料；及其处理' },
          { code: 'A24', name: '烟草；雪茄烟；纸烟；模拟吸烟装置；吸烟者用品' },
          { code: 'A41', name: '服装' },
          { code: 'A42', name: '帽类制品' },
          { code: 'A43', name: '鞋类' },
          { code: 'A44', name: '服饰缝纫用品；珠宝' },
          { code: 'A45', name: '手携物品或旅行品' },
          { code: 'A46', name: '刷类制品' },
          { code: 'A47', name: '家具；家庭用的物品或设备；咖啡磨；香料磨；一般吸尘器' },
          { code: 'A61', name: '医学或兽医学；卫生学' },
          { code: 'A62', name: '救生；消防' },
          { code: 'A63', name: '运动；游戏；娱乐活动' },
          { code: 'A99', name: '本部其他类目中不包括的技术主题[2006.01]' },
        ],
        B: [
          { code: 'B01', name: '一般的物理或化学的方法或装置' },
          { code: 'B02', name: '破碎、磨粉或粉碎；谷物碾磨的预处理' },
          { code: 'B03', name: '用液体或用风力摇床或风力跳汰机分离固体物料；从固体物料或流体中分离固体物料的磁或静电分离；高压电场分离[2006.01]' },
          { code: 'B04', name: '用于实现物理或化学工艺过程的离心装置或离心机' },
          { code: 'B05', name: '一般喷射或雾化；对表面涂覆流体的一般方法[2006.01]' },
          { code: 'B06', name: '一般机械振动的发生或传递' },
          { code: 'B07', name: '将固体从固体中分离；分选' },
          { code: 'B08', name: '清洁' },
          { code: 'B09', name: '固体废物的处理；被污染土壤的再生[2006.01]' },
          { code: 'B21', name: '基本上无切削的金属机械加工；金属冲压' },
          { code: 'B22', name: '铸造；粉末冶金' },
          { code: 'B23', name: '机床；其他类目中不包括的金属加工' },
          { code: 'B24', name: '磨削；抛光' },
          { code: 'B25', name: '手动工具；轻便机动工具；手动器械的手柄；车间设备；机械手' },
          { code: 'B26', name: '手动切割工具；切割；切断' },
          { code: 'B27', name: '木材或类似材料的加工或保存；一般钉钉机或钉U形钉机' },
          { code: 'B28', name: '加工水泥、黏土或石料' },
          { code: 'B29', name: '塑料的加工；一般处于塑性状态物质的加工' },
          { code: 'B30', name: '压力机' },
          { code: 'B31', name: '纸品或纸板或类似纸的方式加工的材料制品制作；纸或纸板或类似纸的方式加工的材料的加工' },
          { code: 'B32', name: '层状产品' },
          { code: 'B33', name: '增材制造技术[2015.01]' },
          { code: 'B41', name: '印刷；排版机；打字机；模印机[2006.01]' },
          { code: 'B42', name: '装订；图册；文件夹；特种印刷品' },
          { code: 'B43', name: '书写或绘图器具；办公用品' },
          { code: 'B44', name: '装饰艺术' },
          { code: 'B60', name: '一般车辆' },
          { code: 'B61', name: '铁路' },
          { code: 'B62', name: '无轨陆用车辆 B62B手动车辆，例如手推车或摇篮车；雪橇（以畜力驱动为特点的入B62C；由驾驶人或发动机推进的雪橇入B62M）' },
          { code: 'B63', name: '船舶或其他水上船只；与船有关的设备' },
          { code: 'B64', name: '飞行器；航空；宇宙航行' },
          { code: 'B65', name: '输送；包装；贮存；搬运薄的或细丝状材料' },
          { code: 'B66', name: '卷扬；提升；牵引' },
          { code: 'B67', name: '开启或封闭瓶子、罐或类似的容器；液体的贮运' },
          { code: 'B68', name: '鞍具；家具罩面' },
          { code: 'B81', name: '微观结构技术[7]' },
          { code: 'B82', name: '超微技术[7]' },
          { code: 'B99', name: '本部其他类目中不包括的技术主题[2006.01]' },
        ],
        C: [
          { code: 'C01', name: '无机化学' },
          { code: 'C02', name: '水、废水、污水或污泥的处理' },
          { code: 'C03', name: '玻璃；矿棉或渣棉' },
          { code: 'C04', name: '水泥；混凝土；人造石；陶瓷；耐火材料〔4〕' },
          { code: 'C05', name: '肥料；肥料制造〔4〕' },
          { code: 'C06', name: '炸药；火柴' },
          { code: 'C07', name: '有机化学〔2〕' },
          { code: 'C08', name: '有机高分子化合物；其制备或化学加工；以其为基料的组合物' },
          { code: 'C09', name: '染料；涂料；抛光剂；天然树脂；黏合剂；其他类目不包含的组合物；其他类目不包含的材料的应用' },
          { code: 'C10', name: '石油、煤气及炼焦工业；含一氧化碳的工业气体；燃料；润滑剂；泥煤' },
          { code: 'C11', name: '动物或植物油、脂、脂肪物质或蜡；由此制取的脂肪酸；洗涤剂；蜡烛' },
          { code: 'C12', name: '生物化学；啤酒；烈性酒；果汁酒；醋；微生物学；酶学；突变或遗传工程' },
          { code: 'C13', name: '糖工业〔4〕' },
          { code: 'C14', name: '小原皮；大原皮；毛皮或皮革' },
          { code: 'C21', name: '铁的冶金' },
          { code: 'C22', name: '冶金；黑色或有色金属合金；合金或有色金属的处理' },
          { code: 'C23', name: '对金属材料的镀覆；用金属材料对材料的镀覆；表面化学处理；金属材料的扩散处理；真空蒸发法、溅射法、离子注入法或化学气相沉积法的一般镀覆；金属材料腐蚀或积垢的一般抑制' },
          { code: 'C25', name: '电解或电泳工艺；其所用设备〔4〕' },
          { code: 'C30', name: '晶体生长[2006.01]' },
          { code: 'C40', name: '组合技术[2006.01]' },
          { code: 'C99', name: '本部其他类目不包括的技术主题[2006.01]' },
        ],
        D: [
          { code: 'D01', name: '天然或化学的线或纤维；纺纱或纺丝' },
          { code: 'D02', name: '纱线；纱线或绳索的机械整理；整经或络经' },
          { code: 'D03', name: '织造' },
          { code: 'D04', name: '编织；花边制作；针织；饰带；非织造布' },
          { code: 'D05', name: '缝纫；绣花；簇绒' },
          { code: 'D06', name: '织物等的处理；洗涤；其他类不包括的柔性材料' },
          { code: 'D07', name: '绳；除电缆以外的缆索' },
          { code: 'D21', name: '造纸；纤维素的生产' },
          { code: 'D99', name: '本部其他类目不包括的技术主题[2006.01]' },
        ],
        E: [
          { code: 'E01', name: '道路、铁路或桥梁的建筑' },
          { code: 'E02', name: '水利工程；基础；疏浚' },
          { code: 'E03', name: '给水；排水' },
          { code: 'E04', name: '建筑物' },
          { code: 'E05', name: '锁；钥匙；门窗零件；保险箱' },
          { code: 'E06', name: '一般门、窗、百叶窗或卷辊遮帘；梯子' },
          { code: 'E21', name: '土层或岩石的钻进；采矿' },
          { code: 'E99', name: '本部其他类目不包括的技术主题[2006.01]' },
        ],
        F: [
          { code: 'F01', name: '一般机器或发动机；一般的发动机装置；蒸汽机' },
          { code: 'F02', name: '燃烧发动机；热气或燃烧生成物的发动机装置' },
          { code: 'F03', name: '液力机械或液力发动机；风力、弹力或重力发动机；其他类目中不包括的产生机械动力或反推力的发动机' },
          { code: 'F04', name: '液体变容式机械；液体泵或弹性流体泵' },
          { code: 'F15', name: '流体压力执行机构；一般液压技术和气动技术' },
          { code: 'F16', name: '工程元件或部件；为产生和保持机器或设备的有效运行的一般措施；一般绝热' },
          { code: 'F17', name: '气体或液体的贮存或分配' },
          { code: 'F21', name: '照明' },
          { code: 'F22', name: '蒸汽的发生' },
          { code: 'F23', name: '燃烧设备；燃烧方法' },
          { code: 'F24', name: '供热；炉灶；通风' },
          { code: 'F25', name: '制冷或冷却；加热和制冷的联合系统；热泵系统；冰的制造或储存；气体的液化或固化' },
          { code: 'F26', name: '干燥' },
          { code: 'F27', name: '炉；窑；烘烤炉；蒸馏炉〔4〕' },
          { code: 'F28', name: '一般热交换' },
          { code: 'F41', name: '武器' },
          { code: 'F42', name: '弹药；爆破' },
          { code: 'F99', name: '本部其他类目不包括的技术主题[2006.01]' },
        ],
        G: [
          { code: 'G01', name: '测量；测试' },
          { code: 'G02', name: '光学' },
          { code: 'G03', name: '摄影术；电影术；利用了光波以外其他波的类似技术；电记录术；全息摄影术〔4〕' },
          { code: 'G04', name: '测时学' },
          { code: 'G05', name: '控制；调节' },
          { code: 'G06', name: '计算；推算或计数' },
          { code: 'G07', name: '核算装置' },
          { code: 'G08', name: '信号装置' },
          { code: 'G09', name: '教育；密码术；显示；广告；印鉴' },
          { code: 'G10', name: '乐器；声学' },
          { code: 'G11', name: '信息存储' },
          { code: 'G12', name: '仪器的结构零部件，或未列入其他类目的其他设备的类似零部件' },
          { code: 'G16', name: '特别适用于特定应用领域的信息通信技术[ICT]' },
          { code: 'G21', name: '核物理；核工程' },
          { code: 'G99', name: '不包含在本部其他类目中的技术主题[2006.01]' },        ],
        H: [
          { code: 'H01', name: '电气元件' },
          { code: 'H02', name: '发电、变电或配电' },
          { code: 'H03', name: '电子电路' },
          { code: 'H04', name: '电通信技术' },
          { code: 'H05', name: '其他类目不包含的电技术' },
          { code: 'H10', name: '半导体器件；其他类目中不包括的电固体器件[2023.01]' },
          { code: 'H99', name: '本部中其他类目不包括的技术主题[2006.01]' },        ]
      }
    };

  },
  computed: {
    filteredSubcategories() {
      return this.selectedCategory ? this.subcategories[this.selectedCategory] : [];
    }
  },
  methods: {
    getCategoryName(category) {
      const names = {
        'A': '生活必需品', 'B': '运输', 'C': '化学、冶金', 'D': '纺织、造纸',
        'E': '固定建筑物', 'F': '机械工程、照明、加热、武器、爆破', 'G': '物理', 'H': '电学'
      };
      return names[category] || '';
    },
    fuzzySearch() {
      this.$router.push('/fuzzy-search-page');
    },
    goToPage(pageId) {
      // 在此处添加跳转逻辑，例如：
      this.$router.push(`/IPC/${pageId}`);
    }
  }
};
</script>

<style scoped>
.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 40px 20px;
  background-color: #ffffff;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
  border-radius: 20px;
  overflow: hidden;
}

.description {
  text-align: center;
  margin-bottom: 30px;
  font-size: 16px;
  color: #2c3e50;
  line-height: 1.6;
  padding: 25px;
  background: linear-gradient(145deg, #f3f4f6, #ffffff);
  border-radius: 15px;
  box-shadow: 10px 10px 20px #d1d9e6, -10px -10px 20px #ffffff;
  position: relative;
}

.description i {
  font-size: 24px;
  color: #3498db;
  margin-bottom: 10px;
}

.search-container {
  display: flex;
  justify-content: center;
  margin-bottom: 40px;
}

.search-button {
  padding: 15px 30px;
  font-size: 18px;
  border-radius: 50px;
  background: linear-gradient(145deg, #3498db, #2980b9);
  border: none;
  color: white;
  display: flex;
  align-items: center;
  transition: all 0.3s ease;
  box-shadow: 5px 5px 10px #d1d9e6, -5px -5px 10px #ffffff;
}

.search-button:hover {
  transform: translateY(-3px);
  box-shadow: 0 7px 14px rgba(52, 152, 219, 0.3);
}

.search-button i {
  margin-right: 10px;
}

.category-section {
  margin-bottom: 40px;
}

.category-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 20px;
  margin-bottom: 20px;
}

.category-button {
  height: 100px;
  border-radius: 15px;
  transition: all 0.3s ease;
  border: none;
  background: linear-gradient(145deg, #f3f4f6, #ffffff);
  box-shadow: 5px 5px 10px #d1d9e6, -5px -5px 10px #ffffff;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  overflow: hidden;
}

.category-button:hover {
  transform: translateY(-5px);
  box-shadow: 0 10px 20px rgba(52, 152, 219, 0.2);
}

.category-letter {
  font-size: 36px;
  font-weight: bold;
  color: #3498db;
  margin-bottom: 5px;
}

.category-name {
  font-size: 14px;
  color: #34495e;
}

.subcategory-section {
  background: linear-gradient(145deg, #f3f4f6, #ffffff);
  border-radius: 15px;
  padding: 30px;
  margin-top: 30px;
  box-shadow: 10px 10px 20px #d1d9e6, -10px -10px 20px #ffffff;
}

.subcategory-section h2 {
  font-size: 28px;
  color: #2c3e50;
  margin-bottom: 25px;
  text-align: center;
  text-shadow: 1px 1px 2px #ffffff;
}

.subcategory-button {
  display: flex;
  justify-content: flex-start;
  align-items: center;
  width: 100%;
  margin-bottom: 15px;
  padding: 15px 20px;
  text-align: left;
  border: none;
  border-radius: 10px;
  transition: all 0.3s ease;
  background: #ffffff;
  box-shadow: 3px 3px 6px #d1d9e6, -3px -3px 6px #ffffff;
}

.subcategory-button:hover {
  background: linear-gradient(145deg, #f3f4f6, #ffffff);
  transform: translateX(5px);
  box-shadow: 5px 5px 10px #d1d9e6, -5px -5px 10px #ffffff;
}

.subcategory-button .code {
  font-weight: bold;
  margin-right: 15px;
  color: #3498db;
  font-size: 18px;
}

.subcategory-button .name {
  color: #34495e;
  font-size: 16px;
}

.fade-enter-active, .fade-leave-active {
  transition: all 0.5s ease;
}
.fade-enter, .fade-leave-to {
  opacity: 0;
  transform: translateY(20px);
}

@media (max-width: 768px) {
  .category-grid {
    grid-template-columns: repeat(2, 1fr);
  }
  .category-button {
    height: 80px;
  }
  .category-letter {
    font-size: 28px;
  }
  .category-name {
    font-size: 12px;
  }
}
</style>
