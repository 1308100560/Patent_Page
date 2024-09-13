<template>
  <div class="container">
    <div class="top-right-content">
      <img src="IPC/srippm.jpg" alt="描述图片" class="top-right-image">
      <span class="top-right-text">srippm</span>
    </div>
    <h1 class="title">农业;林业;畜牧业;狩猎;诱捕;捕鱼</h1>
    <el-select v-model="selectedModel" placeholder="选择模型">
      <el-option
        v-for="model in availableModels"
        :key="model"
        :label="model"
        :value="model"
      />
    </el-select>
    <el-col :span="12">
      <el-card class="input-section" shadow="hover">
        <el-input
          v-model="inputText"
          type="textarea"
          :placeholder="Input_promp"
          :rows="10"
        />
      </el-card>
    </el-col>
    <div class="button-container">
      <el-button type="primary" @click="generatePrompt" :loading="isLoading">开始生成</el-button>
    </div>
    <el-col :span="12" v-if="outputText">
      <el-card class="output-section" shadow="hover">
        <div v-html="outputText"></div>
      </el-card>
    </el-col>
    <el-button v-if="outputText" type="success" @click="downloadDocument">下载文档</el-button>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      inputText: "",
      outputText: '',
      isLoading: false,
      docBytes: null,
      Input_promp: "发明名称：\n" +
        "[请在此处填写发明的名称]\n" +
        "技术方案或解决的问题：\n" +
        "[描述发明的核心技术方案或解决的问题]\n\n\n" +
        "产生的效果：\n" +
        "[列出第一个有益效果或第一个解决的问题]\n" +
        "[列出第二个有益效果或第二个解决的问题]\n" +
        "[根据需要添加更多有益效果或更多解决的问题]\n" +
        "具体实施例：\n" +
        "[描述一个或多个具体实施例，包括关键参数、操作条件等]\n",
      Generate_prompt: "你是一个专利专业人员，请阅读以上专利内容，详细回答我的任何问题，并且用中文回答我，我需要撰写一篇专利交底书，" +
        "请回答以下问题，确保内容详尽，清晰，并符合专利交底书的撰写要求，根据现有内容回答，且字数尽量长，专注于数据，不要说与问题无关的话，不要自己创造问题，" +
        "不要太宽泛，具体到细节领域，直接回答问题内容，以下是我的问题：\n" +
        "发明名称、\n" +
        "技术领域、\n" +
        "背景技术、\n" +
        "发明内容、\n" +
        "技术方案、\n" +
        "特点和有益效果、\n" +
        "附图说明(流程图，结构图)、\n" +
        "实施方式、\n" +
        "实施例（一个包含所有所知数据的流程）、\n" +
        "对比例、\n",
      selectedModel: 'llama2',
      availableModels: ['llama2', 'mistral', 'codellama'],  // 根据你的 ollama 可用模型进行调整
    };
  },
  methods: {
    async generatePrompt() {
      this.isLoading = true;
      try {
        const response = await axios.post('/generate_prompt', {
          input_text: this.inputText,
          generate_prompt: this.Generate_prompt,
          model: this.selectedModel
        });
        this.outputText = response.data.output;
        this.docBytes = response.data.doc_bytes;
      } catch (error) {
        console.error('Error generating prompt:', error);
        this.outputText = 'An error occurred while generating the prompt.';
      }
      this.isLoading = false;
    },
    downloadDocument() {
      if (this.docBytes) {
        const blob = new Blob([this.docBytes], { type: 'application/vnd.openxmlformats-officedocument.wordprocessingml.document' });
        const link = document.createElement('a');
        link.href = window.URL.createObjectURL(blob);
        link.download = 'patent_disclosure.docx';
        link.click();
      }
    }
  },
};
</script>

<style scoped>
.container {
  padding: 20px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  background-color: #b3d8ff;
  min-height: 100vh;
  box-sizing: border-box;
  position: relative;
}

.title {
  font-size: 30px;
  margin-bottom: 30px;
}

.top-right-content {
  position: absolute;
  top: 20px;
  right: 20px;
  display: flex;
  align-items: center;
  z-index: 10;
}

.top-right-image {
  width: 100px;
  height: 20px;
  margin-right: 10px;
}

.top-right-text {
  font-size: 16px;
  color: #333;
}

.content {
  width: 100%;
  max-width: 1800px;
  margin-bottom: 30px;
}

.input-section,
.output-section {
  padding: 20px;
  background-color: #e6f7ff;
  margin-bottom: 20px;
}

.button-container {
  text-align: center;
  margin-bottom: 20px;
}

::v-deep .el-textarea__inner {
  background-color: #f5f5f5 !important;
  border: 1px dashed #b3d8ff !important;
  box-sizing: border-box;
}
</style>
