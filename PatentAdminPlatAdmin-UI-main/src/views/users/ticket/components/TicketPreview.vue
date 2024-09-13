<template>
  <div>
    <el-page-header content="预览" @back="switchPage('index')" />
    <el-collapse v-model="activeNames" style="margin-top: 20px">
      <el-collapse-item name="1">
        <template slot="title">
          <div style="font-size: 20px">工单信息</div>
        </template>
        <div>
          <div>
            <h4>{{ ticket.name }}</h4>
          </div>
          <h3>表单信息</h3>
          <div v-for="key in properties.keys" :key="key.value">
            <h4>{{ key.value }}</h4>
            <p style="font-size: 14px">{{ properties[key.key] }}</p>
          </div>
        </div>
      </el-collapse-item>
      <el-collapse-item name="2">
        <template slot="title">
          <div style="font-size: 20px">流程追踪</div>
        </template>
        <div>
          <ticket-progress :switch-page="switchPage" :ticket="ticket" />
        </div>
      </el-collapse-item>
    </el-collapse>

  </div>
</template>
<script>
import TicketProgress from '@/views/users/ticket/components/TicketProgress.vue'

export default {
  name: 'TicketPreview',
  components: { TicketProgress },
  props: {
    switchPage: {
      type: Function,
      default: () => {
      }
    },

    ticket: {
      type: Object,
      default: () => {
      }
    }
  },
  data() {
    return {
      properties: [],
      activeNames: ['1', '2']
    }
  },
  watch: {
    ticket: {
      handler() {
        this.properties = JSON.parse(this.ticket.properties)
        console.log(this.properties)
      },
      immediate: true
    }
  },
  methods: {
    onSubmit() {
      this.$emit('submit')
    }
  }
}
</script>
<style scoped>
.ticket-preview {
  margin: 30px 0;
}
</style>
