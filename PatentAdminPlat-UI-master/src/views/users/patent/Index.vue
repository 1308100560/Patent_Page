<template>
  <div class="container">
    <el-dialog :visible.sync="achievementsVisible" center title="添加科技成果" width="50%">
      <el-form :model="achievementsForm" :rules="achievementsFormRules">
        <el-form-item label="名称" prop="name">
          <el-input v-model="achievementsForm.name" placeholder="修改名称" />
        </el-form-item>
        <el-form-item label="内容" prop="content">
          <el-input v-model="achievementsForm.content" placeholder="修改内容" type="textarea" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="achievementsForm.description" placeholder="修改备注" type="textarea" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="achievementsVisible = false">取消</el-button>
        <el-button type="primary" @click="handleAddAchievements">确定</el-button>
      </div>
    </el-dialog>
    <div class="my-card" style="width: 100%;padding-top: 5px!important;">
      <el-tabs v-model="activeName">
        <el-tab-pane label="专利管理" name="first">
          <patent-claim-list :show-achievements="showAchievements" :switch-tab="switchTab" />
        </el-tab-pane>
        <el-tab-pane label="科技成果" name="second">
          <achievements
            ref="achievements"
            :show-achievement-edit-desc-form="showAchievementEditDescForm"
            :show-achievements="showAchievements"
            :switch-tab="switchTab"
          />
        </el-tab-pane>
      </el-tabs>
    </div>
  </div>

</template>
<script>
import Achievements from '@/views/users/patent/components/Achievements.vue'
import PatentClaimList from '@/views/users/patent/patentClaimList.vue'
import { createAchievement, updateAchievement } from '@/api/achievements'

export default {
  components: {
    PatentClaimList, Achievements
  },
  data() {
    return {
      achievementsVisible: false,
      activeName: 'first',
      achievementsForm: {
        name: '',
        content: '',
        description: ''
      },
      achievementsFormRules: {
        name: [
          { required: true, message: '请输入名称', trigger: 'blur' }
        ],
        content: [
          { required: true, message: '请输入内容', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    handleAddAchievements() {
      if (this.achievementsForm.name === '') {
        this.$message.error('请填写名称')
        return
      }
      if (this.achievementsForm.content === '') {
        this.$message.error('请填写内容')
        return
      }
      if (this.achievementsId) {
        updateAchievement(this.achievementsId, { properties: this.achievementsForm }).then(res => {
          this.$message.success('修改成功')
          this.achievementsVisible = false
          this.achievementsForm.name = ''
          this.achievementsForm.content = ''
          this.achievementsId = ''
          this.achievementsForm.description = ''
          this.$refs.achievements.getList()
        })
      } else {
        createAchievement({ properties: this.achievementsForm }).then(res => {
          this.$message.success('添加成功')
          this.achievementsVisible = false
          this.achievementsForm.name = ''
          this.achievementsForm.content = ''
          this.achievementsForm.description = ''
          this.$refs.achievements.getList()
          this.switchTab('second')
        })
      }
    },
    showAchievements() {
      this.achievementsVisible = true
    },
    showAchievementEditDescForm(id, achievement) {
      this.achievementsId = id
      this.achievementsForm = JSON.parse(JSON.stringify(achievement))
      this.achievementsVisible = true
    },
    switchTab(name) {
      this.activeName = name
    }
  }
}
</script>
