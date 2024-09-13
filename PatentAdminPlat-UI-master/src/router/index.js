import Vue from 'vue'
import Router from 'vue-router'
/* Layout */
import Layout from '@/layout'

Vue.use(Router)

export const constantRoutes = [{
  path: '/redirect', component: Layout, hidden: true, children: [{
    path: '/redirect/:path(.*)', component: () => import('@/views/redirect/index')
  }]
}, {
  path: '/register', component: () => import('@/views/users/components/RegisterComponent'), hidden: true
}, {
  path: '/auth-redirect', component: () => import('@/views/login/auth-redirect'), hidden: true
}, {
  path: '/404', component: () => import('@/views/error-page/404'), hidden: true
}, {
  path: '/401', component: () => import('@/views/error-page/401'), hidden: true
}, {
  path: '/login', component: () => import('@/views/users/login')
}, {
  path: '/index', component: Layout, children: [{
    path: 'index',
    component: () => import('@/views/users/profile/index'),
    name: 'Profile',
    meta: { title: '个人资料', icon: 'profile' }
  }]
}, {
  path: '/',
  component: Layout,
  meta: { title: '首页', icon: 'el-icon-chat-dot-square' },
  children: [
    {
      path: '/',
      component: () => import('@/views/users/shouye'),
      name: 'Shouye',
      meta: { title: '首页', icon: 'profile' }
    },
    {
      path: 'New_innovation_analysis_report',
      component: () => import('@/views/users/New_innovation_analysis_report'),
      name: 'New_innovation_analysis_report',
      meta: { title: '新创性分析报告', icon: 'el-icon-document' }
    },
    {
      path: 'IPC',
      component: { render: (c) => c('router-view') },
      name: 'Patent_disclosure_document',
      meta: { title: '专利技术交底书', icon: 'el-icon-s-management' },
      children: [
        {
          path: '/',
          component: () => import('@/views/users/IPC/Patent_disclosure_document'),
          name: 'Patent_disclosure_overview',
          meta: { title: '专利技术交底书概览', icon: 'el-icon-document-copy' }
        },
        {
          path: '',
          name: 'IPC_A',
          meta: { title: '人类生活必需品_A', icon: 'el-icon-user' },
          component: { render: (c) => c('router-view') },
          children: [
            {
              path: 'A01',
              name: 'IPC_A01',
              component: () => import('@/views/users/IPC/A/A_01.vue'),
              meta: { title: '农业；林业；畜牧业；狩猎；诱捕；捕鱼', icon: 'el-icon-s-home' }
            },
            {
              path: 'A21',
              name: 'IPC_A21',
              component: () => import('@/views/users/IPC/A/A_21.vue'),
              meta: { title: '焙烤；制作或处理面团的设备；焙烤用面团', icon: 'el-icon-food' }
            },
            {
              path: 'A22',
              name: 'IPC_A22',
              component: () => import('@/views/users/IPC/A/A_22.vue'),
              meta: { title: '屠宰；肉品处理；家禽或鱼的加工', icon: 'el-icon-knife-fork' }
            },
            {
              path: 'A23',
              name: 'IPC_A23',
              component: () => import('@/views/users/IPC/A/A_23.vue'),
              meta: { title: '其他类不包含的食品或食料；及其处理', icon: 'el-icon-dish' }
            },
            {
              path: 'A24',
              name: 'IPC_A24',
              component: () => import('@/views/users/IPC/A/A_24.vue'),
              meta: { title: '烟草；雪茄烟；纸烟；模拟吸烟装置；吸烟者用品', icon: 'el-icon-smoking' }
            },
            {
              path: 'A41',
              name: 'IPC_A41',
              component: () => import('@/views/users/IPC/A/A_41.vue'),
              meta: { title: '服装', icon: 'el-icon-clothes' }
            },
            {
              path: 'A42',
              name: 'IPC_A42',
              component: () => import('@/views/users/IPC/A/A_42.vue'),
              meta: { title: '帽类制品', icon: 'el-icon-hat' }
            },
            {
              path: 'A43',
              name: 'IPC_A43',
              component: () => import('@/views/users/IPC/A/A_43.vue'),
              meta: { title: '鞋类', icon: 'el-icon-shoe' }
            },
            {
              path: 'A44',
              name: 'IPC_A44',
              component: () => import('@/views/users/IPC/A/A_44.vue'),
              meta: { title: '服饰缝纫用品；珠宝', icon: 'el-icon-gem' }
            },
            {
              path: 'A45',
              name: 'IPC_A45',
              component: () => import('@/views/users/IPC/A/A_45.vue'),
              meta: { title: '手携物品或旅行品', icon: 'el-icon-suitcase' }
            },
            {
              path: 'A46',
              name: 'IPC_A46',
              component: () => import('@/views/users/IPC/A/A_46.vue'),
              meta: { title: '刷类制品', icon: 'el-icon-brush' }
            },
            {
              path: 'A47',
              name: 'IPC_A47',
              component: () => import('@/views/users/IPC/A/A_47.vue'),
              meta: { title: '家具；家庭用的物品或设备；咖啡磨；香料磨；一般吸尘器', icon: 'el-icon-house' }
            },
            {
              path: 'A61',
              name: 'IPC_A61',
              component: () => import('@/views/users/IPC/A/A_61.vue'),
              meta: { title: '医学或兽医学；卫生学', icon: 'el-icon-first-aid-kit' }
            },
            {
              path: 'A62',
              name: 'IPC_A62',
              component: () => import('@/views/users/IPC/A/A_62.vue'),
              meta: { title: '救生；消防', icon: 'el-icon-warning' }
            },
            {
              path: 'A63',
              name: 'IPC_A63',
              component: () => import('@/views/users/IPC/A/A_63.vue'),
              meta: { title: '运动；游戏；娱乐活动', icon: 'el-icon-basketball' }
            },
            {
              path: 'A99',
              name: 'IPC_A99',
              component: () => import('@/views/users/IPC/A/A_99.vue'),
              meta: { title: '本部其他类目中不包括的技术主题', icon: 'el-icon-more' }
            }
          ]
        },
        {
          path: '',
          name: 'IPC_B',
          meta: { title: '作业；运输_B', icon: 'el-icon-truck' },
          component: { render: (c) => c('router-view') },
          children: [
            {
              path: 'B01',
              name: 'IPC_B01',
              component: () => import('@/views/users/IPC/B/B_01.vue'),
              meta: { title: '一般的物理或化学的方法或装置', icon: 'el-icon-s-home' }
            },
            {
              path: 'B02',
              name: 'IPC_B02',
              component: () => import('@/views/users/IPC/B/B_02.vue'),
              meta: { title: '破碎、磨粉或粉碎；谷物碾磨的预处理', icon: 'el-icon-food' }
            },
            {
              path: 'B03',
              name: 'IPC_B03',
              component: () => import('@/views/users/IPC/B/B_03.vue'),
              meta: { title: '用液体或用风力摇床或风力跳汰机分离固体物料；从固体物料或流体中分离固体物料的磁或静电分离；高压电场分离', icon: 'el-icon-knife-fork' }
            },
            {
              path: 'B04',
              name: 'IPC_B04',
              component: () => import('@/views/users/IPC/B/B_04.vue'),
              meta: { title: '用于实现物理或化学工艺过程的离心装置或离心机', icon: 'el-icon-dish' }
            },
            {
              path: 'B05',
              name: 'IPC_B05',
              component: () => import('@/views/users/IPC/B/B_05.vue'),
              meta: { title: '一般喷射或雾化；对表面涂覆流体的一般方法', icon: 'el-icon-smoking' }
            },
            {
              path: 'B06',
              name: 'IPC_B06',
              component: () => import('@/views/users/IPC/B/B_06.vue'),
              meta: { title: '一般机械振动的发生或传递', icon: 'el-icon-clothes' }
            },
            {
              path: 'B07',
              name: 'IPC_B07',
              component: () => import('@/views/users/IPC/B/B_07.vue'),
              meta: { title: '将固体从固体中分离；分选', icon: 'el-icon-hat' }
            },
            {
              path: 'B08',
              name: 'IPC_B08',
              component: () => import('@/views/users/IPC/B/B_08.vue'),
              meta: { title: '清洁', icon: 'el-icon-shoe' }
            },
            {
              path: 'B09',
              name: 'IPC_B09',
              component: () => import('@/views/users/IPC/B/B_09.vue'),
              meta: { title: '固体废物的处理；被污染土壤的再生', icon: 'el-icon-gem' }
            },
            {
              path: 'B21',
              name: 'IPC_B21',
              component: () => import('@/views/users/IPC/B/B_21.vue'),
              meta: { title: '基本上无切削的金属机械加工；金属冲压', icon: 'el-icon-suitcase' }
            },
            {
              path: 'B22',
              name: 'IPC_B22',
              component: () => import('@/views/users/IPC/B/B_22.vue'),
              meta: { title: '铸造；粉末冶金', icon: 'el-icon-brush' }
            },
            {
              path: 'B23',
              name: 'IPC_B23',
              component: () => import('@/views/users/IPC/B/B_23.vue'),
              meta: { title: '机床；其他类目中不包括的金属加工', icon: 'el-icon-house' }
            },
            {
              path: 'B24',
              name: 'IPC_B24',
              component: () => import('@/views/users/IPC/B/B_24.vue'),
              meta: { title: '磨削；抛光', icon: 'el-icon-first-aid-kit' }
            },
            {
              path: 'B25',
              name: 'IPC_B25',
              component: () => import('@/views/users/IPC/B/B_25.vue'),
              meta: { title: '手动工具；轻便机动工具；手动器械的手柄；车间设备；机械手', icon: 'el-icon-warning' }
            },
            {
              path: 'B26',
              name: 'IPC_B26',
              component: () => import('@/views/users/IPC/B/B_26.vue'),
              meta: { title: '手动切割工具；切割；切断', icon: 'el-icon-basketball' }
            },
            {
              path: 'B27',
              name: 'IPC_B27',
              component: () => import('@/views/users/IPC/B/B_27.vue'),
              meta: { title: '木材或类似材料的加工或保存；一般钉钉机或钉U形钉机', icon: 'el-icon-more' }
            },
            {
              path: 'B28',
              name: 'IPC_B28',
              component: () => import('@/views/users/IPC/B/B_28.vue'),
              meta: { title: '加工水泥、黏土或石料', icon: 'el-icon-more' }
            },
            {
              path: 'B29',
              name: 'IPC_B29',
              component: () => import('@/views/users/IPC/B/B_29.vue'),
              meta: { title: '塑料的加工；一般处于塑性状态物质的加工', icon: 'el-icon-more' }
            },
            {
              path: 'B30',
              name: 'IPC_B30',
              component: () => import('@/views/users/IPC/B/B_30.vue'),
              meta: { title: '压力机', icon: 'el-icon-more' }
            },
            {
              path: 'B31',
              name: 'IPC_B31',
              component: () => import('@/views/users/IPC/B/B_31.vue'),
              meta: { title: '纸品或纸板或类似纸的方式加工的材料制品制作；纸或纸板或类似纸的方式加工的材料的加工', icon: 'el-icon-more' }
            },
            {
              path: 'B32',
              name: 'IPC_B32',
              component: () => import('@/views/users/IPC/B/B_32.vue'),
              meta: { title: '层状产品', icon: 'el-icon-more' }
            },
            {
              path: 'B33',
              name: 'IPC_B33',
              component: () => import('@/views/users/IPC/B/B_33.vue'),
              meta: { title: '增材制造技术', icon: 'el-icon-more' }
            },
            {
              path: 'B41',
              name: 'IPC_B41',
              component: () => import('@/views/users/IPC/B/B_41.vue'),
              meta: { title: '印刷；排版机；打字机；模印机', icon: 'el-icon-more' }
            },
            {
              path: 'B42',
              name: 'IPC_B42',
              component: () => import('@/views/users/IPC/B/B_42.vue'),
              meta: { title: '装订；图册；文件夹；特种印刷品', icon: 'el-icon-more' }
            },
            {
              path: 'B43',
              name: 'IPC_B43',
              component: () => import('@/views/users/IPC/B/B_43.vue'),
              meta: { title: '书写或绘图器具；办公用品', icon: 'el-icon-more' }
            },
            {
              path: 'B44',
              name: 'IPC_B44',
              component: () => import('@/views/users/IPC/B/B_44.vue'),
              meta: { title: '装饰艺术', icon: 'el-icon-more' }
            },
            {
              path: 'B60',
              name: 'IPC_B60',
              component: () => import('@/views/users/IPC/B/B_60.vue'),
              meta: { title: '一般车辆', icon: 'el-icon-more' }
            },
            {
              path: 'B61',
              name: 'IPC_B61',
              component: () => import('@/views/users/IPC/B/B_61.vue'),
              meta: { title: '铁路', icon: 'el-icon-more' }
            },
            {
              path: 'B62',
              name: 'IPC_B62',
              component: () => import('@/views/users/IPC/B/B_62.vue'),
              meta: { title: '无轨陆用车辆', icon: 'el-icon-more' }
            },
            {
              path: 'B63',
              name: 'IPC_B63',
              component: () => import('@/views/users/IPC/B/B_63.vue'),
              meta: { title: '船舶或其他水上船只；与船有关的设备', icon: 'el-icon-more' }
            },
            {
              path: 'B64',
              name: 'IPC_B64',
              component: () => import('@/views/users/IPC/B/B_64.vue'),
              meta: { title: '飞行器；航空；宇宙航行', icon: 'el-icon-more' }
            },
            {
              path: 'B65',
              name: 'IPC_B65',
              component: () => import('@/views/users/IPC/B/B_65.vue'),
              meta: { title: '输送；包装；贮存；搬运薄的或细丝状材料', icon: 'el-icon-more' }
            },
            {
              path: 'B66',
              name: 'IPC_B66',
              component: () => import('@/views/users/IPC/B/B_66.vue'),
              meta: { title: '卷扬；提升；牵引', icon: 'el-icon-more' }
            },
            {
              path: 'B67',
              name: 'IPC_B67',
              component: () => import('@/views/users/IPC/B/B_67.vue'),
              meta: { title: '开启或封闭瓶子、罐或类似的容器；液体的贮运', icon: 'el-icon-more' }
            },
            {
              path: 'B68',
              name: 'IPC_B68',
              component: () => import('@/views/users/IPC/B/B_68.vue'),
              meta: { title: '鞍具；家具罩面', icon: 'el-icon-more' }
            },
            {
              path: 'B81',
              name: 'IPC_B81',
              component: () => import('@/views/users/IPC/B/B_81.vue'),
              meta: { title: '微观结构技术', icon: 'el-icon-more' }
            },
            {
              path: 'B82',
              name: 'IPC_B82',
              component: () => import('@/views/users/IPC/B/B_82.vue'),
              meta: { title: '超微技术', icon: 'el-icon-more' }
            },
            {
              path: 'B99',
              name: 'IPC_B99',
              component: () => import('@/views/users/IPC/B/B_99.vue'),
              meta: { title: '本部其他类目中不包括的技术主题', icon: 'el-icon-more' }
            }
          ]
        },
        {
          path: '',
          name: 'IPC_C',
          meta: { title: '化学；冶金_C', icon: 'el-icon-test-tube' },
          component: { render: (c) => c('router-view') },
          children: [
            {
              path: 'C01',
              name: 'IPC_C01',
              component: () => import('@/views/users/IPC/C/C_01.vue'),
              meta: { title: '无机化学', icon: 'el-icon-s-home' }
            },
            {
              path: 'C02',
              name: 'IPC_C02',
              component: () => import('@/views/users/IPC/C/C_02.vue'),
              meta: { title: '水、废水、污水或污泥的处理', icon: 'el-icon-water' }
            },
            {
              path: 'C03',
              name: 'IPC_C03',
              component: () => import('@/views/users/IPC/C/C_03.vue'),
              meta: { title: '玻璃；矿棉或渣棉', icon: 'el-icon-glass' }
            },
            {
              path: 'C04',
              name: 'IPC_C04',
              component: () => import('@/views/users/IPC/C/C_04.vue'),
              meta: { title: '水泥；混凝土；人造石；陶瓷；耐火材料', icon: 'el-icon-cement' }
            },
            {
              path: 'C05',
              name: 'IPC_C05',
              component: () => import('@/views/users/IPC/C/C_05.vue'),
              meta: { title: '肥料；肥料制造', icon: 'el-icon-leaf' }
            },
            {
              path: 'C06',
              name: 'IPC_C06',
              component: () => import('@/views/users/IPC/C/C_06.vue'),
              meta: { title: '炸药；火柴', icon: 'el-icon-bomb' }
            },
            {
              path: 'C07',
              name: 'IPC_C07',
              component: () => import('@/views/users/IPC/C/C_07.vue'),
              meta: { title: '有机化学', icon: 'el-icon-bottle' }
            },
            {
              path: 'C08',
              name: 'IPC_C08',
              component: () => import('@/views/users/IPC/C/C_08.vue'),
              meta: { title: '有机高分子化合物；其制备或化学加工；以其为基料的组合物', icon: 'el-icon-sugar' }
            },
            {
              path: 'C09',
              name: 'IPC_C09',
              component: () => import('@/views/users/IPC/C/C_09.vue'),
              meta: { title: '染料；涂料；抛光剂；天然树脂；黏合剂；其他类目不包含的组合物；其他类目不包含的材料的应用', icon: 'el-icon-paint' }
            },
            {
              path: 'C10',
              name: 'IPC_C10',
              component: () => import('@/views/users/IPC/C/C_10.vue'),
              meta: { title: '石油、煤气及炼焦工业；含一氧化碳的工业气体；燃料；润滑剂；泥煤', icon: 'el-icon-oil' }
            },
            {
              path: 'C11',
              name: 'IPC_C11',
              component: () => import('@/views/users/IPC/C/C_11.vue'),
              meta: { title: '动物或植物油、脂、脂肪物质或蜡；由此制取的脂肪酸；洗涤剂；蜡烛', icon: 'el-icon-candle' }
            },
            {
              path: 'C12',
              name: 'IPC_C12',
              component: () => import('@/views/users/IPC/C/C_12.vue'),
              meta: { title: '生物化学；啤酒；烈性酒；果汁酒；醋；微生物学；酶学；突变或遗传工程', icon: 'el-icon-bio' }
            },
            {
              path: 'C13',
              name: 'IPC_C13',
              component: () => import('@/views/users/IPC/C/C_13.vue'),
              meta: { title: '糖工业', icon: 'el-icon-sugar' }
            },
            {
              path: 'C14',
              name: 'IPC_C14',
              component: () => import('@/views/users/IPC/C/C_14.vue'),
              meta: { title: '小原皮；大原皮；毛皮或皮革', icon: 'el-icon-leather' }
            },
            {
              path: 'C21',
              name: 'IPC_C21',
              component: () => import('@/views/users/IPC/C/C_21.vue'),
              meta: { title: '铁的冶金', icon: 'el-icon-metal' }
            },
            {
              path: 'C22',
              name: 'IPC_C22',
              component: () => import('@/views/users/IPC/C/C_22.vue'),
              meta: { title: '冶金；黑色或有色金属合金；合金或有色金属的处理', icon: 'el-icon-gold' }
            },
            {
              path: 'C23',
              name: 'IPC_C23',
              component: () => import('@/views/users/IPC/C/C_23.vue'),
              meta: { title: '对金属材料的镀覆；用金属材料对材料的镀覆；表面化学处理；金属材料的扩散处理；真空蒸发法、溅射法、离子注入法或化学气相沉积法的一般镀覆；金属材料腐蚀或积垢的一般抑制', icon: 'el-icon-coating' }
            },
            {
              path: 'C25',
              name: 'IPC_C25',
              component: () => import('@/views/users/IPC/C/C_25.vue'),
              meta: { title: '电解或电泳工艺；其所用设备', icon: 'el-icon-electrolysis' }
            },
            {
              path: 'C30',
              name: 'IPC_C30',
              component: () => import('@/views/users/IPC/C/C_30.vue'),
              meta: { title: '晶体生长', icon: 'el-icon-crystal' }
            },
            {
              path: 'C40',
              name: 'IPC_C40',
              component: () => import('@/views/users/IPC/C/C_40.vue'),
              meta: { title: '组合技术', icon: 'el-icon-combination' }
            },
            {
              path: 'C99',
              name: 'IPC_C99',
              component: () => import('@/views/users/IPC/C/C_99.vue'),
              meta: { title: '本部其他类目不包括的技术主题', icon: 'el-icon-more' }
            }
          ]
        },
        {
          path: '',
          name: 'IPC_D',
          meta: { title: '纺织；造纸_D', icon: 'el-icon-scissors' },
          component: { render: (c) => c('router-view') },
          children: [
            {
              path: 'D01',
              name: 'IPC_D01',
              component: () => import('@/views/users/IPC/D/D_01.vue'),
              meta: { title: '天然或化学的线或纤维；纺纱或纺丝', icon: 'el-icon-thread' }
            },
            {
              path: 'D02',
              name: 'IPC_D02',
              component: () => import('@/views/users/IPC/D/D_02.vue'),
              meta: { title: '纱线；纱线或绳索的机械整理；整经或络经', icon: 'el-icon-yarn' }
            },
            {
              path: 'D03',
              name: 'IPC_D03',
              component: () => import('@/views/users/IPC/D/D_03.vue'),
              meta: { title: '织造', icon: 'el-icon-weaving' }
            },
            {
              path: 'D04',
              name: 'IPC_D04',
              component: () => import('@/views/users/IPC/D/D_04.vue'),
              meta: { title: '编织；花边制作；针织；饰带；非织造布', icon: 'el-icon-lace' }
            },
            {
              path: 'D05',
              name: 'IPC_D05',
              component: () => import('@/views/users/IPC/D/D_05.vue'),
              meta: { title: '缝纫；绣花；簇绒', icon: 'el-icon-sewing' }
            },
            {
              path: 'D06',
              name: 'IPC_D06',
              component: () => import('@/views/users/IPC/D/D_06.vue'),
              meta: { title: '织物等的处理；洗涤；其他类不包括的柔性材料', icon: 'el-icon-wash' }
            },
            {
              path: 'D07',
              name: 'IPC_D07',
              component: () => import('@/views/users/IPC/D/D_07.vue'),
              meta: { title: '绳；除电缆以外的缆索', icon: 'el-icon-rope' }
            },
            {
              path: 'D21',
              name: 'IPC_D21',
              component: () => import('@/views/users/IPC/D/D_21.vue'),
              meta: { title: '造纸；纤维素的生产', icon: 'el-icon-paper' }
            },
            {
              path: 'D99',
              name: 'IPC_D99',
              component: () => import('@/views/users/IPC/D/D_99.vue'),
              meta: { title: '本部其他类目不包括的技术主题', icon: 'el-icon-more' }
            },
          ]
        },
        {
          path: '',
          name: 'IPC_E',
          meta: { title: '固定建筑物_E', icon: 'el-icon-office-building' },
          component: { render: (c) => c('router-view') },
          children: [
            {
              path: 'E01',
              name: 'IPC_E01',
              component: () => import('@/views/users/IPC/E/E_01.vue'),
              meta: { title: '道路、铁路或桥梁的建筑', icon: 'el-icon-road' }
            },
            {
              path: 'E02',
              name: 'IPC_E02',
              component: () => import('@/views/users/IPC/E/E_02.vue'),
              meta: { title: '水利工程；基础；疏浚', icon: 'el-icon-water' }
            },
            {
              path: 'E03',
              name: 'IPC_E03',
              component: () => import('@/views/users/IPC/E/E_03.vue'),
              meta: { title: '给水；排水', icon: 'el-icon-pipe' }
            },
            {
              path: 'E04',
              name: 'IPC_E04',
              component: () => import('@/views/users/IPC/E/E_04.vue'),
              meta: { title: '建筑物', icon: 'el-icon-building' }
            },
            {
              path: 'E05',
              name: 'IPC_E05',
              component: () => import('@/views/users/IPC/E/E_05.vue'),
              meta: { title: '锁；钥匙；门窗零件；保险箱', icon: 'el-icon-lock' }
            },
            {
              path: 'E06',
              name: 'IPC_E06',
              component: () => import('@/views/users/IPC/E/E_06.vue'),
              meta: { title: '一般门、窗、百叶窗或卷辊遮帘；梯子', icon: 'el-icon-window' }
            },
            {
              path: 'E21',
              name: 'IPC_E21',
              component: () => import('@/views/users/IPC/E/E_21.vue'),
              meta: { title: '土层或岩石的钻进；采矿', icon: 'el-icon-drill' }
            },
            {
              path: 'E99',
              name: 'IPC_E99',
              component: () => import('@/views/users/IPC/E/E_99.vue'),
              meta: { title: '本部其他类目不包括的技术主题', icon: 'el-icon-more' }
            }
          ]
        },
        {
          path: '',
          name: 'IPC_F',
          meta: { title: '机械工程；照明；加热；武器；爆破_F', icon: 'el-icon-cpu' },
          component: { render: (c) => c('router-view') },
          children: [
            {
              path: 'F01',
              name: 'IPC_F01',
              component: () => import('@/views/users/IPC/F/F_01.vue'),
              meta: { title: '一般机器或发动机；一般的发动机装置；蒸汽机', icon: 'el-icon-engine' }
            },
            {
              path: 'F02',
              name: 'IPC_F02',
              component: () => import('@/views/users/IPC/F/F_02.vue'),
              meta: { title: '燃烧发动机；热气或燃烧生成物的发动机装置', icon: 'el-icon-burn' }
            },
            {
              path: 'F03',
              name: 'IPC_F03',
              component: () => import('@/views/users/IPC/F/F_03.vue'),
              meta: { title: '液力机械或液力发动机；风力、弹力或重力发动机；其他类目中不包括的产生机械动力或反推力的发动机', icon: 'el-icon-fluid' }
            },
            {
              path: 'F04',
              name: 'IPC_F04',
              component: () => import('@/views/users/IPC/F/F_04.vue'),
              meta: { title: '液体变容式机械；液体泵或弹性流体泵', icon: 'el-icon-pump' }
            },
            {
              path: 'F15',
              name: 'IPC_F15',
              component: () => import('@/views/users/IPC/F/F_15.vue'),
              meta: { title: '流体压力执行机构；一般液压技术和气动技术', icon: 'el-icon-pressure' }
            },
            {
              path: 'F16',
              name: 'IPC_F16',
              component: () => import('@/views/users/IPC/F/F_16.vue'),
              meta: { title: '工程元件或部件；为产生和保持机器或设备的有效运行的一般措施；一般绝热', icon: 'el-icon-component' }
            },
            {
              path: 'F17',
              name: 'IPC_F17',
              component: () => import('@/views/users/IPC/F/F_17.vue'),
              meta: { title: '气体或液体的贮存或分配', icon: 'el-icon-storage' }
            },
            {
              path: 'F21',
              name: 'IPC_F21',
              component: () => import('@/views/users/IPC/F/F_21.vue'),
              meta: { title: '照明', icon: 'el-icon-lightbulb' }
            },
            {
              path: 'F22',
              name: 'IPC_F22',
              component: () => import('@/views/users/IPC/F/F_22.vue'),
              meta: { title: '蒸汽的发生', icon: 'el-icon-steam' }
            },
            {
              path: 'F23',
              name: 'IPC_F23',
              component: () => import('@/views/users/IPC/F/F_23.vue'),
              meta: { title: '燃烧设备；燃烧方法', icon: 'el-icon-flame' }
            },
            {
              path: 'F24',
              name: 'IPC_F24',
              component: () => import('@/views/users/IPC/F/F_24.vue'),
              meta: { title: '供热；炉灶；通风', icon: 'el-icon-thermometer' }
            },
            {
              path: 'F25',
              name: 'IPC_F25',
              component: () => import('@/views/users/IPC/F/F_25.vue'),
              meta: { title: '制冷或冷却；加热和制冷的联合系统；热泵系统；冰的制造或储存；气体的液化或固化', icon: 'el-icon-cool' }
            },
            {
              path: 'F26',
              name: 'IPC_F26',
              component: () => import('@/views/users/IPC/F/F_26.vue'),
              meta: { title: '干燥', icon: 'el-icon-dry' }
            },
            {
              path: 'F27',
              name: 'IPC_F27',
              component: () => import('@/views/users/IPC/F/F_27.vue'),
              meta: { title: '炉；窑；烘烤炉；蒸馏炉', icon: 'el-icon-oven' }
            },
            {
              path: 'F28',
              name: 'IPC_F28',
              component: () => import('@/views/users/IPC/F/F_28.vue'),
              meta: { title: '一般热交换', icon: 'el-icon-heat' }
            },
            {
              path: 'F41',
              name: 'IPC_F41',
              component: () => import('@/views/users/IPC/F/F_41.vue'),
              meta: { title: '武器', icon: 'el-icon-weapon' }
            },
            {
              path: 'F42',
              name: 'IPC_F42',
              component: () => import('@/views/users/IPC/F/F_42.vue'),
              meta: { title: '弹药；爆破', icon: 'el-icon-ammo' }
            },
            {
              path: 'F99',
              name: 'IPC_F99',
              component: () => import('@/views/users/IPC/F/F_99.vue'),
              meta: { title: '本部其他类目不包括的技术主题', icon: 'el-icon-more' }
            }
          ]
        },
        {
          path: '',
          name: 'IPC_G',
          meta: { title: '物理_G', icon: 'el-icon-atom' },
          component: { render: (c) => c('router-view') },
          children: [
            {
              path: 'G01',
              name: 'IPC_G01',
              component: () => import('@/views/users/IPC/G/G_01.vue'),
              meta: { title: '测量；测试', icon: 'el-icon-measure' }
            },
            {
              path: 'G02',
              name: 'IPC_G02',
              component: () => import('@/views/users/IPC/G/G_02.vue'),
              meta: { title: '光学', icon: 'el-icon-optics' }
            },
            {
              path: 'G03',
              name: 'IPC_G03',
              component: () => import('@/views/users/IPC/G/G_03.vue'),
              meta: { title: '摄影术；电影术；利用了光波以外其他波的类似技术；电记录术；全息摄影术', icon: 'el-icon-photo' }
            },
            {
              path: 'G04',
              name: 'IPC_G04',
              component: () => import('@/views/users/IPC/G/G_04.vue'),
              meta: { title: '测时学', icon: 'el-icon-time' }
            },
            {
              path: 'G05',
              name: 'IPC_G05',
              component: () => import('@/views/users/IPC/G/G_05.vue'),
              meta: { title: '控制；调节', icon: 'el-icon-control' }
            },
            {
              path: 'G06',
              name: 'IPC_G06',
              component: () => import('@/views/users/IPC/G/G_06.vue'),
              meta: { title: '计算；推算或计数', icon: 'el-icon-calculator' }
            },
            {
              path: 'G07',
              name: 'IPC_G07',
              component: () => import('@/views/users/IPC/G/G_07.vue'),
              meta: { title: '核算装置', icon: 'el-icon-computer' }
            },
            {
              path: 'G08',
              name: 'IPC_G08',
              component: () => import('@/views/users/IPC/G/G_08.vue'),
              meta: { title: '信号装置', icon: 'el-icon-signal' }
            },
            {
              path: 'G09',
              name: 'IPC_G09',
              component: () => import('@/views/users/IPC/G/G_09.vue'),
              meta: { title: '教育；密码术；显示；广告；印鉴', icon: 'el-icon-education' }
            },
            {
              path: 'G10',
              name: 'IPC_G10',
              component: () => import('@/views/users/IPC/G/G_10.vue'),
              meta: { title: '乐器；声学', icon: 'el-icon-music' }
            },
            {
              path: 'G11',
              name: 'IPC_G11',
              component: () => import('@/views/users/IPC/G/G_11.vue'),
              meta: { title: '信息存储', icon: 'el-icon-storage' }
            },
            {
              path: 'G12',
              name: 'IPC_G12',
              component: () => import('@/views/users/IPC/G/G_12.vue'),
              meta: { title: '仪器的结构零部件，或未列入其他类目的其他设备的类似零部件', icon: 'el-icon-parts' }
            },
            {
              path: 'G16',
              name: 'IPC_G16',
              component: () => import('@/views/users/IPC/G/G_16.vue'),
              meta: { title: '特别适用于特定应用领域的信息通信技术[ICT]', icon: 'el-icon-ict' }
            },
            {
              path: 'G21',
              name: 'IPC_G21',
              component: () => import('@/views/users/IPC/G/G_21.vue'),
              meta: { title: '核物理；核工程', icon: 'el-icon-nuclear' }
            },
            {
              path: 'G99',
              name: 'IPC_G99',
              component: () => import('@/views/users/IPC/G/G_99.vue'),
              meta: { title: '不包含在本部其他类目中的技术主题', icon: 'el-icon-more' }
            }
          ]
        },
        {
          path: '',
          name: 'IPC_H',
          meta: { title: '电学_H', icon: 'el-icon-lightning' },
          component: { render: (c) => c('router-view') },
          children: [
            {
              path: 'H01',
              name: 'IPC_H01',
              component: () => import('@/views/users/IPC/H/H_01.vue'),
              meta: { title: '电气元件', icon: 'el-icon-electric' }
            },
            {
              path: 'H02',
              name: 'IPC_H02',
              component: () => import('@/views/users/IPC/H/H_02.vue'),
              meta: { title: '发电、变电或配电', icon: 'el-icon-power' }
            },
            {
              path: 'H03',
              name: 'IPC_H03',
              component: () => import('@/views/users/IPC/H/H_03.vue'),
              meta: { title: '电子电路', icon: 'el-icon-circuit' }
            },
            {
              path: 'H04',
              name: 'IPC_H04',
              component: () => import('@/views/users/IPC/H/H_04.vue'),
              meta: { title: '电通信技术', icon: 'el-icon-communication' }
            },
            {
              path: 'H05',
              name: 'IPC_H05',
              component: () => import('@/views/users/IPC/H/H_05.vue'),
              meta: { title: '其他类目不包含的电技术', icon: 'el-icon-electronics' }
            },
            {
              path: 'H10',
              name: 'IPC_H10',
              component: () => import('@/views/users/IPC/H/H_10.vue'),
              meta: { title: '半导体器件；其他类目中不包括的电固体器件', icon: 'el-icon-semiconductor' }
            },
            {
              path: 'H99',
              name: 'IPC_H99',
              component: () => import('@/views/users/IPC/H/H_99.vue'),
              meta: { title: '本部中其他类目不包括的技术主题', icon: 'el-icon-more' }
            }
          ]
        }
      ]
    }
  ]
}, {
  path: '/feed', component: Layout, meta: { title: '帮助与反馈', icon: 'el-icon-chat-dot-square' }, children: [
    {
      path: 'help',
      component: () => import('@/views/users/help/usage.vue'),
      name: 'Help',
      meta: { title: '使用说明', icon: 'el-icon-stopwatch' }
    },
    {
      path: 'feedback',
      component: () => import('@/views/users/help/feed.vue'),
      name: 'Feedback',
      meta: { title: '问题反馈', icon: 'el-icon-message' }
    }
  ]
}
]

/**
 * asyncRoutes
 * the routes that need to be dynamically loaded based on user roles
 */
export const asyncRoutes = []

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }), routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
