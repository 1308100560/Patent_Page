<template>
  <div>
    <header>
      <h3>xx项目</h3>
    </header>
    <div class="form-container">
      <div class="container">
        <form @submit.prevent="submitForm">
          <h2>发明名称</h2>
          <p>
            <input v-model="form.str1" type="text" placeholder="请在此处填写发明的名称" />
          </p>

          <h2>技术方案</h2>
          <p>
            <textarea v-model="form.str2" rows="4" cols="50" placeholder="描述发明的核心技术方案"></textarea>
          </p>

          <h2>产生的效果/解决的问题</h2>
          <p>
            <textarea v-model="form.str3" rows="4" cols="50" placeholder="解决的具体问题或
[列出第一个有益效果]
[列出第二个有益效果]
[根据需要添加更多有益效果]"></textarea>
          </p>

          <h2>具体实施例</h2>
          <p>
            <textarea v-model="form.str4" rows="4" cols="50" placeholder="描述一个或多个具体实施例，包括关键参数、操作条件等"></textarea>
          </p>

          <input type="submit" value="提交" />
        </form>
      </div>
    </div>
    <footer>
      <p>版权</p>
    </footer>
  </div>
</template>

<script>
export default {
  data() {
    return {
      form: {
        str1: '',
        str2: '',
        str3: '',
        str4: ''
      }
    };
  },
  methods: {
    async submitForm() {
      try {
        const response = await fetch('http://127.0.0.1:8002/submit-form', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
          },
          body: new URLSearchParams(this.form).toString()
        });

        console.log('Response status:', response.status);
        console.log('Response body:', await response.text()); // 输出响应内容

        if (response.ok) {
          alert('提交成功');
        } else {
          alert('提交失败');
        }
      } catch (error) {
        console.error('提交失败:', error);
        alert('提交过程中出现错误');
      }
    }
  }
};
</script>

<style scoped>
body {
  font-family: Arial, sans-serif;
  margin: 0;
  padding: 0;
  background-color: #f4f4f4;
}

header {
  background-color: #007BFF;
  color: white;
  padding: 10px 20px;
  text-align: center;
}

nav {
  background-color: #333;
  overflow: hidden;
}

nav a {
  float: left;
  color: white;
  text-align: center;
  padding: 14px 20px;
  text-decoration: none;
}

nav a:hover {
  background-color: #ddd;
  color: black;
}

.form-container {
  margin: auto;
  width: 70%;
  height: 900px;
  padding: 15px;
  border: 1px solid #ccc;
  border-radius: 5px;
  background-color: #f9f9f9;
}

.container {
  padding: 20px 90px;
  margin: 20px;
  height: 760px;
  background-color: white;
  border-radius: 8px;
  box-shadow: 0 0 10px rgba(0,0,0,0.1);
}

footer {
  background-color: #333;
  color: white;
  text-align: center;
  padding: 10px 0;
  position: fixed;
  left: 0;
  bottom: 0;
  width: 100%;
}

h1, h2 {
  color: #007BFF;
}

input[type="text"], textarea {
  width: 100%;
  border-radius: 5px;
  border: 1px solid #ccc;
  padding: 5px;
  font-size: 16px;
}

input[type="submit"] {
  background-color: #007BFF;
  color: white;
  border: none;
  padding: 10px 20px;
  border-radius: 5px;
  cursor: pointer;
}

input[type="submit"]:hover {
  background-color: #0056b3;
}
</style>
