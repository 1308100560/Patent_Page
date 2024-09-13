<template>
  <div>
    <el-page-header content="预览" @back="switchPage('index')" />

    <div class="ticket-preview">
      <h3>工单信息</h3>
      <div>
        <h4>{{ ticket.name }}</h4>
      </div>
      <h3>表单信息</h3>
      <div v-for="key in properties.keys" :key="key">
        <h4>{{ key.value }}</h4>
        <p style="font-size: 14px">{{ properties[key.key] }}</p>
      </div>
    </div>
    <h3>流程追踪</h3>
    <ticket-progress :switch-page="switchPage" :ticket="ticket" />
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
      properties: []
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
